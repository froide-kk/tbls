package md

import "testing"

// extractLogicalName は論理名の切り出しに正規表現 `^[^:\n\r\s　：]+` を使うため
// タブも区切りとして扱われる。一方で残余コメントの除去側でタブが対象から漏れており、
// タブ区切りのコメントでは説明文の先頭にタブが残っていた。
func TestExtractLogicalName(t *testing.T) {
	tests := []struct {
		name        string
		comment     string
		wantLogical string
		wantComment string
	}{
		{"タブ区切り", "論理名\t説明本文", "論理名", "説明本文"},
		{"タブ+半角スペース区切り", "論理名\t 説明本文", "論理名", "説明本文"},
		{"コロン区切り", "論理名:説明本文", "論理名", "説明本文"},
		{"コロン+半角スペース区切り", "論理名: 説明本文", "論理名", "説明本文"},
		{"全角コロン区切り", "論理名：説明本文", "論理名", "説明本文"},
		{"半角スペース区切り", "論理名 説明本文", "論理名", "説明本文"},
		{"全角スペース区切り", "論理名　説明本文", "論理名", "説明本文"},
		{"改行区切り", "論理名\n説明本文", "論理名", "説明本文"},
		{"区切りなし", "論理名のみ", "論理名のみ", ""},
		{"空文字", "", "", ""},
		{"説明側のコロンは分割しない", "論理名: 比率は 1:2 とする", "論理名", "比率は 1:2 とする"},
		{"説明が複数行", "論理名: 一行目\n二行目", "論理名", "一行目\n二行目"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractLogicalName(tt.comment)
			if got.LogicalName != tt.wantLogical {
				t.Errorf("LogicalName: got %q, want %q", got.LogicalName, tt.wantLogical)
			}
			if got.UpdatedComment != tt.wantComment {
				t.Errorf("UpdatedComment: got %q, want %q", got.UpdatedComment, tt.wantComment)
			}
		})
	}
}
