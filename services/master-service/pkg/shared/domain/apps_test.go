package domain

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFullPathParentFromChild(t *testing.T) {
	perm := Permission{Code: "user-service", Childs: []Permission{
		{Code: "member", Childs: []Permission{
			{Code: "getAllMember", Childs: []Permission{
				{Code: "addMember"},
				{Code: "updateMember", Childs: []Permission{
					{Code: "upgradeMember"},
				}},
				{Code: "deleteMember"},
			}},
		}},
		{Code: "merchant", Childs: []Permission{
			{Code: "getAllMerchant", Childs: []Permission{
				{Code: "addMerchant"},
				{Code: "upgradeMerchant"},
			}},
		}},
	}}

	allVisitedPath := perm.GetAllVisitedPath()
	var case1, case2, case3 []string

	for _, perm := range allVisitedPath["upgradeMember"] {
		case1 = append(case1, perm.Code)
	}
	assert.Equal(t, "user-service.member.getAllMember.updateMember", strings.Join(case1, "."))

	for _, perm := range allVisitedPath["upgradeMerchant"] {
		case2 = append(case2, perm.Code)
	}
	assert.Equal(t, "user-service.merchant.getAllMerchant", strings.Join(case2, "."))

	for _, perm := range allVisitedPath["getAllMerchant"] {
		case3 = append(case3, perm.Code)
	}
	assert.Equal(t, "user-service.merchant", strings.Join(case3, "."))
}
