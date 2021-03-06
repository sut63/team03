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
    EntAppointmentEdges,
    EntAppointmentEdgesFromJSON,
    EntAppointmentEdgesFromJSONTyped,
    EntAppointmentEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntAppointment
 */
export interface EntAppointment {
    /**
     * AppointID holds the value of the "AppointID" field.
     * @type {string}
     * @memberof EntAppointment
     */
    appointID?: string;
    /**
     * Datetime holds the value of the "Datetime" field.
     * @type {string}
     * @memberof EntAppointment
     */
    datetime?: string;
    /**
     * Detail holds the value of the "Detail" field.
     * @type {string}
     * @memberof EntAppointment
     */
    detail?: string;
    /**
     * Remark holds the value of the "Remark" field.
     * @type {string}
     * @memberof EntAppointment
     */
    remark?: string;
    /**
     * 
     * @type {EntAppointmentEdges}
     * @memberof EntAppointment
     */
    edges?: EntAppointmentEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntAppointment
     */
    id?: number;
}

export function EntAppointmentFromJSON(json: any): EntAppointment {
    return EntAppointmentFromJSONTyped(json, false);
}

export function EntAppointmentFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntAppointment {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'appointID': !exists(json, 'AppointID') ? undefined : json['AppointID'],
        'datetime': !exists(json, 'Datetime') ? undefined : json['Datetime'],
        'detail': !exists(json, 'Detail') ? undefined : json['Detail'],
        'remark': !exists(json, 'Remark') ? undefined : json['Remark'],
        'edges': !exists(json, 'edges') ? undefined : EntAppointmentEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntAppointmentToJSON(value?: EntAppointment | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'AppointID': value.appointID,
        'Datetime': value.datetime,
        'Detail': value.detail,
        'Remark': value.remark,
        'edges': EntAppointmentEdgesToJSON(value.edges),
        'id': value.id,
    };
}


