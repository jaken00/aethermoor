package world

import (
	"encoding/json"
	"os"
)

type rawNeedEntry struct {
	Resource    string  `json:"resource"`
	Kind        string  `json:"kind"`
	Current     float64 `json:"current"`
	Max         float64 `json:"max"`
	Threshold   float64 `json:"threshold"`
	ConsumeRate float64 `json:"consumeRate"`
}

type rawTemplateEntry struct {
	Produces     []ResourceEntry `json:"produces"`
	Needs        []rawNeedEntry  `json:"needs"`
	ShelterPrefs []string        `json:"shelterPrefs"`
	Aversions    []string        `json:"aversions"`
}

// EntityTemplate is the loaded template (not a live entity)
type EntityTemplate struct {
	TemplateName string
	Produces     []ResourceEntry
	Needs        map[NeedType]*NeedEntry
	ShelterPrefs []string
	Aversions    []AversionEntry
}

// LoadTemplates reads JSON and returns templates (not entities)
func LoadTemplates(path string) (map[string]EntityTemplate, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var raw map[string]rawTemplateEntry
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	out := make(map[string]EntityTemplate)
	for name, r := range raw {
		template := EntityTemplate{
			TemplateName: name,
			Produces:     r.Produces,
			ShelterPrefs: r.ShelterPrefs,
			Needs:        make(map[NeedType]*NeedEntry),
		}

		for _, rn := range r.Needs {
			need := &NeedEntry{
				Resource:    ResourceType(rn.Resource),
				Kind:        NeedType(rn.Kind),
				Current:     rn.Current,
				Max:         rn.Max,
				Threshold:   rn.Threshold,
				ConsumeRate: rn.ConsumeRate,
			}
			template.Needs[NeedType(rn.Kind)] = need
		}

		for _, a := range r.Aversions {
			template.Aversions = append(template.Aversions, AversionEntry{Resource: a})
		}

		out[name] = template
	}

	return out, nil
}

func SpawnEntityFromTemplate(template EntityTemplate, pos Vec2, id string) *Entity {
	entity := &Entity{
		Name:     id,
		Position: &Vec2{XPos: pos.XPos, YPos: pos.YPos},
		Alive:    true,
	}

	// Deep copy slices
	entity.Produces = make([]ResourceEntry, len(template.Produces))
	copy(entity.Produces, template.Produces)

	// Initialize Current to Max for new entities
	for i := range entity.Produces {
		entity.Produces[i].Current = entity.Produces[i].Max
	}

	entity.Needs = make(map[NeedType]*NeedEntry)
	for key, need := range template.Needs {
		needCopy := *need
		entity.Needs[key] = &needCopy
	}

	entity.ShelterPrefs = make([]string, len(template.ShelterPrefs))
	copy(entity.ShelterPrefs, template.ShelterPrefs)

	entity.Aversions = make([]AversionEntry, len(template.Aversions))
	copy(entity.Aversions, template.Aversions)

	return entity
}
