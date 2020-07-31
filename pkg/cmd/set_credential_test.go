/*
 * Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"testing"

	"github.com/ZupIT/ritchie-cli/pkg/credential"
	"github.com/ZupIT/ritchie-cli/pkg/prompt"
)

func Test_setCredentialCmd_runPrompt(t *testing.T) {
	type fields struct {
		Setter        credential.Setter
		file          credential.ReaderWriterPather
		InputText     prompt.InputText
		InputBool     prompt.InputBool
		InputList     prompt.InputList
		InputPassword prompt.InputPassword
	}
	var tests = []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Run with success",
			fields: fields{
				Setter:        credSetterMock{},
				file:          credSettingsMock{},
				InputText:     inputSecretMock{},
				InputBool:     inputFalseMock{},
				InputList:     inputListCredMock{},
				InputPassword: inputPasswordMock{},
			},
			wantErr: false,
		},
		{
			name: "Run with success AddNew",
			fields: fields{
				Setter:        credSetterMock{},
				file:          credSettingsMock{},
				InputText:     inputSecretMock{},
				InputBool:     inputFalseMock{},
				InputList:     inputListCustomMock{credential.AddNew},
				InputPassword: inputPasswordMock{},
			},
			wantErr: false,
		},
		{
			name: "Fail when list return err",
			fields: fields{
				Setter:        credSetterMock{},
				file:          credSettingsMock{},
				InputText:     inputSecretMock{},
				InputBool:     inputFalseMock{},
				InputList:     inputListErrorMock{},
				InputPassword: inputPasswordMock{},
			},
			wantErr: true,
		},
		{
			name: "Fail when text return err",
			fields: fields{
				Setter:        credSetterMock{},
				file:          credSettingsMock{},
				InputText:     inputTextErrorMock{},
				InputBool:     inputFalseMock{},
				InputList:     inputListCustomMock{credential.AddNew},
				InputPassword: inputPasswordMock{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := NewSetCredentialCmd(
				tt.fields.Setter,
				tt.fields.file,
				tt.fields.InputText,
				tt.fields.InputBool,
				tt.fields.InputList,
				tt.fields.InputPassword,
			)
			o.PersistentFlags().Bool("stdin", false, "input by stdin")
			if err := o.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("setCredentialCmd_runPrompt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
