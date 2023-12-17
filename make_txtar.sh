rm -rf /tmp/charmweasel

rm -f /tmp/charmweasel.tar
rm -f /tmp/filelist.txt

{
    rg --files . \
        | grep -vE 'charmweasel$' \
        | grep -v README.org \
        | grep -v make_txtar.sh \
        | grep -v go.sum \
        | grep -v go.mod \
        | grep -v Makefile \
        | grep -v cmd/main.go \
        # | grep -v charmweasel_test.go \
        # | grep -v charmweasel.go \

} | tee /tmp/filelist.txt
tar -cf /tmp/charmweasel.tar -T /tmp/filelist.txt
mkdir -p /tmp/charmweasel
tar xf /tmp/charmweasel.tar -C /tmp/charmweasel
rg --files /tmp/charmweasel
txtar-c /tmp/charmweasel | pbcopy
