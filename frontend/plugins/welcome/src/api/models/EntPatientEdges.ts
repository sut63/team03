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
    EntAppointment,
    EntAppointmentFromJSON,
    EntAppointmentFromJSONTyped,
    EntAppointmentToJSON,
    EntDisease,
    EntDiseaseFromJSON,
    EntDiseaseFromJSONTyped,
    EntDiseaseToJSON,
    EntGender,
    EntGenderFromJSON,
    EntGenderFromJSONTyped,
    EntGenderToJSON,
    EntMedicalCare,
    EntMedicalCareFromJSON,
    EntMedicalCareFromJSONTyped,
    EntMedicalCareToJSON,
    EntMedicalfile,
    EntMedicalfileFromJSON,
    EntMedicalfileFromJSONTyped,
    EntMedicalfileToJSON,
    EntNurse,
    EntNurseFromJSON,
    EntNurseFromJSONTyped,
    EntNurseToJSON,
    EntQueue,
    EntQueueFromJSON,
    EntQueueFromJSONTyped,
    EntQueueToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPatientEdges
 */
export interface EntPatientEdges {
    /**
     * Appointment holds the value of the appointment edge.
     * @type {Array<EntAppointment>}
     * @memberof EntPatientEdges
     */
    appointment?: Array<EntAppointment>;
    /**
     * 
     * @type {EntDisease}
     * @memberof EntPatientEdges
     */
    disease?: EntDisease;
    /**
     * 
     * @type {EntGender}
     * @memberof EntPatientEdges
     */
    gender?: EntGender;
    /**
     * 
     * @type {EntMedicalCare}
     * @memberof EntPatientEdges
     */
    medicalcare?: EntMedicalCare;
    /**
     * Medicalfiles holds the value of the medicalfiles edge.
     * @type {Array<EntMedicalfile>}
     * @memberof EntPatientEdges
     */
    medicalfiles?: Array<EntMedicalfile>;
    /**
     * 
     * @type {EntNurse}
     * @memberof EntPatientEdges
     */
    nurse?: EntNurse;
    /**
     * Queue holds the value of the queue edge.
     * @type {Array<EntQueue>}
     * @memberof EntPatientEdges
     */
    queue?: Array<EntQueue>;
}

export function EntPatientEdgesFromJSON(json: any): EntPatientEdges {
    return EntPatientEdgesFromJSONTyped(json, false);
}

export function EntPatientEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPatientEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'appointment': !exists(json, 'appointment') ? undefined : ((json['appointment'] as Array<any>).map(EntAppointmentFromJSON)),
        'disease': !exists(json, 'disease') ? undefined : EntDiseaseFromJSON(json['disease']),
        'gender': !exists(json, 'gender') ? undefined : EntGenderFromJSON(json['gender']),
        'medicalcare': !exists(json, 'Medicalcare') ? undefined : EntMedicalCareFromJSON(json['Medicalcare']),
        'medicalfiles': !exists(json, 'Medicalfiles') ? undefined : ((json['Medicalfiles'] as Array<any>).map(EntMedicalfileFromJSON)),
        'nurse': !exists(json, 'nurse') ? undefined : EntNurseFromJSON(json['nurse']),
        'queue': !exists(json, 'queue') ? undefined : ((json['queue'] as Array<any>).map(EntQueueFromJSON)),
    };
}

export function EntPatientEdgesToJSON(value?: EntPatientEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'appointment': value.appointment === undefined ? undefined : ((value.appointment as Array<any>).map(EntAppointmentToJSON)),
        'disease': EntDiseaseToJSON(value.disease),
        'gender': EntGenderToJSON(value.gender),
        'medicalcare': EntMedicalCareToJSON(value.medicalcare),
        'medicalfiles': value.medicalfiles === undefined ? undefined : ((value.medicalfiles as Array<any>).map(EntMedicalfileToJSON)),
        'nurse': EntNurseToJSON(value.nurse),
        'queue': value.queue === undefined ? undefined : ((value.queue as Array<any>).map(EntQueueToJSON)),
    };
}


