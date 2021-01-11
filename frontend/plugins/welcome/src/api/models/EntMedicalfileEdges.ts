/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntDentalExpense,
    EntDentalExpenseFromJSON,
    EntDentalExpenseFromJSONTyped,
    EntDentalExpenseToJSON,
    EntDentist,
    EntDentistFromJSON,
    EntDentistFromJSONTyped,
    EntDentistToJSON,
    EntNurse,
    EntNurseFromJSON,
    EntNurseFromJSONTyped,
    EntNurseToJSON,
    EntPatient,
    EntPatientFromJSON,
    EntPatientFromJSONTyped,
    EntPatientToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntMedicalfileEdges
 */
export interface EntMedicalfileEdges {
    /**
     * Dentalexpenses holds the value of the dentalexpenses edge.
     * @type {Array<EntDentalExpense>}
     * @memberof EntMedicalfileEdges
     */
    dentalexpenses?: Array<EntDentalExpense>;
    /**
     * 
     * @type {EntDentist}
     * @memberof EntMedicalfileEdges
     */
    dentist?: EntDentist;
    /**
     * 
     * @type {EntNurse}
     * @memberof EntMedicalfileEdges
     */
    nurse?: EntNurse;
    /**
     * 
     * @type {EntPatient}
     * @memberof EntMedicalfileEdges
     */
    patient?: EntPatient;
}

export function EntMedicalfileEdgesFromJSON(json: any): EntMedicalfileEdges {
    return EntMedicalfileEdgesFromJSONTyped(json, false);
}

export function EntMedicalfileEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntMedicalfileEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'dentalexpenses': !exists(json, 'dentalexpenses') ? undefined : ((json['dentalexpenses'] as Array<any>).map(EntDentalExpenseFromJSON)),
        'dentist': !exists(json, 'dentist') ? undefined : EntDentistFromJSON(json['dentist']),
        'nurse': !exists(json, 'nurse') ? undefined : EntNurseFromJSON(json['nurse']),
        'patient': !exists(json, 'patient') ? undefined : EntPatientFromJSON(json['patient']),
    };
}

export function EntMedicalfileEdgesToJSON(value?: EntMedicalfileEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'dentalexpenses': value.dentalexpenses === undefined ? undefined : ((value.dentalexpenses as Array<any>).map(EntDentalExpenseToJSON)),
        'dentist': EntDentistToJSON(value.dentist),
        'nurse': EntNurseToJSON(value.nurse),
        'patient': EntPatientToJSON(value.patient),
    };
}


