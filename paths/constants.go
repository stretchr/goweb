package paths

var PathSeperator string = "/"
var FileExtensionSeparator string = "."

var RegexDynamicSegment string = "(.*)"
var RegexOptionalPathSeperator string = "(/?)"

/*
  Segments
*/

var segmentDynamicPrefix string = "{"
var segmentDynamicSuffix string = "}"
var segmentOptionalDynamicPrefix string = "["
var segmentOptionalDynamicSuffix string = "]"
var segmentWildcard string = "*"
var segmentCatchAll string = "***"
