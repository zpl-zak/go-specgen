## Enum: ModelTypeEnum

| Value | Description |
| ----- | ----------- |
| Standard = 1 |  Standard mesh type |
| SingleMesh = 4 |  Mesh consisting of bones and skin data |
| Morph |  |
| Sector |  Holds objects |
## Enum: MaterialFlags

| Value | Description |
| ----- | ----------- |
| TEXTUREDIFFUSE = 0x00040000 |  whether diffuse texture is present |
| COLORED = 0x08000000 |  whether to use diffuse color (only applies with diffuse texture) |
| MIPMAPPING = 0x00800000 |  |
| ANIMATEDTEXTUREDIFFUSE = 0x04000000 |  |
| ANIMATEXTEXTUREALPHA = 0x02000000 |  |
| DOUBLESIDEDMATERIAL = 0x10000000 |  whether backface culling should be off |
| ENVIRONMENTMAP = 0x00080000 |  simulates glossy material with environment texture |
| NORMALTEXTUREBLEND = 0x00000100 |  blend between diffuse and environment texture normally |
| MULTIPLYTEXTUREBLEND = 0x00000200 |  blend between diffuse and environment texture by multiplying |
| ADDITIVETEXTUREBLEND = 0x00000400 |  blend between diffuse and environment texture by addition |
| CALCREFLECTTEXTUREY = 0x00001000 |  |
| PROJECTREFLECTTEXTUREY = 0x00002000 |  |
| PROJECTREFLECTTEXTUREZ = 0x00004000 |  |
| ADDITIONALEFFECT = 0x00008000 |  should be ALPHATEXTURE | COLORKEY | ADDITIVEMIXING |
| ALPHATEXTURE = 0x40000000 |  |
| COLORKEY = 0x20000000 |  |
| ADDITIVEMIXING = 0x80000000 |  the object is blended against the world by adding RGB (see street lamps etc.) |

## Spec: Header

| Type | Name | Description |
| ---- | ---- | ----------- |
| int8 | Magic | plain array of 4 elements; Has to be 'PACK' |
| int32 | DirectoryOffset | Offset to the directory |
| int32 | DirectoryLength | Directory length |
## Spec: Directory

| Type | Name | Description |
| ---- | ---- | ----------- |
| int8 | FileName | String consisting of 56 characters; Archived file name |
| int32 | FilePosition |  |
| int32 | FileLength |  |
## Spec: Model

| Type | Name | Description |
| ---- | ---- | ----------- |
| ModelTypeEnum | ModelType | Model type |
| uint16 | FaceGroupCount | Number of face groups |
| FaceGroup | FaceGroups | N definitions of FaceGroup; model's face groups |
## Spec: FaceGroup

| Type | Name | Description |
| ---- | ---- | ----------- |
| int32 | MaterialID | -1 for default |
| uint16 | FaceCount |  |
| int64 | Faces | plain array of N elements;  |

