package main

import (
	"archive/tar"
	"io"
	"path"
	"compress/gzip"
	"compress/bzip2"
	"github.com/ulikunitz/xz"
	"errors"
	"io/fs"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		println("No archives given, exiting...")
		os.Exit(0)
	}
	for _, archive := range os.Args[1:] {
		archiveFile, err := os.OpenFile(archive, os.O_RDONLY, 0755)
		if errors.Is(err, fs.ErrNotExist) {
			println("Archive "+archive+" not found. Skipping...")
			continue
		}
		if err != nil {
			panic(err)
		}

		var baseName string
		switch(path.Ext(archive)) {
		case ".gz":
			reader, err := gzip.NewReader(archiveFile)
			if err == io.EOF {
				println("Archive "+archive+" empty. Skipping")
				continue
			}
			if err != nil {
				panic(err)
			}
			baseName = strings.TrimSuffix(archive, ".tar.gz")

			tarReader := tar.NewReader(reader)
			hdr, err := tarReader.Next()
			if err == io.EOF {
				println("Archive "+archive+" empty. Skipping...")
				continue
			}
			if hdr.Typeflag == tar.TypeDir {
				if err = os.Mkdir(hdr.Name, hdr.FileInfo().Mode().Perm()); err != nil {
					println("Couldn't make "+hdr.Name+" folder for archive "+archive+". Skipping...")
					continue
				}
				for {
					hdr, err = tarReader.Next()
					if err == io.EOF {
						break
					}
					if err != nil {
						panic(err)
					}
					if hdr.Typeflag == tar.TypeDir {
						if err = os.Mkdir(hdr.Name, hdr.FileInfo().Mode().Perm()); err != nil {
							println("Couldn't create directory "+hdr.Name+". Skipping this archive...")
							continue
						}
					} else {
						outFile, err := os.OpenFile(hdr.Name, os.O_WRONLY | os.O_CREATE, hdr.FileInfo().Mode().Perm())
						if err != nil {
							println("Couldn't open file "+hdr.Name+" in archive. Skipping this archive...")
							continue
						}
						if _, err = io.Copy(outFile, tarReader); err != nil {
							outFile.Close()
							panic(err)
						}
						outFile.Close()
					}
				}
			} else {
				if err = os.Mkdir(baseName, 0755); err != nil {
					panic(err)
				}
				outFile, err := os.OpenFile(path.Join(baseName, hdr.Name), os.O_WRONLY | os.O_CREATE, hdr.FileInfo().Mode().Perm())
				if err != nil {
					println("Couldn't open file "+path.Join(baseName, hdr.Name)+" in archive. Skipping this archive...")
					continue
				}
				if _, err = io.Copy(outFile, tarReader); err != nil {
					outFile.Close()
					panic(err)
				}
				outFile.Close()
				for {
					hdr, err = tarReader.Next()
					if err == io.EOF {
						break
					}
					if err != nil {
						panic(err)
					}
					if hdr.Typeflag == tar.TypeDir {
						if err = os.Mkdir(path.Join(baseName, hdr.Name), hdr.FileInfo().Mode().Perm()); err != nil {
							println("Couldn't create directory "+path.Join(baseName, hdr.Name)+". Skipping this archive...")
							continue
						}
					} else {
						outFile, err = os.OpenFile(path.Join(baseName, hdr.Name), os.O_WRONLY | os.O_CREATE, hdr.FileInfo().Mode().Perm())
						if err != nil {
							println("Couldn't open file "+path.Join(baseName, hdr.Name)+" in archive. Skipping this archive...")
							continue
						}
						if _, err = io.Copy(outFile, tarReader); err != nil {
							outFile.Close()
							panic(err)
						}
						outFile.Close()
					}
				}
			}
		case ".bz2":
			reader := bzip2.NewReader(archiveFile)
			baseName = strings.TrimSuffix(archive, ".tar.bz2")
			tarReader := tar.NewReader(reader)
			hdr, err := tarReader.Next()
			if err == io.EOF {
				println("Archive "+archive+" empty. Skipping...")
				continue
			}
			if hdr.Typeflag == tar.TypeDir {
				if err = os.Mkdir(hdr.Name, hdr.FileInfo().Mode().Perm()); err != nil {
					println("Couldn't make "+hdr.Name+" folder for archive "+archive+". Skipping...")
					continue
				}
				for {
					hdr, err = tarReader.Next()
					if err == io.EOF {
						break
					}
					if err != nil {
						panic(err)
					}
					if hdr.Typeflag == tar.TypeDir {
						if err = os.Mkdir(hdr.Name, hdr.FileInfo().Mode().Perm()); err != nil {
							println("Couldn't create directory "+hdr.Name+". Skipping this archive...")
							continue
						}
					} else {
						outFile, err := os.OpenFile(hdr.Name, os.O_WRONLY | os.O_CREATE, hdr.FileInfo().Mode().Perm())
						if err != nil {
							println("Couldn't open file "+hdr.Name+" in archive. Skipping this archive...")
							continue
						}
						if _, err = io.Copy(outFile, tarReader); err != nil {
							outFile.Close()
							panic(err)
						}
						outFile.Close()
					}
				}
			} else {
				if err = os.Mkdir(baseName, 0755); err != nil {
					panic(err)
				}
				outFile, err := os.OpenFile(path.Join(baseName, hdr.Name), os.O_WRONLY | os.O_CREATE, hdr.FileInfo().Mode().Perm())
				if err != nil {
					println("Couldn't open file "+path.Join(baseName, hdr.Name)+" in archive. Skipping this archive...")
					continue
				}
				if _, err = io.Copy(outFile, tarReader); err != nil {
					outFile.Close()
					panic(err)
				}
				outFile.Close()
				for {
					hdr, err = tarReader.Next()
					if err == io.EOF {
						break
					}
					if err != nil {
						panic(err)
					}
					if hdr.Typeflag == tar.TypeDir {
						if err = os.Mkdir(path.Join(baseName, hdr.Name), hdr.FileInfo().Mode().Perm()); err != nil {
							println("Couldn't create directory "+path.Join(baseName, hdr.Name)+". Skipping this archive...")
							continue
						}
					} else {
						outFile, err = os.OpenFile(path.Join(baseName, hdr.Name), os.O_WRONLY | os.O_CREATE, hdr.FileInfo().Mode().Perm())
						if err != nil {
							println("Couldn't open file "+path.Join(baseName, hdr.Name)+" in archive. Skipping this archive...")
							continue
						}
						if _, err = io.Copy(outFile, tarReader); err != nil {
							outFile.Close()
							panic(err)
						}
						outFile.Close()
					}
				}
			}
		case ".xz":
			reader, err := xz.NewReader(archiveFile)
			if err != nil {
				panic(err)
			}
			baseName = strings.TrimSuffix(archive, ".tar.xz")
			tarReader := tar.NewReader(reader)
			hdr, err := tarReader.Next()
			if err == io.EOF {
				println("Archive "+archive+" empty. Skipping...")
				continue
			}
			if hdr.Typeflag == tar.TypeDir {
				if err = os.Mkdir(hdr.Name, hdr.FileInfo().Mode().Perm()); err != nil {
					println("Couldn't make "+hdr.Name+" folder for archive "+archive+". Skipping...")
					continue
				}
				for {
					hdr, err = tarReader.Next()
					if err == io.EOF {
						break
					}
					if err != nil {
						panic(err)
					}
					if hdr.Typeflag == tar.TypeDir {
						if err = os.Mkdir(hdr.Name, hdr.FileInfo().Mode().Perm()); err != nil {
							println("Couldn't create directory "+hdr.Name+". Skipping this archive...")
							continue
						}
					} else {
						outFile, err := os.OpenFile(hdr.Name, os.O_WRONLY | os.O_CREATE, hdr.FileInfo().Mode().Perm())
						if err != nil {
							println("Couldn't open file "+hdr.Name+" in archive. Skipping this archive...")
							continue
						}
						if _, err = io.Copy(outFile, tarReader); err != nil {
							outFile.Close()
							panic(err)
						}
						outFile.Close()
					}
				}
			} else {
				if err = os.Mkdir(baseName, 0755); err != nil {
					panic(err)
				}
				outFile, err := os.OpenFile(path.Join(baseName, hdr.Name), os.O_WRONLY | os.O_CREATE, hdr.FileInfo().Mode().Perm())
				if err != nil {
					println("Couldn't open file "+path.Join(baseName, hdr.Name)+" in archive. Skipping this archive...")
					continue
				}
				if _, err = io.Copy(outFile, tarReader); err != nil {
					outFile.Close()
					panic(err)
				}
				outFile.Close()
				for {
					hdr, err = tarReader.Next()
					if err == io.EOF {
						break
					}
					if err != nil {
						panic(err)
					}
					if hdr.Typeflag == tar.TypeDir {
						if err = os.Mkdir(path.Join(baseName, hdr.Name), hdr.FileInfo().Mode().Perm()); err != nil {
							println("Couldn't create directory "+path.Join(baseName, hdr.Name)+". Skipping this archive...")
							continue
						}
					} else {
						outFile, err = os.OpenFile(path.Join(baseName, hdr.Name), os.O_WRONLY | os.O_CREATE, hdr.FileInfo().Mode().Perm())
						if err != nil {
							println("Couldn't open file "+path.Join(baseName, hdr.Name)+" in archive. Skipping this archive...")
							continue
						}
						if _, err = io.Copy(outFile, tarReader); err != nil {
							outFile.Close()
							panic(err)
						}
						outFile.Close()
					}
				}
			}
		}

	}
}
