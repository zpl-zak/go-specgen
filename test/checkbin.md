## Enum: EPointType

| Value | Description |
| ----- | ----------- |
| Pedestrian = 0x1 |  |
| AI = 0x2 |  |
| Vehicle = 0x4 |  |
| TramStation = 0x8 |  |
| Special = 0x10 |  |
## Enum: ELinkType

| Value | Description |
| ----- | ----------- |
| Pedestrian = 1 |  |
| AI = 2 |  |
| TrainsAndSalinas_Forward = 4 |  |
| TrainsAndSalinas_Reverse = 0x8400 |  |
| Other = 0x1000 |  |

## Spec: Header

| Type | Name | Description |
| ---- | ---- | ----------- |
| uint32 | Magic | Should be 0x1ABCEDF |
| uint32 | PointCount | Number of points |
## Spec: Point

| Type | Name | Description |
| ---- | ---- | ----------- |
| Vector3 | Position |  |
| EPointType | Type |  |
| uint16 | ID |  |
| uint8 | Unk | plain array of 10 elements; Unknown values |
| uint8 | EnterLinks |  |
| uint8 | ExitLinks |  |
## Spec: Link

| Type | Name | Description |
| ---- | ---- | ----------- |
| uint16 | TargetPoint |  |
| ELinkType | LinkType |  |
| float32 | Unk |  |
## Spec: Vector3

| Type | Name | Description |
| ---- | ---- | ----------- |
| float32 | x |  |
| float32 | y |  |
| float32 | z |  |

