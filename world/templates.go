package world

import (
    "encoding/json"
    "os"
)

type rawNeedEntry struct {
    Resource    string  `json:"resource"`
    Threshhold  float64 `json:"threshhold"`
    Capacity    float64 `json:"capacity"`
    ConsumeRate float64 `json:"ConsumeRate"`
    MinInterest float64 `json:"MinInterest"`
}

type rawTemplateEntry struct {
    Produces     []ResourceEntry `json:"produces"`
    Needs        []rawNeedEntry  `json:"needs"`
    ShelterPrefs []string        `json:"shelterPrefs"`
    Aversions    []string        `json:"aversions"`
}

// LoadTemplates reads a JSON file at path and returns a map of template name -> Entity
func LoadTemplates(path string) (map[string]Entity, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }

    var raw map[string]rawTemplateEntry
    if err := json.Unmarshal(data, &raw); err != nil {
        return nil, err
    }

    out := make(map[string]Entity)
    for name, r := range raw {
        e := Entity{
            Name:         name,
            Alive:        true,
            Produces:     r.Produces,
            ShelterPrefs: r.ShelterPrefs,
        }

        for _, rn := range r.Needs {
            need := NeedEntry{
                Resource:    rn.Resource,
                Threshold:   rn.Threshhold,
                Capacity:    rn.Capacity,
                ConsumeRate: rn.ConsumeRate,
                MinInterest: rn.MinInterest,
            }
            e.Needs = append(e.Needs, need)
        }

        for _, a := range r.Aversions {
            e.Aversions = append(e.Aversions, AversionEntry{Resource: a})
        }

        out[name] = e
    }

    return out, nil
}
