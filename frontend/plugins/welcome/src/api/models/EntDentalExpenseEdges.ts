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
    EntMedicalfile,
    EntMedicalfileFromJSON,
    EntMedicalfileFromJSONTyped,
    EntMedicalfileToJSON,
    EntNurse,
    EntNurseFromJSON,
    EntNurseFromJSONTyped,
    EntNurseToJSON,
    EntPriceType,
    EntPriceTypeFromJSON,
    EntPriceTypeFromJSONTyped,
    EntPriceTypeToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDentalExpenseEdges
 */
export interface EntDentalExpenseEdges {
    /**
     * 
     * @type {EntMedicalfile}
     * @memberof EntDentalExpenseEdges
     */
    medicalfile?: EntMedicalfile;
    /**
     * 
     * @type {EntNurse}
     * @memberof EntDentalExpenseEdges
     */
    nurse?: EntNurse;
    /**
     * 
     * @type {EntPriceType}
     * @memberof EntDentalExpenseEdges
     */
    pricetype?: EntPriceType;
}

export function EntDentalExpenseEdgesFromJSON(json: any): EntDentalExpenseEdges {
    return EntDentalExpenseEdgesFromJSONTyped(json, false);
}

export function EntDentalExpenseEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDentalExpenseEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'medicalfile': !exists(json, 'medicalfile') ? undefined : EntMedicalfileFromJSON(json['medicalfile']),
        'nurse': !exists(json, 'nurse') ? undefined : EntNurseFromJSON(json['nurse']),
        'pricetype': !exists(json, 'pricetype') ? undefined : EntPriceTypeFromJSON(json['pricetype']),
    };
}

export function EntDentalExpenseEdgesToJSON(value?: EntDentalExpenseEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'medicalfile': EntMedicalfileToJSON(value.medicalfile),
        'nurse': EntNurseToJSON(value.nurse),
        'pricetype': EntPriceTypeToJSON(value.pricetype),
    };
}

