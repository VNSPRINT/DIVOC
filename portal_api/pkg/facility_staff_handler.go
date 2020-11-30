package pkg

import (
	"github.com/divoc/portal-api/swagger_gen/models"
	log "github.com/sirupsen/logrus"
)

func GetFacilityStaffs(authHeader string) error {
	bearerToken, err := getToken(authHeader)
	_, err = getClaimBody(bearerToken)
	if err != nil {
		log.Errorf("Error while parsing token : %s", bearerToken)
		return err
	}
	return nil
}

func CreateFacilityStaff(staff *models.FacilityStaff, authHeader string) error {
	bearerToken, err := getToken(authHeader)
	claimBody, err := getClaimBody(bearerToken)
	if err != nil {
		log.Errorf("Error while parsing token : %s", bearerToken)
		return err
	}
	userRequest := KeyCloakUserRequest{
		Username: staff.MobileNumber,
		Enabled:  "true",
		Attributes: KeycloakUserAttributes{
			MobileNumber: []string{staff.MobileNumber},
			EmployeeID:   staff.EmployeeID,
			FullName:     staff.Name,
			FacilityCode: claimBody.FacilityCode,
		},
	}
	resp, err := CreateKeycloakUser(userRequest, authHeader)
	log.Info("Created keycloak user ", resp.Response().StatusCode, " ", resp.String())
	if err != nil || !isUserCreatedOrAlreadyExists(resp) {
		log.Errorf("Error while creating keycloak user : %s", staff.MobileNumber)
		return err
	} else {
		log.Info("Setting up roles for the user ", staff.MobileNumber)
		keycloakUserId := getKeycloakUserId(resp, userRequest, authHeader)
		if keycloakUserId != "" {
			_ = addUserToGroup(keycloakUserId, staff.RoleID, authHeader)
		} else {
			log.Error("Unable to map keycloak user id for ", staff.MobileNumber)
		}
	}
	return nil
}
