package approvalProcesses

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseActionLists(t *testing.T) {
	data := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<ApprovalProcess xmlns="http://soap.sforce.com/2006/04/metadata">
    <active>true</active>
    <approvalStep>
        <allowDelegate>true</allowDelegate>
        <approvalActions>
            <action>
                <name>Step_Approved_One</name>
                <type>FieldUpdate</type>
            </action>
            <action>
                <name>Step_Approved_Two</name>
                <type>FieldUpdate</type>
            </action>
        </approvalActions>
        <label>Step One</label>
        <name>Step_One</name>
        <rejectionActions>
            <action>
                <name>Step_Rejected_One</name>
                <type>FieldUpdate</type>
            </action>
            <action>
                <name>Step_Rejected_Two</name>
                <type>Alert</type>
            </action>
        </rejectionActions>
    </approvalStep>
    <finalApprovalActions>
        <action>
            <name>Final_Approved_One</name>
            <type>FieldUpdate</type>
        </action>
        <action>
            <name>Final_Approved_Two</name>
            <type>FieldUpdate</type>
        </action>
    </finalApprovalActions>
    <finalRejectionActions>
        <action>
            <name>Final_Rejected_One</name>
            <type>FieldUpdate</type>
        </action>
    </finalRejectionActions>
    <initialSubmissionActions>
        <action>
            <name>Submitted_One</name>
            <type>FieldUpdate</type>
        </action>
        <action>
            <name>Submitted_Two</name>
            <type>FieldUpdate</type>
        </action>
    </initialSubmissionActions>
    <label>Action Lists</label>
    <recallActions>
        <action>
            <name>Recalled_One</name>
            <type>FieldUpdate</type>
        </action>
        <action>
            <name>Recalled_Two</name>
            <type>FieldUpdate</type>
        </action>
    </recallActions>
</ApprovalProcess>`)

	var process ApprovalProcess
	require.NoError(t, xml.Unmarshal(data, &process))

	require.Len(t, process.ApprovalStep, 1)
	step := process.ApprovalStep[0]
	require.NotNil(t, step.ApprovalActions)
	require.Len(t, step.ApprovalActions.Action, 2)
	assert.Equal(t, "Step_Approved_One", step.ApprovalActions.Action[0].Name.Text)
	assert.Equal(t, "Step_Approved_Two", step.ApprovalActions.Action[1].Name.Text)
	require.NotNil(t, step.RejectionActions)
	require.Len(t, step.RejectionActions.Action, 2)
	assert.Equal(t, "Step_Rejected_One", step.RejectionActions.Action[0].Name.Text)
	assert.Equal(t, "Alert", step.RejectionActions.Action[1].Type.Text)

	require.NotNil(t, process.FinalApprovalActions)
	require.Len(t, process.FinalApprovalActions.Action, 2)
	assert.Equal(t, "Final_Approved_One", process.FinalApprovalActions.Action[0].Name.Text)
	require.NotNil(t, process.FinalRejectionActions)
	require.Len(t, process.FinalRejectionActions.Action, 1)
	assert.Equal(t, "Final_Rejected_One", process.FinalRejectionActions.Action[0].Name.Text)
	require.NotNil(t, process.InitialSubmissionActions)
	require.Len(t, process.InitialSubmissionActions.Action, 2)
	assert.Equal(t, "Submitted_One", process.InitialSubmissionActions.Action[0].Name.Text)
	require.NotNil(t, process.RecallActions)
	require.Len(t, process.RecallActions.Action, 2)
	assert.Equal(t, "Recalled_Two", process.RecallActions.Action[1].Name.Text)
}
