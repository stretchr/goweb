package paths

var PathSeperator string = "/"
var FileExtensionSeparator string = "."

var RegexDynamicSegment string = "(.*)"
var RegexOptionalPathSeperator string = "(/?)"

/*
  Segments
*/

const segmentDynamicPrefix string = "{"
const segmentDynamicSuffix string = "}"
const segmentOptionalDynamicPrefix string = "["
const segmentOptionalDynamicSuffix string = "]"
const segmentWildcard string = "*"
const segmentCatchAll string = "***"

/*
  Public
*/

const MatchAllPaths string = segmentCatchAll
