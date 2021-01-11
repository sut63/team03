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
    EntDegree,
    EntDegreeFromJSON,
    EntDegreeFromJSONTyped,
    EntDegreeToJSON,
    EntExpert,
    EntExpertFromJSON,
    EntExpertFromJSONTyped,
    EntExpertToJSON,
    EntGender,
    EntGenderFromJSON,
    EntGenderFromJSONTyped,
    EntGenderToJSON,
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
 * @interface EntDentistEdges
 */
export interface EntDentistEdges {
    /**
     * Appointment holds the value of the appointment edge.
     * @type {Array<EntAppointment>}
     * @memberof EntDentistEdges
     */
    appointment?: Array<EntAppointment>;
    /**
     * 
     * @type {EntDegree}
     * @memberof EntDentistEdges
     */
    degree?: EntDegree;
    /**
     * 
     * @type {EntExpert}
     * @memberof EntDentistEdges
     */
    expert?: EntExpert;
    /**
     * 
     * @type {EntGender}
     * @memberof EntDentistEdges
     */
    gender?: EntGender;
    /**
     * Medicalfiles holds the value of the medicalfiles edge.
     * @type {Array<EntMedicalfile>}
     * @memberof EntDentistEdges
     */
    medicalfiles?: Array<EntMedicalfile>;
    /**
     * 
     * @type {EntNurse}
     * @memberof EntDentistEdges
     */
    nurse?: EntNurse;
    /**
     * Queue holds the value of the queue edge.
     * @type {Array<EntQueue>}
     * @memberof EntDentistEdges
     */
    queue?: Array<EntQueue>;
}

export function EntDentistEdgesFromJSON(json: any): EntDentistEdges {
    return EntDentistEdgesFromJSONTyped(json, false);
}

export function EntDentistEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDentistEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'appointment': !exists(json, 'appointment') ? undefined : ((json['appointment'] as Array<any>).map(EntAppointmentFromJSON)),
        'degree': !exists(json, 'degree') ? undefined : EntDegreeFromJSON(json['degree']),
        'expert': !exists(json, 'expert') ? undefined : EntExpertFromJSON(json['expert']),
        'gender': !exists(json, 'gender') ? undefined : EntGenderFromJSON(json['gender']),
        'medicalfiles': !exists(json, 'medicalfiles') ? undefined : ((json['medicalfiles'] as Array<any>).map(EntMedicalfileFromJSON)),
        'nurse': !exists(json, 'nurse') ? undefined : EntNurseFromJSON(json['nurse']),
        'queue': !exists(json, 'queue') ? undefined : ((json['queue'] as Array<any>).map(EntQueueFromJSON)),
    };
}

export function EntDentistEdgesToJSON(value?: EntDentistEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'appointment': value.appointment === undefined ? undefined : ((value.appointment as Array<any>).map(EntAppointmentToJSON)),
        'degree': EntDegreeToJSON(value.degree),
        'expert': EntExpertToJSON(value.expert),
        'gender': EntGenderToJSON(value.gender),
        'medicalfiles': value.medicalfiles === undefined ? undefined : ((value.medicalfiles as Array<any>).map(EntMedicalfileToJSON)),
        'nurse': EntNurseToJSON(value.nurse),
        'queue': value.queue === undefined ? undefined : ((value.queue as Array<any>).map(EntQueueToJSON)),
    };
}

