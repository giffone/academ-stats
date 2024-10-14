package zero_one_api

import (
	"academ_stats/internal/domain"
	"fmt"
)

type TopCadets interface {
	TopCadets(moduleID int) ([]byte, error)
}

func (z *zeroOneApi) TopCadets(moduleID int) ([]byte, error) {
	variables := map[string]interface{}{
		"moduleId":       moduleID,
		"moduleName":     domain.ModulePath,
		"checkpointName": domain.CheckpointPath,
		"admissionName":  domain.PiscinePath,
	}

	query := `fragment xpFields on xp_view_aggregate {
	xpTotal: aggregate {
		sum {
			total: amount
		}
	}
	paths: nodes {
		xp: amount
		path: pathByPath {
			attrs: object {
				id
				pathName: name
				language: attrs(path: "language")
			}
		}
		event {
			registrationId
			createdAt
			object{
				eventName: name
			}
		}
	}
}
		  
query topCadets($moduleId: Int!, $moduleName: String!, $checkpointName: String!, $admissionName: String!) {
	cadets: event_user(
		where: {event: {_and: [{path: {_eq: $moduleName}}, {registrationId: {_eq: $moduleId}}]}}
		order_by: [{level: desc_nulls_last}, {userLogin: asc_nulls_last}]
	) {
		userId
		login: userLogin
		level
		user{
			attrs
			registrations(
				where: {registration: {path: {_eq: $admissionName}}}
				order_by: {id: asc_nulls_last}
			){
				admission: registrationId
			}
		}
		journey: user {
			module: xps_aggregate(
				where: {event: {_and: [{path: {_eq: $moduleName}}, {registrationId: {_eq: $moduleId}}]}}
				order_by: [{event: {registrationId: asc_nulls_last}}, {amount: asc_nulls_last}, {pathByPath: {objectId: asc_nulls_last}}]
			) {
				...xpFields
			}
			checkpoint: xps_aggregate(
				where: {event: {path: {_eq: $checkpointName}}}
				order_by: [{event: {registrationId: asc_nulls_last}}, {amount: asc_nulls_last}]
			) {
				...xpFields
			}
			piscine: xps_aggregate(
				where: {event: {_and: [{path: {_neq: $admissionName}}, {pathByPath: {object: {type: {_eq: "piscine"}}}}]}}
				order_by: [{event: {registrationId: asc_nulls_last}}, {amount: asc_nulls_last}]
			) {
				...xpFields
			}
		}
	}
}`

	body, err := z.cli.Run(query, variables)
	if err != nil {
		return nil, fmt.Errorf("run: %w", err)
	}

	return body, nil
}
