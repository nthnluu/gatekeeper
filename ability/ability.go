package ability

type Rule struct {
	Action  string
	Subject string
}

type Ability struct {
	allow []*Rule
	deny  []*Rule
}

func NewAbility() *Ability {
	return &Ability{
		allow: []*Rule{},
		deny:  []*Rule{},
	}
}

func (a *Ability) Can(action string, subject string) {
	a.allow = append(a.allow, &Rule{
		Action:  action,
		Subject: subject,
	})
}

func (a *Ability) Cannot(action string, subject string) {
	a.deny = append(a.deny, &Rule{
		Action:  action,
		Subject: subject,
	})
}

func (a *Ability) Check(action string, subject string) (allow bool) {
	allow = false

	// Check for rules that allow the given action on the given subject
	for _, rule := range a.allow {
		if (rule.Subject == subject) && (rule.Action == action) {
			allow = true
		}
	}

	// Check for rules that explicitly deny the given action on the given subject
	for _, rule := range a.deny {
		if (rule.Subject == subject) && (rule.Action == action) {
			allow = false
		}
	}

	return allow
}