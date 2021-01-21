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
    EntMedicalfileEdges,
    EntMedicalfileEdgesFromJSON,
    EntMedicalfileEdgesFromJSONTyped,
    EntMedicalfileEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntMedicalfile
 */
export interface EntMedicalfile {
    /**
     * AddedTime holds the value of the "AddedTime" field.
     * @type {string}
     * @memberof EntMedicalfile
     */
    addedTime?: string;
    /**
     * Detial holds the value of the "Detial" field.
     * @type {string}
     * @memberof EntMedicalfile
     */
    detial?: string;
    /**
     * DrugAllergy holds the value of the "DrugAllergy" field.
     * @type {string}
     * @memberof EntMedicalfile
     */
    drugAllergy?: string;
    /**
     * Medno holds the value of the "Medno" field.
     * @type {string}
     * @memberof EntMedicalfile
     */
    medno?: string;
    /**
     * 
     * @type {EntMedicalfileEdges}
     * @memberof EntMedicalfile
     */
    edges?: EntMedicalfileEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntMedicalfile
     */
    id?: number;
}

export function EntMedicalfileFromJSON(json: any): EntMedicalfile {
    return EntMedicalfileFromJSONTyped(json, false);
}

export function EntMedicalfileFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntMedicalfile {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'addedTime': !exists(json, 'AddedTime') ? undefined : json['AddedTime'],
        'detial': !exists(json, 'Detial') ? undefined : json['Detial'],
        'drugAllergy': !exists(json, 'DrugAllergy') ? undefined : json['DrugAllergy'],
        'medno': !exists(json, 'Medno') ? undefined : json['Medno'],
        'edges': !exists(json, 'edges') ? undefined : EntMedicalfileEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntMedicalfileToJSON(value?: EntMedicalfile | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'AddedTime': value.addedTime,
        'Detial': value.detial,
        'DrugAllergy': value.drugAllergy,
        'Medno': value.medno,
        'edges': EntMedicalfileEdgesToJSON(value.edges),
        'id': value.id,
    };
}


