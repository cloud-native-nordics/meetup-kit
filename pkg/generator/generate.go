package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/cloud-native-nordics/meetup-kit/pkg/types"
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/yaml"
)

// Options for the generator
// TODO: Make this better :)
type Options struct {
	// SpeakersFile points to the speakers.yaml file
	SpeakersFile string
	// CompaniesFile points to the companies.yaml file
	CompaniesFile string
	// RootDir points to the directory that has all meetup groups as subfolders, each with a meetup.yaml file
	RootDir string
	// DryRun controls whether to actually apply the changes or not
	DryRun bool
	// Validate controls whether to validate the current state of the repo content with the spec
	Validate bool
}

var unmarshal = yaml.UnmarshalStrict

// this maps the locations returned from meetup.com to what we want to use here.
// TODO: Maybe skip this and just use "Århus" directly in our
var cityNameExceptions = map[string]string{
	"Århus": "Aarhus",
}

func Generate(opts *Options) error {
	log.Debugf("generate: %v", *opts)
	cfg, err := load(opts.CompaniesFile, opts.SpeakersFile, opts.RootDir)
	if err != nil {
		return err
	}
	if err := update(cfg); err != nil {
		return err
	}
	out, err := exec(cfg)
	if err != nil {
		return err
	}
	if opts.Validate {
		return validate(out, opts.RootDir)
	}
	return apply(out, opts.RootDir, opts.DryRun)
}

func load(companiesPath, speakersPath, meetupsDir string) (*types.Config, error) {
	log.Debugf("load: %s %s %s", companiesPath, speakersPath, meetupsDir)
	companies := []types.Company{}
	companiesContent, err := ioutil.ReadFile(companiesPath)
	if err != nil {
		return nil, err
	}
	log.Debugf("%s", string(companiesContent))
	if err := unmarshal(companiesContent, &companies); err != nil {
		return nil, err
	}
	speakers := []types.Speaker{}
	speakersContent, err := ioutil.ReadFile(speakersPath)
	if err != nil {
		return nil, err
	}
	if err := unmarshal(speakersContent, &speakers); err != nil {
		return nil, err
	}
	meetupGroups := []types.MeetupGroup{}

	err = filepath.Walk(meetupsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return nil
		}
		// Consider only subdirectories of the root path
		if filepath.Dir(path) != "." {
			return nil
		}
		meetupsFile := filepath.Join(path, "meetup.yaml")
		if _, err := os.Stat(meetupsFile); os.IsNotExist(err) {
			return nil
		} else if err != nil {
			return err
		}
		mg := types.MeetupGroup{}
		mgContent, err := ioutil.ReadFile(meetupsFile)
		if err != nil {
			return err
		}
		if err := unmarshal(mgContent, &mg); err != nil {
			return err
		}
		meetupGroups = append(meetupGroups, mg)
		return nil
	})
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	wg.Add(len(meetupGroups))
	// Run the fetching from the meetup API in parallel for all meetup groups to speed things up
	for i := range meetupGroups {
		go func(mg *types.MeetupGroup) {
			defer wg.Done()
			mg.AutogenMeetupGroup, err = GetMeetupInfoFromAPI(*mg)
			if err != nil {
				log.Fatal(err)
			}
			mg.ApplyGeneratedData()
		}(&meetupGroups[i])
	}
	wg.Wait()

	return &types.Config{
		Speakers:     speakers,
		Companies:    companies,
		MeetupGroups: meetupGroups,
	}, nil
}

func apply(files map[string][]byte, rootDir string, dryRun bool) error {
	log.Debugf("apply: %v %s %t", files, rootDir, dryRun)
	for path, fileContent := range files {
		fullPath := filepath.Join(rootDir, path)
		if err := writeFile(fullPath, fileContent, dryRun); err != nil {
			return err
		}
	}
	return nil
}

func validate(files map[string][]byte, rootDir string) error {
	log.Debugf("validate: %v %s %t", files, rootDir)
	for path, fileContent := range files {
		fullPath := filepath.Join(rootDir, path)
		actual, err := ioutil.ReadFile(fullPath)
		if err != nil {
			return err
		}
		if !bytes.Equal(actual, fileContent) {
			return fmt.Errorf("%s differs from expected state. expected: \"%s\", actual: \"%s\"", fullPath, fileContent, actual)
		}
	}
	log.Info("Validation succeeded!")
	return nil
}

func tmpl(t *template.Template, obj interface{}) ([]byte, error) {
	log.Debugf("tmpl: %v", obj)
	var buf bytes.Buffer
	if err := t.Execute(&buf, obj); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func exec(cfg *types.Config) (map[string][]byte, error) {
	log.Debugf("exec: %v", *cfg)
	result := map[string][]byte{}
	types.ShouldMarshalAutoMeetup = false
	for _, mg := range cfg.MeetupGroups {
		mg.SetMeetupList()
		b, err := tmpl(readmeTmpl, mg)
		if err != nil {
			return nil, err
		}
		city := mg.CityLowercase()
		path := filepath.Join(city, "README.md")
		result[path] = b

		path = filepath.Join(city, "meetup.yaml")
		mg.AutogenMeetupGroup = nil
		meetupYAML, err := yaml.Marshal(mg)
		if err != nil {
			return nil, err
		}
		result[path] = meetupYAML
	}
	companiesYAML, err := yaml.Marshal(cfg.Companies)
	if err != nil {
		return nil, err
	}
	result["companies.yaml"] = companiesYAML
	speakersYAML, err := yaml.Marshal(cfg.Speakers)
	if err != nil {
		return nil, err
	}
	result["speakers.yaml"] = speakersYAML
	readmeBytes, err := tmpl(toplevelTmpl, cfg)
	if err != nil {
		return nil, err
	}
	result["README.md"] = readmeBytes
	types.ShouldMarshalAutoMeetup = true
	configJSON, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return nil, err
	}
	result["types.Config.json"] = configJSON
	stats, err := aggregateStats(cfg)
	if err != nil {
		return nil, err
	}
	statsJSON, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		return nil, err
	}
	result["stats.json"] = statsJSON
	return result, nil
}

func update(cfg *types.Config) error {
	for i := range cfg.MeetupGroups {
		mg := &cfg.MeetupGroups[i]

		calcSponsorTiers(mg)

		for j, m := range mg.Meetups {
			if err := setPresentationTimestamps(&m); err != nil {
				return err
			}
			mg.Meetups[j] = m
		}
	}
	return nil
}

func calcSponsorTiers(mg *types.MeetupGroup) {
	mg.SponsorTiers = map[types.CompanyID]types.SponsorTier{}
	for _, c := range mg.EcosystemMembers {
		if c.Company != nil {
			mg.SponsorTiers[c.ID] = types.SponsorTierEcosystemMember
		}
	}
	for _, m := range mg.Meetups {
		for _, p := range m.Presentations {
			for _, s := range p.Speakers {
				if s.Company.Company != nil {
					mg.SponsorTiers[s.Company.ID] = types.SponsorTierSpeakerProvider
				}
			}
		}
	}
	for _, o := range mg.Organizers {
		if o.Company.Company != nil {
			mg.SponsorTiers[o.Company.ID] = types.SponsorTierMeetup
		}
	}
	for _, m := range mg.Meetups {
		for _, s := range m.Sponsors {
			if s.Company.Company != nil {
				if s.Role == types.SponsorRoleLongterm {
					mg.SponsorTiers[s.Company.ID] = types.SponsorTierLongterm
				} else {
					mg.SponsorTiers[s.Company.ID] = types.SponsorTierMeetup
				}
			}
		}
	}
}

func writeFile(path string, b []byte, dryRun bool) error {
	if dryRun {
		fmt.Printf("Would write file %q with contents \"%s\"\n", path, string(b))
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	return ioutil.WriteFile(path, b, 0644)
}
