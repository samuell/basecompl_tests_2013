def basecompl(l):
    newl = ""
    for i, c in enumerate(l):
        if c == 'A':
            newl += 'T'
        elif c == 'T':
            newl += 'A'
        elif c == 'C':
            newl += 'G'
        elif c == 'G':
            newl += 'C'
    return newl

f = open("Homo_sapiens.GRCh37.67.dna_rm.chromosome.Y.fa","r")
for l in f:
    l = basecompl(l)
    l = basecompl(l)
    l = basecompl(l)
    l = basecompl(l)
    print l
