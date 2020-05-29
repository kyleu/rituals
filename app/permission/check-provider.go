package permission

import (
	"strings"

	"github.com/kyleu/rituals.dev/app/auth"
	"github.com/kyleu/rituals.dev/app/util"
)

func (s *Service) checkAuths(authEnabled bool, svc util.Service, perms Permissions, auths auth.Records) Errors {
	var ret Errors

	if authEnabled {
		for _, prv := range auth.AllProviders {
			ret = append(ret, providerCheck(svc, prv, perms, auths)...)
		}
	}

	return ret
}

func providerCheck(svc util.Service, p *auth.Provider, perms Permissions, auths auth.Records) Errors {
	var authMatches = auths.FindByProvider(p.Key)
	var permMatches = perms.FindByK(p.Key)
	if len(permMatches) == 0 {
		return nil
	}

	var ok = false
	for _, a := range authMatches {
		for _, p := range permMatches {
			split := strings.Split(p.V, ",")
			for _, domain := range split {
				if strings.HasSuffix(a.Email, domain) {
					ok = true
				}
			}
		}
	}

	if ok {
		return nil
	}

	msg := "you must log in using " + p.Title

	var emailDomains []string
	for _, p := range permMatches {
		if len(p.V) > 0 {
			emailDomains = append(emailDomains, p.V)
		}
	}

	if len(emailDomains) > 0 {
		msg += " with email address " + util.OxfordComma(emailDomains, "or")
	}

	msg += " to access this " + svc.Key
	return Errors{{Svc: svc.Key, Provider: p.Key, Match: strings.Join(emailDomains, ","), Message: msg}}
}
