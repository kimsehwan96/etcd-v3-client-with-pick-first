# ETCD

이 디렉터리는 https://github.com/etcd-io/etcd 의 v3.5.12 버전을 clone 하고 필요 없는 파일은 제거한 버전입니다.

원본은 https://github.com/etcd-io/etcd/tree/v3.5.12 여기를 참고하세요. 

우리는 여기서 `etcd/client/v3/internal/resolver/resolver.go` 경로의 43, 44번째 라인의

```go
// Build returns itself for Resolver, because it's both a builder and a resolver.
func (r *EtcdManualResolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r.serviceConfig = cc.ParseServiceConfig(`{"loadBalancingPolicy": "round_robin"}`)
	//r.serviceConfig = cc.ParseServiceConfig(`{"loadBalancingPolicy": "pick_first"}`)
	if r.serviceConfig.Err != nil {
		return nil, r.serviceConfig.Err
	}
	res, err := r.Resolver.Build(target, cc, opts)
	if err != nil {
		return nil, err
	}
	// Populates endpoints stored in r into ClientConn (cc).
	r.updateState()
	return res, nil
}
```
```go
r.serviceConfig = cc.ParseServiceConfig(`{"loadBalancingPolicy": "round_robin"}`) 
```
부분을 `pick_first` 로 바꿔보면서 테스트 해볼 예정입니다.
