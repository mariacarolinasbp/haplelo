
#!/usr/bin/env bash

ARQUIVO_ENTRADA="input_acre.inp"

INPUT="/content/drive/My Drive/$ARQUIVO_ENTRADA"

TMP=$(mktemp --directory /tmp/haplelo.XXXXXXXXXX) ||
  { echo "Failed to create temp directory"; exit 1; }

# Install PHASE at /usr/local/bin/phase if not available already
if ! [ -x "$(command -v phase)" ]; then
  wget --output-document=$TMP/phase.tar.gz http://stephenslab.uchicago.edu/assets/software/phase/phasecode/phase.2.1.1.linux.tar.gz
  tar --gzip --extract --directory=$TMP --file=$TMP/phase.tar.gz phase.2.1.1.linux/PHASE
  sudo mv $TMP/phase.2.1.1.linux/PHASE /usr/local/bin/phase # avoid sudo if possible
fi

# Run three phase proccesses in parallel and wait all end
phase -MS -f1 -S$RANDOM "$INPUT" $TMP/res1 400000 1000 50000 &
# phase -MS -f1 -S$RANDOM $INPUT $TMP/res2 400000 1000 50000 &
# phase -MS -f1 -S$RANDOM $INPUT $TMP/res3 400000 1000 50000 &
wait
# TODO: Check if res1 = res2 = res3

sed --expression='/BEGIN LIST_SUMMARY/,/END LIST_SUMMARY/{ /LIST_SUMMARY/d; p }' \
  --quiet $TMP/res1 | column -t | tr -s ' ' | tr ' ' , > $TMP/halelos.csv

sed --expression='/BEGIN BESTPAIRS_SUMMARY/,/END BESTPAIRS_SUMMARY/{ /BESTPAIRS_SUMMARY/d; p }' \
  --quiet $TMP/res1 | tr ':(),' ' ' | sort | column -t | tr -s ' ' | tr ' ' , > $TMP/pacientes.csv

cat <<- "EOF" > $TMP/modelo.csv
GGGCCCCGC,Normal,*1/*1xN,1
GGGCCCTGC,Indeterminada,*34,1
GGGCCCTCC,Normal,*2D/2xN,1
GGGCCGTCC,Normal,*2A/*35/*2xN,1
GGGCCGTCA,Indeterminado,,
GGGCTCTCC,Reduzida,*17/*17xN,0.5
GGGTCCCCC,Reduzida,*10/*10x2,0.5
GGACCCTCC,Reduzida,*29,0.5
GAGCCCTCC,Reduzida,*41,0.5
AGGTCCCCC,Nula,*4/*4x2,0
AGGTTCCCC,Indeterminado,,
GGGCCCCGA,Reduzida,*9,0.5
EOF

echo "Files available at: $TMP"
