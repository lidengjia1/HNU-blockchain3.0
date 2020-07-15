# 附录 1. 密码学专题 - 大数

## 附录 1. 密码学专题 - 大数

表 1-1 大数

| 物理模拟量                                             | 大数                                                         |
| ------------------------------------------------------ | ------------------------------------------------------------ |
| 每天被闪电杀死的可能性                                 | 90 亿 (![2^{33}](https://math.jianshu.com/math?formula=2%5E%7B33%7D)) 分之一 |
| 赢得国家发行彩票头等奖的可能性                         | 400 万 (![2^{22}](https://math.jianshu.com/math?formula=2%5E%7B22%7D)) 分之一 |
| 赢得国家发行彩票头等奖励并且在同一天被闪电杀死的可能性 | ![2^{55}](https://math.jianshu.com/math?formula=2%5E%7B55%7D) 分之一 |
| 每年淹死的可能性                                       | 59 000 (![2^{16}](https://math.jianshu.com/math?formula=2%5E%7B16%7D)) 分之一 |
| 1993 年在美国交通事故中死亡的可能性                    | 6100 (![2^{13}](https://math.jianshu.com/math?formula=2%5E%7B13%7D)) 分之一 |
| 在美国死于交通事故的可能性                             | 88 (![2^7](https://math.jianshu.com/math?formula=2%5E7)) 分之一 |
| 到下一个冰川年代的时间                                 | 14 000 (![2^{14}](https://math.jianshu.com/math?formula=2%5E%7B14%7D)) 分之一 |
| 到太阳变成新星的时间                                   | ![10^9](https://math.jianshu.com/math?formula=10%5E9) (![2^{30}](https://math.jianshu.com/math?formula=2%5E%7B30%7D)) 年 |
| 行星的年龄                                             | ![10^9](https://math.jianshu.com/math?formula=10%5E9) (![2^{30}](https://math.jianshu.com/math?formula=2%5E%7B30%7D)) 年 |
| 宇宙的年龄                                             | ![10^{10}](https://math.jianshu.com/math?formula=10%5E%7B10%7D) (![2^{34}](https://math.jianshu.com/math?formula=2%5E%7B34%7D)) 年 |
| 行星中的原子数                                         | ![10^{51}](https://math.jianshu.com/math?formula=10%5E%7B51%7D) (![2^{170}](https://math.jianshu.com/math?formula=2%5E%7B170%7D)) |
| 太阳中的原子数                                         | ![10^{57}](https://math.jianshu.com/math?formula=10%5E%7B57%7D) (![2^{190}](https://math.jianshu.com/math?formula=2%5E%7B190%7D)) |
| 银河系中的原子数                                       | ![10^{67}](https://math.jianshu.com/math?formula=10%5E%7B67%7D) (![2^{223}](https://math.jianshu.com/math?formula=2%5E%7B223%7D)) |
| 宇宙中的原子数                                         | ![10^{77}](https://math.jianshu.com/math?formula=10%5E%7B77%7D) (![2^{265}](https://math.jianshu.com/math?formula=2%5E%7B265%7D)) |

## 项目源代码

项目源代码会逐步上传到 Github，地址为 [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)。

## Contributor

1. Windstamp, [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)



作者：furnace
链接：https://www.jianshu.com/p/c1bccf6b9ce3
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。







# 附录 2. 密码学专题 - 数学知识

# 密码学专题 - 数学知识

## 2. 数论

这里仅列出一些对密码学有用的思想，关于数论更详细的知识请参考专业文献。

### 2.1 模运算

本质上，如果 ![a = b + kn](https://math.jianshu.com/math?formula=a%20%3D%20b%20%2B%20kn) 对某些整数 ![k](https://math.jianshu.com/math?formula=k) 成立，那么 ![a \equiv b \ (mod \ n)](https://math.jianshu.com/math?formula=a%20%5Cequiv%20b%20%5C%20(mod%20%5C%20n))。如果 ![a](https://math.jianshu.com/math?formula=a) 为正，![b](https://math.jianshu.com/math?formula=b) 为 0 ~ n，那么你可将 ![b](https://math.jianshu.com/math?formula=b) 看做 ![a](https://math.jianshu.com/math?formula=a) 被 ![n](https://math.jianshu.com/math?formula=n) 整除后的余数。有时 ![b](https://math.jianshu.com/math?formula=b) 叫做 ![a](https://math.jianshu.com/math?formula=a) 模 ![n](https://math.jianshu.com/math?formula=n) 的余数 (residue)。有时 ![a](https://math.jianshu.com/math?formula=a) 叫做与 ![b](https://math.jianshu.com/math?formula=b) 模 ![n](https://math.jianshu.com/math?formula=n) 同余 (congruent) (三元等号 ![\equiv](https://math.jianshu.com/math?formula=%5Cequiv) 表示同余)。

从 ![0 \sim n-1](https://math.jianshu.com/math?formula=0%20%5Csim%20n-1) 的整数组成的集合构成了模 ![n](https://math.jianshu.com/math?formula=n) 的完全剩余集 (complete set of residue)。这意味着，对于每一个整数 ![a](https://math.jianshu.com/math?formula=a)，它的模 ![n](https://math.jianshu.com/math?formula=n) 的余项是从 ![0 \sim n-1](https://math.jianshu.com/math?formula=0%20%5Csim%20n-1) 的某个数。

![a](https://math.jianshu.com/math?formula=a) 模 ![n](https://math.jianshu.com/math?formula=n) 的运算给出了 ![a](https://math.jianshu.com/math?formula=a) 的余数，余数是从 ![0 \sim n-1](https://math.jianshu.com/math?formula=0%20%5Csim%20n-1) 的某个整数，这种运算称为模变换 (modular reduction)。例如，![5 \ mod \ 3 \ = 2](https://math.jianshu.com/math?formula=5%20%5C%20mod%20%5C%203%20%5C%20%3D%202)。

模运算就像普通运算一样，它是可交换的、可结合的和可分配的。而且，简化每一个中间结果的模 ![n](https://math.jianshu.com/math?formula=n) 运算，其作用与先进行全部运算再简化模 ![n](https://math.jianshu.com/math?formula=n) 运算是一样的。
 ![(a+b) mod \ n = ((a \ mod \ n) + (b \ mod \ n)) mod \ n](https://math.jianshu.com/math?formula=(a%2Bb)%20mod%20%5C%20n%20%3D%20((a%20%5C%20mod%20%5C%20n)%20%2B%20(b%20%5C%20mod%20%5C%20n))%20mod%20%5C%20n)

![(a-b) mod \ n = ((a \ mod \ n) - (b \ mod \ n)) mod \ n](https://math.jianshu.com/math?formula=(a-b)%20mod%20%5C%20n%20%3D%20((a%20%5C%20mod%20%5C%20n)%20-%20(b%20%5C%20mod%20%5C%20n))%20mod%20%5C%20n)

![(a \times b) mod \ n = ((a \ mod \ n) \times (b \ mod \ n)) mod \ n](https://math.jianshu.com/math?formula=(a%20%5Ctimes%20b)%20mod%20%5C%20n%20%3D%20((a%20%5C%20mod%20%5C%20n)%20%5Ctimes%20(b%20%5C%20mod%20%5C%20n))%20mod%20%5C%20n)

密码学用了许多模 ![n](https://math.jianshu.com/math?formula=n) 运算，因为像计算离散对数和平方根这样的问题很困难，而模运算可将所有中间结果和最后结果限制在一个范围内，所以用它进行计算比较容易。对一个 ![k](https://math.jianshu.com/math?formula=k) 位的模数 ![n](https://math.jianshu.com/math?formula=n)，任何加、减、乘的中间结果将不会超过 ![2k](https://math.jianshu.com/math?formula=2k) 位长。因此可以用模运算进行指数运算而又不会产生巨大的中间结果。虽然计算某数的乘方并对其取模的运算
 ![a^x \ mod \ n](https://math.jianshu.com/math?formula=a%5Ex%20%5C%20mod%20%5C%20n)

将导致一系列的乘法和除法运算，但有加速运算的方法：一种方法指在最小化模乘法运算的数量；另一种旨在优化单个模乘法运算。因为操作步骤划分后，当完成一串乘法，并且每次都进行模运算后，指数运算就更快，这样就与一般取幂没有多大差别，但当用 200 位的数字进行运行时，情况就不同了。

例如，如果要计算 ![a^8 \ mod \ n](https://math.jianshu.com/math?formula=a%5E8%20%5C%20mod%20%5C%20n)，不要直接进行七次乘法和一个大数的模化简：
 ![(a \times a \times a \times a \times a \times a \times a \times a) mod \ n](https://math.jianshu.com/math?formula=(a%20%5Ctimes%20a%20%5Ctimes%20a%20%5Ctimes%20a%20%5Ctimes%20a%20%5Ctimes%20a%20%5Ctimes%20a%20%5Ctimes%20a)%20mod%20%5C%20n)

相反，应进行三次较小的乘法和三次较小的模化简：
 ![((a^2 mod \ n)^2 mod \ n)^2 mod \ n](https://math.jianshu.com/math?formula=((a%5E2%20mod%20%5C%20n)%5E2%20mod%20%5C%20n)%5E2%20mod%20%5C%20n)

以此类推，
 ![a^{16} mod \ n = (((a^2 mod \ n)^2 mod \ n)^2 mod \ n)^2 mod \ n](https://math.jianshu.com/math?formula=a%5E%7B16%7D%20mod%20%5C%20n%20%3D%20(((a%5E2%20mod%20%5C%20n)%5E2%20mod%20%5C%20n)%5E2%20mod%20%5C%20n)%5E2%20mod%20%5C%20n)

当 ![x](https://math.jianshu.com/math?formula=x) 不是 2 的幂次方时，计算 ![a^x \ mod \ n](https://math.jianshu.com/math?formula=a%5Ex%20%5C%20mod%20%5C%20n) 稍微要难些。可将 ![x](https://math.jianshu.com/math?formula=x) 表示成 2 的幂次方之和：在二进制中，25 是 11001，因此 ![25=2^4 + 2^3 + 2^0](https://math.jianshu.com/math?formula=25%3D2%5E4%20%2B%202%5E3%20%2B%202%5E0)。故：
 ![a^25 \ mod \ n = (a \times a^{24})mod \ n = (a \times a^8 \times a^{16})mod \ n = (a \times ((a^2)^2)^2) \times (((a^2)^2)^2)^2) mod \ n = ((((a^2 \times a)^2)^2)^2 \times a) mod \ n](https://math.jianshu.com/math?formula=a%5E25%20%5C%20mod%20%5C%20n%20%3D%20(a%20%5Ctimes%20a%5E%7B24%7D)mod%20%5C%20n%20%3D%20(a%20%5Ctimes%20a%5E8%20%5Ctimes%20a%5E%7B16%7D)mod%20%5C%20n%20%3D%20(a%20%5Ctimes%20((a%5E2)%5E2)%5E2)%20%5Ctimes%20(((a%5E2)%5E2)%5E2)%5E2)%20mod%20%5C%20n%20%3D%20((((a%5E2%20%5Ctimes%20a)%5E2)%5E2)%5E2%20%5Ctimes%20a)%20mod%20%5C%20n)

注意，上面的公式利用了，![x^n \times y^n = (xy)^n](https://math.jianshu.com/math?formula=x%5En%20%5Ctimes%20y%5En%20%3D%20(xy)%5En)。

适当利用存储的中间结果，只需要 6 次乘法：
 ![(((((((a^2 \ mod \ n) \times a)mod \ n)^2 mod \ n)^2 mod \ n)^2 mod \ n) \times a) mod \ n](https://math.jianshu.com/math?formula=(((((((a%5E2%20%5C%20mod%20%5C%20n)%20%5Ctimes%20a)mod%20%5C%20n)%5E2%20mod%20%5C%20n)%5E2%20mod%20%5C%20n)%5E2%20mod%20%5C%20n)%20%5Ctimes%20a)%20mod%20%5C%20n)

这种算法称为加法链 (addition chaining)，或二进制平方和乘法方法。它用二进制表示了一个简单明了的加法链。算法的 C 语言描述如下：



```c
unsigned long qe2(unsigned long x, unsigned long y, unsigned long n) {
    unsigned long s, t, u;
    int i;

    s = 1; t = x; u = y;
    while (u)
    {
        if (u&1) s = (s * t) % n;
        u >>= 1;
        t = (t * t) % n;
    }
    
    return s;
}
```

另一种递归算法为：



```c
unsigned long fast_exp(unsigned long x, unsigned long y, unsigned long N) {
    unsigned long tmp;
    if (y == 1) return (x % N);

    if ((y & 1) == 0) {
        tmp = fast_exp(x, y/2, N);
        return ((tmp * tmp) % N);
    }
    else {
        tmp = fast_exp(x, (y-1) / 2, N);
        tmp = (tmp * tmp) % N;
        tmp = (tmp * x) % N;
        return (tmp);
    }
}
```

对应的 python 实现如下。



```python
def qe2(x, y, n):
    s = 1
    t = x
    u = y

    while (u):
        if (u&1):
            s = (s * t) % n
        u >>= 1
        t = (t * t) % n
    
    return s
```

另一种递归算法为：



```python
def fast_exp(x, y, N):
    if (y == 1):
        return (x % N)

    if ((y & 1) == 0):
        tmp = fast_exp(x, y/2, N)
        return ((tmp * tmp) % N)
    else:
        tmp = fast_exp(x, (y-1) / 2, N)
        tmp = (tmp * tmp) % N
        tmp = (tmp * x) % N
        return (tmp)
```

如果用 ![k](https://math.jianshu.com/math?formula=k) 表示数 ![x](https://math.jianshu.com/math?formula=x) 中位数的长度，这项技术平均可减少 ![1.5k](https://math.jianshu.com/math?formula=1.5k) 次操作。

### 2.2 整除性与素数

素数是这样一种数：比 1 大，其因子只有 1 和它本身，没有其他数可以整除它。2 是一个素数，其他的素数如 73、2521、2365347734399 和 ![2^{756839}-1](https://math.jianshu.com/math?formula=2%5E%7B756839%7D-1) 等。素数是无限的。密码学，特别是公开密钥密码学常用大的素数 (512 位，甚至更大)。

如果 ![b](https://math.jianshu.com/math?formula=b) 除以 ![a](https://math.jianshu.com/math?formula=a) 余数为 0，则称 ![a](https://math.jianshu.com/math?formula=a) 是 ![b](https://math.jianshu.com/math?formula=b) 的一个因子 (记叙 ![a|b](https://math.jianshu.com/math?formula=a%7Cb)，读作 “a整除b”)。比如，7 是 35 的一个因子，记作 ![7|35](https://math.jianshu.com/math?formula=7%7C35)。如果一个数只有 1 和它自身两个正因子，我们就称这个数是素数。比如，13 是素数，两个因子为 1 和 13。最初几个素数很容易找到：2、3、5、7、11、13...... 如果一个整数大于 1 且不为素数，我们就称为合数。1 既不是素数也不是合数。

下面是关于整除性的一个简单的引理。

引理 1：如果 ![a|b](https://math.jianshu.com/math?formula=a%7Cb) 且 ![b|c](https://math.jianshu.com/math?formula=b%7Cc)，那么 ![a|c](https://math.jianshu.com/math?formula=a%7Cc)。

证明：如果 ![a|b](https://math.jianshu.com/math?formula=a%7Cb)，那么存在整数 ![s](https://math.jianshu.com/math?formula=s) 使得 ![as = b](https://math.jianshu.com/math?formula=as%20%3D%20b) (由 ![b](https://math.jianshu.com/math?formula=b) 能被 ![a](https://math.jianshu.com/math?formula=a) 整除可知 ![b](https://math.jianshu.com/math?formula=b) 是 ![a](https://math.jianshu.com/math?formula=a) 的倍数)；如果 ![b|c](https://math.jianshu.com/math?formula=b%7Cc)，同样存在整数 ![t](https://math.jianshu.com/math?formula=t) 使得 ![bt=c](https://math.jianshu.com/math?formula=bt%3Dc)。综上可知，![c=bt=(as)t=a(st)](https://math.jianshu.com/math?formula=c%3Dbt%3D(as)t%3Da(st))，所以 ![a](https://math.jianshu.com/math?formula=a) 为 ![c](https://math.jianshu.com/math?formula=c) 的一个因子。

引理 2：如果 ![n](https://math.jianshu.com/math?formula=n) 为大于 1 的正整数且 ![d](https://math.jianshu.com/math?formula=d) 为 ![n](https://math.jianshu.com/math?formula=n) 除 1 之外最小的因子，那么 ![d](https://math.jianshu.com/math?formula=d) 是素数。

证明：首先，我们必须保证 ![d](https://math.jianshu.com/math?formula=d) 是被明确定义的。(如果对于某个 ![n](https://math.jianshu.com/math?formula=n)，除 1 之外不存在一个最小的因子，那么 ![d](https://math.jianshu.com/math?formula=d) 的定义就不恰当，引理 2 就毫无意义。) 由于 ![n](https://math.jianshu.com/math?formula=n) 也是 ![n](https://math.jianshu.com/math?formula=n) 的一个因子，而 ![n>1](https://math.jianshu.com/math?formula=n%3E1)，所以 ![n](https://math.jianshu.com/math?formula=n) 至少有一个大于 1 的因子，也必然有一个大于 1 的最小因子。

为证明 ![d](https://math.jianshu.com/math?formula=d) 是素数，我们使用一种标准的数学技巧，称为反证法。为证明结论 X，反证法的一般思路是假设 X 不成立，接着从这个假设推出矛盾；如果假设 X 不成立能够推出矛盾，那么 X 必须是正确的。

在这个例子中，我们假设 ![d](https://math.jianshu.com/math?formula=d) 不是素数，那么 ![d](https://math.jianshu.com/math?formula=d) 肯定存在满足 ![1<e<d](https://math.jianshu.com/math?formula=1%3Ce%3Cd) 的因子 ![e](https://math.jianshu.com/math?formula=e)。但是从引理 1 可知，如果 ![e|d](https://math.jianshu.com/math?formula=e%7Cd) 且 ![d|n](https://math.jianshu.com/math?formula=d%7Cn)，那么 ![e|n](https://math.jianshu.com/math?formula=e%7Cn)，即 ![e](https://math.jianshu.com/math?formula=e) 也是 ![n](https://math.jianshu.com/math?formula=n) 的一个因子且 ![e<d](https://math.jianshu.com/math?formula=e%3Cd)。这样就产生了矛盾，因为 ![d](https://math.jianshu.com/math?formula=d) 被定义为 ![n](https://math.jianshu.com/math?formula=n) 除 1 之外最小的因子，因此我们的假设是错误的，从而 ![d](https://math.jianshu.com/math?formula=d) 是素数。

定理 3 (殴几里得)：素数有无穷多个。

证明：我们仍然使用反证法来证明。假设素数的个数是有限的，那么一个包含所有素数的列表也是有限的，记为 ![p_1,p_2,p_3,...,p_k](https://math.jianshu.com/math?formula=p_1%2Cp_2%2Cp_3%2C...%2Cp_k)，这里 ![k](https://math.jianshu.com/math?formula=k) 表示素数的个数。定义 ![n = p_1p_2p_3...p_k+1](https://math.jianshu.com/math?formula=n%20%3D%20p_1p_2p_3...p_k%2B1)，即 ![n](https://math.jianshu.com/math?formula=n) 为所有素数的乘积加上 1。

考虑 ![n](https://math.jianshu.com/math?formula=n) 除 1 之外的最小因子，我们仍用 ![d](https://math.jianshu.com/math?formula=d) 来表示这个因子。由引理 2 可知，![d](https://math.jianshu.com/math?formula=d) 为素数且 ![d|n](https://math.jianshu.com/math?formula=d%7Cn)；但是在那个有限的素数列表中，没有一个素数是 ![n](https://math.jianshu.com/math?formula=n) 的因子，因为它们都是 ![n-1](https://math.jianshu.com/math?formula=n-1) 的因子，![n](https://math.jianshu.com/math?formula=n) 除以列表中任何一个素数 ![p_i](https://math.jianshu.com/math?formula=p_i) 都会有余数 1，所以 ![d](https://math.jianshu.com/math?formula=d) 为素数且不在列表中。而列表在定义时就包含了所有素数的，这样就出现了矛盾，所以素数的个数是有限的这个假设是错误的，从而可知素数有无穷多个。

### 2.3 最大公因子

两个数互素 (relatively prime) 是指：当它们除了 1 外没有共同的因子。换句话说，如果 ![a](https://math.jianshu.com/math?formula=a) 和 ![n](https://math.jianshu.com/math?formula=n) 的最大公因子 (greatest common divisor) 等于 1，那么可写作：
 ![gcd(a,n) = 1](https://math.jianshu.com/math?formula=gcd(a%2Cn)%20%3D%201)

数 15 和 28 是互素的，15 和 27 不是，而 13 和 500 是。一个素数与它的倍数以外的任何其他数都是互素的。

计算两个数的最大公因子最容易的方法是用殴几里得算法 (Euclid's algorithm)。殴几里德在公元前 300 年所写的《Elements》中描述了这个算法。这个算法并非由他发明，历史学家相信这个算法在当时已有 200 年历史。它是幸存到现在最古老的非凡的算法，至今它仍是完好的。

算法的 C 语言描述如下：



```c
/* returns gcd of x and y */
int gcd(int x, int y) {
    int g;

    if (x < 0)
        x = -x;
    
    if (y < 0)
        y = -y;

    if (x + y == 0)
        exit(1);

    g = y;

    while (x > 0)
    {
        g = x;
        x = y % x;
        y = g;
    }
    
    return g;
}
```

这个算法可以推广为返回由 m 个数组成的 gcd 数组。



```c
/* return the gcd of x1, x2, ..., xm */
int multiple_gcd(int m, int *x) {
    size_t i;
    int g;

    if (m < 1)
        return 0;
    
    g = x[0];

    for (i = 1; i < m; ++i) {
        g = gcd(g, x[i]);

        /* optimization, since for random x(i), g==1 60% of the time: */
        if (g == 1)
            return 1;
    }

    return g;
}
```

对应的 python 实现如下。



```python
# returns gcd of x and y
def gcd(x, y):
    if (x < 0):
        x = -x
    
    if (y < 0):
        y = -y

    if (x + y == 0):
        exit(1)

    g = y

    while (x > 0):
        g = x
        x = y % x
        y = g
    
    return g
```

这个算法可以推广为返回由 m 个数组成的 gcd 数组。



```python
# return the gcd of x1, x2, ..., xm
def multiple_gcd(m, x):
    if (m < 1):
        return 0
    
    g = x[0]

    for i in range(m):
        g = gcd(g, x[i])

        # optimization, since for random x(i), g==1 60% of the time:
        if (g == 1):
            return 1

    return g
```

### 2.4 殴几里得算法

求最大公因子 (GCD) 的算法。

### 2.5 求模逆元

记得逆元 (inverse) 吗？4 的乘法逆元是 1/4，因为 ![4 \times 1 / 4 = 1](https://math.jianshu.com/math?formula=4%20%5Ctimes%201%20%2F%204%20%3D%201)。在模运算的领域，这个问题更复杂：
 ![4 \times x \equiv 1 (mod \ 7)](https://math.jianshu.com/math?formula=4%20%5Ctimes%20x%20%5Cequiv%201%20(mod%20%5C%207))

这个方程等价于寻找一组 ![x](https://math.jianshu.com/math?formula=x) 和 ![k](https://math.jianshu.com/math?formula=k)，以使：
 ![4x = 7k + 1](https://math.jianshu.com/math?formula=4x%20%3D%207k%20%2B%201)

这里 ![x](https://math.jianshu.com/math?formula=x) 和 ![k](https://math.jianshu.com/math?formula=k) 均为整数。

更为一般的问题是寻找一个 ![x](https://math.jianshu.com/math?formula=x)，使得：
 ![1 = (a \times x) mod \ n](https://math.jianshu.com/math?formula=1%20%3D%20(a%20%5Ctimes%20x)%20mod%20%5C%20n)

也可写作：
 ![a^{-1} \equiv x (mod \ n)](https://math.jianshu.com/math?formula=a%5E%7B-1%7D%20%5Cequiv%20x%20(mod%20%5C%20n))

解决模的逆元问题很困难。有时候有一个方案，有时候没有。例如，5 模 14 的逆元是 3：![5 \times 3 = 15 \equiv 1 (mod \ 14)](https://math.jianshu.com/math?formula=5%20%5Ctimes%203%20%3D%2015%20%5Cequiv%201%20(mod%20%5C%2014))。2 模 14 却没有逆元。

一般而论，如果 ![a](https://math.jianshu.com/math?formula=a) 和 ![n](https://math.jianshu.com/math?formula=n) 是互素的，那么 ![a^{-1} \equiv x (mod \ n)](https://math.jianshu.com/math?formula=a%5E%7B-1%7D%20%5Cequiv%20x%20(mod%20%5C%20n)) 有唯一解；如果 ![a](https://math.jianshu.com/math?formula=a) 和 ![n](https://math.jianshu.com/math?formula=n) 不是互素的，那么 ![a^{-1} \equiv x (mod \ n)](https://math.jianshu.com/math?formula=a%5E%7B-1%7D%20%5Cequiv%20x%20(mod%20%5C%20n)) 没有解。如果 ![n](https://math.jianshu.com/math?formula=n) 是素数，那么从 ![1 \thicksim n-1](https://math.jianshu.com/math?formula=1%20%5Cthicksim%20n-1) 的每一个数与 ![n](https://math.jianshu.com/math?formula=n) 都是互素的，且在这个范围内恰好有一个逆元。

一切顺利。现在，怎样找出 ![a](https://math.jianshu.com/math?formula=a) 模 ![n](https://math.jianshu.com/math?formula=n) 的逆元呢？有一系列的方法。殴几里得算法也能计算 ![a](https://math.jianshu.com/math?formula=a) 模 ![n](https://math.jianshu.com/math?formula=n) 的逆元，有时候这叫做扩展殴几里得算法 (extended Euclidean algorithm)。

下面是用 C++ 写的算法：



```cpp
#include <stdlib.h>

#include <iostream>

using namespace std;

#define isEven(x)       ((x & 0x01) == 0)
#define isOdd(x)        (x & 0x01)
#define swap(x,y)       (x ^= y, y ^= x, x ^= y)

void ExtBinEuclid(int *u, int *v, int *u1, int *u2, int *u3) {
    // warning: u and v will be swapped if u < v
    int k, t1, t2, t3;

    if (*u < *v) swap(*u, *v);

    for (k = 0; isEven(*u) && isEven(*v); ++k) {
        *u >>= 1; *v >>= 1;
    }

    *u1 = 1; *u2 = 0; *u3 = *u; t1 = *v; t2 = *u - 1; t3 = *v;
    
    do {
        do {
            if (isEven(*u3)) {
                if (isOdd(*u1) || isOdd(*u2)) {
                    *u1 += *v; *u2 += *u;
                }

                *u1 >>= 1; *u2 >>= 1; *u3 >>= 1;
            }

            if (isEven(t3) || *u3 < t3) {
                swap(*u1, t1); swap(*u2, t2); swap(*u3, t3);
            }
        } while (isEven(*u3));
        
        while (*u1 < t1 || *u2 < t2) {
            *u1 += *v; *u2 += *u;
        }

        *u1 -= t1; *u2 -= t2; *u3 -= t3;
    } while (t3 > 0);
    
    while (*u1 >= *v && *u2 >= *u) {
        *u1 -= *v; *u2 -= *u;
    }
    
    *u1 <<= k; *u2 <<= k; *u3 <<= k;
}

int main(int argc, char **argv) {
    int a, b, gcd;

    if (argc < 3) {
        std::cerr << "Usage: xeuclid u v" << std::endl;
        
        return -1;
    }

    int u = atoi(argv[1]);
    int v = atoi(argv[2]);

    if (u <= 0 || v <= 0) {
        std::cerr << "Arguments must be positive!" << std::endl;

        return -2;
    }

    // warning: u and v will be swapped if u < v
    ExtBinEuclid(&u, &v, &a, &b, &gcd);

    std::cout << a << " * " << u << " + (-"
            << b << ") * " << v << " = " << gcd << std::endl;

    if (gcd == 1)
        std::cout << "the inverse of " << v << " mod " << u << " is: " << u - b << std::endl;
    
    return 0;
}
```

此算法通过迭代运算来实现，对于大的整数，其运行可能较慢。Knuth 指出这个算法完成的除法的平均数目是
 ![0.843 \times log_2(n) + 1.47](https://math.jianshu.com/math?formula=0.843%20%5Ctimes%20log_2(n)%20%2B%201.47)

### 2.6 求系数

殴几里得算法可用于解决下面的一类问题：给出一个包含 ![m](https://math.jianshu.com/math?formula=m) 个变量 ![x_1, x_2, ..., x_m](https://math.jianshu.com/math?formula=x_1%2C%20x_2%2C%20...%2C%20x_m) 的数组，求一个包含 ![m](https://math.jianshu.com/math?formula=m) 个系数 ![u_1, u_2, ..., u_m](https://math.jianshu.com/math?formula=u_1%2C%20u_2%2C%20...%2C%20u_m) 的数组，使得
 ![u_1 \times x_1 + ... + u_m \times x_m = 1](https://math.jianshu.com/math?formula=u_1%20%5Ctimes%20x_1%20%2B%20...%20%2B%20u_m%20%5Ctimes%20x_m%20%3D%201)

### 2.7 费马小定理

如果 ![m](https://math.jianshu.com/math?formula=m) 是一个素数，且 ![a](https://math.jianshu.com/math?formula=a) 不是 ![m](https://math.jianshu.com/math?formula=m) 的倍数，那么，根据费马小定理 (Fermat's little theorem) 有：
 ![a^{m-1} \equiv 1 \ (mod \ m)](https://math.jianshu.com/math?formula=a%5E%7Bm-1%7D%20%5Cequiv%201%20%5C%20(mod%20%5C%20m))

### 2.8 欧拉函数

还有另一种方法计算模 ![n](https://math.jianshu.com/math?formula=n) 的逆元，但不是在任何情况下都能使用。模 ![n](https://math.jianshu.com/math?formula=n) 的余数化简集 (reduced set of residues) 是余数完全集合的子集，与 ![n](https://math.jianshu.com/math?formula=n) 互素。例如，模 12 的余数化简集是 ![{1, 5, 7, 11}](https://math.jianshu.com/math?formula=%7B1%2C%205%2C%207%2C%2011%7D)。如果 ![n](https://math.jianshu.com/math?formula=n) 是素数，那么模 ![n](https://math.jianshu.com/math?formula=n) 的余数化简集是从 ![1 \thicksim n-1](https://math.jianshu.com/math?formula=1%20%5Cthicksim%20n-1) 的所有整数集合。对 ![n](https://math.jianshu.com/math?formula=n) 不等于 1 的数，数 0 不是余数化简集的元素。

欧拉函数 (Euler totient fuction)，也称为欧拉 ![\varphi](https://math.jianshu.com/math?formula=%5Cvarphi) 函数，写作 ![\phi(n)](https://math.jianshu.com/math?formula=%5Cphi(n))，它表示模 ![n](https://math.jianshu.com/math?formula=n) 的余数化简集中元素的数目。换句话说，![\phi(n)](https://math.jianshu.com/math?formula=%5Cphi(n)) 表示与 ![n](https://math.jianshu.com/math?formula=n) 互素的小于 ![n](https://math.jianshu.com/math?formula=n) 的正整数的数目 (![n>1](https://math.jianshu.com/math?formula=n%3E1))。

如果 ![n](https://math.jianshu.com/math?formula=n) 是素数，那么 ![\phi(n)=n-1](https://math.jianshu.com/math?formula=%5Cphi(n)%3Dn-1)；如果 ![n=pq](https://math.jianshu.com/math?formula=n%3Dpq)，且 ![p](https://math.jianshu.com/math?formula=p) 和 ![q](https://math.jianshu.com/math?formula=q) 互素，那么 ![\phi(n)=(p-1)(q-1)](https://math.jianshu.com/math?formula=%5Cphi(n)%3D(p-1)(q-1))。这些数字在随后谈到的公开密钥系统中将再次出现，它们都来自于此。

根据费马小定理的欧拉推广，如果 ![gcd(a,n)=1](https://math.jianshu.com/math?formula=gcd(a%2Cn)%3D1)，那么
 ![a^{\phi(n)} \ mod \ n = 1](https://math.jianshu.com/math?formula=a%5E%7B%5Cphi(n)%7D%20%5C%20mod%20%5C%20n%20%3D%201)

现在计算 ![a](https://math.jianshu.com/math?formula=a) 模 ![n](https://math.jianshu.com/math?formula=n) 很容易：
 ![x = a^{\phi(n)-1} \ mod \ n](https://math.jianshu.com/math?formula=x%20%3D%20a%5E%7B%5Cphi(n)-1%7D%20%5C%20mod%20%5C%20n)

现在计算 ![a](https://math.jianshu.com/math?formula=a) 模 ![n](https://math.jianshu.com/math?formula=n) 很容易：
 ![x = a^{\phi(n)-1} \ mod \ n](https://math.jianshu.com/math?formula=x%20%3D%20a%5E%7B%5Cphi(n)-1%7D%20%5C%20mod%20%5C%20n)

证明：
 ![(a \times x)mod \ n = (a \times a^{\phi(n)-1})mod \ n = a^{\phi(n)} \ mod \ n = 1](https://math.jianshu.com/math?formula=(a%20%5Ctimes%20x)mod%20%5C%20n%20%3D%20(a%20%5Ctimes%20a%5E%7B%5Cphi(n)-1%7D)mod%20%5C%20n%20%3D%20a%5E%7B%5Cphi(n)%7D%20%5C%20mod%20%5C%20n%20%3D%201)

例如，求 5 模 7 的逆元是多少？既然 7 是素数，![\phi(7)=7-1=6](https://math.jianshu.com/math?formula=%5Cphi(7)%3D7-1%3D6)。因此，5 模 7 的逆元是
 ![5^{6-1} mod \ 7 = 5^5 mod \ 7 = 3](https://math.jianshu.com/math?formula=5%5E%7B6-1%7D%20mod%20%5C%207%20%3D%205%5E5%20mod%20%5C%207%20%3D%203)

计算逆元的两种方法都推广到在一般性的问题中求解 ![x](https://math.jianshu.com/math?formula=x) (如果 ![gcd(a,n)=1](https://math.jianshu.com/math?formula=gcd(a%2Cn)%3D1))：
 ![(a \times x) mod \ n = b](https://math.jianshu.com/math?formula=(a%20%5Ctimes%20x)%20mod%20%5C%20n%20%3D%20b)

用欧拉推广公式，解：
 ![x = (b \times a^{\phi(n)-1}) mod \ n](https://math.jianshu.com/math?formula=x%20%3D%20(b%20%5Ctimes%20a%5E%7B%5Cphi(n)-1%7D)%20mod%20%5C%20n)

用殴几里得算法，解：
 ![x = (b \times (a^{-1} mod \ n)) mod \ n](https://math.jianshu.com/math?formula=x%20%3D%20(b%20%5Ctimes%20(a%5E%7B-1%7D%20mod%20%5C%20n))%20mod%20%5C%20n)

通常，殴几里得算法在计算逆元方面比欧拉推广更快，特别是对于 500 位范围内的数。如果 ![gcd(a, n) \neq 1](https://math.jianshu.com/math?formula=gcd(a%2C%20n)%20%5Cneq%201)，并非一切都没用了。这种一般情况而言，![(a \times x) \ mod \ n = b](https://math.jianshu.com/math?formula=(a%20%5Ctimes%20x)%20%5C%20mod%20%5C%20n%20%3D%20b)，可能有多个解或无解。

### 2.9 中国剩余定理

如果已知 ![n](https://math.jianshu.com/math?formula=n) 的素因子，那么就能利用中国剩余定理 (Chinese remainder theorem) 求解整个方程组，这个定理的最初形式是由 1 世纪的中国数学家孙子发现的。

一般而言，如果 ![n](https://math.jianshu.com/math?formula=n) 的素因子可分解为 ![n = p_1 \times p_2 \times ... \times p_t](https://math.jianshu.com/math?formula=n%20%3D%20p_1%20%5Ctimes%20p_2%20%5Ctimes%20...%20%5Ctimes%20p_t)，那么方程组
 ![(x \ mod \ p_i) = a_i \qquad i=1,2,...,t](https://math.jianshu.com/math?formula=(x%20%5C%20mod%20%5C%20p_i)%20%3D%20a_i%20%5Cqquad%20i%3D1%2C2%2C...%2Ct)

有唯一解，这里 ![x<n](https://math.jianshu.com/math?formula=x%3Cn) (注意，有些素数可能不止一次地出现。例如，p_1 可能等于 p_2)。换句话说，一个数 (小于一些素数之积) 被它的余数模这些素数唯一确定。

例如，取素数 3 和 5，取一个数 14，那么 ![14 \ mod \ 3 = 2, \quad 14 \ mod \ 5 = 4](https://math.jianshu.com/math?formula=14%20%5C%20mod%20%5C%203%20%3D%202%2C%20%5Cquad%2014%20%5C%20mod%20%5C%205%20%3D%204)。则小于 ![3 \times 5 = 15](https://math.jianshu.com/math?formula=3%20%5Ctimes%205%20%3D%2015) 且具有上述余数的数只有 14，即由这两个余数唯一地确定了数 14。

如果对任意 ![a<p](https://math.jianshu.com/math?formula=a%3Cp) 和 ![b<q](https://math.jianshu.com/math?formula=b%3Cq) (![p](https://math.jianshu.com/math?formula=p) 和 ![q](https://math.jianshu.com/math?formula=q) 都是素数)，那么，当 ![x < p \times q](https://math.jianshu.com/math?formula=x%20%3C%20p%20%5Ctimes%20q) 时，存在一个唯一的 ![x](https://math.jianshu.com/math?formula=x)，使得
 ![x \equiv a (mod \ p) 且 x \equiv b (mod \ q)](https://math.jianshu.com/math?formula=x%20%5Cequiv%20a%20(mod%20%5C%20p)%20%E4%B8%94%20x%20%5Cequiv%20b%20(mod%20%5C%20q))

为求出这个 ![x](https://math.jianshu.com/math?formula=x)，首先用殴几里得算法找到 ![u](https://math.jianshu.com/math?formula=u)，使得
 ![u \times q \equiv 1 (mod \ p)](https://math.jianshu.com/math?formula=u%20%5Ctimes%20q%20%5Cequiv%201%20(mod%20%5C%20p))

然后计算：
 ![x = (((a - b) \times u) mod \ p) \times q + b](https://math.jianshu.com/math?formula=x%20%3D%20(((a%20-%20b)%20%5Ctimes%20u)%20mod%20%5C%20p)%20%5Ctimes%20q%20%2B%20b)

用 C 语言所写的中国剩余定理如下：



```c
/* r is the number of elements in arrays m and u;
m is the array of (pairwise relatively prime) moduli
u is the array of coefficients
return value is n such than n == u[k]%m[k] (k=0..r-1) and
    n < m[0]*m[1]*...*m[r-1]
*/

/* totient() is left as an exercise to the reader. */

int chinese_remainder(size_t r, int *m, int *u) {
    size_t i;
    int modulus;
    int n;
    modulus = 1;
    for (i = 0; i < r; ++i)
        modulus *= m[i];
    n = 0;
    for (i = 0; i < r; ++i) {
        n += u[i] * modexp(modulus / m[i], totient(m[i]), m[i]);
        n %= modulus;
    }
    
    return n;
}
```

中国剩余定理的一个推论可用于求出一个类似问题的解：如果 ![p](https://math.jianshu.com/math?formula=p) 和 ![q](https://math.jianshu.com/math?formula=q) 都是素数，且 ![p < q](https://math.jianshu.com/math?formula=p%20%3C%20q)，那么存在一个唯一的 ![x < p \times q](https://math.jianshu.com/math?formula=x%20%3C%20p%20%5Ctimes%20q)，使得
 ![a \equiv x (mod \ p) 且 b \equiv x (mod \ q)](https://math.jianshu.com/math?formula=a%20%5Cequiv%20x%20(mod%20%5C%20p)%20%E4%B8%94%20b%20%5Cequiv%20x%20(mod%20%5C%20q))

如果 ![a \geq b \ mod \ p](https://math.jianshu.com/math?formula=a%20%5Cgeq%20b%20%5C%20mod%20%5C%20p)，那么
 ![x = (((a - (b \ mod \ p)) \times u) mod \ p) \times q + b](https://math.jianshu.com/math?formula=x%20%3D%20(((a%20-%20(b%20%5C%20mod%20%5C%20p))%20%5Ctimes%20u)%20mod%20%5C%20p)%20%5Ctimes%20q%20%2B%20b)

如果 ![a < b \ mod \ p](https://math.jianshu.com/math?formula=a%20%3C%20b%20%5C%20mod%20%5C%20p)，那么
 ![x = (((a + p - (b \ mod \ p)) \times u) mod \ p) \times q + b](https://math.jianshu.com/math?formula=x%20%3D%20(((a%20%2B%20p%20-%20(b%20%5C%20mod%20%5C%20p))%20%5Ctimes%20u)%20mod%20%5C%20p)%20%5Ctimes%20q%20%2B%20b)

### 2.10 二次剩余

如果 ![p](https://math.jianshu.com/math?formula=p) 是素数，且 ![a < p](https://math.jianshu.com/math?formula=a%20%3C%20p)，如果
 ![x^{2} \equiv a \ (mod \ p) \qquad 对某些 x 成立](https://math.jianshu.com/math?formula=x%5E%7B2%7D%20%5Cequiv%20a%20%5C%20(mod%20%5C%20p)%20%5Cqquad%20%E5%AF%B9%E6%9F%90%E4%BA%9B%20x%20%E6%88%90%E7%AB%8B)

那么称 ![a](https://math.jianshu.com/math?formula=a) 是对模 ![p](https://math.jianshu.com/math?formula=p) 的二次剩余 (quadratic residue)。

不是所有的 ![a](https://math.jianshu.com/math?formula=a) 的值都满足这个特性。如果 ![a](https://math.jianshu.com/math?formula=a) 是对模 ![n](https://math.jianshu.com/math?formula=n) 的一个二次剩余，那么它必定是对模 ![n](https://math.jianshu.com/math?formula=n) 的所有素因子的二次剩余。例如，如果 ![p = 7](https://math.jianshu.com/math?formula=p%20%3D%207)，那么二次剩余是 1、2 和 4：
 ![1^2 = 1 \equiv 1(mod \ 7)](https://math.jianshu.com/math?formula=1%5E2%20%3D%201%20%5Cequiv%201(mod%20%5C%207))

![2^2 = 4 \equiv 4(mod \ 7)](https://math.jianshu.com/math?formula=2%5E2%20%3D%204%20%5Cequiv%204(mod%20%5C%207))

![3^2 = 9 \equiv 2(mod \ 7)](https://math.jianshu.com/math?formula=3%5E2%20%3D%209%20%5Cequiv%202(mod%20%5C%207))

![4^2 = 16 \equiv 2(mod \ 7)](https://math.jianshu.com/math?formula=4%5E2%20%3D%2016%20%5Cequiv%202(mod%20%5C%207))

![5^2 = 25 \equiv 4(mod \ 7)](https://math.jianshu.com/math?formula=5%5E2%20%3D%2025%20%5Cequiv%204(mod%20%5C%207))

![6^2 = 36 \equiv 1(mod \ 7)](https://math.jianshu.com/math?formula=6%5E2%20%3D%2036%20%5Cequiv%201(mod%20%5C%207))

注意，每一个二次剩余在上面都出现了两次。

没有 ![x](https://math.jianshu.com/math?formula=x) 值可满足下列这些方程的任意一个：
 ![x^2 = 3 (mod \ 7)](https://math.jianshu.com/math?formula=x%5E2%20%3D%203%20(mod%20%5C%207))

![x^2 = 5 (mod \ 7)](https://math.jianshu.com/math?formula=x%5E2%20%3D%205%20(mod%20%5C%207))

![x^2 = 6 (mod \ 7)](https://math.jianshu.com/math?formula=x%5E2%20%3D%206%20(mod%20%5C%207))

对模 7 的二次非剩余 (quadratic nonresidue) 是 3、5 和 6。

很容易证明，当 ![p](https://math.jianshu.com/math?formula=p) 为奇数时，对模 ![p](https://math.jianshu.com/math?formula=p) 的二次剩余数目恰好是 ![(p-1)/2](https://math.jianshu.com/math?formula=(p-1)%2F2)，且与其二次非剩余的数目相同。而且，如果 ![x^2](https://math.jianshu.com/math?formula=x%5E2) 等于二次剩余模 ![p](https://math.jianshu.com/math?formula=p)，那么 ![x^2](https://math.jianshu.com/math?formula=x%5E2) 恰好有两个平方根：其中一个在 ![1 \thicksim (p-1)/2](https://math.jianshu.com/math?formula=1%20%5Cthicksim%20(p-1)%2F2) 之间；另一个在 ![(p+1)/2 \thicksim (p-1)](https://math.jianshu.com/math?formula=(p%2B1)%2F2%20%5Cthicksim%20(p-1)) 之间。这两个平方根中的一个也是模 ![p](https://math.jianshu.com/math?formula=p) 的二次剩余，称为主平方根 (pricipal square root)。

如果 ![n](https://math.jianshu.com/math?formula=n) 是两个素数 ![p](https://math.jianshu.com/math?formula=p) 和 ![q](https://math.jianshu.com/math?formula=q) 之积，那么模 ![n](https://math.jianshu.com/math?formula=n) 恰好有 ![(p-1)(q-1)/4](https://math.jianshu.com/math?formula=(p-1)(q-1)%2F4) 个二次剩余。模 ![n](https://math.jianshu.com/math?formula=n) 的一个二次剩余是模 ![n](https://math.jianshu.com/math?formula=n) 的一个完全平方。这是因为要成为模 ![n](https://math.jianshu.com/math?formula=n) 的平方，其余数必须有模 ![p](https://math.jianshu.com/math?formula=p) 的平方和模 ![q](https://math.jianshu.com/math?formula=q) 的平方。例如，模 35 有 11 个二次剩余：1、4、9、11、14、15、16、21、25、29、30。每一个二次剩余恰好有 4 个平方根。

## 3. 有限域上的离散对数

模指数运算是频繁地用于密码学中的另一种单向函数。计算下面的表达式很容易：
 ![a^x \ mod \ n](https://math.jianshu.com/math?formula=a%5Ex%20%5C%20mod%20%5C%20n)

模指数运算的逆问题是找出一个数的离散对数，这是一个难题：
 ![求解 x，使得 a^x \equiv \ b \ (mod \ n)](https://math.jianshu.com/math?formula=%E6%B1%82%E8%A7%A3%20x%EF%BC%8C%E4%BD%BF%E5%BE%97%20a%5Ex%20%5Cequiv%20%5C%20b%20%5C%20(mod%20%5C%20n))

例如：
 ![如果 3^x \equiv 15 \ mod \ 17，那么 x = 6](https://math.jianshu.com/math?formula=%E5%A6%82%E6%9E%9C%203%5Ex%20%5Cequiv%2015%20%5C%20mod%20%5C%2017%EF%BC%8C%E9%82%A3%E4%B9%88%20x%20%3D%206)

不是所有的离散对数都有解 (记住，只有整数才是合法的解)。发现下面的方程没有解 ![x](https://math.jianshu.com/math?formula=x) 很容易：
 ![3^x \equiv \ 7 \ mod \ 13](https://math.jianshu.com/math?formula=3%5Ex%20%5Cequiv%20%5C%207%20%5C%20mod%20%5C%2013)

对 1024 位的数求离散对数更加困难。

### 3.1 计算有限群中的离散对数

密码设计者对下面三个主要群的离散对数很感兴趣：

- 素数域的乘法群：![GF(p)](https://math.jianshu.com/math?formula=GF(p))。
- 特征为 2 的有限域上的乘法群：![GF(2^n)](https://math.jianshu.com/math?formula=GF(2%5En))。
- 有限域 ![F](https://math.jianshu.com/math?formula=F) 上的椭圆曲线群：![EC(F)](https://math.jianshu.com/math?formula=EC(F))。

许多公开密钥算法的安全性是基于寻找离散对数的，因此对这个问题进行了广泛的研究。

## 项目源代码

项目源代码会逐步上传到 Github，地址为 [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)。

## Contributor

1. Windstamp, [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)



作者：furnace
链接：https://www.jianshu.com/p/068f99486f06
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。







# 1. Hyperledger Fabric - 介绍

[官方文档](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Fwhatis.html) 给出了 Hyperledger Fabric 的介绍。

本文主要是基于上述文档，并根据自己对区块链的认知从文档中提取关键内容并进行重新组织。

## 1. 引言

最主流的两条公链 Bitcoin 和 Ethereum，它们的网络都是公开的不需要许可的 (public permissionless)。即它们的网络对任何人公开，同时参与者是匿名进行交互的。

随着区块链技术的普及，更多具有创新性的企业级应用案例也开始尝试使用区块链技术。然而，这里存在几个比较大的障碍。第一个是，公链技术达不到企业级应用案例需要的性能。第二个是，对参与者的身份标识是硬性要求，如在金融交易中必需的 KYC (Know-Your-Customer) 和 AML (Anti-Money- Laundering)。

对于企业级应用，必须考虑如下需求：

- 参与者必须是可标识的
- 网络必须采用许可制 (permissioned)
- 性能上需要很高的交易吞吐率
- 交易确认还需要低延时
- 与商业交易相关的交易和数据需要隐私性和机密性 (privacy and confidentiality)

虽然也有许多早期的区块链平台目前被用于企业级应用，但是 Hyperledger Fabric 从一开始就是为企业级应用而设计的。下面的内容主要介绍 Hyperledger Fabric 本身与其它区块链平台的差别，以及描述部分体系架构的设计动机。

## 2. Hyperledger Fabric

Hyperledger Fabric 是开源的企业级许可制分布式账本技术 (Distributed Ledger Technology, DLT) 平台，这与其它流行的 DLT 平台或区块链平台有一些关键的不同。

- Hyperledger 是由 Linux Foundation 建立的，Linux Foundation 本身具有非常长和非常成功的开源历史及开源项目。Hyperledger 是由多个技术会员会管理的，Hyperledger Fabric 项目由来自多个组织的人员维护。它的开发者社区包括 35 个组织和将近 200 个开发者。
- Fabric 具有高度的模块化和可配置性的架构，支持对广泛范围企业应用的创新和优化，如：银行、金融、保险、医疗、人力资源、供应链、和数字音乐分发。
- Fabric 是第一个支持采用通用编程语言编写智能合约的 DLT 平台，如 Go, Node.js, Java，而不是限制于领域特定语言 (domain-specific languages, DSL)。这意味着大多数企业已经具备开发智能合约的技能集，而不需要额外的培训来学习一门新的语言。
- Fabric 平台采用的是许可制，意味着与公开无需许可的网络不同的是，参与者已经各自有一定了解，而不是匿名和完全无信任的。这意味着，尽管参与者之间可能不会完全信任对方（例如，他们可能是同一行业的竞争者），但网络可以在治理模型下运行，该治理模型是基于参与者之间确实存在的信任而建立的，例如处理争议的法律协议或框架。
- Fabric 平台最重要的特色之一是它对可插拔共识协议的支持，该协议使平台可以更有效地进行定制，以适应特定的用例和信任模型。例如，当部署在单个企业中或由受信任的机构运营时，完全拜占庭式的容错共识可能被认为是不必要的，并且会严重拖累性能和吞吐量。在这种情况下，崩溃容错 (crash fault-tolerant, CFT) 共识协议可能绰绰有余，而在多方，去中心化的用例中，可能需要更传统的拜占庭容错 (byzantine fault tolerant, BFT) 共识协议。
- Fabric 利用不需要原生加密货币的共识协议来实现代价昂贵的挖矿操作或智能合约的执行。避免使用加密货币会降低一些重大的风险，并且无需进行加密货币挖矿操作就意味着可以以与任何其他分布式系统大致相同的运营成本来部署该平台。

这些差异化设计功能的结合使 Fabric 在交易处理 (transaction processing) 和交易确认延迟 (transaction confirmation latency) 方面成为当今性能最好的平台之一，并实现了交易和智能合约  (在 Fabric 中称为 chaincode) 的隐私和机密性 (privacy and confidentiality)。

下面我们将更详细地探讨这些差异化功能。

## 3. 高度模块化

Hyperledger Fabric 专门设计为具有模块化的体系结构。无论是可插拔共识，可插拔身份管理协议 (例如 LDAP 或 OpenID Connect)，密钥管理协议还是密码库，该平台将可配置性作为它的设计核心，这样可以满足企业应用需求的多样性。

在较高的层次上，Fabric 由以下模块化组件组成：

- ordering service: 可插拔订单服务在交易顺序上达成共识，然后将区块广播给对端节点。可以简单地类比于比特币中的挖矿模块和 P2P 模块。
- membership service provider: 可插拔成员资格服务提供者负责将网络中的实体与加密身份相关联。可以简单地认为是 Fabric 中特有的网络许可模块。
- peer-to-peer gossip service: 可选的点对点传输服务通过向其他对等节点订购服务来分发输出区块。可以简单地类比于比特币中的 P2P 模块。
- chaincode: 智能合约 ("chaincode") 在容器环境 (例如 Docker) 中运行以进行隔离。它们可以用通用编程语言编写，但不能直接访问帐本状态。(?)
- 账本可以配置为支持各种 DBMS。
- endorsement and validation policy enforcement: 可插拔的认可和验证策略实施，可以针对每个应用程序进行独立配置。可以简单地类比于比特币的矿工，但是这里更像是联盟链中的各个组织，而且单个组织可以只和特定的逻辑挂钩，可以简单类比于以太坊 2.0 中的分片技术。

业界普遍认为，没有一个区块链可以全部统治。可以通过多种方式配置 Hyperledger Fabric，以满足多种行业用例的不同解决方案要求。

## 4. 许可制区块链 VS 无需许可制区块链

在无需许可的区块链中，几乎任何人都可以参与，且每个参与者都是匿名的。在这种情况下，除了一定高度之前的区块链状态是不可变的之外，别无其他信任。为了解决这种缺乏信任的场景，无需许可的区块链采用工作量证明 (Proof Of Work, POW) 共识机制，并提供原生的加密货币作为激励。

在许可制区块链中，参与者是具有身份标识的，且大多数还经过了审核，因此整个网络是在具有一定程度信任的治理模型下运行的。许可制区块链提供了一种方法来保护一组具有共同目标但可能不会完全相互信任的实体之间的交互。通过依赖参与者的身份，许可制区块链可以使用更传统的崩溃容错 (CFT) 或拜占庭容错 (BFT) 共识协议，这些协议不需要昂贵挖矿成本。

在这种许可制的情况下，参与者通过智能合约有意引入恶意代码的风险得以降低。首先，参与者是相互了解的，并且遵循针对网络和相关交易类型建立的认可政策，所有活动 (无论是提交应用程序交易，修改网络配置还是部署智能合约) 都记录在区块链上。与完全匿名不同的，这可以根据治理模型的条款轻松地确定有罪的一方并处理恶意事件。

## 5. 智能合约

智能合约，或者在 Fabric 中称为 "chaincode"，是一种受信任的分布式应用程序，可从区块链和对等节点之间的共识机制中获得安全性/信任。用以实现区块链应用程序的业务逻辑。

应用智能合约的三个要点，尤其是将其应用于平台时：

- 网络中能够同时运行许多智能合约
- 它们可以动态部署 (在许多情况下，任何人都可以部署)
- 应用程序代码应被视为不受信任，甚至可能是恶意的

现有的大多数具有智能合约功能的区块链平台都遵循一种订单执行 (order-execute) 架构，其中的共识协议为：

- order: 验证并订购交易，然后将其传播到所有对等节点
- execute: 然后，每个对等节点依次执行交易

order-execute 架构实际上可以在所有现有的区块链系统中找到，从以太坊等公共/非许可平台 (基于 PoW 的共识) 到 [Tendermint](https://links.jianshu.com/go?to=https%3A%2F%2Ftendermint.com%2F)，[Chain](https://links.jianshu.com/go?to=http%3A%2F%2Fchain.com%2F) 和 [Quorum](https://links.jianshu.com/go?to=https%3A%2F%2Fwww.goquorum.com%2F) 等许可平台。

在以 order-execute 架构运行的区块链中执行的智能合约必须具有确定性。否则，可能永远无法达成共识。为了解决非确定性问题，许多平台要求以非标准或特领域的语言 (例如 Solidity) 编写智能合约，以便消除非确定性操作。这阻碍了广泛采用，因为它要求开发人员编写智能合约来学习一种新语言，并可能导致编程错误。

此外，由于所有交易由所有节点顺序执行，因此性能和规模受到限制。智能合约代码在系统中的每个节点上执行的事实要求采取复杂的措施来保护整个系统免受潜在的恶意合约的侵害，以确保整个系统的弹性。

## 6. Fabric 提供全新的方法来执行智能合约

Fabric 为交易引入了一种新的架构，称之为 execute-order-validate。它通过将交易流分为三个步骤来解决 order-execute 模型面临的弹性、灵活性、可伸缩性、性能和机密性挑战：(?)

- execute: 执行交易并检查其正确性，从而认可 (endorsing) 该交易
- order: 通过 (可插拔) 共识协议 order 交易
- validate: 在将交易提交到账本之前，根据特定于应用程序的背书策略 (endorsement policy) 验证交易

这种设计与 order-execute 范式完全不同，Fabric 在达成交易的最终协议之前执行交易。

在 Fabric 中，特定于应用程序的背书策略指定需要其中的哪些对端节点保证给定智能合约的正确执行。因此，每笔交易只需要由满足交易认可策略所需的对端节点的子集执行 (认可) 即可。这允许并行执行，从而提高了系统的整体性能和规模。第一阶段还消除了任何不确定性，因为不一致的结果可以在 order 前滤除。

因为消除了不确定性，所以 Fabric 是第一个启用通用编程语言使用的区块链技术。在 1.4.3 版本中，可以使用 Go、Node.js、Java 编写智能合约。

**注释**

这里尝试对 execute-order-validate 进行类比以加深理解，这里与以太坊 2.0 进行类比。以太坊 2.0 包含分片节点和中继节点，各分片节点之间可以执行不同的合约，各分片内部维护片内本身的共识，再由中继节点维护各分片之间的共识，从而维护以太坊网络的整体共识。那么，这里把 order 服务类比于中继节点，execute 的节点为分片节点，根据业务逻辑划分成不同的分片节点，即联盟链的组织节点。之间，通过 channel 进行连接。

## 7. 隐私性和机密性

正如前面已经讨论的那样，在一个利用 POW 作为共识模型的公共无许可的区块链网络中，交易在每个节点上执行。这意味着合约本身和其所处理的交易数据都不会保密。每个交易及其实现的代码对于网络中的每个节点都是可见的。在这种情况下，我们已经将合约和数据的机密性换成了 POW 交付的拜占庭容错共识。

对于许多企业级应用而言，缺乏机密性可能会成为问题。例如，在供应链合作伙伴网络中，可能会为某些消费者提供优惠价格，以巩固关系或促进额外销售。如果每个参与者都能看到每份合约和交易，那么就不可能在完全透明的网络中维持这种业务关系–每个人都希望获得优惠的价格！

再举一个例子，考虑证券行业，在该行业中，建立仓位 (或出售仓位) 的交易者不希望竞争对手知道这一点，否则他们将寻求参与竞争，从而削弱了交易者的竞争能力。

为了解决企业级应用对隐私和机密性的需求，区块链平台采用了多种方法。所有的平台都有其取舍。

加密数据是提供机密性的一种方法。但是，在利用 POW 达成共识的无许可网络中，加密数据位于每个节点上。如果有足够的时间和计算资源，则可能会破坏加密。对于许多企业级应用而言，其信息可能遭到破坏的风险是无法接受的。

零知识证明 (Zero knowledge proofs, ZKP) 是为解决此问题而正在探索的另一个研究领域，这里的权衡是，目前计算 ZKP 需要大量时间和计算资源。因此，在这种情况下的权衡是为了保密。

的利用可替代共识机制的许可制上下文中，人们可能会探索一些将机密信息仅分配给授权节点的方法。

Hyperledger Fabric 是采用许可制的平台，可通过其通道 (channel) 架构实现机密性。本质上，Fabric 网络上的参与者可以在参与者的子集之间建立一个“通道”，该通道应被授予特定交易集的可见性。将此视为网络覆盖。因此，只有那些参与通道的节点才能访问智能合约 (链码, chaincode) 和交易的数据，从而保留了两者的隐私和机密性。

为了改善其隐私和机密性功能，Fabric 增加了对私有数据 (private data) 的支持，并在未来开发可用的零知识证明 (ZKP)。随着它的可用，将对此进行更多介绍。

## 8. 可插拔的共识机制

交易的顺序被委托给模块化组件以实现共识，该组件在逻辑上与执行交易并维护帐本的对端节点分离。具体来说就是订购 (ordering) 服务。由于共识是模块化的，因此可以根据特定部署或解决方案的信任假设量身定制其实现。这种模块化体系结构允许平台依赖完善的工具包来进行 CFT (崩溃容错) 或 BFT (拜占庭容错) 排序。

Fabric 当前提供两种 CFT 订购服务 (ordering service) 实现。第一个基于 [Raft 协议](https://links.jianshu.com/go?to=https%3A%2F%2Fraft.github.io%2Fraft.pdf) 的 [etcd 库](https://links.jianshu.com/go?to=https%3A%2F%2Fetcd.io%2F)。另一个是 [Kafka](https://links.jianshu.com/go?to=https%3A%2F%2Fkafka.apache.org%2F) (内部使用 [Zookeeper](https://links.jianshu.com/go?to=https%3A%2F%2Fzookeeper.apache.org%2F))。有关当前可用订购服务 (ordering service) 的信息，请查看[有关订购的概念性文档](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Forderer%2Fordering_service.html)。

还要注意，它们不是互斥的。Fabric 网络可以具有支持不同应用程序或应用程序需求的多种订购服务。

**注释**

共识机制在与对端节点独立的模块组件中实现。

## 9. 性能和可扩展性

区块链平台的性能可能会受到许多变量的影响，例如交易大小，区块大小，网络规模，以及硬件限制等。Hyperledger 社区目前正在性能和规模工作组内制定 [一套措施草案](https://links.jianshu.com/go?to=https%3A%2F%2Fdocs.google.com%2Fdocument%2Fd%2F1DQ6PqoeIH0pCNJSEYiw7JVbExDvWh_ZRVhWkuioG4k0%2Fedit%23heading%3Dh.av4vusatnjg6)。以及称为 [Hyperledger Caliper](https://links.jianshu.com/go?to=https%3A%2F%2Fwiki.hyperledger.org%2Fprojects%2Fcaliper) 的基准测试框架的相应实现。

尽管这项工作仍在继续发展，应被视为衡量区块链平台性能和规模特征的权威，但 IBM Research 的一个团队发表了一篇 [同行评审论文](https://links.jianshu.com/go?to=https%3A%2F%2Farxiv.org%2Fabs%2F1801.10228v1)，评估了 Hyperledger Fabric 的体系结构和性能。本文提供了关于 Fabric 架构的深入讨论，然后使用 Hyperledger Fabric v1.1 的预发行版报告了团队对该平台的性能评估。

研究团队所做的基准测试工作为 Fabric v1.1.0 发行版带来了许多性能改进，使平台的整体性能比 v1.0.0 发行版提高了一倍以上。

## 10. 结论

任何对区块链平台的认真评估都应在其短名单中包括 Hyperledger Fabric。

结合起来，Fabric 的差异化功能使其成为用于许可制区块链的高度可扩展系统，支持灵活的信任假设，使该平台能够支持从政府，金融，供应链物流，医疗保健等广泛的行业用例。

更重要的是，Hyperledger Fabric 是 (当前) 十个 Hyperledger 项目中最活跃的。该平台周围的社区正在稳步增长，并且每个后续版本提供的创新远远超过了其他任何企业区块链平台。

## Reference

- Introduction, [https://hyperledger-fabric.readthedocs.io/en/latest/whatis.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Fwhatis.html)
- Introduction, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/whatis.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fwhatis.html)
- Tendermint, [https://tendermint.com/](https://links.jianshu.com/go?to=https%3A%2F%2Ftendermint.com%2F)
- Chain, [http://chain.com/](https://links.jianshu.com/go?to=http%3A%2F%2Fchain.com%2F)
- Quorum, [https://www.goquorum.com/](https://links.jianshu.com/go?to=https%3A%2F%2Fwww.goquorum.com%2F)
- Raft 协议, [https://raft.github.io/raft.pdf](https://links.jianshu.com/go?to=https%3A%2F%2Fraft.github.io%2Fraft.pdf)
- etcd 库, [https://etcd.io/](https://links.jianshu.com/go?to=https%3A%2F%2Fetcd.io%2F)
- Kafka, [https://kafka.apache.org/](https://links.jianshu.com/go?to=https%3A%2F%2Fkafka.apache.org%2F)
- Zookeeper, [https://zookeeper.apache.org/](https://links.jianshu.com/go?to=https%3A%2F%2Fzookeeper.apache.org%2F)

## 项目源代码

项目源代码会逐步上传到 Github，地址为 [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)。

## Contributor

1. Windstamp, [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)



作者：furnace
链接：https://www.jianshu.com/p/f1eba8f48176
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。



# 构建第一个 Hyperledger Fabric 网络

Hyperledger Fabric 最新的文档基于版本是 v2.0 Alpha release。由于示例中相关的 docker 镜像的版本是 v1.4.3，因为相关文档需要参考的版本为 v1.4.3。这些文档链接在本文的 Reference 部分都有涉及。

## 1. 本文目的

参考文档 [Building Your First Network](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Fbuild_network.html) 搭建第一个 Hyperledger Fabric 网络。本示例 [fabric-samples](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fhyperledger%2Ffabric-samples) 需要的 hyperledger fabric 相关的二进制程序安装参考文档 [Install Samples, Binaries and Docker Images](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Finstall.html)。而安装 hyperledger fabric 相关的二进制程序所需要的一些前置软件需求，请参考文档 [Prerequisites](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Fprereqs.html)。

## 2. 构建 Hyperledger Fabric 网络

根据 fabric-samples/first-network 示例搭建第一个 Hyperledger Fabric 网络。这里，主要是基于提供的脚本 byfn.sh 来快速搭建。脚本 byfn.sh 提供了许多可配置的参数，如果不提供则使用默认值。这里为了简化，先使用默认值。

需要注意的是，下面的命令需要在 first-network 目录下运行，但是调用的 hyperledger fabric 相关的二进制程序却是在 fabric-samples/bin 目录。

### 2.1 脚本 byfn.sh 的详细说明



```bash
Usage:
byfn.sh <mode> [-c <channel name>] [-t <timeout>] [-d <delay>] [-f <docker-compose-file>] [-s <dbtype>] [-l <language>] [-o <consensus-type>] [-i <imagetag>] [-v]"
  <mode> - one of 'up', 'down', 'restart', 'generate' or 'upgrade'"
    - 'up' - bring up the network with docker-compose up"
    - 'down' - clear the network with docker-compose down"
    - 'restart' - restart the network"
    - 'generate' - generate required certificates and genesis block"
    - 'upgrade'  - upgrade the network from version 1.3.x to 1.4.0"
  -c <channel name> - channel name to use (defaults to \"mychannel\")"
  -t <timeout> - CLI timeout duration in seconds (defaults to 10)"
  -d <delay> - delay duration in seconds (defaults to 3)"
  -f <docker-compose-file> - specify which docker-compose file use (defaults to docker-compose-cli.yaml)"
  -s <dbtype> - the database backend to use: goleveldb (default) or couchdb"
  -l <language> - the chaincode language: golang (default), node, or java"
  -o <consensus-type> - the consensus-type of the ordering service: solo (default), kafka, or etcdraft"
  -i <imagetag> - the tag to be used to launch the network (defaults to \"latest\")"
  -v - verbose mode"
byfn.sh -h (print this message)"

Typically, one would first generate the required certificates and
genesis block, then bring up the network. e.g.:"

  byfn.sh generate -c mychannel"
  byfn.sh up -c mychannel -s couchdb"
  byfn.sh up -c mychannel -s couchdb -i 1.4.0"
  byfn.sh up -l node"
  byfn.sh down -c mychannel"
  byfn.sh upgrade -c mychannel"

Taking all defaults:"
      byfn.sh generate"
      byfn.sh up"
      byfn.sh down"
```

### 2.2 使用脚本 byfn.sh 构建 Fabric 网络的详细步骤

#### Step 1. Generate Network Artifacts



```ruby
$ ./byfn.sh generate
```

#### Step 2. Bring Up the Network



```ruby
$ sudo ./byfn.sh up
```

#### Step 3. Bring Down the Network



```ruby
$ sudo ./byfn.sh down
```

## 3. 拓展

在这一部分，将把上述通过脚本 byfn.sh 构建的过程每一个关键步骤进行详细解构，并通过手动执行命令的方式来构建 hyperledger fabric 网络。

### 3.0 前置操作

下面的操作需要将容器 cli 的 FABRIC_LOGGING_SPEC 日志级别从 INFO 调整为 DEBUG。

first-network/docker-compose-cli.yaml



```php
cli:
  container_name: cli
  image: hyperledger/fabric-tools:$IMAGE_TAG
  tty: true
  stdin_open: true
  environment:
    - GOPATH=/opt/gopath
    - CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock
    - FABRIC_LOGGING_SPEC=DEBUG
    #- FABRIC_LOGGING_SPEC=INFO
```

### 3.1 Crypto Generator

- cryptogen
  - 生成密码组件的工具
  - 文件位置为 fabric-samples/bin/cryptogen
- crypto-config.yaml
  - 配置文件，用于说明示例 first-network 需要生成的密码组件
  - 文件位置为 first-network/crypto-config.yaml
- crypto-config
  - 示例 first-network 生成的密码组件
  - 文件位置为 first-network/crypto-config

### 3.2 Configuration Transaction Generator

- configtxgen
  - 生成构件，如
    - orderer `genesis block`,
    - channel `configuration transaction`,
    - and two `anchor peer transactions` - one for each Peer Org.
  - 文件位置为 fabric-samples/bin/configtxgen
- configtx.yaml.yaml
  - 配置文件，用于说明示例 first-network 需要的构件，一个 Orderer Org (`OrdererOrg`)，和两个 Peer Orgs (`Org1` & `Org2`)。
  - 文件位置为 first-network/configtx.yaml
- channel-artifacts
  - 示例 first-network 生成构件
    - channel.tx
    - genesis.block
    - Org1MSPanchors.tx
    - Org2MSPanchors.tx
  - 文件位置为 first-network/channel-artifacts

### 3.3 通过工具独立运行上述命令

下面，我们通过命令 cryptogen 和 configtxgen  手动生成上述的密码组件和示例构件。

#### 3.3.1 手动生成密码组件

参考脚本 byfn.sh 中的函数 generateCerts().



```ruby
$ ../bin/cryptogen generate --config=./crypto-config.yaml
org1.example.com
org2.example.com
```

最终的密码组件会生成到目录 first-network/crypto-config 中。

#### 3.3.2 手动生成示例构件

参考脚本 byfn.sh 中的函数 generateChannelArtifacts().



```ruby
$ export FABRIC_CFG_PATH=$PWD
$ ../bin/configtxgen -profile TwoOrgsOrdererGenesis -channelID byfn-sys-channel -outputBlock ./channel-artifacts/genesis.block
```

最终的示例构件会生成到目录 first-network/channel-artifacts 中，如 first-network/channel-artifacts/genesis.block。

#### 3.3.3 Create a Channel Configuration Transaction

注意，需要设置环境变量 CHANNEL_NAME。

###### 创建 first-network/channel-artifacts/channel.tx



```bash
# The channel.tx artifact contains the definitions for our sample channel

$ export CHANNEL_NAME=mychannel  && ../bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME
```

###### 创建 first-network/channel-artifacts/Org1MSPanchors.tx



```ruby
$ ../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP
```

###### 创建 first-network/channel-artifacts/Org2MSPanchors.tx



```ruby
$ ../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org2MSP
```

### 3.4 Start the network

如果之前利用脚本 byfn.sh 启动了测试网络，需要先关闭，如执行下列命令



```ruby
$ sudo ./byfn.sh down
```

#### 3.4.1 Start the network



```ruby
$ sudo docker-compose -f docker-compose-cli.yaml up -d
```

或通过下面的命令启动，能够立刻显示日志，但需要另启一个窗口运行下面的 cli 容器。



```ruby
$ sudo docker-compose -f docker-compose-cli.yaml up
```

#### 3.4.2 Create & Join Channel



```ruby
$ sudo docker exec -it cli bash
root@33ab5acc5622:/opt/gopath/src/github.com/hyperledger/fabric/peer#
```

在容器 cli 中需要设置的环境变量。



```ruby
# Environment variables for PEER0

$ CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
$ CORE_PEER_ADDRESS=peer0.org1.example.com:7051
$ CORE_PEER_LOCALMSPID="Org1MSP"
$ CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
```



```ruby
$ export CHANNEL_NAME=mychannel

# the channel.tx file is mounted in the channel-artifacts directory within your CLI container
# as a result, we pass the full path for the file
# we also pass the path for the orderer ca-cert in order to verify the TLS handshake
# be sure to export or replace the $CHANNEL_NAME variable appropriately

$ peer channel create -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```

上述命令将会生成文件 ``，这里是 `mychannel.block`。

###### 将 peer0.org1.example.com 加入到 channel



```ruby
# By default, this joins ``peer0.org1.example.com`` only
# the <CHANNEL_NAME.block> was returned by the previous command
# if you have not modified the channel name, you will join with mychannel.block
# if you have created a different channel name, then pass in the appropriately named block

$ peer channel join -b mychannel.block
```

###### 将 peer0.org2.example.com 加入到 channel



```ruby
$ CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
$ CORE_PEER_ADDRESS=peer0.org2.example.com:9051
$ CORE_PEER_LOCALMSPID="Org2MSP"
$ CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
$ peer channel join -b mychannel.block
```

#### 3.4.3 Update the anchor peers

###### 将 Org1 定义为 peer0.org1.example.com



```ruby
$ CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
$ CORE_PEER_ADDRESS=peer0.org1.example.com:7051
$ CORE_PEER_LOCALMSPID="Org1MSP"
$ CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
$ peer channel update -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/Org1MSPanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```

###### 将 Org2 定义为 peer0.org2.example.com



```ruby
$ CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
$ CORE_PEER_ADDRESS=peer0.org2.example.com:9051
$ CORE_PEER_LOCALMSPID="Org2MSP"
$ CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
$ peer channel update -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/Org2MSPanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```

#### 3.4.4 Install & Instantiate Chaincode

Chaincode 代码支持 golang, node, java 语言，默认为 golang，本示例中全部使用默认值。

###### Install peer0 in Org1

首先设置 Org1 相关的环境变量。



```ruby
$ CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
$ CORE_PEER_ADDRESS=peer0.org2.example.com:9051
$ CORE_PEER_LOCALMSPID="Org2MSP"
$ CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
```



```ruby
# this installs the Go chaincode. For go chaincode -p takes the relative path from $GOPATH/src
$ peer chaincode install -n mycc -v 1.0 -p github.com/chaincode/chaincode_example02/go/
```

###### Install peer0 in Org2

首先设置 Org2 相关的环境变量。



```ruby
# Environment variables for PEER0 in Org2

$ CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
$ CORE_PEER_ADDRESS=peer0.org2.example.com:9051
$ CORE_PEER_LOCALMSPID="Org2MSP"
$ CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
```



```ruby
# this installs the Go chaincode. For go chaincode -p takes the relative path from $GOPATH/src
$ peer chaincode install -n mycc -v 1.0 -p github.com/chaincode/chaincode_example02/go/
```

###### Instantiate the chaincode on the channel



```ruby
# be sure to replace the $CHANNEL_NAME environment variable if you have not exported it
# if you did not install your chaincode with a name of mycc, then modify that argument as well
$ export CHANNEL_NAME=mychannel
$ peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc -v 1.0 -c '{"Args":["init","a", "100", "b","200"]}' -P "AND ('Org1MSP.peer','Org2MSP.peer')"
```

注意，`-P "AND ('Org1MSP.peer','Org2MSP.peer')"` 中的 `AND` 可以改为 `OR`，这意味着只需要 Org1 或 Org2 其中一个组织背书就可以了。

#### 3.4.5 Query



```ruby
# be sure to set the -C and -n flags appropriately

$ peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["query","a"]}'
```

#### 3.4.6 Invoke



```swift
# be sure to set the -C and -n flags appropriately

$ peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"Args":["invoke","a","b","10"]}'
```

#### 3.4.7 Query



```ruby
# be sure to set the -C and -n flags appropriately

$ peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["query","a"]}'
```

返回结果



```rust
Query Result: 90
```

#### 3.4.8 Install peer1 in Org2

首先设置好环境变量。



```ruby
# Environment variables for PEER1 in Org2

$ CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
$ CORE_PEER_ADDRESS=peer1.org2.example.com:10051
$ CORE_PEER_LOCALMSPID="Org2MSP"
$ CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/tls/ca.crt
```



```ruby
# this installs the Go chaincode. For go chaincode -p takes the relative path from $GOPATH/src
$ peer chaincode install -n mycc -v 1.0 -p github.com/chaincode/chaincode_example02/go/
```

#### 3.4.9 Query by peer1 in Org2

首先需要将 peer1 in Org2 加入到 channel，之后才能响应查询。



```ruby
$ CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
$ CORE_PEER_ADDRESS=peer1.org2.example.com:10051
$ CORE_PEER_LOCALMSPID="Org2MSP"
$ CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/tls/ca.crt
$ peer channel join -b mychannel.block
```



```ruby
# be sure to set the -C and -n flags appropriately

$ peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["query","a"]}'
```

返回结果



```rust
Query Result: 90
```

#### 3.4.10 How do I see these transactions?



```ruby
$ sudo docker logs -f cli
```

#### 3.4.11 How can I see the chaincode logs?



```ruby
$ sudo docker logs dev-peer0.org2.example.com-mycc-1.0
ex02 Init
Aval = 100, Bval = 200
ex02 Invoke
Query Response:{"Name":"a","Amount":"100"}
ex02 Invoke
Aval = 90, Bval = 210
ex02 Invoke
Query Response:{"Name":"a","Amount":"90"}

$ sudo docker logs dev-peer0.org1.example.com-mycc-1.0
ex02 Invoke
Aval = 90, Bval = 210

$ sudo docker logs dev-peer1.org2.example.com-mycc-1.0
ex02 Invoke
Query Response:{"Name":"a","Amount":"90"}
```

### 3.5 Using CouchDB



```ruby
$ sudo docker-compose -f docker-compose-cli.yaml -f docker-compose-couch.yaml up -d
```

## Reference

- Building Your First Network, [https://hyperledger-fabric.readthedocs.io/en/latest/build_network.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Fbuild_network.html)
- Building Your First Network - v1.4, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/build_network.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fbuild_network.html)
- Prerequisites, [https://hyperledger-fabric.readthedocs.io/en/latest/prereqs.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Fprereqs.html)
- Install Samples, Binaries and Docker Images, [https://hyperledger-fabric.readthedocs.io/en/latest/install.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Finstall.html)
- fabric-samples, [https://github.com/hyperledger/fabric-samples](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fhyperledger%2Ffabric-samples)
- x.509 certificates and public key infrastructure, [https://en.wikipedia.org/wiki/Public_key_infrastructure](https://links.jianshu.com/go?to=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FPublic_key_infrastructure)
- cryptogen, [https://hyperledger-fabric.readthedocs.io/en/latest/commands/cryptogen.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Fcommands%2Fcryptogen.html)
- configtxgen, [https://hyperledger-fabric.readthedocs.io/en/latest/commands/configtxgen.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Fcommands%2Fconfigtxgen.html)

## 项目源代码

项目源代码会逐步上传到 Github，地址为 [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)。

## Contributor

1. Windstamp, [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)



作者：furnace
链接：https://www.jianshu.com/p/bd7d283a5e36
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。



# 2.Hyperledger Fabric 专题 - 核心技术构件

## 共享账本

Hyperledger Fabric 具有一个账本子系统，该子系统包括两个组件：世界状态 (world state) 和交易日志 (transaction log)。每个参与者都有一份他们所属的那个 Hyperledger Fabric 网络的账本副本。

世界状态组件描述帐本在给定时间点的状态。这是帐本的数据库。交易日志组件记录所有导致世界状态当前值的交易，这是世界状态的更新历史。帐本是世界状态数据库和交易日志历史记录的组合。

帐本对于世界状态的数据存储采用可替换机制。默认情况下，这是一个 LevelDB 键值存储数据库。交易日志不需要是可插拔的。它仅记录区块链网络使用的帐本数据库的前后值。

## 智能合约 (Smart Contract)

Hyperledger Fabric 智能合约以链码编写，并在区块链外部的应用程序需要与账本交互时由该应用程序调用。在大多数情况下，链码 (chaincode) 仅与帐本的数据库组件，世界状态 (例如，查询状态) 交互，而不与交易日志交互。

链码可以用几种编程语言实现。当前，支持 Go 和 Node.js 还有 Java。

## 对端节点

Fabric 网络中的账本实例和链码实例都存储在对端节点上，管理员和应用程序都需要通过对端节点与 Fabric 网络交互。对端节点是 Fabric 网络中最重要的物理载体。

## 管理员

具有管理 Fabric 网络的某些权限，能够操作 Fabric 网络。

## 应用程序

外部应用程序，能够交互 Fabric 网络上的链码，或者访问 Fabric 的状态和交易日志。

## 交易排序器 (?)

交易排序器用于 Fabric 网络智能合约执行模型 execute-order-validate 中的 order 部分。

是共识算法中的一部分。

## MSP

MSP (Membership Service Provider) 是 Hyperledger Fabric 中用于达成网络中各成员之间互相信任的构件，类似于比特币中的 POW (Proof Of Work) 机制。

注意，这里 MSP 并不具备达成数据共识的能力，它只是提供了 Fabric 网络中各对端节点之间的互信能力。

## 通道

Hyperledger Fabric 通过通道 (Channel) 允许一组参与者创建单独的交易账本。

这是一个特别重要的选项，因为网络中的某些参与者可能是竞争对手，同时他们也不希望进行的每笔交易 (例如，他们为某些参与者而不是其他人提供的特殊价格) 都被所有参与者知道，因为，只有这些参与者 (没有其他参与者) 拥有该通道的交易账本。

## 隐私性

根据网络的需求，企业对企业 (B2B) 网络的参与者可能对他们共享多少信息非常敏感。对于其他网络，隐私将不是头等大事。

Hyperledger Fabric 支持以隐私 (使用通道) 为主要操作要求的网络以及相对开放的网络。

## 共识机制

交易必须按它们发生的顺序写入账本，即使它们可能会处于网络中由参与者构成的不同的组中。为此，必须建立交易顺序，并且必须采用一种方法来拒绝错误 (或恶意) 插入帐本中的不良交易。

这是计算机科学领域中经过深入研究的领域，有很多方法可以实现，但都需要权衡取舍。例如，实用拜占庭式容错 (Practical Byzantine Fault Tolerance, PBFT) 可以为文件副本复制提供相互通信的机制，即使在出现损坏的情况下，也可以使每个副本保持一致。或者，在比特币中，排序通过称为挖矿的过程进行，在此过程中，竞争的计算机竞相解决一个密码难题，该难题定义了随后所有进程所依据的顺序。

Hyperledger Fabric 旨在允许网络参与者选择一种共识机制，以最能代表参与者之间存在的关系。与隐私一样，也有各种各样的需求。从关系高度结构化的网络到点对点的网络。

Hyperledger Fabric 共识机制目前包括 SOLO (?)，Kafka 和 Raft。

## Hyperledger Fabric 的特性

- 数据存储可以支持多种格式
- 共识机制是可替换的
- 支持不同的 MSP

## Reference

1. Docs » Key Concepts » Introduction, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/blockchain.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fblockchain.html)

## 项目源代码

项目源代码会逐步上传到 Github，地址为 [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)。

## Contributor

1. Windstamp, [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)



作者：furnace
链接：https://www.jianshu.com/p/aee8c7a5879f
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。



# 3. Hyperledger Fabric 专题 - Identity

## 1. 身份识别 (Identity) 概述

区块链网络中的不同参与者包括对端节点 (peer)，交易排序器 (orderer)，客户端应用程序，管理员等。这些参与者 (网络内部或外部能够使用服务的活动元素) 中的每一个都有封装在 X.509 数字证书 (X.509 digital certificate) 中的数字身份 (digital identity)。这些身份确实很重要，因为它们确定了对资源的确切权限以及对参与者在区块链网络中拥有的信息的访问权限。

数字身份还具有一些其他属性，Fabric 可使用这些属性来确定权限，并且 Fabric 为身份和关联的属性的并集定义了特殊名称-主体 (principal)。主体就像用户 ID (userID) 或组 ID (groupID)  一样，但是更加灵活，因为主体可以包含参与者身份的各种属性，例如参与者的组织，组织单位，角色，甚至参与者的特定身份。当我们谈论主体时，它们是确定其权限的属性。

为了使身份可验证，它必须来自受信任的权威机构。成员资格服务提供商 (Membership Service Provider, MSP)  使得这些可以在 Fabric 中实现。更具体地说，MSP 是用于该组织定义管理有效身份规则的组件。 Fabric 中的默认 MSP 实现使用 X.509 证书作为身份标识，并采用传统的公共密钥基础结构 (Public Key Infrastructure, PKI) 层次模型。

## 2. 一个说明如何使用身份识别的简单场景

想象一下，你去一家超市买了一些杂货。在结帐时，你会看到一个标牌，上面只接受 Visa，Mastercard 和 AMEX 卡。如果你尝试使用其他卡付款 (我们称其为 "ImagineCard")，则该卡是否真实以及帐户中是否有足够的资金都无关紧要，因为它们都不会被接受。

拥有有效的信用卡是不够的 - 还必须要商店接受它！ PKI 和 MSP 以相同的方式一起工作 - PKI 提供了一个身份列表，而 MSP 则表示网络中的哪些参与者是参给定组织的成员。

PKI 证书颁发机构和 MSP 提供了类似的功能组合。 PKI 就像卡提供商一样，它分配许多不同类型的可验证身份。另一方面，MSP 就像商店接受的卡提供商列表一样，确定哪些身份是商店支付网络的受信任成员 (参与者)。 MSP 将可验证的身份转换为区块链网络的成员。

## PKI 介绍

PKI 有四个关键要素：

- 数字证书 (Digital Certificates)
- 公钥和私钥 (Public and Private Keys)
- 证书颁发机构 (Certificate Authorities)
- 证书吊销列表 (Certificate Revocation Lists)

关于 PKI 的详细介绍可以参考文档 [What are PKIs?](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fidentity%2Fidentity.html%23what-are-pkis)。

## Fabric CA

因为 CA 非常重要，所以 Fabric 提供了内置的 CA 组件，可让你在自己构建的区块链网络中创建 CA。该组件 (称为 Fabric CA) 是私有的根 CA 提供者，能够管理具有 X.509 证书形式的 Fabric 参与者的数字身份。由于 Fabric CA 是针对 Fabric 的根 CA 需求的自定义 CA，因此它无法提供 SSL 证书供浏览器中的常规/自动使用。但是，由于必须使用某些 CA 来管理身份 (即使在测试环境中也是如此)，因此可以将 Fabric CA 用于提供和管理证书。也可以 (并且完全合适) 使用公共/商业根或中间 CA 进行标识。

更详细的内容可以 [参考 Fabric CA 文档](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric-ca.readthedocs.io%2Fen%2Frelease-1.4%2F)。

既然你已经了解了 PKI 如何通过信任链提供可验证的身份，那么下一步就是了解如何使用这些身份来表示区块链网络的受信任成员。这就是会员服务提供商 (Membership Service Provider, MSP)  发挥作用的地方 - 它确定了在区块链网络中给定组织的成员的参与方。

更详细的内容可以 [参考 MSP 文档](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fmembership%2Fmembership.html)



## Reference

1. Docs » Key Concepts » Identity, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/identity/identity.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fidentity%2Fidentity.html)
2. Docs » Welcome to Hyperledger Fabric CA (Certificate Authority), [https://hyperledger-fabric-ca.readthedocs.io/en/release-1.4/](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric-ca.readthedocs.io%2Fen%2Frelease-1.4%2F)
3. Docs » Key Concepts » Membership, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/membership/membership.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fmembership%2Fmembership.html)

## 项目源代码

项目源代码会逐步上传到 Github，地址为 [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)。

## Contributor

1. Windstamp, [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)



作者：furnace
链接：https://www.jianshu.com/p/af693e235b6e
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。



# 4.Hyperledger Fabric 专题 - PKI

关于 PKI (Public Key Infrastructure) 可以参考文档 [What are PKIs?](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fidentity%2Fidentity.html%23what-are-pkis)。

## Reference

1. What are PKIs?, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/identity/identity.html#what-are-pkis](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fidentity%2Fidentity.html%23what-are-pkis)

## 项目源代码

项目源代码会逐步上传到 Github，地址为 [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)。

## Contributor

1. Windstamp, [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)



作者：furnace
链接：https://www.jianshu.com/p/28adfb32f9ee
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。



# 5.Hyperledger Fabric 专题 - MSP

如果你通读了有关 [身份标识](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fidentity%2Fidentity.html) 的文档，你将了解 PKI 如何通过信任链提供可验证的身份。现在，让我们看看如何使用这些身份标识来表示区块链网络的受信任成员。

这就是成员资格服务提供商 (Membership Service Provider, MSP) 发挥作用的地方 - 它通过列出其成员的身份或通过标识哪些 CA 具有权限为其成员发布有效身份，或者 (通常是这样) 通过将两者结合使用，来确定哪些根 CA 和中间 CA 是可信的并且可以定义可信域 (例如组织) 的成员。

MSP 的功能不仅限于列出谁是网络参与者或通道 (channel) 成员。MSP 可以识别参与者可能在 MSP 代表的组织范围内扮演的特定角色 (例如，管理员或作为子组织组的成员)，并作为基础以在网络和通道上下文中定义访问权限 (例如，通道管理者，读取者，写入者)。

将 MSP 的配置发布到相应组织成员参与的所有通道 (以通道 MSP 的形式)。除了通道 MSP 外，对端节点，交易排序器和客户端还维护本地 MSP，以在通道上下文之外对成员消息进行身份验证，并定义对特定组件的许可权限 (例如，该组件可以在对端节点上安装链码)。

另外，如 [身份标识](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fidentity%2Fidentity.html) 文档中所讨论的，MSP 可以识别已吊销的身份标识列表，但是我们将讨论该过程如何扩展到 MSP。

我们将马上进行关于通道和本地 MSP 的更多讨论。不过，现在让我们先看看 MSP 的一般功能。

## 1. 将 MSP 映射到组织

组织由一组受管理的成员构成。它可以像跨国公司一样大，也可以像花店一样小。对于组织而言，最重要的是，他们在单个 MSP 下管理成员。请注意，这与 X.509 证书中定义的组织概念不同，我们稍后将讨论。

由于组织与其 MSP 之间的排他关系，因此以组织名称命名 MSP 是很合理的，这是大多数策略配置中都会采用的约定。例如，组织 ORG1 可能有一个称为 ORG1-MSP 之类的MSP。在某些情况下，组织可能需要多个成员资格组 - 例如，在组织之间使用通道执行非常不同的业务功能的情况下。在这些情况下，一个组织拥有多个 MSP 并相应地命名它们是有意义的，例如 ORG2-MSP-NATIONAL 和 ORG2-MSP-GOVERNMENT，这反映了与政府监管通道相比，ORA2 在 NATIONAL 销售通道中信任不同的成员资格根源。

![img](https:////upload-images.jianshu.io/upload_images/6280489-be761130dfecc4cc.png?imageMogr2/auto-orient/strip|imageView2/2/w/980/format/webp)

image

上图显示了一个组织两种不同的 MSP 配置。第一个配置显示了 MSP 与组织之间的典型关系 - 单个 MSP 定义组织的成员列表。在第二种配置中，使用不同的 MSP 代表具有国家、国际、和政府从属关系的不同组织组。

## 2. 组织单位和 MSP

一个组织通常划分成多个组织单位 (OU)，每个组织单位都有一定的职责集。例如，组织 ORG1 可能同时具有 ORG1-MANUFACTURING 和 ORG1-DISTRIBUTION  OU，以反映这些单独的业务线。当 CA 颁发 X.509 证书时，证书中的 OU 字段会指定身份所属的业务范围。

稍后我们将看到 OU 如何帮助控制将组织中的哪些部分视为区块链网络的成员。例如，只有来自 ORG1-MANUFACTURING OU 的身份标识才可以访问通道，而 ORG1-DISTRIBUTION 则不能。

最后，尽管这是对 OU 的轻微滥用，但是有时联盟中的不同组织可以使用它们来区分彼此。在这种情况下，不同的组织将相同的根 CA 和中间层 CA 用于其信任链，但会分配 OU 字段以标识每个组织的成员。稍后，我们还将介绍如何配置 MSP 来实现此目的。

## 3. 通道 MSP 和本地 MSP

MSP 出现在区块链网络中的两个位置：通道配置 (通道 MSP) 和角色所在位置的本地 (本地MSP)。本地 MSP 是为客户端 (用户) 和节点（对端节点和交易排序器）定义的。节点本地 MSP 定义了该节点的权限 (例如，对端节点的管理员是谁)。用户的本地 MSP 允许用户端在其交易中作为通道成员 (例如，在链码交易中) 或作为系统中特定角色的所有者 (例如，配置交易组织管理员) 进行身份验证。

每个节点和用户都必须定义一个本地 MSP，因为它定义了谁在该级别具有管理或参与权 (对端节点管理员不一定是通道管理员，反之亦然)。

相反，通道 MSP 在通道级别定义了管理权和参与权。每个参与通道的组织都必须为其定义通道 MSP。通道上的对端节点和交易排序器将共享相同的通道 MSP 视图，因此将能够正确地验证通道参与者。这意味着，如果组织希望加入通道，则需要在通道配置中包含一个包含组织成员信任链的 MSP。否则，来自该组织身份的交易将被拒绝。

本地 MSP 和通道 MSP 之间的主要区别不是它们的功能 (两者都将身份标识转变为角色)，而是它们的范围。

![img](https:////upload-images.jianshu.io/upload_images/6280489-ace42117ef6a989c.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

上图显示了本地和通道 MSP。每个对端节点的信任域 (例如，组织) 由对端节点的本地 MSP (例如 ORG1 或 ORG2) 定义。通过将组织的 MSP 添加到通道配置中，可以实现在通道上代表组织。例如，此图的通道同时由 ORG1 和 ORG2 管理。类似的原理适用于网络，交易排序器和用户，但为简单起见，此处未显示。

你可能会发现，通过查看区块链管理员安装和实例化智能合约时会发生什么，了解使用本地和通道 MSP 很有帮助，如上图所示。

管理员 B 以 RCA1 发布的身份标识连接到对端节点，并存储在其本地 MSP 中。当 B 尝试在对端节点上安装智能合约时，对端节点会检查其本地 MSP ORG1-MSP，以验证 B 的身份确实是 ORG1 的成员。成功的验证将使安装命令成功完成。随后，B 希望在通道上实例化智能合约。因为这是通道操作，所以通道上的所有组织都必须对此表示同意。因此，对端节点必须先检查通道的 MSP，然后才能成功提交此命令。(其他事情也必须发生，但目前就集中在上面提到的内容。)

本地 MSP 仅在它们应用到的节点或用户的文件系统上定义。因此，在物理上和逻辑上，每个节点或用户只有一个本地 MSP。但是，由于通道 MSP 可用于通道中的所有节点，因此它们在通道配置中逻辑定义一次。但是，通道 MSP 也会在该通道中每个节点的文件系统上实例化，并通过共识保持同步。因此，尽管在每个节点的本地文件系统上都有每个通道 MSP 的副本，但从逻辑上讲，通道 MSP 驻留在通道或网络上并由通道或网络维护。

## 4. MSP 级别

MSP 划分为通道 MSP 和本地 MSP 反映了组织管理的需求，如管理本地资源 (例如对端节点或交易排序器节点)，通道资源 (例如帐本，智能合约)，和在通道或网络级别运行的联盟。可以将这些 MSP 划分为不同的级别，这会很有帮助，较高级别的 MSP 与网络管理相关，而较低级别的 MSP 处理身份管理的私有资源。 MSP 在每个管理级别都是必需的 - 必须为网络，通道，对端节点，交易排序器和用户定义 MSP。

![img](https:////upload-images.jianshu.io/upload_images/6280489-59ef37ff8523f55b.png?imageMogr2/auto-orient/strip|imageView2/2/w/1015/format/webp)

image

上图显示了 MSP 级别。对端节点和交易排序器的 MSP 是本地的，而某个通道 (包括网络配置通道) 的 MSP 在该通道的所有参与者之间共享。在此图中，网络配置通道由 ORG1 管理，但另一个应用程序通道可以由 ORG1 和 ORG2 管理。对端节点是 ORG2 的成员并由其管理，而 ORG1 管理交易排序器。 ORG1 信任来自 RCA1 的身份，而 ORG2 信任来自 RCA2 的身份。请注意，这些是管理身份，反映了谁可以管理这些组件。因此，当 ORG1 管理网络时，ORG2.MSP 确实存在于网络定义中。

- 网络 MSP (Network MSP)：网络的配置通过定义参与组织的 MSP 来定义谁是网络中的成员，以及授权这些成员中的哪些成员可以执行管理任务 (例如，创建通道)。
- 通道 MSP (Channel MSP)：通道要分别维护其成员的 MSP，这一点很重要。通道在一组特定的组织之间提供私人通信，而这些组织又对该组织具有管理控制权。在该通道的 MSP 上下文中解释的通道政策定义了谁有能力参与该通道上的某些操作，例如添加组织或实例化链码。请注意，管理通道的权限与管理网络配置通道 (或任何其他通道) 的能力之间没有必要的关系。管理权限存在于所管理的范围内 (除非规则另有规定，请参见下面有关 ROLE 属性的讨论)。
- 对端节点 MSP (Peer MSP)：此本地 MSP 在每个对端节点的文件系统上定义，并且每个对端节点都有一个 MSP 实例。从概念上讲，它执行与通道 MSP 完全相同的功能，但限制是它仅适用于定义它的对端节点。使用对端节点的本地 MSP 评估授权的操作示例是在对端节点上安装链码。
- 交易排序器 MSP (Order MSP)：与对端节点 MSP 一样，交易排序器本地 MSP 也定义在节点的文件系统上，并且仅适用于该节点。像对端节点一样，交易排序器也由单个组织拥有，因此具有单个 MSP 来列出其信任的参与者或节点。

## 5. MSP 结构

到目前为止，你已经看到，MSP 的最重要元素是用于建立相应组织中参与者或节点成员身份的根或中间 CA 的规范。但是，还有更多元素与这两个元素结合使用以辅助成员资格功能。

![img](https:////upload-images.jianshu.io/upload_images/6280489-c36a6c6872dffc08.png?imageMogr2/auto-orient/strip|imageView2/2/w/838/format/webp)

image

上图显示了本地 MSP 如何存储在本地文件系统上。即使通道 MSP 的物理结构不是完全按照这种方式构成的，但仍是考虑它们的有用方法。

如你所见，MSP 有九个要素。在目录结构中最容易想到这些元素，其中 MSP 名称是根文件夹名称，每个子文件夹代表 MSP 配置的不同元素。

让我们更详细地描述这些文件夹，看看它们为什么很重要。

- 根 CA (Root CAs)：此文件夹包含由此 MSP 代表的组织信任的根 CA 的自签名 X.509 证书的列表。此 MSP 文件夹中必须至少有一个 Root CA X.509 证书。

这是最重要的文件夹，因为它标识了必须从其派生所有其他证书才能被视为相应组织的成员的 CA。

- 中间层 CA (Intermediate CAs)：此文件夹包含此组织信任的中间层 CA 的 X.509 证书的列表。每个证书必须由 MSP 中的一个根 CA 签名，或由其发行 CA 链最终引回到受信任的根 CA 的中间层 CA 签名。

中间层 CA 可能代表组织的不同部门 (例如 ORG1 的 ORG1-MANUFACTURING 和 ORG1-DISTRIBUTION 所做的事情)，也可能代表组织本身 (如果商业 CA 用于组织的身份管理，则可能是这种情况)。在后一种情况下，中间层 CA 可以用来表示组织细分。在 [这里](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fmsp.html)，你可以找到有关 MSP 配置最佳实践的更多信息。请注意，可能有一个没有中间层 CA 的正常运行的网络，在这种情况下，此文件夹将为空。

与 “根 CA” 文件夹类似，此文件夹定义了必须颁发证书才能被视为组织成员的 CA。

- 组织单位 (Organizational Units, OUs)：这些列在 `$ FABRIC_CFG_PATH/msp/config.yaml` 文件中，并包含一个组织单位列表，其成员被认为是此 MSP 代表的组织的一部分。当你希望将组织的成员限制为持有具有特定 OU 的身份 (由 MSP 指定的 CA 之一签名) 的成员时，此功能特别有用。

指定 OU 是可选的。如果未列出任何 OU，则 MSP 一部分的所有身份 (由根 CA 和中间层 CA 文件夹标识) 将被视为组织的成员。

- 管理员 (Administrators)：此文件夹包含身份标识列表，这些身份标识定义了具有该组织管理员角色的参与者。对于标准 MSP 类型，此列表中应该有一个或多个 X.509 证书。

值得注意的是，仅仅因为参与者具有管理员角色，并不意味着他们可以管理特定资源！给定身份在管理系统方面的实际能力由管理系统资源的策略确定。例如，通道策略可能指定 ORG1-MANUFACTURING 管理员有权向该通道添加新组织，而 ORG1-DISTRIBUTION 管理员则无此权利。

即使 X.509 证书具有 `ROLE` 属性 (例如，指定参与者是管理员)，也指的是参与者在其组织内而不是在区块链网络上的角色。这类似于 OU 属性的用途，如果已定义，则它是指参与者在组织中的位置。

如果已为该通道编写了策略，以允许组织 (或某些组织) 中的任何管理员执行某些通道功能 (例如实例化链码) 的权限，则可以使用 `ROLE` 属性在该通道级别授予管理权限。这样，组织角色可以赋予网络角色。

- 吊销证书 (Revoked Certificates)：如果参与者的身份标识已被吊销，则有关身份的标识信息 (而不是身份本身) 保存在此文件夹中。对于基于 X.509 的身份标识，这些标识符是称为主题密钥标识符 (Subject Key Identifier, SKI) 和授权访问标识符 (Authority Access Identifier, AKI) 的字符串对，并且只要使用 X.509 证书就对其进行检查以确保该证书未被撤销。

此列表在概念上与 CA 的证书吊销列表 (Certificate Revocation List, CRL) 相同，但也与组织的成员资格吊销有关。结果，MSP 的管理员 (本地或通道) 可以通过向 CA 的更新的 CRL 发布由其颁发的吊销证书来快速吊销组织中的参与者或节点。此“列表列表”是可选的。仅当证书被吊销时，它才会被填充。

- 节点身份 (Node Identity)：此文件夹包含节点的标识身份，即，与 `KeyStore` 的内容结合使用的加密材料，将允许节点在发送到其通道和网络其他参与者的消息中对自身进行身份验证。对于基于 X.509 的身份，此文件夹包含 X.509 证书。这是对端节点在交易提议响应中放置的证书，例如，表明对端节点已经背书了 (endorse) 该证书 - 随后可以在验证时根据结果交易的背书策略 (endorsement policy) 检查该证书。

对于本地 MSP，此文件夹是必需的，并且该节点必须完全有一个 X.509 证书。它不用于通道 MSP。

- 专用密钥的密钥库 (KeyStore)：此文件夹是为对端节点或交易排序器节点的本地 MSP (或在客户端的本地 MSP中) 定义的，并且包含节点的签名密钥。该密钥在密码上匹配包含在 “节点标识” 文件夹中的节点的标识，并用于对数据进行签名 - 例如，作为签署阶段的一部分，对交易建议响应进行签名。

此文件夹对于本地 MSP 是必需的，并且必须仅包含一个私钥。显然，对此文件夹的访问必须仅限于对对端节点具有管理责任的用户的身份。

通道 MSP 的配置不包括此文件夹，因为通道 MSP 仅旨在提供身份验证功能而不是签名功能。

- TLS 根 CA (TLS Root CA)：此文件夹包含此组织信任用于 TLS 通信的根 CA 的自签名 X.509 证书的列表。 TLS 通信的一个示例是当对端节点需要连接到交易排序器以便接收账本更新时。

MSP TLS 信息与网络内部的节点有关 - 换句话说，对端节点和交易排序器，而不是消耗网络的应用程序和管理。

此文件夹中至少必须有一个 TLS 根 CA X.509 证书。

- TLS 中间层 CA (TLS Intermediate CA)：此文件夹包含此 MSP 代表的组织用于 TLS 通信的中间层 CA 证书 CA 列表。当商业 CA 用于组织的 TLS 证书时，此文件夹特别有用。与成员资格中间层 CA 相似，指定中间层 TLS CA 是可选的。

有关 TLS 的更多信息，请单击 [此处](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fenable_tls.html)。

如果你已经阅读过本文档以及我们的有关 [身份标识](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fidentity%2Fidentity.html) 的文档，那么你应该对身份标识和成员身份在 Hyperledger Fabric 中的工作方式有很好的了解。你已经了解了如何使用 PKI 和 MSP 来识别在区块链网络中进行协作的参与者。你已经了解了证书，公钥/私钥和信任根的工作原理，以及 MSP 的物理和逻辑结构。

## Reference

1. Docs » Key Concepts » Membershipm, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/membership/membership.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fmembership%2Fmembership.html)
2. Docs » Key Concepts » Identity, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/identity/identity.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fidentity%2Fidentity.html)
3. Docs » Operations Guides » Membership Service Providers (MSP), [https://hyperledger-fabric.readthedocs.io/en/release-1.4/msp.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fmsp.html)
4. Docs » Operations Guides » Securing Communication With Transport Layer Security (TLS), [https://hyperledger-fabric.readthedocs.io/en/release-1.4/enable_tls.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fenable_tls.html)

## 项目源代码

项目源代码会逐步上传到 Github，地址为 [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)。

## Contributor

1. Windstamp, [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)



作者：furnace
链接：https://www.jianshu.com/p/389801c31c07
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。



# 6.Hyperledger Fabric 专题 - Peers

区块链网络主要由一组对端节点 (peer nodes, or simply, peers) 组成。对端节点是网络的基本元素，因为它们托管帐本和智能合约。回想一下，账本一成不变地记录了智能合约生成的所有交易 (Hyperledger Fabric 中的交易包含在链码中，稍后会详细介绍)。智能合约和账本分别用于封装网络中的共享过程 (shared process) 和共享信息 (shared information)。对端节点的这些方面使它们成为了解 Fabric 网络的良好起点。

区块链网络的其他元素当然很重要：帐本和智能合约，交易排序器，策略，通道，应用程序，组织，身份标识和成员资格服务提供商，你可以在其专用部分中阅读有关它们的更多信息。本节重点介绍对端节点及其与 Fabric 网络中其他元素的关系。

![img](https:////upload-images.jianshu.io/upload_images/6280489-997eb29fe6024e0d.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

区块链网络由对端节点组成，每个对端节点可以保存账本的副本和智能合约的副本。在此示例中，网络 N 由对端节点 P1，P2 和 P3 组成，每个对端节点 P1，P2 和 P3 维护自己的分布式账本 L1 实例。 P1，P2 和 P3 使用相同的链码 S1 来访问其分布式账本的副本。

可以创建，启动，停止，重新配置甚至删除对端节点。它们公开了一组 API，使管理员和应用程序可以与其提供的服务进行交互。我们将在本节中详细了解这些服务。

## 1. 术语

Fabric 通过称为链码 (chaincode) 的技术概念来实现智能合约(smart contract) - 只是访问账本的一段代码，以受支持的编程语言中的任一种编写。在本主题中，我们通常使用术语“链码”，但是如果你更习惯智能合约这个术语，可以随时将其作为智能合约来阅读。这是同一件事！如果你想了解有关链码和智能合约的更多信息，请查看我们关于智能合约和链码的文档。

## 2. 分布式账本和链码

让我们再详细一点。我们可以看到，同时托管帐本和链码的是对端节点。更准确地说，对端节点实际上是托管帐本实例和链码实例的宿主。请注意，这在 Fabric 网络中提供了有意的冗余 - 避免了单点故障。在本节的后面，我们将详细了解区块链网络的分布式和去中心化性质。

[图片上传失败...(image-4d3467-1573120639360)]

对端节点托管账本实例和链码实例。在此示例中，P1 承载帐本 L1 的实例和链码 S1 的实例。单个对端节点主机上可以托管许多帐本和链码。

由于对端节点是帐本和链码的宿主，因此应用程序和管理员如果要访问这些资源，必须与对端节点进行交互。这就是为什么将对等方视为Fabric网络的最基本组成部分。首次创建同位体时，既没有分类帐，也没有链码。稍后我们将介绍如何在同级上创建分类帐以及如何安装链码。

### 2.1 多个账本

对端节点可以承载多个帐本，这很有用，因为它允许灵活的系统设计。最简单的配置是让对端节点管理一个帐本，但是对端节点在需要时托管两个或多个帐本绝对是合适的。

![img](https:////upload-images.jianshu.io/upload_images/6280489-1645df0f60a25f74.png?imageMogr2/auto-orient/strip|imageView2/2/w/796/format/webp)

image

承载多个帐本的对端节点。对端节点主机托管一个或多个帐本，每个帐本具有应用于它们的零个或多个链码。在此示例中，我们可以看到对端节点 P1 承载帐本 L1 和 L2。使用链码 S1 访问帐本 L1。另一方面，可以使用链码 S1 和 S2 访问帐本 L2。

尽管对端节点完全有可能托管一个账本实例，而不托管任何访问该账本的链码，但很少有对端节点是以这种方式配置的。绝大多数对端节点将至少安装一个链码，可以查询或更新对等点的分类帐实例。值得一提的是，无论用户是否安装了供外部应用程序使用的链码，对端节点也都始终存在着特殊的系统链码 (system chaincode)。在本主题中不会详细讨论这些内容。

### 2.2 多个链码

对端节点拥有的账本数量与可以访问该账本的链码数量之间没有固定的关系。一个对端节点可能有许多可用的链码和帐本。

![img](https:////upload-images.jianshu.io/upload_images/6280489-fa0cc48d07039a1d.png?imageMogr2/auto-orient/strip|imageView2/2/w/797/format/webp)

image

上图是一个对端节点托管多个链码的示例。每个帐本可以有许多访问它的链码。在此示例中，我们可以看到对端节点 P1 承载帐本 L1 和 L2，其中 L1 由链码 S1 和 S2 访问，而 L2 由 S1 和 S3 访问。我们可以看到 S1 可以访问 L1 和 L2。

我们稍后将看到当在对端节点上托管多个帐本或多个链码时，Fabric 中的通道概念为什么如此重要的原因。

## 3. 应用程序和对端节点

现在我们将解释应用程序是如何通过与对端节点交互来访问账本的。账本查询交互涉及一个应用程序和一个对端节点之间简单的三步对话。账本更新交互稍微复杂一点，需要两个额外的步骤。我们将简化这些步骤一点点地帮助你开始使用 Fabric，但不要担心 - 最重要的是理解应用程序和对端节点交互在账本查询和账本更新之间的差异。

当应用程序需要访问帐本和链码时，它们总是连接到对端节点。Fabric 的软件开发套件 (Software Development Kit, SDK) 使程序员可以轻松实现它 - 其 API 使应用程序可以连接到对端节点，调用链码以生成交易，将交易提交到网络 (该交易将被排序并提交到分布式帐本) 以及当此过程完成时接收事件。

通过与一个对端节点连接，应用程序可以执行链码来查询或更新帐本。帐本查询交易的结果将立即返回，而帐本更新涉及应用程序，对端节点和交易排序器之间更复杂的交互。让我们更详细地研究一下。

![img](https:////upload-images.jianshu.io/upload_images/6280489-3dfc17d3e53472c7.png?imageMogr2/auto-orient/strip|imageView2/2/w/844/format/webp)

image

上图显示，对端节点通过与交易排序器连接以确保每个对端节点上的账本都是最新的。在此示例中，应用程序 A 连接到 P1 并调用链码 S1 以查询或更新帐本 L1。P1 调用 S1 以生成包含查询结果或帐本更新提案的响应。应用程序 A 收到提案响应，并且对于查询该过程现在已完成。为了进行更新，A 根据所有响应构建一个交易，然后将其发送给 O1 进行交易排序。 O1 将整个网络中的交易收集到多个区块中，并将其分配给所有的对端节点，包括 P1。 P1 在将交易应用到 L1 之前先进行验证。一旦 L1 更新，P1 就会生成一个事件，该事件由 A 接收，表示整个流程已经完成。

对端节点可以立即将查询结果返回给应用程序，因为满足查询所需的所有账本信息都在对端节点的本地副本中。对端节点从不需要与其他对端节点协商以响应来自应用程序的查询。但是，应用程序可以连接到一个或多个对端节点以发出查询。例如，以在多个对端节之间证实其结果，或如果怀疑信息有可能过时则可以从不同的对端节点获取最新结果。在上图中，你可以看到帐本查询是一个简单的三步过程。

更新交易以与查询交易相同的方式启动，但有两个额外的步骤。尽管更新帐本的应用程序也连接到对端节点以调用链码，但与帐本查询应用程序不同，单个对端节点此时无法执行帐本更新，因为其他对端节点必须首先同意更改，这一过程称为共识。因此，对端节点向应用程序返回了一项提案 (proposed) 的更新 - 该对端节点将在其他对端节点事先同意的情况下应用该更新。第一个额外的步骤 (第四步) 要求应用程序将一组适当的匹配的提案更新集发送到整个网络，以作为对端节点对各自账本的承诺的交易。这是通过应用程序使用交易排序器将交易打包成区块并将其分发到整个网络来实现的，在将它们应用于每个对端节点的本地帐本副本之前，可以在其中进行验证。由于整个交易排序器过程需要一些时间 (几秒钟) 才能完成，因此将异步通知应用程序，如步骤 5 所示。

在本节的后面，你将了解有关此交易排序器过程的详细特性的更多信息 - 有关此过程的详细信息，请参见 [交易流](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Ftxflow.html) 主题。

## 4. 对端节点和通道

尽管本节的重点是对端节点而不是通道，但值得花一些时间来了解对端节点如何通过通道与彼此以及与应用程序进行交互。通过通道，区块链网络中的一组组件可以进行私下通信和交易。

这些组件通常是对端节点，交易排序器节点和应用程序，并且通过加入通道，它们同意协作以共同共享和管理与该通道关联的帐本的相同副本。从概念上讲，你可以将通道视为与朋友组相似 (尽管通道的成员当然不需要成为朋友！)。一个人可能有几组朋友，每组都有他们一起做的活动。这些小组可能是完全独立的 (一群工作朋友，而不是一群爱好朋友)，或者它们之间可能会有一些交叉。但是，每个组都是其自己的实体，具有某种“规则”。

![img](https:////upload-images.jianshu.io/upload_images/6280489-c8476d9b40e95f1f.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

上图显示，通道允许一组特定的对端节点和应用程序在区块链网络内相互通信。在此示例中，应用程序 A 可以使用通道 C 直接与对端节点 P1 和 P2 通信。你可以将通道视为特定应用程序和对端节点之间进行通信的路径。 (为简单起见，交易排序器未在此图中显示，但必须存在于正常运行的网络中。)

我们发现通道的存在方式与对端节点不同，因此将通道视为由一组物理对端节点形成的逻辑结构更为合适。理解这一点至关重要 - 对端节点提供访问和管理通道的控制点。

## 5. 对端节点和组织

既然你了解了对端节点及其与帐本，链码和通道的关系，现在你将能够看到多个组织如何一起组成一个区块链网络。

区块链网络由组织的集合而不是单个组织管理。对端节点对于这种分布式网络的构建至关重要，因为对端节点网络由这些组织拥有，并且是这些组织的网络连接点。

![img](https:////upload-images.jianshu.io/upload_images/6280489-3fcf3c9d03d7f2c4.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

上图显示了，具有多个组织的区块链网络中的对端节点。区块链网络是由拥有不同的对端节点的组织建立和贡献的。在此示例中，我们看到四个组织贡献了八个对端节点组成一个网络。通道 C 连接网络 N 中的五个对端节点 - P1，P3，P5，P7 和 P8。这些组织拥有的其他对端节点尚未加入此通道，但通常已加入至少一个其他通道。由特定组织开发的应用程序将连接到其各自组织的对端节点以及不同组织的对端节点。同样，为简单起见，该图中未显示交易排序器节点。

了解区块链网络的形成过程非常重要。网络由向其贡献资源的多个组织组成和管理。对端节点是我们在本主题中讨论的资源，但是组织提供的资源不仅仅是对端节点。这里有一个原则在起作用 - 如果组织没有将各自的资源投入到集体网络中，那么网络实际上就不存在。而且，网络随着这些协作组织提供的资源而增长和收缩。

在上面的示例中，你可以看到 (除了交易排序器服务之外) 没有集中式资源 - 如果组织不提供对端节点资源，则网络 N 将不存在。这反映了这样一个事实：除非有组织提供构成该网络的资源，否则该网络不存在任何意义。此外，网络不依赖于任何单个组织，只要一个组织存在，它就会继续存在，无论其他组织可能来去去去。这是网络去中心化意味着什么的核心。

如上例所示，不同组织中的应用程序可能相同也可能不同。那是因为这完全取决于组织的应用程序如何处理其对端节点的帐本副本。这意味着应用程序和表示逻辑在组织之间可能会有所不同，即使它们各自的对端节点托管的账本数据完全相同。

应用程序可以连接到其组织中的对端节点，也可以连接到另一个组织中的对端节点，具体取决于所需的账本交互的性质。对于账本查询交互，应用程序通常会连接到其组织的对端节点。对于账本更新交互，我们将在后面看到为什么应用程序需要连接到代表背书 (endorse) 账本更新所需的每个组织的对端节点。

## 6. 对端节点和身份标识

既然你已经了解了来自不同组织的对端节点如何组成一个区块链网络，那么值得花一些时间来了解如何将其对端节点由管理员分配给组织。

对端节点具有通过特定证书颁发机构通过数字证书分配给他们的身份标识。你可以在本指南的其他地方阅读有关 X.509 数字证书如何工作的更多信息，但就目前而言，数字证书就像是 ID 卡，可以提供有关对端节点的大量可验证信息。网络中的每个对端节点都由其所属组织的管理员分配了数字证书。

![img](https:////upload-images.jianshu.io/upload_images/6280489-fc9b77632e0704eb.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

当对端节点连接到一个通道，它的数字证书通过通道 MSP 标识其所属的组织。在此示例中，P1 和 P2 是由有 CA1 颁发的身份标识。通道 C 根据其通道配置中的策略，确定应该使用 ORG1.MSP 将来自 CA1 的身份标识与 Org1 关联。同样，P3 和 P4 被 ORG2.MSP 标识为 Org2 的一部分。

每当对端节点使用通道连接到区块链网络时，通道配置中的策略都会使用对端节点的身份标识来确定其权利。身份标识到组织的映射由称为成员资格服务提供商 (Membership Service Provider, MSP) 的组件提供 - 它确定如何将对端节点分配给特定组织中的特定角色，并因此获得对区块链资源的适当访问。此外，对端节点只能由单个组织拥有，因此与单个 MSP 关联。我们将在本节的后面部分详细了解对端节点访问控制，并且在本指南的其他部分中将有一个完整的关于 MSP 和访问控制策略的部分。但就目前而言，可以将 MSP 视为在区块链网络中的个人身份和特定组织角色之间提供链接。

离题一会儿，对端节点以及一切从他们的数字证书和 MSP 与区块链网络交互获取其组织的身份。对端节点，应用程序，终端用户，管理员和交易排序器，必须有一个身份标识和相关的 MSP，如果他们想与区块链网络交互。我们将使用身份标识与区块链网络交互的每个实体命名为主体 (principal)。你可以在本指南中的其他地方了解更多有关主体和组织的信息，但是到目前为止，你所了解的知识还远远不够，可以继续理解对端节点！

最后，请注意，对端节点的物理位置并不重要 - 它可以位于云中，也可以位于组织之一拥有的数据中心中，也可以位于本地计算机上 - 与之相关联的身份标识将其标识为由特定组织拥有。在上面的示例中，P3 可以托管在 Org1 的数据中心中，但是只要与它关联的数字证书由 CA2 颁发，那么它就由 Org2 拥有。

## 7. 对端节点和交易排序器

我们已经看到，对端节点构成了区块链网络，托管帐本和智能合约的基础，可以由对端节点连接的应用程序查询和更新。但是，应用程序和对端节点彼此交互以确保每个对端节点的帐本保持一致的机制是由称为交易排序器 (orderer) 的特殊节点来确保的，现在我们转向这些节点。

更新交易与查询交易完全不同，因为单个对端节点自身无法更新帐本 - 更新需要网络中其他对端节点的同意。对端节点要求网络中的其他对端节点批准帐本更新，然后才能将其应用于对端节点的本地帐本。此过程称为共识，与简单查询相比，此过程需要更长的时间才能完成。但是，当所有需要批准交易的对端节点都批准了该交易并将交易提交到帐本时，对端节点将通知其连接的应用程序帐本已更新。在本部分中，你将获得有关对端节点和交易排序器如何管理共识过程的更多详细信息。

具体来说，想要更新账本的应用程序涉及 3 个阶段的过程，这确保了区块链网络中的所有对端节点都保持其账本彼此一致。在第一阶段，应用程序与背书对端节点 (endosring peer) 的子集一起工作，每个背书对端节点都向应用程序提供对账本更新提案的背书，但不将提案的更新应用于其账本副本。在第二阶段，将这些单独的背书 (endorsement) 作为交易收集在一起，并打包成区块。在最后阶段，这些区块被分配回每个对端节点，在将每个交易应用于该对端节点的帐本副本之前，都已经对其进行了验证。

正如你将看到的，交易排序器节点是此过程的核心，所以让我们更详细地研究一下应用程序和对端节点如何使用交易排序器来生成账本更新，这些更新可以一致地应用于分布式复制账本。

### 7.1 第 1 阶段 - 提案

交易工作流程的第 1 阶段涉及应用程序和一组对端节点之间的交互 - 它不涉及交易排序器。第 1 阶段仅涉及一个应用程序，这个应用程序请求不同组织的背书对端节点 (endorsing peer) 同意提案的链码调用结果。

从第 1 阶段开始，应用程序会生成一个交易提案，然后将其发送给每个必需的对端节点组以进行背书。然后，这些背书对端节点中的每一个都使用交易提案独立执行链码以生成交易提案响应。它不会将此更新应用于帐本，而只是对其进行签名并将其返回给应用程序。一旦应用程序收到了足够数量的已签名提案响应，交易流程的第 1 阶段便完成了。让我们更详细地研究这个阶段。

![img](https:////upload-images.jianshu.io/upload_images/6280489-b5645b75386f9435.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

交易提案由返回提案响应背书的对端节点独立执行。在此示例中，应用程序 A1 生成交易 T1 提议 P，并将其发送到通道 C 上的对端节点 P1 和对端节点 P2。P1 使用交易 T1 提案 P 执行 S1 生成交易 T1 响应 R1 并由 E1 背书。P2 独立地使用交易 T1 提案 P 执行 S1，并生成交易 T1 响应 R2，并由 E2 背书。应用程序 A1 收到两个对交易 T1 的响应背书，即 E1 和 E2。

在初始化阶段，应用程序选择一组对端节点以生成一组提案的账本更新。应用程序选择了哪些对端节点？好吧，这取决于背书策略 (为链码定义)，背书策略 (endorsement policy) 定义了需要在网络接受之前对提案的账本更改进行背书的组织集合。从字面上看，这就是达成共识的意思 - 每个重要组织都必须已批准提案的帐本更改，然后才能将其接受到任何对端节点的帐本中。

对端节点通过添加其数字签名并使用其私钥对整个有效负载进行签名来背书提案响应。此背书可以随后用于证明该组织的对端节点产生了特定的响应。在我们的示例中，如果对端节点 P1 由组织 Org1 拥有，则背书 E1 对应于一个数字证明，即 “由 Org1 的对端节点 P1 提供了账本 L1 上的交易 T1 的响应 R1！”。

当应用程序收到足够的对端节点签署的提案响应时，第 1 阶段结束。我们注意到，对于同一个交易提案，不同的对端节点可以向应用程序返回不同的因此不一致的交易响应。可能只是结果是在不同时间使用不同状态的帐本在不同的对端节点上生成的，在这种情况下，应用程序可以简单地请求更新的提案响应。虽然出现这种结果的可能性较小，但更严重的是结果可能会有所不同，因为链码是不确定的。非确定性是链码和账本的敌人，如果发生，则表明提案的交易存在严重问题，因为显然不能将不一致的结果应用于账本。单个对端节点无法知道他们的交易结果是不确定的，因此必须先收集交易响应以进行比较，然后才能检测到不确定性。(严格来说，这还不够，但是我们将讨论推迟到交易部分，在此部分将详细讨论不确定性。)

在第 1 阶段结束时，应用程序可以随意丢弃不一致的交易响应，从而有效地尽早终止交易工作流。稍后我们将看到，如果应用程序尝试使用一组不一致的交易响应来更新帐本，则它将被拒绝。

### 7.2 第 2 阶段 - 将交易排序和打包成区块

交易工作流程的第 2 阶段是打包阶段。交易排序器对于此过程至关重要 - 它从许多应用程序接收包含交易提案响应背书的交易，并将交易排序进区块。有关排序和打包阶段的更多详细信息，请查看我们有关 [排序阶段的概念性信息](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Forderer%2Fordering_service.html%23phase-two-ordering-and-packaging-transactions-into-blocks)。

### 7.3 第 3 阶段 - 验证和提交

在第 2 阶段结束时，我们看到交易排序器负责简单但至关重要的过程，这些过程包括收集提案的交易更新，进行排序并将它们打包成区块，以便分发给对端节点。

交易工作流程的最后阶段涉及从交易排序器到对端节点的区块分发和后续验证，在这里可以将它们应用于帐本。具体来说，在每个对端节点，对一个区块中的每个交易都进行验证，以确保在将其应用于账本之前，所有相关组织都一致认可该交易。失败的交易将保留以进行审核，但不会应用于帐本。

![img](https:////upload-images.jianshu.io/upload_images/6280489-ed349f51f8b3c875.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

交易排序器节点的第二个作用是将区块分配给对端节点。在该示例中，交易排序器 O1 将区块 B2 分配给对端节点 P1 和对端节点 P2。对端节点 P1 处理区块 B2，从而将新区块添加到 P1 上的帐本 L1。并行地，对端节点 P2 处理区块 B2，导致将新区块添加到 P2 上的帐本 L1。一旦该过程完成，账本 L1 就已在对端节点 P1 和 P2 上进行了一致更新，并且每个账本可以通知连接的应用程序交易已被处理。

第 3 阶段从交易排序器将区块分配给与其连接的所有对端节点开始。对端节点通过通道连接到交易排序器，以便在生成新区块时，将向连接到交易排序器的所有对端节点发送新区块的副本。每个对端节点将独立地但与通道上的每个其他对端节点以完全相同的方式处理此区块。这样，我们将看到帐本可以保持一致。还值得注意的是，并非每个对端节点都需要连接到交易排序器，对端节点可以使用 gossip 协议将区块级联到其他对端节点，后者(?)也可以独立处理它们。但是，让我们将讨论留给其他时间！

收到区块后，对端节点将按照它出现在区块中的顺序处理每个交易。对于每笔交易，每个对端节点将根据生成交易的链码的背书策略，验证该交易是否已被所需的组织背书。例如，某些交易可能只需要由单个组织背书，而其他交易可能需要多次背书才能被视为有效。此验证过程将验证所有相关组织均已产生相同的输出或结果。还要注意，此验证与第 1 阶段中的背书检查不同，在第 1 阶段中，应用程序是从背书对端节点接收响应并做出发送提案交易的决定。如果应用程序通过发送错误的交易违反了背书策略，则对端节点仍然可以在第 3 阶段的验证过程中拒绝交易。

如果交易已正确背书，则对端节点将尝试将其应用于帐本。为此，对端节点必须执行帐本一致性检查，以验证生成提案更新时帐本的当前状态与帐本的状态兼容。即使交易已得到完全认可，这也不总是可能的。例如，另一笔交易可能已更新账本中的同一资产，因此该交易更新不再有效，因此无法再应用。这样，由于每个对端节点遵循相同的验证规则，因此它们在整个网络中保持一致。

对端节点成功验证了每个单独的交易后，它将更新帐本。失败的交易不应用于帐本，但保留它们以进行审核，与成功的交易一样。这意味着对端节点区块几乎与从交易排序器接收到的区块完全相同，除了该区块中每个交易的有效或无效指示符。

我们还注意到，第 3 阶段不需要运行链码 - 仅在第 1 阶段已经完成，这很重要。这意味着链码仅在背书节点上可用，而不是在整个区块链网络上可用。这通常很有用，因为它可以使链码的逻辑对背书组织保密。这与链码的输出 (交易提案响应) 相反，链码的输出与通道中的每个对端节点共享，无论他们是否背书交易。支持对端节点的这种专业化旨在帮助实现可伸缩性。

最后，每次将一个区块提交给对端节点的帐本时，该对端节点都会生成一个适当的事件。区块事件包括完整的区块内容，而区块交易事件仅包含摘要信息，例如区块中的每个交易是否已验证或无效。链码执行产生的链码事件也可以在此时发布。应用程序可以注册这些事件类型，以便在事件发生时得到通知。这些通知结束了交易工作流的第 3 阶段即最后阶段。

总而言之，第 3 阶段将看到由交易排序器生成的区块始终应用于帐本。严格按顺序将交易划分为区块，每个对端节点都可以验证交易更新是否在整个区块链网络中得到一致应用。

### 7.4 交易排序器和共识

交易流程的整个过程称为共识，因为在交易排序器的调解下，所有对端节点都已就交易的顺序和内容达成了共识。共识是一个多步骤的过程，仅当过程完成时才通知应用程序账本更新，这可能在不同的对端节点上略有不同。

我们将在以后的交易排序器主题中更详细地讨论交易排序器，但现在，将交易排序器视为从应用程序收集和分发提案的账本更新以供对端节点验证并包含在账本中的节点。

就是这样！现在，我们已经完成了对对端节点及其与 Fabric 相关的其他组件的浏览。我们已经看到，在许多方面，对端节点是最基本的元素 - 它们形成网络，托管链码和帐本，处理交易提案和响应，并通过始终向其应用交易更新来使帐本保持最新。

## Reference

- Docs » Key Concepts » Peers, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/peers/peers.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fpeers%2Fpeers.html)
- Docs » Key Concepts » Smart Contracts and Chaincode, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/smartcontract/smartcontract.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fsmartcontract%2Fsmartcontract.html)
- Docs » Architecture Reference » Transaction Flow, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/txflow.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Ftxflow.html)
- Docs » Key Concepts » The Ordering Service, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/orderer/ordering_service.html#phase-two-ordering-and-packaging-transactions-into-blocks](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Forderer%2Fordering_service.html%23phase-two-ordering-and-packaging-transactions-into-blocks)

## 项目源代码

项目源代码会逐步上传到 Github，地址为 [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)。

## Contributor

1. Windstamp, [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)



作者：furnace
链接：https://www.jianshu.com/p/33572ccbef98
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。



# Hyperledger Fabric 专题 - 构建第一个 Hyperledger Fabric 网络

###### 注解

> 本文中的指令已经通过验证，可与提供的 tar 文件中的最新稳定 Docker 镜像和预编译的实用程序一起使用。如果使用当前主分支中的镜像或工具运行这些命令，则可能会看到配置和 panic 错误。

> Hyperledger Fabric 最新的文档基于版本是 v2.0 Alpha release。由于示例中相关的 docker 镜像的版本是 v1.4.3，因为相关文档需要参考的版本为 v1.4.3。这些文档链接在本文的 Reference 部分都有涉及。

构建你的第一个网络 (Building Your First Network, BYFN) 方案将提供一个示例 Hyperledger Fabric 网络，该网络由两个组织组成，每个组织维护两个对端节点。尽管可以使用其他交易排序服务实现，但默认情况下还是部署 Solo 交易排序服务。

## 1. 安装先决条件

在开始之前，如果你尚未这样做，则不妨检查一下是否已在要开发区块链应用程序和/或运行 Hyperledger Fabric 的平台上安装了所有 [必备软件](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Fprereqs.html)。

你还需要 [安装示例，二进制文件和 Docker 映像](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Finstall.html)。你会注意到，`fabric-samples` 仓库中包含许多示例。我们将使用 `first-network` 示例。现在打开该子目录。



```bash
cd fabric-samples/first-network
```

###### 注解

> 本文档中提供的命令必须从 `fabric-samples` 仓库克隆的 `first-network` 子目录中运行。如果选择从其他位置运行命令，则提供的各种脚本将无法找到二进制文件。

## 2. 立即运行

我们提供了一个带有完整注释的脚本 —  byfn.sh，该脚本利用这些 Docker 镜像来快速引导 Hyperledger Fabric 网络，该网络默认情况下由代表两个不同组织的四个对端节点和一个交易排序器节点组成。它还将启动一个容器来运行脚本化执行，该脚本化执行将对端节点加入到通道中，部署链码并根据已部署的链码推动交易的执行。

这是 byfn.sh 脚本的帮助文本：



```bash
Usage:
byfn.sh <mode> [-c <channel name>] [-t <timeout>] [-d <delay>] [-f <docker-compose-file>] [-s <dbtype>] [-l <language>] [-o <consensus-type>] [-i <imagetag>] [-v]"
  <mode> - one of 'up', 'down', 'restart', 'generate' or 'upgrade'"
    - 'up' - bring up the network with docker-compose up"
    - 'down' - clear the network with docker-compose down"
    - 'restart' - restart the network"
    - 'generate' - generate required certificates and genesis block"
    - 'upgrade'  - upgrade the network from version 1.3.x to 1.4.0"
  -c <channel name> - channel name to use (defaults to \"mychannel\")"
  -t <timeout> - CLI timeout duration in seconds (defaults to 10)"
  -d <delay> - delay duration in seconds (defaults to 3)"
  -f <docker-compose-file> - specify which docker-compose file use (defaults to docker-compose-cli.yaml)"
  -s <dbtype> - the database backend to use: goleveldb (default) or couchdb"
  -l <language> - the chaincode language: golang (default), node, or java"
  -o <consensus-type> - the consensus-type of the ordering service: solo (default), kafka, or etcdraft"
  -i <imagetag> - the tag to be used to launch the network (defaults to \"latest\")"
  -v - verbose mode"
byfn.sh -h (print this message)"

Typically, one would first generate the required certificates and
genesis block, then bring up the network. e.g.:"

  byfn.sh generate -c mychannel"
  byfn.sh up -c mychannel -s couchdb"
  byfn.sh up -c mychannel -s couchdb -i 1.4.0"
  byfn.sh up -l node"
  byfn.sh down -c mychannel"
  byfn.sh upgrade -c mychannel"

Taking all defaults:"
      byfn.sh generate"
      byfn.sh up"
      byfn.sh down"
```

如果选择不提供标志参数，则脚本将使用默认值。

### 2.1 生成网络组件

准备好尝试了吗？好吧！执行以下命令：



```ruby
$ cd /path/to/fabric-samples/first-network
$ ./byfn.sh generate
```

你将看到有关发生的情况的简短说明，以及是/否命令行提示符。响应 y 或按返回键以执行所描述的操作。



```php
Generating certs and genesis block for channel 'mychannel' with CLI timeout of '10' seconds and CLI delay of '3' seconds
Continue? [Y/n] y
proceeding ...
/Users/xxx/dev/fabric-samples/bin/cryptogen

##########################################################
##### Generate certificates using cryptogen tool #########
##########################################################
org1.example.com
2017-06-12 21:01:37.334 EDT [bccsp] GetDefault -> WARN 001 Before using BCCSP, please call InitFactories(). Falling back to bootBCCSP.
...

/Users/xxx/dev/fabric-samples/bin/configtxgen
##########################################################
#########  Generating Orderer Genesis block ##############
##########################################################
2017-06-12 21:01:37.558 EDT [common/configtx/tool] main -> INFO 001 Loading configuration
2017-06-12 21:01:37.562 EDT [msp] getMspConfig -> INFO 002 intermediate certs folder not found at [/Users/xxx/dev/byfn/crypto-config/ordererOrganizations/example.com/msp/intermediatecerts]. Skipping.: [stat /Users/xxx/dev/byfn/crypto-config/ordererOrganizations/example.com/msp/intermediatecerts: no such file or directory]
...
2017-06-12 21:01:37.588 EDT [common/configtx/tool] doOutputBlock -> INFO 00b Generating genesis block
2017-06-12 21:01:37.590 EDT [common/configtx/tool] doOutputBlock -> INFO 00c Writing genesis block

#################################################################
### Generating channel configuration transaction 'channel.tx' ###
#################################################################
2017-06-12 21:01:37.634 EDT [common/configtx/tool] main -> INFO 001 Loading configuration
2017-06-12 21:01:37.644 EDT [common/configtx/tool] doOutputChannelCreateTx -> INFO 002 Generating new channel configtx
2017-06-12 21:01:37.645 EDT [common/configtx/tool] doOutputChannelCreateTx -> INFO 003 Writing new channel tx

#################################################################
#######    Generating anchor peer update for Org1MSP   ##########
#################################################################
2017-06-12 21:01:37.674 EDT [common/configtx/tool] main -> INFO 001 Loading configuration
2017-06-12 21:01:37.678 EDT [common/configtx/tool] doOutputAnchorPeersUpdate -> INFO 002 Generating anchor peer update
2017-06-12 21:01:37.679 EDT [common/configtx/tool] doOutputAnchorPeersUpdate -> INFO 003 Writing anchor peer update

#################################################################
#######    Generating anchor peer update for Org2MSP   ##########
#################################################################
2017-06-12 21:01:37.700 EDT [common/configtx/tool] main -> INFO 001 Loading configuration
2017-06-12 21:01:37.704 EDT [common/configtx/tool] doOutputAnchorPeersUpdate -> INFO 002 Generating anchor peer update
2017-06-12 21:01:37.704 EDT [common/configtx/tool] doOutputAnchorPeersUpdate -> INFO 003 Writing anchor peer update
```

第一步将为我们的各种网络实体生成所有证书和密钥，用于交易排序服务的创世区块以及配置 [通道](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23channel) 所需的配置交易的集合。

### 2.2 建立网络

接下来，你可以使用以下命令之一启动网络：



```ruby
$ cd /path/to/fabric-samples/first-network
$ sudo ./byfn.sh up
```

上面的命令将编译 Golang 链码镜像并旋转相应的容器。 Go 是默认的链码语言，但是还支持 Node.js 和 Java 链码。如果你想使用 Node.js 链代码来完成本教程，请改用以下命令：



```bash
# we use the -l flag to specify the chaincode language
# forgoing the -l flag will default to Golang

./byfn.sh up -l node
```

###### 注解

> 有关 Node.js 填充程序的更多信息，请参阅其 [文档](https://links.jianshu.com/go?to=https%3A%2F%2Ffabric-shim.github.io%2Frelease-1.4%2Ffabric-shim.ChaincodeInterface.html%3Fredirect%3Dtrue)。

###### 注解

> 有关 Java 填充程序的更多信息，请参阅其 [文档](https://links.jianshu.com/go?to=https%3A%2F%2Ffabric-chaincode-java.github.io%2Forg%2Fhyperledger%2Ffabric%2Fshim%2FChaincode.html)。

使示例使用 Java 链码运行，你必须指定 -l java，如下所示：



```ruby
$ cd /path/to/fabric-samples/first-network
$ sudo ./byfn.sh up -l java
```

###### 注解

> 不要同时运行这些命令。除非你关闭并重新建立网络，否则只能尝试一种语言。

除了支持多种链码语言外，你还可以发出一个标志，该标志将显示一个五节点的 Raft 交易排序服务或 Kafka 交易排序服务，而不是一个节点的 Solo 交易排序服务。有关当前支持的交易排序服务实现的更多信息，请查看 [交易排序服务](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Forderer%2Fordering_service.html)。

要使用 Raft 交易排序服务启动网络，请执行：



```ruby
$ cd /path/to/fabric-samples/first-network
$ sudo ./byfn.sh up -o etcdraft
```

要使用 Kafka 交易排序服务建立网络，请执行：



```ruby
$ cd /path/to/fabric-samples/first-network
$ sudo ./byfn.sh up -o kafka
```

再次提示你是否要继续还是中止。回应 y 或按回车键：



```ruby
Starting for channel 'mychannel' with CLI timeout of '10' seconds and CLI delay of '3' seconds
Continue? [Y/n]
proceeding ...
Creating network "net_byfn" with the default driver
Creating peer0.org1.example.com
Creating peer1.org1.example.com
Creating peer0.org2.example.com
Creating orderer.example.com
Creating peer1.org2.example.com
Creating cli


 ____    _____      _      ____    _____
/ ___|  |_   _|    / \    |  _ \  |_   _|
\___ \    | |     / _ \   | |_) |   | |
 ___) |   | |    / ___ \  |  _ <    | |
|____/    |_|   /_/   \_\ |_| \_\   |_|

Channel name : mychannel
Creating channel...
```

日志将从此处继续。这将启动所有容器，然后驱动一个完整的端到端应用程序场景。成功完成后，它将在你的终端窗口中报告以下内容：



```ruby
Query Result: 90
2017-05-16 17:08:15.158 UTC [main] main -> INFO 008 Exiting.....
===================== Query successful on peer1.org2 on channel 'mychannel' =====================

===================== All GOOD, BYFN execution completed =====================


 _____   _   _   ____
| ____| | \ | | |  _ \
|  _|   |  \| | | | | |
| |___  | |\  | | |_| |
|_____| |_| \_| |____/
```

你可以滚动浏览这些日志以查看各种交易。如果未得到此结果，请跳至 [故障排除](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fbuild_network.html%23troubleshoot) 部分，让我们看看我们是否可以帮助你发现问题所在。

### 2.3 关闭网络

最后，让我们关闭所有组件，以便我们可以每次探索一次网络设置。以下内容将杀死你的容器，删除加密组件和四个组件，并从 Docker 注册表中删除链码镜像：



```ruby
$ cd /path/to/fabric-samples/first-network
$ sudo ./byfn.sh down
```

再次提示你继续，用 y 响应或按回车键：



```csharp
Stopping with channel 'mychannel' and CLI timeout of '10'
Continue? [Y/n] y
proceeding ...
WARNING: The CHANNEL_NAME variable is not set. Defaulting to a blank string.
WARNING: The TIMEOUT variable is not set. Defaulting to a blank string.
Removing network net_byfn
468aaa6201ed
...
Untagged: dev-peer1.org2.example.com-mycc-1.0:latest
Deleted: sha256:ed3230614e64e1c83e510c0c282e982d2b06d148b1c498bbdcc429e2b2531e91
...
```

如果你想进一步了解基础工具和引导机制，请继续阅读。在接下来的这些部分中，我们将逐步介绍构建一个功能齐全的 Hyperledger Fabric 网络的各个步骤和要求。

###### 注解

> 下面概述的手动步骤假定 cli 容器中的 FABRIC_LOGGING_SPEC 设置为 DEBUG。你可以通过修改 `first-network` 目录中的 `docker-compose-cli.yaml` 文件来进行设置。例如



```php
cli:
  container_name: cli
  image: hyperledger/fabric-tools:$IMAGE_TAG
  tty: true
  stdin_open: true
  environment:
    - GOPATH=/opt/gopath
    - CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock
    - FABRIC_LOGGING_SPEC=DEBUG
    #- FABRIC_LOGGING_SPEC=INFO
```

## 3. 密钥生成器

我们将使用 `cryptogen` 工具为我们的各种网络实体生成密钥材料 (x509 证书和签名密钥)。这些证书代表身份，它们允许在我们的实体进行通信和交易时进行签名/验证身份授权。

### 3.1 它是如何工作的？

Cryptogen 使用了一个文件 - `crypto-config.yaml`，该文件包含网络拓扑，并允许我们为组织以及属于这些组织的组件生成一组证书和密钥。每个组织都有一个唯一的根证书 (`ca-cert`)，它将特定组件 (对端节点和交易排序器) 绑定到该组织。通过为每个组织分配唯一的 CA 证书，我们正在模仿一个典型的网络，参与的 [成员](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23member) 将使用其自己的证书颁发机构。Hyperledger Fabric 中的交易和通信由实体的私钥 (密钥库) 签名，然后通过公钥 (signcerts) 进行验证。

你会在此文件中注意到一个 `count` 变量。我们使用它来指定每个组织的对端节点数；在我们的案例中，每个单位有两个对端节点。现在，我们不会深入研究 [x.509 证书和公钥基础结构](https://links.jianshu.com/go?to=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FPublic_key_infrastructure) 的细节。如果您有兴趣，可以自己阅读这些主题。

在运行 `cryptogen` 工具之后，生成的证书和密钥将保存到名为 `crypto-config` 的文件夹中。请注意，`crypto-config.yaml` 文件列出了五个与交易排序器组织相关的交易排序器。尽管 `cryptogen` 工具将为所有这五个交易排序器创建证书，除非使用了 `Raft` 或 `Kafka` 交易排序服务，否则这些交易排序器中只有一个将用于 Solo 交易排序服务实现中，并用于创建系统通道和 `mychannel`。

## 4. 配置交易生成器

`configtxgen` 工具用于创建四个配置组件：

- 交易排序器的创始区块 (`genesis block`)，
- 通道配置交易 (`configuration transaction`)，
- 和两个锚点对端节点交易 (`anchor peer transactions`) - 每个对端节点组织一个。

有关此工具功能的完整说明，请参见 `configtxgen`。

交易排序器区块是交易排序服务的 [创世区块](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23genesis-block)，并且在 [通道](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23channel) 创建时将通道 配置交易文件广播到交易排序器。顾名思义，锚定对端节点交易 (anchor peer transaction) 指定此通道上每个组织的 [锚定对端节点](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23anchor-peer)。

### 4.1 它是如何工作的？

Configtxgen 使用一个文件 `configtx.yaml`，其中包含示例网络的定义。有三个成员 - 一个交易排序器组织 (OrdererOrg) 和两个对端节点组织 (Org1 ＆ Org2)，每个组织都管理和维护两个对端节点。该文件还指定了一个联盟 - `SampleConsortium` - 由我们的两个对端节点组织组成。请特别注意此文件底部的 "Profiles" 部分。你会注意到我们有几个唯一的 profiles。一些值得注意的地方：

- `TwoOrgsOrdererGenesis`：为 Solo 交易排序服务生成创世区块。
- `SampleMultiNodeEtcdRaft`：生成 Raft 交易排序服务的创世区块。仅在发出 `-o` 标志并指定 `etcdraft` 时使用。
- `SampleDevModeKafka`：生成 Kafka 交易排序服务的创世区块。仅在发出 `-o` 标志并指定 `kafka` 时使用。
- `TwoOrgsChannel`：为我们的通道 `mychannel` 生成创世区块。

这些标题很重要，因为在创建组件时，我们会将它们作为参数传递。

###### 注解

> 请注意，我们的 `SampleConsortium` 是在系统级配置文件中定义的，然后由我们的通道级配置文件引用。通道存在于联盟的权限范围内，所有联盟都必须在整个网络范围内定义。

该文件还包含两个值得注意的附加规范。首先，我们为每个对端节点组织指定锚定对端节点 (peer0.org1.example.com & peer0.org2.example.com)。其次，我们指向每个组织 MSP 的目录位置，从而允许我们将每个组织的根证书存储在交易排序器的创始区块中。这是一个关键的概念。现在，与交易排序服务通信的任何网络实体都可以验证其数字签名。

## 5. 运行上述工具

你可以使用 `configtxgen` 和 `cryptogen` 命令手动生成证书/密钥和各种配置工件。或者，你可以尝试修改 `byfn.sh` 脚本以实现目标。

### 5.1 手动生成组件

你可以参考 `byfn.sh` 脚本中的 `generateCerts` 函数，以获取生成将用于 `crypto-config.yaml` 文件中定义的网络配置的证书所必需的命令。但是，为方便起见，我们还将在此处提供参考。

首先，让我们运行 `cryptogen` 工具。我们的二进制文件位于 `bin` 目录中，因此我们需要提供工具所在的相对路径。



```ruby
$ ../bin/cryptogen generate --config=./crypto-config.yaml
```

你应该在终端中看到以下内容：



```css
org1.example.com
org2.example.com
```

证书和密钥 (即 MSP 材料) 将输出到 `first-network` 目录的子目录 - `crypto-config`。

接下来，我们需要告诉 `configtxgen` 工具在哪里寻找需要读取的 `configtx.yaml` 文件。我们将在当前的工作目录中告诉它：



```bash
$ export FABRIC_CFG_PATH=$PWD
```

然后，我们将调用 `configtxgen` 工具来创建交易排序器的创始区块：



```ruby
$ ../bin/configtxgen -profile TwoOrgsOrdererGenesis -channelID byfn-sys-channel -outputBlock ./channel-artifacts/genesis.block
```

要输出 Raft 交易排序服务的创始区块，此命令应为：



```ruby
$ ../bin/configtxgen -profile SampleMultiNodeEtcdRaft -channelID byfn-sys-channel -outputBlock ./channel-artifacts/genesis.block
```

请注意此处使用的 `SampleMultiNodeEtcdRaft` profile。

要输出 `Kafka` 交易排序服务的创世区块，请执行：



```ruby
$ ../bin/configtxgen -profile SampleDevModeKafka -channelID byfn-sys-channel -outputBlock ./channel-artifacts/genesis.block
```

如果你未使用 `Raft` 或 `Kafka`，则应看到类似于以下内容的输出：



```css
2017-10-26 19:21:56.301 EDT [common/tools/configtxgen] main -> INFO 001 Loading configuration
2017-10-26 19:21:56.309 EDT [common/tools/configtxgen] doOutputBlock -> INFO 002 Generating genesis block
2017-10-26 19:21:56.309 EDT [common/tools/configtxgen] doOutputBlock -> INFO 003 Writing genesis block
```

###### 注解

> 交易排序器的创世区块和我们将要创建的后续组件将输出到该项目的 `channel-artifacts` 子目录中。上面命令中的 channelID 是系统通道的名称。

### 5.2 创建通道配置交易

接下来，我们需要创建通道交易组件。确保替换 `$CHANNEL_NAME` 或将 `CHANNEL_NAME` 设置为可在以下说明中使用的环境变量：



```bash
# The channel.tx artifact contains the definitions for our sample channel

$ export CHANNEL_NAME=mychannel  && ../bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME
```

请注意，如果你使用的是 `Raft` 或 `Kafka` 交易排序服务，则不必对通道发出特殊命令。在创建网络的创世纪区块时，`TwoOrgsChannel` 配置文件将使用你指定的交易排序服务配置。

如果你不使用 `Raft` 或 `Kafka` 交易排序服务，则应在终端中看到类似于以下内容的输出：



```css
2017-10-26 19:24:05.324 EDT [common/tools/configtxgen] main -> INFO 001 Loading configuration
2017-10-26 19:24:05.329 EDT [common/tools/configtxgen] doOutputChannelCreateTx -> INFO 002 Generating new channel configtx
2017-10-26 19:24:05.329 EDT [common/tools/configtxgen] doOutputChannelCreateTx -> INFO 003 Writing new channel tx
```

接下来，我们将在正在构建的通道上为 `Org1` 定义锚点。同样，请确保替换 `$CHANNEL_NAME` 或为以下命令设置环境变量。终端输出将模拟通道交易组件的输出：



```ruby
$ ../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP
```

现在，我们将在同一通道上为 `Org2` 定义锚点：



```ruby
$ ../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org2MSP
```

## 6. 启动网络

###### 注意

> 如果你以前运行过 `byfn.sh` 示例，请确保在继续操作之前已关闭测试网络 (请参阅 [关闭网络](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fbuild_network.html%23bring-down-the-network))。



```ruby
$ sudo ./byfn.sh down
```

我们将利用脚本来启动我们的网络。 docker-compose 文件引用了我们先前下载的镜像，并使用我们先前生成的 `genesis.block` 引导交易排序器程序。

我们希望手动检查命令，以显示每个调用的语法和功能。

首先，我们开始我们的网络：



```ruby
$ sudo docker-compose -f docker-compose-cli.yaml up -d
```

如果要查看网络的实时日志，请不要提供 -d 标志。如果让日志流传输，则需要打开第二个终端以执行 CLI 调用。

### 6.1 创建 & 加入通道

回想一下，我们在上面的“创建通道配置交易”部分中使用 `configtxgen` 工具创建了通道配置交易。你可以重复该过程，以使用传递给 `configtxgen` 工具的 `configtx.yaml` 中的相同或不同配置文件来创建其他通道配置交易。然后，你可以重复本节中定义的过程以在网络中建立其他通道。

我们将使用 `docker exec` 命令进入 CLI 容器：



```bash
$ sudo docker exec -it cli bash
```

如果成功，你应该看到以下内容：



```ruby
root@33ab5acc5622:/opt/gopath/src/github.com/hyperledger/fabric/peer#
```

为了使以下 CLI 命令起作用，我们需要在命令前添加以下四个环境变量。`peer0.org1.example.com` 的这些变量被内嵌到 CLI 容器中，因此我们可以在不传递它们的情况下进行操作。但是，如果要将调用发送给其他对端节点或交易排序器，请在进行任何 CLI 调用时覆盖环境变量，如以下示例所示：



```ruby
# Environment variables for PEER0

$ CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
$ CORE_PEER_ADDRESS=peer0.org1.example.com:7051
$ CORE_PEER_LOCALMSPID="Org1MSP"
$ CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
```

接下来，作为创建通道请求的一部分，我们将把在 [创建通道配置交易](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fbuild_network.html%23createchanneltx) 部分 (我们称为 `channel.tx`) 中创建生成的通道配置交易组件传递给交易排序器。

我们使用 `-c` 标志指定我们的通道名称，并使用 `-f` 标志指定我们的通道配置交易。在这种情况下，它是 `channel.tx`，但是你可以使用其他名称挂载自己的配置交易。再次，我们将在 CLI 容器中设置 `CHANNEL_NAME` 环境变量，以便我们不必显式传递此参数。通道名称必须全部为小写字母，少于 250 个字符，并且与正则表达式 `[a-z][a-z0-9 .-] *` 相匹配。



```ruby
$ export CHANNEL_NAME=mychannel

# the channel.tx file is mounted in the channel-artifacts directory within your CLI container
# as a result, we pass the full path for the file
# we also pass the path for the orderer ca-cert in order to verify the TLS handshake
# be sure to export or replace the $CHANNEL_NAME variable appropriately

$ peer channel create -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```

###### 注解

> 请注意我们在此命令中传递的 `--cafile`。这是交易排序器根证书的本地路径，使我们可以验证 TLS 握手。

此命令返回一个创世区块 - `` - 我们将使用它来加入通道。它包含 `channel.tx` 中指定的配置信息。如果你未对默认通道名称进行任何修改，则该命令将返回一个名为 `mychannel.block` 的原型。

###### 注意

> 这些手动命令的其余部分将保留在 CLI 容器中。当定位到除 `peer0.org1.example.com` 以外的对端节点对象时，你还必须记住在所有命令前加上相应的环境变量。

现在，将 `peer0.org1.example.com` 加入该通道。



```ruby
# By default, this joins ``peer0.org1.example.com`` only
# the <CHANNEL_NAME.block> was returned by the previous command
# if you have not modified the channel name, you will join with mychannel.block
# if you have created a different channel name, then pass in the appropriately named block

$ peer channel join -b mychannel.block
```

你可以通过对在上面的 [创建和加入通道](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fbuild_network.html%23peerenvvars) 部分中使用的四个环境变量进行适当的更改，使其他对端节点加入通道。

与其加入每个对端节点，不如简单地加入 `peer0.org2.example.com`，以便我们可以正确更新通道中的锚点对端节点定义。由于我们将覆盖内嵌到 CLI 容器中的默认环境变量，因此该完整命令如下：



```ruby
$ CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
$ CORE_PEER_ADDRESS=peer0.org2.example.com:9051
$ CORE_PEER_LOCALMSPID="Org2MSP"
$ CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
$ peer channel join -b mychannel.block
```

###### 注解

> 在 v1.4.1 之前，docker 网络中的所有对端节点都使用端口 `7051`。如果使用 v1.4.1 之前的版本的 fabric-samples，请在本教程中将所有出现的 `CORE_PEER_ADDRESS` 修改为使用端口 `7051`。

另外，你可以选择单独设置这些环境变量，而不是传入整个字符串。设置完后，你只需再次发出 `peer channel join` 命令，CLI 容器将代表 `peer0.org2.example.com`。

### 6.2 更新锚定对端节点

以下命令是通道更新，它们将传播到通道的定义。从本质上讲，我们在通道的创始区块上方添加了其他配置信息。请注意，我们不是在修改创世区块，而只是将增量添加到将定义锚点对端节点的链中。

更新通道定义以将 Org1 的锚点对端节点定义为 `peer0.org1.example.com`：



```ruby
$ CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
$ CORE_PEER_ADDRESS=peer0.org1.example.com:7051
$ CORE_PEER_LOCALMSPID="Org1MSP"
$ CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
$ peer channel update -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/Org1MSPanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```

现在更新通道定义，以将 Org2 的锚点对端节点定义为 `peer0.org2.example.com`。与 Org2 对端节点的 `peer channel join` 命令相同，我们需要在此调用前加上适当的环境变量。



```ruby
$ CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
$ CORE_PEER_ADDRESS=peer0.org2.example.com:9051
$ CORE_PEER_LOCALMSPID="Org2MSP"
$ CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
$ peer channel update -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/Org2MSPanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```

### 6.3 安装并实例化链码

###### 注解

> 我们将利用一个简单的现有链码。要了解如何编写自己的链码，请参阅 [链码开发者](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fchaincode4ade.html) 教程。

应用程序通过链码与区块链账本进行交互。因此，我们需要在将执行并背书我们交易的每个对端节点上安装链码，然后在通道上实例化链码。

首先，将示例 Go，Node.js 或 Java 链码安装到 Org1 中的 peer0 节点上。这些命令将指定的源代码样式放置到我们对端节点的文件系统上。

###### 注解

> 每个链码名称和版本只能安装一个版本的源代码。源代码在链码名称和版本的上下文中存在于对端节点的文件系统中，它与语言无关。同样，实例化的链码容器将反映对端节点上已安装的任何语言。

首先设置 Org1 相关的环境变量。



```ruby
$ CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
$ CORE_PEER_ADDRESS=peer0.org1.example.com:7051
$ CORE_PEER_LOCALMSPID="Org1MSP"
$ CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
```

###### Golang



```ruby
# this installs the Go chaincode. For go chaincode -p takes the relative path from $GOPATH/src
$ peer chaincode install -n mycc -v 1.0 -p github.com/chaincode/chaincode_example02/go/
```

###### Node.js



```ruby
# this installs the Node.js chaincode
# make note of the -l flag to indicate "node" chaincode
# for node chaincode -p takes the absolute path to the node.js chaincode
$ peer chaincode install -n mycc -v 1.0 -l node -p /opt/gopath/src/github.com/chaincode/chaincode_example02/node/
```

###### Java



```ruby
# make note of the -l flag to indicate "java" chaincode
# for java chaincode -p takes the absolute path to the java chaincode
$ peer chaincode install -n mycc -v 1.0 -l java -p /opt/gopath/src/github.com/chaincode/chaincode_example02/java/
```

当我们实例化通道上的链码时，将设置背书策略以要求来自 Org1 和 Org2 中的对端节点的背书。因此，我们还需要在 Org2 的对端节点上安装链码。

修改以下四个环境变量以对 Org2 中的 peer0 发出安装命令：



```ruby
# Environment variables for PEER0 in Org2

$ CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
$ CORE_PEER_ADDRESS=peer0.org2.example.com:9051
$ CORE_PEER_LOCALMSPID="Org2MSP"
$ CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
```

现在，将示例 Go，Node.js 或 Java 链码安装到 Org2 中的 peer0 上。这些命令将指定的源代码样式放置到我们对端节点的文件系统上。



```ruby
# this installs the Go chaincode. For go chaincode -p takes the relative path from $GOPATH/src
$ peer chaincode install -n mycc -v 1.0 -p github.com/chaincode/chaincode_example02/go/
```

###### Node.js



```ruby
# this installs the Node.js chaincode
# make note of the -l flag to indicate "node" chaincode
# for node chaincode -p takes the absolute path to the node.js chaincode
$ peer chaincode install -n mycc -v 1.0 -l node -p /opt/gopath/src/github.com/chaincode/chaincode_example02/node/
```

###### Java



```ruby
# make note of the -l flag to indicate "java" chaincode
# for java chaincode -p takes the absolute path to the java chaincode
$ peer chaincode install -n mycc -v 1.0 -l java -p /opt/gopath/src/github.com/chaincode/chaincode_example02/java/
```

接下来，实例化通道上的链码。这将初始化通道上的链码，为链码设置背书策略，并为目标对端节点启动链码容器。注意 `-P` 参数。这是我们的策略，其中我们针对要验证的链码指定了交易所需的背书级别。

在下面的命令中，你会注意到我们将策略指定为 `-P "AND ('Org1MSP.peer','Org2MSP.peer')"`。这意味着我们需要来自 Org1 和 Org2 的对端节点的背书 (即两次背书)。如果我们将语法更改为 OR，则只需要一个背书即可。

###### Golang



```ruby
# be sure to replace the $CHANNEL_NAME environment variable if you have not exported it
# if you did not install your chaincode with a name of mycc, then modify that argument as well
$ export CHANNEL_NAME=mychannel
$ peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc -v 1.0 -c '{"Args":["init","a", "100", "b","200"]}' -P "AND ('Org1MSP.peer','Org2MSP.peer')"
```

###### Node.js

###### 注解

> Node.js 链码的实例化大约需要一分钟。该命令未挂起；而是在编译镜像时安装了 fabric-shim 层。



```ruby
# be sure to replace the $CHANNEL_NAME environment variable if you have not exported it
# if you did not install your chaincode with a name of mycc, then modify that argument as well
# notice that we must pass the -l flag after the chaincode name to identify the language

$ peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc -l node -v 1.0 -c '{"Args":["init","a", "100", "b","200"]}' -P "AND ('Org1MSP.peer','Org2MSP.peer')"
```

###### Java

###### 注解

> 请注意，Java 链码实例化可能会花费一些时间，因为它会在 Java 环境中编译链码并下载 docker 容器。



```swift
$ peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc -l java -v 1.0 -c '{"Args":["init","a", "100", "b","200"]}' -P "AND ('Org1MSP.peer','Org2MSP.peer')"
```

有关策略实现的更多详细信息，请参阅 [背书策略](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Fendorsement-policies.html) 文档。

如果你希望其他对端节点与帐本进行交互，则需要将它们加入通道，并将链码源的相同名称，版本和语言安装到适当的对端节点的文件系统上。一旦每个对端节点尝试与该特定链码进行交互，就会为每个对端节点启动一个链码容器。同样，请注意 Node.js 镜像的编译速度较慢。

一旦链码已在通道上实例化，我们就可以放弃 `l` 标志。我们只需要输入通道标识符和链码的名称即可。

### 6.4 Query

让我们查询 `a` 的值，以确保链码已正确实例化并填充了状态数据库。查询的语法如下：



```ruby
# be sure to set the -C and -n flags appropriately

$ peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["query","a"]}'
100
```

### 6.5 Invoke

现在，将 `10` 从 `a` 移到 `b`。此交易将创建一个新区块并更新状态数据库。调用的语法如下：



```swift
# be sure to set the -C and -n flags appropriately

$ peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"Args":["invoke","a","b","10"]}'
```

### 6.6 Query

让我们确认之前的调用已正确执行。我们将键 `a` 的值初始化为 `100`，并在之前的调用中删除了 `10`。因此，针对 `a` 的查询应返回 `90`。查询的语法如下。



```ruby
# be sure to set the -C and -n flags appropriately

$ peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["query","a"]}'
```

我们应该看到以下内容：



```rust
Query Result: 90
```

随意重新开始并操作键值对和后续调用。

### 6.7 Install

现在，我们将链码安装在第三个对端节点 Org2 中的 peer1 上。修改以下四个环境变量以对 Org2 中的 peer1 发出安装命令：



```ruby
# Environment variables for PEER1 in Org2

$ CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
$ CORE_PEER_ADDRESS=peer1.org2.example.com:10051
$ CORE_PEER_LOCALMSPID="Org2MSP"
$ CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/tls/ca.crt
```

现在，将示例 Go，Node.js 或 Java 链码安装到 Org2 中的 peer1 上。这些命令将指定的源代码样式放置到我们对端节点的文件系统上。

###### Golang



```ruby
# this installs the Go chaincode. For go chaincode -p takes the relative path from $GOPATH/src
$ peer chaincode install -n mycc -v 1.0 -p github.com/chaincode/chaincode_example02/go/
```

###### Node.js



```ruby
# this installs the Node.js chaincode
# make note of the -l flag to indicate "node" chaincode
# for node chaincode -p takes the absolute path to the node.js chaincode
$ peer chaincode install -n mycc -v 1.0 -l node -p /opt/gopath/src/github.com/chaincode/chaincode_example02/node/
```

###### Java



```ruby
# make note of the -l flag to indicate "java" chaincode
# for java chaincode -p takes the absolute path to the java chaincode
$ peer chaincode install -n mycc -v 1.0 -l java -p /opt/gopath/src/github.com/chaincode/chaincode_example02/java/
```

### 6.8 Query

我们确认可以将查询发布到 Org2 中的 Peer1。我们将键 `a` 的值初始化为 `100`，并在之前的调用中删除了 `10`。因此，针对 `a` 的查询仍应返回 `90`。

Org2 中的 peer1 必须首先加入通道，然后它才能响应查询。可以通过执行以下命令来加入通道：



```ruby
$ CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
$ CORE_PEER_ADDRESS=peer1.org2.example.com:10051
$ CORE_PEER_LOCALMSPID="Org2MSP"
$ CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/tls/ca.crt
$ peer channel join -b mychannel.block
```

连接命令返回后，可以发出查询。查询的语法如下。



```ruby
# be sure to set the -C and -n flags appropriately

$ peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["query","a"]}'
```

我们应该看到以下内容：



```rust
Query Result: 90
```

随意重新开始并操作键值对和后续调用。

### 6.9 幕后发生了什么？

###### 注解

> 这些步骤描述了由 `./byfn.sh up` 运行 `script.sh` 的情况。使用 `./byfn.sh down` 清理网络，并确保此命令处于活动状态。然后使用相同的 `docker-compose` 提示符再次启动网络。

- 脚本 - `script.sh` - 内嵌在 CLI 容器中。该脚本根据提供的通道名称驱动 `createChannel` 命令，并使用 `channel.tx` 文件进行通道配置。
- `createChannel` 的输出是一个创世区块 - ` .block` - 它存储在对端节点的文件系统中，并包含从 `channel.tx` 指定的通道配置。
- 所有四个对端节点均执行 `joinChannel` 命令，该命令将先前生成的创世区块作为输入。此命令指示对端节点加入 `` 并创建以 `.block` 开头的链。
- 现在，我们有一个由四个对端节点和两个组织组成的通道。这是我们的 `TwoOrgsChannel` profile。
- `peer0.org1.example.com` 和 `peer1.org1.example.com` 属于 Org1。`peer0.org2.example.com` 和 `peer1.org2.example.com` 属于 Org2。
- 这些关系通过 `crypto-config.yaml` 定义，并且 MSP 路径在我们的 docker compose 中指定。
- 然后，更新 Org1MSP （peer0.org1.example.com) 和 Org2MSP (peer0.org2.example.com) 的锚点对端节点。为此，我们将 `Org1MSPanchors.tx` 和 `Org2MSPanchors.tx` 组件与通道名称一起传递给交易排序服务。
- 链码 - `chaincode_example02` - 安装在 `peer0.org1.example.com` 和 `peer0.org2.example.com` 上。
- 然后在 `mychannel` 上实例化链码。实例化将链码添加到通道，启动目标对端节点的容器，并初始化与链码关联的键值对。此示例的初始值为 ["a","100" "b","200"]。此实例化导致启动以 `dev-peer0.org2.example.com-mycc-1.0` 开头的容器。
- 实例化还为背书策略传递了一个参数。该策略定义为 `-P "AND ('Org1MSP.peer','Org2MSP.peer')"`，这意味着任何交易都必须由与 Org1 和 Org2 绑定的对端节点认可。
- 针对 "a" 的值查询将发布到 `peer0.org2.example.com`。实例化链码时，启动了一个名为 `dev-peer0.org2.example.com-mycc-1.0` 的 Org2 peer0 的容器。返回查询结果。没有发生写操作，因此对 "a" 的查询仍将返回值 "100"。
- 调用被发送到 `peer0.org1.example.com` 和 `peer0.org2.example.com`，以将 "10" 从 "a" 移动到 "b"。
- 查询发送到 `peer0.org2.example.com`，以获取 "a" 的值。返回值 90，正确反映了先前的交易，在该交易中，键 "a" 的值被修改了 10。
- 链码 - `chaincode_example02` - 安装在 `peer1.org2.example.com` 上。
- 查询发送到 `peer1.org2.example.com` 以获取 "a" 的值。这将启动名为 `dev-peer1.org2.example.com-mycc-1.0` 的第三个链码容器。返回值 90，正确反映了先前的交易，在该交易中，键 "a" 的值被修改了 10。

### 6.10 这说明了什么？

必须在对端节点上安装链码，以使其能够对帐本成功执行读/写操作。此外，直到针对该链码执行了初始化或传统交易 (读/写) (例如查询 "a" 的值) 后，对端节点的链码容器才启动。交易导致容器启动。而且，通道中的所有对端节点都维护账本的精确副本，该副本包括将区块，不可变的顺序记录存储在区块中的区块链，以及维护当前状态快照的状态数据库。这包括那些未安装链码的对端节点 (如上例中的 peer1.org1.example.com`)。最终，由于安装了链码，因此可以对其进行访问 (例如上例中的 peer1.org2.example.com)，因为该链码已被实例化。

### 6.11 如何查看这些交易？

检查 CLI Docker 容器的日志。



```ruby
$ sudo docker logs -f cli
```

你应该看到以下输出：



```ruby
2017-05-16 17:08:01.366 UTC [msp] GetLocalMSP -> DEBU 004 Returning existing local MSP
2017-05-16 17:08:01.366 UTC [msp] GetDefaultSigningIdentity -> DEBU 005 Obtaining default signing identity
2017-05-16 17:08:01.366 UTC [msp/identity] Sign -> DEBU 006 Sign: plaintext: 0AB1070A6708031A0C08F1E3ECC80510...6D7963631A0A0A0571756572790A0161
2017-05-16 17:08:01.367 UTC [msp/identity] Sign -> DEBU 007 Sign: digest: E61DB37F4E8B0D32C9FE10E3936BA9B8CD278FAA1F3320B08712164248285C54
Query Result: 90
2017-05-16 17:08:15.158 UTC [main] main -> INFO 008 Exiting.....
===================== Query successful on peer1.org2 on channel 'mychannel' =====================

===================== All GOOD, BYFN execution completed =====================


 _____   _   _   ____
| ____| | \ | | |  _ \
|  _|   |  \| | | | | |
| |___  | |\  | | |_| |
|_____| |_| \_| |____/
```

你可以滚动浏览这些日志以查看各种交易。

### 6.12 如何查看链码日志？

检查各个链码容器，以查看针对每个容器执行的单独交易。这是每个容器的合并输出：



```ruby
$ sudo docker logs dev-peer0.org2.example.com-mycc-1.0
ex02 Init
Aval = 100, Bval = 200
ex02 Invoke
Query Response:{"Name":"a","Amount":"100"}
ex02 Invoke
Aval = 90, Bval = 210
ex02 Invoke
Query Response:{"Name":"a","Amount":"90"}

$ sudo docker logs dev-peer0.org1.example.com-mycc-1.0
ex02 Invoke
Aval = 90, Bval = 210

$ sudo docker logs dev-peer1.org2.example.com-mycc-1.0
ex02 Invoke
Query Response:{"Name":"a","Amount":"90"}
```

## 7. 了解 Docker Compose 拓扑

BYFN 示例向我们提供了两种 Docker Compose 文件，这两种文件都是从 `docker-compose-base.yaml` (位于 `base` 文件夹中) 扩展的。我们的第一个版本 `docker-compose-cli.yaml` 为我们提供了一个 CLI 容器以及一个交易排序器和四个对端节点。我们将此文件用于此页面上的全部说明。

###### 注意

> 本节的其余部分介绍了为 SDK 设计的 docker-compose 文件。有关运行这些测试的详细信息，请参考 [Node SDK](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fhyperledger%2Ffabric-sdk-node) 仓库。

第二种形式 `docker-compose-e2e.yaml` 被构造为使用 Node.js SDK 运行端到端测试。除了可以使用 SDK 之外，它的主要区别还在于，fabric-ca 服务器具有容器。因此，我们能够将 REST 调用发送到组织 CA，以进行用户注册和登记。

如果要在不首先运行 `byfn.sh` 脚本的情况下使用 `docker-compose-e2e.yaml`，则我们将需要进行四个小修改。我们需要指向组织的 CA 的私钥。你可以在 `crypto-config` 文件夹中找到这些值。例如，要找到 Org1 的私钥，我们将遵循以下路径 - `crypto-config/peerOrganizations/org1.example.com/ca/`。私钥是一个长哈希值，后跟 `_sk`。 Org2 的路径为 - `crypto-config/peerOrganizations/org2.example.com/ca/`。

在 `docker-compose-e2e.yaml` 中，为 ca0 和 ca1 更新 FABRIC_CA_SERVER_TLS_KEYFILE 变量。你还需要编辑命令中提供的路径以启动 ca 服务器。你为每个 CA 容器提供相同的私钥两次。

## 8. 使用 CouchDB

可以将状态数据库从默认 (goleveldb) 切换到 CouchDB。CouchDB 可以使用相同的链码功能，但是，还具有对建模为 JSON 数据 的链码状态数据库执行丰富和复杂查询的功能。

要使用 CouchDB 代替默认数据库 (goleveldb)，请遵循前面概述的用于生成工件的相同过程，除了在启动网络时还要通过 docker-compose-couch.yaml：



```ruby
$ sudo docker-compose -f docker-compose-cli.yaml -f docker-compose-couch.yaml up -d
```

chaincode_example02 现在应该可以在底层使用 CouchDB 了。

###### 注解

> 如果你选择实现 fabric-couchdb 容器端口到主机端口的映射，请确保你了解安全隐患。在开发环境中端口的映射使 CouchDB REST API 可用，并允许通过 CouchDB Web 界面 (Fauxton) 可视化数据库。生产环境可能会避免实施端口映射，以限制外部对 CouchDB 容器的访问。

你可以按照上述步骤针对 chaincode_example02 链码使用 CouchDB 状态数据库，但是，为了行使 CouchDB 查询功能，你将需要使用数据编码为 JSON 的链码 (例如 marbles02)。你可以在 `fabric/examples/chaincode/go` 目录中找到 marbles02 链码。

我们将按照上述创建和加入通道一节中所述的相同过程来创建和加入通道。将同伴加入通道后，请执行以下步骤与 marbles02 链码进行交互：

- 在 `peer0.org1.example.com` 上安装并实例化链码：



```ruby
# be sure to modify the $CHANNEL_NAME variable accordingly for the instantiate command

$ peer chaincode install -n marbles -v 1.0 -p github.com/chaincode/marbles02/go
$ peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n marbles -v 1.0 -c '{"Args":["init"]}' -P "OR ('Org1MSP.peer','Org2MSP.peer')"
```

创建一些大理石并将其移动：



```swift
# be sure to modify the $CHANNEL_NAME variable accordingly

$ peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n marbles -c '{"Args":["initMarble","marble1","blue","35","tom"]}'
$ peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n marbles -c '{"Args":["initMarble","marble2","red","50","tom"]}'
$ peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n marbles -c '{"Args":["initMarble","marble3","blue","70","tom"]}'
$ peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n marbles -c '{"Args":["transferMarble","marble2","jerry"]}'
$ peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n marbles -c '{"Args":["transferMarblesBasedOnColor","blue","jerry"]}'
$ peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n marbles -c '{"Args":["delete","marble1"]}'
```

如果选择在 docker-compose 中映射 CouchDB 端口，则现在可以通过打开浏览器并导航至以下 URL，通过 CouchDB Web 界面 (Fauxton) 查看状态数据库：

[http://localhost:5984/_utils](https://links.jianshu.com/go?to=http%3A%2F%2Flocalhost%3A5984%2F_utils)

你应该会看到一个名为 `mychannel` (或你唯一的通道名称) 的数据库以及其中的文档。

###### 注意

> 对于以下命令，请确保适当地更新 $CHANNEL_NAME 变量。

你可以通过 CLI 运行常规查询 (例如，读取 `marble2`)：



```swift
$ peer chaincode query -C $CHANNEL_NAME -n marbles -c '{"Args":["readMarble","marble2"]}'
```

输出应显示 `marble2` 的详细信息：



```rust
Query Result: {"color":"red","docType":"marble","name":"marble2","owner":"jerry","size":50}
```

你可以检索特定大理石的历史记录 - 例如 marble1：



```swift
$ peer chaincode query -C $CHANNEL_NAME -n marbles -c '{"Args":["getHistoryForMarble","marble1"]}'
```

输出应显示在 marble1 上的交易：



```rust
Query Result: [{"TxId":"1c3d3caf124c89f91a4c0f353723ac736c58155325f02890adebaa15e16e6464", "Value":{"docType":"marble","name":"marble1","color":"blue","size":35,"owner":"tom"}},{"TxId":"755d55c281889eaeebf405586f9e25d71d36eb3d35420af833a20a2f53a3eefd", "Value":{"docType":"marble","name":"marble1","color":"blue","size":35,"owner":"jerry"}},{"TxId":"819451032d813dde6247f85e56a89262555e04f14788ee33e28b232eef36d98f", "Value":}]
```

你还可以对数据内容执行丰富的查询，例如通过所有者 jerry 查询大理石字段：



```swift
$ peer chaincode query -C $CHANNEL_NAME -n marbles -c '{"Args":["queryMarblesByOwner","jerry"]}'
```

输出应显示 `jerry` 拥有的两种大理石：



```rust
Query Result: [{"Key":"marble2", "Record":{"color":"red","docType":"marble","name":"marble2","owner":"jerry","size":50}},{"Key":"marble3", "Record":{"color":"blue","docType":"marble","name":"marble3","owner":"jerry","size":70}}]
```

## 9. 为什么选择 CouchDB

CouchDB 是一种 NoSQL 解决方案。它是一个面向文档的数据库，其中文档字段存储为键值映射。字段可以是简单的键值对，列表或映射。除了 LevelDB 支持的键/复合键/键范围查询外，CouchDB 还支持完整的数据丰富查询功能，例如针对整个区块链数据的非键查询，因为其数据内容以 JSON 格式存储，并且完全可查询的。因此，CouchDB 可以满足 LevelDB 不支持的许多用例的链码，审计，报告要求。

CouchDB 还可以增强区块链中合规性和数据保护的安全性。因为它能够通过过滤和屏蔽交易中的各个属性来实现字段级安全性，并且仅在需要时才授权只读权限。

另外，CouchDB 属于 CAP 定理的 AP 类型 (可用性和分区容错)。它使用具有最终一致性的主-主复制模型。可以在 CouchDB 文档的最终一致性 (Eventual Consistency) 页面上找到更多信息。但是，在每个 Fabric 对端节点，没有数据库副本，可以保证对数据库的写入是一致且持久的 (不是最终一致性)。

CouchDB是第一个用于Fabric的外部可插入状态数据库，并且可能并且应该有其他外部数据库选项。例如，IBM为其区块链启用关系数据库。并且可能还需要CP类型（一致性和分区容差）数据库，以便在不保证应用程序级别的情况下实现数据一致性。

## 10. 关于数据持久性的说明

如果需要在对端节点容器或 CouchDB 容器上保持数据持久性，一种选择是将 docker-host 中的目录挂载到容器中的相关目录中。例如，你可以在 `docker-compose-base.yaml` 文件的对端节点容器规范中添加以下两行：



```ruby
volumes:
 - /var/hyperledger/peer0:/var/hyperledger/production
```

对于 CouchDB 容器，你可以在 CouchDB 容器规范中添加以下两行：



```ruby
volumes:
 - /var/hyperledger/couchdb0:/opt/couchdb/data
```

## 11. 故障排除

- 始终重新启动网络。使用以下命令删除组件：密钥，容器和链码镜像：



```ruby
$ sudo ./byfn.sh down
```

###### 注意

> 如果不删除旧的容器和镜像，将会看到错误。

- 如果你看到 Docker 错误，请首先检查你的 Docker 版本 (先决条件)，然后尝试重新启动 Docker 进程。Docker 的问题通常无法立即识别。例如，你可能会看到由于无法访问安装在容器中的加密组件而导致的错误。

  如果它们仍然存在，请删除你的镜像并从头开始：



```ruby
$ docker rm -f $(docker ps -aq)
$ docker rmi -f $(docker images -q)
```

- 如果在创建，实例化，调用或查询命令时看到错误，请确保已正确更新了通道名称和链码名称。提供的示例命令中包含占位符值。
- 如果看到以下错误：



```jsx
Error: Error endorsing chaincode: rpc error: code = 2 desc = Error installing chaincode code mycc:1.0(chaincode /var/hyperledger/production/chaincodes/mycc.1.0 exits)
```

你可能有以前运行的链码镜像 (例如 `dev-peer1.org2.example.com-mycc-1.0` 或 `dev-peer0.org1.example.com-mycc-1.0`）。删除它们，然后重试。



```ruby
$ sudo docker rmi -f $(docker images | grep peer[0-9]-peer[0-9] | awk '{print $3}')
```

- 如果你看到类似以下内容的内容：



```tsx
Error connecting: rpc error: code = 14 desc = grpc: RPC failed fast due to transport failure
Error: rpc error: code = 14 desc = grpc: RPC failed fast due to transport failure
```

确保针对重新标记为 "last" 的 "1.0.0" 运行网络。

- 如果看到以下错误：



```dart
[configtx/tool/localconfig] Load -> CRIT 002 Error reading configuration: Unsupported Config Type ""
panic: Error reading configuration: Unsupported Config Type ""
```

然后，你没有正确设置 `FABRIC_CFG_PATH` 环境变量。 configtxgen 工具需要此变量才能找到 `configtx.yaml`。返回并执行 `export FABRIC_CFG_PATH=$PWD`，然后重新创建侈的通道工件。

- 要清理网络，请使用 `down` 选项：



```ruby
$ sudo ./byfn.sh down
```

- 如果看到错误消息指出你仍然具有 "active endpoints"，请修剪 Docker 网络。这将清除以前的网络，并为你提供一个全新的环境：



```ruby
$ sudo docker network prune
```

你将看到以下消息：



```csharp
WARNING! This will remove all networks not used by at least one container.
Are you sure you want to continue? [y/N]
```

选择 `y`。

- 如果看到类似于以下内容的错误：



```jsx
/bin/bash: ./scripts/script.sh: /bin/bash^M: bad interpreter: No such file or directory
```

确保有问题的文件 （在本示例中为 script.sh) 以 Unix 格式编码。这很可能是由于未在 Git 配置中将 `core.autocrlf` 设置为 `false` 引起的 (请参阅 Windows 其他功能)。有几种解决方法。例如，如果你有权访问 vim 编辑器，请打开文件：



```undefined
vim ./fabric-samples/first-network/scripts/script.sh
```

然后通过执行以下 vim 命令来更改其格式：



```bash
:set ff=unix
```

###### 注意

> 如果仍然看到错误，请在 [Hyperledger Rocket Chat](https://links.jianshu.com/go?to=https%3A%2F%2Fchat.hyperledger.org%2Fhome) 或 [StackOverflow](https://links.jianshu.com/go?to=https%3A%2F%2Fstackoverflow.com%2Fquestions%2Ftagged%2Fhyperledger-fabric) 的结构问题通道上共享日志。

## 附录 1. 常用 Docker 命令

### 1.1 进入 Docker 容器

可以通过下面的命令进入 Docker 容器 cli 查看。



```bash
$ sudo docker exec -it cli /bin/sh
```

### 1.2 查看 Docker 容器日志

可以通过下面的命令查看 Docker 容器 cli 的日志。



```ruby
$ sudo docker logs -f cli
```

## Reference

- Building Your First Network, [https://hyperledger-fabric.readthedocs.io/en/latest/build_network.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Fbuild_network.html)
- Building Your First Network - v1.4, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/build_network.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fbuild_network.html)
- Prerequisites, [https://hyperledger-fabric.readthedocs.io/en/latest/prereqs.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Fprereqs.html)
- Install Samples, Binaries and Docker Images, [https://hyperledger-fabric.readthedocs.io/en/latest/install.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Finstall.html)
- fabric-samples, [https://github.com/hyperledger/fabric-samples](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fhyperledger%2Ffabric-samples)
- Docs » Glossary, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/glossary.html#channel](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23channel)
- Hyperledger Fabric Node.js Contract and Shim, [https://fabric-shim.github.io/release-1.4/index.html](https://links.jianshu.com/go?to=https%3A%2F%2Ffabric-shim.github.io%2Frelease-1.4%2Findex.html)
- fabric-chaincode-java, [https://fabric-chaincode-java.github.io/org/hyperledger/fabric/shim/Chaincode.html](https://links.jianshu.com/go?to=https%3A%2F%2Ffabric-chaincode-java.github.io%2Forg%2Fhyperledger%2Ffabric%2Fshim%2FChaincode.html)
- Docs » Glossary, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/glossary.html#member](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23member)
- Docs » Key Concepts » The Ordering Service, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/orderer/ordering_service.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Forderer%2Fordering_service.html)
- x.509 certificates and public key infrastructure, [https://en.wikipedia.org/wiki/Public_key_infrastructure](https://links.jianshu.com/go?to=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FPublic_key_infrastructure)
- cryptogen, [https://hyperledger-fabric.readthedocs.io/en/latest/commands/cryptogen.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Fcommands%2Fcryptogen.html)
- configtxgen, [https://hyperledger-fabric.readthedocs.io/en/latest/commands/configtxgen.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Fcommands%2Fconfigtxgen.html)
- Docs » Glossary, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/glossary.html#genesis-block](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23genesis-block)
- Docs » Glossary, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/glossary.html#channel](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23channel)
- Docs » Glossary, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/glossary.html#anchor-peer](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23anchor-peer)
- Docs » Tutorials » Chaincode for Developers, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/chaincode4ade.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fchaincode4ade.html)
- Docs » Operations Guides » Endorsement policies, [https://hyperledger-fabric.readthedocs.io/en/latest/endorsement-policies.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Fendorsement-policies.html)
- hyperledger/fabric-sdk-node, [https://github.com/hyperledger/fabric-sdk-node](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fhyperledger%2Ffabric-sdk-node)
- Hyperledger Chat, [https://chat.hyperledger.org/home](https://links.jianshu.com/go?to=https%3A%2F%2Fchat.hyperledger.org%2Fhome)
- Questions tagged [hyperledger-fabric]
   , [https://stackoverflow.com/questions/tagged/hyperledger-fabric](https://links.jianshu.com/go?to=https%3A%2F%2Fstackoverflow.com%2Fquestions%2Ftagged%2Fhyperledger-fabric)

## 项目源代码

项目源代码会逐步上传到 Github，地址为 [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)。

## Contributor

1. Windstamp, [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)



作者：furnace
链接：https://www.jianshu.com/p/2c8e15a973e8
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。



# 5.1 Hyperledger Fabric 专题 - 构建第一个 Hyperledger Fabric 网络



Hyperledger Fabric 最新的文档基于版本是 v2.0 Alpha release。由于示例中相关的 docker 镜像的版本是 v1.4.3，因为相关文档需要参考的版本为 v1.4.3。这些文档链接在本文的 Reference 部分都有涉及。

## 1. 本文目的

参考文档 [Building Your First Network](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Fbuild_network.html) 搭建第一个 Hyperledger Fabric 网络。本示例 [fabric-samples](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fhyperledger%2Ffabric-samples) 需要的 hyperledger fabric 相关的二进制程序安装参考文档 [Install Samples, Binaries and Docker Images](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Finstall.html)。而安装 hyperledger fabric 相关的二进制程序所需要的一些前置软件需求，请参考文档 [Prerequisites](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Fprereqs.html)。

## 2. 构建 Hyperledger Fabric 网络

根据 fabric-samples/first-network 示例搭建第一个 Hyperledger Fabric 网络。这里，主要是基于提供的脚本 byfn.sh 来快速搭建。脚本 byfn.sh 提供了许多可配置的参数，如果不提供则使用默认值。这里为了简化，先使用默认值。

需要注意的是，下面的命令需要在 first-network 目录下运行，但是调用的 hyperledger fabric 相关的二进制程序却是在 fabric-samples/bin 目录。

### 2.1 脚本 byfn.sh 的详细说明



```bash
Usage:
byfn.sh <mode> [-c <channel name>] [-t <timeout>] [-d <delay>] [-f <docker-compose-file>] [-s <dbtype>] [-l <language>] [-o <consensus-type>] [-i <imagetag>] [-v]"
  <mode> - one of 'up', 'down', 'restart', 'generate' or 'upgrade'"
    - 'up' - bring up the network with docker-compose up"
    - 'down' - clear the network with docker-compose down"
    - 'restart' - restart the network"
    - 'generate' - generate required certificates and genesis block"
    - 'upgrade'  - upgrade the network from version 1.3.x to 1.4.0"
  -c <channel name> - channel name to use (defaults to \"mychannel\")"
  -t <timeout> - CLI timeout duration in seconds (defaults to 10)"
  -d <delay> - delay duration in seconds (defaults to 3)"
  -f <docker-compose-file> - specify which docker-compose file use (defaults to docker-compose-cli.yaml)"
  -s <dbtype> - the database backend to use: goleveldb (default) or couchdb"
  -l <language> - the chaincode language: golang (default), node, or java"
  -o <consensus-type> - the consensus-type of the ordering service: solo (default), kafka, or etcdraft"
  -i <imagetag> - the tag to be used to launch the network (defaults to \"latest\")"
  -v - verbose mode"
byfn.sh -h (print this message)"

Typically, one would first generate the required certificates and
genesis block, then bring up the network. e.g.:"

  byfn.sh generate -c mychannel"
  byfn.sh up -c mychannel -s couchdb"
  byfn.sh up -c mychannel -s couchdb -i 1.4.0"
  byfn.sh up -l node"
  byfn.sh down -c mychannel"
  byfn.sh upgrade -c mychannel"

Taking all defaults:"
      byfn.sh generate"
      byfn.sh up"
      byfn.sh down"
```

### 2.2 使用脚本 byfn.sh 构建 Fabric 网络的详细步骤

#### Step 1. Generate Network Artifacts



```ruby
$ ./byfn.sh generate
```

#### Step 2. Bring Up the Network



```ruby
$ sudo ./byfn.sh up
```

#### Step 3. Bring Down the Network



```ruby
$ sudo ./byfn.sh down
```

## 3. 拓展

在这一部分，将把上述通过脚本 byfn.sh 构建的过程每一个关键步骤进行详细解构，并通过手动执行命令的方式来构建 hyperledger fabric 网络。

### 3.0 前置操作

下面的操作需要将容器 cli 的 FABRIC_LOGGING_SPEC 日志级别从 INFO 调整为 DEBUG。

first-network/docker-compose-cli.yaml



```php
cli:
  container_name: cli
  image: hyperledger/fabric-tools:$IMAGE_TAG
  tty: true
  stdin_open: true
  environment:
    - GOPATH=/opt/gopath
    - CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock
    - FABRIC_LOGGING_SPEC=DEBUG
    #- FABRIC_LOGGING_SPEC=INFO
```

### 3.1 Crypto Generator

- cryptogen
  - 生成密码组件的工具
  - 文件位置为 fabric-samples/bin/cryptogen
- crypto-config.yaml
  - 配置文件，用于说明示例 first-network 需要生成的密码组件
  - 文件位置为 first-network/crypto-config.yaml
- crypto-config
  - 示例 first-network 生成的密码组件
  - 文件位置为 first-network/crypto-config

### 3.2 Configuration Transaction Generator

- configtxgen
  - 生成构件，如
    - orderer `genesis block`,
    - channel `configuration transaction`,
    - and two `anchor peer transactions` - one for each Peer Org.
  - 文件位置为 fabric-samples/bin/configtxgen
- configtx.yaml.yaml
  - 配置文件，用于说明示例 first-network 需要的构件，一个 Orderer Org (`OrdererOrg`)，和两个 Peer Orgs (`Org1` & `Org2`)。
  - 文件位置为 first-network/configtx.yaml
- channel-artifacts
  - 示例 first-network 生成构件
    - channel.tx
    - genesis.block
    - Org1MSPanchors.tx
    - Org2MSPanchors.tx
  - 文件位置为 first-network/channel-artifacts

### 3.3 通过工具独立运行上述命令

下面，我们通过命令 cryptogen 和 configtxgen  手动生成上述的密码组件和示例构件。

#### 3.3.1 手动生成密码组件

参考脚本 byfn.sh 中的函数 generateCerts().



```ruby
$ ../bin/cryptogen generate --config=./crypto-config.yaml
org1.example.com
org2.example.com
```

最终的密码组件会生成到目录 first-network/crypto-config 中。

#### 3.3.2 手动生成示例构件

参考脚本 byfn.sh 中的函数 generateChannelArtifacts().



```ruby
$ export FABRIC_CFG_PATH=$PWD
$ ../bin/configtxgen -profile TwoOrgsOrdererGenesis -channelID byfn-sys-channel -outputBlock ./channel-artifacts/genesis.block
```

最终的示例构件会生成到目录 first-network/channel-artifacts 中，如 first-network/channel-artifacts/genesis.block。

#### 3.3.3 Create a Channel Configuration Transaction

注意，需要设置环境变量 CHANNEL_NAME。

###### 创建 first-network/channel-artifacts/channel.tx



```bash
# The channel.tx artifact contains the definitions for our sample channel

$ export CHANNEL_NAME=mychannel  && ../bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME
```

###### 创建 first-network/channel-artifacts/Org1MSPanchors.tx



```ruby
$ ../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP
```

###### 创建 first-network/channel-artifacts/Org2MSPanchors.tx



```ruby
$ ../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org2MSP
```

### 3.4 Start the network

如果之前利用脚本 byfn.sh 启动了测试网络，需要先关闭，如执行下列命令



```ruby
$ sudo ./byfn.sh down
```

#### 3.4.1 Start the network



```ruby
$ sudo docker-compose -f docker-compose-cli.yaml up -d
```

或通过下面的命令启动，能够立刻显示日志，但需要另启一个窗口运行下面的 cli 容器。



```ruby
$ sudo docker-compose -f docker-compose-cli.yaml up
```

#### 3.4.2 Create & Join Channel



```ruby
$ sudo docker exec -it cli bash
root@33ab5acc5622:/opt/gopath/src/github.com/hyperledger/fabric/peer#
```

在容器 cli 中需要设置的环境变量。



```ruby
# Environment variables for PEER0

$ CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
$ CORE_PEER_ADDRESS=peer0.org1.example.com:7051
$ CORE_PEER_LOCALMSPID="Org1MSP"
$ CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
```



```ruby
$ export CHANNEL_NAME=mychannel

# the channel.tx file is mounted in the channel-artifacts directory within your CLI container
# as a result, we pass the full path for the file
# we also pass the path for the orderer ca-cert in order to verify the TLS handshake
# be sure to export or replace the $CHANNEL_NAME variable appropriately

$ peer channel create -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```

上述命令将会生成文件 ``，这里是 `mychannel.block`。

###### 将 peer0.org1.example.com 加入到 channel



```ruby
# By default, this joins ``peer0.org1.example.com`` only
# the <CHANNEL_NAME.block> was returned by the previous command
# if you have not modified the channel name, you will join with mychannel.block
# if you have created a different channel name, then pass in the appropriately named block

$ peer channel join -b mychannel.block
```

###### 将 peer0.org2.example.com 加入到 channel



```ruby
$ CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
$ CORE_PEER_ADDRESS=peer0.org2.example.com:9051
$ CORE_PEER_LOCALMSPID="Org2MSP"
$ CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
$ peer channel join -b mychannel.block
```

#### 3.4.3 Update the anchor peers

###### 将 Org1 定义为 peer0.org1.example.com



```ruby
$ CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
$ CORE_PEER_ADDRESS=peer0.org1.example.com:7051
$ CORE_PEER_LOCALMSPID="Org1MSP"
$ CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
$ peer channel update -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/Org1MSPanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```

###### 将 Org2 定义为 peer0.org2.example.com



```ruby
$ CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
$ CORE_PEER_ADDRESS=peer0.org2.example.com:9051
$ CORE_PEER_LOCALMSPID="Org2MSP"
$ CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
$ peer channel update -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/Org2MSPanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```

#### 3.4.4 Install & Instantiate Chaincode

Chaincode 代码支持 golang, node, java 语言，默认为 golang，本示例中全部使用默认值。

###### Install peer0 in Org1

首先设置 Org1 相关的环境变量。



```ruby
$ CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
$ CORE_PEER_ADDRESS=peer0.org2.example.com:9051
$ CORE_PEER_LOCALMSPID="Org2MSP"
$ CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
```



```ruby
# this installs the Go chaincode. For go chaincode -p takes the relative path from $GOPATH/src
$ peer chaincode install -n mycc -v 1.0 -p github.com/chaincode/chaincode_example02/go/
```

###### Install peer0 in Org2

首先设置 Org2 相关的环境变量。



```ruby
# Environment variables for PEER0 in Org2

$ CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
$ CORE_PEER_ADDRESS=peer0.org2.example.com:9051
$ CORE_PEER_LOCALMSPID="Org2MSP"
$ CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
```



```ruby
# this installs the Go chaincode. For go chaincode -p takes the relative path from $GOPATH/src
$ peer chaincode install -n mycc -v 1.0 -p github.com/chaincode/chaincode_example02/go/
```

###### Instantiate the chaincode on the channel



```ruby
# be sure to replace the $CHANNEL_NAME environment variable if you have not exported it
# if you did not install your chaincode with a name of mycc, then modify that argument as well
$ export CHANNEL_NAME=mychannel
$ peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc -v 1.0 -c '{"Args":["init","a", "100", "b","200"]}' -P "AND ('Org1MSP.peer','Org2MSP.peer')"
```

注意，`-P "AND ('Org1MSP.peer','Org2MSP.peer')"` 中的 `AND` 可以改为 `OR`，这意味着只需要 Org1 或 Org2 其中一个组织背书就可以了。

#### 3.4.5 Query



```ruby
# be sure to set the -C and -n flags appropriately

$ peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["query","a"]}'
```

#### 3.4.6 Invoke



```swift
# be sure to set the -C and -n flags appropriately

$ peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"Args":["invoke","a","b","10"]}'
```

#### 3.4.7 Query



```ruby
# be sure to set the -C and -n flags appropriately

$ peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["query","a"]}'
```

返回结果



```rust
Query Result: 90
```

#### 3.4.8 Install peer1 in Org2

首先设置好环境变量。



```ruby
# Environment variables for PEER1 in Org2

$ CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
$ CORE_PEER_ADDRESS=peer1.org2.example.com:10051
$ CORE_PEER_LOCALMSPID="Org2MSP"
$ CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/tls/ca.crt
```



```ruby
# this installs the Go chaincode. For go chaincode -p takes the relative path from $GOPATH/src
$ peer chaincode install -n mycc -v 1.0 -p github.com/chaincode/chaincode_example02/go/
```

#### 3.4.9 Query by peer1 in Org2

首先需要将 peer1 in Org2 加入到 channel，之后才能响应查询。



```ruby
$ CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
$ CORE_PEER_ADDRESS=peer1.org2.example.com:10051
$ CORE_PEER_LOCALMSPID="Org2MSP"
$ CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/tls/ca.crt
$ peer channel join -b mychannel.block
```



```ruby
# be sure to set the -C and -n flags appropriately

$ peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["query","a"]}'
```

返回结果



```rust
Query Result: 90
```

#### 3.4.10 How do I see these transactions?



```ruby
$ sudo docker logs -f cli
```

#### 3.4.11 How can I see the chaincode logs?



```ruby
$ sudo docker logs dev-peer0.org2.example.com-mycc-1.0
ex02 Init
Aval = 100, Bval = 200
ex02 Invoke
Query Response:{"Name":"a","Amount":"100"}
ex02 Invoke
Aval = 90, Bval = 210
ex02 Invoke
Query Response:{"Name":"a","Amount":"90"}

$ sudo docker logs dev-peer0.org1.example.com-mycc-1.0
ex02 Invoke
Aval = 90, Bval = 210

$ sudo docker logs dev-peer1.org2.example.com-mycc-1.0
ex02 Invoke
Query Response:{"Name":"a","Amount":"90"}
```

### 3.5 Using CouchDB



```ruby
$ sudo docker-compose -f docker-compose-cli.yaml -f docker-compose-couch.yaml up -d
```

## Reference

- Building Your First Network, [https://hyperledger-fabric.readthedocs.io/en/latest/build_network.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Fbuild_network.html)
- Building Your First Network - v1.4, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/build_network.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fbuild_network.html)
- Prerequisites, [https://hyperledger-fabric.readthedocs.io/en/latest/prereqs.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Fprereqs.html)
- Install Samples, Binaries and Docker Images, [https://hyperledger-fabric.readthedocs.io/en/latest/install.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Finstall.html)
- fabric-samples, [https://github.com/hyperledger/fabric-samples](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fhyperledger%2Ffabric-samples)
- x.509 certificates and public key infrastructure, [https://en.wikipedia.org/wiki/Public_key_infrastructure](https://links.jianshu.com/go?to=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FPublic_key_infrastructure)
- cryptogen, [https://hyperledger-fabric.readthedocs.io/en/latest/commands/cryptogen.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Fcommands%2Fcryptogen.html)
- configtxgen, [https://hyperledger-fabric.readthedocs.io/en/latest/commands/configtxgen.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Flatest%2Fcommands%2Fconfigtxgen.html)

## 项目源代码

项目源代码会逐步上传到 Github，地址为 [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)。

## Contributor

1. Windstamp, [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)



作者：furnace
链接：https://www.jianshu.com/p/185c648480b0
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。





# Hyperledger Fabric 专题 - Transaction Flow

本专题文档概述了标准资产交换过程中发生的交易机制。该方案包括买卖萝卜的两个客户 A 和 B。他们每个人在网络上都有一个对端节点，通过它们发送交易并与账本进行交互。

![img](https:////upload-images.jianshu.io/upload_images/6280489-6ce8ac684661768c.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

## 0. 一些前提假设

此流程假定已建立并运行通道。应用程序用户已经在组织的证书颁发机构 (Certificate Authority, CA) 中注册，并收到了必要的加密材料，用于对网络进行身份验证。

链码 (包含一组代表萝卜市场初始状态的键值对) 安装在对端节点上，并在通道上实例化。链码包含定义了一组交易指令和买卖萝卜价格的逻辑。还为此链代码设置了一个背书策略 (endorsement policy)，指出 peerA 和 peerB 都必须背书 (endorse) 任何交易。

## 1. 客户 A 发起交易

![img](https:////upload-images.jianshu.io/upload_images/6280489-a8ec3d618f121caa.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

发生了什么？客户端 A 正在发送购买萝卜的请求。该请求以分别代表客户端 A 和客户端 B 的对端节点 A 和对端节点 B 为目标。背书策略声明两个对端节点都必须背书任何交易，因此该请求将发送到对端节点 A 和对端节点 B。

接下来，构造交易提案。应用程序使用 SDK (Node.js，Java，Python) 及其 API 来生成交易提案。该提案即包含某些输入参数的链码函数调用，用于读取和/或更新帐本。

可以通过 SDK 将交易提案打包为适当架构的格式 (用于 gRPC 的 protocol buffer)，并使用用户的加密证书为该交易提案生成唯一的签名。

## 2. 背书对端节点验证签名并执行交易

![img](https:////upload-images.jianshu.io/upload_images/6280489-635ff09920c9da6a.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

背书对端节点验证：(1) 交易提案格式是否正确，(2) 过去未提交过此交易提案 (重放攻击保护)，（3）签名是否有效 (使用MSP) 和 (4) 提交者（在示例中为客户 A）是否被授权在该通道上执行提案的操作 (即，每个背书对端节点都确保提交者满足该通道的 Writer 策略)。背书对端节点将交易提案输入作为所调用链码函数的参数。然后针对当前状态数据库执行链码以产生包括响应值，读取集和写入集 (即表示要创建或更新资产的键/值对) 的交易结果。此时，不会对帐本进行任何更新。这些值的集合以及背书对端节点的签名作为“提案响应”传递回 SDK，该 SDK 解析有效负载供应用程序使用。

###### 注解

MSP 是对端节点组件，允许对端节点验证来自客户端的交易请求并签署交易结果 (背书)。写入策略 (writing policy) 是在通道创建时定义的，并确定哪些用户有权向该通道提交交易。有关 MSP 的更多信息，请查看我们的 [MSP](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fmembership%2Fmembership.html) 文档。

## 3. 检查提案响应

![img](https:////upload-images.jianshu.io/upload_images/6280489-2d8a9591da06edfe.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

该应用程序验证背书对端节点的签名，并比较提案响应以确定提案响应是否相同。如果仅通过链码查询账本，则应用程序将检查查询响应，并且通常不将交易提交给交易排序服务。如果客户端应用程序打算将交易提交给交易排序服务以更新帐本，则应用程序将确定提交之前是否已满足指定的背书策略 (即 peerA 和 peerB 都背书了)。该体系结构使得即使应用程序选择不检查响应或以其他方式转发未背书的交易，背书策略仍将由对端节点强制执行，并在提交验证阶段得到维护。

## 4. 客户将背书组合成交易

![img](https:////upload-images.jianshu.io/upload_images/6280489-7c02eca516b65ea8.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

该应用程序在“交易消息”中将交易提案和响应“广播”到交易排序服务。交易将包含读/写集，背书对端节点的签名和通道 ID。交易排序服务不需要检查交易的全部内容即可执行其操作，它仅从网络中的所有通道接收交易，由通道按时间对它们进行排序，并为每个通道创建交易区块。

## 5. 交易已验证并提交

![img](https:////upload-images.jianshu.io/upload_images/6280489-117221dd817f947f.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

交易区块被分发给通道上的所有对端节点。验证区块中的交易以确保实现了背书策略，并确保自经过交易执行生成以来，读取集变量的帐本状态没有发生变化。区块中的交易被标记为有效或无效。

## 6. 更新账本

![img](https:////upload-images.jianshu.io/upload_images/6280489-09371f5851b60bfd.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

每个对端节点都会将该区块添加到通道链中，并且对于每个有效交易，写集都将提交到当前状态数据库。发出一个事件，以通知客户端应用程序交易 (调用) 已被不可变地添加到链上，并通知交易是已验证还是无效。

###### 注解

应用程序应在提交交易后侦听交易事件，例如，通过使用自动侦听交易事件的 `submitTransaction` API。如果不监听交易事件，你将不会知道你的交易是否实际上已被排序，验证并提交到账本。

请参考 [序列化流程图](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Farch-deep-dive.html%23swimlane) 以更好地了解交易流程。

## Reference

1. Docs » Architecture Reference » Transaction Flow, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/txflow.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Ftxflow.html)
2. Docs » Key Concepts » Membership, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/membership/membership.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fmembership%2Fmembership.html)
3. Docs » Architecture Reference » Architecture Origins, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/arch-deep-dive.html#swimlane](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Farch-deep-dive.html%23swimlane)

## 项目源代码

项目源代码会逐步上传到 Github，地址为 [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)。

## Contributor

1. Windstamp, [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)



作者：furnace
链接：https://www.jianshu.com/p/db04e13163cf
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

# Hyperledger Fabric 专题 - 编写第一个 Hyperledger Fabric 应用

注意，本示例选择 javascript 版本。

###### 注解

> 如果你还不熟悉 Fabric 网络的基本架构，则可能需要先访问 [关键概念](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fkey_concepts.html) 部分，然后再继续。

> 同样值得注意的是，本文是 Fabric 应用程序的入门教程，使用简单的智能合约和应用程序。要更深入地了解 Fabric 应用程序和智能合约，请查看我们的 [开发应用程序](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fdevelopapps%2Fdeveloping_applications.html) 部分或 [商业论文](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Ftutorial%2Fcommercial_paper.html) 教程。

在本教程中，我们将查看一些示例程序，以了解 Fabric 应用程序的工作方式。这些应用程序及其使用的智能合约统称为 `FabCar`。它们为了解 Hyperledger Fabric 区块链提供了一个很好的起点。你将学习如何编写应用程序和智能合约以查询和更新帐本，以及如何使用证书颁发机构 (Certificate Authority, CA) 生成 X.509 证书，以供与许可制区块链进行交互的应用程序使用。

我们将使用应用程序 SDK (在 [应用程序](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fdevelopapps%2Fapplication.html) 主题中进行了详细描述) 来调用智能合约，该智能合约使用智能合约 SDK
 来查询和更新帐本 (在 [智能合约处理](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fdevelopapps%2Fsmartcontract.html) 部分中进行了详细介绍)。

我们将经历三个主要步骤：

1. 设置开发环境。我们的应用程序需要与之交互的网络，因此我们将获得一个智能合约和应用程序将使用的基本网络。

![img](https:////upload-images.jianshu.io/upload_images/6280489-13d2f3c86d8725a0.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

1. 了解智能合约示例 FabCar。我们将检查智能合约，以了解其中的交易以及应用程序如何使用它们查询和更新账本。
2. 开发一个使用 FabCar 的示例应用程序。我们的应用程序将使用 FabCar 智能合约来查询和更新账本上的汽车资产。我们将深入研究应用程序及其创建的交易代码，包括查询汽车，查询一系列汽车以及创建新汽车。

完成本教程后，你应该对如何结合应用程序和智能合约与 Fabric 网络中对端节点上托管的账本副本进行交互有基本了解。

###### 注解

> 这些应用程序还与 [服务发现](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fdiscovery-overview.html) 和 [私有数据](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fprivate-data%2Fprivate-data.html) 兼容，尽管我们不会明确显示如何使用我们的应用程序来利用这些功能。

## 1. 搭建区块链网络

###### 注解

> 下一部分将要求你位于 `fabric-samples` 本地克隆的子目录  `first-network` 中。

如果你已经完成了 [构建第一个网络](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fbuild_network.html) 的准备工作，那么你已下载 `fabric-samples` 并有了一个运行的网络。在运行本教程之前，必须停止该网络：



```ruby
$ cd /path/to/fabric-samples/first-network
$ sudo ./byfn.sh down
```

如果你之前已经运行过本教程，请使用以下命令杀死任何陈旧或活动的容器。请注意，这将删除你所有与 Fabric 无关的容器。



```ruby
$ sudo docker rm -f $(sudo docker ps -aq)
$ sudo docker rmi -f $(sudo docker images | grep fabcar | awk '{print $3}')
```

如果你没有开发环境以及网络和应用程序随附的构件，请访问 [先决条件](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fprereqs.html) 页面，并确保在计算机上安装了必要的依赖项。

接下来，如果你还没有这样做，请访问 [安装示例，二进制文件和 Docker 映像](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Finstall.html) 页面，并按照提供的说明进行操作。克隆了 `fabric-samples` 并下载了最新的稳定 Fabric 映像和可用实用程序后，请返回本教程。

### 1.1 启动网络

###### 注解

> 下一部分将要求你位于 `fabric-samples` 本地克隆的子目录  `fabcarc` 中。

> 本教程演示了 FabCar 智能合约和应用程序的 JavaScript 版本，但是 `fabric-samples` 仓库还包含此示例的 Java 和 TypeScript 版本。要尝试 Java 或 TypeScript 版本，请将下面的 `./startFabric.sh` 的 `javascript` 参数更改为 `java` 或 `typescript`，然后按照写入终端的说明进行操作。

使用 `startFabric.sh` Shell 脚本启动网络。该命令将启动一个由对端网络，交易排序，证书颁发机构等组成的区块链网络。它还将安装并实例化 `FabCar` 智能合约的 JavaScript 版本，供我们的应用程序用来访问帐本。在学习本教程时，我们将详细了解这些组件。



```ruby
$ cd /path/to/fabric-samples/fabcar
$ sudo ./startFabric.sh javascript
```

注意，上述脚本需要几分钟才能执行完成，而且还需要稳定的网络环境，建议在 AWS EC2 上实验。

同时，还需要注意权限问题，因为 docker 需要 root 权限，因此整个目录 `first-samples` 的所有者有可能已经是 root:root。然后你用非 root 的用户在进行写操作时有可能出现权限问题。

好了，你现在已经建立并运行了一个示例网络，并且已安装和实例化了 `FabCar` 智能合约。让我们先安装我们的应用程序先决条件，以便我们可以尝试一下，并了解一切如何协同工作。

### 1.2 安装应用程序

###### 注解

> 下一部分将要求你位于 `fabric-samples` 本地克隆的子目录  `fabcarc/javascript` 中。

运行以下命令以安装应用程序的 Fabric 依赖。大约需要一分钟才能完成：



```ruby
$ cd /path/to/fabric-samples/fabcar/javascript
$ npm install
```

此过程将安装 package.json 中定义的关键应用程序依赖项。其中最重要的是 `fabric-network` 类。它使应用程序能够使用身份标识，钱包，和网关来连接到通道，提交交易并等待通知。本教程还使用 `fabric-ca-client` 类为具有各自证书颁发机构的用户注册，生成有效的身份，然后由 `fabric-network` 类方法使用该身份标识。

等 npm 安装完成后，一切就绪，可以运行该应用程序。对于本教程，你将主要使用 `fabcar/javascript` 目录中的应用程序 JavaScript 文件。让我们看一下里面的东西：



```go
$ ls
enrollAdmin.js  node_modules       package.json  registerUser.js
invoke.js       package-lock.json  query.js      wallet
```

在 `fabcar/typescript` 目录中有其他程序语言的文件。使用 JavaScript 示例后，你可以阅读这些内容 - 原理相同。

## 2. 注册管理员用户

###### 注解

> 以下两节涉及与证书颁发机构的通信。你可能会发现通过运行新的 shell 终端并运行 sudo docker logs -f ca.example.com 在运行即将到来的程序时流式传输 CA 日志很有用。



```ruby
$ sudo docker logs -f ca.example.com
```

但示例中，CA 的 docker 容器应该是 ca_peerOrg1 和 ca_peerOrg2，而非 ca.example.com。测试下来，管理员用户 admin 使用的是 ca_peerOrg1。



```ruby
$ sudo docker logs -f ca_peerOrg1
$ sudo docker logs -f ca_peerOrg2
```

当我们创建网络时，创建了一个管理员用户 (字面上称为 admin) 作为证书颁发机构 (Certificate Authority, CA) 的登记员 (rigistrar)。我们的第一步是使用 `enroll.js` 程序为 `admin` 生成私钥，公钥和 X.509 证书。此过程使用证书签名请求 (Certificate Signing Request, CSR) - 首先在本地生成私钥和公钥，然后将公钥发送到 CA，CA 返回编码的证书以供应用程序使用。然后将这三个凭证存储在钱包中，使我们能够充当 CA 的管理员。

随后，我们将注册并登记一个新的应用程序用户，我们的应用程序将使用该用户与区块链进行交互。

让我们注册用户 `admin`：



```ruby
$ cd /path/to/fabric-samples/fabcar/javascript
$ node enrollAdmin.js
```

此命令已将 CA 管理员的凭据存储在 `wallet` 目录中。

## 3. 注册并登记 `user1`

现在我们已经有了管理员的凭据，可以在电子钱包中注册一个新用户 `user1`，该用户将用于查询和更新帐本：



```ruby
$ cd /path/to/fabric-samples/fabcar/javascript
$ node registerUser.js
```

与 `admin` 注册类似，此程序使用 CSR 来注册 `user1` 并将其凭据与 `admin` 的凭据一起存储在钱包中。现在，我们有两个独立用户的身份 `admin` 和 `user1`，我们的应用程序使用了这些身份。

是时候与分类帐进行交互了 ......

## 4. 查询账本

区块链网络中的每个对端节点都托管账本的副本，应用程序可以通过调用智能合约来查询账本，该合约查询账本的最新值并将其返回给应用程序。

这是查询工作方式的简化表示：



![img](https:////upload-images.jianshu.io/upload_images/6280489-d02a36b687ca8672.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

应用程序使用查询从帐本读取数据。最常见的查询涉及帐本中数据的当前值 - 其世界状态。世界状态表示为一组键值对，应用程序可以查询数据以获取单个键或多个键。此外，账本世界状态可以配置为使用像 CouchDB 这样的数据库，当将键值建模为 JSON 数据时，该数据库支持复杂的查询。当寻找与某些关键字匹配特定值的所有资产时，这将非常有用。例如，所有具有特定所有者的汽车。

首先，我们运行 `query.js` 程序以返回帐本中所有汽车的清单。该程序使用我们的第二个身份 `user1` 来访问帐本：



```ruby
$ cd /path/to/fabric-samples/fabcar/javascript
$ node query.js
```

输出如下所示：



```csharp
Wallet path: ...fabric-samples/fabcar/javascript/wallet
Transaction has been evaluated, result is:
[{"Key":"CAR0", "Record":{"colour":"blue","make":"Toyota","model":"Prius","owner":"Tomoko"}},
{"Key":"CAR1", "Record":{"colour":"red","make":"Ford","model":"Mustang","owner":"Brad"}},
{"Key":"CAR2", "Record":{"colour":"green","make":"Hyundai","model":"Tucson","owner":"Jin Soo"}},
{"Key":"CAR3", "Record":{"colour":"yellow","make":"Volkswagen","model":"Passat","owner":"Max"}},
{"Key":"CAR4", "Record":{"colour":"black","make":"Tesla","model":"S","owner":"Adriana"}},
{"Key":"CAR5", "Record":{"colour":"purple","make":"Peugeot","model":"205","owner":"Michel"}},
{"Key":"CAR6", "Record":{"colour":"white","make":"Chery","model":"S22L","owner":"Aarav"}},
{"Key":"CAR7", "Record":{"colour":"violet","make":"Fiat","model":"Punto","owner":"Pari"}},
{"Key":"CAR8", "Record":{"colour":"indigo","make":"Tata","model":"Nano","owner":"Valeria"}},
{"Key":"CAR9", "Record":{"colour":"brown","make":"Holden","model":"Barina","owner":"Shotaro"}}]
```

让我们仔细看看这个程序。使用编辑器 (例如 atom 或 Visual Studio) 打开 `query.js`。

该应用程序首先从 `fabric-network` 模块中给作用域引入两个关键类： `FileSystemWallet` 和 `Gateway`。这些类将用于在钱包中定位 `user1` 身份，并将其用于连接到网络：



```jsx
const { FileSystemWallet, Gateway } = require('fabric-network');
```

该应用程序使用网关连接到网络：



```csharp
const gateway = new Gateway();
await gateway.connect(ccp, { wallet, identity: 'user1' });
```

此代码创建一个新的网关，然后使用它将应用程序连接到网络。`ccp` 描述了网关将使用钱包中的标识 `user1` 访问的网络。查看如何从 `../../basic-network/connection.json` 加载 `ccp` 并将其解析为 `JSON` 文件：



```jsx
const ccpPath = path.resolve(__dirname, '..', '..', 'basic-network', 'connection.json');
const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
const ccp = JSON.parse(ccpJSON);
```

如果你想进一步了解连接配置文件的结构以及它如何定义网络，请查看 [连接配置文件主题](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fdevelopapps%2Fconnectionprofile.html)。

一个网络可以分为多个通道，下一行重要的代码行将应用程序连接到网络中的特定通道 `mychannel`：



```csharp
const network = await gateway.getNetwork('mychannel');
```

在此通道中，我们可以访问智能合约 `fabcar` 与帐本进行交互：



```csharp
const contract = network.getContract('fabcar');
```

`fabcar` 中有许多不同的交易，并且我们的应用程序最初使用 `queryAllCars` 交易来访问帐本世界状态数据：



```csharp
const result = await contract.evaluateTransaction('queryAllCars');
```

`evaluateTransaction` 方法表示与区块链网络中智能合约的最简单交互之一。它只是选择一个在连接配置文件中定义的对端节点，然后将请求发送到该对端节点，并在此对其进行评估。智能合约会查询对端节点账本副本上的所有汽车，并将结果返回给应用程序。此交互不会导致更新帐本。

## 5. 智能合约 FabCar

让我们看一下智能合约 `FabCar` 中的交易。导航到 `fabric-samples` 根目录下的 `chaincode/fabcar/javascript/lib` 子目录，然后在编辑器中打开 `fabcar.js`。

查看如何使用 `Contract` 类定义智能合约：



```dart
class FabCar extends Contract {...
```

在该类结构中，你将看到我们定义了以下交易：`initLedger`，`queryCar`，`queryAllCars`，`createCar` 和 `changeCarOwner`。例如：



```csharp
async queryCar(ctx, carNumber) {...}
async queryAllCars(ctx) {...}
```

让我们仔细看看 `queryAllCars` 交易，看看它如何与账本交互。



```csharp
async queryAllCars(ctx) {

  const startKey = 'CAR0';
  const endKey = 'CAR999';

  const iterator = await ctx.stub.getStateByRange(startKey, endKey);
```

此代码定义了 `queryAllCars` 将要从账本中检索的汽车范围。查询将返回介于 `CAR0` 和 `CAR999` 之间的每辆汽车 - 总共 1000 辆汽车 (假设每个键均已正确标记)。其余代码遍历查询结果，并将其打包为应用程序的 JSON。

下面是应用程序如何调用智能合约中的不同交易的图示。每个交易使用一组广泛的 API (例如 `getStateByRange`) 与帐本进行交互。你可以详细了解有关这些 [API](https://links.jianshu.com/go?to=https%3A%2F%2Ffabric-shim.github.io%2Fmaster%2Findex.html%3Fredirect%3Dtrue) 的更多信息。

![img](https:////upload-images.jianshu.io/upload_images/6280489-8c4662f85335f02a.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

我们可以看到 `queryAllCars` 交易和另一个名为 `createCar` 的交易。我们将在本教程的后面部分中使用它来更新帐本，并向区块链添加一个新区块。

但首先，请返回 `query` 程序，然后将 `evaluateTransaction` 请求更改为查询 `CAR4`。`query` 程序现在应如下所示：



```csharp
const result = await contract.evaluateTransaction('queryCar', 'CAR4');
```

保存程序，然后导航回到 `fabcar/javascript` 目录。现在再次运行查询程序：



```ruby
$ cd /path/to/fabric-samples/fabcar/javascript
$ node query.js
```

你应该看到以下内容：



```csharp
Wallet path: ...fabric-samples/fabcar/javascript/wallet
Transaction has been evaluated, result is:
{"colour":"black","make":"Tesla","model":"S","owner":"Adriana"}
```

如果你回头查看交易 `queryAllCars` 的结果，你会发现 `CAR4` 是 Adriana 的黑色 Tesla model S，这是在此处返回的结果。

我们可以使用 `queryCar` 交易通过其关键字 (例如 CAR0) 查询任何汽车，并获取与该汽车相对应的任何品牌，型号，颜色和所有者。

此时，你应该熟悉智能合约中的基本查询交易和查询程序中的少数参数。

是时候更新账本了......

## 6. 更新账本

现在，我们已经完成了一些账本查询并添加了一些代码，我们可以更新账本了。我们可以进行很多潜在的更新，但让我们首先创建一辆新车。

从应用程序的角度来看，更新帐本很简单。应用程序将交易提交到区块链网络，并且在验证和提交交易后，应用程序会收到有关交易成功的通知。在幕后，这涉及共识过程，通过该过程，区块链网络的不同组件将共同努力，以确保对账本的每个提案更新均有效并以一致且一致的顺序执行。

![img](https:////upload-images.jianshu.io/upload_images/6280489-93601fd4ae9b9057.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

在上方，你可以查看使得此过程正常运行的主要组件。除了每个托管一个帐本副本以及可选的智能合约副本的多个对端节点之外，网络还包含一个交易排序服务。交易排序服务协调网络的交易；它以明确定义的顺序创建包含交易数据的区块，这些交易源自连接到网络的所有不同应用程序。

我们对帐本的第一次更新将创建一辆新车。我们有一个单独的程序称为 `invoke.js`，将用于更新帐本。与查询一样，使用编辑器打开程序并导航到代码块，我们在其中构建交易并将其提交给网络：



```csharp
await contract.submitTransaction('createCar', 'CAR12', 'Honda', 'Accord', 'Black', 'Tom');
```

了解应用程序如何调用智能合约交易 `createCar` 与名为 `Tom` 的所有者创建 `black Honda Accord`。我们在这里使用 `CAR12` 作为识别键 (identifying key)，只是为了表明我们不需要使用顺序键 (sequential keys)。

保存并运行程序：



```ruby
$ cd /path/to/fabric-samples/fabcar/javascript
$ node invoke.js
```

如果调用成功，你将看到以下输出：



```bash
Wallet path: ...fabric-samples/fabcar/javascript/wallet
2018-12-11T14:11:40.935Z - info: [TransactionEventHandler]: _strategySuccess: strategy success for transaction "9076cd4279a71ecf99665aed0ed3590a25bba040fa6b4dd6d010f42bb26ff5d1"
Transaction has been submitted
```

请注意，`invoke` 应用程序如何使用 `submitTransaction` API 而不是 `evaluateTransaction` 与区块链网络进行交互。



```csharp
await contract.submitTransaction('createCar', 'CAR12', 'Honda', 'Accord', 'Black', 'Tom');
```

`submitTransaction` 比 `evaluateTransaction` 更复杂。 SDK 不会与单个对端节点进行交互，而是将 `submitTransaction` 提案发送到区块链网络中组织需要的每个对端节点。这些对端节点中的每一个都将使用该提案执行所请求的智能合约，以生成交易响应，并签名并返回给 SDK。 SDK 将所有已签名的交易响应收集到一个交易中，然后将其发送给交易排序服务器。交易排序服务器将来自每个应用程序的交易收集并排序成一个交易区块。然后，将这些区块分配给网络中的每个对端节点，在此对每个交易进行验证和提交。最后，通知 SDK，使其可以将控制权返回给应用程序。

###### 注解

> `submitTransaction` 还包括一个监听器，该监听器进行检查以确保交易已通过验证并已提交到账本。应用程序应该利用提交易监听器，或者利用诸如 `submitTransaction` 之类的 API 来完成此任务。否则，你的交易可能无法成功排序，验证并提交到账本。

`submitTransaction` 为应用程序完成所有这些工作！应用程序，智能合约，对端节点和交易排序服务一起工作以使帐本在整个网络中保持一致的过程称为共识，本节对此进行了详细说明。

要查看此交易已被写入帐本，请返回 `query.js` 并将参数从 `CAR4` 更改为 `CAR12`。

换句话说，更改下面代码：



```csharp
const result = await contract.evaluateTransaction('queryCar', 'CAR4');
```

到



```csharp
const result = await contract.evaluateTransaction('queryCar', 'CAR12');
```

再次保存，然后查询：



```ruby
$ cd /path/to/fabric-samples/fabcar/javascript
$ node query.js
```

将返回以下内容：



```csharp
Wallet path: ...fabric-samples/fabcar/javascript/wallet
Transaction has been evaluated, result is:
{"colour":"Black","make":"Honda","model":"Accord","owner":"Tom"}
```

恭喜你！你已经创建了一辆汽车，并验证了它已记录在帐本中！

因此，既然我们已经做到了，那么就说 Tom 很慷慨，他想将 Honda Accord 送给一个叫 Dave 的人。

为此，请返回 `invoke.js` 并将智能合约交易从 `createCar` 更改为 `changeCarOwner`，并在输入参数中进行相应的更改：



```csharp
await contract.submitTransaction('changeCarOwner', 'CAR12', 'Dave');
```

第一个参数 `CAR12` 标识将要更改所有者的汽车。第二个参数 `Dave` 定义了汽车的新所有者。

保存并再次执行程序：



```ruby
$ cd /path/to/fabric-samples/fabcar/javascript
$ node invoke.js
```

现在，让我们再次查询帐本，并确保 Dave 现在已与 CAR12 键相关联：



```ruby
$ cd /path/to/fabric-samples/fabcar/javascript
$ node query.js
```

它应该返回以下结果：



```csharp
Wallet path: ...fabric-samples/fabcar/javascript/wallet
Transaction has been evaluated, result is:
{"colour":"Black","make":"Honda","model":"Accord","owner":"Dave"}
```

`CAR12` 的所有权已从 `Tom` 更改为 `Dave`。

###### 注解

> 在实际应用中，智能合约可能会具有一些访问控制逻辑。例如，仅某些授权用户可以创建新车，而只有车主可以将车转让给其他人。

## 7. 总结

现在，我们已经完成了一些查询和一些更新，你应该对应用程序如何使用智能合约查询或更新账本与区块链网络进行交互有了很好的了解。你已经了解了智能合约，API 和 SDK 在查询和更新中扮演的角色的基本知识，并且应该对如何使用各种类型的应用程序执行其他业务任务和操作有所了解。

## 8. 额外资源

正如我们在简介中所述，我们有一整节关于 [Developing Applications](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fdevelopapps%2Fdeveloping_applications.html)，其中包括有关智能合约，流程和数据设计的深入信息，使用更深入的 [Commercial Paper](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Ftutorial%2Fcommercial_paper.html) 的教程以及与应用程序的开发相关的大量其他材料。

## Reference

- Docs » Tutorials » Writing Your First Application, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/write_first_app.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fwrite_first_app.html)
- Docs » Key Concepts, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/key_concepts.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fkey_concepts.html)
- Docs » Developing Applications, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/developapps/developing_applications.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fdevelopapps%2Fdeveloping_applications.html)
- Docs » Tutorials » Commercial paper tutorial, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/tutorial/commercial_paper.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Ftutorial%2Fcommercial_paper.html)
- Docs » Developing Applications » Application, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/developapps/application.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fdevelopapps%2Fapplication.html)
- Docs » Developing Applications » Smart Contract Processing, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/developapps/smartcontract.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fdevelopapps%2Fsmartcontract.html)
- Docs » Architecture Reference » Service Discovery, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/discovery-overview.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fdiscovery-overview.html)
- Docs » Key Concepts » Private data, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/private-data/private-data.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fprivate-data%2Fprivate-data.html)
- Docs » Tutorials » Building Your First Network, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/build_network.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fbuild_network.html)
- Docs » Getting Started » Prerequisites, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/prereqs.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fprereqs.html)
- Docs » Getting Started » Install Samples, Binaries and Docker Images, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/install.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Finstall.html)
- Docs » Developing Applications » Application design elements » Connection Profile, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/developapps/connectionprofile.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fdevelopapps%2Fconnectionprofile.html)
- Hyperledger Fabric Node.js Contract and Shim, [https://fabric-shim.github.io/master/index.html?redirect=true](https://links.jianshu.com/go?to=https%3A%2F%2Ffabric-shim.github.io%2Fmaster%2Findex.html%3Fredirect%3Dtrue)

## 项目源代码

项目源代码会逐步上传到 Github，地址为 [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)。

## Contributor

1. Windstamp, [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)



作者：furnace
链接：https://www.jianshu.com/p/b65c582beb9d
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。



# Hyperledger Fabric 专题 - Hyperledger Fabric Model

本专题文档概述了 Hyperledger Fabric 的关键设计特征，这些特征可以实现 Fabric 对全面但可自定义的企业区块链解决方案的承诺：

- 资产 (Asset) - 资产定义允许通过网络交换几乎所有具有货币价值的东西，从完整食品到古董车再到货币期货。
- 链码 (Chaincode) - 链码执行从交易排序中划分出来，限制了节点类型之间的信任和验证级别，并优化了网络可伸缩性和性能。
- 帐本特征 - 不变的共享帐本对每个通道的整个交易历史进行编码，并包括类似 SQL 的查询功能，以进行有效的审计和争议解决。
- 隐私 (Privacy) - 通道和私人数据集可实现私有和机密的多方交易，这通常是竞争企业和受管制的行业在同一网络上交换资产时需要的。
- 安全和 MSP (Security & Membership Services) - 会员许可制提供了一个受信任的区块链网络，参与者知道所有交易都可以由授权的监管机构和审计员检测和追踪。
- 共识 (Consensus) - 达成共识的独特方法可实现企业所需的灵活性和可扩展性。

## 1. 资产

资产的范围从有形的 (房地产和硬件) 到无形的 (合同和知识产权)。 Hyperledger Fabric 提供了使用链码交易修改资产的功能。

资产在 Hyperledger Fabric 中表示为键值对的集合，状态更改记录为通道账本中的交易。资产可以二进制和/或 JSON 形式表示。

## 2. 链码

链码是定义一项或多项资产的软件，以及用于修改资产的交易指令。换句话说，这就是业务逻辑。链码定义了用于读取或更改键值对或其他状态数据库信息的规则。链码函数针对帐本的当前状态数据库执行，并通过交易提案启动。链码执行会产生一组键值写操作 (写集)，这些键值写操作可以提交给网络，并应用于所有对端节点的帐本中。

## 3. 账本功能

账本是 Fabric 中所有状态转换的有序和防篡改的记录。状态转换是参与方提交的链码调用 (“交易”) 的结果。每笔交易都会产生一组资产键值对，这些键值对在创建，更新或删除时将被提交到账本。

帐本由一个区块链 (“链”) 和一个状态数据库组成，该区块链将不可变的顺序记录存储在区块中，并维护一个状态数据库。每个通道有一个帐本。每个对端节点都为其所属的每个通道维护一个帐本的副本。

Fabric 帐本的一些功能：

- 使用基于键的查询，范围查询和组合键查询来查询和更新帐本。
- 使用富查询语言的只读查询 (如果使用 CouchDB 作为状态数据库)。
- 只读历史记录查询 — 查询键的帐本历史记录，从而启用数据起源场景。
- 交易版本包括链码读取的键/值 (读集) 和链码写入的键/值  (写集)。
- 交易包含每个背书对端节点的签名，并提交给交易排序服务。
- 交易被分为几大块，并从交易排序服务分发到通道上的对端节点。
- 对端节点根据背书策略验证交易并执行。
- 在添加区块之前，执行版本检查，以确保自链码执行以来，已读取资产的状态未更改。
- 一旦交易被验证并被提交，就不可改变。
- 通道的帐本包含一个配置块，用于定义策略，访问控制列表和其他相关信息。
- 通道包含成员资格服务提供者实例，允许从不同的证书颁发机构派生加密材料。

有关数据库，存储结构和“查询能力”的更深入了解，请参阅 [帐本](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Ffabric_model.html) 专题文档。

## 4. 隐私

Hyperledger Fabric 在每个通道的基础上使用不变的帐本，以及可以操纵和修改资产当前状态 (即更新键值对) 的链码。账本存在于通道范围内 - 可以在整个网络中共享 (假设每个参与者都在一个公共通道上工作) - 或可以将其私有化以仅包括一组特定的参与者。

在后一种情况下，这些参与者将创建一个单独的通道，从而隔离他们的交易和帐本。为了解决想要弥合总体透明度和隐私之间的差距的方案，只能在需要访问资产状态以执行读写的对端节点上安装链码 (换句话说，如果未在对端节点上安装链码, 它将无法与帐本正确连接)。

当该通道上的组织子集需要对其交易数据保密时，可以使用私有数据集合将这些数据隔离在逻辑上与通道帐本分离的私有数据库中，该数据库只能由组织的授权子集访问。

因此，通道使交易对于更广泛的网络而言是不公开的，而集合则对通道上的组织子集之间的数据保持不公开。

为了进一步模糊数据，可以在将交易发送到交易排序服务并将区块添加到账本之前，使用 AES 等通用加密算法对链码中的值进行加密 (部分或全部)。一旦加密数据已写入帐本，则只有拥有用于生成密文的相应密钥的用户才能对其解密。有关链码加密的更多详细信息，请参阅 [链码开发人员](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fchaincode4ade.html) 主题。

有关如何在区块链网络上实现隐私的更多详细信息，请参见 [私有数据](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fprivate-data-arch.html) 主题。

## 5. 安全和 MSP

Hyperledger Fabric 支持所有参与者都具有已知身份的交易网络。公钥基础结构 (Public Key Infrastructure, PKI) 用于生成与组织，网络组件以及最终用户或客户端应用程序绑定的加密证书。结果，可以在更广泛的网络和通道级别上操纵和控制数据访问控制。Hyperledger Fabrc 的这种“许可制”的概念，再加上通道的存在和能力，帮助解决那些需要隐私和保密性要求非常高的情况。

查看 [会员服务提供商 ( Membership Service Providers , MSP) 主题](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fmsp.html)，以更好地了解加密实现，并签名，验证，在 Hyperledger Fabric 用于身份验证的方法。

## 6. 共识

在分布式帐本技术中，共识最近已成为单一功能内特定算法的同义词。但是，共识不仅包括简单地同意交易顺序，而且这种差异在 Hyperledger Fabric 中得到了体现，它在整个交易流程 (从提案和背书到交易排序，验证和承诺) 中的基本作用得到了强调。简而言之，共识被定义为对包含一个区块的一组交易的正确性的全面验证。

区块交易的顺序和结果满足明确的政策标准检查后，才能最终达成共识。这些检查发生在交易的生命周期中，包括使用背书策略来指示哪些特定成员必须背书某个交易类，以及系统链码以确保这些策略得到执行和维护。在作出承诺之前，对端节点将使用这些系统链码来确保存在足够的背书，并且它们是从适当的实体派生的。此外，在将包含交易的任何区块追加到账本之前，将进行版本控制检查，在此期间将对账本的当前状态达成一致。最终检查可以防止重复使用操作和其他可能危害数据完整性的威胁，并可以针对非静态变量执行功能。

除了进行大量的背书，有效性和版本检查外，还在交易流程的所有方向上进行持续的身份验证。访问控制列表是在网络的分层层上实现的 (将交易排序服务分布到各个通道)，有效载荷在交易提案通过不同的体系结构组件时被重复签名，验证和认证。总而言之，共识不仅限于一批交易的商定顺序。相反，它是一项总体特征，它是交易从提案到承诺过程中不断进行的验证的副产品。

查看 [交易流程图](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Ftxflow.html) 以直观表示共识。

## Reference

- Docs » Key Concepts » Hyperledger Fabric Model, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/fabric_model.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Ffabric_model.html)
- Docs » Ledger, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/ledger.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fledger.html)
- Docs » Tutorials » Chaincode for Developers, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/chaincode4ade.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fchaincode4ade.html)
- Docs » Architecture Reference » Private Data, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/private-data-arch.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fprivate-data-arch.html)
- Docs » Operations Guides » Membership Service Providers (MSP), [https://hyperledger-fabric.readthedocs.io/en/release-1.4/msp.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fmsp.html)
- Docs » Architecture Reference » Transaction Flow, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/txflow.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Ftxflow.html)

## 项目源代码

项目源代码会逐步上传到 Github，地址为 [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)。

## Contributor

1. Windstamp, [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)



作者：furnace
链接：https://www.jianshu.com/p/f78ef5730816
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。



# 2.1 Hyperledger Fabric - 核心概念 - 智能合约和链码

# 核心概念 - 智能合约和链码

从应用程序开发人员的角度来看，智能合约与帐本一起构成了 Hyperledger Fabric 区块链系统的核心。账本保存有关一组业务对象的当前和历史状态的事实，而智能合约定义可执行逻辑，生成添加到账本的新事实。管理员通常使用链码对相关的智能合约进行分组以进行部署，但也可以将其用于 Fabric 的低级系统编程。在本主题中，我们将重点介绍为什么存在智能合约和链码以及如何以及何时使用它们。

在本主题中，我们将介绍：

- 什么是智能合约
- 术语说明
- 智能合约和帐三个
- 如何开发智能合约
- 背书策略的重要性
- 有效交易
- 智能合约和通道
- 智能合约之间的通信
- 什么是系统链码？

## 1. 智能合约

在业务彼此之间进行交易之前，他们必须定义一组通用合约，涵盖通用条款，数据，规则，概念定义和流程。这些合约放在一起，构成了控制交易双方之间所有交互的业务模型。

![img](https:////upload-images.jianshu.io/upload_images/6280489-273e516a091881cf.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

智能合约以可执行代码定义了不同组织之间的规则。应用程序调用智能合约以生成记录在账本上的交易。

使用区块链网络，我们可以将这些合约转换为可执行程序 (在业界称为智能合约)，以开辟各种新的可能性。这是因为智能合约可以为任何类型的业务对象实施治理规则，以便在执行智能合约时可以自动执行这些规则。例如，智能合约可以确保在指定的时间范围内完成新车交付，或者确保按照预先安排的条款释放资金，从而分别改善了货物或资金的流动。但是，最重要的是，智能合约的执行比人工业务流程要高效得多。

在上图中，我们可以看到 ORG1 和 ORG2 这两个组织如何定义了汽车智能合约来查询，转让和更新汽车。这些组织的应用程序调用此智能合约在业务流程中执行约定的步骤，例如将特定汽车的所有权从 ORG1 转移到 ORG2。

## 2. 术语

Hyperledger Fabric 用户经常互换使用术语智能合约和链码。通常，智能合约定义了控制世界状态中包含的业务对象生命周期的交易逻辑。然后将其打包成链码，然后将其部署到区块链网络。可以将智能合约视为管理交易，而链码则可以控制如何打包智能合约以进行部署。

![img](https:////upload-images.jianshu.io/upload_images/6280489-7391e6c5d46bbcfa.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

智能合约在链码中定义。可以在同一链码中定义多个智能合约。部署链码后，其中的所有智能合约都可用于应用程序。

在上图中，我们可以看到一个包含三个智能合约的车辆 (`vehicle`) 链码：汽车 (`cars`)，船只 (`boats`) 和卡车 (`trucks`)。我们还可以看到包含四个智能合约的保险 (`insurance`) 链码：保单 (`policy`)，责任 (`liablility`)，银团 (`syndication`) 和证券化 (`securitization`)。在这两种情况下，这些合同均涉及与车辆和保险有关的业务流程的关键方面。在本主题中，我们将以汽车 (`car`) 合约为例。我们可以看到，智能合约是与特定业务流程相关的特定于域的程序，而链码是一组用于安装和实例化的相关智能合约的技术容器。

## 3. 账本

在最简单的层次上，区块链一成不变地记录更新账本中状态的交易。智能合约以编程方式访问账本的两个不同部分：一个区块链，它不变地记录所有交易的历史记录；一个世界状态，保存着这些状态的当前值缓存，因为它是对象通常需要的当前值。

智能合约主要在世界状态下放置 (put)，获取 (get) 和删除 (delete) 状态，并且还可以查询不可变的区块链交易记录。

- 获取 (get) 通常代表查询以检索有关业务对象当前状态的信息。
- 放置 (put) 通常会创建新的业务对象或修改帐本世界状态中的现有业务对象。
- 删除 (delete) 通常表示从帐本的当前状态中删除业务对象，但不删除其历史记录。

智能合约有许多可用的 [API](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fdevelopapps%2Ftransactioncontext.html%23structure)。至关重要的是，在所有情况下，无论交易是在世界状态下创建，读取，更新还是删除业务对象，区块链都包含这些更改的 [不可变记录](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fledger%2Fledger.html)。

## 4. 开发

智能合约是应用程序开发的重点，正如我们已经看到的，可以在单个链码中定义一个或多个智能合约。将链码部署到网络后，该网络中的所有组织均可使用其所有智能合约。这意味着只有管理员才需要担心链码。其他人都可以根据智能合约进行思考。

智能合约的核心是一组交易定义。例如，查看 `fabcar.js`，你可以在其中看到创建新车的智能合约交易：



```csharp
async createCar(ctx, carNumber, make, model, color, owner) {

    const car = {
        color,
        docType: 'car',
        make,
        model,
        owner,
    };

    await ctx.stub.putState(carNumber, Buffer.from(JSON.stringify(car)));
}
```

你可以在 [编写第一个应用程序教程](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fwrite_first_app.html) 中了解有关 Fabcar 智能合约的更多信息。

智能合约可以描述与多组织决策中的数据不变性相关的几乎无限的业务用例。智能合约开发人员的工作是采用可能控制金融价格或交付条件的现有业务流程，并以诸如 JavaScript，GOLANG 或 Java 的编程语言将其表示为智能合约。智能合约审核员越来越多地练习将数百年的法律语言转换为编程语言所需的法律和技术技能。你可以在 [开发应用程序主题](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fdevelopapps%2Fdeveloping_applications.html) 中了解如何设计和开发智能合约。

## 5. 背书

与每个链码相关联的是一种背书策略，该策略适用于其中定义的所有智能合约。背书非常重要；它指示区块链网络中的哪些组织必须签署由给定智能合约生成的交易，才能将该交易宣布为有效。

![img](https:////upload-images.jianshu.io/upload_images/6280489-c812227119683f77.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

每个智能合约都有与其相关的背书策略。该背书策略可确定哪些组织必须批准智能合约生成的交易，然后才能将其确认为有效交易。

示例性背书策略可能定义了参与区块链网络的四个组织中的三个必须在一笔交易被视为有效之前对其进行签名。所有有效或无效交易都将添加到分布式帐本中，但只有有效交易会更新世界状态。

如果背书策略指定必须由多个组织签署交易，则必须由足够多的组织来执行智能合约，以便生成有效的交易。在上面的示例中，转让汽车的智能合约交易需要由 `ORG1` 和 `ORG2` 执行并签名，以使其有效。

背书策略使 Hyperledger Fabric 与以太坊或比特币等其他区块链有所不同。在这些系统中，网络中的任何节点都可以生成有效的交易。Hyperledger Fabric 更现实地模拟了现实世界；交易必须由网络中的受信任组织验证。例如，管理组织必须签署有效的 `issueIdentity` 交易，或者汽车的买 (`buyer`) 卖 (`seller`) 双方都必须签署汽车 (`car`) 转让交易。背书策略旨在使 Hyperledger Fabric 能够更好地对这些类型的实际交互进行建模。

最后，背书策略只是 Hyperledger Fabric 中策略的一个示例。可以定义其他策略以标识谁可以查询或更新帐本，或从网络中添加或删除参与者。一般而言，尽管不是一成不变的，但策略应由区块链网络中的联盟组织事先达成协议。实际上，策略本身可以定义可以更改策略的规则。尽管是高级主题，但也可以在 Fabric 提供的规则之外定义 [自定义背书策略](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fpluggable_endorsement_and_validation.html) 规则。

## 6. 有效交易

当智能合约执行时，它在区块链网络中的一个组织拥有的对端节点上运行。合约采用一组称为交易提案的输入参数，并将其与程序逻辑结合使用以读取和写入帐本。对世界状态的更改被捕获为交易提案响应 (或仅仅是交易响应)，其中包含一个读写集 (read-write set)，其中包含已读取的状态以及如果交易有效将要写入的新状态。请注意，执行智能合约时世界状态不会更新！

[图片上传失败...(image-4ff111-1575546689004)]

所有交易都具有由一组组织签名的标识符，提案和响应。所有交易均记录在区块链上，无论其有效与否，但只有有效交易才能更新世界状态。

检查 `car transfer` 交易。你可以看到在 ORG1 和 ORG2 之间进行汽车转移的交易 t3。查看交易如何输入 {CAR1，ORG1，ORG2} 并输出 {CAR1.owner = ORG1，CAR1.owner = ORG2}，表示所有者从 ORG1 变为 ORG2。请注意，输入是由应用程序的组织 ORG1 签名的，输出是由背书策略 ORG1 和 ORG2 标识的两个组织签名的。这些签名是通过使用每个参与者的私钥生成的，意味着网络中的任何人都可以验证网络中的所有参与者是否都同意交易细节。

分两个阶段验证分发给网络中所有对端节点的交易。首先，检查交易，以确保有足够的组织根据背书策略签署了该交易。其次，检查以确保世界状态的当前值与由对端节点签名的交易的读取集相匹配；没有中间更新。如果交易通过了这两个测试，则将其标记为有效。所有交易都会添加到区块链历史记录中，无论有效与否，但只有有效交易才会导致世界状态的更新。

在我们的示例中，t3 是有效交易，因此 CAR1 的所有者已更新为 ORG2。但是，t4 (未显示) 是无效交易，因此，尽管将其记录在帐本中，但世界状态并未更新，CAR2 仍归 ORG2 所有。

最后，要了解如何在世界状态下使用智能合约或链码，请阅读 [链码命名空间主题](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fdevelopapps%2Fchaincodenamespace.html)。

## 7. 通道

Hyperledger Fabric 允许组织通过通道同时参与多个独立的区块链网络。通过加入多个通道，组织可以参与所谓网络的网络。通道可有效共享基础架构，同时保持数据和通信的隐私。它们足够独立，可以帮助组织与不同的交易对手分离其工作量，但又足够集成，可以在必要时协调独立的活动。

![img](https:////upload-images.jianshu.io/upload_images/6280489-04d896f983878286.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

通道在一组组织之间提供了完全独立的通信机制。在通道上实例化链码时，将为其定义背书策略。链码中的所有智能合约均可供该通道上的应用程序使用。

管理员在通道上实例化链码时为其定义了背书策略，并且可以在链码升级时对其进行更改。背书策略同样适用于在部署到通道的同一链码中定义的所有智能合约。这也意味着可以将单个智能合约部署到具有不同背书策略的不同通道。

在上面的示例中，汽车合约已部署到 VEHICLE 通道，保险合约已部署到 INSURANCE 通道。汽车合约的背书策略要求 ORG1 和 ORG2 在被视为有效之前签署交易，而保险合约的背书策略仅要求 ORG3 签署有效交易。 ORG1 参与两个网络，即 VEHICLE 通道和 INSURANCE 网络，并且可以分别通过 ORG2 和 ORG3 协调这两个网络之间的活动。

## 8. 互通

智能合约可以在同一通道内和不同通道之间调用其他智能合约。这样，他们可以读取和写入由于智能合约命名空间而无法访问的世界状态数据。

合约间通信存在局限性，在 [链码命名空间](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fdevelopapps%2Fchaincodenamespace.html%23cross-chaincode-access) 主题中已对此进行了全面介绍。

## 9. 系统链码

链码中定义的智能合约对一组区块链组织之间达成共识的业务流程的域相关规则进行编码。但是，链码还可以定义与领域无关的系统交互相对应的低级程序代码，而这些交互与业务流程的智能合约无关。

以下是不同类型的系统链码及其相关的缩写：

- 生命周期系统链码 (Lifecycle system chaincode, LSCC) 在所有对端节点中运行，以处理程序包签名，安装，实例化和升级链码请求。你可以阅读有关 LSCC 实现此 [过程](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fchaincode4noah.html%23chaincode-lifecycle) 的更多信息。
- 配置系统链码 (Configuration system chaincode, CSCC) 在所有对端节点中运行，以处理对通道配置的更改，例如策略更新。你可以在以下链码 [主题](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fconfigtx.html%23configuration-updates) 中阅读有关此过程的更多信息。
- 查询系统链码 (Query system chaincode, QSCC) 在所有对端节点运行提供账本的 API，其中包括区块查询，交易查询等，你可以在交易方面的 [话题](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fdevelopapps%2Ftransactioncontext.html) 了解更多有关这些帐本 API。
- 背书系统链码 (Endorsement system chaincode, ESCC) 在背书对端节点以加密方式签署交易响应时运行。你可以阅读有关 ESCC 如何实施此 [过程](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fpeers%2Fpeers.html%23phase-1-proposal) 的更多信息。
- 验证系统链码 (Validation system chaincode, VSCC) 验证交易，包括检查背书策略和读写集版本控制。你可以阅读有关 LSCC 实现此 [过程](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fpeers%2Fpeers.html%23phase-3-validation) 的更多信息。

底层 Fabric 开发人员和管理员可以修改这些系统链码以供自己使用。但是，系统链码的开发和管理是一项专门的活动，与智能合约的开发完全分开，通常不需要。对系统链码的更改必须格外小心，因为它们对于 Hyperledger Fabric 网络的正常运行至关重要。例如，如果未正确开发系统链码，则一个对端节点可能会与另一对端节点不同地更新其世界状态或区块链的副本。缺乏共识是账本分叉的一种形式，这是非常不理想的情况。

## Reference

- Docs » Key Concepts » Smart Contracts and Chaincode, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/smartcontract/smartcontract.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fsmartcontract%2Fsmartcontract.html)
- Docs » Developing Applications » Application design elements » Transaction context, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/developapps/transactioncontext.html#structure](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fdevelopapps%2Ftransactioncontext.html%23structure)
- Docs » Key Concepts » Ledger, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/ledger/ledger.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fledger%2Fledger.html)
- Docs » Tutorials » Writing Your First Application, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/write_first_app.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fwrite_first_app.html)
- Docs » Developing Applications, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/developapps/developing_applications.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fdevelopapps%2Fdeveloping_applications.html)
- Docs » Operations Guides » Pluggable transaction endorsement and validation, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/pluggable_endorsement_and_validation.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fpluggable_endorsement_and_validation.html)
- Docs » Developing Applications » Application design elements » Chaincode namespace, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/developapps/chaincodenamespace.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fdevelopapps%2Fchaincodenamespace.html)
- Docs » Developing Applications » Application design elements » Chaincode namespace, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/developapps/chaincodenamespace.html#cross-chaincode-access](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fdevelopapps%2Fchaincodenamespace.html%23cross-chaincode-access)

## 项目源代码

项目源代码会逐步上传到 Github，地址为 [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)。

## Contributor

1. Windstamp, [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)



作者：furnace
链接：https://www.jianshu.com/p/88a8575b96df
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。



# 2.2 Hyperledger Fabric - 核心概念 - 账本

# 核心概念 - 账本 (Ledger)

帐本是 Hyperledger Fabric 中的关键概念。它存储有关业务对象的重要事实信息；对象属性的当前值，以及导致这些当前值的交易历史。

在本主题中，我们将介绍：

- 什么是帐本？
- 存储有关业务对象的事实
- 区块链帐本
- 世界状态
- 区块链数据结构
- 区块如何存储在区块链中
- 交易
- 世界状态数据库选项
- Fabcar 示例帐本
- 帐本和命名空间
- 帐本和通道

## 1. 什么是帐本？

帐本包含业务的当前状态作为交易日志。最早的欧洲和中国帐本可追溯到大约 1000 年前，而苏美尔人则在 4000 年前就有石头帐本 - 但让我们从一个更新的例子开始吧！

你可能习惯于查看自己的银行帐户。对你而言，最重要的是可用余额 - 它是你当前及时的支出能力。如果要查看余额的产生方式，则可以查看确定余额的交易贷方和借方。这是一个账本的活生生的例子 - 一个状态 (你的银行存款余额)，以及一组确定其状态的有序交易 (贷方和借方)。 Hyperledger Fabric 是由这两个相同的组件构成 - 提供一套账本状态的当前值，并捕获确定这些状态的交易历史。

## 2. 帐本，事实和状态

帐本实际上不存储业务对象，而是存储有关这些对象的事实。当我们说“我们将业务对象存储在帐本中”时，我们的真正意思是我们正在记录有关对象当前状态的事实以及导致当前状态的交易历史的事实。在日益数字化的世界中，感觉就像我们在看一个对象，而不是关于对象的事实。对于数字对象，它可能存在于外部数据存储中；我们存储在帐本中的事实使我们能够识别其位置以及有关他的其他关键信息。

尽管有关业务对象当前状态的事实可能会发生变化，但是有关它的事实历史是不可变的，可以对其进行添加，但不能进行追溯更改。我们将看到如何将区块链视为关于业务对象事实的不变历史，这是一种简单而有效的理解方式。

现在，让我们仔细看一下 Hyperledger Fabric 帐本结构！

## 3. 帐本

在 Hyperledger Fabric 中，账本由两个不同但相关的部分组成 - 世界状态和区块链。这些中的每一个都代表有关一组业务对象的一组事实。

首先，有一个世界状态 - 一个数据库，其中存储了一组帐本状态的当前值的缓存。世界状态使程序可以轻松地直接访问状态的当前值，而不必通过遍历整个交易日志来计算状态值。默认情况下，帐本状态表示为键值对，稍后我们将看到 Hyperledger Fabric 如何在这方面提供灵活性。世界状态可以频繁更改，因为可以创建，更新和删除状态。

其次，有一个区块链 – 交易日志，记录了导致当前世界状态的所有更改。交易收集在附加到区块链的区块内部，使你能够了解导致当前世界状态变化的历史。区块链数据结构与世界状态非常不同，因为一旦写入，就无法修改；这是一成不变的。

![img](https:////upload-images.jianshu.io/upload_images/6280489-1db52a59ba8a71c9.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

账本 L 由区块链 B 和世界状态 W 组成，其中区块链 B 确定世界状态 W。我们也可以说世界状态 W 源自区块链 B。

最好在 Hyperledger Fabric 网络中考虑一个逻辑帐本。实际上，网络维护一个帐本的多个副本 – 通过称为共识的过程，它们与其他所有副本保持一致。术语“分布式账本技术” (Distributed Ledger Technology, DLT) 通常与这种帐本相关联 - 这种帐本在逻辑上是单一的，但是在整个网络中分布有许多一致的副本。

现在，让我们更详细地研究世界状态和区块链数据结构。

## 4. 世界状态

世界状态将业务对象属性的当前值保留为唯一的帐本状态。这很有用，因为程序通常需要对象的当前值；遍历整个区块链来计算对象的当前值将很麻烦 - 你只需直接从世界状态中获取即可。

![img](https:////upload-images.jianshu.io/upload_images/6280489-ca510344d507b6e0.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

包含两个状态的帐本世界状态。第一种状态是：key = CAR1 和 value = Audi。第二个状态具有更复杂的值：key = CAR2 和 value = {model：BMW，color = red，owner = Jane}。两种状态的版本均为 0。

帐本状态记录有关特定业务对象的一组事实。我们的示例显示了两个汽车 CAR1 和 CAR2 的帐本状态，每个汽车都有一个键和一个值。应用程序可以调用智能合约，该合约使用简单的帐本 API 来获取，放置和删除状态。注意状态值可以简单 (Audi…) 或复杂 (类型：BMW…)。通常会查询世界状态来检索具有某些属性的对象，例如查找所有红色宝马。

世界状态被实现为数据库。这很有意义，因为数据库为状态的有效存储和检索提供了一组丰富的运算符。稍后我们将看到 Hyperledger Fabric 可以配置为使用不同的世界状态数据库来满足不同类型的状态值和应用程序所需的访问模式的需求，例如在复杂的查询中。

应用程序提交捕获世界状态变化的交易，这些交易最终被提交到账本区块链。 Hyperledger Fabric SDK 将应用程序与该共识机制的细节隔离开来。它们仅调用智能合约，并在交易已包含在区块链中时 (无论有效还是无效) 得到通知。关键设计要点是，只有由所需的一组背书组织签名的交易才会导致对世界状态的更新。如果交易没有得到足够的背书人的签名，则不会导致世界状态的改变。你可以阅读有关应用程序如何使用 [智能合约](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fsmartcontract%2Fsmartcontract.html) 以及如何 [开发应用程序](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fdevelopapps%2Fdeveloping_applications.html) 的更多信息。

你还会注意到，一个状态有一个版本号，在上图中，状态 CAR1 和 CAR2 的起始版本为 0。Hyperledger Fabric 内部使用的版本号，每次状态更改时都会递增。每当状态更新时都会检查版本，以确保当前状态与背书时的版本匹配。这确保了世界状态正在按预期变化；没有并发更新。

最后，当首次创建帐本时，世界状态为空。因为任何代表世界状态有效改变的交易都记录在区块链上，这意味着可以随时从区块链重新生成世界状态。这可能非常方便 – 例如，创建对端节点时会自动生成世界状态。此外，如果对端节点异常失败，则可以在接受交易之前在对端节点重新启动时重新生成世界状态。

## 5. 区块链

现在，让我们将注意力从世界状态转移到区块链上。世界状态包含与一组业务对象的当前状态有关的一组事实，而区块链是有关这些对象如何到达其当前状态的事实的历史记录。区块链记录了每个帐本状态的每个先前版本以及更改方式。

区块链被构造为互连区块的顺序日志，其中每个区块包含一系列交易，每个交易代表对世界状态的查询或更新。交易排序的确切机制在 [其他地方](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fpeers%2Fpeers.html%23peers-and-orderers) 讨论；重要的是，在首次由称为交易排序服务的 Hyperledger Fabric 组件创建区块时，便要建立区块排序以及区块内的交易排序。

每个区块头均包含该区块交易的哈希值，以及前一个区块头的哈希值。这样，账本上的所有交易都被排序并通过密码链接在一起。这种哈希和链接使帐本数据非常安全。即使托管账本的一个节点被篡改，它也无法说服所有其他节点其具有“正确的”区块链，因为账本分布在整个网络中的独立节点。

与使用数据库的世界状态相反，区块链始终被实现为文件。这是一个明智的设计选择，因为区块链数据结构严重偏向极少量的简单操作。追加到区块链的末尾是主要操作，而查询在当前是相对不频繁的操作。

让我们更详细地了解区块链的结构。

[图片上传失败...(image-5a8c3f-1575546721769)]

包含区块 B0，B1，B2，B3 的区块链 B。B0 是区块链中的第一个区块，即创世区块。

在上图中，我们可以看到区块 B2 的数据 D2 包含其所有交易：T5，T6，T7。

最重要的是，B2 具有一个区块头 H2，它包含 D2 中所有交易的加密哈希以及先前区块 B1 中的等效哈希。通过这种方式，区块彼此之间有着千丝万缕的联系，这就是术语“区块链”如此巧妙地捕捉到的！

最后，如你在图中看到的那样，区块链中的第一个区块称为创世块。尽管它不包含任何用户交易，但它是帐本的起点。相反，它包含一个配置交易，该交易包含网络通道的初始状态 (未显示)。当我们在文档中讨论区块链网络和 [通道](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fchannels.html) 时，我们将更详细地讨论创世区块。

## 6. 区块

让我们仔细看一下区块的结构。它包括三个部分。

### 6.1 区块头

本节包含三个字段，在创建区块时写入。

- 区块编号 (Block number)：一个整数，从 0 (创世区块) 开始，并随添加到区块链的每个新区块增加 1。
- 当前区块哈希 (Current Block Hash)：当前区块中包含的所有交易的哈希。
- 上一个区块哈希 (Previous Block Hash)：来自区块链上一个区块的哈希。

这些字段是通过对区块数据进行密码哈希处理而在内部派生的。他们确保每个区块都与其邻居密不可分，从而导致账本不可改变。

![img](https:////upload-images.jianshu.io/upload_images/6280489-143d8b6ff6e540be.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

区块头详细信息。区块 B2 的头部 H2 由区块编号 2，当前区块数据 D2 的哈希 CH2，以及来自前一个区块，区块编号 1 的哈希 PH1 组成。

### 6.2 区块数据

本节包含按顺序排列的交易列表。在交易排序服务创建区块时写入。这些交易具有丰富但直接的结构，我们将在本主题的后面部分进行介绍。

### 6.3 区块元数据

此部分包含写入区块的时间，以及区块写入者的证书，公钥和签名。随后，区块提交者还为每个交易添加有效/无效指示符，然而此信息不包括在哈希中，因为创建区块时会创建该信息。

## 7. 交易

如我们所见，交易捕获了世界状态的变化。让我们看一下详细的阿悄块数据结构，该结构在一个区块中包含交易。

![img](https:////upload-images.jianshu.io/upload_images/6280489-a7b30df33920ed40.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

交易明细。区块 B1 的区块数据 D1 中的交易 T4 由交易头 H4，交易签名 S4，交易提案 P4，交易提案响应 R4 和背书列表 E4 组成。

在上面的示例中，我们可以看到以下字段：

- 交易头 (Header)

  由 H4 举例说明的本节捕获有关交易的一些基本元数据，例如，相关链码的名称及其版本。

- 签名 (Signature)

  由 S4 说明的该部分包含由客户端应用程序创建的加密签名。此字段用于检查交易明细是否未被篡改，因为它需要应用程序的私钥来生成它。

- 提案 (Proposal)

  用 P4 表示的该字段对应用程序提供给智能合约的输入参数进行编码，该参数用于创建更新账本的提案。当智能合约运行时，此提案提供一组输入参数，这些输入参数与当前的世界状态一起确定新的世界状态。

- 响应 (Response)

  用 R4 表示的此部分将世界状态的前后值捕获为读写集 (Read Write set, RW-set)。它是智能合约的输出，如果交易成功通过验证，它将被应用于账本以更新世界状态。

- 背书 (Endorsement)

  如图 E4 所示，这是来自每个所需组织的足以满足背书策略的已签名交易响应的列表。你会注意到，虽然交易中仅包含一项交易响应，但其中有多项背书。那是因为每个背书都有效地编码了其组织的特定交易响应 - 意味着无需包含任何与足够背书不匹配的交易响应，因为它将被拒绝为无效交易，并且不会更新世界状态。

总结了交易的主要字段 – 还有其他字段，但是你必须了解这些必不可少的字段，才能对账本数据结构有深入的了解。

## 8. 世界状态数据库选项

世界状态从物理上实现为数据库，以提供简单有效的存储和账本状态检索。如我们所见，帐本状态可以具有简单值或复合值，并且为了适应这种情况，世界状态数据库的实现方式可能会有所不同，从而可以有效地实现这些值。世界状态数据库的选项当前包括 LevelDB 和 CouchDB。

LevelDB 是默认值，当帐本状态是简单的键/值对时尤其适用。 LevelDB 数据库与网络节点紧密位于同一位置 – 嵌入在同一操作系统进程中。

当帐本状态被构造为 JSON 文档时，CouchDB 是一个特别合适的选择，因为 CouchDB 支持丰富的查询和更新业务交易中经常发现的更丰富的数据类型。在实现方面，CouchDB 在单独的操作系统进程中运行，但是对端节点和 CouchDB 实例之间仍然存在 1:1 的关系。所有这些对于智能合约都是看不见的。有关 CouchDB 的更多信息，请参见 [CouchDB as the StateDatabase](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fcouchdb_as_state_database.html)。

在 LevelDB 和 CouchDB 中，我们看到了 Hyperledger Fabric 的一个重要方面 – 它是可插入的。世界状态数据库可以是关系数据存储，图形存储或时间数据库。这为可以有效访问的帐本状态类型提供了极大的灵活性，从而使 Hyperledger Fabric 可以解决许多不同类型的问题。

## 9. 帐本示例：fabcar

当我们结束帐本主题时，让我们看一个示例帐本。如果你运行了 [fabcar 示例应用程序](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fwrite_first_app.html)，那么你已经创建了此帐本。

fabcar 示例应用程序创建了一组 10 辆汽车，每辆汽车具有唯一的标识；不同的颜色 (color)，品牌 (make)，型号 (model) 和所有者 (owner)。创建前四辆汽车后，帐本如下所示。

[图片上传失败...(image-2f4c81-1575546721769)]

账本 L 包含一个世界状态 W 和一个区块链 B。W 包含四个具有键的状态：CAR1，CAR2，CAR3 和 CAR4。 B 包含两个区块 0 和 1。区块 1 包含四个交易：T1，T2，T3，T4。

我们可以看到世界状态包含对应于 CAR0，CAR1，CAR2 和 CAR3 的状态。 CAR0 的值表示它是 Tomomo 目前拥有的蓝色 Toyota Prius，我们可以看到其他汽车的相似状态和值。此外，我们可以看到所有汽车状态均为版本号 0，表明这是它们的起始版本号 - 自创建以来尚未对其进行更新。

我们还可以看到，区块链包含两个区块。区块 0 是创世区块，尽管它不包含任何与汽车相关的交易。但是，区块 1 包含交易 T1，T2，T3，T4，这些交易对应于在世界状态下为 CAR0 到 CAR3 创建初始状态的交易。我们可以看到区块 1 链接到区块 0。

我们没有显示区块或交易中的其他字段，特别是区块头和哈希。如果你对这些信息的确切细节感兴趣，可以在文档的其他地方找到专用的参考主题。它为你提供了整个区块的完整可工作的示例，其中包含详细的交易细节 - 但到目前为止，你已经对 Hyperledger Fabric 帐本有了扎实的概念性理解。做得好！

## 10. 命名空间

即使我们已经将帐本呈现为一个单一的世界状态和单个区块链，但这还是有点过分简化了。实际上，每个链码都有自己的世界状态，该世界状态与所有其他链码是分开的。世界状态位于命名空间中，因此只有相同链码内的智能合约才能访问给定的命名空间。

区块链未命名空间。它包含来自许多不同的智能合约命名空间的交易。你可以在 [本主题](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fledger%2Fdevelopapps%2Fchaincodenamespace.html) 中阅读有关链码命名空间的更多信息。

现在让我们看看如何在 Hyperledger Fabric 通道中应用命名空间的概念。

## 11. 通道

在 Hyperledger Fabric 中，每个 [通道](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fchannels.html) 都有一个完全独立的帐本。这意味着完全独立的区块链，以及完全独立的世界状态，包括命名空间。应用程序和智能合约可以在通道之间进行通信，以便可以在它们之间访问帐本信息。

你可以在 [本主题](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fledger%2Fdevelopapps%2Fchaincodenamespace.html%23channel) 中阅读有关帐本如何与通道一起使用的更多信息。

## 12. 更多信息

请参阅 [交易流](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Ftxflow.html)，[读写集语义](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Freadwrite.html) 和 [CouchDB as the StateDatabase](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fcouchdb_as_state_database.html) 主题，以更深入地了交易务流，并发控制和世界状态数据库。

## Reference

- Docs » Key Concepts » Ledger, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/ledger/ledger.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fledger%2Fledger.html)
- Docs » Key Concepts » Smart Contracts and Chaincode, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/smartcontract/smartcontract.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fsmartcontract%2Fsmartcontract.html)
- Docs » Developing Applications, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/developapps/developing_applications.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fdevelopapps%2Fdeveloping_applications.html)
- Docs » Key Concepts » Peers, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/peers/peers.html#peers-and-orderers](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fpeers%2Fpeers.html%23peers-and-orderers)
- Docs » Architecture Reference » Channels, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/channels.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fchannels.html)
- Docs » Architecture Reference » CouchDB as the State Database, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/couchdb_as_state_database.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fcouchdb_as_state_database.html)
- Docs » Tutorials » Writing Your First Application, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/write_first_app.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fwrite_first_app.html)
- Docs » Architecture Reference » Transaction Flow, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/txflow.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Ftxflow.html)
- Docs » Architecture Reference » Read-Write set semantics, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/readwrite.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Freadwrite.html)

## 项目源代码

项目源代码会逐步上传到 Github，地址为 [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)。

## Contributor

1. Windstamp, [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)



作者：furnace
链接：https://www.jianshu.com/p/8379f7e3963e
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。



# 2.3 Hyperledger Fabric - 核心概念 - 区块链网络

# 核心概念 - 区块链网络

本主题将从概念上描述 Hyperledger Fabric 如何允许组织在形成区块链网络方面进行协作。如果你是架构师，管理员或开发人员，则可以使用本主题深入了解 Hyperledger Fabric 区块链网络中的主要结构和流程组件。本主题将使用一个易于管理的示例，介绍区块链网络中的所有主要组件。在理解了这个示例之后，你可以阅读文档中其他地方有关这些组件的更多详细信息，或者尝试 [构建示例网络](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fbuild_network.html)。

阅读本主题并了解策略的概念之后，你将对组织为建立控制已部署的 Hyperledger Fabric 网络的策略而需要做出的决策有深入的了解。你还将了解组织如何使用声明性策略 (Hyperledger Fabric 的一项关键功能) 来管理网络演进。简而言之，你将了解 Hyperledger Fabric 的主要技术组件以及组织需要做出的决策。

## 1. 什么是区块链网络？

区块链网络是一种技术基础架构，可为应用程序提供帐本和智能合约 (链码) 服务。首先，智能合约用于生成交易，随后将交易分配到网络中的每个对端点，在该对端节点中，它们一成不变地记录在账本副本中。应用程序的用户可能是使用客户端应用程序或区块链网络管理员的最终用户。

在大多数情况下，多个组织会作为一个联盟聚集在一起以形成网络，并且它们的权限由最初配置网络时联盟所同意的一组策略确定。此外，网络策略会随着联盟组织的一致而随时间变化，这在我们讨论修改策略的概念时会发现。

## 2. 示例网络

在开始之前，让我们告诉你们我们的目标！这是代表示例网络最终状态的图表。

不必担心这看起来很复杂！在研究本主题时，我们将逐步构建网络，以便你了解组织 R1，R2，R3 和 R4 如何为网络贡献基础结构以帮助形成网络。该基础架构实现了区块链网络，并受组成网络的组织 (例如，可以添加新组织的组织) 所同意的策略的约束。你将发现应用程序如何使用由区块链网络提供的帐本和智能合约服务。

![img](https:////upload-images.jianshu.io/upload_images/6280489-020e798493ddef76.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

R1，R2，R3 和 R4 这四个组织共同决定并签署协议，他们将建立和利用 Hyperledger Fabric 网络。R4 已被指定为网络发起者 – 已被授予建立网络初始版本的权力。R4 无意在网络上执行业务交易。 R1 和 R2 以及 R2 和 R3 都需要在整个网络内进行隐私通信。组织 R1 有一个客户端应用程序，可以在通道 C1 内执行业务交易。组织 R2 有一个客户端应用程序，可以在通道 C1 和 C2 中完成类似的工作。组织 R3 有一个客户端应用程序，可以在通道 C2 上执行此操作。对端节点 P1 维护与 C1 关联的帐本 L1 的副本。对端节点 P2 维护与 C1 关联的帐本 L1 的副本和与 C2 关联的帐本 L2 的副本。对端节点 P3 维护与 C2 关联的帐本 L2 的副本。该网络根据网络配置 NC4 中指定的策略规则进行管理，该网络受组织 R1 和 R4 的控制。通道 C1 根据通道配置 CC1 中指定的策略规则进行管理，该通道受组织 R1 和 R2 的控制。通道 C2 根据通道配置 CC2 中指定的策略规则进行管理，该通道受组织 R2 和 R3 的控制。有一个交易排序服务 O4，它充当 N 的网络管理点，并使用系统通道。交易排序服务还支持应用程序通道 C1 和 C2，以便将交易进行排序到区块中并进行分发。四个组织中的每一个都有一个首选的证书颁发机构。

## 3. 创建网络

让我们从头开始，为网络建立基础：

![img](https:////upload-images.jianshu.io/upload_images/6280489-24ae912d4c32b641.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

启动交易排序器后便形成了网络。在我们的示例网络 N 中，包括单个节点 O4 的交易排序服务是根据网络配置 NC4 配置的，该网络配置向组织 R4 提供管理权限。在网络级别，证书颁发机构 CA4 用于将身份分配给 R4 组织的管理员和网络节点。

我们可以看到，定义网络 N 的第一件事 ，即交易排序服务 O4。将交易排序服务视为网络的初始管理点很有帮助。按照事先约定，O4 最初由组织 R4 中的管理员配置和启动，并托管在 R4 中。配置 NC4 包含描述网络管理功能的初始集合的策略。最初，将其设置为仅通过网络授予 R4 权限。我们将在后面看到这些将会改变，但是目前 R4 是网络的唯一成员。

### 3.1 证书颁发机构

你还可以看到证书颁发机构 CA4，它用于向管理员和网络节点颁发证书。 CA4 在我们的网络中起着关键作用，因为它分发 X.509 证书，该证书可用于将组件标识为属于组织 R4。由 CA 颁发的证书还可以用于签署交易，以表明组织认可交易结果 – 这是将其接受到帐本中的前提。让我们更详细地研究 CA 的这两个方面。

首先，区块链网络的不同组件使用证书将自己标识为来自特定组织。这就是为什么通常有不止一个 CA 支持一个区块链网络 – 不同的组织经常使用不同的 CA。我们将在网络中使用四个 CA，每个组织使用一个。确实，CA 非常重要，因此 Hyperledger Fabric 为你提供了一个内置的 CA (称为 Fabric-CA) 来帮助你前进，尽管在实践中，组织将选择使用自己的 CA。

证书到成员组织的映射是通过称为 [成员资格服务提供程序 (Membership Services Provider, MSP)](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23membership-services) 的结构实现的。网络配置 NC4 使用命名的 MSP 来标识由 CA4 分配的证书的属性，该证书将证书持有者与组织 R4 相关联。然后，NC4 可以在策略中使用此 MSP 名称，以向 R4 的参与者授予对网络资源的特定权限。这种策略的一个示例是确定 R4 中可以将新成员组织添加到网络的管理员。我们不会在这些图表上显示 MSP，因为它们只会使它们混乱，但是它们非常重要。

其次，我们稍后将看到 CA 颁发的证书如何成为 [交易](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23transaction) 生成和验证过程的核心。具体来说，X.509 证书用于客户端应用程序 [交易提案](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23proposal) 和智能合约 [交易响应](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23response) 中以进行数字签名 [交易](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23transaction)。随后，托管分类帐副本的网络节点在接受交易到分类帐之前会验证交易签名是否有效。

让我们回顾一下示例区块链网络的基本结构。证书颁发机构 CA4 定义了一组用户访问的资源，即网络 N，该用户对网络 N 中资源的一组权限，如网络配置 NC4 中包含的策略所述。当我们配置并启动交易排序服务节点 O4 时，所有这些都变为现实。

## 4. 添加网络管理员

最初将 NC4 配置为仅允许 R4 用户通过网络进行管理。在下一个阶段，我们将允许组织 R1 用户管理网络。让我们看看网络如何发展：

![img](https:////upload-images.jianshu.io/upload_images/6280489-ce2c79966db56919.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

组织 R4 更新网络配置，使组织 R1 也成为管理员。此后，R1 和 R4 在网络配置上享有同等的权利。

我们看到增加了一个新的组织 R1 作为管理员 - R1 和 R4 现在在网络上拥有平等的权利。我们还可以看到已添加了证书颁发机构 CA1 – 可用于识别 R1 组织中的用户。此后，R1 和 R4 的用户都可以管理网络。

尽管交易排序器节点 O4 在 R4 的基础结构上运行，但是 R1 对其拥有共享的管理权限，只要它可以获取网络访问权限即可。这意味着R1 或 R4 可以更新网络配置 NC4，以允许 R2 组织进行网络子集的操作。这样，即使 R4 正在运行交易排序服务，并且 R1 对其具有完全的管理权限，R2 仍具有创建新联盟的有限权限。

以最简单的形式，交易排序服务是网络中的单个节点，这就是你在示例中看到的内容。交易排序服务通常是多节点的，可以配置为在不同组织中具有不同的节点。例如，我们可以在 R4 中运行 O4 并将其连接到组织 R1 中单独的交易排序器节点 O2。这样，我们将拥有一个多节点，多组织的管理结构。

在本主题的稍后部分，我们将讨论交易排序服务，但现在，仅将交易排序服务视为一个管理点，即可为不同组织提供对网络的受控访问。

## 5. 定义联盟

尽管现在可以通过 R1 和 R4 管理网络，但是几乎没有什么可以完成的。我们需要做的第一件事是定义一个联盟。这个词的字面意思是“拥有共同命运的团体”，因此对于区块链网络中的一组组织来说是一个适当的选择。

让我们看看是如何定义联盟的：
 [图片上传失败...(image-4fac71-1575546776088)]

网络管理员定义了一个联盟 X1，该联盟包含两个成员，组织 R1 和 R2。该联盟定义存储在网络配置 NC4 中，并将在网络开发的下一阶段使用。CA1 和 CA2 是这些组织各自的证书颁发机构。

由于 NC4 的配置方式，只有 R1 或 R4 可以创建新的联盟。该图显示了添加的新联盟 X1，该联盟将 R1 和 R2 定义为其组成组织。我们还可以看到已经添加了 CA2 以从 R2 识别用户。请注意，一个联盟可以有任意数量的组织成员 – 我们刚刚显示了两个，因为它是最简单的配置。

联盟为何重要？我们可以看到，一个联盟定义了网络中彼此共享交易需求的组织集合 - 在这种情况下为 R1 和 R2。如果组织有共同的目标，将它们组合在一起真的很有意义，而这正是正在发生的事情。

网络虽然由单个组织启动，但现在由更大的一组组织控制。我们可以以这种方式启动它，让 R1，R2 和 R4 拥有共享控制权，但是这种构建过程使它更易于理解。

现在，我们将使用联盟 X1 创建 Hyperledger Fabric 区块链中非常重要的部分 - 通道。

## 6. 为联盟创建通道

因此，让我们创建 Fabric 区块链网络的关键部分 - 一个通道。通道是主要的通信机制，联盟的成员可以通过该机制相互通信。网络中可以有多个通道，但现在，我们将从一个通道开始。

让我们看看如何将第一个通道添加到网络中：

[图片上传失败...(image-cc2b7d-1575546776088)]

使用联盟定义 X1 为 R1 和 R2 创建了通道 C1。通道由完全独立于网络配置的通道配置 CC1 控制。CC1 由对 C1 拥有同等权利的 R1 和 R2 管理。R4 在 CC1 中没有任何权利。

通道 C1 为联盟 X1 提供了私有的通信机制。我们可以看到通道 C1 已连接到交易排序服务 O4，但是没有附加任何内容。在网络开发的下一阶段，我们将连接客户端应用程序和对端节点等组件。但是在这一点上，通道代表着未来连接的潜力。

即使通道 C1 是网络 N 的一部分，也可以从中区分出来。还要注意，组织 R3 和 R4 不在此通道中 – 用于 R1 和 R2 之间的交易处理。在上一步中，我们了解了 R4 如何授予 R1 创建新联盟的权限。值得一提的是 R4 还允许 R1 创建通道！在此图中，可能是组织 R1 或 R4 创建了通道 C1。同样，请注意，一个通道可以连接任意数量的组织 - 我们已经展示了两个，因为它是最简单的配置。

再次注意通道 C1 与网络配置 NC4 如何具有完全独立的配置 CC1。CC1 包含用于控制 R1 和 R2 在通道 C1 上的权限的策略 - 正如我们所看到的，R3 和 R4 在此通道上没有权限。如果 R3 和 R4 由 R1 或 R2 添加到通道配置 CC1 中的适当策略，则 R3 和 R4 只能与 C1 交互。一个示例是定义谁可以向通道添加新组织。特别要注意的是，R4 不能将自己添加到通道 C1 中 - 它必须并且只能由 R1 或 R2 授权。

为什么通道如此重要？通道之所以有用，是因为它们为联盟成员之间的私有通信和私有数据提供了一种机制。通道可在其他通道和网络中提供保密性。Hyperledger Fabric 在这方面非常强大，因为它允许组织共享基础结构并同时保持私有状态。这里没有矛盾 - 网络中的不同联盟将需要适当共享不同的信息和流程，而通道则提供了一种有效的机制来做到这一点。通道可有效共享基础架构，同时保持数据和通信的隐私。

我们还可以看到，一旦创建了一个通道，它实际上就是“脱离网络”。从此时到将来，只有在通道配置中明确指定的组织才能对其进行控制。同样，从此时开始对网络配置 NC4 的任何更新都不会直接影响通道配置 CC1。例如，如果联盟定义 X1 更改，它将不会影响通道 C1 的成员。通道之所以有用，是因为它们允许构成通道的组织之间进行隐私通信。此外，通道中的数据与网络的其余部分 (包括其他通道) 完全隔离。

顺便说一句，还有一个特殊的系统通道，供交易排序服务使用。它的行为与常规通道完全相同，因此有时会称为应用程序通道。通常，我们无需担心该通道，但稍后在 [本主题](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fnetwork%2Fnetwork.html%23the-ordering-service) 中我们将对此进行更多讨论。

## 7. 对端节点和账本

现在开始使用该通道将区块链网络和组织组件连接在一起。在网络开发的下一阶段，我们可以看到我们的网络 N 刚刚获得了两个新组件，即对端节点 P1 和帐本实例 L1。

[图片上传失败...(image-de8a46-1575546776088)]

对端节点 P1 已加入通道 C1。P1 物理上存放帐本 L1 的副本。P1 和 O4 可以使用通道 C1 相互通信。

对端节点是托管区块链帐本副本的网络组件！最后，我们开始看到一些可识别的区块链组件！P1 在网络中的目的纯粹是托管帐本 L1 的副本，以供其他人访问。我们可以认为 L1 物理上托管在 P1 上，但逻辑上托管在通道 C1 上。当我们向通道添加更多对端节点时，我们会更清楚地看到这个想法。

P1 配置的关键部分是 CA1 发出的 X.509 身份，该身份将 P1 与组织 R1 相关联。一旦 P1 启动，它就可以使用交易排序器 O4 加入通道 C1。O4 收到此加入请求后，会使用通道配置 CC1 来确定 P1 在此通道上的权限。例如，CC1 确定 P1 是否可以向帐本 L1 读取和/或写入信息。

请注意，对端节点是如何由拥有它们的组织加入通道的，尽管我们仅添加了一个对端节点，但我们将看到网络中多个通道上如何有多个对端节点。稍后我们将看到对端节点可以扮演不同的角色。

## 8. 应用程序和智能合约链码

现在，通道 C1 上有一个帐本，我们可以开始连接客户端应用程序，以使用帐本的主要对象 (对端节点) 提供的某些服务！

请注意网络的增长方式：

[图片上传失败...(image-e3c111-1575546776088)]

智能合约 S5 已安装到 P1 上。组织 R1 中的客户端应用程序 A1 可以使用 S5 通过对端节点 P1 访问帐本。A1，P1 和 O4 都加入了通道 C1，即它们都可以利用该通道提供的通信设施。

在网络开发的下一阶段，我们可以看到客户端应用程序 A1 可以使用通道 C1 连接到特定的网络资源 – 在这种情况下，A1 可以同时连接到对端节点 P1 和交易排序器节点 O4。再次，查看通道如何在网络和组织组件之间进行通信。就像对端节点和交易排序器一样，客户端应用程序将具有将其与组织关联的身份。在我们的示例中，客户端应用程序 A1 与组织 R1 关联，尽管它在 Fabric 区块链网络之外，但仍通过通道 C1 连接到它。

现在看来，A1 可以直接通过 P1 访问帐本 L1，但实际上，所有访问都是通过称为智能合约链码 S5 的特殊程序进行管理的。将 S5 定义为帐本的所有常见访问模式，S5 提供了一套明确定义的方式，可以查询或更新帐本 L1。简而言之，客户应用程序 A1 必须通过智能合约 S5 才能到达帐本 L1！

可以由每个组织中的应用程序开发人员创建智能合约链码，以实施联盟成员共享的业务流程。智能合约用于帮助生成交易，随后可以将交易分配到网络中的每个节点。我们稍后再讨论这个想法，当网络更大时，将更容易理解。现在，需要了解的重要一点是，到此为止，必须在智能合约上执行两次操作。它必须已安装，然后实例化。

### 8.1 安装智能合约

开发了智能合约 S5 之后，组织 R1 中的管理员必须将其安装到对端节点 P1 上。这是一个简单的操作，发生之后，P1 完全了解 S5。具体来说，P1 可以看到 S5 的实现逻辑 - 它用于访问帐本 L1 的程序代码。我们将其与仅描述 S5 的输入和输出而不考虑其实现的 S5 接口进行对比。

当组织在一个通道中有多个对端节点时，它可以选择在其上安装智能合约的对端节点。它不需要在每个对端节点上安装智能合约。

### 8.2 实例化智能合约

但是，仅仅因为 P1 已经安装了 S5，连接到通道 C1 的其他组件却没有意识到，它必须首先在通道 C1 上实例化。在我们的示例中，它只有一个对端节点 P1，组织 R1 中的管理员必须使用 P1 在通道 C1 上实例化 S5。实例化之后，通道 C1 上的每个组件都知道 S5 的存在。在我们的示例中，这意味着 S5 现在可以由客户端应用程序 A1 调用！

请注意，尽管通道上的每个组件现在都可以访问 S5，但它们无法看到其程序逻辑。对于已安装它的那些节点，它仍然是私有的，在我们的示例中，这意味着 P1。从概念上讲，这意味着实例化的是智能合约接口，与已安装的智能合约实现相反。加强这个想法，安装智能合约显示了我们如何将其物理托管在对端节点，而实例化智能合约则显示了我们如何将其视为逻辑上由通道托管。

### 8.3 背书策略

实例化阶段提供的最重要附加信息是 [背书策略](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23endorsement-policy)。它描述了哪些组织必须批准交易，其他组织才能将其接受到其账本副本上。在我们的示例网络中，只有 R1 或 R2 背书交易，才能将交易接受到帐本 L1 上。

实例化的行为将背书策略放置在通道配置 CC1 中。它使通道的任何成员都可以访问它。你可以在 [交易流主题](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Ftxflow.html) 中阅读有关背书策略的更多信息。

### 8.4 调用智能合约

一旦将智能合约安装在对端节点上并在通道上实例化，即可由客户端应用程序 [调用](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23invoke)。客户端应用程序通过向智能合约背书策略指定的组织所拥有的对端节点发送交易提案来完成此任务。交易提案用作智能合约的输入，智能合约使用它来生成背书的交易响应，该响应由对端节点返回到客户端应用程序。

这些交易响应与交易提案打包在一起，形成一个完全背书的交易，可以分发到整个网络。稍后，我们将对其进行更详细的介绍。现在，足以了解应用程序如何调用智能合约以生成背书的交易。

在网络开发的这一阶段，我们可以看到组织 R1 正在完全参与网络。它的应用程序 - 从 A1 开始 - 可以通过智能合约 S5 访问帐本 L1，以生成将由 R1 背书的交易，因此由于它们符合背书策略而被接受到帐本中。

## 9. 网络完成

回想一下，我们的目标是为 X1 联盟 (组织 R1 和 R2) 创建一个通道。网络开发的下一阶段将看到组织 R2 将其基础结构添加到网络中。

让我们看看网络是如何发展的：

![img](https:////upload-images.jianshu.io/upload_images/6280489-11a46360493cec46.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

通过增加组织 R2 的基础结构，网络得到了发展。具体来说，R2 添加了对端节点 P2，该对端节点托管账本 L1 的副本和链码 S5。 P2 也与应用程序 A2 一样加入了通道 C1。A2 和 P2 使用来自 CA2 的证书进行标识。所有这些意味着应用程序 A1 和 A2 都可以使用对端节点 P1 或 P2 在 C1 上调用 S5。

我们可以看到组织 R2 在通道 C1 上添加了一个对端节点 P2。P2 还托管帐本 L1 和智能合约 S5 的副本。我们可以看到 R2 还添加了可以通过通道 C1 连接到网络的客户端应用程序 A2。为此，组织 R2 中的管理员已创建对端节点 P2，并将其加入通道 C1，方法与 R1 中的管理员相同。

我们已经建立了第一个可操作的网络！在网络开发的现阶段，我们拥有一个通道，组织 R1 和 R2 可以彼此完全进行交易。具体来说，这意味着应用程序 A1 和 A2 可以使用智能合约 S5 和账本 L1 在通道 C1 上生成交易。

### 9.1 生成和接受交易

与始终承载帐本副本的对端节点相反，我们看到有两种不同的对端节点：那些拥有智能合约的和那些没有智能合约的。在我们的网络中，每个对端节点都承载智能合约的副本，但是在较大的网络中，将有更多的对端节点不承载智能合约的副本。对端节点只能在安装了智能合约的情况下运行它，但是通过连接到通道，它可以知道智能合约的接口。

你不应该将没有安装智能合约的对端节点视为劣等节点。具有智能合约的对端节点更具有特殊的能力 - 能够产生交易。注意，所有对端节点都可以验证并随后接受或拒绝交易到其账本 L1 的副本上。但是，只有安装了智能合约的对端节点才能参与交易背书的过程，这对生成有效交易至关重要。

我们在本主题中无需担心如何生成，分发和接受交易的确切细节，足以了解我们拥有一个区块链网络，组织 R1 和 R2 可以将这些信息和流程共享为账本捕获的交易。在其他主题中，我们将学到更多有关交易，账本，智能合约的信息。

### 9.2 对端节点的类型

在 Hyperledger Fabric 中，尽管所有对端节点都相同，但是它们可以根据网络的配置方式承担多个角色。现在，我们对典型的网络拓扑有了足够的了解，可以描述这些角色。

- [确认对端节点 (Committing peer)](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23commitment)。通道中的每个对端节点都是提交对端节点。它接收生成交易的区块，随后将这些交易进行验证，然后再将它们提交到对端节点的帐本副本，作为追加操作。

- [背书对端节点 (Endorsing peer)](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23endorsement)。如果安装了智能合约，则拥有智能合约的每个对端节点都可以成为背书节点。但是，实际上要成为背书节点，客户端应用程序必须使用对端节点上的智能合约来生成数字签名的交易响应。背书对端节点一词是对此事实的明确引用。

  智能合约的背书策略可确定组织，其对端节点应先对生成的交易进行数字签名，然后才能将其接受到提交的对端节点的账本副本中。

这是对端节点的两种主要类型。对端节点可以采用其他两个角色：

- [领导对端节点 (Leader peer)](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23leading-peer)。当组织在一个通道中有多个对端节点时，领导对端节点是负责将交易从交易排序器分发到组织中其他提交对端节点的节点。对端节点可以选择参与静态或动态领导选择。

  因此，从领导者的角度考虑两组对端节点是有帮助的 - 那些具有静态领导者选择的对端节点和那些具有动态领导者选择的对端节点。对于静态集，可以将零个或多个对端节点配置为领导者。对于动态集，将由该集选出一个对端节点作为领导者。此外，在动态集中，如果领导者对端节点失败，则其余对端节点将重新选举领导者。

  这意味着组织的对端节点可以让一个或多个领导者连接到交易排序服务。这可以帮助提高处理大量交易的大型网络的弹性和可伸缩性。

- [锚点对端节点 (Anchor peer)](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23anchor-peer)。如果对端节点需要与另一个组织中的对端节点进行通信，则可以使用该组织的通道配置中定义的锚点对端节点之一。一个组织可以为其定义零个或多个锚点，并且锚点可以帮助处理许多不同的跨组织通信场景。

注意，一个对端节点可以同时是一个提交对端节点，背书对端节点，领导对端节点和锚点对端节点！仅锚点对端节点是可选的 – 出于所有实际目的，总会有一个领导对端节点，至少一个背书 对端节点和至少一个提交对端节点。

### 9.3 安装而不实例化

与组织 R1 类似，组织 R2 必须将智能合约 S5 安装到其对等节点 P2 上。这很明显 – 如果应用程序 A1 或 A2 希望在对端节点 P2 上使用 S5 来生成事务，则必须首先存在它，安装是发生这种情况的机制。此时，对端节点 P2 具有智能合约和帐本的物理副本，像 P1 一样，它可以在其帐本 L1 的副本上生成并接受交易。

但是，与组织 R1 相比，组织 R2 不需要在通道 C1 上实例化智能合约 S5。这是因为组织 R1 已在通道上实例化了 S5。实例化只需要发生一次。随后加入该通道的任何对端节点都知道智能合约 S5 可用于该通道。这一事实反映了这样一个事实，即账本 L1 和智能合约实际上以物理方式存在于对端节点上，并且以逻辑方式存在于通道上。R2 只是将 L1 和 S5 的另一个物理实例添加到网络。

在我们的网络中，我们可以看到通道 C1 连接了两个客户端应用程序，两个对端节点和一个交易排序服务。由于只有一个通道，因此只有一个逻辑帐本可以与这些组件进行交互。对端节点 P1 和 P2 具有帐本 L1 的相同副本。智能合约 S5 的副本通常将使用相同的编程语言完全相同地实现，但如果不同，则它们在语义上必须等效。

我们可以看到，将对端节点小心地添加到网络可以帮助支持增加的吞吐量，稳定性和弹性。例如，网络中更多的对端节点将允许更多的应用程序连接到它，如果计划内或计划外的停机，组织中的多个对端节将提供额外的弹性。

这一切都意味着可以配置支持各种操作目标的复杂拓扑 – 网络可以达到的规模没有理论上的限制。此外，单个组织内的对端节点有效地发现并彼此通信的技术机制 ([gossip 协议](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fgossip.html%23gossip-protocol)) 将容纳大量对端节点，以支持此类拓扑。

仔细使用网络和通道策略可以使大型网络得到良好管理。组织可以自由地将对端节点添加到网络，只要它们符合网络约定的策略即可。网络和通道策略在自治和控制之间建立了平衡，这是去中心化网络的特征。

## 10. 简化示意图

现在，我们将简化用于表示示例区块链网络的示意图。随着网络规模的扩大，最初用于帮助我们了解通道的线路将变得很繁琐。想象一下，如果添加另一个对端节点或客户端应用程序或另一个通道，图将变得多么复杂？

这就是我们将在一分钟内要做的事情，因此在我们这样做之前，让我们简化示意图。这是到目前为止我们开发的网络的简化表示：

![img](https:////upload-images.jianshu.io/upload_images/6280489-535fda6e5a076ed8.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

该图显示了与网络 N 中的通道 C1 有关的事实，如下所示：客户端应用程序 A1 和 A2 可以使用通道 C1 与对端节点 P1 和 P2 以及交易排序器 O4 进行通信。对端节点 P1 和 P2 可以使用通道 C1 的通信服务。交易排序服务 O4 可以利用通道 C1 的通信服务。通道配置 CC1 适用于通道 C1。

请注意，通过用连接点替换通道线来简化网络图，显示为蓝色圆圈，其中包括通道号。没有信息丢失。此表示形式更具伸缩性，因为它消除了交叉线。这使我们可以更清楚地表示更大的网络。我们通过专注于组件和通道之间的连接点，而不是通道本身来实现了这种简化。

## 11. 添加另一个联盟

在网络开发的下一阶段，我们介绍组织 R3。我们将为组织 R2 和 R3 提供一个独立的应用程序通道，使他们可以彼此进行交易。该应用程序通道将与先前定义的通道完全分开，因此 R2 和 R3 交易可以保持私有。

让我们回到网络层级，为 R2 和 R3 定义一个新的联盟 X2：

![img](https:////upload-images.jianshu.io/upload_images/6280489-deb45afd924a90b6.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

组织 R1 或 R4 的网络管理员添加了新的联盟定义 X2，其中包括组织 R2 和 R3。这将用于为 X2 定义一个新通道。

请注意，网络现在定义了两个联盟：组织 R1 和 R2 使用 X1，组织 R2 和 R3 使用 X2。引入了联盟 X2，以便能够为 R2 和 R3 创建新的通道。

只有在网络配置策略 NC4 中明确标识为具有适当权限的组织才能创建新通道，即 R1 或 R4。这是策略的示例，该策略将可以在网络级别管理资源的组织与可以在通道级别管理资源的组织区分开来。看到这些策略在起作用，有助于我们理解为什么 Hyperledger Fabric 具有复杂的分层策略结构。

实际上，联盟定义 X2 已添加到网络配置 NC4 中。我们将在文档的其他地方讨论此操作的确切机制。

## 12. 添加新的通道

现在，使用这个新的联盟定义 X2 创建一个新通道 C2。为了帮助你进一步了解更简单的通道符号，我们使用了两种视觉样式 - 通道 C1 用蓝色圆形端点表示，而通道 C2 用红色连接线表示：

![img](https:////upload-images.jianshu.io/upload_images/6280489-d0de176ec8cfeddd.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

使用联盟定义 X2 为 R2 和 R3 创建了一个新的通道 C2。该通道具有完全独立于网络配置 NC4 和通道配置 CC1 的通道配置 CC2。通道 C2 由 R2 和 R3 管理，它们具有 CC2 中策略定义的对 C2 相同的权限。R1 和 R4 在 CC2 中均未定义任何权限。

通道 C2 为联盟 X2 提供了专用的通信机制。同样，请注意组织是以联盟形式组织起来的，以及构成通道。通道配置 CC2 现在包含管理通道资源的策略，并通过通道 C2 向组织 R2 和 R3 分配管理权限。它仅由 R2 和 R3 管理，R1 和 R4 在通道 C2 中没有权限。例如，随后可以更新通道配置 CC2 以添加组织来支持网络增长，但这只能由 R2 或 R3 完成。

请注意，通道配置 CC1 和 CC2 如何保持彼此完全独立，并与网络配置 NC4 完全独立。再次，我们看到了 Hyperledger Fabric 网络的去中心化性质，创建通道 C2 后，组织 R2 和 R3 会独立于其他网络元素来管理它。通道策略始终保持彼此独立，并且只能由有权在通道中这样做的组织进行更改。

随着网络和通道的发展，网络和通道的配置也将随之发展。有一个过程可以通过受控的方式完成此过程 - 涉及配置交易，这些交易捕获对这些配置的更改。每次配置更改都会产生一个新的配置区块交易，在 [本主题的后面](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fnetwork%2Fnetwork.html%23the-ordering-serivce)，我们将看到如何验证和接受这些区块以分别创建更新的网络和通道配置。

### 12.1 网络和通道配置

在整个示例网络中，我们看到了网络和通道配置的重要性。这些配置很重要，因为它们封装了网络成员同意的策略，这些策略为控制对网络资源的访问提供了共享引用。网络和通道配置还包含有关网络和通道组成的事实，例如联盟名称及其组织。

例如，当首先使用交易排序服务节点 O4 形成网络时，其行为由网络配置 NC4 控制。NC4 的初始配置仅包含允许组织 R4 管理网络资源的策略。随后将 NC4 更新为还允许 R1 管理网络资源。进行此更改后，组织 R1 或 R4 中连接到 O4 的任何管理员都将具有网络管理权限，因为这是 NC4 网络配置中允许的策略。在内部，交易排序服务中的每个节点都会记录网络配置中的每个通道，以便在网络级别上记录每个创建的通道。

这意味着尽管交易排序服务节点 O4 是创建联盟 X1 和 X2 以及通道 C1 和 C2 的参与者，但网络的智能包含在 O4 遵循的网络配置 NC4中。只要 O4 表现良好，并且在处理网络资源时正确执行 NC4 中定义的策略，我们的网络就会按照所有组织的同意行事。在许多方面，NC4 比 O4 更重要，因为它最终控制了网络访问。

相同的原则适用于有关对端节点的通道配置。在我们的网络中，P1 和 P2 同样是好的参与者。当对端节点 P1 和 P2 与客户端应用程序 A1 或 A2 交互时，它们每个都使用在通道配置 CC1 中定义的策略来控制对通道 C1 资源的访问。

例如，如果 A1 要访问对端节点 P1 或 P2 上的智能合约链码 S5，则每个对端节点都使用其 CC1 副本确定 A1 可以执行的操作。例如，可以根据 CC1 中定义的策略允许 A1 从账本 L1 读取或写入数据。稍后我们将在通道及其通道配置 CC2 中看到针对参与者的相同模式。再次，我们可以看到，尽管对端节点和应用程序是网络中的关键角色，但它们在通道中的行为更多地由通道配置策略决定，而不是其他任何因素。

最后，了解物理上如何实现网络和通道配置将很有帮助。我们可以看到网络和通道配置在逻辑上是唯一的 - 只有一个网络配置，只有一个通道配置。这个很重要，访问网络或通道的每个组件都必须对授予不同组织的权限有共同的了解。

即使从逻辑上讲只有一个配置，但实际上它被构成网络或通道的每个节点复制并保持一致。例如，在我们的网络中，对端节点 P1 和 P2 都具有通道配置 CC1 的副本，而到网络完全完成时，对端节点 P2 和 P3 都将具有通道配置 CC2 的副本。类似地，交易排序服务节点 O4 具有网络配置的副本，但是在 [多节点配置](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fnetwork%2Fnetwork.html%23the-ordering-service) 中，每个交易排序服务节点将具有其自己的网络配置副本。

使用用于用户交易 (但用于配置交易) 的相同区块链技术，可以使网络和通道配置保持一致。要更改网络或通道配置，管理员必须提交配置交易以更改网络或通道配置。必须由适当策略中确定负责配置更改的组织签名。这项策略称为 mod_policy，我们将在 [后面讨论](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fnetwork%2Fnetwork.html%23changing-policy)。

的确，交易排序服务节点运行着一个微型区块链，通过我们前面提到的系统通道连接。使用系统通道交易排序服务节点可以分发网络配置交易。这些交易用于在每个交易排序服务节点上合作维护网络配置的一致副本。以类似的方式，应用程序通道中的对端节点可以分发通道配置交易。同样，这些交易用于在每个对端节点上维护通道配置的一致副本。

通过在物理上分布在逻辑上唯一的对象之间的这种平衡是 Hyperledger Fabric 中的常见模式。例如，逻辑上单一的对象 (如网络配置) 实际上是在一组交易排序服务节点之间进行物理复制的。我们还会在通道配置，帐本以及某种程度上将智能合约安装在多个位置，但它们的接口逻辑上存在于通道级别的情况下看到它。你会在 Hyperledger Fabric 中一次又一次地看到这种模式，这种模式使 Hyperledger Fabric 既去中心化又可管理。

## 13. 添加另一个对端节点

现在，组织 R3 可以完全参与通道 C2 了，让我们将其基础结构组件添加到渠道中。我们不会一次添加一个组件，而是一次添加一个对端节点，其帐本的本地副本，智能合约和客户端应用程序！

让我们看一下添加了组织 R3 组件的网络：

![img](https:////upload-images.jianshu.io/upload_images/6280489-51faec930ba97d9e.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

该图显示了与网络 N 中的通道 C1 和 C2 有关的事实，如下所示：客户端应用程序 A1 和 A2 可以使用通道 C1 与对端节点 P1 和 P2 进行通信，以及交易排序服务 O4。客户应用程序 A3 可以使用通道 C2 与对端节点 P3 和交易排序服务 O4 进行通信。交易排序服务 O4 可以利用通道 C1 和 C2 的通信服务。通道配置 CC1 适用于通道 C1，CC2 适用于通道 C2。

首先，请注意，由于对端节点 P3 连接到通道 C2，因此它对使用通道 C1 的那些对端节点具有不同的帐本 L2。帐本 L2 有效地限定在通道 C2 上。帐本 L1 是完全独立的，它的作用域是通道 C1。这是有道理的 – 通道 C2 的目的是在联盟 X2 的成员之间提供隐私通信，而帐本 L2 是其交易的私有存储。

以类似的方式，安装在对端节点 P3 上并实例化在通道 C2 上的智能合约 S6 用于提供对账本 L2 的受控访问。应用程序 A3 现在可以使用通道 C2 调用由智能合约 S6 提供的服务，以生成可以接受到网络中账本 L2 的每个副本上的交易。

此时，我们有了一个网络，其中定义了两个完全独立的通道。这些通道为组织之间的相互交易提供了独立管理的设施。同样，这是工作中的分权，我们在控制和自治之间取得平衡。这是通过将策略应用到受不同组织控制并影响不同组织的通道来实现的。

## 14. 将一个对端节点加入到多个通道

在网络开发的最后阶段，让我们把重点放回到组织 R2。我们可以利用 R2 是 X1 和 X2 联盟成员的事实，将 R2 加入多个通道。

[图片上传失败...(image-712d77-1575546776088)]

该图显示了与网络 N 中的通道 C1 和 C2 有关的事实，如下所示：客户端应用程序 A1 可以使用通道 C1 与对端节点 P1 和 P2 进行通信，以及交易排序服务 O4。客户端应用程序 A2 可以使用通道 C1 与对端节点 P1 和 P2 进行通信，并使用通道 C2 与对端节点 P2 和 P3 进行通信以及交易排序服务 O4。客户端应用程序 A3 可以使用通道 C2 与对端节点 P3 和 P2 以及交易排序服务 O4 进行通信。交易排序服务 O4 可以利用通道 C1 和 C2 的通信服务。通道配置 CC1 适用于通道 C1，CC2 适用于通道 C2。

我们可以看到 R2 是网络中的一个特殊组织，因为它是两个应用程序通道中唯一的组织！它能够与通道 C1 上的组织 R1 进行交易，同时它还可以与其他通道 C2 上的组织 R3 进行交易。

请注意，对端节点 P2 如何为通道 C1 安装了智能合约 S5，并为通道 C2 安装了智能合约 S6。对端节点 P2 通过不同账本的不同智能合约同时是两个通道的正式成员。

这是一个非常强大的概念 - 通道既提供了组织分离的机制，又提供了组织之间协作的机制。一直以来，此基础结构是由一组独立的组织提供并在它们之间共享的。

同样重要的是要注意，对端节点 P2 的行为受其进行交易的通道的控制非常不同。具体来说，通道配置 CC1 中包含的策略规定了 P2 在通道 C1 中进行交易时可用于的操作，而通道配置 CC2 中的策略则控制了 P2 在通道 C2 中的行为。

同样，这是理想的 –  R2 和 R1 同意了通道 C1 的规则，而 R2 和 R3 同意了通道 C2 的规则。这些规则是在各自的通道策略中捕获的 - 通道中的每个组件都可以并且必须使用它们来强制执行正确的行为，这已经达成共识。

同样，我们可以看到客户端应用程序 A2 现在能够在通道 C1 和 C2 上进行交易。同样，它也将由适当通道配置中的策略控制。顺便说一句，请注意客户端应用程序 A2 和对端节点 P2 正在使用混合的可视词汇表 - 线路和连接。你可以看到它们是等效的。它们是相似的图示。

### 14.1 交易排序服务

细心的读者可能会注意到交易排序服务节点似乎是一个集中式组件。它最初用于创建网络，并连接到网络中的每个通道。即使我们在控制交易排序器的网络配置策略 NC4 中添加了 R1 和 R4，该节点仍在 R4 的基础架构上运行。在去中心化的世界里，这看起来是错误的！

不用担心我们的示例网络显示了最简单的交易排序服务配置，以帮助你了解网络管理点的概念。实际上，交易排序服务本身也可以完全去中心！前面我们提到过，交易排序服务可能由不同组织拥有的许多单个节点组成，因此让我们看看如何在示例网络中完成该工作。

让我们看一个更现实的交易排序服务节点配置：

![img](https:////upload-images.jianshu.io/upload_images/6280489-b6c06ca8ee72f324.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

多个组织的交易排序服务。交易排序服务包括交易排序服务节点 O1 和 O4。 O1 由组织 R1 提供，节点 O4 由组织 R4 提供。网络配置 NC4 为组织 R1 和 R4 的参与者定义了网络资源权限。

我们可以看到，该订购服务已完全去中心化 – 在组织 R1 中运行，在组织 R4 中运行。网络配置策略 NC4 允许 R1 和 R4 在网络资源上享有同等的权利。来自组织 R1 和 R4 的客户端应用程序和对端节点可以通过连接到节点 O1 或节点 O4 来管理网络资源，因为这两个节点的行为方式相同，如网络配置 NC4 中的策略所定义。实际上，来自特定组织的参与者往往会使用其上级组织提供的基础架构，但是并非总是如此。

### 14.2 去中心化交易分发

交易排序服务不仅是网络的管理点，而且还提供了另一个关键功能 – 它是交易的分发点。交易排序服务是一个组件，它从应用程序中收集已背书的交易并将其排序到交易区块中，然后将其分发到通道中的每个对端节点。在这些提交对端节点的每个对端节点，记录交易 (有效或无效)，并适当更新其本地帐本副本。

请注意，交易排序服务节点 O4 在通道 C1 上的作用与在网络 N 上的作用是非常不同的。在通道级别进行操作时，O4 的作用是在通道 C1 内收集交易并分配区块。它根据通道配置 CC1 中定义的策略执行此操作。相反，在网络级别执行操作时，O4 的作用是根据网络配置 NC4 中定义的策略为网络资源提供管理点。再次注意，这些角色是如何分别由通道和网络配置中的不同策略定义的。这应该向你增强 Hyperledger Fabric 中基于声明策略的配置的重要性。策略既定义了联盟的每个成员，又用于控制联盟的每个成员的行为。

我们可以看到，交易排序服务与 Hyperledger Fabric 中的其他组件一样，是完全去中心化的组件。无论是充当网络管理点，还是充当通道中的块分配器，都可以根据需要在网络中的多个组织中分布其节点。

### 14.3 改变策略

在整个示例网络的探索过程中，我们已经了解了控制系统中参与者行为的策略的重要性。我们仅讨论了一些可用的策略，但是可以声明性地定义许多策略来控制行为的各个方面。这些单独的策略在文档的其他地方进行了讨论。

最重要的是，Hyperledger Fabric 提供了独特而强大的策略，允许网络和通道管理员自行管理策略更改！基本的哲学思想是，无论是在组织内部还是组织之间发生的变化，还是由外部监管机构施加的变化，变化都是一个常数。例如，新组织可以加入通道，或者现有组织的权限可以增加或减少。让我们研究一下 Hyperledger Fabric 中更改策略的实施方式。

他们的主要理解点是，策略更改由策略本身内的策略管理。修改策略 (简称 mod_policy) 是管理更改的网络或通道配置中的一流策略。让我们举两个简单的例子，说明我们如何使用 mod_policy 来管理网络中的更改！

第一个示例是最初建立网络时的情况。此时，仅组织 R4 被允许管理网络。实际上，这是通过使 R4 成为网络配置 NC4 中定义的唯一拥有网络资源许可的组织来实现的。此外，NC4 的 mod_policy 仅提及组织 R4 – 仅允许 R4 更改此配置。

然后，我们对网络 N 进行了演进，以允许组织 R1 来管理网络。 R4 通过将 R1 添加到用于通道创建和联盟创建的策略中来做到这一点。由于此更改，R1 能够定义联盟 X1 和 X2，并创建通道 C1 和 C2。 R1 对网络配置中的通道和联盟策略具有同等的管理权限。

但是，R4 可以通过网络配置为 R1 提供更多的权限！ R4 可以将 R1 添加到 mod_policy 中，以便 R1 也能够管理网络策略的更改。

第二个权限比第一个权限强大得多，因为 R1 现在可以完全控制网络配置 NC4！这意味着 R1 原则上可以从网络中删除 R4 的管理权限。实际上，R4 将配置 mod_policy，以便 R4 也需要批准更改，或者 mod_policy 中的所有组织都必须批准更改。有足够的灵活性可以使 mod_policy 复杂到可以支持所需的任何更改过程。

这是在工作的 mod_policy，它使基本配置可以优雅地演变为复杂的配置。在所有相关组织的同意下，这种情况一直存在。 mod_policy 的行为类似于网络或通道配置中的其他所有策略，它定义了一组允许更改mod_policy 本身的组织。

在本小节中，我们仅涉及到策略和 mod_policy 的作用。在策略主题中将对它进行更详细的讨论，但现在让我们回到完成的网络中！

## 15. 完整的网络

让我们使用一致的图示来回顾一下我们的网络。我们使用更紧凑的图示对其进行了稍微的重组，因为它可以更好地适应更大的拓扑：

[图片上传失败...(image-858e1-1575546776088)]

在此图中，我们看到 Fabric 区块链网络由两个应用程序通道和一个交易排序通道组成。组织 R1 和 R4 负责交易排序通道，组织 R1 和 R2 负责蓝色应用程序通道，而组织 R2 和 R3 负责红色应用程序通道。客户端应用程序 A1 是组织 R1 的组件，而 CA1 是其证书颁发机构。请注意，组织 R2 的对端节点 P2 可以使用蓝色和红色应用程序通道的通信设施。每个应用程序通道都有自己的通道配置，在这种情况下为 CC1 和 CC2。系统通道的通道配置是网络配置 NC4 的一部分。

我们的概念之旅已结束，以构建示例 Hyperledger Fabric 区块链网络。我们创建了一个具有两个通道和三个对端节点，四个智能合约和一个交易排序服务的四个组织的网络。它由四个证书颁发机构支持。它为三个客户端应用程序提供帐本和智能合约服务，他们可以通过两个通道与之交互。花一点时间浏览图中的网络详细信息，然后随时阅读本主题以增强你的知识，或者转到更详细的主题。

## 16. 网络组件摘要

以下是我们讨论过的网络组件的简要摘要：

- [帐本 (Ledger)](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23ledger)。每个通道一个。由 [区块链](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23block) 和 [世界状态](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23world-state) 组成。
- [智能合约 (又名链码) (Smart Contract, aka chaincode)](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23smart-contract)
- [对端节点 (Peer node)](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23peer)
- [交易排序服务 (Ordering service)](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23ordering-service)
- [通道 (Channel)](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23channel)
- [证书颁发机构 (Certificate Authority)](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html%23hyperledger-fabric-ca)

## 17. 网络汇总

在本主题中，我们已经看到了不同的组织如何共享其基础架构以提供集成的 Hyperledger Fabric 区块链网络。我们已经看到了如何将集体基础结构组织成提供独立管理的私有通信机制的通道。我们已经了解了如何通过使用来自各自证书颁发机构的证书来将诸如客户端应用程序，管理员，对端节点和交易排序器之类的参与者识别为来自不同组织。反过来，我们已经看到了定义这些组织参与者对网络和通道资源所拥有的同意许可的政策的重要性。

## Reference

- Docs » Key Concepts » Blockchain network, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/network/network.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fnetwork%2Fnetwork.html)
- Docs » Tutorials » Building Your First Network, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/build_network.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fbuild_network.html)
- Docs » Glossary, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/glossary.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fglossary.html)
- Docs » Architecture Reference » Transaction Flow, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/txflow.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Ftxflow.html)

## 项目源代码

项目源代码会逐步上传到 Github，地址为 [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)。

## Contributor

1. Windstamp, [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)



3人点赞



[hl.fabric.doc]()





作者：furnace
链接：https://www.jianshu.com/p/0a5dbc53159e
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

# 2.4 Hyperledger Fabric - 核心概念 - 交易排序服务

# 核心概念 - 交易排序服务

本主题从概念上介绍了交易排序的概念，交易排序器如何与对端节点交互，他们在交易流中扮演的角色以及交易排序服务的当前可用实现的概述，尤其着重于 Raft 交易排序服务的实现。

## 1. 什么是交易排序

以太坊和比特币等许多分布式区块链，是无需许可制的，这意味着任何节点都可以参与共识过程，在此过程中，交易被排序并打包成区块。因此，这些系统依赖于概率共识算法 (probabilistic consensus algorithm)，该算法最终可以很高的概率确保达到帐本的一致性，但是仍然容易受到帐本分叉的影响，因为网络中的不同参与者对接受的交易顺序具有不同的视图。

Hyperledger Fabric 的工作原理有所不同。它具有一种称为交易排序器的节点来执行交易排序，该节点与其他节点一起构成交易排序服务。由于 Fabric 的设计依赖于确定性共识算法 (deterministic consensus algorithm)，因此对端节点验证的由交易排序服务生成的任何区块都可以保证是最终且正确的。账本无法像在其他许多分布式区块链中那样分叉。

除了提高确定性之外，将链码执行的背书 (在对端点发生) 与交易排序分开可以为 Fabric 提供性能和可伸缩性方面的优势，消除了在同一节点执行交易和交易排序时可能出现的瓶颈。

## 2. 交易排序器节点和通道配置

除其交易排序角色外，交易排序器还维护被允许创建通道的组织的列表。该组织列表称为“联盟”，该列表本身保留在“交易排序器系统通道”的配置中。默认情况下，该列表及其所处的通道只能由交易排序器管理员编辑。请注意，交易排序服务可能会保留其中几个列表，这使该联盟成为 Fabric 多租户的工具。

交易排序器还对通道实施基本的访问控制，从而限制了谁可以对其进行读写数据以及谁可以对其进行配置。请记住，有权修改通道中的配置元素的人员必须受相关管理员在创建联盟或通道时设置的策略的约束。配置交易由交易排序器处理，因为交易排序器需要知道当前策略集才能执行其基本形式的访问控制。在这种情况下，交易排序器将处理配置更新，以确保请求者具有适当的管理权限。如果是这样，则交易排序器针对现有配置验证更新请求，生成新的配置交易，并将其打包到一个区块中，该区块中继到通道上的所有对端节点。然后，对端节点处理配置交易，以验证交易排序器批准的修改确实确实满足通道中定义的策略。

## 3. 交易排序器节点和身份标识

与区块链网络交互的所有事物 (包括对端节点，应用程序，管理员和交易排序器) 都从其数字证书和其成员资格服务提供商 (Membership Service Provider, MSP) 定义中获取其组织身份。

有关身份和 MSP 的更多信息，请查看有关 [身份](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fidentity%2Fidentity.html) 和 [MSP](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fmembership%2Fmembership.html) 的文档。

就像对端节点一样，交易排序节点属于组织。与对端节点一样，每个组织都应使用单独的证书颁发机构 (Certificate Authority, CA)。此 CA 是否将充当根 CA，还是你选择部署根 CA，然后选择与该根 CA 关联的中间 CA，取决于你。

## 4. 交易排序器和交易流程

### 4.1 第一阶段：提案

我们从 [对端节点](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fpeers%2Fpeers.html) 主题中看到，它们构成了区块链网络的基础，托管了帐本，应用程序可以通过智能合约对其进行查询和更新。

具体来说，想要更新帐本的应用程序涉及三个阶段的过程，以确保区块链网络中的所有对端节点都保持其帐本彼此一致。

在第一阶段，客户端应用程序将交易提案发送给对端节点的子集，该对端节点将调用智能合约以产生提案的账本更新，然后对结果进行背书。背书的对端节点目前不将提案的更新应用于其账本副本。相反，背书的对端节点将提案响应返回到客户端应用程序。背书的交易提案将最终在第二阶段按顺序排列，然后分发给所有对端节点以进行最终验证并在第三阶段进行提交。

要深入了解第一阶段，请参考 [对端节点](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fpeers%2Fpeers.html) 主题。

### 4.2 第二阶段：将交易排序和打包成区块

在交易的第一阶段完成之后，客户端应用程序已收到来自一组对端节点的背书交易提案响应。现在该进行交易的第二阶段了。

在此阶段，应用程序客户端将包含背书交易提案响应的交易提交到交易排序服务节点。交易排序服务创建交易区块，最终将这些交易区块分发给通道上的所有对端节点，以进行最终验证并在第三阶段进行提交。

交易排序服务节点同时接收来自许多不同应用程序客户端的交易。这些交易排序服务节点一起工作以共同形成交易排序服务。它的工作是将提交的交易的批处理按定义明确的顺序排列，然后打包成区块。这些区块将成为区块链的区块！

区块中的交易数量取决于与区块的所需大小和最大经过时间有关的通道配置参数 (准确地说是 BatchSize 和 BatchTimeout 参数)。然后，将这些区块保存到交易排序器的帐本中，并分发给已加入该通道的所有对端节点。如果对端节点此时恰好处于关闭状态，或者稍后加入了该通道，则在重新连接到交易排序服务节点后，或者与另一对端节点 gossiping，它将接收到区块。我们将在第三阶段看到对端节点如何处理此区块。

![img](https:////upload-images.jianshu.io/upload_images/6280489-b978d0537180340c.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

交易排序节点的第一个角色是打包提案的帐本更新。在上面的示例中，应用程序 A1 将由 E1 和 E2 背书的交易 T1 发送给交易排序器 O1。并行地，应用程序 A2 将由 E1 背书的交易 T2 发送到交易排序器 O1。O1 将来自应用程序 A1 的交易 T1 和来自应用程序 A2 的交易 T2 以及来自网络中其他应用程序的其他交易打包到区块 B2 中。我们可以看到，在 B2 中，交易顺序为 T1，T2，T3，T4，T6，T5 – 可能不是这些交易到达交易排序器的顺序！ (此示例显示了非常简化的交易排序服务配置，其中只有一个交易排序器节点。)

值得注意的是，一个区块中的交易顺序不一定与交易排序服务所接收的交易顺序相同，因为可能有多个交易排序服务节点大约在同一时间接收交易。重要的是，交易排序服务将交易置于严格的顺序中，对端节点将在验证和提交交易时使用此顺序。

区块内交易的这种严格排序使 Hyperledger Fabric 与其他区块链略有不同，在其他区块链中，同一笔交易可以打包成多个不同的区块，竞争形成一个链。在 Hyperledger Fabric 中，交易排序服务生成的区块是最终的。一旦将交易写入一个区块，就可以确保其在账本中的位置。正如我们之前所说，Hyperledger Fabric 的确定性意味着没有账本分叉 - 验证的交易将永远不会被还原或丢弃。

我们还可以看到，尽管对端节点执行智能合约并处理交易，但交易排序器绝对不会这样做。到达交易排序器的每个授权交易都被打包在一个区块中 - 交易排序器不对交易的内容做出判断 (如前所述，通道配置交易除外)。

在第二阶段结束时，我们看到交易排序器负责简单但至关重要的过程，这些过程包括收集提案的交易更新，进行交易排序并将它们打包成区块以便分发。

### 4.3 第三阶段：验证和提交

交易工作流程的第三阶段涉及从交易排序器到对端节点的区块的分发和后续验证，在这里可以将它们应用于帐本。

阶段 3 从交易排序器将区块分配给与其连接的所有对端节点开始。还值得注意的是，并非每个对端节点都需要连接到交易排序器，对端节点可以使用 [gossip](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fgossip.html) 协议将区块级联到其他对端节点。

每个对端节点将独立地但以确定性的方式验证分布式区块，以确保账本保持一致。具体来说，通道中的每个对端节点都将验证区块中的每个交易，以确保其已被所需组织的对端节点认可，其背书相匹配，并且未被其他最近提交的交易 (可能已经在交易中进行的交易) 使之无效。无效的交易仍保留在交易排序器创建的不可变区块中，但它们被对端节点标记为无效，并且不会更新帐本的状态。

![img](https:////upload-images.jianshu.io/upload_images/6280489-937fe8c7af3df260.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image

交易排序节点的第二个作用是将区块分配给对端节点。在该示例中，交易排序器 O1 将区块 B2 分配给对端节点 P1 和对端节点 P2。对端节点 P1 处理区块 B2，从而将新区块添加到 P1 上的帐本 L1。并行地，对端节点 P2 处理区块 B2，导致将新区块添加到 P2 上的帐本 L1。一旦该过程完成，账本 L1 就已在对端节点 P1 和 P2 上进行了一致更新，并且每个账本可以通知连接的应用程序交易已被处理。

总而言之，在第三阶段中，交易排序服务生成的区块始终应用于帐本。严格按顺序将交易划分为区块，每个对端节点都可以验证交易更新是否在整个区块链网络中得到一致应用。

要深入了解第三阶段，请返回 [对端节点](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fpeers%2Fpeers.html) 主题。

## 5. 交易排序服务实现

尽管当前可用的每个交易排序服务都以相同的方式处理交易和配置更新，但是仍然存在几种不同的实现方式，用于在交易排序服务节点之间对严格的交易排序达成共识。

有关如何启动交易排序节点的信息 (无论将使用哪种实现节点)，请查阅 [我们有关启动交易排序节点的文档](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Forderer_deploy.html)。

- Solo

  交易排序服务的 Solo 实现被恰当地命名：它仅具有一个交易排序节点。结果，它不是，而且永远不会是容错的。因此，无法将 Solo 实现用于生产，但它们是测试应用程序和智能合约或创建概念证明 (Proof of Concept) 的理想选择。但是，如果你想将此 PoC 网络扩展到生产环境中，则可能要从一个单节点 Raft 群集开始，因为它可能已重新配置为添加其他节点。

- Raft

  从 v1.4.1 开始，Raft 是一项新的功能，它是基于 `etcd` 中 Raft 协议实现的崩溃容错 (Carsh Fault Tolerant, CFT) 交易排序服务。Raft 遵循“领导者和跟随者”模型，其中 (每个通道) 选举领导者节点，并由跟随者复制其决策。与基于 Kafka 的交易排序服务相比，Raft 交易排序服务应该更易于设置和管理，并且其设计允许不同的组织将节点贡献给分布式交易排序服务。

- Kafka

  类似于基于 Raft 的交易排序，Apache Kafka 是一种 CFT 实现，它使用“领导者和跟随者”节点配置。Kafka 利用 ZooKeeper 集合进行管理。从 Fabric v1.0 开始可以使用基于 Kafka 的交易排序服务，但是许多用户可能会发现管理 Kafka 群集的额外管理开销令人生畏或不受欢迎。

## 6. Solo

如上所述，开发测试，开发或概念证明网络时，Solo 交易排序服务是一个不错的选择。因此，它是我们 [构建你的第一个网络教程](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fbuild_network.html) 中部署的默认交易排序服务），因为从其他网络组件的角度来看，Solo 交易排序服务与更复杂的 Kafka 和 Raft 实施相同地处理交易，同时节省了管理费用，维护和升级多个节点和群集的开销。由于 Solo 交易排序服务不具有崩溃容错功能，因此永远不应将其视为生产区块链网络的可行替代方案。对于只希望从单个交易排序节点开始但将来可能会增长的网络，单个节点 Raft 群集是一个更好的选择。

## 7. Raft

有关如何配置 Raft 交易排序服务的信息，请查看我们的 [有关配置 Raft 交易排序服务的文档](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fraft_configuration.html)。

生产网络的交易排序服务选择，已建立的 Raft 协议的 Fabric 实现使用“领导者和跟随者”模型，在该模型中，领导者是在通道中的交易排序节点之间动态选举的 (这种节点集合称为“同意集”)，并且该领导者将消息复制到跟随者节点。只要有剩余的大多数交易排序节点 (称为“仲裁”)，系统就可以承受包括领导节点在内的所有节点的损失，因此 Raft 被称为“故障容错” (Crash Fault Tolerant, CFT)。换句话说，如果一个通道中有三个节点，则它可以承受一个节点的丢失 (剩下两个)。如果通道中有五个节点，则可能会丢失两个节点 (剩下三个剩余节点)。

从它们提供给网络或通道的服务的角度来看，Raft 与现有的基于 Kafka 的交易排序服务 (我们将在后面讨论)相似。它们都是采用领导者和跟随者设计的 CFT 交易排序服务。如果你是应用程序开发人员，智能合约开发人员或对端节点管理员，则不会注意到基于 Raft 和 Kafka 的交易排序服务之间的功能差异。但是，有一些主要差异值得考虑，尤其是如果你打算管理交易排序服务：

- Raft 更容易安装。尽管 Kafka 拥有许多崇拜者，但即使是那些崇拜者，他们也通常会承认部署 Kafka 集群及其 ZooKeeper 集成可能很棘手，这需要在 Kafka 基础架构和设置方面具有高水平的专业知识。此外，与 Raft 相比，Kafka 需要管理的组件更多，这意味着出错的地方更多。 Kafka 有其自己的版本，必须与你的交易排序器协调。使用 Raft，一切都将嵌入到你的交易排序节点中。
- Kafka 和 Zookeeper 并非旨在跨大型网络运行。它们被设计为 CFT，但应在紧密的主机中运行。这就是说，实际上，你需要一个组织来运行 Kafka 集群。鉴于此，在使用 Kafka (Fabric 支持) 的情况下，让交易排序节点由不同的组织运行不会带来很多分散性，因为节点将全部进入同一组织控制下的同一 Kafka 集群。借助 Raft，每个组织都可以拥有自己的交易排序节点，参与交易排序服务，从而导致系统更加分散。
- Raft 本身受支持。尽管基于 Kafka 的交易排序服务当前与 Fabric 兼容，但要求用户获取必需的镜像并了解如何自行使用 Kafka 和 ZooKeeper。同样，对 Kafka 相关问题的支持是通过 [Apache](https://links.jianshu.com/go?to=https%3A%2F%2Fkafka.apache.org%2F)，Kafka 的开源开发人员，而不是 Hyperledger Fabric 来进行的。另一方面，已经开发了 Fabric Raft 实现，并将在 Fabric 开发人员社区及其支持机构中提供支持。
- Kafka 使用服务器池 (称为 “Kafka 经纪人”)，而交易排序器组织的管理员指定他们要在特定通道上使用多少个节点，Raft 允许用户指定将哪个交易排序节点部署到哪个通道。这样，对端节点组织可以确保，如果他们还拥有一个交易排序器，则将该节点作为该通道的交易排序服务的一部分，而不是信任并依靠中央管理员来管理 Kafka 节点。
- Raft 是 Fabric 开发拜占庭式容错 (Byzantine Fault Tolerant, BFT) 交易排序服务的第一步。就像我们将看到的那样，Raft 开发中的某些决定是由此驱动的。如果你对 BFT 感兴趣，则学习如何使用 Raft 应该可以简化过渡过程。

注意：与 Solo 和 Kafka 相似，Raft 交易排序服务在收到回执后会丢失交易。例如，如果领导者大约在追随者提供收据确认的同时崩溃。因此，应用程序客户端无论如何都应在对端节点上侦听交易提交事件 (以检查交易的有效性)，但应格外小心，以确保客户端也能容忍超时，该超时不会在配置的时间范围内落实交易。取决于应用程序，可能希望在这种超时情况下重新提交交易或收集一组新的背书。

### 7.1 Raft 概念

尽管 Raft 提供了许多与 Kafka 相同的功能 - 尽管采用了更简单易用的包装 - 但在 Kafka 的掩盖下其功能却大不相同，并向 Fabric 引入了许多新概念或对现有概念的扭曲。

日志条目 (Log Entry)。 Raft 交易排序服务中的主要工作单元是“日志条目”，此类条目的完整顺序称为“日志”。如果大多数成员 (换言之为法定人数) 同意条目及其顺序，则我们认为日志是一致的，从而使复制了各种交易排序器的日志。

同意集 (Contenter set)。交易排序节点积极参与给定通道的共识机制，并接收该通道的复制日志。这可以是所有可用节点 (在单个群集中或在组成系统通道的多个群集中)，也可以是那些节点的子集。

有限状态机 (Finite-State Machine, FSM)。 Raft 中的每个交易排序节点都有一个 FSM，并共同使用它们来确保各个交易排序节点中的日志顺序是确定性的 (以相同顺序编写)。

法定人数 (Quorum)。描述需要确认提案以允许交易排序交易的最小同意者数量。对于每个同意集，这是大多数节点。在具有五个节点的群集中，三个群集必须可用。如果由于某种原因无法达到法定数量的节点，则交易排序服务集群将无法用于通道上的读取和写入操作，并且无法提交任何新日志。

领导者 (Leader)。这不是一个新概念 - 正如我们所说的，Kafka 还使用了领导者 - 但至关重要的是要了解，在任何给定时间，通道的同意者集会选举一个节点作为领导者 (稍后我们将在 Raft 中描述这种情况是如何发生的）。负责人负责摄取新的日志条目，将它们复制到跟随者交易排序节点，并管理何时将条目视为已提交。这不是交易排序器的特殊类型。这只是交易排序器可能在某些时候扮演的角色，而在其他情况下 (视情况而定) 则没有。

跟随者 (Follower)。同样，这不是一个新概念，但了解跟随者的关键是跟随者从领导者那里接收日志并确定性地复制它们，以确保日志保持一致。正如我们在领导者选举部分中所看到的那样，跟随者还会从领导者那里收到“心跳”消息。如果领导者在可配置的时间内停止发送这些消息，则跟随者将发起领导者选举，其中一个将被选举为新领导者。

### 7.2 交易流程中的 Raft

每个通道都在 Raft 协议的单独实例上运行，该协议允许每个实例选举不同的领导者。在群集由不同组织控制的交易排序节点组成的用例中，此配置还允许进一步分散服务。尽管所有 Raft 节点都必须是系统通道的一部分，但不一定必须是所有应用程序通道的一部分。通道创建者 (和通道管理员) 可以选择可用交易排序器的子集，并根据需要添加或移除交易排序节点 (只要一次仅添加或移除一个节点)。

尽管此配置以冗余心跳消息和 goroutine 的形式创建了更多开销，但为 BFT 奠定了必要的基础。

在 Raft 中，交易 (以提议或配置更新的形式) 由接收交易的交易排序节点自动路由到该通道的当前负责人。这意味着对端节点和应用程序不需要在任何特定时间知道谁是领导者节点。仅交易排序节点需要知道。

完成交易排序器验证检查后，将按照我们的交易流程第二阶段中的说明对交易进行排序，打包，成区块同意和分发。

### 7.3 体系结构注解

###### Raft 中领导人选举的工作方式

尽管选举领导者的过程是在交易排序器的内部流程中进行的，但值得注意的是该流程是如何进行的。

Raft 节点始终处于以下三种状态之一：跟随者，候选者或领导者。所有节点最初都是作为跟随者开始的。在这种状态下，他们可以接受来自领导者的日志条目 (如果已当选)，或为领导者投票。如果在设定的时间段内 (例如，五秒钟) 未接收到日志条目或心跳，则节点会自动升级为候选状态。在候选状态下，节点向其他节点请求投票。如果候选人获得法定人数的选票，则将其晋升为领导人。领导者必须接受新的日志条目并将其复制到跟随者。

有关领导者选举过程工作方式的直观图示，请查看 [The Secret Lives of Data](https://links.jianshu.com/go?to=http%3A%2F%2Fthesecretlivesofdata.com%2Fraft%2F)。

###### 快照

如果交易排序节点出现故障，它如何获取重新启动时丢失的日志？

尽管可以无限期地保留所有日志，但是为了节省磁盘空间，Raft 使用了一个称为“快照”的过程，用户可以在其中定义将在日志中保留多少字节的数据。此数据量将符合一定数量的区块 (取决于区块中的数据量。请注意，快照中仅存储完整的区块)。

例如，假设滞后的副本 R1 刚刚重新连接到网络。它的最新区块是 100。领导者 L 在区块 196，并配置为以在这种情况下表示 20 个区块的数据量进行快照。因此，R1 将从 L 接收区块 180，然后发出对区块 101 至 180 的传递请求。然后，将通过常规 Raft 协议将区块 180 至 196 复制到 R1。

### 7.4 Kafka

Fabric 支持的另一个崩溃容错交易排序服务是对 Kafka 分布式流平台的改编，可以用作交易排序节点的集群。你可以在 [Apache Kafka](https://links.jianshu.com/go?to=https%3A%2F%2Fkafka.apache.org%2Fintro) 网站上了解有关 Kafka 的更多信息，但在较高层次上，Kafka 使用与 Raft 相同的概念“领导者和跟随者”配置，在该配置中，交易 (被 Kafka 称为“消息”) 从领导者复制而来。节点到跟随者节点。如果领导者节点发生故障，则跟随者之一将成为领导者，交易排序可以继续进行，从而确保了容错能力，就像 Raft 一样。

Kafka 集群的管理，包括任务的协调，集群成员，访问控制和控制器选举等，均由 ZooKeeper 集成及其相关 API 进行处理。

众所周知，Kafka 集群和 ZooKeeper 集成很难设置，因此我们的文档假定你对 Kafka 和 ZooKeeper 有一定的了解。如果你决定在没有专业知识的情况下使用 Kafka，则在尝试使用基于 Kafka 的交易排序服务之前，至少应完成[《 Kafka快速入门》](https://links.jianshu.com/go?to=https%3A%2F%2Fkafka.apache.org%2Fquickstart) 指南的前六个步骤。你也可以查阅此样本配置文件，以简要了解 Kafka 和 ZooKeeper 的合理默认值。

要了解如何启动基于 Kafka 的交易排序服务，请查看 [我们有关 Kafka 的文档](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fkafka.html)。

## Reference

- Docs » Key Concepts » The Ordering Service, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/orderer/ordering_service.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Forderer%2Fordering_service.html)
- Docs » Key Concepts » Identity, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/identity/identity.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fidentity%2Fidentity.html)
- Docs » Key Concepts » Membership, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/membership/membership.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fmembership%2Fmembership.html)
- Docs » Key Concepts » Peers, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/peers/peers.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fpeers%2Fpeers.html)
- Docs » Architecture Reference » Gossip data dissemination protocol, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/gossip.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fgossip.html)
- Docs » Operations Guides » Setting up an ordering node, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/orderer_deploy.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Forderer_deploy.html)
- Docs » Tutorials » Building Your First Network, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/build_network.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fbuild_network.html)
- Docs » Operations Guides » Configuring and operating a Raft ordering service, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/raft_configuration.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fraft_configuration.html)
- [https://kafka.apache.org/](https://links.jianshu.com/go?to=https%3A%2F%2Fkafka.apache.org%2F)
- The Secret Lives of Data, [http://thesecretlivesofdata.com/raft/](https://links.jianshu.com/go?to=http%3A%2F%2Fthesecretlivesofdata.com%2Fraft%2F)
- Docs » Operations Guides » Bringing up a Kafka-based Ordering Service, [https://hyperledger-fabric.readthedocs.io/en/release-1.4/kafka.html](https://links.jianshu.com/go?to=https%3A%2F%2Fhyperledger-fabric.readthedocs.io%2Fen%2Frelease-1.4%2Fkafka.html)

## 项目源代码

项目源代码会逐步上传到 Github，地址为 [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)。

## Contributor

1. Windstamp, [https://github.com/windstamp](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fwindstamp)



0人点赞



[hl.fabric.doc]()





作者：furnace
链接：https://www.jianshu.com/p/77bfe2aca32d
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。























