## Header

| Type | Name | Description |
| ---- | ---- | ----------- |
| int8 | Magic | plain array of 4 elements; Has to be 'PACK' |
| int32 | DirectoryOffset | Offset to the directory |
| int32 | DirectoryLength | Directory length |
## Directory

| Type | Name | Description |
| ---- | ---- | ----------- |
| int8 | FileName | String consisting of 56 characters; Archived file name |
| int32 | FilePosition |  |
| int32 | FileLength |  |
## Model

| Type | Name | Description |
| ---- | ---- | ----------- |
| uint32 | ModelType | Model type |
| uint16 | FaceGroupCount | Number of face groups |
| FaceGroup | FaceGroups | N definitions of FaceGroup; model's face groups |
## FaceGroup

| Type | Name | Description |
| ---- | ---- | ----------- |
| int32 | MaterialID | -1 for default |
| uint16 | FaceCount |  |
| int64 | Faces | plain array of N elements;  |
