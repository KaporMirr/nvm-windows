
	}

	
// 	if ignore {
// 		if !utf8.ValidString(bs) {
// 			v := make([]rune, 0, len(bs))
// 			for i, r := range bs {
// 				if r == utf8.RuneError {
// 					_, size := utf8.DecodeRuneInString(bs[i:])
// 					if size == 1 {
// 						continue
// 					}
// 				}
// 				v = append(v, r)
// 			}
// 			bs = string(v)
// 		}
// 	}

// 	if cs == "UTF-8" {
// 		return bs, nil
// 	}

// 	converter, err := iconv.NewConverter(cs, "UTF-8")
// 	if err != nil {
// 		err = errors.New("Failed to convert " + cs + " to UTF-8: " + err.Error())
// 		return bs, err
// 	}

// 	return converter.ConvertString(bs)
// }
