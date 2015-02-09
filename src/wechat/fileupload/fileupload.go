/*
 * jQuery File Upload Plugin GAE Go Example 3.2.0
 * https://github.com/blueimp/jQuery-File-Upload
 *
 * Copyright 2011, Sebastian Tschan
 * https://blueimp.net
 *
 * Licensed under the MIT license:
 * http://www.opensource.org/licenses/MIT
 */

package fileupload

import (
	"fmt"

	"net/http"
)

func Fileupload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("文件")
}
