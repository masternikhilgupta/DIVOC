package services

import (
	"encoding/json"
	"testing"
)

func TestMosipSignature(t *testing.T) {
	otpReq := OTPRequest{
		Id:               "mosip.identity.otp",
		Version:          "1.0",
		TransactionId:    "1234567890",
		RequestTime:      "2022-05-11T18:00:37.914Z",
		IndividualId:     "1234567",
		IndividualIdType: "VID",
		OtpChannel:       []string{"email"},
	}
	otpReqStr, _ := json.Marshal(otpReq)

	// dummy keys for testing
	publicKeyPem := "-----BEGIN CERTIFICATE-----\nMIIFcDCCA1gCCQDImfyNShfG0DANBgkqhkiG9w0BAQsFADB6MQswCQYDVQQGEwJJTjESMBAGA1UECAwJS2FybmF0YWthMRIwEAYDVQQHDAlCYW5nYWxvcmUxDjAMBgNVBAoMBURJVk9DMQowCAYDVQQLDAEuMScwJQYDVQQDDB5odHRwczovL2RlbW8tZGl2b2MuZWdvdi5vcmcuaW4wHhcNMjIwNTExMTc1NTE3WhcNMjMwNTExMTc1NTE3WjB6MQswCQYDVQQGEwJJTjESMBAGA1UECAwJS2FybmF0YWthMRIwEAYDVQQHDAlCYW5nYWxvcmUxDjAMBgNVBAoMBURJVk9DMQowCAYDVQQLDAEuMScwJQYDVQQDDB5odHRwczovL2RlbW8tZGl2b2MuZWdvdi5vcmcuaW4wggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQDJuEFJNAPPDT8Sj38H5dSkliSZlmm6YOqKfPJ5WMbY9qZ5Lx8DoDr5Q/H11Dnoi0Ik89BrrmzUDUUvtuz0byM0oiw/qmQpnRZGcBZ23gH/9vCyRPWShqhlc9wmUfMgNtEqWbxSQlGK6JTriV5txW21Fh46HZvoAlS3f4v8BW3dfR8/6GgyLKdKpFTXSjEoEMmyZZRh5u7aUymCdO1gtr7wZXzmsvCe+LshwS3uaCMUQ8yhsVPtuLL+9ZHmoXcrOvjS60oymByHyFK95nP6uXnW9n7AtsFElgxYkph1xbX8ow+kwRq+3QO/lrs20ztMSeC/vUIw5HJfFeahtbTaAxGAC3RM0EaMITd9HxNtc+bsQiAOWNN4ED8aeZqQ1gSCVKr1hOeHZE7PqO+t36JhgHKDfzyP4rSQeH21A2hYNfwvbKtKsE2kkvcQnjgyyHK2OPtuWalReDvYb9Z7prYdgN1+tm3/Fm9xqP+uHWobcJRJSu78psvMjFRTASiVIh8Q5/Lx9oRMeGKzStG5J+ixOB2Cd/lg8aOX0Ac0zhkcZj1qdQk0eEhTUmh0dmUzARWA1kQ9OlfruPJ/45p0cqpRVLUQZKnhFrC+YSw8ODno3IXSyOimCwNtGt6AXt/qW0ZX20z7soI3dLRtxUGIaonwScNcUFE+uu0a5ho3aZX28UsdjwIDAQABMA0GCSqGSIb3DQEBCwUAA4ICAQC7mBftpG+Ay92KX07HKgJvEt3UHogMK7ofENgWe4b7fvqOtHde8fPotiITJLrCjCHlyaoueOQ8IMRDzIEH+mz9s6NEQPGX70kD2OIrFFunvVmKLCiF3sSyUAq1Qgm2aThB238MkTnEC6C2EBHebz5ncolmVIpWTusExK0eEtw9oRDeIo9jioyrZW8Dt97ND9n1SngsxgcjH1kqGOROwVbgqOqlKrjueT9N8Jbo+Cg57zZZ4npsbmjb21KQXKq2X5RUR95p8saUuZ6GLh/gFVmnFm66qG3HzjnJAXyk9WH+pgp6Pt8HnrCN0jbT2LE0RWh/xMy8wACFihFI3m9TmS7k8ZH6CfAwruK4obvEJoIWNNFq+MtM44l1ydl//726FSp0Le0JYEXsO8O1RJQ39n2IglOs36WGuT1dZ8PBk3gJ/lo+N1Mlro6HGjkrzasykFd3RPTUzVCF+Zd+I67NFOmbjgdRsica4I3h2A4xM1PfOUtbb5TCZeLBLkcW4YKfjHYQ5v728ZhpjdbEarcnuKU34akSMvOinYXIG70Rt3PVjXTl0RW53sQXhTu0+q3l5ccvAQm3Gn/q1z2ofnCFhb5y3a1Hop9IaY2QtdAthTpBe91XT3YnL3i1eQ6mzFDPmQUB+4GNs3/zZHX3M/oOtie1UI560a2u+k4AUn7kShAwdg==\n-----END CERTIFICATE-----";
	privateKeyPem := "-----BEGIN PRIVATE KEY-----\n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 GnEg4hAWcUpapOWBA4lhzKGcBZ2TPw==\n-----END PRIVATE KEY-----";

	type args struct {
		data []byte
		privateKey string
		publicKey string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "shouldSign",
			args: args{
				otpReqStr,
				privateKeyPem,
				publicKeyPem,
			},
			want: "eyJhbGciOiJSUzI1NiIsIng1YyI6WyJNSUlGY0RDQ0ExZ0NDUURJbWZ5TlNoZkcwREFOQmdrcWhraUc5dzBCQVFzRkFEQjZNUXN3Q1FZRFZRUUdFd0pKVGpFU01CQUdBMVVFQ0F3SlMyRnlibUYwWVd0aE1SSXdFQVlEVlFRSERBbENZVzVuWVd4dmNtVXhEakFNQmdOVkJBb01CVVJKVms5RE1Rb3dDQVlEVlFRTERBRXVNU2N3SlFZRFZRUUREQjVvZEhSd2N6b3ZMMlJsYlc4dFpHbDJiMk11WldkdmRpNXZjbWN1YVc0d0hoY05Nakl3TlRFeE1UYzFOVEUzV2hjTk1qTXdOVEV4TVRjMU5URTNXakI2TVFzd0NRWURWUVFHRXdKSlRqRVNNQkFHQTFVRUNBd0pTMkZ5Ym1GMFlXdGhNUkl3RUFZRFZRUUhEQWxDWVc1bllXeHZjbVV4RGpBTUJnTlZCQW9NQlVSSlZrOURNUW93Q0FZRFZRUUxEQUV1TVNjd0pRWURWUVFEREI1b2RIUndjem92TDJSbGJXOHRaR2wyYjJNdVpXZHZkaTV2Y21jdWFXNHdnZ0lpTUEwR0NTcUdTSWIzRFFFQkFRVUFBNElDRHdBd2dnSUtBb0lDQVFESnVFRkpOQVBQRFQ4U2ozOEg1ZFNrbGlTWmxtbTZZT3FLZlBKNVdNYlk5cVo1THg4RG9EcjVRL0gxMURub2kwSWs4OUJycm16VURVVXZ0dXowYnlNMG9pdy9xbVFwblJaR2NCWjIzZ0gvOXZDeVJQV1NocWhsYzl3bVVmTWdOdEVxV2J4U1FsR0s2SlRyaVY1dHhXMjFGaDQ2SFp2b0FsUzNmNHY4QlczZGZSOC82R2d5TEtkS3BGVFhTakVvRU1teVpaUmg1dTdhVXltQ2RPMWd0cjd3Wlh6bXN2Q2UrTHNod1MzdWFDTVVROHloc1ZQdHVMTCs5Wkhtb1hjck92alM2MG95bUJ5SHlGSzk1blA2dVhuVzluN0F0c0ZFbGd4WWtwaDF4Ylg4b3cra3dScSszUU8vbHJzMjB6dE1TZUMvdlVJdzVISmZGZWFodGJUYUF4R0FDM1JNMEVhTUlUZDlIeE50Yytic1FpQU9XTk40RUQ4YWVacVExZ1NDVktyMWhPZUhaRTdQcU8rdDM2SmhnSEtEZnp5UDRyU1FlSDIxQTJoWU5md3ZiS3RLc0Uya2t2Y1Fuamd5eUhLMk9QdHVXYWxSZUR2WWI5WjdwcllkZ04xK3RtMy9GbTl4cVArdUhXb2JjSlJKU3U3OHBzdk1qRlJUQVNpVkloOFE1L0x4OW9STWVHS3pTdEc1SitpeE9CMkNkL2xnOGFPWDBBYzB6aGtjWmoxcWRRazBlRWhUVW1oMGRtVXpBUldBMWtROU9sZnJ1UEovNDVwMGNxcFJWTFVRWktuaEZyQytZU3c4T0RubzNJWFN5T2ltQ3dOdEd0NkFYdC9xVzBaWDIwejdzb0kzZExSdHhVR0lhb253U2NOY1VGRSt1dTBhNWhvM2FaWDI4VXNkandJREFRQUJNQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUNBUUM3bUJmdHBHK0F5OTJLWDA3SEtnSnZFdDNVSG9nTUs3b2ZFTmdXZTRiN2Z2cU90SGRlOGZQb3RpSVRKTHJDakNIbHlhb3VlT1E4SU1SRHpJRUgrbXo5czZORVFQR1g3MGtEMk9JckZGdW52Vm1LTENpRjNzU3lVQXExUWdtMmFUaEIyMzhNa1RuRUM2QzJFQkhlYno1bmNvbG1WSXBXVHVzRXhLMGVFdHc5b1JEZUlvOWppb3lyWlc4RHQ5N05EOW4xU25nc3hnY2pIMWtxR09ST3dWYmdxT3FsS3JqdWVUOU44SmJvK0NnNTd6Wlo0bnBzYm1qYjIxS1FYS3EyWDVSVVI5NXA4c2FVdVo2R0xoL2dGVm1uRm02NnFHM0h6am5KQVh5azlXSCtwZ3A2UHQ4SG5yQ04wamJUMkxFMFJXaC94TXk4d0FDRmloRkkzbTlUbVM3azhaSDZDZkF3cnVLNG9idkVKb0lXTk5GcStNdE00NGwxeWRsLy83MjZGU3AwTGUwSllFWHNPOE8xUkpRMzluMklnbE9zMzZXR3VUMWRaOFBCazNnSi9sbytOMU1scm82SEdqa3J6YXN5a0ZkM1JQVFV6VkNGK1pkK0k2N05GT21iamdkUnNpY2E0STNoMkE0eE0xUGZPVXRiYjVUQ1plTEJMa2NXNFlLZmpIWVE1djcyOFpocGpkYkVhcmNudUtVMzRha1NNdk9pbllYSUc3MFJ0M1BWalhUbDBSVzUzc1FYaFR1MCtxM2w1Y2N2QVFtM0duL3ExejJvZm5DRmhiNXkzYTFIb3A5SWFZMlF0ZEF0aFRwQmU5MVhUM1luTDNpMWVRNm16RkRQbVFVQis0R05zMy96WkhYM00vb090aWUxVUk1NjBhMnUrazRBVW43a1NoQXdkZz09Il19.eyJpZCI6Im1vc2lwLmlkZW50aXR5Lm90cCIsInZlcnNpb24iOiIxLjAiLCJ0cmFuc2FjdGlvbklEIjoiMTIzNDU2Nzg5MCIsInJlcXVlc3RUaW1lIjoiMjAyMi0wNS0xMVQxODowMDozNy45MTRaIiwiaW5kaXZpZHVhbElkIjoiMTIzNDU2NyIsImluZGl2aWR1YWxJZFR5cGUiOiJWSUQiLCJvdHBDaGFubmVsIjpbImVtYWlsIl19.mEJyGITU8Tl3dswDPdVpQArhCE8jMv5z7Hts5XTli-PGw47Ins0iXrl4Oys1TJmfXpDrLi-J4umQTHTP2_E4LmUFwxnDNHL4qEEUlBP4S7a6RhOjozuLMjxnv1HTLlyXJi0wFkDCjkjrV4T3VP-2UEr8ucwHotgxwgVvzbWiQqEe_JVpETtXDTUWDCm_CQ7ZzDrgPaiMbHdtT7PDBKnDuZoS98rGZilMLTZjg2qPN9InggV6zHpH5DWxnbWgpz0120u2I-BH1dOvi54gGiew8oF7erjnhsaMQsbJaKde2nP3a2O012BbdilKnLz8Uxz52P90ZOpPJU3uMbQmGWUdbO2rGmmZnuOd8qQSovycYwwLlRfO2jF1SZXut7d63Rww4hPoUqls6grgfnRrjRF7K4HB-D1V0BwOngRPyLI_0Kdm18RDpgo4LHA4xhceipK7BWctlBtfDQVOCFHgjSp4LyQUVOI0_7bENzpazbW5os69AUyCQx5EgRtAXwwwG7y_wHYdqdVxV_YU6kngoH9n0vz01aapY1tUCC0pq2TQDTIkNpxRnB2zEKT1q0bdDeYvbla16mviROTSOPL55ruebNIlaV7-aKdwES8zJ1n0LJv8ahfDWOO5-zzxjkql9gXQf3TuLIBsbhm7DD0WLr-okvHM-R6IJ7XxsNvt-72uQTE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSignature(tt.args.data, tt.args.privateKey, tt.args.publicKey); got != tt.want {
				t.Errorf("getSignature() = %v, want %v", got, tt.want)
			}
		})
	}
}

