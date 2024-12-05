package pkg

import "errors"

var (
	ErrStatusForbidden     = errors.New("forbidden")
	ErrStatusInternalError = errors.New("internal server error")
	ErrNoPrivilege         = errors.New("no permission to doing this task")

	// Authentication
	ErrEmailAlreadyExists       = errors.New("email already exists")
	ErrPhoneNumberAlreadyExists = errors.New("phone number already exists")
	ErrUserNotFound             = errors.New("user not found")
	ErrPasswordInvalid          = errors.New("password invalid")
	ErrOTPInvalid               = errors.New("otp invalid")
	ErrNeedToVerify             = errors.New("verify account false")
	ErrUserAlreadyVerified      = errors.New("user already verified")

	// Upload Cloudinary
	ErrUploadCloudinary = errors.New("upload cloudinary server error")

	// admin
	ErrAdminNotFound = errors.New("admin not found")
	ErrRole          = errors.New("role must be admin or super admin")

	// Report
	ErrReportNotFound = errors.New("report not found")

	// Date
	ErrDateFormat = errors.New("invalid date format")

	// Manage Task
	ErrTaskStepsNull           = errors.New("steps cannot be null")
	ErrTaskNotFound            = errors.New("task not found")
	ErrParsedTime              = errors.New("start date or end data is invalid")
	ErrThumbnail               = errors.New("thumbnail is required")
	ErrThumbnailMaximum        = errors.New("thumbnail must be one image")
	ErrUserTaskAlreadyAccepted = errors.New("cannot reject user task because it already accepted")

	// User Task
	ErrImageTaskNull                = errors.New("image task cannot be null")
	ErrUserTaskExist                = errors.New("user task already exist")
	ErrUserTaskNotFound             = errors.New("user task not found")
	ErrUserTaskDone                 = errors.New("user task already done")
	ErrTaskCannotBeFollowed         = errors.New("task cannot be followed")
	ErrUserNoHasTask                = errors.New("user has no task")
	ErrImagesExceed                 = errors.New("image exceed limit")
	ErrUserTaskNotReject            = errors.New("user task not reject")
	ErrUserTaskAlreadyReject        = errors.New("user task already reject")
	ErrUserTaskAlreadyApprove       = errors.New("user task already approve")
	ErrTaskStepNotFound             = errors.New("task step not found")
	ErrTaskStepDone                 = errors.New("task step already done")
	ErrUserTaskStepNotFound         = errors.New("user task step not found")
	ErrUserTaskNotCompleted         = errors.New("task step not completed")
	ErrStepNotInOrder               = errors.New("task step must be completed in order")
	ErrUserTaskStepAlreadyCompleted = errors.New("user task step already completed")

	// manage achievement
	ErrAchievementLevelAlreadyExist = errors.New("achievement level already exist")
	ErrAchievementNotFound          = errors.New("achievement not found")
	ErrBadge                        = errors.New("badge is required")
	ErrBadgeMaximum                 = errors.New("badge must be one image")

	// Custom Data
	ErrCustomDataNotFound = errors.New("custom data not found")

	// manage video
	ErrVideoTitleAlreadyExist        = errors.New("video title already exist")
	ErrVideoCategoryNameAlreadyExist = errors.New("video category name already exist")
	ErrNoVideoIdFoundOnUrl           = errors.New("no video id found on url")
	ErrVideoNotFound                 = errors.New("video not found")
	ErrVideoService                  = errors.New("video service error")
	ErrApiYouTube                    = errors.New("api youtube error")
	ErrParsingUrl                    = errors.New("parsing url error")
	ErrVideoCategory                 = errors.New("content category is required")
	ErrVideoTrashCategory            = errors.New("waste category is required")
	ErrNameCategoryVideoNotFound     = errors.New("name content category not found")
	ErrNameTrashCategoryNotFound     = errors.New("name waste category not found")

	// user achievement
	ErrUserNotHasHistoryPoint = errors.New("user not has history points")

	// About Us
	ErrAboutUsCategoryNotFound = errors.New("about us with that category not found")

	// Article
	ErrArticleNotFound         = errors.New("article not found")
	ErrCategoryArticleNotFound = errors.New("invalid category type")

	// Error file
	ErrFileTooLarge    = errors.New("upload image size must less than 2MB")
	ErrInvalidFileType = errors.New("invalid file type")
	ErrOpenFile        = errors.New("failed to open file")
)
