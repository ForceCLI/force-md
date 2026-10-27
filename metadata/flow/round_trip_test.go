package flow

import (
	"testing"

	"github.com/ForceCLI/force-md/internal"
)

func assertFlowRoundTrip(t *testing.T, contents string) {
	t.Helper()
	f := openFlow(t, contents)
	out, err := internal.Marshal(f)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(out) != contents {
		t.Errorf("round trip changed the flow:\n%s", out)
	}
}

// A flow migrated from a workflow rule keeps its metrics flag, the rule name,
// the metadata values on conditions and input assignments, the nil location
// of its update element, and its empty text elements written open.
func TestMigratedFlowRoundTrip(t *testing.T) {
	assertFlowRoundTrip(t, `<?xml version="1.0" encoding="UTF-8"?>
<Flow xmlns="http://soap.sforce.com/2006/04/metadata" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
    <actionCalls>
        <name>Post_To_Chatter</name>
        <label>Post To Chatter</label>
        <locationX>0</locationX>
        <locationY>0</locationY>
        <actionName>chatterPost</actionName>
        <actionType>chatterPost</actionType>
        <flowTransactionModel>CurrentTransaction</flowTransactionModel>
        <inputParameters>
            <name>text</name>
            <value>
                <elementReference>Message</elementReference>
                <inputConfiguratorMode>Resource</inputConfiguratorMode>
            </value>
        </inputParameters>
        <inputParameters>
            <name>recipientAddresses</name>
            <value/>
        </inputParameters>
        <nameSegment>chatterPost</nameSegment>
        <offset>0</offset>
        <storeOutputAutomatically>true</storeOutputAutomatically>
        <versionString>1.0.0</versionString>
    </actionCalls>
    <apiVersion>55.0</apiVersion>
    <areMetricsLoggedToDataCloud>false</areMetricsLoggedToDataCloud>
    <decisions>
        <name>myDecision</name>
        <label>myDecision</label>
        <locationX>50</locationX>
        <locationY>0</locationY>
        <defaultConnectorLabel>default</defaultConnectorLabel>
        <rules>
            <name>myRule_1</name>
            <conditionLogic>and</conditionLogic>
            <conditions>
                <processMetadataValues>
                    <name>inputDataType</name>
                    <value>
                        <stringValue></stringValue>
                    </value>
                </processMetadataValues>
                <leftValueReference>myVariable_current.Active__c</leftValueReference>
                <operator>EqualTo</operator>
                <rightValue>
                    <booleanValue>false</booleanValue>
                </rightValue>
            </conditions>
            <label>Active is False</label>
        </rules>
    </decisions>
    <label>Account Active is False</label>
    <migratedFromWorkflowRuleName>Account Active is False</migratedFromWorkflowRuleName>
    <processType>AutoLaunchedFlow</processType>
    <recordCreates>
        <name>Create_Task</name>
        <label>Create Task</label>
        <locationX>0</locationX>
        <locationY>0</locationY>
        <inputAssignments>
            <processMetadataValues>
                <name>dataType</name>
                <value>
                    <stringValue>String</stringValue>
                </value>
            </processMetadataValues>
            <field>Subject</field>
            <value>
                <stringValue>Follow up</stringValue>
            </value>
        </inputAssignments>
        <object>Task</object>
    </recordCreates>
    <recordUpdates>
        <processMetadataValues>
            <name>extraTypeInfo</name>
        </processMetadataValues>
        <description></description>
        <name>mainUpdate</name>
        <label>mainUpdate</label>
        <locationX xsi:nil="true"/>
        <locationY xsi:nil="true"/>
        <inputReference>$Record</inputReference>
    </recordUpdates>
    <status>Active</status>
    <textTemplates>
        <description>Message for the account owner</description>
        <name>Message</name>
        <isViewedAsPlainText>true</isViewedAsPlainText>
        <text>Account is inactive</text>
    </textTemplates>
</Flow>
`)
}

// A screen flow keeps the style properties of its fields at each level of
// nesting, the screen's style settings, a dynamic choice set's output
// assignments, multi-line help text, and the legacy steps.
func TestScreenFlowRoundTrip(t *testing.T) {
	assertFlowRoundTrip(t, `<?xml version="1.0" encoding="UTF-8"?>
<Flow xmlns="http://soap.sforce.com/2006/04/metadata">
    <apiVersion>60.0</apiVersion>
    <dynamicChoiceSets>
        <name>Contacts</name>
        <dataType>String</dataType>
        <displayField>Name</displayField>
        <object>Contact</object>
        <outputAssignments>
            <assignToReference>ContactId</assignToReference>
            <field>Id</field>
        </outputAssignments>
        <outputAssignments>
            <assignToReference>ContactEmail</assignToReference>
            <field>Email</field>
        </outputAssignments>
        <valueField>Id</valueField>
    </dynamicChoiceSets>
    <label>Event Form</label>
    <processType>Flow</processType>
    <screens>
        <name>Event_Form</name>
        <label>Event Form</label>
        <locationX>0</locationX>
        <locationY>0</locationY>
        <allowBack>true</allowBack>
        <allowFinish>true</allowFinish>
        <allowPause>true</allowPause>
        <fields>
            <name>Section</name>
            <fieldType>RegionContainer</fieldType>
            <fields>
                <name>Column</name>
                <fieldType>Region</fieldType>
                <fields>
                    <name>Attendees</name>
                    <dataType>Number</dataType>
                    <fieldText>Attendees</fieldText>
                    <fieldType>InputField</fieldType>
                    <isRequired>false</isRequired>
                    <scale>0</scale>
                    <styleProperties>
                        <verticalAlignment>
                            <stringValue>top</stringValue>
                        </verticalAlignment>
                        <width>
                            <stringValue>12</stringValue>
                        </width>
                    </styleProperties>
                </fields>
                <inputParameters>
                    <name>width</name>
                    <value>
                        <stringValue>12</stringValue>
                    </value>
                </inputParameters>
                <isRequired>false</isRequired>
                <styleProperties>
                    <verticalAlignment>
                        <stringValue>top</stringValue>
                    </verticalAlignment>
                    <width>
                        <stringValue>12</stringValue>
                    </width>
                </styleProperties>
            </fields>
            <isRequired>false</isRequired>
            <regionContainerType>SectionWithoutHeader</regionContainerType>
        </fields>
        <fields>
            <name>Event_Type</name>
            <dataType>String</dataType>
            <fieldText>Event Type</fieldText>
            <fieldType>InputField</fieldType>
            <helpText>Outreach – meetings
External Event – conferences</helpText>
            <isRequired>false</isRequired>
            <styleProperties>
                <verticalAlignment>
                    <stringValue>top</stringValue>
                </verticalAlignment>
                <width>
                    <stringValue>12</stringValue>
                </width>
            </styleProperties>
        </fields>
        <showFooter>true</showFooter>
        <showHeader>true</showHeader>
        <styleSettings>
            <propertyName>--flowruntime-c-flow-header-color</propertyName>
            <propertyValue>
                <stringValue>#777</stringValue>
            </propertyValue>
            <scope>Header</scope>
        </styleSettings>
    </screens>
    <status>Draft</status>
    <steps>
        <description>Initial event registration.</description>
        <name>Event_Form_Stage_1</name>
        <label>Event Form Stage 1</label>
        <locationX>91</locationX>
        <locationY>75</locationY>
        <connectors>
            <targetReference>Event_Form</targetReference>
        </connectors>
    </steps>
</Flow>
`)
}

// A screen keeps the actions it runs while displayed and the triggers that
// name the events which run them.
func TestScreenActionsAndTriggersRoundTrip(t *testing.T) {
	assertFlowRoundTrip(t, `<?xml version="1.0" encoding="UTF-8"?>
<Flow xmlns="http://soap.sforce.com/2006/04/metadata">
    <apiVersion>64.0</apiVersion>
    <label>Inquiry</label>
    <processType>Flow</processType>
    <screens>
        <name>Screen_Identity</name>
        <label>Identity</label>
        <locationX>0</locationX>
        <locationY>0</locationY>
        <actions>
            <name>Zip_Lookup</name>
            <actionName>Inquiry_Zip_Lookup</actionName>
            <actionType>flow</actionType>
            <inputParameters>
                <name>zip</name>
                <value>
                    <elementReference>CallerZip</elementReference>
                </value>
            </inputParameters>
            <label>Zip Lookup</label>
            <nameSegment>Inquiry_Zip_Lookup</nameSegment>
        </actions>
        <allowBack>true</allowBack>
        <allowFinish>true</allowFinish>
        <allowPause>false</allowPause>
        <showFooter>true</showFooter>
        <showHeader>false</showHeader>
        <triggers>
            <eventName>flow__screeninit</eventName>
            <eventSource>Screen_Identity</eventSource>
            <handlers>
                <screenActionName>Zip_Lookup</screenActionName>
            </handlers>
        </triggers>
        <triggers>
            <eventName>flow__screenfieldattributechange</eventName>
            <eventSource>CallerZip</eventSource>
            <handlers>
                <screenActionName>Zip_Lookup</screenActionName>
            </handlers>
        </triggers>
    </screens>
</Flow>
`)
}
