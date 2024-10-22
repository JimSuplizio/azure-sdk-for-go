//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package generated

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"time"
)

// MarshalXML implements the xml.Marshaller interface for type AccessPolicy.
func (a AccessPolicy) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias AccessPolicy
	aux := &struct {
		*alias
		Expiry *dateTimeRFC3339 `xml:"Expiry"`
		Start  *dateTimeRFC3339 `xml:"Start"`
	}{
		alias:  (*alias)(&a),
		Expiry: (*dateTimeRFC3339)(a.Expiry),
		Start:  (*dateTimeRFC3339)(a.Start),
	}
	return enc.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type AccessPolicy.
func (a *AccessPolicy) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type alias AccessPolicy
	aux := &struct {
		*alias
		Expiry *dateTimeRFC3339 `xml:"Expiry"`
		Start  *dateTimeRFC3339 `xml:"Start"`
	}{
		alias: (*alias)(a),
	}
	if err := dec.DecodeElement(aux, &start); err != nil {
		return err
	}
	a.Expiry = (*time.Time)(aux.Expiry)
	a.Start = (*time.Time)(aux.Start)
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type FileProperty.
func (f FileProperty) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias FileProperty
	aux := &struct {
		*alias
		ChangeTime     *dateTimeRFC3339 `xml:"ChangeTime"`
		CreationTime   *dateTimeRFC3339 `xml:"CreationTime"`
		LastAccessTime *dateTimeRFC3339 `xml:"LastAccessTime"`
		LastModified   *dateTimeRFC1123 `xml:"Last-Modified"`
		LastWriteTime  *dateTimeRFC3339 `xml:"LastWriteTime"`
	}{
		alias:          (*alias)(&f),
		ChangeTime:     (*dateTimeRFC3339)(f.ChangeTime),
		CreationTime:   (*dateTimeRFC3339)(f.CreationTime),
		LastAccessTime: (*dateTimeRFC3339)(f.LastAccessTime),
		LastModified:   (*dateTimeRFC1123)(f.LastModified),
		LastWriteTime:  (*dateTimeRFC3339)(f.LastWriteTime),
	}
	return enc.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type FileProperty.
func (f *FileProperty) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type alias FileProperty
	aux := &struct {
		*alias
		ChangeTime     *dateTimeRFC3339 `xml:"ChangeTime"`
		CreationTime   *dateTimeRFC3339 `xml:"CreationTime"`
		LastAccessTime *dateTimeRFC3339 `xml:"LastAccessTime"`
		LastModified   *dateTimeRFC1123 `xml:"Last-Modified"`
		LastWriteTime  *dateTimeRFC3339 `xml:"LastWriteTime"`
	}{
		alias: (*alias)(f),
	}
	if err := dec.DecodeElement(aux, &start); err != nil {
		return err
	}
	f.ChangeTime = (*time.Time)(aux.ChangeTime)
	f.CreationTime = (*time.Time)(aux.CreationTime)
	f.LastAccessTime = (*time.Time)(aux.LastAccessTime)
	f.LastModified = (*time.Time)(aux.LastModified)
	f.LastWriteTime = (*time.Time)(aux.LastWriteTime)
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type FilesAndDirectoriesListSegment.
func (f FilesAndDirectoriesListSegment) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias FilesAndDirectoriesListSegment
	aux := &struct {
		*alias
		Directories *[]*Directory `xml:"Directory"`
		Files       *[]*File      `xml:"File"`
	}{
		alias: (*alias)(&f),
	}
	if f.Directories != nil {
		aux.Directories = &f.Directories
	}
	if f.Files != nil {
		aux.Files = &f.Files
	}
	return enc.EncodeElement(aux, start)
}

// MarshalXML implements the xml.Marshaller interface for type ListHandlesResponse.
func (l ListHandlesResponse) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias ListHandlesResponse
	aux := &struct {
		*alias
		Handles *[]*Handle `xml:"Entries>Handle"`
	}{
		alias: (*alias)(&l),
	}
	if l.Handles != nil {
		aux.Handles = &l.Handles
	}
	return enc.EncodeElement(aux, start)
}

// MarshalXML implements the xml.Marshaller interface for type ListSharesResponse.
func (l ListSharesResponse) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias ListSharesResponse
	aux := &struct {
		*alias
		Shares *[]*Share `xml:"Shares>Share"`
	}{
		alias: (*alias)(&l),
	}
	if l.Shares != nil {
		aux.Shares = &l.Shares
	}
	return enc.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type Share.
func (s *Share) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type alias Share
	aux := &struct {
		*alias
		Metadata additionalProperties `xml:"Metadata"`
	}{
		alias: (*alias)(s),
	}
	if err := dec.DecodeElement(aux, &start); err != nil {
		return err
	}
	s.Metadata = (map[string]*string)(aux.Metadata)
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type ShareFileRangeList.
func (s ShareFileRangeList) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias ShareFileRangeList
	aux := &struct {
		*alias
		ClearRanges *[]*ClearRange `xml:"ClearRange"`
		Ranges      *[]*FileRange  `xml:"Range"`
	}{
		alias: (*alias)(&s),
	}
	if s.ClearRanges != nil {
		aux.ClearRanges = &s.ClearRanges
	}
	if s.Ranges != nil {
		aux.Ranges = &s.Ranges
	}
	return enc.EncodeElement(aux, start)
}

// MarshalJSON implements the json.Marshaller interface for type SharePermission.
func (s SharePermission) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "format", s.Format)
	populate(objectMap, "permission", s.Permission)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SharePermission.
func (s *SharePermission) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "format":
			err = unpopulate(val, "Format", &s.Format)
			delete(rawMsg, key)
		case "permission":
			err = unpopulate(val, "Permission", &s.Permission)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type ShareProperties.
func (s ShareProperties) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias ShareProperties
	aux := &struct {
		*alias
		AccessTierChangeTime          *dateTimeRFC1123 `xml:"AccessTierChangeTime"`
		DeletedTime                   *dateTimeRFC1123 `xml:"DeletedTime"`
		LastModified                  *dateTimeRFC1123 `xml:"Last-Modified"`
		NextAllowedQuotaDowngradeTime *dateTimeRFC1123 `xml:"NextAllowedQuotaDowngradeTime"`
	}{
		alias:                         (*alias)(&s),
		AccessTierChangeTime:          (*dateTimeRFC1123)(s.AccessTierChangeTime),
		DeletedTime:                   (*dateTimeRFC1123)(s.DeletedTime),
		LastModified:                  (*dateTimeRFC1123)(s.LastModified),
		NextAllowedQuotaDowngradeTime: (*dateTimeRFC1123)(s.NextAllowedQuotaDowngradeTime),
	}
	return enc.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type ShareProperties.
func (s *ShareProperties) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type alias ShareProperties
	aux := &struct {
		*alias
		AccessTierChangeTime          *dateTimeRFC1123 `xml:"AccessTierChangeTime"`
		DeletedTime                   *dateTimeRFC1123 `xml:"DeletedTime"`
		LastModified                  *dateTimeRFC1123 `xml:"Last-Modified"`
		NextAllowedQuotaDowngradeTime *dateTimeRFC1123 `xml:"NextAllowedQuotaDowngradeTime"`
	}{
		alias: (*alias)(s),
	}
	if err := dec.DecodeElement(aux, &start); err != nil {
		return err
	}
	s.AccessTierChangeTime = (*time.Time)(aux.AccessTierChangeTime)
	s.DeletedTime = (*time.Time)(aux.DeletedTime)
	s.LastModified = (*time.Time)(aux.LastModified)
	s.NextAllowedQuotaDowngradeTime = (*time.Time)(aux.NextAllowedQuotaDowngradeTime)
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StorageError.
func (s StorageError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "AuthenticationErrorDetail", s.AuthenticationErrorDetail)
	populate(objectMap, "Message", s.Message)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StorageError.
func (s *StorageError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "AuthenticationErrorDetail":
			err = unpopulate(val, "AuthenticationErrorDetail", &s.AuthenticationErrorDetail)
			delete(rawMsg, key)
		case "Message":
			err = unpopulate(val, "Message", &s.Message)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type StorageServiceProperties.
func (s StorageServiceProperties) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type alias StorageServiceProperties
	aux := &struct {
		*alias
		CORS *[]*CORSRule `xml:"Cors>CorsRule"`
	}{
		alias: (*alias)(&s),
	}
	if s.CORS != nil {
		aux.CORS = &s.CORS
	}
	return enc.EncodeElement(aux, start)
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
