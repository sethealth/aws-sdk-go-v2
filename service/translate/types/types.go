// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// The custom terminology applied to the input text by Amazon Translate for the
// translated text response. This is optional in the response and will only be
// present if you specified terminology input in the request. Currently, only one
// terminology can be applied per TranslateText request.
type AppliedTerminology struct {

	// The name of the custom terminology applied to the input text by Amazon Translate
	// for the translated text response.
	Name *string

	// The specific terms of the custom terminology applied to the input text by Amazon
	// Translate for the translated text response. A maximum of 250 terms will be
	// returned, and the specific terms applied will be the first 250 terms in the
	// source text.
	Terms []*Term
}

// The encryption key used to encrypt the custom terminologies used by Amazon
// Translate.
type EncryptionKey struct {

	// The Amazon Resource Name (ARN) of the encryption key being used to encrypt the
	// custom terminology.
	//
	// This member is required.
	Id *string

	// The type of encryption key used by Amazon Translate to encrypt custom
	// terminologies.
	//
	// This member is required.
	Type EncryptionKeyType
}

// The input configuration properties for requesting a batch translation job.
type InputDataConfig struct {

	// Describes the format of the data that you submit to Amazon Translate as input.
	// You can specify one of the following multipurpose internet mail extension (MIME)
	// types:
	//
	// * text/html: The input data consists of one or more HTML files. Amazon
	// Translate translates only the text that resides in the html element in each
	// file.
	//
	// * text/plain: The input data consists of one or more unformatted text
	// files. Amazon Translate translates every character in this type of input.
	//
	// *
	// application/vnd.openxmlformats-officedocument.wordprocessingml.document: The
	// input data consists of one or more Word documents (.docx).
	//
	// *
	// application/vnd.openxmlformats-officedocument.presentationml.presentation: The
	// input data consists of one or more PowerPoint Presentation files (.pptx).
	//
	// *
	// application/vnd.openxmlformats-officedocument.spreadsheetml.sheet: The input
	// data consists of one or more Excel Workbook files (.xlsx).
	//
	// If you structure
	// your input data as HTML, ensure that you set this parameter to text/html. By
	// doing so, you cut costs by limiting the translation to the contents of the html
	// element in each file. Otherwise, if you set this parameter to text/plain, your
	// costs will cover the translation of every character.
	//
	// This member is required.
	ContentType *string

	// The URI of the AWS S3 folder that contains the input file. The folder must be in
	// the same Region as the API endpoint you are calling.
	//
	// This member is required.
	S3Uri *string
}

// The number of documents successfully and unsuccessfully processed during a
// translation job.
type JobDetails struct {

	// The number of documents that could not be processed during a translation job.
	DocumentsWithErrorsCount *int32

	// The number of documents used as input in a translation job.
	InputDocumentsCount *int32

	// The number of documents successfully processed during a translation job.
	TranslatedDocumentsCount *int32
}

// The output configuration properties for a batch translation job.
type OutputDataConfig struct {

	// The URI of the S3 folder that contains a translation job's output file. The
	// folder must be in the same Region as the API endpoint that you are calling.
	//
	// This member is required.
	S3Uri *string
}

// The term being translated by the custom terminology.
type Term struct {

	// The source text of the term being translated by the custom terminology.
	SourceText *string

	// The target text of the term being translated by the custom terminology.
	TargetText *string
}

// The data associated with the custom terminology.
type TerminologyData struct {

	// The file containing the custom terminology data. Your version of the AWS SDK
	// performs a Base64-encoding on this field before sending a request to the AWS
	// service. Users of the SDK should not perform Base64-encoding themselves.
	//
	// This member is required.
	File []byte

	// The data format of the custom terminology. Either CSV or TMX.
	//
	// This member is required.
	Format TerminologyDataFormat
}

// The location of the custom terminology data.
type TerminologyDataLocation struct {

	// The location of the custom terminology data.
	//
	// This member is required.
	Location *string

	// The repository type for the custom terminology data.
	//
	// This member is required.
	RepositoryType *string
}

// The properties of the custom terminology.
type TerminologyProperties struct {

	// The Amazon Resource Name (ARN) of the custom terminology.
	Arn *string

	// The time at which the custom terminology was created, based on the timestamp.
	CreatedAt *time.Time

	// The description of the custom terminology properties.
	Description *string

	// The encryption key for the custom terminology.
	EncryptionKey *EncryptionKey

	// The time at which the custom terminology was last update, based on the
	// timestamp.
	LastUpdatedAt *time.Time

	// The name of the custom terminology.
	Name *string

	// The size of the file used when importing a custom terminology.
	SizeBytes *int32

	// The language code for the source text of the translation request for which the
	// custom terminology is being used.
	SourceLanguageCode *string

	// The language codes for the target languages available with the custom
	// terminology file. All possible target languages are returned in array.
	TargetLanguageCodes []*string

	// The number of terms included in the custom terminology.
	TermCount *int32
}

// Provides information for filtering a list of translation jobs. For more
// information, see ListTextTranslationJobs.
type TextTranslationJobFilter struct {

	// Filters the list of jobs by name.
	JobName *string

	// Filters the list of jobs based by job status.
	JobStatus JobStatus

	// Filters the list of jobs based on the time that the job was submitted for
	// processing and returns only the jobs submitted after the specified time. Jobs
	// are returned in descending order, newest to oldest.
	SubmittedAfterTime *time.Time

	// Filters the list of jobs based on the time that the job was submitted for
	// processing and returns only the jobs submitted before the specified time. Jobs
	// are returned in ascending order, oldest to newest.
	SubmittedBeforeTime *time.Time
}

// Provides information about a translation job.
type TextTranslationJobProperties struct {

	// The Amazon Resource Name (ARN) of an AWS Identity Access and Management (IAM)
	// role that granted Amazon Translate read access to the job's input data.
	DataAccessRoleArn *string

	// The time at which the translation job ended.
	EndTime *time.Time

	// The input configuration properties that were specified when the job was
	// requested.
	InputDataConfig *InputDataConfig

	// The number of documents successfully and unsuccessfully processed during the
	// translation job.
	JobDetails *JobDetails

	// The ID of the translation job.
	JobId *string

	// The user-defined name of the translation job.
	JobName *string

	// The status of the translation job.
	JobStatus JobStatus

	// An explanation of any errors that may have occured during the translation job.
	Message *string

	// The output configuration properties that were specified when the job was
	// requested.
	OutputDataConfig *OutputDataConfig

	// The language code of the language of the source text. The language must be a
	// language supported by Amazon Translate.
	SourceLanguageCode *string

	// The time at which the translation job was submitted.
	SubmittedTime *time.Time

	// The language code of the language of the target text. The language must be a
	// language supported by Amazon Translate.
	TargetLanguageCodes []*string

	// A list containing the names of the terminologies applied to a translation job.
	// Only one terminology can be applied per StartTextTranslationJob request at this
	// time.
	TerminologyNames []*string
}
