# SportScrypt--去中心化的体育赛事投注平台

[![img](https://upload.jianshu.io/users/upload_avatars/11095853/cd9773c8-f92f-4ea0-a968-c4ceab066a7b.jpg?imageMogr2/auto-orient/strip|imageView2/1/w/96/h/96/format/webp)](https://www.jianshu.com/u/be70cbe10bb2)

[区块链技术布道者](https://www.jianshu.com/u/be70cbe10bb2)关注

2018.09.26 17:45:18字数 8,031阅读 486

# 1、Introduction

预测市场的设计者必须做出很多设计和实施来决定平台的最终效用。对于SportCrypt我们相信我们已经在简单、有效和具有特色之间取得了平衡。本文档解释了SportCrypt平台技术层面的设计和实现。

## 1.1、Sports Betting 体育赛事投注

尽管SportCrypt合约不是特定于体育设计的，体育只是我们最初的焦点所在。但是我们认为，体育赛事对于我们的工具来说具有以下优势：

- 体育竞猜是一个还没有被现有系统充分服务的庞大市场。当前的系统正遭受着不可信任、缓慢、昂贵的支付处理、监管障碍、高交易佣金、重大的交易对手风险等问题。我们相信在分布式应用里面，所有这些问题都可以被解决。
- 大部分体育赛事的结果是客观事实。这些数据都可以通过观看电视或者官方发布的数据进行确认。正因为这样，分布式的预言机不像在其他应用中那么重要。由于任何对官方结果的修改或误导都会直接对交易所产生很明显的影响，因此交易所是没有动机去操作的。

## 1.2、ETH-Denominated ETH计价

在SportCrypt上，所有的交易都是直接以ETH交易的。不需要任何其他的tokens。

因为ETH是迄今为止以太坊网络里流动性最好的资产，满足条件的交易会被记入以太坊区块，因此我们认为添加一个新的指定用途的tokens是没有必要的。用户不仅需要关心怎么样、从哪里获取这些tokens，而且需要支付市场价差和交易费用。

尽管获取和使用ETH对于非技术人员有一些障碍。然而，从统计学上来说，和其他的使投注者不得不在不同平台跳来跳去才能进行投注，SportCrypt平台还是好多了。事实上，整个行业以体育竞猜为目的资金系统建立在一种非正统的方式上。

## 1.3、Custody of Funds 资金的监管

使用智能合约进行投注的主要优势就是解决了对手方风险。假定合约不存在bug，交易者甚至不能操作账户余额或者锁定在交易里的资金。同样也不能冻结资金或者撤回交易。唯一的恶意交易的方式是误导赛事的结果，这样的话，所有的参与者都会被影响，同时会有无可争辩的证据被记录在区块链上。我们将会在Oracles的章节进一步讨论这个话题。

无论由于什么原因SportCrypt平台消失了，或者没有终结一个赛事，资金将会在一个确定的等待期后自动返还。详细介绍参考下面的资金恢复章节。

由于平台本身没有能力逆转或取消任何的行为，因此平台也必然不能冻结资金或者插手交易。所有交易都会执行完毕。然而，我们要维护自己良好的商业信誉。任何由于错误感到遭受损失的用户都请及时联系我们。

## 1.4、Fees。费率

我们最初的计划是不收取交易费用或者佣金。参与到SportCrypt上所需要的唯一的费用是以太坊平台需要的gas费用。查看下面关于效用的章节计算大概需要的gas消耗。

由于交易没有扣除费用和佣金，因此在SportCrypt上用户之间的交换完全是零和的。假设这具有积极合法的效果。我们相信通过SportCrypt进行的投注和两个人私下进行单独的投注具有同样的特性。

## 1.5、Fixed Odds 固定赔率

所有在SportCrypt上的交易均为固定赔率。这意味着所有交易的赔率都是事先被知晓并达成一致的，确立以后不可被改变。当然，后续的交易可以在新的赔率下交易。在之前的交易相同的方向上增加新的交易可以增加或者降低头寸的平均成本，增加相反方向上的头寸可以全部或部分关闭已经在盈利或者受损的头寸，在这种情况下，账户的余额变动会被马上记录。

固定赔率投注和派利分成法投注不同，派利分成法投注是一个将所有投注放在一个池子里，因此赔率只有投注结束才可以计算得知(通常是在赛事开始之前)。

## 1.6 In-game Trading 游戏中交易

在比赛结束之前的任何时刻都可以提供订单进行交易。允许用户在赛事之前，赛事之中添加新的头寸或者退出现有的头寸。不想等待最终结果的用户甚至可以在赛事完成之后进行交易直接获得他们的余额(需要一点成本)。

我们将这个特性视作SportCrypt很重要的一个方面。可以在中场休息的时候进行交易，甚至在进行中交易，使得交易过程变得更加刺激。然而，由于以太坊分布式共识的实现特性，某些注意事项还是需要考虑的。当有两个或者更多的有冲突的交易被广播到网络上时，哪一个会被矿工首先记账是不确定的。例如，当意外发生时，拥有一个比较好的订单列表的做市商可能会尝试取消订单，同时可能会有一个或者更多的参与者尝试对它交易。交易是否达成，谁的交易会被达成都是不确定的。

更进一步，不仅交易的顺序是不确定的，还有可能受到gas价格的影响，这为交易增加了一个维度，将会对复杂的交易者更有吸引力。那些不希望在他们的交易中受到gas价格影响的交易者，建议限制游戏中交易在中场休息、超时、电视暂停时等，并且谨慎使用订单到期参数。

## 1.7、Partial(局部) Trades 部分交易

订单可以提供任意数量的ETH交易，订单可以被一个或多个用户完成或部分完成。一旦交易建立，交易的双方都拥有头寸，这意味着一方将在一个结果下获利，另一方将在另一个结果下获利。然而，如果有任何一方想要退出头寸，头寸将会部分或全部卖给想要进入市场的参与者，以不同的价格(具有盈利或者损失)或者以相同的价格(没有盈利或损失)。

## 1.8、Collateral 抵押

为了进入交易，用户需要在SportCrypt合约有余额。余额与以太坊的公钥绑定，智能合约会确保只有余额相符的用户可以选择撤回或者交易。

在SportCrypt上的抵押相对于其他平台需要更加灵活，在设计上考虑到了做市商。创建订单不需要用户的账户增加余额。实际上用户基于同一账户余额可以创建很多订单。只有当交易执行的时候资金才会被锁定。如果余额减少影响到了其他未执行订单的偿付，这些订单会被自动的减少或取消。

由于用户拥有余额只是为了进入交易，事实上余额为0的时候也可以创建订单。除非用户抵押资金、宣称他们赢得了投注，或者在不同的赛事中关闭已经存在的头寸，否则订单不会对任何人可见。这时，订单会出现在订单列表里。

另外，现有头寸相反方向的交易可以以现有头寸作为抵押。正因为此，给定一个事件的头寸和一个空余额的账户，订单也可以被创建，如果被填充头寸将部分或全部关闭。更进一步，关闭头寸的收益可以用来创建新的相反方向的头寸，甚至在同样的交易中也是可以的。

# 2、Matchs 赛事

在对一个赛事进行交易之前，赛事的细节将会被审查。这些细节代表着特定赛事的规则，应该在交易之前充分理解。

## 2.1、Match Details

赛事在SportCrypt上的发布是以UTF-8编码的JSON格式发布的。下面是一个例子：



```json
{

    "type": "sports/nfl/game",

    "event": {

        "kickoff": "1509296400",

        "spread": "-2.5",

        "team1": "OAK",

        "team2": "BUF"

    },

    "contractAddr": "b9fea0142cd54bd0a8238cba4a286f5a1a261692",

    "nonce": "ul36TwZFutiL9nTlpHkMV5",

    "cancelPrice": "50",

    "recoveryWeeks": "12"

}
```

- type: type用来对赛事进行分类分层，用“/”分隔。
- event: 嵌套的对象包含了赛事的特性。
- – kickoff: The unix timestamp of when the match is scheduled to begin.
- – name: The special name of the match, for example “Super Bowl” (op-tional).
- – spread: The point spread for this match (explained below).
- – team1: The short-form identifier(标示) for the visiting-team of this match.
- – team2: The short-form identifier for the home-team of this match.
- contractAddr: 合约地址
- nonce: 一个阻止预测ID和创建其他相同赛事的随机值
- cancelPrice: 如果由于赛事被取消而导致赛事的结果没法确定，这是赛事的最终的价格(通常通过交易，也可能通过资金恢复)
- recoveryWeeks: 第一次交易过后的数周没有结果资金可以以取消价的价格恢复。

所有的值为字符串的形式，包括点差、价格、时间戳和周数。以太坊地址以小写的十六进制存放，没有0x的前缀。

## 2.2、Match IDs

为了计算256位的赛事ID，赛事详情的JSON以字母表顺序排序，压缩(不需要的空格被移除)，然后用keccak256哈希算法计算哈希值。

将结果hash截去两个字节，然后拼接上两个uint8字节：第一个字节编码了赛事的cancelPrice，第二个字节编码了recoveryWeeks 参数。

结果就是在区块链上和订单签名时使用的match ID。

展示的时候通常以十六进制的字符串展示，例如：8c1705e212fd2d369e57e0012fa1e3083705cd71a871af21f6ce6230cfcd320c 前六十个字符代表从match details的hash中截取的30个字节。接下来的两个字节分别代表cancelPrice为50，recoveryWeeks为12。

在创建订单或交易之前，用户应该确认match ID被正确的计算了。我们的前端UI实现自动做了这件事。

## 2.3、Point Spreads

根据赛事的type，可能会有相关的point spread。这是一个或正或负的数，在评估赛事的结果之前添加到主场的最终分数上。这样一来，即使不同技能层次的比赛双方，也可以以相近的赔率进行交易，同时平局也不可能发生，因为在大部分体育赛事结果分通常为整数，而point-spreads包含分数。SportCrypt根据初始发布的拉斯维加斯或离岸线来选择point-spreads。

使用上面的例子，赛事具有形如“OAK@BUF-2.5”的简称。简称代表Oakland Raiders对战Buffalo Bills。“@”字符表示Oakland为客场，Buffalo为主场。“-2.5”表明主场的分数最后要加上-2.5。在这个例子中，对于主场是不利的，意味着，不考虑point-spreads，Buffalo被认为更可能赢得比赛。

在SportCrypt平台上，合约假定总是客场队是否赢得point spread。因此，为了确定合约的输出，将-2.5加到Buffalo的最终分数上与Oakland的得分比较。如果Oakland的分数大于最终Buffalo的，则假定为真，否则假定为假。

# 3、Prices and Odds

理解价格和赔率对于在SportCrypt上交易盈利是重要的，在其他的地方也是如此。同时以不同的prices/odds买入和卖出是做市商盈利的方式，准确的评估赛事的结果的可能性对比发布的价格(或者自己创造价格)是交易者的盈利方式。

在SportCrypt平台上，任何人可以是做市商或交易者，或者既是做市商又是交易者。

## 3.1、Implied Probability

在SportCrypt上每个合约的价值为0-100的整数。这与传统体育投注中的赔率相对应。它反应了对结果发生的感知到的机会，因此交易者需要冒一定的风险才能赚钱一定的金额。

赛事投注最终结束，如果合约假设为真，合约价值为100，如果结果为假则合约价值为0。在结束之前参与者可以选择0到100之间的一个值作为合约的价格。

选用0到100的范围是为了和百分比相对应。因此，这种情况下的赔率成为隐含可能赔率。在体育投注领域隐含概率最重要的推广者是现在已经不存在的网站TradeSports.com，因此在本介绍中我们把隐含概率定价的二进制选择作为“TradeSport Model”。

## 3.2、Bid-Ask Spread

最低的要价和最高的投标价之间的差称为买卖差(bid-ask spread)。这个买卖差和前面讨论的点差(point spread)没有关系。

因为做市商总是尝试以低价买入高价卖出，他们更倾向于较大的买卖差。反过来，交易者由于要支付市场价，他们在执行交易的时候更倾向于较小的买卖差。

热门交易事件的买卖差通常比冷门事件要小。这是因为做市商通过提供更小的买卖价差竞争，交易者通常创建略比做市商略低的价格订单，希望通过直接卖给对方交易者的方法避免支付买卖价差。

在一个中心化的交易场景中，通常不会出现负的价差(买价大于卖价)因为这样的交叉订单会被直接成交。然而，在SportCrypt上，负的买卖价差是可能存在的，因为订单环节和执行环节是解耦的。负的买卖价差将是罕见的，因为这里面有套利的机会。一些投机者可以同时以较低的报价买入以较高的卖价卖出从而获得差价。以太坊具有一个有用的特性可以使这种情况原子的发生，两笔交易都被执行或者两笔交易都不被执行。

和存在负的买卖价差一样，SportCrypt上也会存在为0的买卖价差。这时候套利者就没有了套利的机会，这时的买卖价差可以看做是交易执行所需要的gas。

## 3.3、Amount at Risk

创建交易的时候，参与者同意0到100(不包含100)的一个整数作为交易的价格。他们在交易中所面临的风险大小取决于这个价格，可以用下面的公式进行计算：

![img](https://upload-images.jianshu.io/upload_images/11095853-5c95a29af955de53.png!thumbnail?imageMogr2/auto-orient/strip|imageView2/2/w/1130/format/webp)

## 3.4、Finalization Prices

![img](https://upload-images.jianshu.io/upload_images/11095853-0602aa0be278f4b0.png!thumbnail?imageMogr2/auto-orient/strip|imageView2/2/w/1182/format/webp)

最终价格通常是100或者0。这意味着卖方或者卖方最后会一方获得全部的资金，另一方什么也得不到。

然而，在极少数的情况下，比赛没有确定的结果，将会以cancelPrice作为最终价格。cancelPrice是赛事合约详情的一项，应该在交易模型中考虑。通常情况下cancelPrice为50，然而如果预测的合约的最初价值有很大的不同cancelPrice可能不一样。“money-line”的赛事就是这样的例子(没有point-spreads的赛事)。

## 3.5、Expected Value

有了隐含概率，很明显最好的操作就是低价买高价卖。

根据结果输出的可能性给出一个精确的估计，我们就可以在低于此估计价的时候买入或以高于此估计价卖出。

![img](https://upload-images.jianshu.io/upload_images/11095853-64b19c7df1390467.png!thumbnail?imageMogr2/auto-orient/strip|imageView2/2/w/1132/format/webp)

使用期望值，标准的资金管理技术，例如凯利公式就可以使用了。

## 3.6、Odds Conversion Examples

隐含概率相对于其他的赔率描述有几个优势。然而，当前的体育赛事投注者习惯于各种各样的形式，因此我们会做一个简单的对于赔率的转变。

![img](https://upload-images.jianshu.io/upload_images/11095853-3baeac2785b42b8a.png!thumbnail?imageMogr2/auto-orient/strip|imageView2/2/w/1184/format/webp)

# 4、Off-Chain Mechanics

SportCrypt使用链上链下混合的模式。赛事详情，赛事最终价格，订单的建立都是在链下进行，交易的发生是在链上进行。我们称这种模式为“EtherDelta Model”。

我们也在评估修改这种模式为集中匹配通过链上交易进行设置。我们叫这种模式为“IDEX model”，在本节中我们不深入讨论。

## 4.1、Match Creation

当赛事用来提供交易，它们被添加到我们的后台系统并通过websocket将它们推送给所有连接的客户端。创造赛事不需要任何的链上交易。这是SportCrypt的一个主要的可扩展优势，它可以减少交易的花费，允许我们同时进行多个赛事而不用担心以太坊gas的花费和区块挖矿的等待。

## 4.2、Match Finalization

同样，终结一个赛事也不需要链上行为。交易的操作者将会注册一个标明对应ID在某一个确定价格终结的信息。和订单类似，签名指定了特定的合约地址。当参与者领取奖金时，参与者会将签名信息和签名一起发送到区块链。如果还没有其他人终结合约，签名生效合约终结。结果是第一个声明获得奖励的参与者需要额外支付10k的gas费用。但是，我们觉得这是值得的，因为这可以减少为了终结合约需要进行的交换的操作开销。

## 4.3、Orders

在选定一个赛事进行交易之后，市场参与者将创建并签名的订单通过websocket上传到我们的链下订单列表。

The orders are tightly packed into 3 uint256 values:这些订单被压缩进3个uint256的值：



```csharp
[0]: 32-byte match ID

[1]: 32-byte amount in wei

[2]:  5-byte expiry

      5-byte nonce

      1-byte price

      1-byte direction

     20-byte address
```

- match ID: The hashed match details, as described above.
- amount: The maximum total amount at risk authorized by the order creator.
- expiry: Unix timestamp after which the order ceases to be valid.
- nonce: A random value which allows multiple otherwise(在其他方面) identical(相同的) orders to be issued, and prevents order ID prediction.
- price: A value from 1-99 representing the price authorized by the order creator.
- direction: Either 0 for a sell order, or 1 for a buy order.
- address: The address of the order creator.

由于Solidity堆栈深度的限制和节约gas的花费对数据的压缩是必要的。SportCrypt96个字节的订单已经是比较节约了。

订单ID是合约地址的hash链接上这三个uint256的值。

在提交到订单列表之后，订单和签名被广播到所有监听对应赛事的连接的客户端。

当一个市场参与者选择执行交易的时候，订单和签名会被发送到区块链执行。智能合约计算订单ID并且验证订单创建这的地址是否满足提供的签名。

# 5、On-Chain Mechanics

本章介绍智能合约在一个理想化的数学机制下的核心交易机制大纲。实际上实施起来的结构会有所不同，以免会有四舍五入的损失、实现高效的不变的断言，并优化gas的使用。描述不具有权威性，以智能合约代码为准。

## 5.1、Trading

交易函数接收一个订单，订单创建者的签名和一个金额。所有除了金额以外的交易细节都由订单参数确定。这个金额是交易将要接受的最大风险数额，智能合约会以尽可能接近该数额的风险价值位置创建交易。

在正常操作过程中，智能合约被设计为不会抛出异常，因为这会消耗完所有提供的gas。相反，交易会以0金额的形式完成，在这种情况下LogTradeError会将日志记录下来。

赛事结束之后的交易将会被智能合约判定为无效。

## 5.2、Positions

头寸是交易的结果，它们保存着合约一旦结束，哪一个账户有资格获得资金。负的头寸值代表着空头，正值代表着多头。头寸代表着声明的资金总额，不是风险的多空数额。

## 5.3、Effective Balances

为了创建订单，用户通常需要在他们的独立账户里有充足的资金，与它们面临的风险相匹配。然而如果用户建立交易的赛事，用户已经拥有头寸，交易的方向为所拥有头寸的相反方向，头寸本身也可以作为抵押。用这种方式，用户可以关闭现有的头寸。

![img](https://upload-images.jianshu.io/upload_images/11095853-16cd7c506cfd7be8.png!thumbnail?imageMogr2/auto-orient/strip|imageView2/2/w/1188/format/webp)

有效余额双方都要计算。这些值随后被用来决定交易中可以被使用的金额，包括现存的订单数量，交易的最大金额，和订单的价格。

## 5.4、Position Updates

当交易创建，参与账户的头寸通过增加或减少总交易金额，总金额来进行更新，公式如下

![img](https://upload-images.jianshu.io/upload_images/11095853-03a714ace7d59efa.png!thumbnail?imageMogr2/auto-orient/strip|imageView2/2/w/1192/format/webp)

在最终结束之后，等式就不再成立，因为胜利者会在声明的过程中将它们重置为0。失败的参与者已经没有理由支付gas费用来设置它们的头寸为0。

## 5.5、Balance Updates

![img](https://upload-images.jianshu.io/upload_images/11095853-e27af98d93987992.png!thumbnail?imageMogr2/auto-orient/strip|imageView2/2/w/1178/format/webp)

这些公式给人的感觉就是头寸增加借记账户余额头寸减少贷记账户余额。头寸被出售或者购买，借记或贷记的大小依赖于交易达成的价格。

如果没有预先存在的头寸，两方余额都要减少对应的金额用来弥补交易所需要承担的风险。然而如果交易双方或一方拥有头寸，且它们朝着拥有头寸相反方向交易，余额将会增加。事情并不总是如此，因为一个单笔交易可以关闭现有头寸并额外增加相反方向的头寸。

## 5.6、Order Amount Decrease

当交易进行，对应订单ID的填充金额将随着交易确认的风险大小增加！

填充金额和订单金额之间的差被用于随后的交易，当决定订单的剩余多少。

当订单取消时，填充金额上升到和订单金额相同。

## 5.7、Order Cancellation

和EtherDelta相同，订单取消必须通过智能的cancelOrder函数。不幸的是，这意味着取消订单必须支付gas并且取消操作不是马上完成的。

很多情况下这可以通过短期订单过期值来避免。一旦区块的时间戳超过了订单过期的时间戳，订单就有效的取消了不需要使用交易取消订单了。

## 5.8、Funds Recovery

正如在赛事ID章节解释的那样，cancelPrice和recoveryWeeks适用于智能合约，是因为它们直接嵌入到了matchID中。这种不希望看到的带内信号和hash强度的减少对于实现本章节描述的资金恢复特性是必须的。

注意到我们截断了256位的Hash值，留下了240位仍然是很安全的。

如果SportCrypt没有终结比赛，用户需要等到第一笔交易过后的recoveryWeeks周，赛事过期。在那个时间点上，每个用户会调用智能合约的recoverFunds函数。智能合约将会以cancelPrice的价格终结赛事，允许用户赎回他们的资金。以周为单位，最大为256，给我们很好的选择颗粒度和接近五年的选择范围。相对于其他的时间单位例如月，周均匀的具有7天的市长，自公元前46年就通过了日历改革。

当然，我们希望这个功能永远不会被使用。它为用户提供了一个保证即使SportCrypt完全消失了，资金仍然可以以赛事详情里面规定好的cancelPrice价格恢复。

# 6、Efficiency

SportCrypt从一开始就设计为一个非常轻巧的项目。通过减少操作花费，我们提供的节约对我们的用户来说有重大的意义。

## 6.1、Smart Contract

正如前面讨论的那样，SportCrypt系统尽可能的设计为在链下运作：

- Match creation and finalization is done off-chain.
- Orders are created and advertised off-chain.

链上进行的操作我们以尽可能少的花费进行：

- There are no loops in contract code (except read-only views).
- Orders are tightly packed to reduce calldata gas consumption.
- The contract never throws in normal operation.

![img](https://upload-images.jianshu.io/upload_images/11095853-efcf32fe8a578713.png!thumbnail?imageMogr2/auto-orient/strip|imageView2/2/w/1166/format/webp)

尽管gas的花费是固定的，用户实际上的支付取决于他们发送交易时gas的价格。

例如最贵的操作花费了140K的gas。假定gas的价格为1Gwei，这个交易将需要0.00014ETH。在每个ETH需要300USD的汇率下，交易需要花费0.042$。

由于gas的花费独立于交易金额，摊销将会使花费的比率尽可能的降低。在这个例子中一个$5的交易需要的gas不到1%。

## 6.2、Order-Book

尽管余额，交易，订单填充和头寸都被严格记录在区块链上，我们仍然需要运行服务，最重要的是提供赛事详情和订单列表。

我们尝试最小化我们的操作成本：

- Hybrid multi-threaded/non-blocking C++ implementation.
- Communicates with clients using compressed websocket push to minimize latency and bandwidth usage.
- In-process, memory-mapped DB to store persistent data.

## 6.3、User Interface

代码的最后部分我们考虑了代码在客户端端浏览器上的运行效率。在我们前端我们使用尽量少的内存，CPU和带宽：

- Modern ES6+, React, Webpack code-base.
- Querying the smart contract with constant(常量) calls amortizes(摊销) RPC overhead by batching requests.

我们的用户接口不需要与SportCrypt交互，我们计划开发移动端和命令行应用。

# 7、Oracles

就像介绍里提到的，SportCrypt目前实现了一个中心化的预言机。尽管有人说有方式可以达到分布式的预言机，我们还不清楚这种系统的效用，如果是这样，他们必须足够依赖才可以。

对于预测市场来说体育赛事的好处在于，对于诚实的各方来说达成赛事结果的一致是直截了当的，不管是通过观看比赛本身还是依靠官方发布的分数。一般的预测市场遭遇着我们认为的“朝鲜导弹问题的影响”。在Intrade.com上，有一个关于北朝鲜是否会在特定时期发射导弹的合约。他们做了。这被广泛的报道并被业内的专家认为是真的。但是合约本身需要美国政府的确认，但是没有任何确认发布。这使得Intrade处于一种尴尬的境地:要么违反自己公布的合同规则，要么以公认的错误价格敲定合同。

朝鲜导弹问题影响着中心化和分布式的预言机，解决方案尚不清晰。对于一个中心化的预言机，预言机本身可以单方面决定结果，或者服从于其他的团体。例如，在朝鲜导弹问题的例子中，Intrade服从美国政府的决定。另一方面，分布式的预言机服从群众的智慧，也许是通过经济激励鼓励人们报道他们认为真的事情。然而，我们认为这些系统还只是停留在理论层面，尚未证实。

在SportCrypt，我们尝试建立一个有声誉可信任的预言机。如果我们错误报道赛事，用户就会掌握证据，因为它是永久性存储在区块链上的。另外，我们没办法针对不同的用户报道特定的结果：一个错误的结果价格会影响任何一个参与者。

也就是说，如果分布式式预言机证明本身是可被信任的，准时的，经济有效的，我们会整合到SportCrypt平台。

最后，由于SportCrypt的体系结构，它可能支持多组匹配，两个或者更多的预言机对一个最终价格预先达成一致才终结赛事。如果市场需要更高的预言保障，我们会在随后的日子探索这种方式。

# 8、Conclusion

- 我们相信为体育赛事建立一个分散式的预测市场已经准备就绪
- TradeSports已经证明仿照金融交易的体育博彩平台的可行性
- 比特币已经将区块链作为一个去信任化，广布全球，不可停止的货币概念

以太坊已经创建了区块链的一种实现，该实现允许资金按照定制的规则进行分配

通过SportCrypt，我们将对体育赛事交易的热情与数十年的金融市场经验以及高效的软件开发相结合。作为做市商、交易者和体育迷，我们创造了我们一直希望存在的系统，我们邀请全世界加入我们。