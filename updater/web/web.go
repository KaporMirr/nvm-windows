 )

		path := filepath.J)

		}

		if f.FileInfo().IsDir() {
			
			os.ĺ´´´(filepat
 err
			}

				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
