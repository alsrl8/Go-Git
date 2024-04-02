# Go-Git

이 문서는 Go-Git 프로그램에 대한 내용을 다룹니다.

이 프로그램은 Git Repository의 Commit Log를 읽고, 내용을 종합하는 프로그램입니다.

## Config

이 프로그램은 같은 경로의 `gogit.config` 파일을 설정 정보로 사용합니다.
`rootDir` field를 git commit을 조회할 root directory로 지정합니다.

```
{key} = {value}

rootDir = C:\Users\user\
```