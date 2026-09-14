
CREATE TABLE insert_small_block_performance
(
    `x` String, 
    `a` DEFAULT SHA256(x), 
    `b` DEFAULT SHA256(toString(a)), 
    `c` DEFAULT SHA256(toString(b)), 
    `d` DEFAULT SHA256(toString(c)), 
    `e` DEFAULT SHA256(toString(d)), 
    `f` DEFAULT SHA256(toString(e)), 
    `g` DEFAULT SHA256(toString(f)), 
    `h` DEFAULT SHA256(toString(g)), 
    `i` DEFAULT SHA256(toString(h)), 
    `j` DEFAULT SHA256(toString(i)), 
    `k` DEFAULT SHA256(toString(j)), 
    `l` DEFAULT SHA256(toString(k)), 
    `m` DEFAULT SHA256(toString(l)), 
    `n` DEFAULT SHA256(toString(m)), 
    `o` DEFAULT SHA256(toString(n)), 
    `p` DEFAULT SHA256(toString(o)), 
    `q` DEFAULT SHA256(toString(p)), 
    `r` DEFAULT SHA256(toString(q)), 
    `s` DEFAULT SHA256(toString(r)), 
    `t` DEFAULT SHA256(toString(s)), 
    `u` DEFAULT SHA256(toString(t)), 
    `v` DEFAULT SHA256(toString(u)), 
    `w` DEFAULT SHA256(toString(v))
)
ENGINE = Null;
    