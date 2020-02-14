package y2bcaptions

import (
	"bytes"
	"testing"
)

//https://www.youtube.com/watch?v=Ap0huJwyT7g&list=RDAp0huJwyT7g&start_radio=1`

func TestGetSubtitleInfo(t *testing.T) {
	captions, err := GetSubtitleInfoFromId("Ap0huJwyT7g")
	if err != nil {
		t.Fatal(err)
	}
	if captions.BaseUrl == "" {
		t.Fail()
	}
	if len(captions.CaptionTrack) < 5 {
		t.Fail()
	}
}

func TestCaptions_Download(t *testing.T) {
	captions, err := GetSubtitleInfoFromId("Ap0huJwyT7g")
	if err != nil {
		t.FailNow()
	}
	buf := new(bytes.Buffer)
	captions.Download(captions.CaptionTrack[0].VssId, "", buf)
	if Ap0huJwyT7g_Srt0 != buf.String() {
		t.Fail()
	}
}


var (
	Ap0huJwyT7g_Srt0 = `0
00:00:22,500 --> 00:00:25,340
J'entendais que Blanche Neige s'enfuyait

1
00:00:26,280 --> 00:00:28,880
Et que le Petite Chaperon Rouge a peur du loup

2
00:00:29,500 --> 00:00:32,360
J'entendais que le Chapelier Fou aime bien Alice

3
00:00:32,960 --> 00:00:35,660
Et que le Vilain Petit Canard deviendra un cygne

4
00:00:36,580 --> 00:00:39,260
J'entendais que Peter Pan ne grandit jamais

5
00:00:39,840 --> 00:00:42,540
Et que Jack a une harpe magique

6
00:00:43,640 --> 00:00:46,280
J'entendais qu'il y a une Maison en Pain d'épice dans les bois

7
00:00:46,760 --> 00:00:50,240
Et que Cendrillon a perdu sa Pantoufle de Verre préférée

8
00:00:50,520 --> 00:00:52,760
Seule la rivière sage sait

9
00:00:53,680 --> 00:00:56,380
que Blanche Neige s'ennuyait trop dans le château

10
00:00:57,040 --> 00:00:59,560
que le Petit Chaperon Rouge porte une cape rouge

11
00:01:00,660 --> 00:01:03,340
qui la retient de se transformer en un loup

12
00:01:03,660 --> 00:01:09,580
Il y a toujours une rivière d'une couleur iridescente qui souffle à travers le village enchanté

13
00:01:10,900 --> 00:01:14,340
gâté par la magie sauvage

14
00:01:14,560 --> 00:01:16,900
mais se tord et tombe amoureux

15
00:01:17,880 --> 00:01:20,180
La rivière court sans fin

16
00:01:20,760 --> 00:01:23,640
coule vers un lac rempli de mémoires

17
00:01:24,800 --> 00:01:27,780
Laissons alors ce "Il était une fois"

18
00:01:27,960 --> 00:01:31,340
s'emporter vers le "Heureux pour toujours"

19
00:01:47,240 --> 00:01:49,960
J'entendais que la Belle au Bois dormant était maudite dans un sommeil profond

20
00:01:50,540 --> 00:01:53,260
que la Petite Sirène regardait depuis le palais doré

21
00:01:53,580 --> 00:01:56,380
J'entendais qu'Apollo se pencha vers le Soleil

22
00:01:56,920 --> 00:01:59,720
que les Tigres à dents de sabre tassent la plaine verte

23
00:02:01,000 --> 00:02:03,240
J'entendais que Pinocchio ne dit jamais la vérité

24
00:02:04,400 --> 00:02:06,560
que les Nains ont des boîtes remplis de bijoux

25
00:02:07,520 --> 00:02:10,240
J'entendais que l'Arbre de la vie grandit sur une falaise

26
00:02:11,200 --> 00:02:13,840
que les Chaussons Rouges dansent sans la moindre fatigue

27
00:02:14,940 --> 00:02:16,940
Seule la sage rivière sait

28
00:02:18,000 --> 00:02:21,040
que La Belle au bois dormant s'échappait de la souffrance

29
00:02:21,180 --> 00:02:24,060
que la Petite Sirène utilisait la lumière du soleil comme eyeshadow

30
00:02:24,440 --> 00:02:27,120
et embrassait des bulles

31
00:02:27,900 --> 00:02:34,020
Il y a toujours une rivière de couleur iridescente qui souffle à travers le village enchanté

32
00:02:35,160 --> 00:02:38,040
gâté par la magie sauvage

33
00:02:38,200 --> 00:02:40,700
mais se tord et tombe amoureux

34
00:02:42,100 --> 00:02:44,760
La rivière court sans fin

35
00:02:44,940 --> 00:02:47,980
coule vers le lac remplis de mémoires

36
00:02:48,120 --> 00:02:51,920
Laissons alors ce "Il était une fois"

37
00:02:52,080 --> 00:02:54,720
s'emporter vers le "Heureux pour toujours"

38
00:02:55,340 --> 00:03:01,660
Il y a toujours une rivière mystérieuse qui souffle vers le village enchanté

39
00:03:02,580 --> 00:03:05,580
divisant rêves et réalité

40
00:03:06,100 --> 00:03:08,660
les permet de se rencontrer et se fusionner

41
00:03:09,520 --> 00:03:11,880
La rivière court sans fin

42
00:03:12,760 --> 00:03:15,280
coule vers un lac remplis de mémoires

43
00:03:16,360 --> 00:03:19,040
Laissons alors ce "Il était une fois"

44
00:03:19,300 --> 00:03:22,240
s'emporter vers le "Heureux pour toujours"

45
00:03:22,660 --> 00:03:24,840
mais reste inattendu

`
)
