package y2bcaptions

import (
	"encoding/json"
	"testing"
)

func TestJson3_ToSrt(t *testing.T) {
	j3 := new(Json3)
	err := json.Unmarshal([]byte(subtitleJson3), j3)
	if err != nil {
		t.Fatal("unmarshal failed")
	}
	if string(j3.ToSrt()) != subtitleSrt {
		t.Fail()
	}
	j32 := new(Json3)
	err = json.Unmarshal([]byte(subtitleJson3_2), j32)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(j32.ToSrt()))
}

var (
	subtitleJson3 = `{
  "wireMagic": "pb3",
  "pens": [
    {
    }
  ],
  "wsWinStyles": [
    {
    }
  ],
  "wpWinPositions": [
    {
    }
  ],
  "events": [
    {
      "tStartMs": 210,
      "dDurationMs": 3830,
      "segs": [
        {
          "utf8": "Hi, I'm Vanessa from SpeakEnglishWithVanessa.com."
        }
      ]
    },
    {
      "tStartMs": 4040,
      "dDurationMs": 5090,
      "segs": [
        {
          "utf8": "Today, I have some good news and some bad\nnews."
        }
      ]
    },
    {
      "tStartMs": 9130,
      "dDurationMs": 1679,
      "segs": [
        {
          "utf8": "Let's start with the bad news."
        }
      ]
    },
    {
      "tStartMs": 10809,
      "dDurationMs": 8171,
      "segs": [
        {
          "utf8": "The bad news is that you probably spent a\nlot of time in boring, middle school and high"
        }
      ]
    },
    {
      "tStartMs": 18980,
      "dDurationMs": 1910,
      "segs": [
        {
          "utf8": "school English classes."
        }
      ]
    },
    {
      "tStartMs": 20890,
      "dDurationMs": 6250,
      "segs": [
        {
          "utf8": "You felt a lot of stress, a lot of worry about\nlearning English because it just wasn't fun."
        }
      ]
    },
    {
      "tStartMs": 27140,
      "dDurationMs": 1760,
      "segs": [
        {
          "utf8": "You had to study for test."
        }
      ]
    },
    {
      "tStartMs": 28900,
      "dDurationMs": 2980,
      "segs": [
        {
          "utf8": "You had to memorize list of vocabulary."
        }
      ]
    },
    {
      "tStartMs": 31880,
      "dDurationMs": 7740,
      "segs": [
        {
          "utf8": "This is definitely not an enjoyable or effective\nway to learn anything, especially English,"
        }
      ]
    },
    {
      "tStartMs": 39620,
      "dDurationMs": 2059,
      "segs": [
        {
          "utf8": "but I have some good news."
        }
      ]
    },
    {
      "tStartMs": 41679,
      "dDurationMs": 7500,
      "segs": [
        {
          "utf8": "Now, you can choose your own teacher, your\nown materials, your own method, the style."
        }
      ]
    },
    {
      "tStartMs": 49179,
      "dDurationMs": 6011,
      "segs": [
        {
          "utf8": "Now, you are an adult and you can choose what\nyou do with English."
        }
      ]
    },
    {
      "tStartMs": 55190,
      "dDurationMs": 6150,
      "segs": [
        {
          "utf8": "This is exactly what is the best way to learn\nwhen you choose what's the best style for"
        }
      ]
    },
    {
      "tStartMs": 61340,
      "dDurationMs": 1059,
      "segs": [
        {
          "utf8": "you."
        }
      ]
    },
    {
      "tStartMs": 62399,
      "dDurationMs": 6571,
      "segs": [
        {
          "utf8": "On Skype, I've taught over 2000 English lessons\nto English learners around the world."
        }
      ]
    },
    {
      "tStartMs": 68970,
      "dDurationMs": 1650,
      "segs": [
        {
          "utf8": "You know what?"
        }
      ]
    },
    {
      "tStartMs": 70620,
      "dDurationMs": 4340,
      "segs": [
        {
          "utf8": "Almost everyone makes the same two mistakes."
        }
      ]
    },
    {
      "tStartMs": 74960,
      "dDurationMs": 1780,
      "segs": [
        {
          "utf8": "What are those two mistakes?"
        }
      ]
    },
    {
      "tStartMs": 76740,
      "dDurationMs": 5680,
      "segs": [
        {
          "utf8": "The first one is that they often use textbook\nexpressions."
        }
      ]
    },
    {
      "tStartMs": 82420,
      "dDurationMs": 6010,
      "segs": [
        {
          "utf8": "These expressions are ones that you learned\nin the textbook, but native speakers don't"
        }
      ]
    },
    {
      "tStartMs": 88430,
      "dDurationMs": 2850,
      "segs": [
        {
          "utf8": "really use them in daily conversation."
        }
      ]
    },
    {
      "tStartMs": 91280,
      "dDurationMs": 6390,
      "segs": [
        {
          "utf8": "Native speakers can understand what they mean,\nbut it sounds a little bit robotic or a little"
        }
      ]
    },
    {
      "tStartMs": 97670,
      "dDurationMs": 3470,
      "segs": [
        {
          "utf8": "bit like a textbook, because that's where\nyou learned it."
        }
      ]
    },
    {
      "tStartMs": 101140,
      "dDurationMs": 2839,
      "segs": [
        {
          "utf8": "I know you don't want to speak textbook English."
        }
      ]
    },
    {
      "tStartMs": 103979,
      "dDurationMs": 6680,
      "segs": [
        {
          "utf8": "I don't want you to speak textbook English,\nso you need to do something different."
        }
      ]
    },
    {
      "tStartMs": 110659,
      "dDurationMs": 2441,
      "segs": [
        {
          "utf8": "Make a huge change."
        }
      ]
    },
    {
      "tStartMs": 113100,
      "dDurationMs": 2659,
      "segs": [
        {
          "utf8": "Let's talk about the second mistake."
        }
      ]
    },
    {
      "tStartMs": 115759,
      "dDurationMs": 6761,
      "segs": [
        {
          "utf8": "The second mistake that almost all of my Skype\nstudents make is there are a lot of daily"
        }
      ]
    },
    {
      "tStartMs": 122520,
      "dDurationMs": 6230,
      "segs": [
        {
          "utf8": "common expressions that native speakers use\nbut they don't use, they don't know about"
        }
      ]
    },
    {
      "tStartMs": 128750,
      "dDurationMs": 4210,
      "segs": [
        {
          "utf8": "and maybe they don't even understand until\nI have to explain it."
        }
      ]
    },
    {
      "tStartMs": 132960,
      "dDurationMs": 7610,
      "segs": [
        {
          "utf8": "These expressions you definitely need to know\nif you want to understand movies, TV shows,"
        }
      ]
    },
    {
      "tStartMs": 140570,
      "dDurationMs": 5760,
      "segs": [
        {
          "utf8": "communicate in a natural way, and so that\nyou can feel more confident because you know"
        }
      ]
    },
    {
      "tStartMs": 146330,
      "dDurationMs": 3530,
      "segs": [
        {
          "utf8": "that you are using natural English."
        }
      ]
    },
    {
      "tStartMs": 149860,
      "dDurationMs": 4280,
      "segs": [
        {
          "utf8": "That's why I've created this course, 50 Natural\nEnglish Expressions."
        }
      ]
    },
    {
      "tStartMs": 154140,
      "dDurationMs": 7070,
      "segs": [
        {
          "utf8": "I've made a list of all of the English expressions\nthat my English students don't know or don't"
        }
      ]
    },
    {
      "tStartMs": 161210,
      "dDurationMs": 5620,
      "segs": [
        {
          "utf8": "use, and I made them into lesson files especially\nfor you."
        }
      ]
    },
    {
      "tStartMs": 166830,
      "dDurationMs": 2300,
      "segs": [
        {
          "utf8": "I know that a lot of you are busy."
        }
      ]
    },
    {
      "tStartMs": 169130,
      "dDurationMs": 4400,
      "segs": [
        {
          "utf8": "You have a lot of things going on in your\nlife, so these lessons you can download."
        }
      ]
    },
    {
      "tStartMs": 173530,
      "dDurationMs": 6390,
      "segs": [
        {
          "utf8": "You can listen to them wherever you are and\nyou can use it and learn it and have fun."
        }
      ]
    },
    {
      "tStartMs": 179920,
      "dDurationMs": 5370,
      "segs": [
        {
          "utf8": "I love hearing what members of this course\nhave to say about their experience."
        }
      ]
    },
    {
      "tStartMs": 185290,
      "dDurationMs": 5300,
      "segs": [
        {
          "utf8": "I think that you shouldn't only listen to\nwhat I'm saying about it, but also listen"
        }
      ]
    },
    {
      "tStartMs": 190590,
      "dDurationMs": 4600,
      "segs": [
        {
          "utf8": "to other English learners, because they know\nyour situation exactly."
        }
      ]
    },
    {
      "tStartMs": 195190,
      "dDurationMs": 6130,
      "segs": [
        {
          "utf8": "They know what it's like to feel frustrated,\nto feel upset about your level and they also"
        }
      ]
    },
    {
      "tStartMs": 201320,
      "dDurationMs": 3040,
      "segs": [
        {
          "utf8": "want to improve their English just like you."
        }
      ]
    },
    {
      "tStartMs": 204360,
      "dDurationMs": 3960,
      "segs": [
        {
          "utf8": "Let's take a look at two people who are part\nof this course."
        }
      ]
    },
    {
      "tStartMs": 208320,
      "dDurationMs": 2780,
      "segs": [
        {
          "utf8": "Two members and see what they have to say."
        }
      ]
    },
    {
      "tStartMs": 211100,
      "dDurationMs": 5370,
      "segs": [
        {
          "utf8": "I'm going to read what Vera from Brazil had\nto say about her experience in the course,"
        }
      ]
    },
    {
      "tStartMs": 216470,
      "dDurationMs": 2090,
      "segs": [
        {
          "utf8": "50 Natural English Expressions."
        }
      ]
    },
    {
      "tStartMs": 218560,
      "dDurationMs": 5270,
      "segs": [
        {
          "utf8": "She said, \"Well, for me, the way the course\nwas set up is perfect because we have the"
        }
      ]
    },
    {
      "tStartMs": 223830,
      "dDurationMs": 6290,
      "segs": [
        {
          "utf8": "explanation of expressions, examples of how\nto use them, a story to understand and fix"
        }
      ]
    },
    {
      "tStartMs": 230120,
      "dDurationMs": 5570,
      "segs": [
        {
          "utf8": "what you've learned and still have the opportunity\nto use them making videos or posting comments"
        }
      ]
    },
    {
      "tStartMs": 235690,
      "dDurationMs": 1310,
      "segs": [
        {
          "utf8": "in the Facebook group.\""
        }
      ]
    },
    {
      "tStartMs": 237000,
      "dDurationMs": 1000,
      "segs": [
        {
          "utf8": "Thanks Vera."
        }
      ]
    },
    {
      "tStartMs": 238000,
      "dDurationMs": 1810,
      "segs": [
        {
          "utf8": "That's a really honest comment."
        }
      ]
    },
    {
      "tStartMs": 239810,
      "dDurationMs": 3100,
      "segs": [
        {
          "utf8": "Let's take a look at what Jean-Jacques from\nFrance has to say."
        }
      ]
    },
    {
      "tStartMs": 242910,
      "dDurationMs": 5510,
      "segs": [
        {
          "utf8": "He said, \"Even if you are a beginner or an\nintermediate English learner, Teacher Vanessa's"
        }
      ]
    },
    {
      "tStartMs": 248420,
      "dDurationMs": 1370,
      "segs": [
        {
          "utf8": "method is useful."
        }
      ]
    },
    {
      "tStartMs": 249790,
      "dDurationMs": 2390,
      "segs": [
        {
          "utf8": "You will learn with simplicity."
        }
      ]
    },
    {
      "tStartMs": 252180,
      "dDurationMs": 5519,
      "segs": [
        {
          "utf8": "If you want to accelerate the process of learning,\nacquire the 50 Natural English Expressions"
        }
      ]
    },
    {
      "tStartMs": 257699,
      "dDurationMs": 1000,
      "segs": [
        {
          "utf8": "Course.\""
        }
      ]
    },
    {
      "tStartMs": 258699,
      "dDurationMs": 3981,
      "segs": [
        {
          "utf8": "Thanks Jean-Jacques, I think that's really\na good point, simplicity."
        }
      ]
    },
    {
      "tStartMs": 262680,
      "dDurationMs": 6070,
      "segs": [
        {
          "utf8": "You don't need to study really complex grammar\nrules or feel frustrated and upset about complex"
        }
      ]
    },
    {
      "tStartMs": 268750,
      "dDurationMs": 5750,
      "segs": [
        {
          "utf8": "things in English, you can learn English with\nsimplicity and not feel stressed and enjoy"
        }
      ]
    },
    {
      "tStartMs": 274500,
      "dDurationMs": 1670,
      "segs": [
        {
          "utf8": "it."
        }
      ]
    },
    {
      "tStartMs": 276170,
      "dDurationMs": 4190,
      "segs": [
        {
          "utf8": "That's why I've created this course, 50 Natural\nEnglish Expressions."
        }
      ]
    },
    {
      "tStartMs": 280360,
      "dDurationMs": 7269,
      "segs": [
        {
          "utf8": "I want you to be able to use these expressions\nthat are not textbook expressions, and that"
        }
      ]
    },
    {
      "tStartMs": 287629,
      "dDurationMs": 2831,
      "segs": [
        {
          "utf8": "native speakers use all the time."
        }
      ]
    },
    {
      "tStartMs": 290460,
      "dDurationMs": 5729,
      "segs": [
        {
          "utf8": "You are definitely going to hear these expressions\nin movies, TV shows, in natural conversations"
        }
      ]
    },
    {
      "tStartMs": 296189,
      "dDurationMs": 5051,
      "segs": [
        {
          "utf8": "and when you use them, you'll know I'm using\nnatural English."
        }
      ]
    },
    {
      "tStartMs": 301240,
      "dDurationMs": 2720,
      "segs": [
        {
          "utf8": "The confidence in speaking will grow."
        }
      ]
    },
    {
      "tStartMs": 303960,
      "dDurationMs": 6600,
      "segs": [
        {
          "utf8": "Right now, I'm going to share my screen with\nyou so that you can see inside the course"
        }
      ]
    },
    {
      "tStartMs": 310560,
      "dDurationMs": 2409,
      "segs": [
        {
          "utf8": "and see exactly what it is."
        }
      ]
    },
    {
      "tStartMs": 312969,
      "dDurationMs": 7011,
      "segs": [
        {
          "utf8": "I'm going to show you the course as well as\nthe Facebook group that you can join to interact"
        }
      ]
    },
    {
      "tStartMs": 319980,
      "dDurationMs": 4409,
      "segs": [
        {
          "utf8": "with me and also interact with other members\nof the group."
        }
      ]
    },
    {
      "tStartMs": 324389,
      "dDurationMs": 5291,
      "segs": [
        {
          "utf8": "This is a great place to be able to use English\nand feel secure."
        }
      ]
    },
    {
      "tStartMs": 329680,
      "dDurationMs": 6450,
      "segs": [
        {
          "utf8": "You can feel confident using English in that\nenvironment because it's a private group."
        }
      ]
    },
    {
      "tStartMs": 336130,
      "dDurationMs": 5529,
      "segs": [
        {
          "utf8": "No one else on Facebook, except for members\nof the course and members of this group can"
        }
      ]
    },
    {
      "tStartMs": 341659,
      "dDurationMs": 2711,
      "segs": [
        {
          "utf8": "participate and can see your content."
        }
      ]
    },
    {
      "tStartMs": 344370,
      "dDurationMs": 6069,
      "segs": [
        {
          "utf8": "I'm going to show you exactly what the course\nis on my screen, and as well, what the Facebook"
        }
      ]
    },
    {
      "tStartMs": 350439,
      "dDurationMs": 1051,
      "segs": [
        {
          "utf8": "group is."
        }
      ]
    },
    {
      "tStartMs": 351490,
      "dDurationMs": 2250,
      "segs": [
        {
          "utf8": "Let's take a look."
        }
      ]
    },
    {
      "tStartMs": 353740,
      "dDurationMs": 1000,
      "segs": [
        {
          "utf8": "All right."
        }
      ]
    },
    {
      "tStartMs": 354740,
      "dDurationMs": 4350,
      "segs": [
        {
          "utf8": "Here, we are on my computer screen and I want\nto show you inside the course."
        }
      ]
    },
    {
      "tStartMs": 359090,
      "dDurationMs": 4889,
      "segs": [
        {
          "utf8": "Here, we have the five different sections\nof the course."
        }
      ]
    },
    {
      "tStartMs": 363979,
      "dDurationMs": 8240,
      "segs": [
        {
          "utf8": "I grouped these expressions into five different\nsections; family, travel, food, work and education"
        }
      ]
    },
    {
      "tStartMs": 372219,
      "dDurationMs": 4251,
      "segs": [
        {
          "utf8": "so that you can more easily remember them\nin a context."
        }
      ]
    },
    {
      "tStartMs": 376470,
      "dDurationMs": 3439,
      "segs": [
        {
          "utf8": "You can use each expression in a variety of\ncircumstances."
        }
      ]
    },
    {
      "tStartMs": 379909,
      "dDurationMs": 6410,
      "segs": [
        {
          "utf8": "The first thing that I want you to do when\nyou open the course is to click on the, \"Start"
        }
      ]
    },
    {
      "tStartMs": 386319,
      "dDurationMs": 1000,
      "segs": [
        {
          "utf8": "here."
        }
      ]
    },
    {
      "tStartMs": 387319,
      "dDurationMs": 3820,
      "segs": [
        {
          "utf8": "How to use the 50 Natural English Expressions\nProgram,\" PDF file."
        }
      ]
    },
    {
      "tStartMs": 391139,
      "dDurationMs": 1990,
      "segs": [
        {
          "utf8": "I have it already opened down here."
        }
      ]
    },
    {
      "tStartMs": 393129,
      "dDurationMs": 1720,
      "segs": [
        {
          "utf8": "I'd like to share this with you."
        }
      ]
    },
    {
      "tStartMs": 394849,
      "dDurationMs": 4940,
      "segs": [
        {
          "utf8": "Here, we have a lot of information that you\ncan use throughout the course."
        }
      ]
    },
    {
      "tStartMs": 399789,
      "dDurationMs": 2951,
      "segs": [
        {
          "utf8": "Here, we have the Facebook group link."
        }
      ]
    },
    {
      "tStartMs": 402740,
      "dDurationMs": 7149,
      "segs": [
        {
          "utf8": "There are all of the expressions, the course\noutline, and these are some general tips for"
        }
      ]
    },
    {
      "tStartMs": 409889,
      "dDurationMs": 1291,
      "segs": [
        {
          "utf8": "using the course."
        }
      ]
    },
    {
      "tStartMs": 411180,
      "dDurationMs": 6340,
      "segs": [
        {
          "utf8": "Downloading, listening, writing, speaking,\nand this is a weekly study guide."
        }
      ]
    },
    {
      "tStartMs": 417520,
      "dDurationMs": 4530,
      "segs": [
        {
          "utf8": "In this weekly study guide, this is, of course,\nyour choice whether or not you want to use"
        }
      ]
    },
    {
      "tStartMs": 422050,
      "dDurationMs": 6419,
      "segs": [
        {
          "utf8": "it, but I wanted to give you my suggestion\nso that you don't feel overwhelmed and try"
        }
      ]
    },
    {
      "tStartMs": 428469,
      "dDurationMs": 4850,
      "segs": [
        {
          "utf8": "to learn all these 50 expressions at the same\ntime."
        }
      ]
    },
    {
      "tStartMs": 433319,
      "dDurationMs": 5731,
      "segs": [
        {
          "utf8": "I want you to take a look at this first when\nyou download the course, and then you will"
        }
      ]
    },
    {
      "tStartMs": 439050,
      "dDurationMs": 3860,
      "segs": [
        {
          "utf8": "have each section."
        }
      ]
    },
    {
      "tStartMs": 442910,
      "dDurationMs": 6159,
      "segs": [
        {
          "utf8": "These pronunciation files, these are bonus\nfiles that are included if you joined during"
        }
      ]
    },
    {
      "tStartMs": 449069,
      "dDurationMs": 4951,
      "segs": [
        {
          "utf8": "the opening week or if you joined during one\nof my special promotions."
        }
      ]
    },
    {
      "tStartMs": 454020,
      "dDurationMs": 6949,
      "segs": [
        {
          "utf8": "Depending on when you joined, you might or\nmight not have this pronunciation file."
        }
      ]
    },
    {
      "tStartMs": 460969,
      "dDurationMs": 1480,
      "segs": [
        {
          "utf8": "Let's take a look at the first one."
        }
      ]
    },
    {
      "tStartMs": 462449,
      "dDurationMs": 1881,
      "segs": [
        {
          "utf8": "Let's open up family."
        }
      ]
    },
    {
      "tStartMs": 464330,
      "dDurationMs": 2260,
      "segs": [
        {
          "utf8": "All right."
        }
      ]
    },
    {
      "tStartMs": 466590,
      "dDurationMs": 2659,
      "segs": [
        {
          "utf8": "Inside family, you will not have these here."
        }
      ]
    },
    {
      "tStartMs": 469249,
      "dDurationMs": 2131,
      "segs": [
        {
          "utf8": "These are the official files."
        }
      ]
    },
    {
      "tStartMs": 471380,
      "dDurationMs": 2749,
      "segs": [
        {
          "utf8": "These are just my practice files."
        }
      ]
    },
    {
      "tStartMs": 474129,
      "dDurationMs": 7540,
      "segs": [
        {
          "utf8": "You will have all of the files for family\nas well as the transcript."
        }
      ]
    },
    {
      "tStartMs": 481669,
      "dDurationMs": 3890,
      "segs": [
        {
          "utf8": "Let's open up number four."
        }
      ]
    },
    {
      "tStartMs": 485559,
      "dDurationMs": 1000,
      "segs": [
        {
          "utf8": "All right."
        }
      ]
    },
    {
      "tStartMs": 486559,
      "dDurationMs": 3970,
      "segs": [
        {
          "utf8": "When you click on the file, you can play it\non your computer."
        }
      ]
    },
    {
      "tStartMs": 490529,
      "dDurationMs": 1860,
      "segs": [
        {
          "utf8": "You can play it on your phone."
        }
      ]
    },
    {
      "tStartMs": 492389,
      "dDurationMs": 1460,
      "segs": [
        {
          "utf8": "You can play it on your tablet."
        }
      ]
    },
    {
      "tStartMs": 493849,
      "dDurationMs": 5231,
      "segs": [
        {
          "utf8": "You can play this on any device and you can\nsimply download it onto that device."
        }
      ]
    },
    {
      "tStartMs": 499080,
      "dDurationMs": 6670,
      "segs": [
        {
          "utf8": "This is the audio listening lesson, and you\ncan also read the transcripts."
        }
      ]
    },
    {
      "tStartMs": 505750,
      "dDurationMs": 3039,
      "segs": [
        {
          "utf8": "Let's take a look at the transcript."
        }
      ]
    },
    {
      "tStartMs": 508789,
      "dDurationMs": 2201,
      "segs": [
        {
          "utf8": "This is what you will see for the transcript."
        }
      ]
    },
    {
      "tStartMs": 510990,
      "dDurationMs": 4480,
      "segs": [
        {
          "utf8": "This is an exact copy of everything that's\nin the lesson."
        }
      ]
    },
    {
      "tStartMs": 515470,
      "dDurationMs": 5759,
      "segs": [
        {
          "utf8": "You can follow the lesson and you can see\nall of the sample sentences, the questions,"
        }
      ]
    },
    {
      "tStartMs": 521229,
      "dDurationMs": 2871,
      "segs": [
        {
          "utf8": "the story."
        }
      ]
    },
    {
      "tStartMs": 524100,
      "dDurationMs": 4740,
      "segs": [
        {
          "utf8": "In the story, the word that you learned for\nthat lesson is going to be bold."
        }
      ]
    },
    {
      "tStartMs": 528840,
      "dDurationMs": 1580,
      "segs": [
        {
          "utf8": "It's going to be bold."
        }
      ]
    },
    {
      "tStartMs": 530420,
      "dDurationMs": 5810,
      "segs": [
        {
          "utf8": "Here, you can see these words are a little\nbit stronger, bolder than the other ones."
        }
      ]
    },
    {
      "tStartMs": 536230,
      "dDurationMs": 4600,
      "segs": [
        {
          "utf8": "That way, you can more easily see the expression\nin the story."
        }
      ]
    },
    {
      "tStartMs": 540830,
      "dDurationMs": 4240,
      "segs": [
        {
          "utf8": "You have the story and at the end, I ask a\nfew questions."
        }
      ]
    },
    {
      "tStartMs": 545070,
      "dDurationMs": 6090,
      "segs": [
        {
          "utf8": "These questions you can answer in your head,\nout loud or in our Facebook group."
        }
      ]
    },
    {
      "tStartMs": 551160,
      "dDurationMs": 5590,
      "segs": [
        {
          "utf8": "I recommend taking some time to answer in\nthe Facebook group so that you can participate"
        }
      ]
    },
    {
      "tStartMs": 556750,
      "dDurationMs": 1210,
      "segs": [
        {
          "utf8": "and get feedback."
        }
      ]
    },
    {
      "tStartMs": 557960,
      "dDurationMs": 1560,
      "segs": [
        {
          "utf8": "That's so important."
        }
      ]
    },
    {
      "tStartMs": 559520,
      "dDurationMs": 3980,
      "segs": [
        {
          "utf8": "Let's take a look now at how you can join\nthe Facebook group."
        }
      ]
    },
    {
      "tStartMs": 563500,
      "dDurationMs": 1490,
      "segs": [
        {
          "utf8": "What is the Facebook group?"
        }
      ]
    },
    {
      "tStartMs": 564990,
      "dDurationMs": 1750,
      "segs": [
        {
          "utf8": "I've got a link down here."
        }
      ]
    },
    {
      "tStartMs": 566740,
      "dDurationMs": 1510,
      "segs": [
        {
          "utf8": "Let's open this up."
        }
      ]
    },
    {
      "tStartMs": 568250,
      "dDurationMs": 2940,
      "segs": [
        {
          "utf8": "All right, this is the Facebook group."
        }
      ]
    },
    {
      "tStartMs": 571190,
      "dDurationMs": 4490,
      "segs": [
        {
          "utf8": "At the moment, there are only the testers\ninvolved in the group, but as soon as the"
        }
      ]
    },
    {
      "tStartMs": 575680,
      "dDurationMs": 7970,
      "segs": [
        {
          "utf8": "course opens on February 5th, 2016, there\nwill be more members hopefully including you"
        }
      ]
    },
    {
      "tStartMs": 583650,
      "dDurationMs": 2300,
      "segs": [
        {
          "utf8": "joining this course."
        }
      ]
    },
    {
      "tStartMs": 585950,
      "dDurationMs": 5590,
      "segs": [
        {
          "utf8": "Inside the group, there is a lot of information\nthat you can use to interact."
        }
      ]
    },
    {
      "tStartMs": 591540,
      "dDurationMs": 1420,
      "segs": [
        {
          "utf8": "You can tell a story."
        }
      ]
    },
    {
      "tStartMs": 592960,
      "dDurationMs": 1280,
      "segs": [
        {
          "utf8": "You can send videos."
        }
      ]
    },
    {
      "tStartMs": 594240,
      "dDurationMs": 7970,
      "segs": [
        {
          "utf8": "A lot of other English learners have posted\nintroduction videos and stories of what's"
        }
      ]
    },
    {
      "tStartMs": 602210,
      "dDurationMs": 2330,
      "segs": [
        {
          "utf8": "going on in their lives, what they're doing."
        }
      ]
    },
    {
      "tStartMs": 604540,
      "dDurationMs": 3680,
      "segs": [
        {
          "utf8": "You can see a lot of interesting information."
        }
      ]
    },
    {
      "tStartMs": 608220,
      "dDurationMs": 6710,
      "segs": [
        {
          "utf8": "This is a safe place for you to interact as\npart of our community."
        }
      ]
    },
    {
      "tStartMs": 614930,
      "dDurationMs": 4720,
      "segs": [
        {
          "utf8": "Thanks so much for learning English with me\nand I look forward to seeing you inside the"
        }
      ]
    },
    {
      "tStartMs": 619650,
      "dDurationMs": 1000,
      "segs": [
        {
          "utf8": "course."
        }
      ]
    },
    {
      "tStartMs": 620650,
      "dDurationMs": 10,
      "segs": [
        {
          "utf8": "Goodbye."
        }
      ]
    }
  ]
}`
	subtitleSrt = `0
00:00:00,210 --> 00:00:04,040
Hi, I'm Vanessa from SpeakEnglishWithVanessa.com.

1
00:00:04,040 --> 00:00:09,130
Today, I have some good news and some bad
news.

2
00:00:09,130 --> 00:00:10,809
Let's start with the bad news.

3
00:00:10,809 --> 00:00:18,980
The bad news is that you probably spent a
lot of time in boring, middle school and high

4
00:00:18,980 --> 00:00:20,890
school English classes.

5
00:00:20,890 --> 00:00:27,140
You felt a lot of stress, a lot of worry about
learning English because it just wasn't fun.

6
00:00:27,140 --> 00:00:28,900
You had to study for test.

7
00:00:28,900 --> 00:00:31,880
You had to memorize list of vocabulary.

8
00:00:31,880 --> 00:00:39,620
This is definitely not an enjoyable or effective
way to learn anything, especially English,

9
00:00:39,620 --> 00:00:41,679
but I have some good news.

10
00:00:41,679 --> 00:00:49,179
Now, you can choose your own teacher, your
own materials, your own method, the style.

11
00:00:49,179 --> 00:00:55,190
Now, you are an adult and you can choose what
you do with English.

12
00:00:55,190 --> 00:01:01,340
This is exactly what is the best way to learn
when you choose what's the best style for

13
00:01:01,340 --> 00:01:02,399
you.

14
00:01:02,399 --> 00:01:08,970
On Skype, I've taught over 2000 English lessons
to English learners around the world.

15
00:01:08,970 --> 00:01:10,620
You know what?

16
00:01:10,620 --> 00:01:14,960
Almost everyone makes the same two mistakes.

17
00:01:14,960 --> 00:01:16,740
What are those two mistakes?

18
00:01:16,740 --> 00:01:22,420
The first one is that they often use textbook
expressions.

19
00:01:22,420 --> 00:01:28,430
These expressions are ones that you learned
in the textbook, but native speakers don't

20
00:01:28,430 --> 00:01:31,280
really use them in daily conversation.

21
00:01:31,280 --> 00:01:37,670
Native speakers can understand what they mean,
but it sounds a little bit robotic or a little

22
00:01:37,670 --> 00:01:41,140
bit like a textbook, because that's where
you learned it.

23
00:01:41,140 --> 00:01:43,979
I know you don't want to speak textbook English.

24
00:01:43,979 --> 00:01:50,659
I don't want you to speak textbook English,
so you need to do something different.

25
00:01:50,659 --> 00:01:53,100
Make a huge change.

26
00:01:53,100 --> 00:01:55,759
Let's talk about the second mistake.

27
00:01:55,759 --> 00:02:02,520
The second mistake that almost all of my Skype
students make is there are a lot of daily

28
00:02:02,520 --> 00:02:08,750
common expressions that native speakers use
but they don't use, they don't know about

29
00:02:08,750 --> 00:02:12,960
and maybe they don't even understand until
I have to explain it.

30
00:02:12,960 --> 00:02:20,570
These expressions you definitely need to know
if you want to understand movies, TV shows,

31
00:02:20,570 --> 00:02:26,330
communicate in a natural way, and so that
you can feel more confident because you know

32
00:02:26,330 --> 00:02:29,860
that you are using natural English.

33
00:02:29,860 --> 00:02:34,140
That's why I've created this course, 50 Natural
English Expressions.

34
00:02:34,140 --> 00:02:41,210
I've made a list of all of the English expressions
that my English students don't know or don't

35
00:02:41,210 --> 00:02:46,830
use, and I made them into lesson files especially
for you.

36
00:02:46,830 --> 00:02:49,130
I know that a lot of you are busy.

37
00:02:49,130 --> 00:02:53,530
You have a lot of things going on in your
life, so these lessons you can download.

38
00:02:53,530 --> 00:02:59,920
You can listen to them wherever you are and
you can use it and learn it and have fun.

39
00:02:59,920 --> 00:03:05,290
I love hearing what members of this course
have to say about their experience.

40
00:03:05,290 --> 00:03:10,590
I think that you shouldn't only listen to
what I'm saying about it, but also listen

41
00:03:10,590 --> 00:03:15,190
to other English learners, because they know
your situation exactly.

42
00:03:15,190 --> 00:03:21,320
They know what it's like to feel frustrated,
to feel upset about your level and they also

43
00:03:21,320 --> 00:03:24,360
want to improve their English just like you.

44
00:03:24,360 --> 00:03:28,320
Let's take a look at two people who are part
of this course.

45
00:03:28,320 --> 00:03:31,100
Two members and see what they have to say.

46
00:03:31,100 --> 00:03:36,470
I'm going to read what Vera from Brazil had
to say about her experience in the course,

47
00:03:36,470 --> 00:03:38,560
50 Natural English Expressions.

48
00:03:38,560 --> 00:03:43,830
She said, "Well, for me, the way the course
was set up is perfect because we have the

49
00:03:43,830 --> 00:03:50,120
explanation of expressions, examples of how
to use them, a story to understand and fix

50
00:03:50,120 --> 00:03:55,690
what you've learned and still have the opportunity
to use them making videos or posting comments

51
00:03:55,690 --> 00:03:57,000
in the Facebook group."

52
00:03:57,000 --> 00:03:58,000
Thanks Vera.

53
00:03:58,000 --> 00:03:59,810
That's a really honest comment.

54
00:03:59,810 --> 00:04:02,910
Let's take a look at what Jean-Jacques from
France has to say.

55
00:04:02,910 --> 00:04:08,420
He said, "Even if you are a beginner or an
intermediate English learner, Teacher Vanessa's

56
00:04:08,420 --> 00:04:09,790
method is useful.

57
00:04:09,790 --> 00:04:12,180
You will learn with simplicity.

58
00:04:12,180 --> 00:04:17,699
If you want to accelerate the process of learning,
acquire the 50 Natural English Expressions

59
00:04:17,699 --> 00:04:18,699
Course."

60
00:04:18,699 --> 00:04:22,680
Thanks Jean-Jacques, I think that's really
a good point, simplicity.

61
00:04:22,680 --> 00:04:28,750
You don't need to study really complex grammar
rules or feel frustrated and upset about complex

62
00:04:28,750 --> 00:04:34,500
things in English, you can learn English with
simplicity and not feel stressed and enjoy

63
00:04:34,500 --> 00:04:36,170
it.

64
00:04:36,170 --> 00:04:40,360
That's why I've created this course, 50 Natural
English Expressions.

65
00:04:40,360 --> 00:04:47,629
I want you to be able to use these expressions
that are not textbook expressions, and that

66
00:04:47,629 --> 00:04:50,460
native speakers use all the time.

67
00:04:50,460 --> 00:04:56,189
You are definitely going to hear these expressions
in movies, TV shows, in natural conversations

68
00:04:56,189 --> 00:05:01,240
and when you use them, you'll know I'm using
natural English.

69
00:05:01,240 --> 00:05:03,960
The confidence in speaking will grow.

70
00:05:03,960 --> 00:05:10,560
Right now, I'm going to share my screen with
you so that you can see inside the course

71
00:05:10,560 --> 00:05:12,969
and see exactly what it is.

72
00:05:12,969 --> 00:05:19,980
I'm going to show you the course as well as
the Facebook group that you can join to interact

73
00:05:19,980 --> 00:05:24,389
with me and also interact with other members
of the group.

74
00:05:24,389 --> 00:05:29,680
This is a great place to be able to use English
and feel secure.

75
00:05:29,680 --> 00:05:36,130
You can feel confident using English in that
environment because it's a private group.

76
00:05:36,130 --> 00:05:41,659
No one else on Facebook, except for members
of the course and members of this group can

77
00:05:41,659 --> 00:05:44,370
participate and can see your content.

78
00:05:44,370 --> 00:05:50,439
I'm going to show you exactly what the course
is on my screen, and as well, what the Facebook

79
00:05:50,439 --> 00:05:51,490
group is.

80
00:05:51,490 --> 00:05:53,740
Let's take a look.

81
00:05:53,740 --> 00:05:54,740
All right.

82
00:05:54,740 --> 00:05:59,090
Here, we are on my computer screen and I want
to show you inside the course.

83
00:05:59,090 --> 00:06:03,979
Here, we have the five different sections
of the course.

84
00:06:03,979 --> 00:06:12,219
I grouped these expressions into five different
sections; family, travel, food, work and education

85
00:06:12,219 --> 00:06:16,470
so that you can more easily remember them
in a context.

86
00:06:16,470 --> 00:06:19,909
You can use each expression in a variety of
circumstances.

87
00:06:19,909 --> 00:06:26,319
The first thing that I want you to do when
you open the course is to click on the, "Start

88
00:06:26,319 --> 00:06:27,319
here.

89
00:06:27,319 --> 00:06:31,139
How to use the 50 Natural English Expressions
Program," PDF file.

90
00:06:31,139 --> 00:06:33,129
I have it already opened down here.

91
00:06:33,129 --> 00:06:34,849
I'd like to share this with you.

92
00:06:34,849 --> 00:06:39,789
Here, we have a lot of information that you
can use throughout the course.

93
00:06:39,789 --> 00:06:42,740
Here, we have the Facebook group link.

94
00:06:42,740 --> 00:06:49,889
There are all of the expressions, the course
outline, and these are some general tips for

95
00:06:49,889 --> 00:06:51,180
using the course.

96
00:06:51,180 --> 00:06:57,520
Downloading, listening, writing, speaking,
and this is a weekly study guide.

97
00:06:57,520 --> 00:07:02,050
In this weekly study guide, this is, of course,
your choice whether or not you want to use

98
00:07:02,050 --> 00:07:08,469
it, but I wanted to give you my suggestion
so that you don't feel overwhelmed and try

99
00:07:08,469 --> 00:07:13,319
to learn all these 50 expressions at the same
time.

100
00:07:13,319 --> 00:07:19,050
I want you to take a look at this first when
you download the course, and then you will

101
00:07:19,050 --> 00:07:22,910
have each section.

102
00:07:22,910 --> 00:07:29,069
These pronunciation files, these are bonus
files that are included if you joined during

103
00:07:29,069 --> 00:07:34,020
the opening week or if you joined during one
of my special promotions.

104
00:07:34,020 --> 00:07:40,969
Depending on when you joined, you might or
might not have this pronunciation file.

105
00:07:40,969 --> 00:07:42,449
Let's take a look at the first one.

106
00:07:42,449 --> 00:07:44,330
Let's open up family.

107
00:07:44,330 --> 00:07:46,590
All right.

108
00:07:46,590 --> 00:07:49,249
Inside family, you will not have these here.

109
00:07:49,249 --> 00:07:51,380
These are the official files.

110
00:07:51,380 --> 00:07:54,129
These are just my practice files.

111
00:07:54,129 --> 00:08:01,669
You will have all of the files for family
as well as the transcript.

112
00:08:01,669 --> 00:08:05,559
Let's open up number four.

113
00:08:05,559 --> 00:08:06,559
All right.

114
00:08:06,559 --> 00:08:10,529
When you click on the file, you can play it
on your computer.

115
00:08:10,529 --> 00:08:12,389
You can play it on your phone.

116
00:08:12,389 --> 00:08:13,849
You can play it on your tablet.

117
00:08:13,849 --> 00:08:19,080
You can play this on any device and you can
simply download it onto that device.

118
00:08:19,080 --> 00:08:25,750
This is the audio listening lesson, and you
can also read the transcripts.

119
00:08:25,750 --> 00:08:28,789
Let's take a look at the transcript.

120
00:08:28,789 --> 00:08:30,990
This is what you will see for the transcript.

121
00:08:30,990 --> 00:08:35,470
This is an exact copy of everything that's
in the lesson.

122
00:08:35,470 --> 00:08:41,229
You can follow the lesson and you can see
all of the sample sentences, the questions,

123
00:08:41,229 --> 00:08:44,100
the story.

124
00:08:44,100 --> 00:08:48,840
In the story, the word that you learned for
that lesson is going to be bold.

125
00:08:48,840 --> 00:08:50,420
It's going to be bold.

126
00:08:50,420 --> 00:08:56,230
Here, you can see these words are a little
bit stronger, bolder than the other ones.

127
00:08:56,230 --> 00:09:00,830
That way, you can more easily see the expression
in the story.

128
00:09:00,830 --> 00:09:05,070
You have the story and at the end, I ask a
few questions.

129
00:09:05,070 --> 00:09:11,160
These questions you can answer in your head,
out loud or in our Facebook group.

130
00:09:11,160 --> 00:09:16,750
I recommend taking some time to answer in
the Facebook group so that you can participate

131
00:09:16,750 --> 00:09:17,960
and get feedback.

132
00:09:17,960 --> 00:09:19,520
That's so important.

133
00:09:19,520 --> 00:09:23,500
Let's take a look now at how you can join
the Facebook group.

134
00:09:23,500 --> 00:09:24,990
What is the Facebook group?

135
00:09:24,990 --> 00:09:26,740
I've got a link down here.

136
00:09:26,740 --> 00:09:28,250
Let's open this up.

137
00:09:28,250 --> 00:09:31,190
All right, this is the Facebook group.

138
00:09:31,190 --> 00:09:35,680
At the moment, there are only the testers
involved in the group, but as soon as the

139
00:09:35,680 --> 00:09:43,650
course opens on February 5th, 2016, there
will be more members hopefully including you

140
00:09:43,650 --> 00:09:45,950
joining this course.

141
00:09:45,950 --> 00:09:51,540
Inside the group, there is a lot of information
that you can use to interact.

142
00:09:51,540 --> 00:09:52,960
You can tell a story.

143
00:09:52,960 --> 00:09:54,240
You can send videos.

144
00:09:54,240 --> 00:10:02,210
A lot of other English learners have posted
introduction videos and stories of what's

145
00:10:02,210 --> 00:10:04,540
going on in their lives, what they're doing.

146
00:10:04,540 --> 00:10:08,220
You can see a lot of interesting information.

147
00:10:08,220 --> 00:10:14,930
This is a safe place for you to interact as
part of our community.

148
00:10:14,930 --> 00:10:19,650
Thanks so much for learning English with me
and I look forward to seeing you inside the

149
00:10:19,650 --> 00:10:20,650
course.

150
00:10:20,650 --> 00:10:20,660
Goodbye.

`
)

var (
	subtitleJson3_2 = `{
  "wireMagic": "pb3",
  "pens": [ {
  
  }, {
    "fcForeColor": 15066597
  }, {
    "fcForeColor": 13421772
  } ],
  "wsWinStyles": [ {
  
  }, {
    "mhModeHint": 2,
    "juJustifCode": 0,
    "sdScrollDir": 3
  } ],
  "wpWinPositions": [ {
  
  }, {
    "apPoint": 6,
    "ahHorPos": 20,
    "avVerPos": 100,
    "rcRows": 2,
    "ccCols": 40
  } ],
  "events": [ {
    "tStartMs": 0,
    "dDurationMs": 2974800,
    "id": 1,
    "wpWinPosId": 1,
    "wsWinStyleId": 1
  }, {
    "tStartMs": 0,
    "dDurationMs": 5130,
    "wWinId": 1,
    "segs": [ {
      "utf8": "in",
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 89,
      "acAsrConf": 252
    }, {
      "utf8": " comments",
      "tOffsetMs": 210,
      "acAsrConf": 237
    }, {
      "utf8": " all",
      "tOffsetMs": 690,
      "acAsrConf": 217
    }, {
      "utf8": " right",
      "tOffsetMs": 960,
      "acAsrConf": 217
    }, {
      "utf8": " oh",
      "tOffsetMs": 1380,
      "acAsrConf": 157
    }, {
      "utf8": " good",
      "tOffsetMs": 1860,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 2639,
      "acAsrConf": 153
    } ]
  }, {
    "tStartMs": 2990,
    "dDurationMs": 2140,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 3000,
    "dDurationMs": 4490,
    "wWinId": 1,
    "segs": [ {
      "utf8": "think",
      "acAsrConf": 252
    }, {
      "utf8": " we",
      "tOffsetMs": 419,
      "acAsrConf": 94
    }, {
      "utf8": " found",
      "tOffsetMs": 659,
      "acAsrConf": 246
    }, {
      "utf8": " some",
      "tOffsetMs": 840,
      "acAsrConf": 252
    }, {
      "utf8": " comments",
      "tOffsetMs": 1230,
      "acAsrConf": 252
    }, {
      "utf8": " all",
      "tOffsetMs": 1680,
      "acAsrConf": 123
    }, {
      "utf8": " right",
      "tOffsetMs": 1859,
      "acAsrConf": 226
    } ]
  }, {
    "tStartMs": 5120,
    "dDurationMs": 2370,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 5130,
    "dDurationMs": 5880,
    "wWinId": 1,
    "segs": [ {
      "utf8": "so",
      "acAsrConf": 252
    }, {
      "utf8": " we",
      "tOffsetMs": 209,
      "acAsrConf": 251
    }, {
      "utf8": " were",
      "tOffsetMs": 360,
      "acAsrConf": 251
    }, {
      "utf8": " going",
      "tOffsetMs": 450,
      "acAsrConf": 221
    }, {
      "utf8": " to",
      "tOffsetMs": 629,
      "acAsrConf": 252
    }, {
      "utf8": " get",
      "tOffsetMs": 749,
      "acAsrConf": 216
    }, {
      "utf8": " started",
      "tOffsetMs": 870,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1110,
      "acAsrConf": 200
    }, {
      "utf8": " all",
      "tOffsetMs": 1499,
      "acAsrConf": 235
    } ]
  }, {
    "tStartMs": 7480,
    "dDurationMs": 3530,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 7490,
    "dDurationMs": 6430,
    "wWinId": 1,
    "segs": [ {
      "utf8": "change",
      "acAsrConf": 217
    }, {
      "utf8": " with",
      "tOffsetMs": 1000,
      "acAsrConf": 210
    }, {
      "utf8": " you",
      "tOffsetMs": 1330,
      "acAsrConf": 252
    }, {
      "utf8": " over",
      "tOffsetMs": 1450,
      "acAsrConf": 249
    }, {
      "utf8": " here",
      "tOffsetMs": 1689,
      "acAsrConf": 224
    }, {
      "utf8": " all",
      "tOffsetMs": 2310,
      "acAsrConf": 225
    }, {
      "utf8": " right",
      "tOffsetMs": 3310,
      "acAsrConf": 218
    }, {
      "utf8": " so",
      "tOffsetMs": 3370,
      "acAsrConf": 102
    } ]
  }, {
    "tStartMs": 11000,
    "dDurationMs": 2920,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 11010,
    "dDurationMs": 6720,
    "wWinId": 1,
    "segs": [ {
      "utf8": "if",
      "acAsrConf": 208
    }, {
      "utf8": " you",
      "tOffsetMs": 299,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 330,
      "acAsrConf": 201
    }, {
      "utf8": " any",
      "tOffsetMs": 629,
      "acAsrConf": 252
    }, {
      "utf8": " questions",
      "tOffsetMs": 810,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1290,
      "acAsrConf": 252
    }, {
      "utf8": " can",
      "tOffsetMs": 1589,
      "acAsrConf": 202
    }, {
      "utf8": " write",
      "tOffsetMs": 1919,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 13910,
    "dDurationMs": 3820,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 13920,
    "dDurationMs": 6000,
    "wWinId": 1,
    "segs": [ {
      "utf8": "them",
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 60,
      "acAsrConf": 134
    }, {
      "utf8": " the",
      "tOffsetMs": 270,
      "acAsrConf": 202
    }, {
      "utf8": " same",
      "tOffsetMs": 570,
      "acAsrConf": 252
    }, {
      "utf8": " box",
      "tOffsetMs": 810,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 1139,
      "acAsrConf": 186
    }, {
      "utf8": " the",
      "tOffsetMs": 2240,
      "acAsrConf": 240
    }, {
      "utf8": " home",
      "tOffsetMs": 3240,
      "acAsrConf": 206
    }, {
      "utf8": " page",
      "tOffsetMs": 3480,
      "acAsrConf": 206
    } ]
  }, {
    "tStartMs": 17720,
    "dDurationMs": 2200,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 17730,
    "dDurationMs": 3809,
    "wWinId": 1,
    "segs": [ {
      "utf8": "oh",
      "acAsrConf": 232
    }, {
      "utf8": " it's",
      "tOffsetMs": 270,
      "acAsrConf": 252
    }, {
      "utf8": " great",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 690,
      "acAsrConf": 236
    }, {
      "utf8": " meet",
      "tOffsetMs": 750,
      "acAsrConf": 172
    }, {
      "utf8": " you",
      "tOffsetMs": 1020,
      "acAsrConf": 252
    }, {
      "utf8": " too",
      "tOffsetMs": 1139,
      "acAsrConf": 226
    }, {
      "utf8": " and",
      "tOffsetMs": 1170,
      "acAsrConf": 198
    }, {
      "utf8": " I",
      "tOffsetMs": 1799,
      "acAsrConf": 245
    }, {
      "utf8": " hope",
      "tOffsetMs": 1980,
      "acAsrConf": 227
    } ]
  }, {
    "tStartMs": 19910,
    "dDurationMs": 1629,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 19920,
    "dDurationMs": 2790,
    "wWinId": 1,
    "segs": [ {
      "utf8": "that",
      "acAsrConf": 226
    }, {
      "utf8": " we",
      "tOffsetMs": 240,
      "acAsrConf": 249
    }, {
      "utf8": " will",
      "tOffsetMs": 449,
      "acAsrConf": 200
    }, {
      "utf8": " be",
      "tOffsetMs": 630,
      "acAsrConf": 211
    }, {
      "utf8": " able",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 840,
      "acAsrConf": 206
    }, {
      "utf8": " give",
      "tOffsetMs": 1109,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1290,
      "acAsrConf": 252
    }, {
      "utf8": " some",
      "tOffsetMs": 1410,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 21529,
    "dDurationMs": 1181,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 21539,
    "dDurationMs": 3361,
    "wWinId": 1,
    "segs": [ {
      "utf8": "great",
      "acAsrConf": 252
    }, {
      "utf8": " information",
      "tOffsetMs": 181,
      "acAsrConf": 206
    }, {
      "utf8": " today",
      "tOffsetMs": 301,
      "acAsrConf": 209
    } ]
  }, {
    "tStartMs": 22700,
    "dDurationMs": 2200,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 22710,
    "dDurationMs": 4770,
    "wWinId": 1,
    "segs": [ {
      "utf8": "thanks",
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 390,
      "acAsrConf": 228
    }, {
      "utf8": " your",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " patience",
      "tOffsetMs": 659,
      "acAsrConf": 252
    }, {
      "utf8": " today",
      "tOffsetMs": 1380,
      "acAsrConf": 249
    }, {
      "utf8": " we're",
      "tOffsetMs": 1950,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 24890,
    "dDurationMs": 2590,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 24900,
    "dDurationMs": 5430,
    "wWinId": 1,
    "segs": [ {
      "utf8": "going",
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 150,
      "acAsrConf": 210
    }, {
      "utf8": " talk",
      "tOffsetMs": 240,
      "acAsrConf": 198
    }, {
      "utf8": " about",
      "tOffsetMs": 450,
      "acAsrConf": 252
    }, {
      "utf8": " conversation",
      "tOffsetMs": 480,
      "acAsrConf": 203
    }, {
      "utf8": " English",
      "tOffsetMs": 1580,
      "acAsrConf": 175
    } ]
  }, {
    "tStartMs": 27470,
    "dDurationMs": 2860,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 27480,
    "dDurationMs": 5370,
    "wWinId": 1,
    "segs": [ {
      "utf8": "tips",
      "acAsrConf": 240
    }, {
      "utf8": " because",
      "tOffsetMs": 510,
      "acAsrConf": 207
    }, {
      "utf8": " I",
      "tOffsetMs": 1260,
      "acAsrConf": 235
    }, {
      "utf8": " know",
      "tOffsetMs": 1350,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 1410,
      "acAsrConf": 252
    }, {
      "utf8": " lot",
      "tOffsetMs": 1530,
      "acAsrConf": 219
    }, {
      "utf8": " of",
      "tOffsetMs": 1889,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1920,
      "acAsrConf": 236
    }, {
      "utf8": " want",
      "tOffsetMs": 2250,
      "acAsrConf": 249
    }, {
      "utf8": " to",
      "tOffsetMs": 2670,
      "acAsrConf": 218
    } ]
  }, {
    "tStartMs": 30320,
    "dDurationMs": 2530,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 30330,
    "dDurationMs": 6389,
    "wWinId": 1,
    "segs": [ {
      "utf8": "be",
      "acAsrConf": 252
    }, {
      "utf8": " fluent",
      "tOffsetMs": 150,
      "acAsrConf": 221
    }, {
      "utf8": " speakers",
      "tOffsetMs": 660,
      "acAsrConf": 242
    }, {
      "utf8": " and",
      "tOffsetMs": 1289,
      "acAsrConf": 250
    }, {
      "utf8": " speak",
      "tOffsetMs": 1679,
      "acAsrConf": 252
    }, {
      "utf8": " naturally",
      "tOffsetMs": 2310,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 32840,
    "dDurationMs": 3879,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 32850,
    "dDurationMs": 7440,
    "wWinId": 1,
    "segs": [ {
      "utf8": "smoothly",
      "acAsrConf": 248
    }, {
      "utf8": " right",
      "tOffsetMs": 869,
      "acAsrConf": 249
    }, {
      "utf8": " right",
      "tOffsetMs": 1500,
      "acAsrConf": 243
    }, {
      "utf8": " yes",
      "tOffsetMs": 2250,
      "acAsrConf": 241
    }, {
      "utf8": " so",
      "tOffsetMs": 2760,
      "acAsrConf": 227
    }, {
      "utf8": " Dan",
      "tOffsetMs": 3060,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 3389,
      "acAsrConf": 252
    }, {
      "utf8": " here",
      "tOffsetMs": 3660,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 36709,
    "dDurationMs": 3581,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 36719,
    "dDurationMs": 6360,
    "wWinId": 1,
    "segs": [ {
      "utf8": "today",
      "acAsrConf": 221
    }, {
      "utf8": " to",
      "tOffsetMs": 241,
      "acAsrConf": 217
    }, {
      "utf8": " help",
      "tOffsetMs": 1081,
      "acAsrConf": 252
    }, {
      "utf8": " me",
      "tOffsetMs": 1561,
      "acAsrConf": 163
    }, {
      "utf8": " explain",
      "tOffsetMs": 1860,
      "acAsrConf": 232
    }, {
      "utf8": " some",
      "tOffsetMs": 2610,
      "acAsrConf": 138
    }, {
      "utf8": " things",
      "tOffsetMs": 2941,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 2971,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 40280,
    "dDurationMs": 2799,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 40290,
    "dDurationMs": 4470,
    "wWinId": 1,
    "segs": [ {
      "utf8": "to",
      "acAsrConf": 252
    }, {
      "utf8": " give",
      "tOffsetMs": 210,
      "acAsrConf": 252
    }, {
      "utf8": " another",
      "tOffsetMs": 390,
      "acAsrConf": 252
    }, {
      "utf8": " perspective",
      "tOffsetMs": 720,
      "acAsrConf": 248
    }, {
      "utf8": " because",
      "tOffsetMs": 1760,
      "acAsrConf": 202
    }, {
      "utf8": " I",
      "tOffsetMs": 2760,
      "acAsrConf": 249
    } ]
  }, {
    "tStartMs": 43069,
    "dDurationMs": 1691,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 43079,
    "dDurationMs": 4831,
    "wWinId": 1,
    "segs": [ {
      "utf8": "hope",
      "acAsrConf": 239
    }, {
      "utf8": " I",
      "tOffsetMs": 421,
      "acAsrConf": 250
    }, {
      "utf8": " have",
      "tOffsetMs": 601,
      "acAsrConf": 248
    }, {
      "utf8": " a",
      "tOffsetMs": 811,
      "acAsrConf": 250
    }, {
      "utf8": " lot",
      "tOffsetMs": 841,
      "acAsrConf": 250
    }, {
      "utf8": " of",
      "tOffsetMs": 1140,
      "acAsrConf": 157
    }, {
      "utf8": " valuable",
      "tOffsetMs": 1291,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 44750,
    "dDurationMs": 3160,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 44760,
    "dDurationMs": 5790,
    "wWinId": 1,
    "segs": [ {
      "utf8": "information",
      "acAsrConf": 136
    }, {
      "utf8": " to",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " share",
      "tOffsetMs": 900,
      "acAsrConf": 207
    }, {
      "utf8": " but",
      "tOffsetMs": 1260,
      "acAsrConf": 224
    }, {
      "utf8": " another",
      "tOffsetMs": 1889,
      "acAsrConf": 247
    }, {
      "utf8": " person",
      "tOffsetMs": 2760,
      "acAsrConf": 246
    } ]
  }, {
    "tStartMs": 47900,
    "dDurationMs": 2650,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 47910,
    "dDurationMs": 4289,
    "wWinId": 1,
    "segs": [ {
      "utf8": "and",
      "acAsrConf": 228
    }, {
      "utf8": " hearing",
      "tOffsetMs": 570,
      "acAsrConf": 252
    }, {
      "utf8": " that",
      "tOffsetMs": 1110,
      "acAsrConf": 252
    }, {
      "utf8": " dialogue",
      "tOffsetMs": 1290,
      "acAsrConf": 232
    }, {
      "utf8": " together",
      "tOffsetMs": 1950,
      "acAsrConf": 246
    }, {
      "utf8": " is",
      "tOffsetMs": 2190,
      "acAsrConf": 206
    } ]
  }, {
    "tStartMs": 50540,
    "dDurationMs": 1659,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 50550,
    "dDurationMs": 5189,
    "wWinId": 1,
    "segs": [ {
      "utf8": "really",
      "acAsrConf": 208
    }, {
      "utf8": " valuable",
      "tOffsetMs": 419,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 630,
      "acAsrConf": 155
    }, {
      "utf8": " that's",
      "tOffsetMs": 1079,
      "acAsrConf": 226
    }, {
      "utf8": " what",
      "tOffsetMs": 1380,
      "acAsrConf": 252
    }, {
      "utf8": " we're",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 52189,
    "dDurationMs": 3550,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 52199,
    "dDurationMs": 6421,
    "wWinId": 1,
    "segs": [ {
      "utf8": "going",
      "acAsrConf": 207
    }, {
      "utf8": " to",
      "tOffsetMs": 121,
      "acAsrConf": 252
    }, {
      "utf8": " talk",
      "tOffsetMs": 151,
      "acAsrConf": 211
    }, {
      "utf8": " about",
      "tOffsetMs": 360,
      "acAsrConf": 155
    }, {
      "utf8": " today",
      "tOffsetMs": 511,
      "acAsrConf": 243
    }, {
      "utf8": " so",
      "tOffsetMs": 750,
      "acAsrConf": 170
    }, {
      "utf8": " we",
      "tOffsetMs": 1471,
      "acAsrConf": 252
    }, {
      "utf8": " will",
      "tOffsetMs": 1891,
      "acAsrConf": 200
    }, {
      "utf8": " put",
      "tOffsetMs": 2540,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 55729,
    "dDurationMs": 2891,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 55739,
    "dDurationMs": 4561,
    "wWinId": 1,
    "segs": [ {
      "utf8": "this",
      "acAsrConf": 226
    }, {
      "utf8": " over",
      "tOffsetMs": 210,
      "acAsrConf": 125
    }, {
      "utf8": " here",
      "tOffsetMs": 660,
      "acAsrConf": 78
    }, {
      "utf8": " and",
      "tOffsetMs": 1021,
      "acAsrConf": 209
    }, {
      "utf8": " get",
      "tOffsetMs": 1230,
      "acAsrConf": 249
    }, {
      "utf8": " started",
      "tOffsetMs": 1921,
      "acAsrConf": 186
    }, {
      "utf8": " all",
      "tOffsetMs": 2401,
      "acAsrConf": 84
    }, {
      "utf8": " right",
      "tOffsetMs": 2611,
      "acAsrConf": 207
    } ]
  }, {
    "tStartMs": 58610,
    "dDurationMs": 1690,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 58620,
    "dDurationMs": 5399,
    "wWinId": 1,
    "segs": [ {
      "utf8": "so",
      "acAsrConf": 255
    }, {
      "utf8": " they",
      "tOffsetMs": 300,
      "acAsrConf": 226
    }, {
      "utf8": " maybe",
      "tOffsetMs": 419,
      "acAsrConf": 251
    }, {
      "utf8": " should",
      "tOffsetMs": 599,
      "acAsrConf": 238
    }, {
      "utf8": " have",
      "tOffsetMs": 900,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 930,
      "acAsrConf": 209
    }, {
      "utf8": " mind",
      "tOffsetMs": 1200,
      "acAsrConf": 241
    }, {
      "utf8": " map",
      "tOffsetMs": 1410,
      "acAsrConf": 241
    } ]
  }, {
    "tStartMs": 60290,
    "dDurationMs": 3729,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 60300,
    "dDurationMs": 5609,
    "wWinId": 1,
    "segs": [ {
      "utf8": "right",
      "acAsrConf": 232
    }, {
      "utf8": " yes",
      "tOffsetMs": 360,
      "acAsrConf": 200
    }, {
      "utf8": " so",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    }, {
      "utf8": " do",
      "tOffsetMs": 1320,
      "acAsrConf": 255
    }, {
      "utf8": " you",
      "tOffsetMs": 1649,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 1890,
      "acAsrConf": 222
    }, {
      "utf8": " the",
      "tOffsetMs": 2250,
      "acAsrConf": 217
    }, {
      "utf8": " download",
      "tOffsetMs": 2910,
      "acAsrConf": 238
    }, {
      "utf8": " I",
      "tOffsetMs": 3270,
      "acAsrConf": 148
    } ]
  }, {
    "tStartMs": 64009,
    "dDurationMs": 1900,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 64019,
    "dDurationMs": 7231,
    "wWinId": 1,
    "segs": [ {
      "utf8": "gave",
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 241,
      "acAsrConf": 252
    }, {
      "utf8": " download",
      "tOffsetMs": 271,
      "acAsrConf": 251
    }, {
      "utf8": " which",
      "tOffsetMs": 901,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 1081,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 1290,
      "acAsrConf": 252
    }, {
      "utf8": " mind",
      "tOffsetMs": 1320,
      "acAsrConf": 244
    }, {
      "utf8": " map",
      "tOffsetMs": 1710,
      "acAsrConf": 244
    } ]
  }, {
    "tStartMs": 65899,
    "dDurationMs": 5351,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 65909,
    "dDurationMs": 10350,
    "wWinId": 1,
    "segs": [ {
      "utf8": "thanks",
      "acAsrConf": 208
    }, {
      "utf8": " for",
      "tOffsetMs": 301,
      "acAsrConf": 165
    }, {
      "utf8": " asking",
      "tOffsetMs": 601,
      "acAsrConf": 203
    }, {
      "utf8": " it",
      "tOffsetMs": 1051,
      "acAsrConf": 216
    }, {
      "utf8": " is",
      "tOffsetMs": 1671,
      "acAsrConf": 133
    }, {
      "utf8": " on",
      "tOffsetMs": 2671,
      "acAsrConf": 248
    }, {
      "utf8": " outline",
      "tOffsetMs": 3381,
      "acAsrConf": 241
    }, {
      "utf8": " with",
      "tOffsetMs": 4381,
      "acAsrConf": 234
    } ]
  }, {
    "tStartMs": 71240,
    "dDurationMs": 5019,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 71250,
    "dDurationMs": 8880,
    "wWinId": 1,
    "segs": [ {
      "utf8": "some",
      "acAsrConf": 239
    }, {
      "utf8": " blanks",
      "tOffsetMs": 979,
      "acAsrConf": 241
    }, {
      "utf8": " where",
      "tOffsetMs": 1979,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 2250,
      "acAsrConf": 252
    }, {
      "utf8": " can",
      "tOffsetMs": 2580,
      "acAsrConf": 252
    }, {
      "utf8": " fill",
      "tOffsetMs": 3439,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 4439,
      "acAsrConf": 252
    }, {
      "utf8": " any",
      "tOffsetMs": 4650,
      "acAsrConf": 221
    } ]
  }, {
    "tStartMs": 76249,
    "dDurationMs": 3881,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 76259,
    "dDurationMs": 6150,
    "wWinId": 1,
    "segs": [ {
      "utf8": "answers",
      "acAsrConf": 252
    }, {
      "utf8": " that",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " we",
      "tOffsetMs": 1021,
      "acAsrConf": 252
    }, {
      "utf8": " talked",
      "tOffsetMs": 1171,
      "acAsrConf": 224
    }, {
      "utf8": " about",
      "tOffsetMs": 1411,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1441,
      "acAsrConf": 115
    }, {
      "utf8": " if",
      "tOffsetMs": 2271,
      "acAsrConf": 131
    }, {
      "utf8": " you",
      "tOffsetMs": 3271,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 80120,
    "dDurationMs": 2289,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 80130,
    "dDurationMs": 4980,
    "wWinId": 1,
    "segs": [ {
      "utf8": "want",
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 239,
      "acAsrConf": 208
    }, {
      "utf8": " take",
      "tOffsetMs": 450,
      "acAsrConf": 252
    }, {
      "utf8": " notes",
      "tOffsetMs": 900,
      "acAsrConf": 252
    }, {
      "utf8": " or",
      "tOffsetMs": 1230,
      "acAsrConf": 226
    }, {
      "utf8": " if",
      "tOffsetMs": 1620,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1739,
      "acAsrConf": 252
    }, {
      "utf8": " want",
      "tOffsetMs": 1919,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 2160,
      "acAsrConf": 208
    } ]
  }, {
    "tStartMs": 82399,
    "dDurationMs": 2711,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 82409,
    "dDurationMs": 5011,
    "wWinId": 1,
    "segs": [ {
      "utf8": "remember",
      "acAsrConf": 252
    }, {
      "utf8": " what",
      "tOffsetMs": 421,
      "acAsrConf": 226
    }, {
      "utf8": " we",
      "tOffsetMs": 780,
      "acAsrConf": 252
    }, {
      "utf8": " said",
      "tOffsetMs": 931,
      "acAsrConf": 252
    }, {
      "utf8": " but",
      "tOffsetMs": 1171,
      "acAsrConf": 199
    }, {
      "utf8": " you",
      "tOffsetMs": 1951,
      "acAsrConf": 240
    }, {
      "utf8": " want",
      "tOffsetMs": 2401,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 2640,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 85100,
    "dDurationMs": 2320,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 85110,
    "dDurationMs": 7020,
    "wWinId": 1,
    "segs": [ {
      "utf8": "listen",
      "acAsrConf": 152
    }, {
      "utf8": " at",
      "tOffsetMs": 570,
      "acAsrConf": 227
    }, {
      "utf8": " the",
      "tOffsetMs": 780,
      "acAsrConf": 252
    }, {
      "utf8": " same",
      "tOffsetMs": 930,
      "acAsrConf": 200
    }, {
      "utf8": " time",
      "tOffsetMs": 1140,
      "acAsrConf": 248
    }, {
      "utf8": " it",
      "tOffsetMs": 1439,
      "acAsrConf": 249
    }, {
      "utf8": " might",
      "tOffsetMs": 1799,
      "acAsrConf": 252
    }, {
      "utf8": " be",
      "tOffsetMs": 2070,
      "acAsrConf": 208
    } ]
  }, {
    "tStartMs": 87410,
    "dDurationMs": 4720,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 87420,
    "dDurationMs": 8510,
    "wWinId": 1,
    "segs": [ {
      "utf8": "helpful",
      "acAsrConf": 242
    }, {
      "utf8": " to",
      "tOffsetMs": 269,
      "acAsrConf": 207
    }, {
      "utf8": " follow",
      "tOffsetMs": 600,
      "acAsrConf": 250
    }, {
      "utf8": " the",
      "tOffsetMs": 1559,
      "acAsrConf": 252
    }, {
      "utf8": " outline",
      "tOffsetMs": 2460,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 2610,
      "acAsrConf": 157
    }, {
      "utf8": " in",
      "tOffsetMs": 3239,
      "acAsrConf": 137
    }, {
      "utf8": " the",
      "tOffsetMs": 4170,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 92120,
    "dDurationMs": 3810,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 92130,
    "dDurationMs": 8460,
    "wWinId": 1,
    "segs": [ {
      "utf8": "on",
      "acAsrConf": 237
    }, {
      "utf8": " the",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " Google+",
      "tOffsetMs": 1290,
      "acAsrConf": 239
    }, {
      "utf8": " Google",
      "tOffsetMs": 2099,
      "acAsrConf": 236
    }, {
      "utf8": " Hangouts",
      "tOffsetMs": 2640,
      "acAsrConf": 250
    }, {
      "utf8": " page",
      "tOffsetMs": 3000,
      "acAsrConf": 255
    }, {
      "utf8": " I",
      "tOffsetMs": 3390,
      "acAsrConf": 89
    } ]
  }, {
    "tStartMs": 95920,
    "dDurationMs": 4670,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 95930,
    "dDurationMs": 9130,
    "wWinId": 1,
    "segs": [ {
      "utf8": "gave",
      "acAsrConf": 247
    }, {
      "utf8": " a",
      "tOffsetMs": 1000,
      "acAsrConf": 248
    }, {
      "utf8": " link",
      "tOffsetMs": 1329,
      "acAsrConf": 237
    }, {
      "utf8": " a",
      "tOffsetMs": 2170,
      "acAsrConf": 237
    }, {
      "utf8": " Google",
      "tOffsetMs": 2520,
      "acAsrConf": 125
    }, {
      "utf8": " Drive",
      "tOffsetMs": 3520,
      "acAsrConf": 252
    }, {
      "utf8": " link",
      "tOffsetMs": 3759,
      "acAsrConf": 236
    }, {
      "utf8": " where",
      "tOffsetMs": 4329,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 100580,
    "dDurationMs": 4480,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 100590,
    "dDurationMs": 6599,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 252
    }, {
      "utf8": " can",
      "tOffsetMs": 29,
      "acAsrConf": 252
    }, {
      "utf8": " download",
      "tOffsetMs": 209,
      "acAsrConf": 211
    }, {
      "utf8": " the",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " mind",
      "tOffsetMs": 1349,
      "acAsrConf": 248
    }, {
      "utf8": " map",
      "tOffsetMs": 1650,
      "acAsrConf": 248
    }, {
      "utf8": " so",
      "tOffsetMs": 2419,
      "acAsrConf": 241
    }, {
      "utf8": " I",
      "tOffsetMs": 3419,
      "acAsrConf": 220
    }, {
      "utf8": " hope",
      "tOffsetMs": 3840,
      "acAsrConf": 238
    } ]
  }, {
    "tStartMs": 105050,
    "dDurationMs": 2139,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 105060,
    "dDurationMs": 6629,
    "wWinId": 1,
    "segs": [ {
      "utf8": "that",
      "acAsrConf": 185
    }, {
      "utf8": " you",
      "tOffsetMs": 300,
      "acAsrConf": 250
    }, {
      "utf8": " have",
      "tOffsetMs": 449,
      "acAsrConf": 229
    }, {
      "utf8": " been",
      "tOffsetMs": 629,
      "acAsrConf": 203
    }, {
      "utf8": " able",
      "tOffsetMs": 780,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1019,
      "acAsrConf": 239
    }, {
      "utf8": " do",
      "tOffsetMs": 1199,
      "acAsrConf": 223
    }, {
      "utf8": " that",
      "tOffsetMs": 1500,
      "acAsrConf": 216
    }, {
      "utf8": " if",
      "tOffsetMs": 1710,
      "acAsrConf": 249
    } ]
  }, {
    "tStartMs": 107179,
    "dDurationMs": 4510,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 107189,
    "dDurationMs": 7021,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 930,
      "acAsrConf": 252
    }, {
      "utf8": " not",
      "tOffsetMs": 1970,
      "acAsrConf": 239
    }, {
      "utf8": " been",
      "tOffsetMs": 2970,
      "acAsrConf": 232
    }, {
      "utf8": " able",
      "tOffsetMs": 3390,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 3601,
      "acAsrConf": 221
    }, {
      "utf8": " do",
      "tOffsetMs": 3811,
      "acAsrConf": 221
    }, {
      "utf8": " that",
      "tOffsetMs": 4081,
      "acAsrConf": 211
    }, {
      "utf8": " if",
      "tOffsetMs": 4261,
      "acAsrConf": 243
    } ]
  }, {
    "tStartMs": 111679,
    "dDurationMs": 2531,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 111689,
    "dDurationMs": 5011,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you're",
      "acAsrConf": 252
    }, {
      "utf8": " not",
      "tOffsetMs": 121,
      "acAsrConf": 63
    }, {
      "utf8": " able",
      "tOffsetMs": 331,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 511,
      "acAsrConf": 234
    }, {
      "utf8": " do",
      "tOffsetMs": 691,
      "acAsrConf": 223
    }, {
      "utf8": " that",
      "tOffsetMs": 901,
      "acAsrConf": 201
    }, {
      "utf8": " go",
      "tOffsetMs": 1191,
      "acAsrConf": 250
    }, {
      "utf8": " to",
      "tOffsetMs": 2191,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 2220,
      "acAsrConf": 204
    } ]
  }, {
    "tStartMs": 114200,
    "dDurationMs": 2500,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 114210,
    "dDurationMs": 5040,
    "wWinId": 1,
    "segs": [ {
      "utf8": "main",
      "acAsrConf": 227
    }, {
      "utf8": " page",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 659,
      "acAsrConf": 216
    }, {
      "utf8": " this",
      "tOffsetMs": 720,
      "acAsrConf": 188
    }, {
      "utf8": " video",
      "tOffsetMs": 1229,
      "acAsrConf": 236
    }, {
      "utf8": " and",
      "tOffsetMs": 1830,
      "acAsrConf": 200
    }, {
      "utf8": " you'll",
      "tOffsetMs": 2220,
      "acAsrConf": 252
    }, {
      "utf8": " be",
      "tOffsetMs": 2339,
      "acAsrConf": 202
    } ]
  }, {
    "tStartMs": 116690,
    "dDurationMs": 2560,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 116700,
    "dDurationMs": 5640,
    "wWinId": 1,
    "segs": [ {
      "utf8": "able",
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 120,
      "acAsrConf": 220
    }, {
      "utf8": " find",
      "tOffsetMs": 360,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 540,
      "acAsrConf": 200
    }, {
      "utf8": " I",
      "tOffsetMs": 779,
      "acAsrConf": 217
    }, {
      "utf8": " will",
      "tOffsetMs": 959,
      "acAsrConf": 216
    }, {
      "utf8": " also",
      "tOffsetMs": 1470,
      "acAsrConf": 232
    }, {
      "utf8": " send",
      "tOffsetMs": 1919,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 2160,
      "acAsrConf": 237
    }, {
      "utf8": " to",
      "tOffsetMs": 2370,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 119240,
    "dDurationMs": 3100,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 119250,
    "dDurationMs": 6740,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 252
    }, {
      "utf8": " try",
      "tOffsetMs": 680,
      "acAsrConf": 246
    }, {
      "utf8": " to",
      "tOffsetMs": 1680,
      "acAsrConf": 252
    }, {
      "utf8": " figure",
      "tOffsetMs": 1710,
      "acAsrConf": 202
    }, {
      "utf8": " out",
      "tOffsetMs": 2130,
      "acAsrConf": 166
    }, {
      "utf8": " how",
      "tOffsetMs": 2280,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 2460,
      "acAsrConf": 217
    }, {
      "utf8": " send",
      "tOffsetMs": 2520,
      "acAsrConf": 215
    }, {
      "utf8": " it",
      "tOffsetMs": 2850,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 2909,
      "acAsrConf": 145
    } ]
  }, {
    "tStartMs": 122330,
    "dDurationMs": 3660,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 122340,
    "dDurationMs": 3650,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 252
    }, {
      "utf8": " now",
      "tOffsetMs": 120,
      "acAsrConf": 201
    }, {
      "utf8": " let's",
      "tOffsetMs": 750,
      "acAsrConf": 236
    }, {
      "utf8": " see",
      "tOffsetMs": 1440,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 127789,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 127799,
    "dDurationMs": 12351,
    "wWinId": 1,
    "segs": [ {
      "utf8": "all",
      "acAsrConf": 229
    }, {
      "utf8": " right",
      "tOffsetMs": 990,
      "acAsrConf": 229
    }, {
      "utf8": " I",
      "tOffsetMs": 1320,
      "acAsrConf": 194
    }, {
      "utf8": " will",
      "tOffsetMs": 1531,
      "acAsrConf": 252
    }, {
      "utf8": " try",
      "tOffsetMs": 1621,
      "acAsrConf": 165
    }, {
      "utf8": " to",
      "tOffsetMs": 2220,
      "acAsrConf": 207
    }, {
      "utf8": " send",
      "tOffsetMs": 2311,
      "acAsrConf": 236
    }, {
      "utf8": " this",
      "tOffsetMs": 2731,
      "acAsrConf": 223
    }, {
      "utf8": " link",
      "tOffsetMs": 4701,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 5701,
      "acAsrConf": 154
    } ]
  }, {
    "tStartMs": 135610,
    "dDurationMs": 4540,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 135620,
    "dDurationMs": 6310,
    "wWinId": 1,
    "segs": [ {
      "utf8": "think",
      "acAsrConf": 240
    }, {
      "utf8": " they",
      "tOffsetMs": 1000,
      "acAsrConf": 252
    }, {
      "utf8": " may",
      "tOffsetMs": 1149,
      "acAsrConf": 252
    }, {
      "utf8": " be",
      "tOffsetMs": 1330,
      "acAsrConf": 215
    }, {
      "utf8": " stuck",
      "tOffsetMs": 1390,
      "acAsrConf": 219
    }, {
      "utf8": " all",
      "tOffsetMs": 2729,
      "acAsrConf": 180
    }, {
      "utf8": " ranks",
      "tOffsetMs": 3729,
      "acAsrConf": 222
    } ]
  }, {
    "tStartMs": 140140,
    "dDurationMs": 1790,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 140150,
    "dDurationMs": 5320,
    "wWinId": 1,
    "segs": [ {
      "utf8": "downloading",
      "acAsrConf": 236
    }, {
      "utf8": " oh",
      "tOffsetMs": 1000,
      "acAsrConf": 227
    }, {
      "utf8": " there",
      "tOffsetMs": 1150,
      "acAsrConf": 227
    }, {
      "utf8": " are",
      "tOffsetMs": 1419,
      "acAsrConf": 230
    }, {
      "utf8": " so",
      "tOffsetMs": 1479,
      "acAsrConf": 252
    }, {
      "utf8": " many",
      "tOffsetMs": 1660,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 141920,
    "dDurationMs": 3550,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 141930,
    "dDurationMs": 5970,
    "wWinId": 1,
    "segs": [ {
      "utf8": "questions",
      "acAsrConf": 236
    }, {
      "utf8": " that's",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " right",
      "tOffsetMs": 870,
      "acAsrConf": 252
    }, {
      "utf8": " yes",
      "tOffsetMs": 1940,
      "acAsrConf": 146
    }, {
      "utf8": " yeah",
      "tOffsetMs": 2940,
      "acAsrConf": 226
    }, {
      "utf8": " we",
      "tOffsetMs": 3270,
      "acAsrConf": 252
    }, {
      "utf8": " don't",
      "tOffsetMs": 3419,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 145460,
    "dDurationMs": 2440,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 145470,
    "dDurationMs": 4650,
    "wWinId": 1,
    "segs": [ {
      "utf8": "know",
      "acAsrConf": 194
    }, {
      "utf8": " how",
      "tOffsetMs": 210,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 360,
      "acAsrConf": 206
    }, {
      "utf8": " send",
      "tOffsetMs": 420,
      "acAsrConf": 252
    }, {
      "utf8": " things",
      "tOffsetMs": 720,
      "acAsrConf": 206
    }, {
      "utf8": " to",
      "tOffsetMs": 1019,
      "acAsrConf": 200
    }, {
      "utf8": " you",
      "tOffsetMs": 1349,
      "acAsrConf": 252
    }, {
      "utf8": " yes",
      "tOffsetMs": 1560,
      "acAsrConf": 237
    }, {
      "utf8": " okay",
      "tOffsetMs": 1950,
      "acAsrConf": 203
    } ]
  }, {
    "tStartMs": 147890,
    "dDurationMs": 2230,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 147900,
    "dDurationMs": 4470,
    "wWinId": 1,
    "segs": [ {
      "utf8": "we",
      "acAsrConf": 226
    }, {
      "utf8": " will",
      "tOffsetMs": 330,
      "acAsrConf": 220
    }, {
      "utf8": " figure",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 690,
      "acAsrConf": 247
    }, {
      "utf8": " out",
      "tOffsetMs": 930,
      "acAsrConf": 252
    }, {
      "utf8": " don't",
      "tOffsetMs": 1110,
      "acAsrConf": 210
    }, {
      "utf8": " worry",
      "tOffsetMs": 1410,
      "acAsrConf": 206
    }, {
      "utf8": " let's",
      "tOffsetMs": 1619,
      "acAsrConf": 217
    } ]
  }, {
    "tStartMs": 150110,
    "dDurationMs": 2260,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 150120,
    "dDurationMs": 4380,
    "wWinId": 1,
    "segs": [ {
      "utf8": "just",
      "acAsrConf": 222
    }, {
      "utf8": " try",
      "tOffsetMs": 210,
      "acAsrConf": 236
    }, {
      "utf8": " to",
      "tOffsetMs": 390,
      "acAsrConf": 201
    }, {
      "utf8": " follow",
      "tOffsetMs": 449,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 660,
      "acAsrConf": 208
    }, {
      "utf8": " outline",
      "tOffsetMs": 990,
      "acAsrConf": 252
    }, {
      "utf8": " yeah",
      "tOffsetMs": 1080,
      "acAsrConf": 201
    }, {
      "utf8": " yeah",
      "tOffsetMs": 1860,
      "acAsrConf": 182
    } ]
  }, {
    "tStartMs": 152360,
    "dDurationMs": 2140,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 152370,
    "dDurationMs": 5699,
    "wWinId": 1,
    "segs": [ {
      "utf8": "we'll",
      "acAsrConf": 249
    }, {
      "utf8": " only",
      "tOffsetMs": 240,
      "acAsrConf": 133
    }, {
      "utf8": " have",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 720,
      "acAsrConf": 200
    }, {
      "utf8": " outline",
      "tOffsetMs": 1229,
      "acAsrConf": 252
    }, {
      "utf8": " try",
      "tOffsetMs": 1589,
      "acAsrConf": 166
    }, {
      "utf8": " to",
      "tOffsetMs": 2040,
      "acAsrConf": 216
    } ]
  }, {
    "tStartMs": 154490,
    "dDurationMs": 3579,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 154500,
    "dDurationMs": 13170,
    "wWinId": 1,
    "segs": [ {
      "utf8": "follow",
      "acAsrConf": 234
    }, {
      "utf8": " with",
      "tOffsetMs": 239,
      "acAsrConf": 236
    }, {
      "utf8": " us",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " yes",
      "tOffsetMs": 660,
      "acAsrConf": 204
    }, {
      "utf8": " yes",
      "tOffsetMs": 1700,
      "acAsrConf": 209
    }, {
      "utf8": " wonderful",
      "tOffsetMs": 2700,
      "acAsrConf": 233
    } ]
  }, {
    "tStartMs": 158059,
    "dDurationMs": 9611,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 158069,
    "dDurationMs": 13111,
    "wWinId": 1,
    "segs": [ {
      "utf8": "oh",
      "acAsrConf": 75
    }, {
      "utf8": " all",
      "tOffsetMs": 4461,
      "acAsrConf": 147
    }, {
      "utf8": " right",
      "tOffsetMs": 5461,
      "acAsrConf": 234
    }, {
      "utf8": " are",
      "tOffsetMs": 5811,
      "acAsrConf": 147
    }, {
      "utf8": " we",
      "tOffsetMs": 6811,
      "acAsrConf": 252
    }, {
      "utf8": " good",
      "tOffsetMs": 6841,
      "acAsrConf": 213
    }, {
      "utf8": " yes",
      "tOffsetMs": 7021,
      "acAsrConf": 219
    }, {
      "utf8": " all",
      "tOffsetMs": 8571,
      "acAsrConf": 204
    }, {
      "utf8": " right",
      "tOffsetMs": 9571,
      "acAsrConf": 219
    } ]
  }, {
    "tStartMs": 167660,
    "dDurationMs": 3520,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 167670,
    "dDurationMs": 8280,
    "wWinId": 1,
    "segs": [ {
      "utf8": "if",
      "acAsrConf": 158
    }, {
      "utf8": " you",
      "tOffsetMs": 270,
      "acAsrConf": 231
    }, {
      "utf8": " are",
      "tOffsetMs": 480,
      "acAsrConf": 171
    }, {
      "utf8": " watching",
      "tOffsetMs": 659,
      "acAsrConf": 151
    }, {
      "utf8": " can",
      "tOffsetMs": 1170,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1440,
      "acAsrConf": 218
    }, {
      "utf8": " please",
      "tOffsetMs": 1649,
      "acAsrConf": 252
    }, {
      "utf8": " mute",
      "tOffsetMs": 2510,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 171170,
    "dDurationMs": 4780,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 171180,
    "dDurationMs": 7440,
    "wWinId": 1,
    "segs": [ {
      "utf8": "your",
      "acAsrConf": 165
    }, {
      "utf8": " sound",
      "tOffsetMs": 320,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1320,
      "acAsrConf": 226
    }, {
      "utf8": " turn",
      "tOffsetMs": 1770,
      "acAsrConf": 227
    }, {
      "utf8": " off",
      "tOffsetMs": 2460,
      "acAsrConf": 222
    }, {
      "utf8": " your",
      "tOffsetMs": 2820,
      "acAsrConf": 202
    }, {
      "utf8": " video",
      "tOffsetMs": 3450,
      "acAsrConf": 208
    }, {
      "utf8": " that",
      "tOffsetMs": 4050,
      "acAsrConf": 218
    } ]
  }, {
    "tStartMs": 175940,
    "dDurationMs": 2680,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 175950,
    "dDurationMs": 10319,
    "wWinId": 1,
    "segs": [ {
      "utf8": "would",
      "acAsrConf": 238
    }, {
      "utf8": " be",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " great",
      "tOffsetMs": 420,
      "acAsrConf": 255
    }, {
      "utf8": " that",
      "tOffsetMs": 890,
      "acAsrConf": 202
    }, {
      "utf8": " would",
      "tOffsetMs": 1890,
      "acAsrConf": 216
    }, {
      "utf8": " be",
      "tOffsetMs": 2069,
      "acAsrConf": 232
    }, {
      "utf8": " great",
      "tOffsetMs": 2190,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 2430,
      "acAsrConf": 93
    } ]
  }, {
    "tStartMs": 178610,
    "dDurationMs": 7659,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 178620,
    "dDurationMs": 8670,
    "wWinId": 1,
    "segs": [ {
      "utf8": "think",
      "acAsrConf": 216
    }, {
      "utf8": " yes",
      "tOffsetMs": 4610,
      "acAsrConf": 237
    }, {
      "utf8": " yes",
      "tOffsetMs": 5610,
      "acAsrConf": 201
    }, {
      "utf8": " all",
      "tOffsetMs": 6030,
      "acAsrConf": 124
    }, {
      "utf8": " right",
      "tOffsetMs": 6270,
      "acAsrConf": 233
    }, {
      "utf8": " so",
      "tOffsetMs": 6330,
      "acAsrConf": 167
    }, {
      "utf8": " first",
      "tOffsetMs": 6780,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 7320,
      "acAsrConf": 244
    } ]
  }, {
    "tStartMs": 186259,
    "dDurationMs": 1031,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 186269,
    "dDurationMs": 2881,
    "wWinId": 1,
    "segs": [ {
      "utf8": "first",
      "acAsrConf": 201
    }, {
      "utf8": " things",
      "tOffsetMs": 181,
      "acAsrConf": 220
    }, {
      "utf8": " we're",
      "tOffsetMs": 481,
      "acAsrConf": 252
    }, {
      "utf8": " going",
      "tOffsetMs": 661,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 780,
      "acAsrConf": 252
    }, {
      "utf8": " talk",
      "tOffsetMs": 810,
      "acAsrConf": 206
    }, {
      "utf8": " about",
      "tOffsetMs": 901,
      "acAsrConf": 106
    } ]
  }, {
    "tStartMs": 187280,
    "dDurationMs": 1870,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 187290,
    "dDurationMs": 3960,
    "wWinId": 1,
    "segs": [ {
      "utf8": "first",
      "acAsrConf": 187
    }, {
      "utf8": " I'm",
      "tOffsetMs": 690,
      "acAsrConf": 248
    }, {
      "utf8": " going",
      "tOffsetMs": 900,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1080,
      "acAsrConf": 205
    }, {
      "utf8": " give",
      "tOffsetMs": 1199,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1380,
      "acAsrConf": 250
    }, {
      "utf8": " a",
      "tOffsetMs": 1470,
      "acAsrConf": 252
    }, {
      "utf8": " little",
      "tOffsetMs": 1500,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 189140,
    "dDurationMs": 2110,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 189150,
    "dDurationMs": 5550,
    "wWinId": 1,
    "segs": [ {
      "utf8": "personal",
      "acAsrConf": 252
    }, {
      "utf8": " introduction",
      "tOffsetMs": 559,
      "acAsrConf": 249
    }, {
      "utf8": " so",
      "tOffsetMs": 1559,
      "acAsrConf": 233
    }, {
      "utf8": " we",
      "tOffsetMs": 1800,
      "acAsrConf": 252
    }, {
      "utf8": " can",
      "tOffsetMs": 1949,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 191240,
    "dDurationMs": 3460,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 191250,
    "dDurationMs": 5609,
    "wWinId": 1,
    "segs": [ {
      "utf8": "introduce",
      "acAsrConf": 252
    }, {
      "utf8": " ourselves",
      "tOffsetMs": 299,
      "acAsrConf": 229
    }, {
      "utf8": " some",
      "tOffsetMs": 859,
      "acAsrConf": 146
    }, {
      "utf8": " of",
      "tOffsetMs": 1859,
      "acAsrConf": 226
    }, {
      "utf8": " you",
      "tOffsetMs": 2040,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 2450,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 194690,
    "dDurationMs": 2169,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 194700,
    "dDurationMs": 4409,
    "wWinId": 1,
    "segs": [ {
      "utf8": "already",
      "acAsrConf": 252
    }, {
      "utf8": " met",
      "tOffsetMs": 330,
      "acAsrConf": 220
    }, {
      "utf8": " us",
      "tOffsetMs": 780,
      "acAsrConf": 217
    }, {
      "utf8": " some",
      "tOffsetMs": 1020,
      "acAsrConf": 173
    }, {
      "utf8": " of",
      "tOffsetMs": 1379,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1530,
      "acAsrConf": 252
    }, {
      "utf8": " I've",
      "tOffsetMs": 1680,
      "acAsrConf": 238
    }, {
      "utf8": " talked",
      "tOffsetMs": 1890,
      "acAsrConf": 246
    } ]
  }, {
    "tStartMs": 196849,
    "dDurationMs": 2260,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 196859,
    "dDurationMs": 7321,
    "wWinId": 1,
    "segs": [ {
      "utf8": "with",
      "acAsrConf": 255
    }, {
      "utf8": " on",
      "tOffsetMs": 91,
      "acAsrConf": 209
    }, {
      "utf8": " skype",
      "tOffsetMs": 301,
      "acAsrConf": 252
    }, {
      "utf8": " but",
      "tOffsetMs": 630,
      "acAsrConf": 141
    }, {
      "utf8": " if",
      "tOffsetMs": 841,
      "acAsrConf": 189
    }, {
      "utf8": " you",
      "tOffsetMs": 1350,
      "acAsrConf": 252
    }, {
      "utf8": " haven't",
      "tOffsetMs": 1621,
      "acAsrConf": 215
    }, {
      "utf8": " met",
      "tOffsetMs": 1951,
      "acAsrConf": 252
    }, {
      "utf8": " us",
      "tOffsetMs": 2041,
      "acAsrConf": 226
    } ]
  }, {
    "tStartMs": 199099,
    "dDurationMs": 5081,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 199109,
    "dDurationMs": 7981,
    "wWinId": 1,
    "segs": [ {
      "utf8": "I'm",
      "acAsrConf": 216
    }, {
      "utf8": " Vanessa",
      "tOffsetMs": 571,
      "acAsrConf": 242
    }, {
      "utf8": " and",
      "tOffsetMs": 1291,
      "acAsrConf": 78
    }, {
      "utf8": " you",
      "tOffsetMs": 1731,
      "acAsrConf": 223
    }, {
      "utf8": " are",
      "tOffsetMs": 2731,
      "acAsrConf": 252
    }, {
      "utf8": " I'm",
      "tOffsetMs": 2880,
      "acAsrConf": 166
    }, {
      "utf8": " Dan",
      "tOffsetMs": 3361,
      "acAsrConf": 220
    }, {
      "utf8": " hi",
      "tOffsetMs": 3780,
      "acAsrConf": 233
    } ]
  }, {
    "tStartMs": 204170,
    "dDurationMs": 2920,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 204180,
    "dDurationMs": 5160,
    "wWinId": 1,
    "segs": [ {
      "utf8": "this",
      "acAsrConf": 130
    }, {
      "utf8": " is",
      "tOffsetMs": 389,
      "acAsrConf": 224
    }, {
      "utf8": " my",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " husband",
      "tOffsetMs": 630,
      "acAsrConf": 252
    }, {
      "utf8": " Dan",
      "tOffsetMs": 1020,
      "acAsrConf": 217
    }, {
      "utf8": " and",
      "tOffsetMs": 1230,
      "acAsrConf": 243
    }, {
      "utf8": " I",
      "tOffsetMs": 1559,
      "acAsrConf": 236
    }, {
      "utf8": " run",
      "tOffsetMs": 2250,
      "acAsrConf": 255
    }, {
      "utf8": " the",
      "tOffsetMs": 2790,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 207080,
    "dDurationMs": 2260,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 207090,
    "dDurationMs": 4709,
    "wWinId": 1,
    "segs": [ {
      "utf8": "website",
      "acAsrConf": 245
    }, {
      "utf8": " speak",
      "tOffsetMs": 179,
      "acAsrConf": 243
    }, {
      "utf8": " English",
      "tOffsetMs": 750,
      "acAsrConf": 242
    }, {
      "utf8": " with",
      "tOffsetMs": 1140,
      "acAsrConf": 250
    }, {
      "utf8": " Vanessa",
      "tOffsetMs": 1380,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1709,
      "acAsrConf": 206
    }, {
      "utf8": " I",
      "tOffsetMs": 2220,
      "acAsrConf": 241
    } ]
  }, {
    "tStartMs": 209330,
    "dDurationMs": 2469,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 209340,
    "dDurationMs": 5519,
    "wWinId": 1,
    "segs": [ {
      "utf8": "get",
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 299,
      "acAsrConf": 252
    }, {
      "utf8": " chance",
      "tOffsetMs": 330,
      "acAsrConf": 203
    }, {
      "utf8": " to",
      "tOffsetMs": 660,
      "acAsrConf": 206
    }, {
      "utf8": " speak",
      "tOffsetMs": 989,
      "acAsrConf": 252
    }, {
      "utf8": " with",
      "tOffsetMs": 1229,
      "acAsrConf": 252
    }, {
      "utf8": " many",
      "tOffsetMs": 1290,
      "acAsrConf": 217
    }, {
      "utf8": " of",
      "tOffsetMs": 2040,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 2220,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 211789,
    "dDurationMs": 3070,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 211799,
    "dDurationMs": 5041,
    "wWinId": 1,
    "segs": [ {
      "utf8": "each",
      "acAsrConf": 236
    }, {
      "utf8": " week",
      "tOffsetMs": 660,
      "acAsrConf": 216
    }, {
      "utf8": " and",
      "tOffsetMs": 1051,
      "acAsrConf": 249
    }, {
      "utf8": " make",
      "tOffsetMs": 1321,
      "acAsrConf": 252
    }, {
      "utf8": " videos",
      "tOffsetMs": 1891,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 2310,
      "acAsrConf": 239
    }, {
      "utf8": " have",
      "tOffsetMs": 2761,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 3030,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 214849,
    "dDurationMs": 1991,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 214859,
    "dDurationMs": 7731,
    "wWinId": 1,
    "segs": [ {
      "utf8": "chance",
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 271,
      "acAsrConf": 203
    }, {
      "utf8": " really",
      "tOffsetMs": 511,
      "acAsrConf": 252
    }, {
      "utf8": " help",
      "tOffsetMs": 660,
      "acAsrConf": 236
    }, {
      "utf8": " you",
      "tOffsetMs": 1111,
      "acAsrConf": 252
    }, {
      "utf8": " learn",
      "tOffsetMs": 1291,
      "acAsrConf": 252
    }, {
      "utf8": " English",
      "tOffsetMs": 1591,
      "acAsrConf": 223
    } ]
  }, {
    "tStartMs": 216830,
    "dDurationMs": 5760,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 216840,
    "dDurationMs": 11070,
    "wWinId": 1,
    "segs": [ {
      "utf8": "and",
      "acAsrConf": 231
    }, {
      "utf8": " we",
      "tOffsetMs": 269,
      "acAsrConf": 243
    }, {
      "utf8": " live",
      "tOffsetMs": 1080,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 1320,
      "acAsrConf": 236
    }, {
      "utf8": " the",
      "tOffsetMs": 1649,
      "acAsrConf": 252
    }, {
      "utf8": " eastern",
      "tOffsetMs": 2489,
      "acAsrConf": 252
    }, {
      "utf8": " US",
      "tOffsetMs": 3179,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 3630,
      "acAsrConf": 205
    }, {
      "utf8": " about",
      "tOffsetMs": 4380,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 222580,
    "dDurationMs": 5330,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 222590,
    "dDurationMs": 7840,
    "wWinId": 1,
    "segs": [ {
      "utf8": "15",
      "acAsrConf": 238
    }, {
      "utf8": " hours",
      "tOffsetMs": 1000,
      "acAsrConf": 252
    }, {
      "utf8": " south",
      "tOffsetMs": 1420,
      "acAsrConf": 206
    }, {
      "utf8": " of",
      "tOffsetMs": 2350,
      "acAsrConf": 252
    }, {
      "utf8": " New",
      "tOffsetMs": 2850,
      "acAsrConf": 217
    }, {
      "utf8": " York",
      "tOffsetMs": 3850,
      "acAsrConf": 252
    }, {
      "utf8": " City",
      "tOffsetMs": 4149,
      "acAsrConf": 252
    }, {
      "utf8": " very",
      "tOffsetMs": 4450,
      "acAsrConf": 166
    }, {
      "utf8": " far",
      "tOffsetMs": 4959,
      "acAsrConf": 204
    } ]
  }, {
    "tStartMs": 227900,
    "dDurationMs": 2530,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 227910,
    "dDurationMs": 5189,
    "wWinId": 1,
    "segs": [ {
      "utf8": "from",
      "acAsrConf": 252
    }, {
      "utf8": " New",
      "tOffsetMs": 240,
      "acAsrConf": 207
    }, {
      "utf8": " York",
      "tOffsetMs": 419,
      "acAsrConf": 251
    }, {
      "utf8": " but",
      "tOffsetMs": 630,
      "acAsrConf": 217
    }, {
      "utf8": " in",
      "tOffsetMs": 1500,
      "acAsrConf": 227
    }, {
      "utf8": " the",
      "tOffsetMs": 1770,
      "acAsrConf": 252
    }, {
      "utf8": " same",
      "tOffsetMs": 2010,
      "acAsrConf": 229
    }, {
      "utf8": " time",
      "tOffsetMs": 2220,
      "acAsrConf": 252
    }, {
      "utf8": " zone",
      "tOffsetMs": 2490,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 230420,
    "dDurationMs": 2679,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 230430,
    "dDurationMs": 8399,
    "wWinId": 1,
    "segs": [ {
      "utf8": "so",
      "acAsrConf": 135
    }, {
      "utf8": " we",
      "tOffsetMs": 630,
      "acAsrConf": 252
    }, {
      "utf8": " live",
      "tOffsetMs": 839,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 989,
      "acAsrConf": 158
    }, {
      "utf8": " a",
      "tOffsetMs": 1110,
      "acAsrConf": 252
    }, {
      "utf8": " city",
      "tOffsetMs": 1200,
      "acAsrConf": 252
    }, {
      "utf8": " called",
      "tOffsetMs": 1380,
      "acAsrConf": 228
    }, {
      "utf8": " Asheville",
      "tOffsetMs": 1669,
      "acAsrConf": 217
    } ]
  }, {
    "tStartMs": 233089,
    "dDurationMs": 5740,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 233099,
    "dDurationMs": 10860,
    "wWinId": 1,
    "segs": [ {
      "utf8": "North",
      "acAsrConf": 252
    }, {
      "utf8": " Carolina",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 841,
      "acAsrConf": 224
    }, {
      "utf8": " we",
      "tOffsetMs": 2871,
      "acAsrConf": 242
    }, {
      "utf8": " just",
      "tOffsetMs": 3871,
      "acAsrConf": 255
    }, {
      "utf8": " moved",
      "tOffsetMs": 4621,
      "acAsrConf": 244
    }, {
      "utf8": " here",
      "tOffsetMs": 5101,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 238819,
    "dDurationMs": 5140,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 238829,
    "dDurationMs": 7020,
    "wWinId": 1,
    "segs": [ {
      "utf8": "about",
      "acAsrConf": 252
    }, {
      "utf8": " six",
      "tOffsetMs": 121,
      "acAsrConf": 252
    }, {
      "utf8": " months",
      "tOffsetMs": 1190,
      "acAsrConf": 179
    }, {
      "utf8": " ago",
      "tOffsetMs": 2190,
      "acAsrConf": 252
    }, {
      "utf8": " from",
      "tOffsetMs": 2310,
      "acAsrConf": 177
    }, {
      "utf8": " Korea",
      "tOffsetMs": 3301,
      "acAsrConf": 218
    }, {
      "utf8": " and",
      "tOffsetMs": 4231,
      "acAsrConf": 247
    }, {
      "utf8": " I",
      "tOffsetMs": 4530,
      "acAsrConf": 221
    } ]
  }, {
    "tStartMs": 243949,
    "dDurationMs": 1900,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 243959,
    "dDurationMs": 6120,
    "wWinId": 1,
    "segs": [ {
      "utf8": "see",
      "acAsrConf": 252
    }, {
      "utf8": " we",
      "tOffsetMs": 241,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 421,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 451,
      "acAsrConf": 92
    }, {
      "utf8": " comment",
      "tOffsetMs": 601,
      "acAsrConf": 252
    }, {
      "utf8": " from",
      "tOffsetMs": 930,
      "acAsrConf": 252
    }, {
      "utf8": " your",
      "tOffsetMs": 1140,
      "acAsrConf": 64
    }, {
      "utf8": " son",
      "tOffsetMs": 1471,
      "acAsrConf": 222
    }, {
      "utf8": " Kim",
      "tOffsetMs": 1650,
      "acAsrConf": 226
    } ]
  }, {
    "tStartMs": 245839,
    "dDurationMs": 4240,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 245849,
    "dDurationMs": 6450,
    "wWinId": 1,
    "segs": [ {
      "utf8": "who",
      "acAsrConf": 252
    }, {
      "utf8": " says",
      "tOffsetMs": 211,
      "acAsrConf": 252
    }, {
      "utf8": " he",
      "tOffsetMs": 450,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 901,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 931,
      "acAsrConf": 206
    }, {
      "utf8": " Korea",
      "tOffsetMs": 1320,
      "acAsrConf": 255
    }, {
      "utf8": " too",
      "tOffsetMs": 1651,
      "acAsrConf": 128
    }, {
      "utf8": " oh",
      "tOffsetMs": 1950,
      "acAsrConf": 135
    }, {
      "utf8": " yeah",
      "tOffsetMs": 2331,
      "acAsrConf": 244
    }, {
      "utf8": " go",
      "tOffsetMs": 3420,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 250069,
    "dDurationMs": 2230,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 250079,
    "dDurationMs": 5280,
    "wWinId": 1,
    "segs": [ {
      "utf8": "Korea",
      "acAsrConf": 252
    }, {
      "utf8": " yes",
      "tOffsetMs": 361,
      "acAsrConf": 139
    }, {
      "utf8": " I",
      "tOffsetMs": 1020,
      "acAsrConf": 255
    }, {
      "utf8": " think",
      "tOffsetMs": 1171,
      "acAsrConf": 73
    }, {
      "utf8": " I",
      "tOffsetMs": 1440,
      "acAsrConf": 239
    }, {
      "utf8": " see",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    }, {
      "utf8": " there's",
      "tOffsetMs": 1770,
      "acAsrConf": 252
    }, {
      "utf8": " also",
      "tOffsetMs": 1981,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 252289,
    "dDurationMs": 3070,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 252299,
    "dDurationMs": 5400,
    "wWinId": 1,
    "segs": [ {
      "utf8": "some",
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 361,
      "acAsrConf": 200
    }, {
      "utf8": " you",
      "tOffsetMs": 481,
      "acAsrConf": 252
    }, {
      "utf8": " here",
      "tOffsetMs": 571,
      "acAsrConf": 227
    }, {
      "utf8": " from",
      "tOffsetMs": 901,
      "acAsrConf": 243
    }, {
      "utf8": " Brazil",
      "tOffsetMs": 1351,
      "acAsrConf": 226
    }, {
      "utf8": " excellent",
      "tOffsetMs": 1910,
      "acAsrConf": 145
    }, {
      "utf8": " I",
      "tOffsetMs": 2910,
      "acAsrConf": 165
    } ]
  }, {
    "tStartMs": 255349,
    "dDurationMs": 2350,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 255359,
    "dDurationMs": 5611,
    "wWinId": 1,
    "segs": [ {
      "utf8": "think",
      "acAsrConf": 219
    }, {
      "utf8": " some",
      "tOffsetMs": 181,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 331,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 361,
      "acAsrConf": 207
    }, {
      "utf8": " from",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " Europe",
      "tOffsetMs": 810,
      "acAsrConf": 238
    }, {
      "utf8": " it's",
      "tOffsetMs": 1201,
      "acAsrConf": 207
    }, {
      "utf8": " very",
      "tOffsetMs": 1981,
      "acAsrConf": 235
    } ]
  }, {
    "tStartMs": 257689,
    "dDurationMs": 3281,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 257699,
    "dDurationMs": 3660,
    "wWinId": 1,
    "segs": [ {
      "utf8": "late",
      "acAsrConf": 239
    }, {
      "utf8": " at",
      "tOffsetMs": 271,
      "acAsrConf": 252
    }, {
      "utf8": " night",
      "tOffsetMs": 421,
      "acAsrConf": 135
    }, {
      "utf8": " for",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 841,
      "acAsrConf": 226
    }, {
      "utf8": " wonderful",
      "tOffsetMs": 1021,
      "acAsrConf": 178
    }, {
      "utf8": " from",
      "tOffsetMs": 2060,
      "acAsrConf": 201
    }, {
      "utf8": " all",
      "tOffsetMs": 3060,
      "acAsrConf": 205
    } ]
  }, {
    "tStartMs": 260960,
    "dDurationMs": 399,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 260970,
    "dDurationMs": 2939,
    "wWinId": 1,
    "segs": [ {
      "utf8": "over",
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 30,
      "acAsrConf": 213
    } ]
  }, {
    "tStartMs": 261349,
    "dDurationMs": 2560,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 261359,
    "dDurationMs": 4650,
    "wWinId": 1,
    "segs": [ {
      "utf8": "world",
      "acAsrConf": 252
    }, {
      "utf8": " hello",
      "tOffsetMs": 210,
      "acAsrConf": 216
    }, {
      "utf8": " world",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " yes",
      "tOffsetMs": 720,
      "acAsrConf": 66
    }, {
      "utf8": " yes",
      "tOffsetMs": 1471,
      "acAsrConf": 252
    }, {
      "utf8": " as",
      "tOffsetMs": 1921,
      "acAsrConf": 203
    }, {
      "utf8": " some",
      "tOffsetMs": 2101,
      "acAsrConf": 237
    }, {
      "utf8": " from",
      "tOffsetMs": 2340,
      "acAsrConf": 205
    } ]
  }, {
    "tStartMs": 263899,
    "dDurationMs": 2110,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 263909,
    "dDurationMs": 4260,
    "wWinId": 1,
    "segs": [ {
      "utf8": "Russia",
      "acAsrConf": 214
    }, {
      "utf8": " excellent",
      "tOffsetMs": 451,
      "acAsrConf": 154
    }, {
      "utf8": " excellent",
      "tOffsetMs": 1440,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 265999,
    "dDurationMs": 2170,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 266009,
    "dDurationMs": 4351,
    "wWinId": 1,
    "segs": [ {
      "utf8": "so",
      "acAsrConf": 252
    }, {
      "utf8": " then",
      "tOffsetMs": 421,
      "acAsrConf": 230
    }, {
      "utf8": " would",
      "tOffsetMs": 870,
      "acAsrConf": 230
    }, {
      "utf8": " you",
      "tOffsetMs": 1141,
      "acAsrConf": 248
    }, {
      "utf8": " like",
      "tOffsetMs": 1261,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1291,
      "acAsrConf": 187
    }, {
      "utf8": " tell",
      "tOffsetMs": 1650,
      "acAsrConf": 252
    }, {
      "utf8": " them",
      "tOffsetMs": 1921,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 268159,
    "dDurationMs": 2201,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 268169,
    "dDurationMs": 4470,
    "wWinId": 1,
    "segs": [ {
      "utf8": "anything",
      "acAsrConf": 228
    }, {
      "utf8": " interesting",
      "tOffsetMs": 231,
      "acAsrConf": 217
    }, {
      "utf8": " about",
      "tOffsetMs": 1231,
      "acAsrConf": 200
    }, {
      "utf8": " our",
      "tOffsetMs": 1470,
      "acAsrConf": 241
    }, {
      "utf8": " city",
      "tOffsetMs": 1770,
      "acAsrConf": 208
    } ]
  }, {
    "tStartMs": 270350,
    "dDurationMs": 2289,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 270360,
    "dDurationMs": 5549,
    "wWinId": 1,
    "segs": [ {
      "utf8": "maybe",
      "acAsrConf": 208
    }, {
      "utf8": " some",
      "tOffsetMs": 540,
      "acAsrConf": 252
    }, {
      "utf8": " people",
      "tOffsetMs": 779,
      "acAsrConf": 248
    }, {
      "utf8": " are",
      "tOffsetMs": 929,
      "acAsrConf": 202
    }, {
      "utf8": " curious",
      "tOffsetMs": 1200,
      "acAsrConf": 221
    }, {
      "utf8": " our",
      "tOffsetMs": 1470,
      "acAsrConf": 204
    }, {
      "utf8": " city",
      "tOffsetMs": 1980,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 272629,
    "dDurationMs": 3280,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 272639,
    "dDurationMs": 6391,
    "wWinId": 1,
    "segs": [ {
      "utf8": "is",
      "acAsrConf": 251
    }, {
      "utf8": " quite",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " different",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " from",
      "tOffsetMs": 1291,
      "acAsrConf": 252
    }, {
      "utf8": " most",
      "tOffsetMs": 1740,
      "acAsrConf": 219
    }, {
      "utf8": " cities",
      "tOffsetMs": 2490,
      "acAsrConf": 236
    } ]
  }, {
    "tStartMs": 275899,
    "dDurationMs": 3131,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 275909,
    "dDurationMs": 4741,
    "wWinId": 1,
    "segs": [ {
      "utf8": "around",
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 271,
      "acAsrConf": 252
    }, {
      "utf8": " world",
      "tOffsetMs": 630,
      "acAsrConf": 252
    }, {
      "utf8": " it's",
      "tOffsetMs": 660,
      "acAsrConf": 207
    }, {
      "utf8": " very",
      "tOffsetMs": 1201,
      "acAsrConf": 143
    }, {
      "utf8": " small",
      "tOffsetMs": 1470,
      "acAsrConf": 218
    }, {
      "utf8": " so",
      "tOffsetMs": 1921,
      "acAsrConf": 201
    }, {
      "utf8": " if",
      "tOffsetMs": 2671,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 279020,
    "dDurationMs": 1630,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 279030,
    "dDurationMs": 4500,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you're",
      "acAsrConf": 252
    }, {
      "utf8": " from",
      "tOffsetMs": 150,
      "acAsrConf": 227
    }, {
      "utf8": " Korea",
      "tOffsetMs": 359,
      "acAsrConf": 236
    }, {
      "utf8": " you",
      "tOffsetMs": 569,
      "acAsrConf": 208
    }, {
      "utf8": " wouldn't",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    }, {
      "utf8": " see",
      "tOffsetMs": 1439,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 1590,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 280640,
    "dDurationMs": 2890,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 280650,
    "dDurationMs": 6750,
    "wWinId": 1,
    "segs": [ {
      "utf8": "city",
      "acAsrConf": 252
    }, {
      "utf8": " like",
      "tOffsetMs": 299,
      "acAsrConf": 252
    }, {
      "utf8": " this",
      "tOffsetMs": 509,
      "acAsrConf": 217
    }, {
      "utf8": " but",
      "tOffsetMs": 750,
      "acAsrConf": 226
    }, {
      "utf8": " it",
      "tOffsetMs": 1109,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 1229,
      "acAsrConf": 252
    }, {
      "utf8": " very",
      "tOffsetMs": 1259,
      "acAsrConf": 147
    }, {
      "utf8": " we",
      "tOffsetMs": 1759,
      "acAsrConf": 216
    }, {
      "utf8": " would",
      "tOffsetMs": 2759,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 283520,
    "dDurationMs": 3880,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 283530,
    "dDurationMs": 6840,
    "wWinId": 1,
    "segs": [ {
      "utf8": "say",
      "acAsrConf": 252
    }, {
      "utf8": " charming",
      "tOffsetMs": 150,
      "acAsrConf": 252
    }, {
      "utf8": " very",
      "tOffsetMs": 750,
      "acAsrConf": 207
    }, {
      "utf8": " interesting",
      "tOffsetMs": 2240,
      "acAsrConf": 241
    }, {
      "utf8": " city",
      "tOffsetMs": 3240,
      "acAsrConf": 176
    }, {
      "utf8": " so",
      "tOffsetMs": 3539,
      "acAsrConf": 255
    } ]
  }, {
    "tStartMs": 287390,
    "dDurationMs": 2980,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 287400,
    "dDurationMs": 5729,
    "wWinId": 1,
    "segs": [ {
      "utf8": "there",
      "acAsrConf": 246
    }, {
      "utf8": " are",
      "tOffsetMs": 449,
      "acAsrConf": 246
    }, {
      "utf8": " not",
      "tOffsetMs": 569,
      "acAsrConf": 217
    }, {
      "utf8": " many",
      "tOffsetMs": 810,
      "acAsrConf": 207
    }, {
      "utf8": " people",
      "tOffsetMs": 1049,
      "acAsrConf": 247
    }, {
      "utf8": " but",
      "tOffsetMs": 1380,
      "acAsrConf": 217
    }, {
      "utf8": " there's",
      "tOffsetMs": 2370,
      "acAsrConf": 250
    } ]
  }, {
    "tStartMs": 290360,
    "dDurationMs": 2769,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 290370,
    "dDurationMs": 7829,
    "wWinId": 1,
    "segs": [ {
      "utf8": "many",
      "acAsrConf": 252
    }, {
      "utf8": " things",
      "tOffsetMs": 780,
      "acAsrConf": 227
    }, {
      "utf8": " to",
      "tOffsetMs": 1740,
      "acAsrConf": 252
    }, {
      "utf8": " do",
      "tOffsetMs": 1979,
      "acAsrConf": 252
    }, {
      "utf8": " we",
      "tOffsetMs": 2130,
      "acAsrConf": 200
    }, {
      "utf8": " live",
      "tOffsetMs": 2430,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 2609,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 2639,
      "acAsrConf": 76
    } ]
  }, {
    "tStartMs": 293119,
    "dDurationMs": 5080,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 293129,
    "dDurationMs": 8820,
    "wWinId": 1,
    "segs": [ {
      "utf8": "mountains",
      "acAsrConf": 201
    }, {
      "utf8": " and",
      "tOffsetMs": 241,
      "acAsrConf": 136
    }, {
      "utf8": " by",
      "tOffsetMs": 1201,
      "acAsrConf": 240
    }, {
      "utf8": " rivers",
      "tOffsetMs": 1741,
      "acAsrConf": 248
    }, {
      "utf8": " and",
      "tOffsetMs": 2400,
      "acAsrConf": 233
    }, {
      "utf8": " there",
      "tOffsetMs": 3590,
      "acAsrConf": 227
    }, {
      "utf8": " is",
      "tOffsetMs": 4590,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 4831,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 298189,
    "dDurationMs": 3760,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 298199,
    "dDurationMs": 7620,
    "wWinId": 1,
    "segs": [ {
      "utf8": "lot",
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 421,
      "acAsrConf": 252
    }, {
      "utf8": " local",
      "tOffsetMs": 451,
      "acAsrConf": 213
    }, {
      "utf8": " business",
      "tOffsetMs": 1381,
      "acAsrConf": 252
    }, {
      "utf8": " here",
      "tOffsetMs": 2041,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 2370,
      "acAsrConf": 202
    }, {
      "utf8": " everybody",
      "tOffsetMs": 2750,
      "acAsrConf": 248
    } ]
  }, {
    "tStartMs": 301939,
    "dDurationMs": 3880,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 301949,
    "dDurationMs": 6960,
    "wWinId": 1,
    "segs": [ {
      "utf8": "is",
      "acAsrConf": 216
    }, {
      "utf8": " very",
      "tOffsetMs": 270,
      "acAsrConf": 247
    }, {
      "utf8": " interested",
      "tOffsetMs": 560,
      "acAsrConf": 207
    }, {
      "utf8": " in",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    }, {
      "utf8": " living",
      "tOffsetMs": 2060,
      "acAsrConf": 207
    }, {
      "utf8": " locally",
      "tOffsetMs": 3060,
      "acAsrConf": 252
    }, {
      "utf8": " we",
      "tOffsetMs": 3330,
      "acAsrConf": 204
    } ]
  }, {
    "tStartMs": 305809,
    "dDurationMs": 3100,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 305819,
    "dDurationMs": 5761,
    "wWinId": 1,
    "segs": [ {
      "utf8": "say",
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 181,
      "acAsrConf": 252
    }, {
      "utf8": " going",
      "tOffsetMs": 261,
      "acAsrConf": 207
    }, {
      "utf8": " to",
      "tOffsetMs": 1261,
      "acAsrConf": 203
    }, {
      "utf8": " the",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    }, {
      "utf8": " farmers",
      "tOffsetMs": 1710,
      "acAsrConf": 252
    }, {
      "utf8": " market",
      "tOffsetMs": 2190,
      "acAsrConf": 207
    }, {
      "utf8": " or",
      "tOffsetMs": 2761,
      "acAsrConf": 249
    } ]
  }, {
    "tStartMs": 308899,
    "dDurationMs": 2681,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 308909,
    "dDurationMs": 6051,
    "wWinId": 1,
    "segs": [ {
      "utf8": "going",
      "acAsrConf": 236
    }, {
      "utf8": " to",
      "tOffsetMs": 780,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 931,
      "acAsrConf": 252
    }, {
      "utf8": " local",
      "tOffsetMs": 1021,
      "acAsrConf": 252
    }, {
      "utf8": " restaurant",
      "tOffsetMs": 1260,
      "acAsrConf": 207
    }, {
      "utf8": " and",
      "tOffsetMs": 1980,
      "acAsrConf": 212
    }, {
      "utf8": " stuff",
      "tOffsetMs": 2370,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 311570,
    "dDurationMs": 3390,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 311580,
    "dDurationMs": 7019,
    "wWinId": 1,
    "segs": [ {
      "utf8": "like",
      "acAsrConf": 252
    }, {
      "utf8": " that",
      "tOffsetMs": 149,
      "acAsrConf": 222
    }, {
      "utf8": " yeah",
      "tOffsetMs": 389,
      "acAsrConf": 236
    }, {
      "utf8": " yeah",
      "tOffsetMs": 780,
      "acAsrConf": 203
    }, {
      "utf8": " I",
      "tOffsetMs": 1170,
      "acAsrConf": 211
    }, {
      "utf8": " work",
      "tOffsetMs": 1649,
      "acAsrConf": 249
    }, {
      "utf8": " at",
      "tOffsetMs": 2220,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 2399,
      "acAsrConf": 252
    }, {
      "utf8": " local",
      "tOffsetMs": 2429,
      "acAsrConf": 252
    }, {
      "utf8": " uh",
      "tOffsetMs": 2790,
      "acAsrConf": 107
    } ]
  }, {
    "tStartMs": 314950,
    "dDurationMs": 3649,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 314960,
    "dDurationMs": 6489,
    "wWinId": 1,
    "segs": [ {
      "utf8": "coffee",
      "acAsrConf": 232
    }, {
      "utf8": " shop",
      "tOffsetMs": 1000,
      "acAsrConf": 203
    }, {
      "utf8": " so",
      "tOffsetMs": 1359,
      "acAsrConf": 187
    }, {
      "utf8": " I",
      "tOffsetMs": 1780,
      "acAsrConf": 239
    }, {
      "utf8": " am",
      "tOffsetMs": 2049,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 2169,
      "acAsrConf": 252
    }, {
      "utf8": " barista",
      "tOffsetMs": 2199,
      "acAsrConf": 252
    }, {
      "utf8": " right",
      "tOffsetMs": 2560,
      "acAsrConf": 123
    }, {
      "utf8": " now",
      "tOffsetMs": 3429,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 318589,
    "dDurationMs": 2860,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 318599,
    "dDurationMs": 6211,
    "wWinId": 1,
    "segs": [ {
      "utf8": "and",
      "acAsrConf": 233
    }, {
      "utf8": " maybe",
      "tOffsetMs": 361,
      "acAsrConf": 234
    }, {
      "utf8": " someday",
      "tOffsetMs": 1051,
      "acAsrConf": 205
    }, {
      "utf8": " I'll",
      "tOffsetMs": 1470,
      "acAsrConf": 207
    }, {
      "utf8": " open",
      "tOffsetMs": 1981,
      "acAsrConf": 230
    }, {
      "utf8": " up",
      "tOffsetMs": 2310,
      "acAsrConf": 252
    }, {
      "utf8": " my",
      "tOffsetMs": 2430,
      "acAsrConf": 199
    }, {
      "utf8": " own",
      "tOffsetMs": 2491,
      "acAsrConf": 99
    } ]
  }, {
    "tStartMs": 321439,
    "dDurationMs": 3371,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 321449,
    "dDurationMs": 5640,
    "wWinId": 1,
    "segs": [ {
      "utf8": "coffee",
      "acAsrConf": 245
    }, {
      "utf8": " shop",
      "tOffsetMs": 541,
      "acAsrConf": 166
    }, {
      "utf8": " so",
      "tOffsetMs": 930,
      "acAsrConf": 244
    }, {
      "utf8": " this",
      "tOffsetMs": 1470,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 2310,
      "acAsrConf": 252
    }, {
      "utf8": " what",
      "tOffsetMs": 2490,
      "acAsrConf": 252
    }, {
      "utf8": " kind",
      "tOffsetMs": 2671,
      "acAsrConf": 220
    }, {
      "utf8": " of",
      "tOffsetMs": 3090,
      "acAsrConf": 204
    }, {
      "utf8": " what",
      "tOffsetMs": 3240,
      "acAsrConf": 206
    } ]
  }, {
    "tStartMs": 324800,
    "dDurationMs": 2289,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 324810,
    "dDurationMs": 4319,
    "wWinId": 1,
    "segs": [ {
      "utf8": "we",
      "acAsrConf": 252
    }, {
      "utf8": " do",
      "tOffsetMs": 150,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 300,
      "acAsrConf": 252
    }, {
      "utf8": " Asheville",
      "tOffsetMs": 449,
      "acAsrConf": 231
    }, {
      "utf8": " there's",
      "tOffsetMs": 900,
      "acAsrConf": 252
    }, {
      "utf8": " more",
      "tOffsetMs": 1169,
      "acAsrConf": 252
    }, {
      "utf8": " but",
      "tOffsetMs": 1380,
      "acAsrConf": 248
    }, {
      "utf8": " it's",
      "tOffsetMs": 1680,
      "acAsrConf": 255
    } ]
  }, {
    "tStartMs": 327079,
    "dDurationMs": 2050,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 327089,
    "dDurationMs": 4140,
    "wWinId": 1,
    "segs": [ {
      "utf8": "an",
      "acAsrConf": 252
    }, {
      "utf8": " interesting",
      "tOffsetMs": 211,
      "acAsrConf": 252
    }, {
      "utf8": " little",
      "tOffsetMs": 450,
      "acAsrConf": 145
    }, {
      "utf8": " town",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    }, {
      "utf8": " yeah",
      "tOffsetMs": 1260,
      "acAsrConf": 164
    }, {
      "utf8": " there's",
      "tOffsetMs": 1800,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 329119,
    "dDurationMs": 2110,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 329129,
    "dDurationMs": 4921,
    "wWinId": 1,
    "segs": [ {
      "utf8": "a",
      "acAsrConf": 252
    }, {
      "utf8": " lot",
      "tOffsetMs": 61,
      "acAsrConf": 245
    }, {
      "utf8": " to",
      "tOffsetMs": 271,
      "acAsrConf": 252
    }, {
      "utf8": " do",
      "tOffsetMs": 421,
      "acAsrConf": 252
    }, {
      "utf8": " here",
      "tOffsetMs": 540,
      "acAsrConf": 252
    }, {
      "utf8": " it's",
      "tOffsetMs": 570,
      "acAsrConf": 201
    }, {
      "utf8": " a",
      "tOffsetMs": 1081,
      "acAsrConf": 252
    }, {
      "utf8": " small",
      "tOffsetMs": 1171,
      "acAsrConf": 252
    }, {
      "utf8": " place",
      "tOffsetMs": 1320,
      "acAsrConf": 217
    }, {
      "utf8": " and",
      "tOffsetMs": 1650,
      "acAsrConf": 83
    } ]
  }, {
    "tStartMs": 331219,
    "dDurationMs": 2831,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 331229,
    "dDurationMs": 5370,
    "wWinId": 1,
    "segs": [ {
      "utf8": "not",
      "acAsrConf": 252
    }, {
      "utf8": " many",
      "tOffsetMs": 150,
      "acAsrConf": 252
    }, {
      "utf8": " foreign",
      "tOffsetMs": 451,
      "acAsrConf": 252
    }, {
      "utf8": " travelers",
      "tOffsetMs": 1171,
      "acAsrConf": 219
    }, {
      "utf8": " come",
      "tOffsetMs": 1921,
      "acAsrConf": 252
    }, {
      "utf8": " here",
      "tOffsetMs": 2250,
      "acAsrConf": 252
    }, {
      "utf8": " but",
      "tOffsetMs": 2581,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 334040,
    "dDurationMs": 2559,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 334050,
    "dDurationMs": 4890,
    "wWinId": 1,
    "segs": [ {
      "utf8": "if",
      "acAsrConf": 245
    }, {
      "utf8": " you",
      "tOffsetMs": 299,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 539,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " chance",
      "tOffsetMs": 690,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 1019,
      "acAsrConf": 238
    }, {
      "utf8": " recommend",
      "tOffsetMs": 1200,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 1679,
      "acAsrConf": 252
    }, {
      "utf8": " we",
      "tOffsetMs": 1859,
      "acAsrConf": 217
    } ]
  }, {
    "tStartMs": 336589,
    "dDurationMs": 2351,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 336599,
    "dDurationMs": 7891,
    "wWinId": 1,
    "segs": [ {
      "utf8": "will",
      "acAsrConf": 252
    }, {
      "utf8": " take",
      "tOffsetMs": 391,
      "acAsrConf": 235
    }, {
      "utf8": " you",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 961,
      "acAsrConf": 207
    }, {
      "utf8": " a",
      "tOffsetMs": 1171,
      "acAsrConf": 252
    }, {
      "utf8": " nice",
      "tOffsetMs": 1201,
      "acAsrConf": 227
    }, {
      "utf8": " restaurant",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    }, {
      "utf8": " how",
      "tOffsetMs": 1740,
      "acAsrConf": 77
    } ]
  }, {
    "tStartMs": 338930,
    "dDurationMs": 5560,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 338940,
    "dDurationMs": 9240,
    "wWinId": 1,
    "segs": [ {
      "utf8": "about",
      "acAsrConf": 219
    }, {
      "utf8": " that",
      "tOffsetMs": 300,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1159,
      "acAsrConf": 201
    }, {
      "utf8": " I",
      "tOffsetMs": 2159,
      "acAsrConf": 223
    }, {
      "utf8": " I",
      "tOffsetMs": 2550,
      "acAsrConf": 241
    }, {
      "utf8": " want",
      "tOffsetMs": 3649,
      "acAsrConf": 250
    }, {
      "utf8": " to",
      "tOffsetMs": 4649,
      "acAsrConf": 245
    }, {
      "utf8": " tell",
      "tOffsetMs": 4770,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 4979,
      "acAsrConf": 255
    }, {
      "utf8": " first",
      "tOffsetMs": 5129,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 344480,
    "dDurationMs": 3700,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 344490,
    "dDurationMs": 6329,
    "wWinId": 1,
    "segs": [ {
      "utf8": "some",
      "acAsrConf": 164
    }, {
      "utf8": " about",
      "tOffsetMs": 840,
      "acAsrConf": 252
    }, {
      "utf8": " my",
      "tOffsetMs": 1200,
      "acAsrConf": 245
    }, {
      "utf8": " personal",
      "tOffsetMs": 1649,
      "acAsrConf": 243
    }, {
      "utf8": " journey",
      "tOffsetMs": 2310,
      "acAsrConf": 168
    }, {
      "utf8": " having",
      "tOffsetMs": 2690,
      "acAsrConf": 228
    } ]
  }, {
    "tStartMs": 348170,
    "dDurationMs": 2649,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 348180,
    "dDurationMs": 6389,
    "wWinId": 1,
    "segs": [ {
      "utf8": "difficult",
      "acAsrConf": 233
    }, {
      "utf8": " frustrating",
      "tOffsetMs": 900,
      "acAsrConf": 252
    }, {
      "utf8": " boring",
      "tOffsetMs": 1260,
      "acAsrConf": 251
    }, {
      "utf8": " language",
      "tOffsetMs": 2159,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 350809,
    "dDurationMs": 3760,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 350819,
    "dDurationMs": 5491,
    "wWinId": 1,
    "segs": [ {
      "utf8": "classes",
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 511,
      "acAsrConf": 245
    }, {
      "utf8": " going",
      "tOffsetMs": 1460,
      "acAsrConf": 219
    }, {
      "utf8": " from",
      "tOffsetMs": 2460,
      "acAsrConf": 252
    }, {
      "utf8": " that",
      "tOffsetMs": 2731,
      "acAsrConf": 208
    }, {
      "utf8": " situation",
      "tOffsetMs": 2820,
      "acAsrConf": 146
    } ]
  }, {
    "tStartMs": 354559,
    "dDurationMs": 1751,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 354569,
    "dDurationMs": 4081,
    "wWinId": 1,
    "segs": [ {
      "utf8": "which",
      "acAsrConf": 252
    }, {
      "utf8": " many",
      "tOffsetMs": 270,
      "acAsrConf": 245
    }, {
      "utf8": " of",
      "tOffsetMs": 840,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1081,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 1291,
      "acAsrConf": 200
    }, {
      "utf8": " probably",
      "tOffsetMs": 1500,
      "acAsrConf": 217
    } ]
  }, {
    "tStartMs": 356300,
    "dDurationMs": 2350,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 356310,
    "dDurationMs": 6300,
    "wWinId": 1,
    "segs": [ {
      "utf8": "experienced",
      "acAsrConf": 250
    }, {
      "utf8": " with",
      "tOffsetMs": 840,
      "acAsrConf": 252
    }, {
      "utf8": " English",
      "tOffsetMs": 960,
      "acAsrConf": 236
    }, {
      "utf8": " boring",
      "tOffsetMs": 1379,
      "acAsrConf": 213
    }, {
      "utf8": " classes",
      "tOffsetMs": 1949,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 358640,
    "dDurationMs": 3970,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 358650,
    "dDurationMs": 6000,
    "wWinId": 1,
    "segs": [ {
      "utf8": "grammar",
      "acAsrConf": 209
    }, {
      "utf8": " classes",
      "tOffsetMs": 780,
      "acAsrConf": 167
    }, {
      "utf8": " to",
      "tOffsetMs": 1470,
      "acAsrConf": 216
    }, {
      "utf8": " being",
      "tOffsetMs": 2280,
      "acAsrConf": 241
    }, {
      "utf8": " able",
      "tOffsetMs": 2819,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 3120,
      "acAsrConf": 252
    }, {
      "utf8": " speak",
      "tOffsetMs": 3359,
      "acAsrConf": 236
    }, {
      "utf8": " a",
      "tOffsetMs": 3750,
      "acAsrConf": 233
    } ]
  }, {
    "tStartMs": 362600,
    "dDurationMs": 2050,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 362610,
    "dDurationMs": 3450,
    "wWinId": 1,
    "segs": [ {
      "utf8": "foreign",
      "acAsrConf": 252
    }, {
      "utf8": " language",
      "tOffsetMs": 239,
      "acAsrConf": 252
    }, {
      "utf8": " fluently",
      "tOffsetMs": 390,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1290,
      "acAsrConf": 165
    }, {
      "utf8": " the",
      "tOffsetMs": 1919,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 364640,
    "dDurationMs": 1420,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 364650,
    "dDurationMs": 4560,
    "wWinId": 1,
    "segs": [ {
      "utf8": "language",
      "acAsrConf": 191
    }, {
      "utf8": " that",
      "tOffsetMs": 389,
      "acAsrConf": 219
    }, {
      "utf8": " I",
      "tOffsetMs": 569,
      "acAsrConf": 255
    }, {
      "utf8": " learned",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 750,
      "acAsrConf": 157
    }, {
      "utf8": " speak",
      "tOffsetMs": 1199,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 366050,
    "dDurationMs": 3160,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 366060,
    "dDurationMs": 6529,
    "wWinId": 1,
    "segs": [ {
      "utf8": "fluently",
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 419,
      "acAsrConf": 211
    }, {
      "utf8": " French",
      "tOffsetMs": 840,
      "acAsrConf": 236
    }, {
      "utf8": " and",
      "tOffsetMs": 1349,
      "acAsrConf": 202
    }, {
      "utf8": " if",
      "tOffsetMs": 1710,
      "acAsrConf": 252
    }, {
      "utf8": " any",
      "tOffsetMs": 2579,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 2819,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 2909,
      "acAsrConf": 213
    }, {
      "utf8": " are",
      "tOffsetMs": 3120,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 369200,
    "dDurationMs": 3389,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 369210,
    "dDurationMs": 6540,
    "wWinId": 1,
    "segs": [ {
      "utf8": "French",
      "acAsrConf": 177
    }, {
      "utf8": " speakers",
      "tOffsetMs": 389,
      "acAsrConf": 252
    }, {
      "utf8": " out",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " there",
      "tOffsetMs": 900,
      "acAsrConf": 252
    }, {
      "utf8": " bond",
      "tOffsetMs": 929,
      "acAsrConf": 201
    }, {
      "utf8": " Rue",
      "tOffsetMs": 1440,
      "acAsrConf": 131
    }, {
      "utf8": " and",
      "tOffsetMs": 1679,
      "acAsrConf": 249
    } ]
  }, {
    "tStartMs": 372579,
    "dDurationMs": 3171,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 372589,
    "dDurationMs": 6190,
    "wWinId": 1,
    "segs": [ {
      "utf8": "when",
      "acAsrConf": 237
    }, {
      "utf8": " as",
      "tOffsetMs": 1000,
      "acAsrConf": 248
    }, {
      "utf8": " many",
      "tOffsetMs": 1241,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 1300,
      "acAsrConf": 164
    }, {
      "utf8": " you",
      "tOffsetMs": 1991,
      "acAsrConf": 244
    }, {
      "utf8": " saw",
      "tOffsetMs": 2110,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 2140,
      "acAsrConf": 152
    }, {
      "utf8": " my",
      "tOffsetMs": 2500,
      "acAsrConf": 252
    }, {
      "utf8": " free",
      "tOffsetMs": 2651,
      "acAsrConf": 252
    }, {
      "utf8": " ebook",
      "tOffsetMs": 2980,
      "acAsrConf": 246
    } ]
  }, {
    "tStartMs": 375740,
    "dDurationMs": 3039,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 375750,
    "dDurationMs": 5669,
    "wWinId": 1,
    "segs": [ {
      "utf8": "I",
      "acAsrConf": 216
    }, {
      "utf8": " talked",
      "tOffsetMs": 479,
      "acAsrConf": 252
    }, {
      "utf8": " about",
      "tOffsetMs": 1080,
      "acAsrConf": 167
    }, {
      "utf8": " I",
      "tOffsetMs": 1379,
      "acAsrConf": 244
    }, {
      "utf8": " used",
      "tOffsetMs": 1680,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 2219,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 2339,
      "acAsrConf": 252
    }, {
      "utf8": " Spanish",
      "tOffsetMs": 2520,
      "acAsrConf": 231
    } ]
  }, {
    "tStartMs": 378769,
    "dDurationMs": 2650,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 378779,
    "dDurationMs": 4831,
    "wWinId": 1,
    "segs": [ {
      "utf8": "classes",
      "acAsrConf": 239
    }, {
      "utf8": " where",
      "tOffsetMs": 540,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 841,
      "acAsrConf": 240
    }, {
      "utf8": " teacher",
      "tOffsetMs": 1681,
      "acAsrConf": 252
    }, {
      "utf8": " just",
      "tOffsetMs": 1950,
      "acAsrConf": 252
    }, {
      "utf8": " talked",
      "tOffsetMs": 2100,
      "acAsrConf": 226
    } ]
  }, {
    "tStartMs": 381409,
    "dDurationMs": 2201,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 381419,
    "dDurationMs": 4381,
    "wWinId": 1,
    "segs": [ {
      "utf8": "the",
      "acAsrConf": 252
    }, {
      "utf8": " whole",
      "tOffsetMs": 180,
      "acAsrConf": 148
    }, {
      "utf8": " time",
      "tOffsetMs": 300,
      "acAsrConf": 211
    }, {
      "utf8": " we",
      "tOffsetMs": 750,
      "acAsrConf": 236
    }, {
      "utf8": " just",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    }, {
      "utf8": " listened",
      "tOffsetMs": 1231,
      "acAsrConf": 237
    }, {
      "utf8": " and",
      "tOffsetMs": 1710,
      "acAsrConf": 246
    }, {
      "utf8": " read",
      "tOffsetMs": 1891,
      "acAsrConf": 210
    } ]
  }, {
    "tStartMs": 383600,
    "dDurationMs": 2200,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 383610,
    "dDurationMs": 4890,
    "wWinId": 1,
    "segs": [ {
      "utf8": "the",
      "acAsrConf": 252
    }, {
      "utf8": " textbook",
      "tOffsetMs": 149,
      "acAsrConf": 236
    }, {
      "utf8": " and",
      "tOffsetMs": 690,
      "acAsrConf": 252
    }, {
      "utf8": " that's",
      "tOffsetMs": 720,
      "acAsrConf": 202
    }, {
      "utf8": " boring",
      "tOffsetMs": 1559,
      "acAsrConf": 243
    }, {
      "utf8": " right",
      "tOffsetMs": 1950,
      "acAsrConf": 232
    } ]
  }, {
    "tStartMs": 385790,
    "dDurationMs": 2710,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 385800,
    "dDurationMs": 3740,
    "wWinId": 1,
    "segs": [ {
      "utf8": "yeah",
      "acAsrConf": 249
    }, {
      "utf8": " yeah",
      "tOffsetMs": 929,
      "acAsrConf": 219
    }, {
      "utf8": " you're",
      "tOffsetMs": 1260,
      "acAsrConf": 244
    }, {
      "utf8": " boring",
      "tOffsetMs": 1560,
      "acAsrConf": 224
    }, {
      "utf8": " just",
      "tOffsetMs": 1919,
      "acAsrConf": 202
    }, {
      "utf8": " thinking",
      "tOffsetMs": 2339,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 388490,
    "dDurationMs": 1050,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 388500,
    "dDurationMs": 4610,
    "wWinId": 1,
    "segs": [ {
      "utf8": "about",
      "acAsrConf": 146
    }, {
      "utf8": " it",
      "tOffsetMs": 270,
      "acAsrConf": 127
    }, {
      "utf8": " is",
      "tOffsetMs": 419,
      "acAsrConf": 203
    }, {
      "utf8": " more",
      "tOffsetMs": 539,
      "acAsrConf": 241
    } ]
  }, {
    "tStartMs": 389530,
    "dDurationMs": 3580,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 389540,
    "dDurationMs": 7350,
    "wWinId": 1,
    "segs": [ {
      "utf8": "yes",
      "acAsrConf": 252
    }, {
      "utf8": " yeah",
      "tOffsetMs": 390,
      "acAsrConf": 237
    }, {
      "utf8": " it's",
      "tOffsetMs": 1200,
      "acAsrConf": 252
    }, {
      "utf8": " not",
      "tOffsetMs": 1380,
      "acAsrConf": 252
    }, {
      "utf8": " very",
      "tOffsetMs": 1530,
      "acAsrConf": 170
    }, {
      "utf8": " easy",
      "tOffsetMs": 1970,
      "acAsrConf": 209
    }, {
      "utf8": " to",
      "tOffsetMs": 2970,
      "acAsrConf": 196
    }, {
      "utf8": " learn",
      "tOffsetMs": 3360,
      "acAsrConf": 211
    } ]
  }, {
    "tStartMs": 393100,
    "dDurationMs": 3790,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 393110,
    "dDurationMs": 6350,
    "wWinId": 1,
    "segs": [ {
      "utf8": "language",
      "acAsrConf": 236
    }, {
      "utf8": " for",
      "tOffsetMs": 690,
      "acAsrConf": 216
    }, {
      "utf8": " a",
      "tOffsetMs": 1470,
      "acAsrConf": 252
    }, {
      "utf8": " long",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    }, {
      "utf8": " time",
      "tOffsetMs": 1830,
      "acAsrConf": 223
    }, {
      "utf8": " in",
      "tOffsetMs": 2190,
      "acAsrConf": 252
    }, {
      "utf8": " school",
      "tOffsetMs": 2520,
      "acAsrConf": 252
    }, {
      "utf8": " yeah",
      "tOffsetMs": 2940,
      "acAsrConf": 207
    } ]
  }, {
    "tStartMs": 396880,
    "dDurationMs": 2580,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 396890,
    "dDurationMs": 5040,
    "wWinId": 1,
    "segs": [ {
      "utf8": "or",
      "acAsrConf": 252
    }, {
      "utf8": " be",
      "tOffsetMs": 180,
      "acAsrConf": 252
    }, {
      "utf8": " able",
      "tOffsetMs": 330,
      "acAsrConf": 211
    }, {
      "utf8": " to",
      "tOffsetMs": 570,
      "acAsrConf": 218
    }, {
      "utf8": " speak",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " fluently",
      "tOffsetMs": 1020,
      "acAsrConf": 252
    }, {
      "utf8": " from",
      "tOffsetMs": 1230,
      "acAsrConf": 142
    }, {
      "utf8": " school",
      "tOffsetMs": 2010,
      "acAsrConf": 245
    } ]
  }, {
    "tStartMs": 399450,
    "dDurationMs": 2480,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 399460,
    "dDurationMs": 5320,
    "wWinId": 1,
    "segs": [ {
      "utf8": "unfortunately",
      "acAsrConf": 205
    }, {
      "utf8": " yeah",
      "tOffsetMs": 1000,
      "acAsrConf": 162
    }, {
      "utf8": " and",
      "tOffsetMs": 1150,
      "acAsrConf": 252
    }, {
      "utf8": " I'm",
      "tOffsetMs": 1390,
      "acAsrConf": 243
    }, {
      "utf8": " sure",
      "tOffsetMs": 1780,
      "acAsrConf": 202
    }, {
      "utf8": " you",
      "tOffsetMs": 2080,
      "acAsrConf": 244
    }, {
      "utf8": " are",
      "tOffsetMs": 2440,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 401920,
    "dDurationMs": 2860,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 401930,
    "dDurationMs": 5640,
    "wWinId": 1,
    "segs": [ {
      "utf8": "very",
      "acAsrConf": 136
    }, {
      "utf8": " familiar",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " with",
      "tOffsetMs": 750,
      "acAsrConf": 216
    }, {
      "utf8": " that",
      "tOffsetMs": 1080,
      "acAsrConf": 252
    }, {
      "utf8": " situation",
      "tOffsetMs": 1230,
      "acAsrConf": 167
    }, {
      "utf8": " so",
      "tOffsetMs": 2010,
      "acAsrConf": 207
    }, {
      "utf8": " in",
      "tOffsetMs": 2550,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 404770,
    "dDurationMs": 2800,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 404780,
    "dDurationMs": 4560,
    "wWinId": 1,
    "segs": [ {
      "utf8": "my",
      "acAsrConf": 252
    }, {
      "utf8": " case",
      "tOffsetMs": 180,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 240,
      "acAsrConf": 206
    }, {
      "utf8": " didn't",
      "tOffsetMs": 890,
      "acAsrConf": 252
    }, {
      "utf8": " really",
      "tOffsetMs": 1890,
      "acAsrConf": 237
    }, {
      "utf8": " learn",
      "tOffsetMs": 2130,
      "acAsrConf": 252
    }, {
      "utf8": " Spanish",
      "tOffsetMs": 2280,
      "acAsrConf": 239
    } ]
  }, {
    "tStartMs": 407560,
    "dDurationMs": 1780,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 407570,
    "dDurationMs": 5670,
    "wWinId": 1,
    "segs": [ {
      "utf8": "very",
      "acAsrConf": 252
    }, {
      "utf8": " well",
      "tOffsetMs": 210,
      "acAsrConf": 228
    }, {
      "utf8": " because",
      "tOffsetMs": 300,
      "acAsrConf": 204
    }, {
      "utf8": " that",
      "tOffsetMs": 840,
      "acAsrConf": 252
    }, {
      "utf8": " situation",
      "tOffsetMs": 930,
      "acAsrConf": 234
    }, {
      "utf8": " was",
      "tOffsetMs": 1590,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 409330,
    "dDurationMs": 3910,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 409340,
    "dDurationMs": 6840,
    "wWinId": 1,
    "segs": [ {
      "utf8": "just",
      "acAsrConf": 252
    }, {
      "utf8": " boring",
      "tOffsetMs": 30,
      "acAsrConf": 219
    }, {
      "utf8": " and",
      "tOffsetMs": 690,
      "acAsrConf": 215
    }, {
      "utf8": " not",
      "tOffsetMs": 1370,
      "acAsrConf": 248
    }, {
      "utf8": " helpful",
      "tOffsetMs": 2370,
      "acAsrConf": 252
    }, {
      "utf8": " or",
      "tOffsetMs": 2820,
      "acAsrConf": 218
    }, {
      "utf8": " useful",
      "tOffsetMs": 3210,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 3450,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 413230,
    "dDurationMs": 2950,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 413240,
    "dDurationMs": 4650,
    "wWinId": 1,
    "segs": [ {
      "utf8": "was",
      "acAsrConf": 252
    }, {
      "utf8": " pretty",
      "tOffsetMs": 90,
      "acAsrConf": 218
    }, {
      "utf8": " stressful",
      "tOffsetMs": 450,
      "acAsrConf": 226
    }, {
      "utf8": " in",
      "tOffsetMs": 870,
      "acAsrConf": 230
    }, {
      "utf8": " tests",
      "tOffsetMs": 1350,
      "acAsrConf": 238
    }, {
      "utf8": " and",
      "tOffsetMs": 1980,
      "acAsrConf": 201
    }, {
      "utf8": " the",
      "tOffsetMs": 2160,
      "acAsrConf": 255
    } ]
  }, {
    "tStartMs": 416170,
    "dDurationMs": 1720,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 416180,
    "dDurationMs": 4500,
    "wWinId": 1,
    "segs": [ {
      "utf8": "teacher",
      "acAsrConf": 255
    }, {
      "utf8": " would",
      "tOffsetMs": 210,
      "acAsrConf": 252
    }, {
      "utf8": " always",
      "tOffsetMs": 450,
      "acAsrConf": 72
    }, {
      "utf8": " ask",
      "tOffsetMs": 690,
      "acAsrConf": 252
    }, {
      "utf8": " us",
      "tOffsetMs": 1020,
      "acAsrConf": 226
    }, {
      "utf8": " really",
      "tOffsetMs": 1320,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 417880,
    "dDurationMs": 2800,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 417890,
    "dDurationMs": 6000,
    "wWinId": 1,
    "segs": [ {
      "utf8": "difficult",
      "acAsrConf": 252
    }, {
      "utf8": " questions",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1290,
      "acAsrConf": 215
    }, {
      "utf8": " it",
      "tOffsetMs": 1560,
      "acAsrConf": 216
    }, {
      "utf8": " wasn't",
      "tOffsetMs": 2310,
      "acAsrConf": 169
    }, {
      "utf8": " a",
      "tOffsetMs": 2580,
      "acAsrConf": 252
    }, {
      "utf8": " good",
      "tOffsetMs": 2640,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 420670,
    "dDurationMs": 3220,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 420680,
    "dDurationMs": 4710,
    "wWinId": 1,
    "segs": [ {
      "utf8": "environment",
      "acAsrConf": 216
    }, {
      "utf8": " so",
      "tOffsetMs": 570,
      "acAsrConf": 219
    }, {
      "utf8": " with",
      "tOffsetMs": 900,
      "acAsrConf": 238
    }, {
      "utf8": " French",
      "tOffsetMs": 1380,
      "acAsrConf": 236
    }, {
      "utf8": " I",
      "tOffsetMs": 1980,
      "acAsrConf": 204
    }, {
      "utf8": " decided",
      "tOffsetMs": 2310,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 423880,
    "dDurationMs": 1510,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 423890,
    "dDurationMs": 5550,
    "wWinId": 1,
    "segs": [ {
      "utf8": "that",
      "acAsrConf": 243
    }, {
      "utf8": " I",
      "tOffsetMs": 150,
      "acAsrConf": 217
    }, {
      "utf8": " wanted",
      "tOffsetMs": 480,
      "acAsrConf": 210
    }, {
      "utf8": " to",
      "tOffsetMs": 810,
      "acAsrConf": 164
    }, {
      "utf8": " do",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 1110,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 1230,
      "acAsrConf": 252
    }, {
      "utf8": " little",
      "tOffsetMs": 1350,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 425380,
    "dDurationMs": 4060,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 425390,
    "dDurationMs": 10110,
    "wWinId": 1,
    "segs": [ {
      "utf8": "differently",
      "acAsrConf": 214
    }, {
      "utf8": " so",
      "tOffsetMs": 540,
      "acAsrConf": 153
    }, {
      "utf8": " I",
      "tOffsetMs": 1380,
      "acAsrConf": 249
    }, {
      "utf8": " went",
      "tOffsetMs": 1790,
      "acAsrConf": 232
    }, {
      "utf8": " to",
      "tOffsetMs": 2790,
      "acAsrConf": 200
    }, {
      "utf8": " the",
      "tOffsetMs": 3330,
      "acAsrConf": 252
    }, {
      "utf8": " French",
      "tOffsetMs": 3750,
      "acAsrConf": 249
    } ]
  }, {
    "tStartMs": 429430,
    "dDurationMs": 6070,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 429440,
    "dDurationMs": 8940,
    "wWinId": 1,
    "segs": [ {
      "utf8": "teachers",
      "acAsrConf": 227
    }, {
      "utf8": " office",
      "tOffsetMs": 1400,
      "acAsrConf": 217
    }, {
      "utf8": " and",
      "tOffsetMs": 2400,
      "acAsrConf": 167
    }, {
      "utf8": " I",
      "tOffsetMs": 3050,
      "acAsrConf": 246
    }, {
      "utf8": " asked",
      "tOffsetMs": 4880,
      "acAsrConf": 252
    }, {
      "utf8": " her",
      "tOffsetMs": 5880,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 435490,
    "dDurationMs": 2890,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 435500,
    "dDurationMs": 6180,
    "wWinId": 1,
    "segs": [ {
      "utf8": "questions",
      "acAsrConf": 252
    }, {
      "utf8": " just",
      "tOffsetMs": 690,
      "acAsrConf": 202
    }, {
      "utf8": " one",
      "tOffsetMs": 1350,
      "acAsrConf": 217
    }, {
      "utf8": " on",
      "tOffsetMs": 1650,
      "acAsrConf": 217
    }, {
      "utf8": " one",
      "tOffsetMs": 1800,
      "acAsrConf": 152
    }, {
      "utf8": " and",
      "tOffsetMs": 2070,
      "acAsrConf": 239
    }, {
      "utf8": " you",
      "tOffsetMs": 2280,
      "acAsrConf": 251
    }, {
      "utf8": " know",
      "tOffsetMs": 2760,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 438370,
    "dDurationMs": 3310,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 438380,
    "dDurationMs": 5880,
    "wWinId": 1,
    "segs": [ {
      "utf8": "what",
      "acAsrConf": 208
    }, {
      "utf8": " that",
      "tOffsetMs": 180,
      "acAsrConf": 252
    }, {
      "utf8": " was",
      "tOffsetMs": 360,
      "acAsrConf": 252
    }, {
      "utf8": " amazing",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " because",
      "tOffsetMs": 1050,
      "acAsrConf": 203
    }, {
      "utf8": " she",
      "tOffsetMs": 1490,
      "acAsrConf": 226
    }, {
      "utf8": " was",
      "tOffsetMs": 2490,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 3270,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 441670,
    "dDurationMs": 2590,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 441680,
    "dDurationMs": 4590,
    "wWinId": 1,
    "segs": [ {
      "utf8": "patient",
      "acAsrConf": 123
    }, {
      "utf8": " with",
      "tOffsetMs": 540,
      "acAsrConf": 227
    }, {
      "utf8": " me",
      "tOffsetMs": 930,
      "acAsrConf": 125
    }, {
      "utf8": " and",
      "tOffsetMs": 1170,
      "acAsrConf": 248
    }, {
      "utf8": " just",
      "tOffsetMs": 1530,
      "acAsrConf": 239
    }, {
      "utf8": " told",
      "tOffsetMs": 1830,
      "acAsrConf": 226
    }, {
      "utf8": " me",
      "tOffsetMs": 2070,
      "acAsrConf": 252
    }, {
      "utf8": " things",
      "tOffsetMs": 2280,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 444250,
    "dDurationMs": 2020,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 444260,
    "dDurationMs": 3960,
    "wWinId": 1,
    "segs": [ {
      "utf8": "like",
      "acAsrConf": 165
    }, {
      "utf8": " oh",
      "tOffsetMs": 270,
      "acAsrConf": 216
    }, {
      "utf8": " if",
      "tOffsetMs": 450,
      "acAsrConf": 248
    }, {
      "utf8": " you",
      "tOffsetMs": 960,
      "acAsrConf": 244
    }, {
      "utf8": " don't",
      "tOffsetMs": 1260,
      "acAsrConf": 227
    }, {
      "utf8": " understand",
      "tOffsetMs": 1470,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 446260,
    "dDurationMs": 1960,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 446270,
    "dDurationMs": 4830,
    "wWinId": 1,
    "segs": [ {
      "utf8": "everything",
      "acAsrConf": 252
    }, {
      "utf8": " eh",
      "tOffsetMs": 270,
      "acAsrConf": 137
    }, {
      "utf8": " it's",
      "tOffsetMs": 900,
      "acAsrConf": 250
    }, {
      "utf8": " no",
      "tOffsetMs": 1170,
      "acAsrConf": 252
    }, {
      "utf8": " problem",
      "tOffsetMs": 1350,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 448210,
    "dDurationMs": 2890,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 448220,
    "dDurationMs": 5910,
    "wWinId": 1,
    "segs": [ {
      "utf8": "let's",
      "acAsrConf": 252
    }, {
      "utf8": " just",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " talk",
      "tOffsetMs": 990,
      "acAsrConf": 252
    }, {
      "utf8": " together",
      "tOffsetMs": 1290,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1650,
      "acAsrConf": 181
    }, {
      "utf8": " try",
      "tOffsetMs": 2190,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 2460,
      "acAsrConf": 252
    }, {
      "utf8": " say",
      "tOffsetMs": 2490,
      "acAsrConf": 218
    } ]
  }, {
    "tStartMs": 451090,
    "dDurationMs": 3040,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 451100,
    "dDurationMs": 6240,
    "wWinId": 1,
    "segs": [ {
      "utf8": "anything",
      "acAsrConf": 202
    }, {
      "utf8": " and",
      "tOffsetMs": 540,
      "acAsrConf": 155
    }, {
      "utf8": " that",
      "tOffsetMs": 960,
      "acAsrConf": 240
    }, {
      "utf8": " was",
      "tOffsetMs": 1230,
      "acAsrConf": 248
    }, {
      "utf8": " exactly",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    }, {
      "utf8": " what",
      "tOffsetMs": 2040,
      "acAsrConf": 205
    }, {
      "utf8": " I",
      "tOffsetMs": 2820,
      "acAsrConf": 182
    } ]
  }, {
    "tStartMs": 454120,
    "dDurationMs": 3220,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 454130,
    "dDurationMs": 6480,
    "wWinId": 1,
    "segs": [ {
      "utf8": "needed",
      "acAsrConf": 121
    }, {
      "utf8": " so",
      "tOffsetMs": 240,
      "acAsrConf": 207
    }, {
      "utf8": " I",
      "tOffsetMs": 480,
      "acAsrConf": 152
    }, {
      "utf8": " had",
      "tOffsetMs": 1010,
      "acAsrConf": 228
    }, {
      "utf8": " some",
      "tOffsetMs": 2010,
      "acAsrConf": 242
    }, {
      "utf8": " classes",
      "tOffsetMs": 2310,
      "acAsrConf": 252
    }, {
      "utf8": " with",
      "tOffsetMs": 2580,
      "acAsrConf": 201
    }, {
      "utf8": " my",
      "tOffsetMs": 3000,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 457330,
    "dDurationMs": 3280,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 457340,
    "dDurationMs": 6960,
    "wWinId": 1,
    "segs": [ {
      "utf8": "French",
      "acAsrConf": 252
    }, {
      "utf8": " teacher",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 960,
      "acAsrConf": 201
    }, {
      "utf8": " then",
      "tOffsetMs": 1380,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 1890,
      "acAsrConf": 252
    }, {
      "utf8": " also",
      "tOffsetMs": 2010,
      "acAsrConf": 252
    }, {
      "utf8": " went",
      "tOffsetMs": 2310,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 2880,
      "acAsrConf": 206
    }, {
      "utf8": " a",
      "tOffsetMs": 3060,
      "acAsrConf": 212
    } ]
  }, {
    "tStartMs": 460600,
    "dDurationMs": 3700,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 460610,
    "dDurationMs": 5910,
    "wWinId": 1,
    "segs": [ {
      "utf8": "local",
      "acAsrConf": 252
    }, {
      "utf8": " French",
      "tOffsetMs": 510,
      "acAsrConf": 226
    }, {
      "utf8": " group",
      "tOffsetMs": 1290,
      "acAsrConf": 222
    }, {
      "utf8": " so",
      "tOffsetMs": 1860,
      "acAsrConf": 231
    }, {
      "utf8": " I",
      "tOffsetMs": 2460,
      "acAsrConf": 227
    }, {
      "utf8": " lived",
      "tOffsetMs": 2700,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 3180,
      "acAsrConf": 245
    }, {
      "utf8": " the",
      "tOffsetMs": 3300,
      "acAsrConf": 252
    }, {
      "utf8": " US",
      "tOffsetMs": 3420,
      "acAsrConf": 123
    } ]
  }, {
    "tStartMs": 464290,
    "dDurationMs": 2230,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 464300,
    "dDurationMs": 5460,
    "wWinId": 1,
    "segs": [ {
      "utf8": "but",
      "acAsrConf": 241
    }, {
      "utf8": " there",
      "tOffsetMs": 390,
      "acAsrConf": 228
    }, {
      "utf8": " were",
      "tOffsetMs": 630,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 900,
      "acAsrConf": 252
    }, {
      "utf8": " lot",
      "tOffsetMs": 930,
      "acAsrConf": 218
    }, {
      "utf8": " of",
      "tOffsetMs": 1170,
      "acAsrConf": 100
    }, {
      "utf8": " French",
      "tOffsetMs": 1380,
      "acAsrConf": 240
    }, {
      "utf8": " speakers",
      "tOffsetMs": 1770,
      "acAsrConf": 142
    } ]
  }, {
    "tStartMs": 466510,
    "dDurationMs": 3250,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 466520,
    "dDurationMs": 5970,
    "wWinId": 1,
    "segs": [ {
      "utf8": "in",
      "acAsrConf": 217
    }, {
      "utf8": " my",
      "tOffsetMs": 300,
      "acAsrConf": 206
    }, {
      "utf8": " city",
      "tOffsetMs": 660,
      "acAsrConf": 251
    }, {
      "utf8": " so",
      "tOffsetMs": 960,
      "acAsrConf": 208
    }, {
      "utf8": " we",
      "tOffsetMs": 1620,
      "acAsrConf": 252
    }, {
      "utf8": " met",
      "tOffsetMs": 2010,
      "acAsrConf": 252
    }, {
      "utf8": " up",
      "tOffsetMs": 2310,
      "acAsrConf": 201
    }, {
      "utf8": " together",
      "tOffsetMs": 2430,
      "acAsrConf": 252
    }, {
      "utf8": " each",
      "tOffsetMs": 2700,
      "acAsrConf": 146
    } ]
  }, {
    "tStartMs": 469750,
    "dDurationMs": 2740,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 469760,
    "dDurationMs": 6000,
    "wWinId": 1,
    "segs": [ {
      "utf8": "week",
      "acAsrConf": 252
    }, {
      "utf8": " at",
      "tOffsetMs": 480,
      "acAsrConf": 248
    }, {
      "utf8": " a",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " cafe",
      "tOffsetMs": 780,
      "acAsrConf": 246
    }, {
      "utf8": " and",
      "tOffsetMs": 1320,
      "acAsrConf": 166
    }, {
      "utf8": " that",
      "tOffsetMs": 1620,
      "acAsrConf": 252
    }, {
      "utf8": " was",
      "tOffsetMs": 2070,
      "acAsrConf": 252
    }, {
      "utf8": " great",
      "tOffsetMs": 2250,
      "acAsrConf": 235
    } ]
  }, {
    "tStartMs": 472480,
    "dDurationMs": 3280,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 472490,
    "dDurationMs": 6060,
    "wWinId": 1,
    "segs": [ {
      "utf8": "because",
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 330,
      "acAsrConf": 218
    }, {
      "utf8": " got",
      "tOffsetMs": 750,
      "acAsrConf": 226
    }, {
      "utf8": " to",
      "tOffsetMs": 1440,
      "acAsrConf": 217
    }, {
      "utf8": " hear",
      "tOffsetMs": 1650,
      "acAsrConf": 214
    }, {
      "utf8": " French",
      "tOffsetMs": 2100,
      "acAsrConf": 225
    }, {
      "utf8": " and",
      "tOffsetMs": 2610,
      "acAsrConf": 205
    }, {
      "utf8": " just",
      "tOffsetMs": 2880,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 475750,
    "dDurationMs": 2800,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 475760,
    "dDurationMs": 5190,
    "wWinId": 1,
    "segs": [ {
      "utf8": "try",
      "acAsrConf": 239
    }, {
      "utf8": " to",
      "tOffsetMs": 450,
      "acAsrConf": 208
    }, {
      "utf8": " speak",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1170,
      "acAsrConf": 249
    }, {
      "utf8": " I",
      "tOffsetMs": 1380,
      "acAsrConf": 244
    }, {
      "utf8": " bet",
      "tOffsetMs": 1680,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 2040,
      "acAsrConf": 252
    }, {
      "utf8": " lot",
      "tOffsetMs": 2160,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 2220,
      "acAsrConf": 144
    }, {
      "utf8": " you",
      "tOffsetMs": 2610,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 478540,
    "dDurationMs": 2410,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 478550,
    "dDurationMs": 6330,
    "wWinId": 1,
    "segs": [ {
      "utf8": "would",
      "acAsrConf": 252
    }, {
      "utf8": " love",
      "tOffsetMs": 240,
      "acAsrConf": 230
    }, {
      "utf8": " to",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " meet",
      "tOffsetMs": 960,
      "acAsrConf": 161
    }, {
      "utf8": " up",
      "tOffsetMs": 1530,
      "acAsrConf": 252
    }, {
      "utf8": " with",
      "tOffsetMs": 1560,
      "acAsrConf": 140
    }, {
      "utf8": " people",
      "tOffsetMs": 1740,
      "acAsrConf": 95
    }, {
      "utf8": " who",
      "tOffsetMs": 2370,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 480940,
    "dDurationMs": 3940,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 480950,
    "dDurationMs": 6270,
    "wWinId": 1,
    "segs": [ {
      "utf8": "aren't",
      "acAsrConf": 186
    }, {
      "utf8": " going",
      "tOffsetMs": 630,
      "acAsrConf": 239
    }, {
      "utf8": " to",
      "tOffsetMs": 1440,
      "acAsrConf": 201
    }, {
      "utf8": " pressure",
      "tOffsetMs": 1920,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 2910,
      "acAsrConf": 252
    }, {
      "utf8": " or",
      "tOffsetMs": 3300,
      "acAsrConf": 252
    }, {
      "utf8": " make",
      "tOffsetMs": 3330,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 3750,
      "acAsrConf": 124
    } ]
  }, {
    "tStartMs": 484870,
    "dDurationMs": 2350,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 484880,
    "dDurationMs": 4620,
    "wWinId": 1,
    "segs": [ {
      "utf8": "feel",
      "acAsrConf": 252
    }, {
      "utf8": " stressed",
      "tOffsetMs": 60,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 660,
      "acAsrConf": 201
    }, {
      "utf8": " we're",
      "tOffsetMs": 1200,
      "acAsrConf": 248
    }, {
      "utf8": " going",
      "tOffsetMs": 1770,
      "acAsrConf": 213
    }, {
      "utf8": " to",
      "tOffsetMs": 2040,
      "acAsrConf": 200
    }, {
      "utf8": " talk",
      "tOffsetMs": 2130,
      "acAsrConf": 218
    } ]
  }, {
    "tStartMs": 487210,
    "dDurationMs": 2290,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 487220,
    "dDurationMs": 4080,
    "wWinId": 1,
    "segs": [ {
      "utf8": "about",
      "acAsrConf": 252
    }, {
      "utf8": " how",
      "tOffsetMs": 30,
      "acAsrConf": 211
    }, {
      "utf8": " to",
      "tOffsetMs": 660,
      "acAsrConf": 216
    }, {
      "utf8": " do",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " that",
      "tOffsetMs": 1200,
      "acAsrConf": 231
    }, {
      "utf8": " a",
      "tOffsetMs": 1410,
      "acAsrConf": 234
    }, {
      "utf8": " little",
      "tOffsetMs": 1620,
      "acAsrConf": 212
    }, {
      "utf8": " bit",
      "tOffsetMs": 1920,
      "acAsrConf": 203
    }, {
      "utf8": " later",
      "tOffsetMs": 2070,
      "acAsrConf": 207
    } ]
  }, {
    "tStartMs": 489490,
    "dDurationMs": 1810,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 489500,
    "dDurationMs": 5340,
    "wWinId": 1,
    "segs": [ {
      "utf8": "but",
      "acAsrConf": 221
    }, {
      "utf8": " yes",
      "tOffsetMs": 420,
      "acAsrConf": 202
    }, {
      "utf8": " so",
      "tOffsetMs": 570,
      "acAsrConf": 203
    }, {
      "utf8": " meeting",
      "tOffsetMs": 780,
      "acAsrConf": 252
    }, {
      "utf8": " up",
      "tOffsetMs": 1020,
      "acAsrConf": 236
    }, {
      "utf8": " with",
      "tOffsetMs": 1260,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 1290,
      "acAsrConf": 204
    }, {
      "utf8": " teacher",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 491290,
    "dDurationMs": 3550,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 491300,
    "dDurationMs": 6630,
    "wWinId": 1,
    "segs": [ {
      "utf8": "helped",
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 480,
      "acAsrConf": 248
    }, {
      "utf8": " lot",
      "tOffsetMs": 570,
      "acAsrConf": 248
    }, {
      "utf8": " meeting",
      "tOffsetMs": 720,
      "acAsrConf": 212
    }, {
      "utf8": " up",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    }, {
      "utf8": " with",
      "tOffsetMs": 1680,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 1820,
      "acAsrConf": 237
    }, {
      "utf8": " local",
      "tOffsetMs": 2820,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 494830,
    "dDurationMs": 3100,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 494840,
    "dDurationMs": 6450,
    "wWinId": 1,
    "segs": [ {
      "utf8": "group",
      "acAsrConf": 245
    }, {
      "utf8": " helped",
      "tOffsetMs": 270,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " lot",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1050,
      "acAsrConf": 208
    }, {
      "utf8": " I",
      "tOffsetMs": 1410,
      "acAsrConf": 255
    }, {
      "utf8": " think",
      "tOffsetMs": 2310,
      "acAsrConf": 203
    }, {
      "utf8": " just",
      "tOffsetMs": 2880,
      "acAsrConf": 246
    } ]
  }, {
    "tStartMs": 497920,
    "dDurationMs": 3370,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 497930,
    "dDurationMs": 5430,
    "wWinId": 1,
    "segs": [ {
      "utf8": "having",
      "acAsrConf": 124
    }, {
      "utf8": " the",
      "tOffsetMs": 420,
      "acAsrConf": 252
    }, {
      "utf8": " bravery",
      "tOffsetMs": 1130,
      "acAsrConf": 204
    }, {
      "utf8": " the",
      "tOffsetMs": 2130,
      "acAsrConf": 216
    }, {
      "utf8": " courage",
      "tOffsetMs": 2490,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 2910,
      "acAsrConf": 252
    }, {
      "utf8": " be",
      "tOffsetMs": 2940,
      "acAsrConf": 176
    } ]
  }, {
    "tStartMs": 501280,
    "dDurationMs": 2080,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 501290,
    "dDurationMs": 3630,
    "wWinId": 1,
    "segs": [ {
      "utf8": "able",
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 180,
      "acAsrConf": 201
    }, {
      "utf8": " speak",
      "tOffsetMs": 360,
      "acAsrConf": 200
    }, {
      "utf8": " is",
      "tOffsetMs": 750,
      "acAsrConf": 216
    }, {
      "utf8": " a",
      "tOffsetMs": 990,
      "acAsrConf": 252
    }, {
      "utf8": " lot",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1380,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 1530,
      "acAsrConf": 234
    }, {
      "utf8": " know",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 1710,
      "acAsrConf": 252
    }, {
      "utf8": " lot",
      "tOffsetMs": 1830,
      "acAsrConf": 236
    } ]
  }, {
    "tStartMs": 503350,
    "dDurationMs": 1570,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 503360,
    "dDurationMs": 3750,
    "wWinId": 1,
    "segs": [ {
      "utf8": "of",
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 30,
      "acAsrConf": 238
    }, {
      "utf8": " you've",
      "tOffsetMs": 240,
      "acAsrConf": 200
    }, {
      "utf8": " asked",
      "tOffsetMs": 420,
      "acAsrConf": 252
    }, {
      "utf8": " great",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " questions",
      "tOffsetMs": 1020,
      "acAsrConf": 252
    }, {
      "utf8": " here",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 504910,
    "dDurationMs": 2200,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 504920,
    "dDurationMs": 4230,
    "wWinId": 1,
    "segs": [ {
      "utf8": "in",
      "acAsrConf": 175
    }, {
      "utf8": " the",
      "tOffsetMs": 390,
      "acAsrConf": 201
    }, {
      "utf8": " question",
      "tOffsetMs": 570,
      "acAsrConf": 138
    }, {
      "utf8": " and",
      "tOffsetMs": 960,
      "acAsrConf": 205
    }, {
      "utf8": " answer",
      "tOffsetMs": 1110,
      "acAsrConf": 205
    }, {
      "utf8": " about",
      "tOffsetMs": 1200,
      "acAsrConf": 252
    }, {
      "utf8": " how",
      "tOffsetMs": 1620,
      "acAsrConf": 196
    }, {
      "utf8": " can",
      "tOffsetMs": 1980,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 507100,
    "dDurationMs": 2050,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 507110,
    "dDurationMs": 3840,
    "wWinId": 1,
    "segs": [ {
      "utf8": "I",
      "acAsrConf": 241
    }, {
      "utf8": " speak",
      "tOffsetMs": 90,
      "acAsrConf": 248
    }, {
      "utf8": " how",
      "tOffsetMs": 390,
      "acAsrConf": 237
    }, {
      "utf8": " can",
      "tOffsetMs": 900,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 1110,
      "acAsrConf": 236
    }, {
      "utf8": " feel",
      "tOffsetMs": 1200,
      "acAsrConf": 252
    }, {
      "utf8": " confident",
      "tOffsetMs": 1380,
      "acAsrConf": 216
    }, {
      "utf8": " and",
      "tOffsetMs": 1950,
      "acAsrConf": 236
    } ]
  }, {
    "tStartMs": 509140,
    "dDurationMs": 1810,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 509150,
    "dDurationMs": 5190,
    "wWinId": 1,
    "segs": [ {
      "utf8": "that",
      "acAsrConf": 200
    }, {
      "utf8": " is",
      "tOffsetMs": 60,
      "acAsrConf": 209
    }, {
      "utf8": " exactly",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " what",
      "tOffsetMs": 810,
      "acAsrConf": 200
    }, {
      "utf8": " we're",
      "tOffsetMs": 1290,
      "acAsrConf": 252
    }, {
      "utf8": " going",
      "tOffsetMs": 1440,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1590,
      "acAsrConf": 213
    }, {
      "utf8": " talk",
      "tOffsetMs": 1650,
      "acAsrConf": 229
    } ]
  }, {
    "tStartMs": 510940,
    "dDurationMs": 3400,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 510950,
    "dDurationMs": 6690,
    "wWinId": 1,
    "segs": [ {
      "utf8": "about",
      "acAsrConf": 164
    }, {
      "utf8": " today",
      "tOffsetMs": 60,
      "acAsrConf": 201
    }, {
      "utf8": " so",
      "tOffsetMs": 510,
      "acAsrConf": 201
    }, {
      "utf8": " let's",
      "tOffsetMs": 870,
      "acAsrConf": 252
    }, {
      "utf8": " get",
      "tOffsetMs": 1320,
      "acAsrConf": 252
    }, {
      "utf8": " started",
      "tOffsetMs": 1590,
      "acAsrConf": 236
    }, {
      "utf8": " yes",
      "tOffsetMs": 2029,
      "acAsrConf": 218
    }, {
      "utf8": " yes",
      "tOffsetMs": 3029,
      "acAsrConf": 234
    } ]
  }, {
    "tStartMs": 514330,
    "dDurationMs": 3310,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 514340,
    "dDurationMs": 6090,
    "wWinId": 1,
    "segs": [ {
      "utf8": "so",
      "acAsrConf": 244
    }, {
      "utf8": " let's",
      "tOffsetMs": 450,
      "acAsrConf": 249
    }, {
      "utf8": " go",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 900,
      "acAsrConf": 252
    }, {
      "utf8": " section",
      "tOffsetMs": 1190,
      "acAsrConf": 216
    }, {
      "utf8": " number",
      "tOffsetMs": 2190,
      "acAsrConf": 238
    }, {
      "utf8": " one",
      "tOffsetMs": 2610,
      "acAsrConf": 198
    }, {
      "utf8": " if",
      "tOffsetMs": 2820,
      "acAsrConf": 242
    }, {
      "utf8": " you",
      "tOffsetMs": 3030,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 517630,
    "dDurationMs": 2800,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 517640,
    "dDurationMs": 6480,
    "wWinId": 1,
    "segs": [ {
      "utf8": "have",
      "acAsrConf": 200
    }, {
      "utf8": " the",
      "tOffsetMs": 270,
      "acAsrConf": 164
    }, {
      "utf8": " mind",
      "tOffsetMs": 600,
      "acAsrConf": 248
    }, {
      "utf8": " map",
      "tOffsetMs": 959,
      "acAsrConf": 248
    }, {
      "utf8": " if",
      "tOffsetMs": 1440,
      "acAsrConf": 242
    }, {
      "utf8": " you",
      "tOffsetMs": 1709,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 1800,
      "acAsrConf": 149
    }, {
      "utf8": " the",
      "tOffsetMs": 2220,
      "acAsrConf": 252
    }, {
      "utf8": " down",
      "tOffsetMs": 2430,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 520420,
    "dDurationMs": 3700,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 520430,
    "dDurationMs": 6780,
    "wWinId": 1,
    "segs": [ {
      "utf8": "load",
      "acAsrConf": 219
    }, {
      "utf8": " the",
      "tOffsetMs": 440,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 1440,
      "acAsrConf": 99
    }, {
      "utf8": " second",
      "tOffsetMs": 1830,
      "acAsrConf": 250
    }, {
      "utf8": " section",
      "tOffsetMs": 2219,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 2400,
      "acAsrConf": 217
    }, {
      "utf8": " how",
      "tOffsetMs": 2760,
      "acAsrConf": 220
    }, {
      "utf8": " to",
      "tOffsetMs": 3210,
      "acAsrConf": 255
    } ]
  }, {
    "tStartMs": 524110,
    "dDurationMs": 3100,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 524120,
    "dDurationMs": 6690,
    "wWinId": 1,
    "segs": [ {
      "utf8": "speak",
      "acAsrConf": 252
    }, {
      "utf8": " fluidly",
      "tOffsetMs": 330,
      "acAsrConf": 126
    }, {
      "utf8": " even",
      "tOffsetMs": 1020,
      "acAsrConf": 207
    }, {
      "utf8": " if",
      "tOffsetMs": 1800,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 2070,
      "acAsrConf": 252
    }, {
      "utf8": " don't",
      "tOffsetMs": 2250,
      "acAsrConf": 225
    }, {
      "utf8": " live",
      "tOffsetMs": 2490,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 2790,
      "acAsrConf": 167
    } ]
  }, {
    "tStartMs": 527200,
    "dDurationMs": 3610,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 527210,
    "dDurationMs": 7500,
    "wWinId": 1,
    "segs": [ {
      "utf8": "an",
      "acAsrConf": 252
    }, {
      "utf8": " english-speaking",
      "tOffsetMs": 150,
      "acAsrConf": 250
    }, {
      "utf8": " country",
      "tOffsetMs": 300,
      "acAsrConf": 142
    }, {
      "utf8": " so",
      "tOffsetMs": 1200,
      "acAsrConf": 207
    }, {
      "utf8": " then",
      "tOffsetMs": 2160,
      "acAsrConf": 236
    }, {
      "utf8": " what",
      "tOffsetMs": 2879,
      "acAsrConf": 239
    } ]
  }, {
    "tStartMs": 530800,
    "dDurationMs": 3910,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 530810,
    "dDurationMs": 6270,
    "wWinId": 1,
    "segs": [ {
      "utf8": "do",
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 270,
      "acAsrConf": 157
    }, {
      "utf8": " think",
      "tOffsetMs": 540,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 1170,
      "acAsrConf": 218
    }, {
      "utf8": " the",
      "tOffsetMs": 1469,
      "acAsrConf": 252
    }, {
      "utf8": " first",
      "tOffsetMs": 1890,
      "acAsrConf": 252
    }, {
      "utf8": " step",
      "tOffsetMs": 2339,
      "acAsrConf": 225
    }, {
      "utf8": " what",
      "tOffsetMs": 3029,
      "acAsrConf": 213
    }, {
      "utf8": " do",
      "tOffsetMs": 3779,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 534700,
    "dDurationMs": 2380,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 534710,
    "dDurationMs": 5220,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 252
    }, {
      "utf8": " think",
      "tOffsetMs": 60,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 390,
      "acAsrConf": 252
    }, {
      "utf8": " first",
      "tOffsetMs": 540,
      "acAsrConf": 252
    }, {
      "utf8": " step",
      "tOffsetMs": 780,
      "acAsrConf": 216
    }, {
      "utf8": " to",
      "tOffsetMs": 1080,
      "acAsrConf": 223
    }, {
      "utf8": " speaking",
      "tOffsetMs": 1740,
      "acAsrConf": 249
    } ]
  }, {
    "tStartMs": 537070,
    "dDurationMs": 2860,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 537080,
    "dDurationMs": 4850,
    "wWinId": 1,
    "segs": [ {
      "utf8": "fluently",
      "acAsrConf": 252
    }, {
      "utf8": " well",
      "tOffsetMs": 390,
      "acAsrConf": 210
    }, {
      "utf8": " it's",
      "tOffsetMs": 1050,
      "acAsrConf": 246
    }, {
      "utf8": " funny",
      "tOffsetMs": 1740,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 2100,
      "acAsrConf": 255
    }, {
      "utf8": " has",
      "tOffsetMs": 2340,
      "acAsrConf": 252
    }, {
      "utf8": " nothing",
      "tOffsetMs": 2460,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 539920,
    "dDurationMs": 2010,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 539930,
    "dDurationMs": 4469,
    "wWinId": 1,
    "segs": [ {
      "utf8": "to",
      "acAsrConf": 252
    }, {
      "utf8": " do",
      "tOffsetMs": 29,
      "acAsrConf": 252
    }, {
      "utf8": " with",
      "tOffsetMs": 210,
      "acAsrConf": 201
    }, {
      "utf8": " speaking",
      "tOffsetMs": 300,
      "acAsrConf": 252
    }, {
      "utf8": " it's",
      "tOffsetMs": 779,
      "acAsrConf": 232
    }, {
      "utf8": " actually",
      "tOffsetMs": 990,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 541920,
    "dDurationMs": 2479,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 541930,
    "dDurationMs": 5020,
    "wWinId": 1,
    "segs": [ {
      "utf8": "listening",
      "acAsrConf": 249
    }, {
      "utf8": " aha",
      "tOffsetMs": 1029,
      "acAsrConf": 153
    }, {
      "utf8": " you're",
      "tOffsetMs": 1360,
      "acAsrConf": 249
    }, {
      "utf8": " correct",
      "tOffsetMs": 1690,
      "acAsrConf": 252
    }, {
      "utf8": " listening",
      "tOffsetMs": 1990,
      "acAsrConf": 198
    } ]
  }, {
    "tStartMs": 544389,
    "dDurationMs": 2561,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 544399,
    "dDurationMs": 4801,
    "wWinId": 1,
    "segs": [ {
      "utf8": "regularly",
      "acAsrConf": 251
    }, {
      "utf8": " yes",
      "tOffsetMs": 781,
      "acAsrConf": 206
    }, {
      "utf8": " English",
      "tOffsetMs": 1081,
      "acAsrConf": 226
    }, {
      "utf8": " so",
      "tOffsetMs": 1380,
      "acAsrConf": 95
    }, {
      "utf8": " the",
      "tOffsetMs": 1861,
      "acAsrConf": 252
    }, {
      "utf8": " first",
      "tOffsetMs": 2071,
      "acAsrConf": 252
    }, {
      "utf8": " thing",
      "tOffsetMs": 2310,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 546940,
    "dDurationMs": 2260,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 546950,
    "dDurationMs": 5639,
    "wWinId": 1,
    "segs": [ {
      "utf8": "is",
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 300,
      "acAsrConf": 229
    }, {
      "utf8": " course",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " listening",
      "tOffsetMs": 810,
      "acAsrConf": 216
    }, {
      "utf8": " I",
      "tOffsetMs": 1680,
      "acAsrConf": 226
    }, {
      "utf8": " know",
      "tOffsetMs": 1800,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 1829,
      "acAsrConf": 206
    }, {
      "utf8": " talk",
      "tOffsetMs": 1949,
      "acAsrConf": 217
    } ]
  }, {
    "tStartMs": 549190,
    "dDurationMs": 3399,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 549200,
    "dDurationMs": 4949,
    "wWinId": 1,
    "segs": [ {
      "utf8": "about",
      "acAsrConf": 135
    }, {
      "utf8": " this",
      "tOffsetMs": 150,
      "acAsrConf": 227
    }, {
      "utf8": " a",
      "tOffsetMs": 449,
      "acAsrConf": 249
    }, {
      "utf8": " lot",
      "tOffsetMs": 630,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1050,
      "acAsrConf": 201
    }, {
      "utf8": " trying",
      "tOffsetMs": 1829,
      "acAsrConf": 244
    }, {
      "utf8": " to",
      "tOffsetMs": 2520,
      "acAsrConf": 207
    }, {
      "utf8": " listen",
      "tOffsetMs": 2970,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 3210,
      "acAsrConf": 237
    } ]
  }, {
    "tStartMs": 552579,
    "dDurationMs": 1570,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 552589,
    "dDurationMs": 4141,
    "wWinId": 1,
    "segs": [ {
      "utf8": "anything",
      "acAsrConf": 200
    }, {
      "utf8": " and",
      "tOffsetMs": 690,
      "acAsrConf": 205
    }, {
      "utf8": " do",
      "tOffsetMs": 1021,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1111,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 1201,
      "acAsrConf": 252
    }, {
      "utf8": " any",
      "tOffsetMs": 1381,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 554139,
    "dDurationMs": 2591,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 554149,
    "dDurationMs": 5401,
    "wWinId": 1,
    "segs": [ {
      "utf8": "recommendations",
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 841,
      "acAsrConf": 252
    }, {
      "utf8": " things",
      "tOffsetMs": 961,
      "acAsrConf": 200
    }, {
      "utf8": " that",
      "tOffsetMs": 1471,
      "acAsrConf": 252
    }, {
      "utf8": " learners",
      "tOffsetMs": 1831,
      "acAsrConf": 242
    } ]
  }, {
    "tStartMs": 556720,
    "dDurationMs": 2830,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 556730,
    "dDurationMs": 5910,
    "wWinId": 1,
    "segs": [ {
      "utf8": "can",
      "acAsrConf": 252
    }, {
      "utf8": " listen",
      "tOffsetMs": 30,
      "acAsrConf": 201
    }, {
      "utf8": " to",
      "tOffsetMs": 600,
      "acAsrConf": 242
    }, {
      "utf8": " well",
      "tOffsetMs": 810,
      "acAsrConf": 246
    }, {
      "utf8": " I",
      "tOffsetMs": 1590,
      "acAsrConf": 243
    }, {
      "utf8": " like",
      "tOffsetMs": 1859,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 2280,
      "acAsrConf": 188
    }, {
      "utf8": " listen",
      "tOffsetMs": 2490,
      "acAsrConf": 226
    }, {
      "utf8": " to",
      "tOffsetMs": 2669,
      "acAsrConf": 231
    } ]
  }, {
    "tStartMs": 559540,
    "dDurationMs": 3100,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 559550,
    "dDurationMs": 3780,
    "wWinId": 1,
    "segs": [ {
      "utf8": "podcasts",
      "acAsrConf": 200
    }, {
      "utf8": " which",
      "tOffsetMs": 710,
      "acAsrConf": 147
    }, {
      "utf8": " are",
      "tOffsetMs": 1710,
      "acAsrConf": 242
    }, {
      "utf8": " radio",
      "tOffsetMs": 1890,
      "acAsrConf": 252
    }, {
      "utf8": " shows",
      "tOffsetMs": 2310,
      "acAsrConf": 252
    }, {
      "utf8": " on",
      "tOffsetMs": 2670,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 2700,
      "acAsrConf": 184
    } ]
  }, {
    "tStartMs": 562630,
    "dDurationMs": 700,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 562640,
    "dDurationMs": 4560,
    "wWinId": 1,
    "segs": [ {
      "utf8": "Internet",
      "acAsrConf": 238
    } ]
  }, {
    "tStartMs": 563320,
    "dDurationMs": 3880,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 563330,
    "dDurationMs": 6569,
    "wWinId": 1,
    "segs": [ {
      "utf8": "so",
      "acAsrConf": 252
    }, {
      "utf8": " just",
      "tOffsetMs": 240,
      "acAsrConf": 236
    }, {
      "utf8": " normal",
      "tOffsetMs": 1080,
      "acAsrConf": 249
    }, {
      "utf8": " people",
      "tOffsetMs": 1860,
      "acAsrConf": 252
    }, {
      "utf8": " talking",
      "tOffsetMs": 2090,
      "acAsrConf": 206
    }, {
      "utf8": " which",
      "tOffsetMs": 3090,
      "acAsrConf": 215
    }, {
      "utf8": " is",
      "tOffsetMs": 3720,
      "acAsrConf": 207
    } ]
  }, {
    "tStartMs": 567190,
    "dDurationMs": 2709,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 567200,
    "dDurationMs": 5730,
    "wWinId": 1,
    "segs": [ {
      "utf8": "perfect",
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 570,
      "acAsrConf": 252
    }, {
      "utf8": " listen",
      "tOffsetMs": 630,
      "acAsrConf": 231
    }, {
      "utf8": " to",
      "tOffsetMs": 1020,
      "acAsrConf": 252
    }, {
      "utf8": " on",
      "tOffsetMs": 1199,
      "acAsrConf": 246
    }, {
      "utf8": " a",
      "tOffsetMs": 1590,
      "acAsrConf": 252
    }, {
      "utf8": " regular",
      "tOffsetMs": 1829,
      "acAsrConf": 217
    }, {
      "utf8": " basis",
      "tOffsetMs": 2250,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 569889,
    "dDurationMs": 3041,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 569899,
    "dDurationMs": 7861,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 187
    }, {
      "utf8": " can",
      "tOffsetMs": 481,
      "acAsrConf": 252
    }, {
      "utf8": " also",
      "tOffsetMs": 661,
      "acAsrConf": 252
    }, {
      "utf8": " watch",
      "tOffsetMs": 810,
      "acAsrConf": 252
    }, {
      "utf8": " TV",
      "tOffsetMs": 1201,
      "acAsrConf": 252
    }, {
      "utf8": " shows",
      "tOffsetMs": 1921,
      "acAsrConf": 252
    }, {
      "utf8": " their",
      "tOffsetMs": 1951,
      "acAsrConf": 87
    }, {
      "utf8": " movies",
      "tOffsetMs": 2611,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 572920,
    "dDurationMs": 4840,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 572930,
    "dDurationMs": 7800,
    "wWinId": 1,
    "segs": [ {
      "utf8": "obviously",
      "acAsrConf": 246
    }, {
      "utf8": " but",
      "tOffsetMs": 590,
      "acAsrConf": 201
    }, {
      "utf8": " sometimes",
      "tOffsetMs": 1730,
      "acAsrConf": 249
    }, {
      "utf8": " the",
      "tOffsetMs": 2730,
      "acAsrConf": 252
    }, {
      "utf8": " dialogue",
      "tOffsetMs": 3180,
      "acAsrConf": 234
    }, {
      "utf8": " the",
      "tOffsetMs": 4080,
      "acAsrConf": 167
    } ]
  }, {
    "tStartMs": 577750,
    "dDurationMs": 2980,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 577760,
    "dDurationMs": 6329,
    "wWinId": 1,
    "segs": [ {
      "utf8": "conversation",
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 750,
      "acAsrConf": 237
    }, {
      "utf8": " not",
      "tOffsetMs": 870,
      "acAsrConf": 200
    }, {
      "utf8": " very",
      "tOffsetMs": 1199,
      "acAsrConf": 234
    }, {
      "utf8": " real",
      "tOffsetMs": 1740,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 2100,
      "acAsrConf": 241
    }, {
      "utf8": " is",
      "tOffsetMs": 2490,
      "acAsrConf": 219
    } ]
  }, {
    "tStartMs": 580720,
    "dDurationMs": 3369,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 580730,
    "dDurationMs": 6780,
    "wWinId": 1,
    "segs": [ {
      "utf8": "funny",
      "acAsrConf": 239
    }, {
      "utf8": " and",
      "tOffsetMs": 479,
      "acAsrConf": 252
    }, {
      "utf8": " interesting",
      "tOffsetMs": 810,
      "acAsrConf": 236
    }, {
      "utf8": " but",
      "tOffsetMs": 1710,
      "acAsrConf": 252
    }, {
      "utf8": " not",
      "tOffsetMs": 2280,
      "acAsrConf": 236
    }, {
      "utf8": " not",
      "tOffsetMs": 2940,
      "acAsrConf": 252
    }, {
      "utf8": " always",
      "tOffsetMs": 3150,
      "acAsrConf": 218
    } ]
  }, {
    "tStartMs": 584079,
    "dDurationMs": 3431,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 584089,
    "dDurationMs": 6061,
    "wWinId": 1,
    "segs": [ {
      "utf8": "the",
      "acAsrConf": 200
    }, {
      "utf8": " most",
      "tOffsetMs": 301,
      "acAsrConf": 252
    }, {
      "utf8": " real",
      "tOffsetMs": 331,
      "acAsrConf": 207
    }, {
      "utf8": " version",
      "tOffsetMs": 1051,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 1891,
      "acAsrConf": 217
    }, {
      "utf8": " conversation",
      "tOffsetMs": 2421,
      "acAsrConf": 241
    } ]
  }, {
    "tStartMs": 587500,
    "dDurationMs": 2650,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 587510,
    "dDurationMs": 4320,
    "wWinId": 1,
    "segs": [ {
      "utf8": "but",
      "acAsrConf": 236
    }, {
      "utf8": " still",
      "tOffsetMs": 210,
      "acAsrConf": 234
    }, {
      "utf8": " very",
      "tOffsetMs": 840,
      "acAsrConf": 122
    }, {
      "utf8": " good",
      "tOffsetMs": 1110,
      "acAsrConf": 207
    }, {
      "utf8": " sure",
      "tOffsetMs": 1410,
      "acAsrConf": 249
    }, {
      "utf8": " sure",
      "tOffsetMs": 1920,
      "acAsrConf": 219
    }, {
      "utf8": " that's",
      "tOffsetMs": 2400,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 590140,
    "dDurationMs": 1690,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 590150,
    "dDurationMs": 3629,
    "wWinId": 1,
    "segs": [ {
      "utf8": "very",
      "acAsrConf": 179
    }, {
      "utf8": " useful",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " especially",
      "tOffsetMs": 450,
      "acAsrConf": 179
    }, {
      "utf8": " since",
      "tOffsetMs": 780,
      "acAsrConf": 166
    }, {
      "utf8": " you",
      "tOffsetMs": 1439,
      "acAsrConf": 252
    }, {
      "utf8": " don't",
      "tOffsetMs": 1559,
      "acAsrConf": 202
    } ]
  }, {
    "tStartMs": 591820,
    "dDurationMs": 1959,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 591830,
    "dDurationMs": 4020,
    "wWinId": 1,
    "segs": [ {
      "utf8": "live",
      "acAsrConf": 204
    }, {
      "utf8": " in",
      "tOffsetMs": 210,
      "acAsrConf": 252
    }, {
      "utf8": " an",
      "tOffsetMs": 330,
      "acAsrConf": 252
    }, {
      "utf8": " english-speaking",
      "tOffsetMs": 420,
      "acAsrConf": 255
    }, {
      "utf8": " country",
      "tOffsetMs": 540,
      "acAsrConf": 207
    }, {
      "utf8": " so",
      "tOffsetMs": 1050,
      "acAsrConf": 149
    } ]
  }, {
    "tStartMs": 593769,
    "dDurationMs": 2081,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 593779,
    "dDurationMs": 3661,
    "wWinId": 1,
    "segs": [ {
      "utf8": "there's",
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 391,
      "acAsrConf": 252
    }, {
      "utf8": " special",
      "tOffsetMs": 451,
      "acAsrConf": 236
    }, {
      "utf8": " rule",
      "tOffsetMs": 810,
      "acAsrConf": 201
    }, {
      "utf8": " that",
      "tOffsetMs": 1351,
      "acAsrConf": 143
    }, {
      "utf8": " I",
      "tOffsetMs": 1591,
      "acAsrConf": 232
    }, {
      "utf8": " want",
      "tOffsetMs": 1651,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1951,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 595840,
    "dDurationMs": 1600,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 595850,
    "dDurationMs": 4609,
    "wWinId": 1,
    "segs": [ {
      "utf8": "tell",
      "acAsrConf": 251
    }, {
      "utf8": " you",
      "tOffsetMs": 179,
      "acAsrConf": 251
    }, {
      "utf8": " about",
      "tOffsetMs": 359,
      "acAsrConf": 252
    }, {
      "utf8": " though",
      "tOffsetMs": 510,
      "acAsrConf": 202
    }, {
      "utf8": " when",
      "tOffsetMs": 1140,
      "acAsrConf": 252
    }, {
      "utf8": " you're",
      "tOffsetMs": 1410,
      "acAsrConf": 237
    } ]
  }, {
    "tStartMs": 597430,
    "dDurationMs": 3029,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 597440,
    "dDurationMs": 8190,
    "wWinId": 1,
    "segs": [ {
      "utf8": "listening",
      "acAsrConf": 236
    }, {
      "utf8": " to",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " TV",
      "tOffsetMs": 480,
      "acAsrConf": 204
    }, {
      "utf8": " shows",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " movies",
      "tOffsetMs": 990,
      "acAsrConf": 197
    }, {
      "utf8": " podcasts",
      "tOffsetMs": 1820,
      "acAsrConf": 249
    }, {
      "utf8": " I",
      "tOffsetMs": 2820,
      "acAsrConf": 206
    } ]
  }, {
    "tStartMs": 600449,
    "dDurationMs": 5181,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 600459,
    "dDurationMs": 11081,
    "wWinId": 1,
    "segs": [ {
      "utf8": "like",
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1000,
      "acAsrConf": 246
    }, {
      "utf8": " call",
      "tOffsetMs": 1151,
      "acAsrConf": 209
    }, {
      "utf8": " it",
      "tOffsetMs": 1361,
      "acAsrConf": 196
    }, {
      "utf8": " the",
      "tOffsetMs": 1481,
      "acAsrConf": 252
    }, {
      "utf8": " 70%",
      "tOffsetMs": 2281,
      "acAsrConf": 201
    }, {
      "utf8": " rule",
      "tOffsetMs": 3281,
      "acAsrConf": 219
    }, {
      "utf8": " so",
      "tOffsetMs": 3761,
      "acAsrConf": 234
    }, {
      "utf8": " the",
      "tOffsetMs": 4331,
      "acAsrConf": 252
    }, {
      "utf8": " 70%",
      "tOffsetMs": 4481,
      "acAsrConf": 248
    } ]
  }, {
    "tStartMs": 605620,
    "dDurationMs": 5920,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 605630,
    "dDurationMs": 10110,
    "wWinId": 1,
    "segs": [ {
      "utf8": "rule",
      "acAsrConf": 201
    }, {
      "utf8": " it",
      "tOffsetMs": 990,
      "acAsrConf": 233
    }, {
      "utf8": " means",
      "tOffsetMs": 1829,
      "acAsrConf": 252
    }, {
      "utf8": " if",
      "tOffsetMs": 2280,
      "acAsrConf": 212
    }, {
      "utf8": " you",
      "tOffsetMs": 2640,
      "acAsrConf": 252
    }, {
      "utf8": " can",
      "tOffsetMs": 3050,
      "acAsrConf": 144
    }, {
      "utf8": " understand",
      "tOffsetMs": 4050,
      "acAsrConf": 252
    }, {
      "utf8": " more",
      "tOffsetMs": 4950,
      "acAsrConf": 229
    } ]
  }, {
    "tStartMs": 611530,
    "dDurationMs": 4210,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 611540,
    "dDurationMs": 7020,
    "wWinId": 1,
    "segs": [ {
      "utf8": "than",
      "acAsrConf": 222
    }, {
      "utf8": " 70%",
      "tOffsetMs": 510,
      "acAsrConf": 226
    }, {
      "utf8": " don't",
      "tOffsetMs": 1500,
      "acAsrConf": 215
    }, {
      "utf8": " use",
      "tOffsetMs": 2400,
      "acAsrConf": 252
    }, {
      "utf8": " subtitles",
      "tOffsetMs": 2820,
      "acAsrConf": 252
    }, {
      "utf8": " don't",
      "tOffsetMs": 3330,
      "acAsrConf": 131
    }, {
      "utf8": " use",
      "tOffsetMs": 3930,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 615730,
    "dDurationMs": 2830,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 615740,
    "dDurationMs": 5550,
    "wWinId": 1,
    "segs": [ {
      "utf8": "your",
      "acAsrConf": 200
    }, {
      "utf8": " native",
      "tOffsetMs": 390,
      "acAsrConf": 252
    }, {
      "utf8": " language",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " subtitles",
      "tOffsetMs": 1200,
      "acAsrConf": 252
    }, {
      "utf8": " don't",
      "tOffsetMs": 1980,
      "acAsrConf": 217
    }, {
      "utf8": " use",
      "tOffsetMs": 2520,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 618550,
    "dDurationMs": 2740,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 618560,
    "dDurationMs": 6390,
    "wWinId": 1,
    "segs": [ {
      "utf8": "English",
      "acAsrConf": 227
    }, {
      "utf8": " subtitles",
      "tOffsetMs": 480,
      "acAsrConf": 237
    }, {
      "utf8": " try",
      "tOffsetMs": 1020,
      "acAsrConf": 216
    }, {
      "utf8": " to",
      "tOffsetMs": 1620,
      "acAsrConf": 252
    }, {
      "utf8": " test",
      "tOffsetMs": 2130,
      "acAsrConf": 252
    }, {
      "utf8": " your",
      "tOffsetMs": 2520,
      "acAsrConf": 224
    } ]
  }, {
    "tStartMs": 621280,
    "dDurationMs": 3670,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 621290,
    "dDurationMs": 6270,
    "wWinId": 1,
    "segs": [ {
      "utf8": "listening",
      "acAsrConf": 201
    }, {
      "utf8": " by",
      "tOffsetMs": 660,
      "acAsrConf": 242
    }, {
      "utf8": " using",
      "tOffsetMs": 1470,
      "acAsrConf": 255
    }, {
      "utf8": " no",
      "tOffsetMs": 2160,
      "acAsrConf": 217
    }, {
      "utf8": " subtitles",
      "tOffsetMs": 2760,
      "acAsrConf": 221
    }, {
      "utf8": " so",
      "tOffsetMs": 3360,
      "acAsrConf": 217
    }, {
      "utf8": " if",
      "tOffsetMs": 3480,
      "acAsrConf": 246
    } ]
  }, {
    "tStartMs": 624940,
    "dDurationMs": 2620,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 624950,
    "dDurationMs": 4620,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 252
    }, {
      "utf8": " understand",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " if",
      "tOffsetMs": 840,
      "acAsrConf": 201
    }, {
      "utf8": " you're",
      "tOffsetMs": 1110,
      "acAsrConf": 228
    }, {
      "utf8": " watching",
      "tOffsetMs": 1590,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 1800,
      "acAsrConf": 227
    }, {
      "utf8": " TV",
      "tOffsetMs": 2280,
      "acAsrConf": 188
    } ]
  }, {
    "tStartMs": 627550,
    "dDurationMs": 2020,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 627560,
    "dDurationMs": 3930,
    "wWinId": 1,
    "segs": [ {
      "utf8": "show",
      "acAsrConf": 203
    }, {
      "utf8": " friends",
      "tOffsetMs": 180,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 540,
      "acAsrConf": 146
    }, {
      "utf8": " lot",
      "tOffsetMs": 899,
      "acAsrConf": 215
    }, {
      "utf8": " of",
      "tOffsetMs": 1170,
      "acAsrConf": 233
    }, {
      "utf8": " you",
      "tOffsetMs": 1290,
      "acAsrConf": 252
    }, {
      "utf8": " love",
      "tOffsetMs": 1410,
      "acAsrConf": 219
    }, {
      "utf8": " to",
      "tOffsetMs": 1649,
      "acAsrConf": 252
    }, {
      "utf8": " listen",
      "tOffsetMs": 1680,
      "acAsrConf": 167
    } ]
  }, {
    "tStartMs": 629560,
    "dDurationMs": 1930,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 629570,
    "dDurationMs": 6170,
    "wWinId": 1,
    "segs": [ {
      "utf8": "to",
      "acAsrConf": 239
    }, {
      "utf8": " watch",
      "tOffsetMs": 90,
      "acAsrConf": 232
    }, {
      "utf8": " friends",
      "tOffsetMs": 420,
      "acAsrConf": 255
    }, {
      "utf8": " so",
      "tOffsetMs": 750,
      "acAsrConf": 200
    }, {
      "utf8": " if",
      "tOffsetMs": 1350,
      "acAsrConf": 252
    }, {
      "utf8": " you're",
      "tOffsetMs": 1560,
      "acAsrConf": 239
    }, {
      "utf8": " watching",
      "tOffsetMs": 1740,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 631480,
    "dDurationMs": 4260,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 631490,
    "dDurationMs": 7980,
    "wWinId": 1,
    "segs": [ {
      "utf8": "friends",
      "acAsrConf": 215
    }, {
      "utf8": " and",
      "tOffsetMs": 690,
      "acAsrConf": 234
    }, {
      "utf8": " you",
      "tOffsetMs": 839,
      "acAsrConf": 252
    }, {
      "utf8": " understand",
      "tOffsetMs": 1080,
      "acAsrConf": 252
    }, {
      "utf8": " about",
      "tOffsetMs": 1710,
      "acAsrConf": 227
    }, {
      "utf8": " 70%",
      "tOffsetMs": 2990,
      "acAsrConf": 206
    } ]
  }, {
    "tStartMs": 635730,
    "dDurationMs": 3740,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 635740,
    "dDurationMs": 7180,
    "wWinId": 1,
    "segs": [ {
      "utf8": "turn",
      "acAsrConf": 204
    }, {
      "utf8": " off",
      "tOffsetMs": 1000,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 1270,
      "acAsrConf": 252
    }, {
      "utf8": " subtitles",
      "tOffsetMs": 1630,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1839,
      "acAsrConf": 218
    }, {
      "utf8": " try",
      "tOffsetMs": 2530,
      "acAsrConf": 249
    }, {
      "utf8": " to",
      "tOffsetMs": 3250,
      "acAsrConf": 252
    }, {
      "utf8": " watch",
      "tOffsetMs": 3310,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 639460,
    "dDurationMs": 3460,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 639470,
    "dDurationMs": 7859,
    "wWinId": 1,
    "segs": [ {
      "utf8": "it",
      "acAsrConf": 252
    }, {
      "utf8": " just",
      "tOffsetMs": 270,
      "acAsrConf": 217
    }, {
      "utf8": " regularly",
      "tOffsetMs": 1400,
      "acAsrConf": 228
    }, {
      "utf8": " with",
      "tOffsetMs": 2400,
      "acAsrConf": 204
    }, {
      "utf8": " no",
      "tOffsetMs": 2609,
      "acAsrConf": 252
    }, {
      "utf8": " subtitles",
      "tOffsetMs": 2820,
      "acAsrConf": 197
    } ]
  }, {
    "tStartMs": 642910,
    "dDurationMs": 4419,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 642920,
    "dDurationMs": 6870,
    "wWinId": 1,
    "segs": [ {
      "utf8": "that's",
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 150,
      "acAsrConf": 252
    }, {
      "utf8": " 70%",
      "tOffsetMs": 359,
      "acAsrConf": 205
    }, {
      "utf8": " rule",
      "tOffsetMs": 900,
      "acAsrConf": 169
    }, {
      "utf8": " so",
      "tOffsetMs": 1530,
      "acAsrConf": 234
    }, {
      "utf8": " under",
      "tOffsetMs": 2130,
      "acAsrConf": 249
    }, {
      "utf8": " 70%",
      "tOffsetMs": 2820,
      "acAsrConf": 227
    }, {
      "utf8": " it's",
      "tOffsetMs": 3409,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 647319,
    "dDurationMs": 2471,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 647329,
    "dDurationMs": 5071,
    "wWinId": 1,
    "segs": [ {
      "utf8": "okay",
      "acAsrConf": 200
    }, {
      "utf8": " it's",
      "tOffsetMs": 630,
      "acAsrConf": 247
    }, {
      "utf8": " no",
      "tOffsetMs": 901,
      "acAsrConf": 252
    }, {
      "utf8": " problem",
      "tOffsetMs": 1260,
      "acAsrConf": 249
    }, {
      "utf8": " to",
      "tOffsetMs": 1561,
      "acAsrConf": 207
    }, {
      "utf8": " use",
      "tOffsetMs": 1921,
      "acAsrConf": 205
    }, {
      "utf8": " English",
      "tOffsetMs": 2221,
      "acAsrConf": 243
    } ]
  }, {
    "tStartMs": 649780,
    "dDurationMs": 2620,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 649790,
    "dDurationMs": 3660,
    "wWinId": 1,
    "segs": [ {
      "utf8": "subtitles",
      "acAsrConf": 215
    }, {
      "utf8": " and",
      "tOffsetMs": 840,
      "acAsrConf": 233
    }, {
      "utf8": " if",
      "tOffsetMs": 989,
      "acAsrConf": 219
    }, {
      "utf8": " you",
      "tOffsetMs": 1799,
      "acAsrConf": 226
    }, {
      "utf8": " don't",
      "tOffsetMs": 2190,
      "acAsrConf": 234
    }, {
      "utf8": " have",
      "tOffsetMs": 2430,
      "acAsrConf": 251
    } ]
  }, {
    "tStartMs": 652390,
    "dDurationMs": 1060,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 652400,
    "dDurationMs": 2880,
    "wWinId": 1,
    "segs": [ {
      "utf8": "subtitles",
      "acAsrConf": 202
    }, {
      "utf8": " it",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " would",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " be",
      "tOffsetMs": 840,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 653440,
    "dDurationMs": 1840,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 653450,
    "dDurationMs": 4950,
    "wWinId": 1,
    "segs": [ {
      "utf8": "boring",
      "acAsrConf": 252
    }, {
      "utf8": " because",
      "tOffsetMs": 390,
      "acAsrConf": 215
    }, {
      "utf8": " you",
      "tOffsetMs": 630,
      "acAsrConf": 171
    }, {
      "utf8": " wouldn't",
      "tOffsetMs": 930,
      "acAsrConf": 239
    }, {
      "utf8": " understand",
      "tOffsetMs": 1290,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 655270,
    "dDurationMs": 3130,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 655280,
    "dDurationMs": 5580,
    "wWinId": 1,
    "segs": [ {
      "utf8": "what's",
      "acAsrConf": 252
    }, {
      "utf8": " happening",
      "tOffsetMs": 210,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " under",
      "tOffsetMs": 870,
      "acAsrConf": 242
    }, {
      "utf8": " 70%",
      "tOffsetMs": 1530,
      "acAsrConf": 218
    }, {
      "utf8": " don't",
      "tOffsetMs": 1980,
      "acAsrConf": 217
    }, {
      "utf8": " feel",
      "tOffsetMs": 2820,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 658390,
    "dDurationMs": 2470,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 658400,
    "dDurationMs": 4410,
    "wWinId": 1,
    "segs": [ {
      "utf8": "bad",
      "acAsrConf": 252
    }, {
      "utf8": " about",
      "tOffsetMs": 240,
      "acAsrConf": 215
    }, {
      "utf8": " using",
      "tOffsetMs": 540,
      "acAsrConf": 217
    }, {
      "utf8": " subtitles",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    }, {
      "utf8": " some",
      "tOffsetMs": 1290,
      "acAsrConf": 200
    }, {
      "utf8": " of",
      "tOffsetMs": 2190,
      "acAsrConf": 252
    }, {
      "utf8": " my",
      "tOffsetMs": 2340,
      "acAsrConf": 233
    } ]
  }, {
    "tStartMs": 660850,
    "dDurationMs": 1960,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 660860,
    "dDurationMs": 5190,
    "wWinId": 1,
    "segs": [ {
      "utf8": "students",
      "acAsrConf": 244
    }, {
      "utf8": " say",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " oh",
      "tOffsetMs": 660,
      "acAsrConf": 232
    }, {
      "utf8": " I'm",
      "tOffsetMs": 930,
      "acAsrConf": 239
    }, {
      "utf8": " so",
      "tOffsetMs": 1290,
      "acAsrConf": 252
    }, {
      "utf8": " sorry",
      "tOffsetMs": 1560,
      "acAsrConf": 243
    } ]
  }, {
    "tStartMs": 662800,
    "dDurationMs": 3250,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 662810,
    "dDurationMs": 4170,
    "wWinId": 1,
    "segs": [ {
      "utf8": "I",
      "acAsrConf": 255
    }, {
      "utf8": " used",
      "tOffsetMs": 120,
      "acAsrConf": 252
    }, {
      "utf8": " English",
      "tOffsetMs": 720,
      "acAsrConf": 207
    }, {
      "utf8": " subtitles",
      "tOffsetMs": 1140,
      "acAsrConf": 252
    }, {
      "utf8": " but",
      "tOffsetMs": 1800,
      "acAsrConf": 207
    }, {
      "utf8": " it's",
      "tOffsetMs": 2160,
      "acAsrConf": 245
    }, {
      "utf8": " not",
      "tOffsetMs": 2910,
      "acAsrConf": 246
    }, {
      "utf8": " a",
      "tOffsetMs": 3210,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 666040,
    "dDurationMs": 940,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 666050,
    "dDurationMs": 2880,
    "wWinId": 1,
    "segs": [ {
      "utf8": "big",
      "acAsrConf": 252
    }, {
      "utf8": " deal",
      "tOffsetMs": 270,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 666970,
    "dDurationMs": 1960,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 666980,
    "dDurationMs": 6000,
    "wWinId": 1,
    "segs": [ {
      "utf8": "that's",
      "acAsrConf": 233
    }, {
      "utf8": " okay",
      "tOffsetMs": 180,
      "acAsrConf": 207
    }, {
      "utf8": " yes",
      "tOffsetMs": 540,
      "acAsrConf": 225
    }, {
      "utf8": " if",
      "tOffsetMs": 810,
      "acAsrConf": 144
    }, {
      "utf8": " you",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    }, {
      "utf8": " understand",
      "tOffsetMs": 1260,
      "acAsrConf": 252
    }, {
      "utf8": " less",
      "tOffsetMs": 1770,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 668920,
    "dDurationMs": 4060,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 668930,
    "dDurationMs": 5070,
    "wWinId": 1,
    "segs": [ {
      "utf8": "than",
      "acAsrConf": 226
    }, {
      "utf8": " 70%",
      "tOffsetMs": 270,
      "acAsrConf": 216
    }, {
      "utf8": " go",
      "tOffsetMs": 720,
      "acAsrConf": 160
    }, {
      "utf8": " ahead",
      "tOffsetMs": 1710,
      "acAsrConf": 252
    }, {
      "utf8": " use",
      "tOffsetMs": 2040,
      "acAsrConf": 214
    }, {
      "utf8": " subtitles",
      "tOffsetMs": 2580,
      "acAsrConf": 240
    }, {
      "utf8": " as",
      "tOffsetMs": 3270,
      "acAsrConf": 202
    }, {
      "utf8": " much",
      "tOffsetMs": 3600,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 672970,
    "dDurationMs": 1030,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 672980,
    "dDurationMs": 4260,
    "wWinId": 1,
    "segs": [ {
      "utf8": "as",
      "acAsrConf": 252
    }, {
      "utf8": " possible",
      "tOffsetMs": 270,
      "acAsrConf": 144
    } ]
  }, {
    "tStartMs": 673990,
    "dDurationMs": 3250,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 674000,
    "dDurationMs": 6090,
    "wWinId": 1,
    "segs": [ {
      "utf8": "that's",
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 630,
      "acAsrConf": 252
    }, {
      "utf8": " 70%",
      "tOffsetMs": 810,
      "acAsrConf": 79
    }, {
      "utf8": " rule",
      "tOffsetMs": 1440,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1620,
      "acAsrConf": 208
    }, {
      "utf8": " what",
      "tOffsetMs": 2100,
      "acAsrConf": 252
    }, {
      "utf8": " about",
      "tOffsetMs": 2400,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 2520,
      "acAsrConf": 250
    } ]
  }, {
    "tStartMs": 677230,
    "dDurationMs": 2860,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 677240,
    "dDurationMs": 4200,
    "wWinId": 1,
    "segs": [ {
      "utf8": "second",
      "acAsrConf": 249
    }, {
      "utf8": " tip",
      "tOffsetMs": 810,
      "acAsrConf": 218
    }, {
      "utf8": " for",
      "tOffsetMs": 1170,
      "acAsrConf": 252
    }, {
      "utf8": " how",
      "tOffsetMs": 1380,
      "acAsrConf": 80
    }, {
      "utf8": " to",
      "tOffsetMs": 1650,
      "acAsrConf": 200
    }, {
      "utf8": " speak",
      "tOffsetMs": 1710,
      "acAsrConf": 252
    }, {
      "utf8": " fluently",
      "tOffsetMs": 2160,
      "acAsrConf": 252
    }, {
      "utf8": " if",
      "tOffsetMs": 2550,
      "acAsrConf": 88
    } ]
  }, {
    "tStartMs": 680080,
    "dDurationMs": 1360,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 680090,
    "dDurationMs": 4320,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 252
    }, {
      "utf8": " don't",
      "tOffsetMs": 390,
      "acAsrConf": 252
    }, {
      "utf8": " live",
      "tOffsetMs": 690,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 870,
      "acAsrConf": 200
    }, {
      "utf8": " an",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " english-speaking",
      "tOffsetMs": 1050,
      "acAsrConf": 146
    } ]
  }, {
    "tStartMs": 681430,
    "dDurationMs": 2980,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 681440,
    "dDurationMs": 5220,
    "wWinId": 1,
    "segs": [ {
      "utf8": "country",
      "acAsrConf": 252
    }, {
      "utf8": " what",
      "tOffsetMs": 620,
      "acAsrConf": 186
    }, {
      "utf8": " do",
      "tOffsetMs": 1620,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1950,
      "acAsrConf": 252
    }, {
      "utf8": " think",
      "tOffsetMs": 2070,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 2340,
      "acAsrConf": 252
    }, {
      "utf8": " on",
      "tOffsetMs": 2460,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 2610,
      "acAsrConf": 252
    }, {
      "utf8": " mind",
      "tOffsetMs": 2760,
      "acAsrConf": 227
    } ]
  }, {
    "tStartMs": 684400,
    "dDurationMs": 2260,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 684410,
    "dDurationMs": 6210,
    "wWinId": 1,
    "segs": [ {
      "utf8": "map",
      "acAsrConf": 227
    }, {
      "utf8": " on",
      "tOffsetMs": 120,
      "acAsrConf": 147
    }, {
      "utf8": " the",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " document",
      "tOffsetMs": 930,
      "acAsrConf": 247
    }, {
      "utf8": " I",
      "tOffsetMs": 1440,
      "acAsrConf": 168
    }, {
      "utf8": " gave",
      "tOffsetMs": 1530,
      "acAsrConf": 245
    }, {
      "utf8": " to",
      "tOffsetMs": 1740,
      "acAsrConf": 201
    }, {
      "utf8": " you",
      "tOffsetMs": 1920,
      "acAsrConf": 235
    }, {
      "utf8": " it",
      "tOffsetMs": 2070,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 686650,
    "dDurationMs": 3970,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 686660,
    "dDurationMs": 5700,
    "wWinId": 1,
    "segs": [ {
      "utf8": "says",
      "acAsrConf": 236
    }, {
      "utf8": " hmm",
      "tOffsetMs": 450,
      "acAsrConf": 128
    }, {
      "utf8": " out",
      "tOffsetMs": 1440,
      "acAsrConf": 230
    }, {
      "utf8": " loud",
      "tOffsetMs": 1650,
      "acAsrConf": 230
    }, {
      "utf8": " is",
      "tOffsetMs": 2370,
      "acAsrConf": 226
    }, {
      "utf8": " the",
      "tOffsetMs": 2850,
      "acAsrConf": 167
    }, {
      "utf8": " second",
      "tOffsetMs": 3150,
      "acAsrConf": 252
    }, {
      "utf8": " step",
      "tOffsetMs": 3510,
      "acAsrConf": 251
    } ]
  }, {
    "tStartMs": 690610,
    "dDurationMs": 1750,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 690620,
    "dDurationMs": 3450,
    "wWinId": 1,
    "segs": [ {
      "utf8": "what",
      "acAsrConf": 236
    }, {
      "utf8": " do",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 750,
      "acAsrConf": 250
    }, {
      "utf8": " think",
      "tOffsetMs": 840,
      "acAsrConf": 232
    }, {
      "utf8": " what",
      "tOffsetMs": 1170,
      "acAsrConf": 220
    }, {
      "utf8": " do",
      "tOffsetMs": 1530,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1620,
      "acAsrConf": 252
    }, {
      "utf8": " think",
      "tOffsetMs": 1650,
      "acAsrConf": 207
    } ]
  }, {
    "tStartMs": 692350,
    "dDurationMs": 1720,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 692360,
    "dDurationMs": 5940,
    "wWinId": 1,
    "segs": [ {
      "utf8": "should",
      "acAsrConf": 186
    }, {
      "utf8": " be",
      "tOffsetMs": 300,
      "acAsrConf": 215
    }, {
      "utf8": " there",
      "tOffsetMs": 360,
      "acAsrConf": 236
    }, {
      "utf8": " I",
      "tOffsetMs": 570,
      "acAsrConf": 142
    }, {
      "utf8": " believe",
      "tOffsetMs": 810,
      "acAsrConf": 224
    }, {
      "utf8": " it",
      "tOffsetMs": 1230,
      "acAsrConf": 252
    }, {
      "utf8": " should",
      "tOffsetMs": 1470,
      "acAsrConf": 252
    }, {
      "utf8": " be",
      "tOffsetMs": 1500,
      "acAsrConf": 206
    } ]
  }, {
    "tStartMs": 694060,
    "dDurationMs": 4240,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 694070,
    "dDurationMs": 6990,
    "wWinId": 1,
    "segs": [ {
      "utf8": "reading",
      "acAsrConf": 226
    }, {
      "utf8": " out",
      "tOffsetMs": 990,
      "acAsrConf": 233
    }, {
      "utf8": " oh",
      "tOffsetMs": 1320,
      "acAsrConf": 213
    }, {
      "utf8": " good",
      "tOffsetMs": 1650,
      "acAsrConf": 217
    }, {
      "utf8": " reading",
      "tOffsetMs": 2430,
      "acAsrConf": 247
    }, {
      "utf8": " in",
      "tOffsetMs": 3360,
      "acAsrConf": 252
    }, {
      "utf8": " your",
      "tOffsetMs": 3510,
      "acAsrConf": 142
    }, {
      "utf8": " head",
      "tOffsetMs": 3690,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 698290,
    "dDurationMs": 2770,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 698300,
    "dDurationMs": 5820,
    "wWinId": 1,
    "segs": [ {
      "utf8": "uh-huh",
      "acAsrConf": 239
    }, {
      "utf8": " I",
      "tOffsetMs": 600,
      "acAsrConf": 213
    }, {
      "utf8": " have",
      "tOffsetMs": 1110,
      "acAsrConf": 239
    }, {
      "utf8": " to",
      "tOffsetMs": 1230,
      "acAsrConf": 243
    }, {
      "utf8": " read",
      "tOffsetMs": 1410,
      "acAsrConf": 252
    }, {
      "utf8": " out",
      "tOffsetMs": 1680,
      "acAsrConf": 252
    }, {
      "utf8": " loud",
      "tOffsetMs": 1890,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1920,
      "acAsrConf": 252
    }, {
      "utf8": " hear",
      "tOffsetMs": 2160,
      "acAsrConf": 202
    } ]
  }, {
    "tStartMs": 701050,
    "dDurationMs": 3070,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 701060,
    "dDurationMs": 6930,
    "wWinId": 1,
    "segs": [ {
      "utf8": "your",
      "acAsrConf": 220
    }, {
      "utf8": " voice",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 570,
      "acAsrConf": 203
    }, {
      "utf8": " practice",
      "tOffsetMs": 1440,
      "acAsrConf": 249
    }, {
      "utf8": " pronouncing",
      "tOffsetMs": 1920,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 2370,
      "acAsrConf": 236
    } ]
  }, {
    "tStartMs": 704110,
    "dDurationMs": 3880,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 704120,
    "dDurationMs": 5760,
    "wWinId": 1,
    "segs": [ {
      "utf8": "pronunciation",
      "acAsrConf": 235
    }, {
      "utf8": " yes",
      "tOffsetMs": 270,
      "acAsrConf": 62
    }, {
      "utf8": " yes",
      "tOffsetMs": 1170,
      "acAsrConf": 210
    }, {
      "utf8": " so",
      "tOffsetMs": 1590,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 2090,
      "acAsrConf": 248
    }, {
      "utf8": " the",
      "tOffsetMs": 3090,
      "acAsrConf": 250
    }, {
      "utf8": " second",
      "tOffsetMs": 3510,
      "acAsrConf": 250
    } ]
  }, {
    "tStartMs": 707980,
    "dDurationMs": 1900,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 707990,
    "dDurationMs": 3960,
    "wWinId": 1,
    "segs": [ {
      "utf8": "step",
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 180,
      "acAsrConf": 227
    }, {
      "utf8": " speaking",
      "tOffsetMs": 570,
      "acAsrConf": 247
    }, {
      "utf8": " fluently",
      "tOffsetMs": 990,
      "acAsrConf": 252
    }, {
      "utf8": " first",
      "tOffsetMs": 1380,
      "acAsrConf": 255
    }, {
      "utf8": " is",
      "tOffsetMs": 1680,
      "acAsrConf": 227
    } ]
  }, {
    "tStartMs": 709870,
    "dDurationMs": 2080,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 709880,
    "dDurationMs": 3810,
    "wWinId": 1,
    "segs": [ {
      "utf8": "listening",
      "acAsrConf": 252
    }, {
      "utf8": " just",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " training",
      "tOffsetMs": 780,
      "acAsrConf": 223
    }, {
      "utf8": " your",
      "tOffsetMs": 1140,
      "acAsrConf": 216
    }, {
      "utf8": " ear",
      "tOffsetMs": 1410,
      "acAsrConf": 239
    }, {
      "utf8": " the",
      "tOffsetMs": 1620,
      "acAsrConf": 231
    } ]
  }, {
    "tStartMs": 711940,
    "dDurationMs": 1750,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 711950,
    "dDurationMs": 4740,
    "wWinId": 1,
    "segs": [ {
      "utf8": "second",
      "acAsrConf": 249
    }, {
      "utf8": " one",
      "tOffsetMs": 330,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 450,
      "acAsrConf": 252
    }, {
      "utf8": " training",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " your",
      "tOffsetMs": 1080,
      "acAsrConf": 201
    }, {
      "utf8": " mouth",
      "tOffsetMs": 1380,
      "acAsrConf": 224
    } ]
  }, {
    "tStartMs": 713680,
    "dDurationMs": 3010,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 713690,
    "dDurationMs": 5370,
    "wWinId": 1,
    "segs": [ {
      "utf8": "muscles",
      "acAsrConf": 207
    }, {
      "utf8": " by",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " reading",
      "tOffsetMs": 990,
      "acAsrConf": 252
    }, {
      "utf8": " something",
      "tOffsetMs": 1620,
      "acAsrConf": 243
    }, {
      "utf8": " out",
      "tOffsetMs": 2220,
      "acAsrConf": 201
    }, {
      "utf8": " loud",
      "tOffsetMs": 2640,
      "acAsrConf": 224
    } ]
  }, {
    "tStartMs": 716680,
    "dDurationMs": 2380,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 716690,
    "dDurationMs": 4710,
    "wWinId": 1,
    "segs": [ {
      "utf8": "and",
      "acAsrConf": 227
    }, {
      "utf8": " of",
      "tOffsetMs": 210,
      "acAsrConf": 217
    }, {
      "utf8": " course",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " if",
      "tOffsetMs": 870,
      "acAsrConf": 251
    }, {
      "utf8": " you",
      "tOffsetMs": 1140,
      "acAsrConf": 252
    }, {
      "utf8": " want",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1740,
      "acAsrConf": 204
    }, {
      "utf8": " be",
      "tOffsetMs": 1860,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 1950,
      "acAsrConf": 252
    }, {
      "utf8": " fluent",
      "tOffsetMs": 1980,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 719050,
    "dDurationMs": 2350,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 719060,
    "dDurationMs": 5970,
    "wWinId": 1,
    "segs": [ {
      "utf8": "speaker",
      "acAsrConf": 252
    }, {
      "utf8": " try",
      "tOffsetMs": 60,
      "acAsrConf": 193
    }, {
      "utf8": " to",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " read",
      "tOffsetMs": 1020,
      "acAsrConf": 252
    }, {
      "utf8": " something",
      "tOffsetMs": 1350,
      "acAsrConf": 252
    }, {
      "utf8": " natural",
      "tOffsetMs": 1620,
      "acAsrConf": 226
    } ]
  }, {
    "tStartMs": 721390,
    "dDurationMs": 3640,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 721400,
    "dDurationMs": 6480,
    "wWinId": 1,
    "segs": [ {
      "utf8": "don't",
      "acAsrConf": 208
    }, {
      "utf8": " read",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " like",
      "tOffsetMs": 900,
      "acAsrConf": 249
    }, {
      "utf8": " CNN",
      "tOffsetMs": 1580,
      "acAsrConf": 216
    }, {
      "utf8": " it's",
      "tOffsetMs": 2580,
      "acAsrConf": 232
    }, {
      "utf8": " not",
      "tOffsetMs": 2940,
      "acAsrConf": 252
    }, {
      "utf8": " very",
      "tOffsetMs": 3300,
      "acAsrConf": 249
    } ]
  }, {
    "tStartMs": 725020,
    "dDurationMs": 2860,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 725030,
    "dDurationMs": 5490,
    "wWinId": 1,
    "segs": [ {
      "utf8": "natural",
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 210,
      "acAsrConf": 203
    }, {
      "utf8": " conversations",
      "tOffsetMs": 630,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1350,
      "acAsrConf": 207
    }, {
      "utf8": " try",
      "tOffsetMs": 2010,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 2430,
      "acAsrConf": 252
    }, {
      "utf8": " read",
      "tOffsetMs": 2460,
      "acAsrConf": 238
    } ]
  }, {
    "tStartMs": 727870,
    "dDurationMs": 2650,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 727880,
    "dDurationMs": 5400,
    "wWinId": 1,
    "segs": [ {
      "utf8": "a",
      "acAsrConf": 223
    }, {
      "utf8": " dialogue",
      "tOffsetMs": 180,
      "acAsrConf": 236
    }, {
      "utf8": " try",
      "tOffsetMs": 810,
      "acAsrConf": 227
    }, {
      "utf8": " to",
      "tOffsetMs": 1170,
      "acAsrConf": 252
    }, {
      "utf8": " read",
      "tOffsetMs": 1230,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 1470,
      "acAsrConf": 238
    }, {
      "utf8": " blog",
      "tOffsetMs": 1590,
      "acAsrConf": 208
    }, {
      "utf8": " post",
      "tOffsetMs": 1920,
      "acAsrConf": 252
    }, {
      "utf8": " read",
      "tOffsetMs": 2310,
      "acAsrConf": 205
    } ]
  }, {
    "tStartMs": 730510,
    "dDurationMs": 2770,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 730520,
    "dDurationMs": 6630,
    "wWinId": 1,
    "segs": [ {
      "utf8": "something",
      "acAsrConf": 236
    }, {
      "utf8": " that's",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 660,
      "acAsrConf": 207
    }, {
      "utf8": " a",
      "tOffsetMs": 900,
      "acAsrConf": 252
    }, {
      "utf8": " conversational",
      "tOffsetMs": 1760,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 733270,
    "dDurationMs": 3880,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 733280,
    "dDurationMs": 6930,
    "wWinId": 1,
    "segs": [ {
      "utf8": "tone",
      "acAsrConf": 136
    }, {
      "utf8": " so",
      "tOffsetMs": 240,
      "acAsrConf": 243
    }, {
      "utf8": " read",
      "tOffsetMs": 930,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 1170,
      "acAsrConf": 252
    }, {
      "utf8": " out",
      "tOffsetMs": 1200,
      "acAsrConf": 203
    }, {
      "utf8": " loud",
      "tOffsetMs": 1740,
      "acAsrConf": 244
    }, {
      "utf8": " a",
      "tOffsetMs": 1800,
      "acAsrConf": 128
    }, {
      "utf8": " short",
      "tOffsetMs": 2580,
      "acAsrConf": 252
    }, {
      "utf8": " story",
      "tOffsetMs": 3120,
      "acAsrConf": 211
    }, {
      "utf8": " a",
      "tOffsetMs": 3450,
      "acAsrConf": 206
    } ]
  }, {
    "tStartMs": 737140,
    "dDurationMs": 3070,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 737150,
    "dDurationMs": 5280,
    "wWinId": 1,
    "segs": [ {
      "utf8": "few",
      "acAsrConf": 222
    }, {
      "utf8": " paragraphs",
      "tOffsetMs": 270,
      "acAsrConf": 252
    }, {
      "utf8": " an",
      "tOffsetMs": 660,
      "acAsrConf": 0
    }, {
      "utf8": " article",
      "tOffsetMs": 1290,
      "acAsrConf": 236
    }, {
      "utf8": " one",
      "tOffsetMs": 1950,
      "acAsrConf": 225
    }, {
      "utf8": " of",
      "tOffsetMs": 2490,
      "acAsrConf": 252
    }, {
      "utf8": " my",
      "tOffsetMs": 2520,
      "acAsrConf": 201
    }, {
      "utf8": " blog",
      "tOffsetMs": 2760,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 740200,
    "dDurationMs": 2230,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 740210,
    "dDurationMs": 4770,
    "wWinId": 1,
    "segs": [ {
      "utf8": "posts",
      "acAsrConf": 84
    }, {
      "utf8": " on",
      "tOffsetMs": 420,
      "acAsrConf": 252
    }, {
      "utf8": " my",
      "tOffsetMs": 570,
      "acAsrConf": 144
    }, {
      "utf8": " website",
      "tOffsetMs": 720,
      "acAsrConf": 203
    }, {
      "utf8": " for",
      "tOffsetMs": 1170,
      "acAsrConf": 252
    }, {
      "utf8": " each",
      "tOffsetMs": 1440,
      "acAsrConf": 236
    }, {
      "utf8": " lesson",
      "tOffsetMs": 1650,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 742420,
    "dDurationMs": 2560,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 742430,
    "dDurationMs": 6060,
    "wWinId": 1,
    "segs": [ {
      "utf8": "I",
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 150,
      "acAsrConf": 252
    }, {
      "utf8": " blog",
      "tOffsetMs": 420,
      "acAsrConf": 242
    }, {
      "utf8": " post",
      "tOffsetMs": 780,
      "acAsrConf": 217
    }, {
      "utf8": " you",
      "tOffsetMs": 1110,
      "acAsrConf": 226
    }, {
      "utf8": " can",
      "tOffsetMs": 1320,
      "acAsrConf": 232
    }, {
      "utf8": " read",
      "tOffsetMs": 1470,
      "acAsrConf": 226
    }, {
      "utf8": " and",
      "tOffsetMs": 1710,
      "acAsrConf": 255
    }, {
      "utf8": " you",
      "tOffsetMs": 2010,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 744970,
    "dDurationMs": 3520,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 744980,
    "dDurationMs": 5790,
    "wWinId": 1,
    "segs": [ {
      "utf8": "help",
      "acAsrConf": 204
    }, {
      "utf8": " your",
      "tOffsetMs": 420,
      "acAsrConf": 252
    }, {
      "utf8": " ear",
      "tOffsetMs": 480,
      "acAsrConf": 217
    }, {
      "utf8": " get",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " used",
      "tOffsetMs": 1260,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1530,
      "acAsrConf": 252
    }, {
      "utf8": " hearing",
      "tOffsetMs": 2030,
      "acAsrConf": 236
    }, {
      "utf8": " your",
      "tOffsetMs": 3030,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 748480,
    "dDurationMs": 2290,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 748490,
    "dDurationMs": 4440,
    "wWinId": 1,
    "segs": [ {
      "utf8": "own",
      "acAsrConf": 208
    }, {
      "utf8": " voice",
      "tOffsetMs": 60,
      "acAsrConf": 252
    }, {
      "utf8": " speaking",
      "tOffsetMs": 270,
      "acAsrConf": 204
    }, {
      "utf8": " English",
      "tOffsetMs": 1200,
      "acAsrConf": 226
    }, {
      "utf8": " and",
      "tOffsetMs": 1620,
      "acAsrConf": 236
    }, {
      "utf8": " that's",
      "tOffsetMs": 1860,
      "acAsrConf": 206
    } ]
  }, {
    "tStartMs": 750760,
    "dDurationMs": 2170,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 750770,
    "dDurationMs": 5070,
    "wWinId": 1,
    "segs": [ {
      "utf8": "really",
      "acAsrConf": 252
    }, {
      "utf8": " important",
      "tOffsetMs": 180,
      "acAsrConf": 233
    }, {
      "utf8": " because",
      "tOffsetMs": 810,
      "acAsrConf": 207
    }, {
      "utf8": " it",
      "tOffsetMs": 1020,
      "acAsrConf": 236
    }, {
      "utf8": " will",
      "tOffsetMs": 1920,
      "acAsrConf": 252
    }, {
      "utf8": " help",
      "tOffsetMs": 2100,
      "acAsrConf": 204
    } ]
  }, {
    "tStartMs": 752920,
    "dDurationMs": 2920,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 752930,
    "dDurationMs": 5550,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 167
    }, {
      "utf8": " have",
      "tOffsetMs": 420,
      "acAsrConf": 252
    }, {
      "utf8": " more",
      "tOffsetMs": 600,
      "acAsrConf": 207
    }, {
      "utf8": " and",
      "tOffsetMs": 990,
      "acAsrConf": 222
    }, {
      "utf8": " more",
      "tOffsetMs": 1350,
      "acAsrConf": 132
    }, {
      "utf8": " confidence",
      "tOffsetMs": 1590,
      "acAsrConf": 145
    }, {
      "utf8": " so",
      "tOffsetMs": 2310,
      "acAsrConf": 231
    }, {
      "utf8": " the",
      "tOffsetMs": 2760,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 755830,
    "dDurationMs": 2650,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 755840,
    "dDurationMs": 6860,
    "wWinId": 1,
    "segs": [ {
      "utf8": "more",
      "acAsrConf": 252
    }, {
      "utf8": " confidence",
      "tOffsetMs": 330,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1080,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 1170,
      "acAsrConf": 169
    }, {
      "utf8": " the",
      "tOffsetMs": 1260,
      "acAsrConf": 226
    }, {
      "utf8": " easier",
      "tOffsetMs": 1950,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 2430,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 758470,
    "dDurationMs": 4230,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 758480,
    "dDurationMs": 9330,
    "wWinId": 1,
    "segs": [ {
      "utf8": "will",
      "acAsrConf": 186
    }, {
      "utf8": " be",
      "tOffsetMs": 150,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 330,
      "acAsrConf": 204
    }, {
      "utf8": " do",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " step",
      "tOffsetMs": 810,
      "acAsrConf": 252
    }, {
      "utf8": " number",
      "tOffsetMs": 1200,
      "acAsrConf": 143
    }, {
      "utf8": " three",
      "tOffsetMs": 1530,
      "acAsrConf": 252
    }, {
      "utf8": " which",
      "tOffsetMs": 2160,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 2880,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 762690,
    "dDurationMs": 5120,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 762700,
    "dDurationMs": 7060,
    "wWinId": 1,
    "segs": [ {
      "utf8": "speaking",
      "acAsrConf": 252
    }, {
      "utf8": " uh-huh",
      "tOffsetMs": 1080,
      "acAsrConf": 249
    }, {
      "utf8": " to",
      "tOffsetMs": 2140,
      "acAsrConf": 242
    }, {
      "utf8": " someone",
      "tOffsetMs": 3100,
      "acAsrConf": 252
    }, {
      "utf8": " else",
      "tOffsetMs": 3610,
      "acAsrConf": 221
    }, {
      "utf8": " yes",
      "tOffsetMs": 4030,
      "acAsrConf": 116
    }, {
      "utf8": " of",
      "tOffsetMs": 4780,
      "acAsrConf": 206
    } ]
  }, {
    "tStartMs": 767800,
    "dDurationMs": 1960,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 767810,
    "dDurationMs": 4770,
    "wWinId": 1,
    "segs": [ {
      "utf8": "speaking",
      "acAsrConf": 239
    }, {
      "utf8": " to",
      "tOffsetMs": 510,
      "acAsrConf": 222
    }, {
      "utf8": " someone",
      "tOffsetMs": 630,
      "acAsrConf": 252
    }, {
      "utf8": " else",
      "tOffsetMs": 840,
      "acAsrConf": 213
    }, {
      "utf8": " in",
      "tOffsetMs": 1140,
      "acAsrConf": 187
    }, {
      "utf8": " English",
      "tOffsetMs": 1410,
      "acAsrConf": 236
    }, {
      "utf8": " of",
      "tOffsetMs": 1800,
      "acAsrConf": 248
    } ]
  }, {
    "tStartMs": 769750,
    "dDurationMs": 2830,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 769760,
    "dDurationMs": 5090,
    "wWinId": 1,
    "segs": [ {
      "utf8": "course",
      "acAsrConf": 252
    }, {
      "utf8": " oh",
      "tOffsetMs": 30,
      "acAsrConf": 74
    }, {
      "utf8": " yes",
      "tOffsetMs": 450,
      "acAsrConf": 236
    }, {
      "utf8": " of",
      "tOffsetMs": 900,
      "acAsrConf": 252
    }, {
      "utf8": " course",
      "tOffsetMs": 1200,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 1230,
      "acAsrConf": 182
    }, {
      "utf8": " want",
      "tOffsetMs": 2070,
      "acAsrConf": 236
    }, {
      "utf8": " to",
      "tOffsetMs": 2520,
      "acAsrConf": 252
    }, {
      "utf8": " learn",
      "tOffsetMs": 2580,
      "acAsrConf": 203
    } ]
  }, {
    "tStartMs": 772570,
    "dDurationMs": 2280,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 772580,
    "dDurationMs": 6840,
    "wWinId": 1,
    "segs": [ {
      "utf8": "English",
      "acAsrConf": 219
    }, {
      "utf8": " so",
      "tOffsetMs": 420,
      "acAsrConf": 120
    }, {
      "utf8": " I'm",
      "tOffsetMs": 870,
      "acAsrConf": 252
    }, {
      "utf8": " going",
      "tOffsetMs": 1080,
      "acAsrConf": 237
    }, {
      "utf8": " to",
      "tOffsetMs": 1350,
      "acAsrConf": 252
    }, {
      "utf8": " speak",
      "tOffsetMs": 1380,
      "acAsrConf": 230
    }, {
      "utf8": " Portuguese",
      "tOffsetMs": 1740,
      "acAsrConf": 250
    } ]
  }, {
    "tStartMs": 774840,
    "dDurationMs": 4580,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 774850,
    "dDurationMs": 6790,
    "wWinId": 1,
    "segs": [ {
      "utf8": "even",
      "acAsrConf": 178
    }, {
      "utf8": " oh",
      "tOffsetMs": 1000,
      "acAsrConf": 0
    }, {
      "utf8": " yes",
      "tOffsetMs": 1590,
      "acAsrConf": 240
    }, {
      "utf8": " so",
      "tOffsetMs": 2590,
      "acAsrConf": 244
    }, {
      "utf8": " if",
      "tOffsetMs": 2920,
      "acAsrConf": 240
    }, {
      "utf8": " you",
      "tOffsetMs": 3310,
      "acAsrConf": 252
    }, {
      "utf8": " want",
      "tOffsetMs": 3730,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 4000,
      "acAsrConf": 252
    }, {
      "utf8": " improve",
      "tOffsetMs": 4090,
      "acAsrConf": 207
    } ]
  }, {
    "tStartMs": 779410,
    "dDurationMs": 2230,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 779420,
    "dDurationMs": 5400,
    "wWinId": 1,
    "segs": [ {
      "utf8": "your",
      "acAsrConf": 252
    }, {
      "utf8": " speaking",
      "tOffsetMs": 150,
      "acAsrConf": 233
    }, {
      "utf8": " first",
      "tOffsetMs": 210,
      "acAsrConf": 166
    }, {
      "utf8": " of",
      "tOffsetMs": 840,
      "acAsrConf": 252
    }, {
      "utf8": " all",
      "tOffsetMs": 990,
      "acAsrConf": 244
    }, {
      "utf8": " listen",
      "tOffsetMs": 1110,
      "acAsrConf": 244
    }, {
      "utf8": " it's",
      "tOffsetMs": 1410,
      "acAsrConf": 215
    }, {
      "utf8": " a",
      "tOffsetMs": 2100,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 781630,
    "dDurationMs": 3190,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 781640,
    "dDurationMs": 4820,
    "wWinId": 1,
    "segs": [ {
      "utf8": "good",
      "acAsrConf": 252
    }, {
      "utf8": " first",
      "tOffsetMs": 180,
      "acAsrConf": 228
    }, {
      "utf8": " step",
      "tOffsetMs": 420,
      "acAsrConf": 202
    }, {
      "utf8": " then",
      "tOffsetMs": 660,
      "acAsrConf": 224
    }, {
      "utf8": " read",
      "tOffsetMs": 1350,
      "acAsrConf": 255
    }, {
      "utf8": " out",
      "tOffsetMs": 1920,
      "acAsrConf": 219
    }, {
      "utf8": " loud",
      "tOffsetMs": 2190,
      "acAsrConf": 238
    }, {
      "utf8": " train",
      "tOffsetMs": 2250,
      "acAsrConf": 220
    } ]
  }, {
    "tStartMs": 784810,
    "dDurationMs": 1650,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 784820,
    "dDurationMs": 4220,
    "wWinId": 1,
    "segs": [ {
      "utf8": "your",
      "acAsrConf": 227
    }, {
      "utf8": " mouth",
      "tOffsetMs": 240,
      "acAsrConf": 249
    }, {
      "utf8": " and",
      "tOffsetMs": 420,
      "acAsrConf": 216
    }, {
      "utf8": " the",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " third",
      "tOffsetMs": 1050,
      "acAsrConf": 240
    } ]
  }, {
    "tStartMs": 786450,
    "dDurationMs": 2590,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 786460,
    "dDurationMs": 6660,
    "wWinId": 1,
    "segs": [ {
      "utf8": "train",
      "acAsrConf": 208
    }, {
      "utf8": " your",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " social",
      "tOffsetMs": 780,
      "acAsrConf": 236
    }, {
      "utf8": " skills",
      "tOffsetMs": 1410,
      "acAsrConf": 252
    }, {
      "utf8": " by",
      "tOffsetMs": 1590,
      "acAsrConf": 252
    }, {
      "utf8": " practicing",
      "tOffsetMs": 2310,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 789030,
    "dDurationMs": 4090,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 789040,
    "dDurationMs": 7710,
    "wWinId": 1,
    "segs": [ {
      "utf8": "speaking",
      "acAsrConf": 200
    }, {
      "utf8": " with",
      "tOffsetMs": 989,
      "acAsrConf": 239
    }, {
      "utf8": " someone",
      "tOffsetMs": 1470,
      "acAsrConf": 215
    }, {
      "utf8": " so",
      "tOffsetMs": 2370,
      "acAsrConf": 240
    }, {
      "utf8": " this",
      "tOffsetMs": 2880,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 3150,
      "acAsrConf": 252
    }, {
      "utf8": " if",
      "tOffsetMs": 3359,
      "acAsrConf": 236
    }, {
      "utf8": " you",
      "tOffsetMs": 3690,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 793110,
    "dDurationMs": 3640,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 793120,
    "dDurationMs": 6029,
    "wWinId": 1,
    "segs": [ {
      "utf8": "feel",
      "acAsrConf": 252
    }, {
      "utf8": " like",
      "tOffsetMs": 360,
      "acAsrConf": 245
    }, {
      "utf8": " you",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " are",
      "tOffsetMs": 870,
      "acAsrConf": 118
    }, {
      "utf8": " up",
      "tOffsetMs": 1080,
      "acAsrConf": 252
    }, {
      "utf8": " high",
      "tOffsetMs": 1310,
      "acAsrConf": 242
    }, {
      "utf8": " beginner",
      "tOffsetMs": 2310,
      "acAsrConf": 211
    }, {
      "utf8": " or",
      "tOffsetMs": 3180,
      "acAsrConf": 222
    } ]
  }, {
    "tStartMs": 796740,
    "dDurationMs": 2409,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 796750,
    "dDurationMs": 7529,
    "wWinId": 1,
    "segs": [ {
      "utf8": "low",
      "acAsrConf": 252
    }, {
      "utf8": " intermediate",
      "tOffsetMs": 300,
      "acAsrConf": 252
    }, {
      "utf8": " it's",
      "tOffsetMs": 1050,
      "acAsrConf": 196
    }, {
      "utf8": " a",
      "tOffsetMs": 1649,
      "acAsrConf": 252
    }, {
      "utf8": " good",
      "tOffsetMs": 1860,
      "acAsrConf": 252
    }, {
      "utf8": " time",
      "tOffsetMs": 2100,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 2370,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 799139,
    "dDurationMs": 5140,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 799149,
    "dDurationMs": 8190,
    "wWinId": 1,
    "segs": [ {
      "utf8": "start",
      "acAsrConf": 128
    }, {
      "utf8": " speaking",
      "tOffsetMs": 630,
      "acAsrConf": 252
    }, {
      "utf8": " if",
      "tOffsetMs": 1021,
      "acAsrConf": 203
    }, {
      "utf8": " you",
      "tOffsetMs": 1471,
      "acAsrConf": 252
    }, {
      "utf8": " can",
      "tOffsetMs": 2221,
      "acAsrConf": 252
    }, {
      "utf8": " understand",
      "tOffsetMs": 2431,
      "acAsrConf": 252
    }, {
      "utf8": " 70%",
      "tOffsetMs": 4130,
      "acAsrConf": 164
    } ]
  }, {
    "tStartMs": 804269,
    "dDurationMs": 3070,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 804279,
    "dDurationMs": 5821,
    "wWinId": 1,
    "segs": [ {
      "utf8": "of",
      "acAsrConf": 248
    }, {
      "utf8": " our",
      "tOffsetMs": 180,
      "acAsrConf": 252
    }, {
      "utf8": " presentation",
      "tOffsetMs": 571,
      "acAsrConf": 239
    }, {
      "utf8": " today",
      "tOffsetMs": 1531,
      "acAsrConf": 252
    }, {
      "utf8": " yes",
      "tOffsetMs": 1821,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 807329,
    "dDurationMs": 2771,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 807339,
    "dDurationMs": 5370,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 252
    }, {
      "utf8": " should",
      "tOffsetMs": 451,
      "acAsrConf": 226
    }, {
      "utf8": " start",
      "tOffsetMs": 781,
      "acAsrConf": 243
    }, {
      "utf8": " speaking",
      "tOffsetMs": 1231,
      "acAsrConf": 244
    }, {
      "utf8": " with",
      "tOffsetMs": 1651,
      "acAsrConf": 203
    }, {
      "utf8": " a",
      "tOffsetMs": 2071,
      "acAsrConf": 241
    }, {
      "utf8": " native",
      "tOffsetMs": 2310,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 810090,
    "dDurationMs": 2619,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 810100,
    "dDurationMs": 3299,
    "wWinId": 1,
    "segs": [ {
      "utf8": "speaker",
      "acAsrConf": 252
    }, {
      "utf8": " speak",
      "tOffsetMs": 30,
      "acAsrConf": 151
    }, {
      "utf8": " with",
      "tOffsetMs": 989,
      "acAsrConf": 238
    }, {
      "utf8": " anyone",
      "tOffsetMs": 1170,
      "acAsrConf": 252
    }, {
      "utf8": " who",
      "tOffsetMs": 1530,
      "acAsrConf": 216
    }, {
      "utf8": " can",
      "tOffsetMs": 1950,
      "acAsrConf": 252
    }, {
      "utf8": " speak",
      "tOffsetMs": 2310,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 812699,
    "dDurationMs": 700,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 812709,
    "dDurationMs": 2401,
    "wWinId": 1,
    "segs": [ {
      "utf8": "English",
      "acAsrConf": 204
    } ]
  }, {
    "tStartMs": 813389,
    "dDurationMs": 1721,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 813399,
    "dDurationMs": 4651,
    "wWinId": 1,
    "segs": [ {
      "utf8": "even",
      "acAsrConf": 227
    }, {
      "utf8": " if",
      "tOffsetMs": 391,
      "acAsrConf": 239
    }, {
      "utf8": " they're",
      "tOffsetMs": 511,
      "acAsrConf": 252
    }, {
      "utf8": " not",
      "tOffsetMs": 690,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 721,
      "acAsrConf": 85
    }, {
      "utf8": " native",
      "tOffsetMs": 871,
      "acAsrConf": 252
    }, {
      "utf8": " speaker",
      "tOffsetMs": 1261,
      "acAsrConf": 157
    }, {
      "utf8": " and",
      "tOffsetMs": 1591,
      "acAsrConf": 231
    } ]
  }, {
    "tStartMs": 815100,
    "dDurationMs": 2950,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 815110,
    "dDurationMs": 5250,
    "wWinId": 1,
    "segs": [ {
      "utf8": "this",
      "acAsrConf": 216
    }, {
      "utf8": " is",
      "tOffsetMs": 270,
      "acAsrConf": 226
    }, {
      "utf8": " maybe",
      "tOffsetMs": 539,
      "acAsrConf": 255
    }, {
      "utf8": " the",
      "tOffsetMs": 960,
      "acAsrConf": 236
    }, {
      "utf8": " most",
      "tOffsetMs": 1169,
      "acAsrConf": 166
    }, {
      "utf8": " challenging",
      "tOffsetMs": 1530,
      "acAsrConf": 242
    }, {
      "utf8": " thing",
      "tOffsetMs": 2460,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 818040,
    "dDurationMs": 2320,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 818050,
    "dDurationMs": 5760,
    "wWinId": 1,
    "segs": [ {
      "utf8": "to",
      "acAsrConf": 252
    }, {
      "utf8": " do",
      "tOffsetMs": 30,
      "acAsrConf": 204
    }, {
      "utf8": " but",
      "tOffsetMs": 390,
      "acAsrConf": 219
    }, {
      "utf8": " it",
      "tOffsetMs": 720,
      "acAsrConf": 206
    }, {
      "utf8": " is",
      "tOffsetMs": 870,
      "acAsrConf": 148
    }, {
      "utf8": " also",
      "tOffsetMs": 1020,
      "acAsrConf": 252
    }, {
      "utf8": " maybe",
      "tOffsetMs": 1200,
      "acAsrConf": 204
    }, {
      "utf8": " the",
      "tOffsetMs": 1770,
      "acAsrConf": 200
    }, {
      "utf8": " most",
      "tOffsetMs": 2130,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 820350,
    "dDurationMs": 3460,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 820360,
    "dDurationMs": 6140,
    "wWinId": 1,
    "segs": [ {
      "utf8": "important",
      "acAsrConf": 236
    }, {
      "utf8": " thing",
      "tOffsetMs": 840,
      "acAsrConf": 207
    }, {
      "utf8": " to",
      "tOffsetMs": 1620,
      "acAsrConf": 252
    }, {
      "utf8": " do",
      "tOffsetMs": 1830,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1979,
      "acAsrConf": 252
    }, {
      "utf8": " try",
      "tOffsetMs": 2220,
      "acAsrConf": 203
    }, {
      "utf8": " to",
      "tOffsetMs": 2430,
      "acAsrConf": 252
    }, {
      "utf8": " we",
      "tOffsetMs": 2460,
      "acAsrConf": 241
    }, {
      "utf8": " say",
      "tOffsetMs": 3270,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 823800,
    "dDurationMs": 2700,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 823810,
    "dDurationMs": 6060,
    "wWinId": 1,
    "segs": [ {
      "utf8": "put",
      "acAsrConf": 231
    }, {
      "utf8": " yourself",
      "tOffsetMs": 360,
      "acAsrConf": 217
    }, {
      "utf8": " out",
      "tOffsetMs": 899,
      "acAsrConf": 252
    }, {
      "utf8": " there",
      "tOffsetMs": 1140,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1310,
      "acAsrConf": 205
    }, {
      "utf8": " try",
      "tOffsetMs": 2310,
      "acAsrConf": 200
    }, {
      "utf8": " to",
      "tOffsetMs": 2550,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 826490,
    "dDurationMs": 3380,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 826500,
    "dDurationMs": 7120,
    "wWinId": 1,
    "segs": [ {
      "utf8": "challenge",
      "acAsrConf": 235
    }, {
      "utf8": " yourself",
      "tOffsetMs": 1000,
      "acAsrConf": 216
    }, {
      "utf8": " and",
      "tOffsetMs": 1779,
      "acAsrConf": 207
    }, {
      "utf8": " fighting",
      "tOffsetMs": 2080,
      "acAsrConf": 236
    }, {
      "utf8": " you",
      "tOffsetMs": 3070,
      "acAsrConf": 252
    }, {
      "utf8": " know",
      "tOffsetMs": 3339,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 829860,
    "dDurationMs": 3760,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 829870,
    "dDurationMs": 7080,
    "wWinId": 1,
    "segs": [ {
      "utf8": "really",
      "acAsrConf": 207
    }, {
      "utf8": " trying",
      "tOffsetMs": 719,
      "acAsrConf": 220
    }, {
      "utf8": " to",
      "tOffsetMs": 1050,
      "acAsrConf": 240
    }, {
      "utf8": " go",
      "tOffsetMs": 1200,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1380,
      "acAsrConf": 252
    }, {
      "utf8": " say",
      "tOffsetMs": 1710,
      "acAsrConf": 236
    }, {
      "utf8": " I'm",
      "tOffsetMs": 2580,
      "acAsrConf": 249
    }, {
      "utf8": " going",
      "tOffsetMs": 2969,
      "acAsrConf": 216
    }, {
      "utf8": " to",
      "tOffsetMs": 3540,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 833610,
    "dDurationMs": 3340,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 833620,
    "dDurationMs": 5310,
    "wWinId": 1,
    "segs": [ {
      "utf8": "speak",
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 360,
      "acAsrConf": 252
    }, {
      "utf8": " somebody",
      "tOffsetMs": 390,
      "acAsrConf": 140
    }, {
      "utf8": " in",
      "tOffsetMs": 839,
      "acAsrConf": 234
    }, {
      "utf8": " English",
      "tOffsetMs": 1230,
      "acAsrConf": 255
    }, {
      "utf8": " and",
      "tOffsetMs": 2070,
      "acAsrConf": 165
    }, {
      "utf8": " we",
      "tOffsetMs": 3000,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 3180,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 836940,
    "dDurationMs": 1990,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 836950,
    "dDurationMs": 6060,
    "wWinId": 1,
    "segs": [ {
      "utf8": "a",
      "acAsrConf": 252
    }, {
      "utf8": " couple",
      "tOffsetMs": 30,
      "acAsrConf": 244
    }, {
      "utf8": " tips",
      "tOffsetMs": 300,
      "acAsrConf": 201
    }, {
      "utf8": " on",
      "tOffsetMs": 570,
      "acAsrConf": 252
    }, {
      "utf8": " how",
      "tOffsetMs": 810,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 930,
      "acAsrConf": 161
    }, {
      "utf8": " do",
      "tOffsetMs": 990,
      "acAsrConf": 252
    }, {
      "utf8": " that",
      "tOffsetMs": 1139,
      "acAsrConf": 252
    }, {
      "utf8": " yes",
      "tOffsetMs": 1259,
      "acAsrConf": 214
    }, {
      "utf8": " yes",
      "tOffsetMs": 1560,
      "acAsrConf": 210
    } ]
  }, {
    "tStartMs": 838920,
    "dDurationMs": 4090,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 838930,
    "dDurationMs": 8460,
    "wWinId": 1,
    "segs": [ {
      "utf8": "so",
      "acAsrConf": 252
    }, {
      "utf8": " we",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 390,
      "acAsrConf": 232
    }, {
      "utf8": " here",
      "tOffsetMs": 599,
      "acAsrConf": 252
    }, {
      "utf8": " there",
      "tOffsetMs": 1110,
      "acAsrConf": 252
    }, {
      "utf8": " are",
      "tOffsetMs": 1650,
      "acAsrConf": 252
    }, {
      "utf8": " several",
      "tOffsetMs": 2540,
      "acAsrConf": 252
    }, {
      "utf8": " tips",
      "tOffsetMs": 3540,
      "acAsrConf": 208
    } ]
  }, {
    "tStartMs": 843000,
    "dDurationMs": 4390,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 843010,
    "dDurationMs": 6960,
    "wWinId": 1,
    "segs": [ {
      "utf8": "three",
      "acAsrConf": 214
    }, {
      "utf8": " tips",
      "tOffsetMs": 750,
      "acAsrConf": 241
    }, {
      "utf8": " for",
      "tOffsetMs": 1260,
      "acAsrConf": 252
    }, {
      "utf8": " how",
      "tOffsetMs": 1410,
      "acAsrConf": 205
    }, {
      "utf8": " you",
      "tOffsetMs": 1860,
      "acAsrConf": 207
    }, {
      "utf8": " can",
      "tOffsetMs": 2460,
      "acAsrConf": 252
    }, {
      "utf8": " find",
      "tOffsetMs": 2790,
      "acAsrConf": 235
    }, {
      "utf8": " someone",
      "tOffsetMs": 3510,
      "acAsrConf": 218
    } ]
  }, {
    "tStartMs": 847380,
    "dDurationMs": 2590,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 847390,
    "dDurationMs": 4350,
    "wWinId": 1,
    "segs": [ {
      "utf8": "to",
      "acAsrConf": 242
    }, {
      "utf8": " speak",
      "tOffsetMs": 240,
      "acAsrConf": 206
    }, {
      "utf8": " with",
      "tOffsetMs": 690,
      "acAsrConf": 252
    }, {
      "utf8": " because",
      "tOffsetMs": 900,
      "acAsrConf": 252
    }, {
      "utf8": " this",
      "tOffsetMs": 1319,
      "acAsrConf": 249
    }, {
      "utf8": " idea",
      "tOffsetMs": 1949,
      "acAsrConf": 206
    }, {
      "utf8": " you're",
      "tOffsetMs": 2340,
      "acAsrConf": 255
    } ]
  }, {
    "tStartMs": 849960,
    "dDurationMs": 1780,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 849970,
    "dDurationMs": 4710,
    "wWinId": 1,
    "segs": [ {
      "utf8": "not",
      "acAsrConf": 226
    }, {
      "utf8": " in",
      "tOffsetMs": 239,
      "acAsrConf": 252
    }, {
      "utf8": " an",
      "tOffsetMs": 420,
      "acAsrConf": 252
    }, {
      "utf8": " english-speaking",
      "tOffsetMs": 510,
      "acAsrConf": 214
    }, {
      "utf8": " country",
      "tOffsetMs": 869,
      "acAsrConf": 228
    }, {
      "utf8": " how",
      "tOffsetMs": 1170,
      "acAsrConf": 203
    } ]
  }, {
    "tStartMs": 851730,
    "dDurationMs": 2950,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 851740,
    "dDurationMs": 7530,
    "wWinId": 1,
    "segs": [ {
      "utf8": "can",
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 60,
      "acAsrConf": 142
    }, {
      "utf8": " find",
      "tOffsetMs": 330,
      "acAsrConf": 252
    }, {
      "utf8": " someone",
      "tOffsetMs": 539,
      "acAsrConf": 208
    }, {
      "utf8": " so",
      "tOffsetMs": 930,
      "acAsrConf": 219
    }, {
      "utf8": " the",
      "tOffsetMs": 1410,
      "acAsrConf": 252
    }, {
      "utf8": " first",
      "tOffsetMs": 1830,
      "acAsrConf": 252
    }, {
      "utf8": " tip",
      "tOffsetMs": 2190,
      "acAsrConf": 236
    }, {
      "utf8": " is",
      "tOffsetMs": 2700,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 854670,
    "dDurationMs": 4600,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 854680,
    "dDurationMs": 6450,
    "wWinId": 1,
    "segs": [ {
      "utf8": "I",
      "acAsrConf": 236
    }, {
      "utf8": " think",
      "tOffsetMs": 840,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1170,
      "acAsrConf": 209
    }, {
      "utf8": " should",
      "tOffsetMs": 1920,
      "acAsrConf": 244
    }, {
      "utf8": " go",
      "tOffsetMs": 2340,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 2880,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 2940,
      "acAsrConf": 252
    }, {
      "utf8": " Internet",
      "tOffsetMs": 3690,
      "acAsrConf": 202
    }, {
      "utf8": " of",
      "tOffsetMs": 4409,
      "acAsrConf": 167
    } ]
  }, {
    "tStartMs": 859260,
    "dDurationMs": 1870,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 859270,
    "dDurationMs": 4920,
    "wWinId": 1,
    "segs": [ {
      "utf8": "course",
      "acAsrConf": 207
    }, {
      "utf8": " we",
      "tOffsetMs": 450,
      "acAsrConf": 207
    }, {
      "utf8": " are",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " here",
      "tOffsetMs": 750,
      "acAsrConf": 185
    }, {
      "utf8": " on",
      "tOffsetMs": 1050,
      "acAsrConf": 201
    }, {
      "utf8": " the",
      "tOffsetMs": 1230,
      "acAsrConf": 228
    }, {
      "utf8": " internet",
      "tOffsetMs": 1439,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 861120,
    "dDurationMs": 3070,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 861130,
    "dDurationMs": 7350,
    "wWinId": 1,
    "segs": [ {
      "utf8": "together",
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 560,
      "acAsrConf": 200
    }, {
      "utf8": " there",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 2130,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 2340,
      "acAsrConf": 252
    }, {
      "utf8": " website",
      "tOffsetMs": 2370,
      "acAsrConf": 200
    }, {
      "utf8": " I",
      "tOffsetMs": 2910,
      "acAsrConf": 200
    } ]
  }, {
    "tStartMs": 864180,
    "dDurationMs": 4300,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 864190,
    "dDurationMs": 10740,
    "wWinId": 1,
    "segs": [ {
      "utf8": "recommend",
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 480,
      "acAsrConf": 239
    }, {
      "utf8": " website",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 930,
      "acAsrConf": 207
    }, {
      "utf8": " meet",
      "tOffsetMs": 1470,
      "acAsrConf": 243
    }, {
      "utf8": " up.com",
      "tOffsetMs": 2460,
      "acAsrConf": 246
    } ]
  }, {
    "tStartMs": 868470,
    "dDurationMs": 6460,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 868480,
    "dDurationMs": 9359,
    "wWinId": 1,
    "segs": [ {
      "utf8": "M",
      "acAsrConf": 202
    }, {
      "utf8": " eet",
      "tOffsetMs": 680,
      "acAsrConf": 217
    }, {
      "utf8": " u",
      "tOffsetMs": 1680,
      "acAsrConf": 0
    }, {
      "utf8": " P",
      "tOffsetMs": 2280,
      "acAsrConf": 202
    }, {
      "utf8": " calm",
      "tOffsetMs": 2609,
      "acAsrConf": 243
    }, {
      "utf8": " and",
      "tOffsetMs": 3540,
      "acAsrConf": 237
    }, {
      "utf8": " meetup",
      "tOffsetMs": 3810,
      "acAsrConf": 200
    }, {
      "utf8": " calm",
      "tOffsetMs": 4770,
      "acAsrConf": 238
    }, {
      "utf8": " is",
      "tOffsetMs": 5190,
      "acAsrConf": 209
    }, {
      "utf8": " great",
      "tOffsetMs": 5450,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 874920,
    "dDurationMs": 2919,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 874930,
    "dDurationMs": 5099,
    "wWinId": 1,
    "segs": [ {
      "utf8": "for",
      "acAsrConf": 235
    }, {
      "utf8": " finding",
      "tOffsetMs": 570,
      "acAsrConf": 252
    }, {
      "utf8": " people",
      "tOffsetMs": 840,
      "acAsrConf": 164
    }, {
      "utf8": " anywhere",
      "tOffsetMs": 1529,
      "acAsrConf": 248
    }, {
      "utf8": " around",
      "tOffsetMs": 2250,
      "acAsrConf": 224
    }, {
      "utf8": " the",
      "tOffsetMs": 2550,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 877829,
    "dDurationMs": 2200,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 877839,
    "dDurationMs": 5370,
    "wWinId": 1,
    "segs": [ {
      "utf8": "world",
      "acAsrConf": 252
    }, {
      "utf8": " who",
      "tOffsetMs": 31,
      "acAsrConf": 200
    }, {
      "utf8": " are",
      "tOffsetMs": 511,
      "acAsrConf": 252
    }, {
      "utf8": " going",
      "tOffsetMs": 541,
      "acAsrConf": 208
    }, {
      "utf8": " to",
      "tOffsetMs": 661,
      "acAsrConf": 154
    }, {
      "utf8": " a",
      "tOffsetMs": 1141,
      "acAsrConf": 252
    }, {
      "utf8": " cafe",
      "tOffsetMs": 1171,
      "acAsrConf": 228
    }, {
      "utf8": " and",
      "tOffsetMs": 1861,
      "acAsrConf": 243
    }, {
      "utf8": " you",
      "tOffsetMs": 2131,
      "acAsrConf": 178
    } ]
  }, {
    "tStartMs": 880019,
    "dDurationMs": 3190,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 880029,
    "dDurationMs": 6241,
    "wWinId": 1,
    "segs": [ {
      "utf8": "can",
      "acAsrConf": 150
    }, {
      "utf8": " just",
      "tOffsetMs": 691,
      "acAsrConf": 247
    }, {
      "utf8": " buy",
      "tOffsetMs": 961,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 1351,
      "acAsrConf": 252
    }, {
      "utf8": " coffee",
      "tOffsetMs": 1651,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1951,
      "acAsrConf": 211
    }, {
      "utf8": " speak",
      "tOffsetMs": 2341,
      "acAsrConf": 244
    }, {
      "utf8": " English",
      "tOffsetMs": 2971,
      "acAsrConf": 255
    } ]
  }, {
    "tStartMs": 883199,
    "dDurationMs": 3071,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 883209,
    "dDurationMs": 4711,
    "wWinId": 1,
    "segs": [ {
      "utf8": "together",
      "acAsrConf": 187
    }, {
      "utf8": " it's",
      "tOffsetMs": 361,
      "acAsrConf": 209
    }, {
      "utf8": " not",
      "tOffsetMs": 1051,
      "acAsrConf": 217
    }, {
      "utf8": " a",
      "tOffsetMs": 1351,
      "acAsrConf": 252
    }, {
      "utf8": " class",
      "tOffsetMs": 1380,
      "acAsrConf": 249
    }, {
      "utf8": " but",
      "tOffsetMs": 1921,
      "acAsrConf": 218
    }, {
      "utf8": " you're",
      "tOffsetMs": 2820,
      "acAsrConf": 236
    } ]
  }, {
    "tStartMs": 886260,
    "dDurationMs": 1660,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 886270,
    "dDurationMs": 3750,
    "wWinId": 1,
    "segs": [ {
      "utf8": "just",
      "acAsrConf": 227
    }, {
      "utf8": " having",
      "tOffsetMs": 300,
      "acAsrConf": 255
    }, {
      "utf8": " a",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " chance",
      "tOffsetMs": 780,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    }, {
      "utf8": " practice",
      "tOffsetMs": 1140,
      "acAsrConf": 135
    } ]
  }, {
    "tStartMs": 887910,
    "dDurationMs": 2110,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 887920,
    "dDurationMs": 5370,
    "wWinId": 1,
    "segs": [ {
      "utf8": "together",
      "acAsrConf": 236
    }, {
      "utf8": " so",
      "tOffsetMs": 210,
      "acAsrConf": 216
    }, {
      "utf8": " I",
      "tOffsetMs": 570,
      "acAsrConf": 242
    }, {
      "utf8": " recommend",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " checking",
      "tOffsetMs": 1260,
      "acAsrConf": 252
    }, {
      "utf8": " out",
      "tOffsetMs": 1440,
      "acAsrConf": 147
    } ]
  }, {
    "tStartMs": 890010,
    "dDurationMs": 3280,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 890020,
    "dDurationMs": 6180,
    "wWinId": 1,
    "segs": [ {
      "utf8": "meetup",
      "acAsrConf": 226
    }, {
      "utf8": " calm",
      "tOffsetMs": 600,
      "acAsrConf": 234
    }, {
      "utf8": " and",
      "tOffsetMs": 1050,
      "acAsrConf": 201
    }, {
      "utf8": " yeah",
      "tOffsetMs": 1819,
      "acAsrConf": 249
    }, {
      "utf8": " do",
      "tOffsetMs": 2819,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 3000,
      "acAsrConf": 255
    }, {
      "utf8": " have",
      "tOffsetMs": 3090,
      "acAsrConf": 245
    }, {
      "utf8": " any",
      "tOffsetMs": 3210,
      "acAsrConf": 200
    } ]
  }, {
    "tStartMs": 893280,
    "dDurationMs": 2920,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 893290,
    "dDurationMs": 5130,
    "wWinId": 1,
    "segs": [ {
      "utf8": "personal",
      "acAsrConf": 142
    }, {
      "utf8": " examples",
      "tOffsetMs": 690,
      "acAsrConf": 252
    }, {
      "utf8": " with",
      "tOffsetMs": 1289,
      "acAsrConf": 252
    }, {
      "utf8": " meetup",
      "tOffsetMs": 1560,
      "acAsrConf": 245
    }, {
      "utf8": " comm",
      "tOffsetMs": 2070,
      "acAsrConf": 228
    }, {
      "utf8": " that",
      "tOffsetMs": 2610,
      "acAsrConf": 247
    } ]
  }, {
    "tStartMs": 896190,
    "dDurationMs": 2230,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 896200,
    "dDurationMs": 4950,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you've",
      "acAsrConf": 252
    }, {
      "utf8": " used",
      "tOffsetMs": 150,
      "acAsrConf": 252
    }, {
      "utf8": " before",
      "tOffsetMs": 420,
      "acAsrConf": 231
    }, {
      "utf8": " um",
      "tOffsetMs": 570,
      "acAsrConf": 92
    }, {
      "utf8": " not",
      "tOffsetMs": 1230,
      "acAsrConf": 236
    }, {
      "utf8": " necessarily",
      "tOffsetMs": 1950,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 898410,
    "dDurationMs": 2740,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 898420,
    "dDurationMs": 5400,
    "wWinId": 1,
    "segs": [ {
      "utf8": "that",
      "acAsrConf": 151
    }, {
      "utf8": " I've",
      "tOffsetMs": 780,
      "acAsrConf": 226
    }, {
      "utf8": " use",
      "tOffsetMs": 1020,
      "acAsrConf": 205
    }, {
      "utf8": " but",
      "tOffsetMs": 1380,
      "acAsrConf": 235
    }, {
      "utf8": " I've",
      "tOffsetMs": 1680,
      "acAsrConf": 125
    }, {
      "utf8": " definitely",
      "tOffsetMs": 1800,
      "acAsrConf": 130
    }, {
      "utf8": " seen",
      "tOffsetMs": 2220,
      "acAsrConf": 143
    } ]
  }, {
    "tStartMs": 901140,
    "dDurationMs": 2680,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 901150,
    "dDurationMs": 5370,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 252
    }, {
      "utf8": " use",
      "tOffsetMs": 330,
      "acAsrConf": 252
    }, {
      "utf8": " meetup",
      "tOffsetMs": 660,
      "acAsrConf": 238
    }, {
      "utf8": " right",
      "tOffsetMs": 1050,
      "acAsrConf": 227
    }, {
      "utf8": " the",
      "tOffsetMs": 1830,
      "acAsrConf": 239
    }, {
      "utf8": " best",
      "tOffsetMs": 2309,
      "acAsrConf": 252
    }, {
      "utf8": " thing",
      "tOffsetMs": 2520,
      "acAsrConf": 216
    } ]
  }, {
    "tStartMs": 903810,
    "dDurationMs": 2710,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 903820,
    "dDurationMs": 5519,
    "wWinId": 1,
    "segs": [ {
      "utf8": "about",
      "acAsrConf": 64
    }, {
      "utf8": " meetup",
      "tOffsetMs": 120,
      "acAsrConf": 252
    }, {
      "utf8": " comm",
      "tOffsetMs": 660,
      "acAsrConf": 208
    }, {
      "utf8": " is",
      "tOffsetMs": 1079,
      "acAsrConf": 129
    }, {
      "utf8": " it's",
      "tOffsetMs": 1410,
      "acAsrConf": 234
    }, {
      "utf8": " just",
      "tOffsetMs": 1620,
      "acAsrConf": 252
    }, {
      "utf8": " people",
      "tOffsetMs": 1709,
      "acAsrConf": 204
    } ]
  }, {
    "tStartMs": 906510,
    "dDurationMs": 2829,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 906520,
    "dDurationMs": 3720,
    "wWinId": 1,
    "segs": [ {
      "utf8": "coming",
      "acAsrConf": 151
    }, {
      "utf8": " together",
      "tOffsetMs": 540,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 689,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 960,
      "acAsrConf": 174
    }, {
      "utf8": " group",
      "tOffsetMs": 1350,
      "acAsrConf": 227
    }, {
      "utf8": " in",
      "tOffsetMs": 1650,
      "acAsrConf": 247
    }, {
      "utf8": " your",
      "tOffsetMs": 1920,
      "acAsrConf": 252
    }, {
      "utf8": " local",
      "tOffsetMs": 1980,
      "acAsrConf": 138
    } ]
  }, {
    "tStartMs": 909329,
    "dDurationMs": 911,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 909339,
    "dDurationMs": 3481,
    "wWinId": 1,
    "segs": [ {
      "utf8": "area",
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 910230,
    "dDurationMs": 2590,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 910240,
    "dDurationMs": 5310,
    "wWinId": 1,
    "segs": [ {
      "utf8": "who",
      "acAsrConf": 252
    }, {
      "utf8": " want",
      "tOffsetMs": 330,
      "acAsrConf": 216
    }, {
      "utf8": " to",
      "tOffsetMs": 599,
      "acAsrConf": 252
    }, {
      "utf8": " do",
      "tOffsetMs": 690,
      "acAsrConf": 217
    }, {
      "utf8": " similar",
      "tOffsetMs": 870,
      "acAsrConf": 209
    }, {
      "utf8": " things",
      "tOffsetMs": 1320,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1440,
      "acAsrConf": 165
    }, {
      "utf8": " it's",
      "tOffsetMs": 2219,
      "acAsrConf": 251
    } ]
  }, {
    "tStartMs": 912810,
    "dDurationMs": 2740,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 912820,
    "dDurationMs": 4350,
    "wWinId": 1,
    "segs": [ {
      "utf8": "very",
      "acAsrConf": 225
    }, {
      "utf8": " simple",
      "tOffsetMs": 150,
      "acAsrConf": 207
    }, {
      "utf8": " very",
      "tOffsetMs": 750,
      "acAsrConf": 231
    }, {
      "utf8": " basic",
      "tOffsetMs": 810,
      "acAsrConf": 163
    }, {
      "utf8": " but",
      "tOffsetMs": 1250,
      "acAsrConf": 216
    }, {
      "utf8": " you",
      "tOffsetMs": 2250,
      "acAsrConf": 166
    }, {
      "utf8": " have",
      "tOffsetMs": 2370,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 2550,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 915540,
    "dDurationMs": 1630,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 915550,
    "dDurationMs": 3380,
    "wWinId": 1,
    "segs": [ {
      "utf8": "be",
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 150,
      "acAsrConf": 252
    }, {
      "utf8": " little",
      "tOffsetMs": 180,
      "acAsrConf": 252
    }, {
      "utf8": " brave",
      "tOffsetMs": 360,
      "acAsrConf": 202
    }, {
      "utf8": " because",
      "tOffsetMs": 750,
      "acAsrConf": 236
    }, {
      "utf8": " you",
      "tOffsetMs": 1170,
      "acAsrConf": 167
    }, {
      "utf8": " have",
      "tOffsetMs": 1350,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1470,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 917160,
    "dDurationMs": 1770,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 917170,
    "dDurationMs": 4160,
    "wWinId": 1,
    "segs": [ {
      "utf8": "put",
      "acAsrConf": 252
    }, {
      "utf8": " yourself",
      "tOffsetMs": 270,
      "acAsrConf": 180
    }, {
      "utf8": " out",
      "tOffsetMs": 659,
      "acAsrConf": 227
    }, {
      "utf8": " there",
      "tOffsetMs": 870,
      "acAsrConf": 252
    }, {
      "utf8": " yeah",
      "tOffsetMs": 900,
      "acAsrConf": 231
    } ]
  }, {
    "tStartMs": 918920,
    "dDurationMs": 2410,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 918930,
    "dDurationMs": 4409,
    "wWinId": 1,
    "segs": [ {
      "utf8": "oh",
      "acAsrConf": 113
    }, {
      "utf8": " there",
      "tOffsetMs": 60,
      "acAsrConf": 190
    }, {
      "utf8": " do",
      "tOffsetMs": 659,
      "acAsrConf": 252
    }, {
      "utf8": " something",
      "tOffsetMs": 840,
      "acAsrConf": 205
    }, {
      "utf8": " uncomfortable",
      "tOffsetMs": 1320,
      "acAsrConf": 250
    }, {
      "utf8": " yeah",
      "tOffsetMs": 1740,
      "acAsrConf": 122
    } ]
  }, {
    "tStartMs": 921320,
    "dDurationMs": 2019,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 921330,
    "dDurationMs": 4949,
    "wWinId": 1,
    "segs": [ {
      "utf8": "it's",
      "acAsrConf": 249
    }, {
      "utf8": " a",
      "tOffsetMs": 540,
      "acAsrConf": 252
    }, {
      "utf8": " little",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " uncomfortable",
      "tOffsetMs": 720,
      "acAsrConf": 206
    }, {
      "utf8": " but",
      "tOffsetMs": 1140,
      "acAsrConf": 252
    }, {
      "utf8": " once",
      "tOffsetMs": 1650,
      "acAsrConf": 227
    }, {
      "utf8": " you",
      "tOffsetMs": 1860,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 923329,
    "dDurationMs": 2950,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 923339,
    "dDurationMs": 4951,
    "wWinId": 1,
    "segs": [ {
      "utf8": "do",
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 211,
      "acAsrConf": 252
    }, {
      "utf8": " you'll",
      "tOffsetMs": 361,
      "acAsrConf": 252
    }, {
      "utf8": " feel",
      "tOffsetMs": 571,
      "acAsrConf": 252
    }, {
      "utf8": " very",
      "tOffsetMs": 750,
      "acAsrConf": 204
    }, {
      "utf8": " proud",
      "tOffsetMs": 810,
      "acAsrConf": 198
    }, {
      "utf8": " yes",
      "tOffsetMs": 1351,
      "acAsrConf": 213
    }, {
      "utf8": " yes",
      "tOffsetMs": 2011,
      "acAsrConf": 236
    }, {
      "utf8": " so",
      "tOffsetMs": 2581,
      "acAsrConf": 205
    } ]
  }, {
    "tStartMs": 926269,
    "dDurationMs": 2021,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 926279,
    "dDurationMs": 3000,
    "wWinId": 1,
    "segs": [ {
      "utf8": "I",
      "acAsrConf": 247
    }, {
      "utf8": " personally",
      "tOffsetMs": 211,
      "acAsrConf": 242
    }, {
      "utf8": " I've",
      "tOffsetMs": 721,
      "acAsrConf": 250
    }, {
      "utf8": " used",
      "tOffsetMs": 1201,
      "acAsrConf": 252
    }, {
      "utf8": " this",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    }, {
      "utf8": " website",
      "tOffsetMs": 1771,
      "acAsrConf": 244
    } ]
  }, {
    "tStartMs": 928280,
    "dDurationMs": 999,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 928290,
    "dDurationMs": 4260,
    "wWinId": 1,
    "segs": [ {
      "utf8": "meetup.com",
      "acAsrConf": 207
    } ]
  }, {
    "tStartMs": 929269,
    "dDurationMs": 3281,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 929279,
    "dDurationMs": 6721,
    "wWinId": 1,
    "segs": [ {
      "utf8": "so",
      "acAsrConf": 225
    }, {
      "utf8": " much",
      "tOffsetMs": 901,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 1351,
      "acAsrConf": 249
    }, {
      "utf8": " fact",
      "tOffsetMs": 1560,
      "acAsrConf": 236
    }, {
      "utf8": " yesterday",
      "tOffsetMs": 1771,
      "acAsrConf": 216
    }, {
      "utf8": " I",
      "tOffsetMs": 2101,
      "acAsrConf": 201
    }, {
      "utf8": " went",
      "tOffsetMs": 2611,
      "acAsrConf": 201
    }, {
      "utf8": " to",
      "tOffsetMs": 3031,
      "acAsrConf": 144
    }, {
      "utf8": " a",
      "tOffsetMs": 3241,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 932540,
    "dDurationMs": 3460,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 932550,
    "dDurationMs": 6330,
    "wWinId": 1,
    "segs": [ {
      "utf8": "local",
      "acAsrConf": 252
    }, {
      "utf8": " French",
      "tOffsetMs": 390,
      "acAsrConf": 198
    }, {
      "utf8": " meetup",
      "tOffsetMs": 1050,
      "acAsrConf": 221
    }, {
      "utf8": " group",
      "tOffsetMs": 1950,
      "acAsrConf": 252
    }, {
      "utf8": " we",
      "tOffsetMs": 2340,
      "acAsrConf": 205
    }, {
      "utf8": " met",
      "tOffsetMs": 2849,
      "acAsrConf": 252
    }, {
      "utf8": " at",
      "tOffsetMs": 3150,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 3390,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 935990,
    "dDurationMs": 2890,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 936000,
    "dDurationMs": 5520,
    "wWinId": 1,
    "segs": [ {
      "utf8": "cafe",
      "acAsrConf": 252
    }, {
      "utf8": " everyone",
      "tOffsetMs": 649,
      "acAsrConf": 166
    }, {
      "utf8": " with",
      "tOffsetMs": 1649,
      "acAsrConf": 166
    }, {
      "utf8": " strangers",
      "tOffsetMs": 1830,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 2310,
      "acAsrConf": 84
    }, {
      "utf8": " didn't",
      "tOffsetMs": 2520,
      "acAsrConf": 249
    } ]
  }, {
    "tStartMs": 938870,
    "dDurationMs": 2650,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 938880,
    "dDurationMs": 5340,
    "wWinId": 1,
    "segs": [ {
      "utf8": "know",
      "acAsrConf": 252
    }, {
      "utf8": " anyone",
      "tOffsetMs": 149,
      "acAsrConf": 202
    }, {
      "utf8": " there",
      "tOffsetMs": 840,
      "acAsrConf": 141
    }, {
      "utf8": " and",
      "tOffsetMs": 1050,
      "acAsrConf": 131
    }, {
      "utf8": " we",
      "tOffsetMs": 1380,
      "acAsrConf": 252
    }, {
      "utf8": " just",
      "tOffsetMs": 2100,
      "acAsrConf": 252
    }, {
      "utf8": " sat",
      "tOffsetMs": 2130,
      "acAsrConf": 157
    }, {
      "utf8": " down",
      "tOffsetMs": 2610,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 941510,
    "dDurationMs": 2710,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 941520,
    "dDurationMs": 5280,
    "wWinId": 1,
    "segs": [ {
      "utf8": "and",
      "acAsrConf": 233
    }, {
      "utf8": " started",
      "tOffsetMs": 540,
      "acAsrConf": 252
    }, {
      "utf8": " talking",
      "tOffsetMs": 720,
      "acAsrConf": 215
    }, {
      "utf8": " hey",
      "tOffsetMs": 1230,
      "acAsrConf": 232
    }, {
      "utf8": " where",
      "tOffsetMs": 1920,
      "acAsrConf": 240
    }, {
      "utf8": " are",
      "tOffsetMs": 2460,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 2490,
      "acAsrConf": 202
    } ]
  }, {
    "tStartMs": 944210,
    "dDurationMs": 2590,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 944220,
    "dDurationMs": 4710,
    "wWinId": 1,
    "segs": [ {
      "utf8": "from",
      "acAsrConf": 255
    }, {
      "utf8": " how",
      "tOffsetMs": 270,
      "acAsrConf": 221
    }, {
      "utf8": " did",
      "tOffsetMs": 660,
      "acAsrConf": 220
    }, {
      "utf8": " you",
      "tOffsetMs": 869,
      "acAsrConf": 252
    }, {
      "utf8": " learn",
      "tOffsetMs": 1020,
      "acAsrConf": 252
    }, {
      "utf8": " French",
      "tOffsetMs": 1170,
      "acAsrConf": 252
    }, {
      "utf8": " how",
      "tOffsetMs": 1500,
      "acAsrConf": 146
    }, {
      "utf8": " long",
      "tOffsetMs": 2309,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 946790,
    "dDurationMs": 2140,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 946800,
    "dDurationMs": 4320,
    "wWinId": 1,
    "segs": [ {
      "utf8": "have",
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 240,
      "acAsrConf": 201
    }, {
      "utf8": " lived",
      "tOffsetMs": 330,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 360,
      "acAsrConf": 203
    }, {
      "utf8": " the",
      "tOffsetMs": 750,
      "acAsrConf": 226
    }, {
      "utf8": " city",
      "tOffsetMs": 900,
      "acAsrConf": 212
    }, {
      "utf8": " and",
      "tOffsetMs": 1170,
      "acAsrConf": 227
    }, {
      "utf8": " you",
      "tOffsetMs": 1440,
      "acAsrConf": 252
    }, {
      "utf8": " could",
      "tOffsetMs": 1830,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 948920,
    "dDurationMs": 2200,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 948930,
    "dDurationMs": 5340,
    "wWinId": 1,
    "segs": [ {
      "utf8": "do",
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 150,
      "acAsrConf": 252
    }, {
      "utf8": " same",
      "tOffsetMs": 300,
      "acAsrConf": 252
    }, {
      "utf8": " thing",
      "tOffsetMs": 480,
      "acAsrConf": 151
    }, {
      "utf8": " in",
      "tOffsetMs": 840,
      "acAsrConf": 233
    }, {
      "utf8": " your",
      "tOffsetMs": 1020,
      "acAsrConf": 252
    }, {
      "utf8": " city",
      "tOffsetMs": 1230,
      "acAsrConf": 227
    }, {
      "utf8": " I've",
      "tOffsetMs": 1590,
      "acAsrConf": 175
    } ]
  }, {
    "tStartMs": 951110,
    "dDurationMs": 3160,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 951120,
    "dDurationMs": 6810,
    "wWinId": 1,
    "segs": [ {
      "utf8": "checked",
      "acAsrConf": 236
    }, {
      "utf8": " many",
      "tOffsetMs": 510,
      "acAsrConf": 202
    }, {
      "utf8": " of",
      "tOffsetMs": 1110,
      "acAsrConf": 252
    }, {
      "utf8": " my",
      "tOffsetMs": 1380,
      "acAsrConf": 216
    }, {
      "utf8": " students",
      "tOffsetMs": 1529,
      "acAsrConf": 252
    }, {
      "utf8": " hometowns",
      "tOffsetMs": 1980,
      "acAsrConf": 207
    }, {
      "utf8": " on",
      "tOffsetMs": 2760,
      "acAsrConf": 137
    } ]
  }, {
    "tStartMs": 954260,
    "dDurationMs": 3670,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 954270,
    "dDurationMs": 6900,
    "wWinId": 1,
    "segs": [ {
      "utf8": "Meetup",
      "acAsrConf": 143
    }, {
      "utf8": " calm",
      "tOffsetMs": 480,
      "acAsrConf": 122
    }, {
      "utf8": " and",
      "tOffsetMs": 990,
      "acAsrConf": 228
    }, {
      "utf8": " almost",
      "tOffsetMs": 1290,
      "acAsrConf": 211
    }, {
      "utf8": " always",
      "tOffsetMs": 2040,
      "acAsrConf": 252
    }, {
      "utf8": " there",
      "tOffsetMs": 2370,
      "acAsrConf": 243
    }, {
      "utf8": " are",
      "tOffsetMs": 3360,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 957920,
    "dDurationMs": 3250,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 957930,
    "dDurationMs": 5700,
    "wWinId": 1,
    "segs": [ {
      "utf8": "English",
      "acAsrConf": 227
    }, {
      "utf8": " meetup",
      "tOffsetMs": 420,
      "acAsrConf": 208
    }, {
      "utf8": " groups",
      "tOffsetMs": 1110,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 1440,
      "acAsrConf": 221
    }, {
      "utf8": " their",
      "tOffsetMs": 1740,
      "acAsrConf": 252
    }, {
      "utf8": " city",
      "tOffsetMs": 2159,
      "acAsrConf": 69
    }, {
      "utf8": " so",
      "tOffsetMs": 2460,
      "acAsrConf": 202
    }, {
      "utf8": " I",
      "tOffsetMs": 3000,
      "acAsrConf": 239
    } ]
  }, {
    "tStartMs": 961160,
    "dDurationMs": 2470,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 961170,
    "dDurationMs": 4410,
    "wWinId": 1,
    "segs": [ {
      "utf8": "recommend",
      "acAsrConf": 215
    }, {
      "utf8": " if",
      "tOffsetMs": 570,
      "acAsrConf": 255
    }, {
      "utf8": " you",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " are",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    }, {
      "utf8": " really",
      "tOffsetMs": 1160,
      "acAsrConf": 200
    }, {
      "utf8": " serious",
      "tOffsetMs": 2160,
      "acAsrConf": 217
    } ]
  }, {
    "tStartMs": 963620,
    "dDurationMs": 1960,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 963630,
    "dDurationMs": 4860,
    "wWinId": 1,
    "segs": [ {
      "utf8": "about",
      "acAsrConf": 252
    }, {
      "utf8": " improving",
      "tOffsetMs": 209,
      "acAsrConf": 216
    }, {
      "utf8": " your",
      "tOffsetMs": 810,
      "acAsrConf": 200
    }, {
      "utf8": " speaking",
      "tOffsetMs": 1230,
      "acAsrConf": 217
    } ]
  }, {
    "tStartMs": 965570,
    "dDurationMs": 2920,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 965580,
    "dDurationMs": 4980,
    "wWinId": 1,
    "segs": [ {
      "utf8": "check",
      "acAsrConf": 239
    }, {
      "utf8": " out",
      "tOffsetMs": 390,
      "acAsrConf": 239
    }, {
      "utf8": " meetup",
      "tOffsetMs": 600,
      "acAsrConf": 201
    }, {
      "utf8": " comm",
      "tOffsetMs": 1110,
      "acAsrConf": 82
    }, {
      "utf8": " I",
      "tOffsetMs": 1530,
      "acAsrConf": 236
    }, {
      "utf8": " don't",
      "tOffsetMs": 1800,
      "acAsrConf": 252
    }, {
      "utf8": " work",
      "tOffsetMs": 2520,
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 2879,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 968480,
    "dDurationMs": 2080,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 968490,
    "dDurationMs": 4550,
    "wWinId": 1,
    "segs": [ {
      "utf8": "their",
      "acAsrConf": 200
    }, {
      "utf8": " website",
      "tOffsetMs": 420,
      "acAsrConf": 236
    }, {
      "utf8": " I",
      "tOffsetMs": 930,
      "acAsrConf": 252
    }, {
      "utf8": " just",
      "tOffsetMs": 1110,
      "acAsrConf": 252
    }, {
      "utf8": " want",
      "tOffsetMs": 1349,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1530,
      "acAsrConf": 252
    }, {
      "utf8": " recommend",
      "tOffsetMs": 1620,
      "acAsrConf": 187
    } ]
  }, {
    "tStartMs": 970550,
    "dDurationMs": 2490,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 970560,
    "dDurationMs": 5460,
    "wWinId": 1,
    "segs": [ {
      "utf8": "it",
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 89,
      "acAsrConf": 168
    }, {
      "utf8": " you",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " as",
      "tOffsetMs": 390,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " great",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " in-person",
      "tOffsetMs": 1380,
      "acAsrConf": 236
    } ]
  }, {
    "tStartMs": 973030,
    "dDurationMs": 2990,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 973040,
    "dDurationMs": 5350,
    "wWinId": 1,
    "segs": [ {
      "utf8": "opportunity",
      "acAsrConf": 231
    }, {
      "utf8": " because",
      "tOffsetMs": 1000,
      "acAsrConf": 244
    }, {
      "utf8": " language",
      "tOffsetMs": 1450,
      "acAsrConf": 206
    }, {
      "utf8": " classes",
      "tOffsetMs": 2320,
      "acAsrConf": 252
    }, {
      "utf8": " are",
      "tOffsetMs": 2830,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 976010,
    "dDurationMs": 2380,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 976020,
    "dDurationMs": 4800,
    "wWinId": 1,
    "segs": [ {
      "utf8": "okay",
      "acAsrConf": 140
    }, {
      "utf8": " but",
      "tOffsetMs": 450,
      "acAsrConf": 219
    }, {
      "utf8": " really",
      "tOffsetMs": 870,
      "acAsrConf": 204
    }, {
      "utf8": " you",
      "tOffsetMs": 1620,
      "acAsrConf": 252
    }, {
      "utf8": " want",
      "tOffsetMs": 1800,
      "acAsrConf": 201
    }, {
      "utf8": " real",
      "tOffsetMs": 2040,
      "acAsrConf": 249
    } ]
  }, {
    "tStartMs": 978380,
    "dDurationMs": 2440,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 978390,
    "dDurationMs": 4470,
    "wWinId": 1,
    "segs": [ {
      "utf8": "experiences",
      "acAsrConf": 245
    }, {
      "utf8": " so",
      "tOffsetMs": 840,
      "acAsrConf": 167
    }, {
      "utf8": " this",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 1199,
      "acAsrConf": 252
    }, {
      "utf8": " real",
      "tOffsetMs": 1380,
      "acAsrConf": 252
    }, {
      "utf8": " experiences",
      "tOffsetMs": 1680,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 980810,
    "dDurationMs": 2050,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 980820,
    "dDurationMs": 4860,
    "wWinId": 1,
    "segs": [ {
      "utf8": "can",
      "acAsrConf": 247
    }, {
      "utf8": " you",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " meet",
      "tOffsetMs": 420,
      "acAsrConf": 186
    }, {
      "utf8": " native",
      "tOffsetMs": 570,
      "acAsrConf": 233
    }, {
      "utf8": " speakers",
      "tOffsetMs": 1110,
      "acAsrConf": 252
    }, {
      "utf8": " there",
      "tOffsetMs": 1530,
      "acAsrConf": 255
    }, {
      "utf8": " too",
      "tOffsetMs": 1769,
      "acAsrConf": 210
    } ]
  }, {
    "tStartMs": 982850,
    "dDurationMs": 2830,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 982860,
    "dDurationMs": 5040,
    "wWinId": 1,
    "segs": [ {
      "utf8": "yes",
      "acAsrConf": 209
    }, {
      "utf8": " if",
      "tOffsetMs": 450,
      "acAsrConf": 217
    }, {
      "utf8": " you",
      "tOffsetMs": 690,
      "acAsrConf": 246
    }, {
      "utf8": " live",
      "tOffsetMs": 990,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 1260,
      "acAsrConf": 233
    }, {
      "utf8": " a",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    }, {
      "utf8": " medium-sized",
      "tOffsetMs": 1620,
      "acAsrConf": 226
    }, {
      "utf8": " city",
      "tOffsetMs": 2490,
      "acAsrConf": 202
    } ]
  }, {
    "tStartMs": 985670,
    "dDurationMs": 2230,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 985680,
    "dDurationMs": 5490,
    "wWinId": 1,
    "segs": [ {
      "utf8": "almost",
      "acAsrConf": 156
    }, {
      "utf8": " always",
      "tOffsetMs": 390,
      "acAsrConf": 144
    }, {
      "utf8": " there's",
      "tOffsetMs": 930,
      "acAsrConf": 252
    }, {
      "utf8": " going",
      "tOffsetMs": 1409,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1710,
      "acAsrConf": 252
    }, {
      "utf8": " be",
      "tOffsetMs": 1800,
      "acAsrConf": 252
    }, {
      "utf8": " native",
      "tOffsetMs": 1860,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 987890,
    "dDurationMs": 3280,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 987900,
    "dDurationMs": 5760,
    "wWinId": 1,
    "segs": [ {
      "utf8": "speakers",
      "acAsrConf": 219
    }, {
      "utf8": " there",
      "tOffsetMs": 750,
      "acAsrConf": 255
    }, {
      "utf8": " so",
      "tOffsetMs": 1080,
      "acAsrConf": 226
    }, {
      "utf8": " yes",
      "tOffsetMs": 1830,
      "acAsrConf": 137
    }, {
      "utf8": " it's",
      "tOffsetMs": 2670,
      "acAsrConf": 205
    }, {
      "utf8": " a",
      "tOffsetMs": 3000,
      "acAsrConf": 252
    }, {
      "utf8": " great",
      "tOffsetMs": 3090,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 991160,
    "dDurationMs": 2500,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 991170,
    "dDurationMs": 5640,
    "wWinId": 1,
    "segs": [ {
      "utf8": "chance",
      "acAsrConf": 203
    }, {
      "utf8": " and",
      "tOffsetMs": 390,
      "acAsrConf": 208
    }, {
      "utf8": " let's",
      "tOffsetMs": 690,
      "acAsrConf": 244
    }, {
      "utf8": " go",
      "tOffsetMs": 1350,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1530,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 1590,
      "acAsrConf": 252
    }, {
      "utf8": " second",
      "tOffsetMs": 2010,
      "acAsrConf": 255
    }, {
      "utf8": " tip",
      "tOffsetMs": 2370,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 993650,
    "dDurationMs": 3160,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 993660,
    "dDurationMs": 5880,
    "wWinId": 1,
    "segs": [ {
      "utf8": "how",
      "acAsrConf": 230
    }, {
      "utf8": " to",
      "tOffsetMs": 390,
      "acAsrConf": 143
    }, {
      "utf8": " find",
      "tOffsetMs": 929,
      "acAsrConf": 249
    }, {
      "utf8": " people",
      "tOffsetMs": 1619,
      "acAsrConf": 252
    }, {
      "utf8": " who",
      "tOffsetMs": 2040,
      "acAsrConf": 207
    }, {
      "utf8": " you",
      "tOffsetMs": 2220,
      "acAsrConf": 106
    }, {
      "utf8": " can",
      "tOffsetMs": 2610,
      "acAsrConf": 252
    }, {
      "utf8": " speak",
      "tOffsetMs": 2790,
      "acAsrConf": 238
    } ]
  }, {
    "tStartMs": 996800,
    "dDurationMs": 2740,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 996810,
    "dDurationMs": 6420,
    "wWinId": 1,
    "segs": [ {
      "utf8": "with",
      "acAsrConf": 226
    }, {
      "utf8": " so",
      "tOffsetMs": 180,
      "acAsrConf": 245
    }, {
      "utf8": " the",
      "tOffsetMs": 360,
      "acAsrConf": 252
    }, {
      "utf8": " first",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " one",
      "tOffsetMs": 719,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 899,
      "acAsrConf": 252
    }, {
      "utf8": " meetup",
      "tOffsetMs": 930,
      "acAsrConf": 201
    }, {
      "utf8": " calm",
      "tOffsetMs": 1860,
      "acAsrConf": 228
    }, {
      "utf8": " the",
      "tOffsetMs": 2250,
      "acAsrConf": 249
    } ]
  }, {
    "tStartMs": 999530,
    "dDurationMs": 3700,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 999540,
    "dDurationMs": 7830,
    "wWinId": 1,
    "segs": [ {
      "utf8": "second",
      "acAsrConf": 249
    }, {
      "utf8": " one",
      "tOffsetMs": 360,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " free",
      "tOffsetMs": 810,
      "acAsrConf": 247
    }, {
      "utf8": " language",
      "tOffsetMs": 1799,
      "acAsrConf": 230
    }, {
      "utf8": " exchanges",
      "tOffsetMs": 2730,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1003220,
    "dDurationMs": 4150,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1003230,
    "dDurationMs": 6720,
    "wWinId": 1,
    "segs": [ {
      "utf8": "online",
      "acAsrConf": 227
    }, {
      "utf8": " so",
      "tOffsetMs": 990,
      "acAsrConf": 249
    }, {
      "utf8": " free",
      "tOffsetMs": 1799,
      "acAsrConf": 252
    }, {
      "utf8": " language",
      "tOffsetMs": 2310,
      "acAsrConf": 252
    }, {
      "utf8": " exchanges",
      "tOffsetMs": 2940,
      "acAsrConf": 252
    }, {
      "utf8": " online",
      "tOffsetMs": 3540,
      "acAsrConf": 218
    } ]
  }, {
    "tStartMs": 1007360,
    "dDurationMs": 2590,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1007370,
    "dDurationMs": 5550,
    "wWinId": 1,
    "segs": [ {
      "utf8": "are",
      "acAsrConf": 212
    }, {
      "utf8": " places",
      "tOffsetMs": 380,
      "acAsrConf": 250
    }, {
      "utf8": " where",
      "tOffsetMs": 1380,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1530,
      "acAsrConf": 213
    }, {
      "utf8": " can",
      "tOffsetMs": 1589,
      "acAsrConf": 125
    }, {
      "utf8": " meet",
      "tOffsetMs": 1740,
      "acAsrConf": 143
    }, {
      "utf8": " up",
      "tOffsetMs": 2190,
      "acAsrConf": 252
    }, {
      "utf8": " with",
      "tOffsetMs": 2339,
      "acAsrConf": 126
    } ]
  }, {
    "tStartMs": 1009940,
    "dDurationMs": 2980,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1009950,
    "dDurationMs": 6509,
    "wWinId": 1,
    "segs": [ {
      "utf8": "someone",
      "acAsrConf": 211
    }, {
      "utf8": " who",
      "tOffsetMs": 360,
      "acAsrConf": 216
    }, {
      "utf8": " wants",
      "tOffsetMs": 420,
      "acAsrConf": 148
    }, {
      "utf8": " to",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " learn",
      "tOffsetMs": 990,
      "acAsrConf": 168
    }, {
      "utf8": " your",
      "tOffsetMs": 1170,
      "acAsrConf": 219
    }, {
      "utf8": " native",
      "tOffsetMs": 2130,
      "acAsrConf": 213
    } ]
  }, {
    "tStartMs": 1012910,
    "dDurationMs": 3549,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1012920,
    "dDurationMs": 5910,
    "wWinId": 1,
    "segs": [ {
      "utf8": "language",
      "acAsrConf": 252
    }, {
      "utf8": " now",
      "tOffsetMs": 859,
      "acAsrConf": 204
    }, {
      "utf8": " personally",
      "tOffsetMs": 1859,
      "acAsrConf": 235
    }, {
      "utf8": " I",
      "tOffsetMs": 2310,
      "acAsrConf": 201
    }, {
      "utf8": " haven't",
      "tOffsetMs": 2880,
      "acAsrConf": 252
    }, {
      "utf8": " had",
      "tOffsetMs": 3210,
      "acAsrConf": 207
    } ]
  }, {
    "tStartMs": 1016449,
    "dDurationMs": 2381,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1016459,
    "dDurationMs": 6931,
    "wWinId": 1,
    "segs": [ {
      "utf8": "much",
      "acAsrConf": 252
    }, {
      "utf8": " success",
      "tOffsetMs": 241,
      "acAsrConf": 252
    }, {
      "utf8": " with",
      "tOffsetMs": 541,
      "acAsrConf": 252
    }, {
      "utf8": " these",
      "tOffsetMs": 961,
      "acAsrConf": 201
    }, {
      "utf8": " websites",
      "tOffsetMs": 1261,
      "acAsrConf": 201
    }, {
      "utf8": " because",
      "tOffsetMs": 1801,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 1018820,
    "dDurationMs": 4570,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1018830,
    "dDurationMs": 7190,
    "wWinId": 1,
    "segs": [ {
      "utf8": "the",
      "acAsrConf": 245
    }, {
      "utf8": " time",
      "tOffsetMs": 720,
      "acAsrConf": 224
    }, {
      "utf8": " zone",
      "tOffsetMs": 1020,
      "acAsrConf": 224
    }, {
      "utf8": " didn't",
      "tOffsetMs": 1350,
      "acAsrConf": 252
    }, {
      "utf8": " work",
      "tOffsetMs": 1740,
      "acAsrConf": 252
    }, {
      "utf8": " or",
      "tOffsetMs": 1980,
      "acAsrConf": 255
    }, {
      "utf8": " we",
      "tOffsetMs": 2629,
      "acAsrConf": 236
    }, {
      "utf8": " I",
      "tOffsetMs": 3629,
      "acAsrConf": 236
    }, {
      "utf8": " came",
      "tOffsetMs": 4110,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1023380,
    "dDurationMs": 2640,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1023390,
    "dDurationMs": 5910,
    "wWinId": 1,
    "segs": [ {
      "utf8": "to",
      "acAsrConf": 252
    }, {
      "utf8": " Skype",
      "tOffsetMs": 240,
      "acAsrConf": 239
    }, {
      "utf8": " and",
      "tOffsetMs": 569,
      "acAsrConf": 233
    }, {
      "utf8": " they",
      "tOffsetMs": 870,
      "acAsrConf": 226
    }, {
      "utf8": " weren't",
      "tOffsetMs": 1049,
      "acAsrConf": 220
    }, {
      "utf8": " there",
      "tOffsetMs": 1319,
      "acAsrConf": 234
    }, {
      "utf8": " so",
      "tOffsetMs": 1620,
      "acAsrConf": 240
    }, {
      "utf8": " I",
      "tOffsetMs": 2370,
      "acAsrConf": 142
    } ]
  }, {
    "tStartMs": 1026010,
    "dDurationMs": 3290,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1026020,
    "dDurationMs": 6309,
    "wWinId": 1,
    "segs": [ {
      "utf8": "don't",
      "acAsrConf": 250
    }, {
      "utf8": " have",
      "tOffsetMs": 1000,
      "acAsrConf": 210
    }, {
      "utf8": " a",
      "tOffsetMs": 1330,
      "acAsrConf": 252
    }, {
      "utf8": " website",
      "tOffsetMs": 1360,
      "acAsrConf": 217
    }, {
      "utf8": " recommendation",
      "tOffsetMs": 1890,
      "acAsrConf": 248
    }, {
      "utf8": " for",
      "tOffsetMs": 2890,
      "acAsrConf": 241
    } ]
  }, {
    "tStartMs": 1029290,
    "dDurationMs": 3039,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1029300,
    "dDurationMs": 4649,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 120,
      "acAsrConf": 252
    }, {
      "utf8": " this",
      "tOffsetMs": 450,
      "acAsrConf": 252
    }, {
      "utf8": " but",
      "tOffsetMs": 600,
      "acAsrConf": 216
    }, {
      "utf8": " some",
      "tOffsetMs": 1250,
      "acAsrConf": 246
    }, {
      "utf8": " people",
      "tOffsetMs": 2250,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 2279,
      "acAsrConf": 143
    }, {
      "utf8": " a",
      "tOffsetMs": 2730,
      "acAsrConf": 219
    }, {
      "utf8": " lot",
      "tOffsetMs": 2789,
      "acAsrConf": 204
    } ]
  }, {
    "tStartMs": 1032319,
    "dDurationMs": 1630,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1032329,
    "dDurationMs": 3990,
    "wWinId": 1,
    "segs": [ {
      "utf8": "of",
      "acAsrConf": 252
    }, {
      "utf8": " success",
      "tOffsetMs": 31,
      "acAsrConf": 227
    }, {
      "utf8": " with",
      "tOffsetMs": 360,
      "acAsrConf": 148
    }, {
      "utf8": " this",
      "tOffsetMs": 721,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 870,
      "acAsrConf": 227
    }, {
      "utf8": " there's",
      "tOffsetMs": 1081,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 1291,
      "acAsrConf": 252
    }, {
      "utf8": " lot",
      "tOffsetMs": 1350,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 1591,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1033939,
    "dDurationMs": 2380,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1033949,
    "dDurationMs": 4921,
    "wWinId": 1,
    "segs": [ {
      "utf8": "language",
      "acAsrConf": 203
    }, {
      "utf8": " exchange",
      "tOffsetMs": 390,
      "acAsrConf": 217
    }, {
      "utf8": " websites",
      "tOffsetMs": 1021,
      "acAsrConf": 218
    }, {
      "utf8": " if",
      "tOffsetMs": 1500,
      "acAsrConf": 139
    }, {
      "utf8": " you",
      "tOffsetMs": 1771,
      "acAsrConf": 252
    }, {
      "utf8": " speak",
      "tOffsetMs": 2041,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1036309,
    "dDurationMs": 2561,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1036319,
    "dDurationMs": 4860,
    "wWinId": 1,
    "segs": [ {
      "utf8": "Korean",
      "acAsrConf": 230
    }, {
      "utf8": " and",
      "tOffsetMs": 510,
      "acAsrConf": 225
    }, {
      "utf8": " someone",
      "tOffsetMs": 870,
      "acAsrConf": 252
    }, {
      "utf8": " wants",
      "tOffsetMs": 1231,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1951,
      "acAsrConf": 252
    }, {
      "utf8": " learn",
      "tOffsetMs": 1981,
      "acAsrConf": 208
    }, {
      "utf8": " Korean",
      "tOffsetMs": 2250,
      "acAsrConf": 226
    } ]
  }, {
    "tStartMs": 1038860,
    "dDurationMs": 2319,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1038870,
    "dDurationMs": 5339,
    "wWinId": 1,
    "segs": [ {
      "utf8": "and",
      "acAsrConf": 250
    }, {
      "utf8": " you",
      "tOffsetMs": 540,
      "acAsrConf": 252
    }, {
      "utf8": " want",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 900,
      "acAsrConf": 252
    }, {
      "utf8": " learn",
      "tOffsetMs": 1020,
      "acAsrConf": 231
    }, {
      "utf8": " English",
      "tOffsetMs": 1170,
      "acAsrConf": 243
    }, {
      "utf8": " great",
      "tOffsetMs": 1380,
      "acAsrConf": 187
    }, {
      "utf8": " you",
      "tOffsetMs": 2040,
      "acAsrConf": 162
    } ]
  }, {
    "tStartMs": 1041169,
    "dDurationMs": 3040,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1041179,
    "dDurationMs": 6750,
    "wWinId": 1,
    "segs": [ {
      "utf8": "can",
      "acAsrConf": 239
    }, {
      "utf8": " exchange",
      "tOffsetMs": 241,
      "acAsrConf": 252
    }, {
      "utf8": " your",
      "tOffsetMs": 780,
      "acAsrConf": 252
    }, {
      "utf8": " languages",
      "tOffsetMs": 1171,
      "acAsrConf": 216
    }, {
      "utf8": " and",
      "tOffsetMs": 1801,
      "acAsrConf": 146
    }, {
      "utf8": " it",
      "tOffsetMs": 2071,
      "acAsrConf": 235
    }, {
      "utf8": " could",
      "tOffsetMs": 2821,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1044199,
    "dDurationMs": 3730,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1044209,
    "dDurationMs": 6720,
    "wWinId": 1,
    "segs": [ {
      "utf8": "work",
      "acAsrConf": 252
    }, {
      "utf8": " fine",
      "tOffsetMs": 151,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 451,
      "acAsrConf": 227
    }, {
      "utf8": " yes",
      "tOffsetMs": 720,
      "acAsrConf": 88
    }, {
      "utf8": " mera",
      "tOffsetMs": 1620,
      "acAsrConf": 248
    }, {
      "utf8": " kaam",
      "tOffsetMs": 2370,
      "acAsrConf": 248
    }, {
      "utf8": " language",
      "tOffsetMs": 2791,
      "acAsrConf": 238
    } ]
  }, {
    "tStartMs": 1047919,
    "dDurationMs": 3010,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1047929,
    "dDurationMs": 4801,
    "wWinId": 1,
    "segs": [ {
      "utf8": "exchange",
      "acAsrConf": 241
    }, {
      "utf8": " websites",
      "tOffsetMs": 541,
      "acAsrConf": 232
    }, {
      "utf8": " and",
      "tOffsetMs": 1081,
      "acAsrConf": 137
    }, {
      "utf8": " the",
      "tOffsetMs": 1521,
      "acAsrConf": 248
    }, {
      "utf8": " third",
      "tOffsetMs": 2521,
      "acAsrConf": 252
    }, {
      "utf8": " thing",
      "tOffsetMs": 2731,
      "acAsrConf": 217
    } ]
  }, {
    "tStartMs": 1050919,
    "dDurationMs": 1811,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1050929,
    "dDurationMs": 5311,
    "wWinId": 1,
    "segs": [ {
      "utf8": "what",
      "acAsrConf": 246
    }, {
      "utf8": " do",
      "tOffsetMs": 661,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 781,
      "acAsrConf": 252
    }, {
      "utf8": " think",
      "tOffsetMs": 871,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 1051,
      "acAsrConf": 252
    }, {
      "utf8": " third",
      "tOffsetMs": 1171,
      "acAsrConf": 252
    }, {
      "utf8": " thing",
      "tOffsetMs": 1351,
      "acAsrConf": 167
    } ]
  }, {
    "tStartMs": 1052720,
    "dDurationMs": 3520,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1052730,
    "dDurationMs": 9569,
    "wWinId": 1,
    "segs": [ {
      "utf8": "dan",
      "acAsrConf": 249
    }, {
      "utf8": " I",
      "tOffsetMs": 300,
      "acAsrConf": 217
    }, {
      "utf8": " would",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " say",
      "tOffsetMs": 1410,
      "acAsrConf": 207
    }, {
      "utf8": " meeting",
      "tOffsetMs": 1829,
      "acAsrConf": 232
    }, {
      "utf8": " a",
      "tOffsetMs": 2460,
      "acAsrConf": 150
    }, {
      "utf8": " wonderful",
      "tOffsetMs": 2730,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1056230,
    "dDurationMs": 6069,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1056240,
    "dDurationMs": 7980,
    "wWinId": 1,
    "segs": [ {
      "utf8": "tutor",
      "acAsrConf": 246
    }, {
      "utf8": " to",
      "tOffsetMs": 1189,
      "acAsrConf": 239
    }, {
      "utf8": " give",
      "tOffsetMs": 2189,
      "acAsrConf": 251
    }, {
      "utf8": " you",
      "tOffsetMs": 3059,
      "acAsrConf": 252
    }, {
      "utf8": " feedback",
      "tOffsetMs": 3240,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 3569,
      "acAsrConf": 217
    }, {
      "utf8": " speak",
      "tOffsetMs": 5059,
      "acAsrConf": 206
    } ]
  }, {
    "tStartMs": 1062289,
    "dDurationMs": 1931,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1062299,
    "dDurationMs": 4531,
    "wWinId": 1,
    "segs": [ {
      "utf8": "English",
      "acAsrConf": 252
    }, {
      "utf8": " with",
      "tOffsetMs": 241,
      "acAsrConf": 199
    }, {
      "utf8": " a",
      "tOffsetMs": 691,
      "acAsrConf": 252
    }, {
      "utf8": " native",
      "tOffsetMs": 721,
      "acAsrConf": 252
    }, {
      "utf8": " speaker",
      "tOffsetMs": 931,
      "acAsrConf": 150
    }, {
      "utf8": " would",
      "tOffsetMs": 1620,
      "acAsrConf": 252
    }, {
      "utf8": " be",
      "tOffsetMs": 1651,
      "acAsrConf": 216
    } ]
  }, {
    "tStartMs": 1064210,
    "dDurationMs": 2620,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1064220,
    "dDurationMs": 5910,
    "wWinId": 1,
    "segs": [ {
      "utf8": "very",
      "acAsrConf": 252
    }, {
      "utf8": " wise",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " sure",
      "tOffsetMs": 839,
      "acAsrConf": 233
    }, {
      "utf8": " so",
      "tOffsetMs": 1410,
      "acAsrConf": 252
    }, {
      "utf8": " meeting",
      "tOffsetMs": 1740,
      "acAsrConf": 252
    }, {
      "utf8": " someone",
      "tOffsetMs": 2040,
      "acAsrConf": 236
    } ]
  }, {
    "tStartMs": 1066820,
    "dDurationMs": 3310,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1066830,
    "dDurationMs": 5550,
    "wWinId": 1,
    "segs": [ {
      "utf8": "one-on-one",
      "acAsrConf": 238
    }, {
      "utf8": " on",
      "tOffsetMs": 450,
      "acAsrConf": 202
    }, {
      "utf8": " skype",
      "tOffsetMs": 1080,
      "acAsrConf": 209
    }, {
      "utf8": " in",
      "tOffsetMs": 1500,
      "acAsrConf": 237
    }, {
      "utf8": " person",
      "tOffsetMs": 1950,
      "acAsrConf": 252
    }, {
      "utf8": " it's",
      "tOffsetMs": 2520,
      "acAsrConf": 223
    }, {
      "utf8": " so",
      "tOffsetMs": 2820,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1070120,
    "dDurationMs": 2260,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1070130,
    "dDurationMs": 6299,
    "wWinId": 1,
    "segs": [ {
      "utf8": "valuable",
      "acAsrConf": 201
    }, {
      "utf8": " to",
      "tOffsetMs": 570,
      "acAsrConf": 252
    }, {
      "utf8": " be",
      "tOffsetMs": 900,
      "acAsrConf": 252
    }, {
      "utf8": " able",
      "tOffsetMs": 1049,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1080,
      "acAsrConf": 145
    }, {
      "utf8": " have",
      "tOffsetMs": 1410,
      "acAsrConf": 187
    }, {
      "utf8": " like",
      "tOffsetMs": 1710,
      "acAsrConf": 230
    }, {
      "utf8": " Dan",
      "tOffsetMs": 2010,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1072370,
    "dDurationMs": 4059,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1072380,
    "dDurationMs": 6929,
    "wWinId": 1,
    "segs": [ {
      "utf8": "said",
      "acAsrConf": 252
    }, {
      "utf8": " feedback",
      "tOffsetMs": 620,
      "acAsrConf": 207
    }, {
      "utf8": " now",
      "tOffsetMs": 1620,
      "acAsrConf": 249
    }, {
      "utf8": " can",
      "tOffsetMs": 2310,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 2669,
      "acAsrConf": 221
    }, {
      "utf8": " explain",
      "tOffsetMs": 2820,
      "acAsrConf": 252
    }, {
      "utf8": " what",
      "tOffsetMs": 3120,
      "acAsrConf": 204
    } ]
  }, {
    "tStartMs": 1076419,
    "dDurationMs": 2890,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1076429,
    "dDurationMs": 5791,
    "wWinId": 1,
    "segs": [ {
      "utf8": "is",
      "acAsrConf": 236
    }, {
      "utf8": " feedback",
      "tOffsetMs": 301,
      "acAsrConf": 238
    }, {
      "utf8": " maybe",
      "tOffsetMs": 870,
      "acAsrConf": 180
    }, {
      "utf8": " some",
      "tOffsetMs": 1681,
      "acAsrConf": 225
    }, {
      "utf8": " of",
      "tOffsetMs": 2041,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 2161,
      "acAsrConf": 249
    }, {
      "utf8": " know",
      "tOffsetMs": 2311,
      "acAsrConf": 172
    }, {
      "utf8": " this",
      "tOffsetMs": 2551,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1079299,
    "dDurationMs": 2921,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1079309,
    "dDurationMs": 6000,
    "wWinId": 1,
    "segs": [ {
      "utf8": "word",
      "acAsrConf": 239
    }, {
      "utf8": " maybe",
      "tOffsetMs": 271,
      "acAsrConf": 231
    }, {
      "utf8": " some",
      "tOffsetMs": 541,
      "acAsrConf": 252
    }, {
      "utf8": " don'ts",
      "tOffsetMs": 811,
      "acAsrConf": 166
    }, {
      "utf8": " yes",
      "tOffsetMs": 1201,
      "acAsrConf": 201
    }, {
      "utf8": " feedback",
      "tOffsetMs": 1681,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 2671,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1082210,
    "dDurationMs": 3099,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1082220,
    "dDurationMs": 6240,
    "wWinId": 1,
    "segs": [ {
      "utf8": "when",
      "acAsrConf": 252
    }, {
      "utf8": " somebody",
      "tOffsetMs": 680,
      "acAsrConf": 252
    }, {
      "utf8": " has",
      "tOffsetMs": 1680,
      "acAsrConf": 246
    }, {
      "utf8": " more",
      "tOffsetMs": 2040,
      "acAsrConf": 252
    }, {
      "utf8": " knowledge",
      "tOffsetMs": 2640,
      "acAsrConf": 210
    }, {
      "utf8": " than",
      "tOffsetMs": 3060,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1085299,
    "dDurationMs": 3161,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1085309,
    "dDurationMs": 5701,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 206
    }, {
      "utf8": " so",
      "tOffsetMs": 271,
      "acAsrConf": 210
    }, {
      "utf8": " is",
      "tOffsetMs": 1051,
      "acAsrConf": 251
    }, {
      "utf8": " better",
      "tOffsetMs": 1201,
      "acAsrConf": 250
    }, {
      "utf8": " at",
      "tOffsetMs": 1801,
      "acAsrConf": 225
    }, {
      "utf8": " speaking",
      "tOffsetMs": 2101,
      "acAsrConf": 252
    }, {
      "utf8": " English",
      "tOffsetMs": 2370,
      "acAsrConf": 154
    }, {
      "utf8": " and",
      "tOffsetMs": 2911,
      "acAsrConf": 234
    } ]
  }, {
    "tStartMs": 1088450,
    "dDurationMs": 2560,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1088460,
    "dDurationMs": 5250,
    "wWinId": 1,
    "segs": [ {
      "utf8": "they",
      "acAsrConf": 238
    }, {
      "utf8": " can",
      "tOffsetMs": 420,
      "acAsrConf": 252
    }, {
      "utf8": " listen",
      "tOffsetMs": 599,
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 810,
      "acAsrConf": 201
    }, {
      "utf8": " your",
      "tOffsetMs": 1290,
      "acAsrConf": 200
    }, {
      "utf8": " mistakes",
      "tOffsetMs": 1440,
      "acAsrConf": 252
    }, {
      "utf8": " or",
      "tOffsetMs": 2099,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1091000,
    "dDurationMs": 2710,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1091010,
    "dDurationMs": 6330,
    "wWinId": 1,
    "segs": [ {
      "utf8": "problems",
      "acAsrConf": 234
    }, {
      "utf8": " or",
      "tOffsetMs": 810,
      "acAsrConf": 204
    }, {
      "utf8": " just",
      "tOffsetMs": 1169,
      "acAsrConf": 232
    }, {
      "utf8": " give",
      "tOffsetMs": 1680,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1890,
      "acAsrConf": 252
    }, {
      "utf8": " any",
      "tOffsetMs": 2039,
      "acAsrConf": 252
    }, {
      "utf8": " general",
      "tOffsetMs": 2250,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1093700,
    "dDurationMs": 3640,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1093710,
    "dDurationMs": 6180,
    "wWinId": 1,
    "segs": [ {
      "utf8": "tips",
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 300,
      "acAsrConf": 229
    }, {
      "utf8": " then",
      "tOffsetMs": 630,
      "acAsrConf": 252
    }, {
      "utf8": " they",
      "tOffsetMs": 1290,
      "acAsrConf": 243
    }, {
      "utf8": " give",
      "tOffsetMs": 1800,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 2040,
      "acAsrConf": 252
    }, {
      "utf8": " feedback",
      "tOffsetMs": 2280,
      "acAsrConf": 234
    }, {
      "utf8": " so",
      "tOffsetMs": 2910,
      "acAsrConf": 163
    } ]
  }, {
    "tStartMs": 1097330,
    "dDurationMs": 2560,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1097340,
    "dDurationMs": 4079,
    "wWinId": 1,
    "segs": [ {
      "utf8": "this",
      "acAsrConf": 248
    }, {
      "utf8": " is",
      "tOffsetMs": 719,
      "acAsrConf": 252
    }, {
      "utf8": " what",
      "tOffsetMs": 959,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1170,
      "acAsrConf": 252
    }, {
      "utf8": " did",
      "tOffsetMs": 1320,
      "acAsrConf": 252
    }, {
      "utf8": " wrong",
      "tOffsetMs": 1500,
      "acAsrConf": 228
    }, {
      "utf8": " this",
      "tOffsetMs": 1770,
      "acAsrConf": 235
    }, {
      "utf8": " is",
      "tOffsetMs": 2190,
      "acAsrConf": 242
    }, {
      "utf8": " what",
      "tOffsetMs": 2400,
      "acAsrConf": 141
    } ]
  }, {
    "tStartMs": 1099880,
    "dDurationMs": 1539,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1099890,
    "dDurationMs": 7110,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 252
    }, {
      "utf8": " did",
      "tOffsetMs": 120,
      "acAsrConf": 252
    }, {
      "utf8": " right",
      "tOffsetMs": 270,
      "acAsrConf": 211
    }, {
      "utf8": " good",
      "tOffsetMs": 570,
      "acAsrConf": 213
    }, {
      "utf8": " job",
      "tOffsetMs": 1230,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1101409,
    "dDurationMs": 5591,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1101419,
    "dDurationMs": 7741,
    "wWinId": 1,
    "segs": [ {
      "utf8": "bad",
      "acAsrConf": 252
    }, {
      "utf8": " job",
      "tOffsetMs": 390,
      "acAsrConf": 255
    }, {
      "utf8": " it's",
      "tOffsetMs": 921,
      "acAsrConf": 226
    }, {
      "utf8": " not",
      "tOffsetMs": 1921,
      "acAsrConf": 187
    }, {
      "utf8": " bad",
      "tOffsetMs": 2101,
      "acAsrConf": 236
    }, {
      "utf8": " job",
      "tOffsetMs": 2491,
      "acAsrConf": 216
    }, {
      "utf8": " attitude",
      "tOffsetMs": 3411,
      "acAsrConf": 70
    }, {
      "utf8": " yes",
      "tOffsetMs": 4581,
      "acAsrConf": 226
    } ]
  }, {
    "tStartMs": 1106990,
    "dDurationMs": 2170,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1107000,
    "dDurationMs": 5820,
    "wWinId": 1,
    "segs": [ {
      "utf8": "yes",
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 270,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 450,
      "acAsrConf": 238
    }, {
      "utf8": " think",
      "tOffsetMs": 480,
      "acAsrConf": 159
    }, {
      "utf8": " one",
      "tOffsetMs": 809,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 1080,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 1110,
      "acAsrConf": 242
    }, {
      "utf8": " most",
      "tOffsetMs": 1320,
      "acAsrConf": 208
    }, {
      "utf8": " important",
      "tOffsetMs": 1530,
      "acAsrConf": 249
    } ]
  }, {
    "tStartMs": 1109150,
    "dDurationMs": 3670,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1109160,
    "dDurationMs": 6980,
    "wWinId": 1,
    "segs": [ {
      "utf8": "things",
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 240,
      "acAsrConf": 74
    }, {
      "utf8": " learning",
      "tOffsetMs": 1550,
      "acAsrConf": 250
    }, {
      "utf8": " English",
      "tOffsetMs": 2550,
      "acAsrConf": 243
    }, {
      "utf8": " is",
      "tOffsetMs": 2970,
      "acAsrConf": 252
    }, {
      "utf8": " giving",
      "tOffsetMs": 3180,
      "acAsrConf": 187
    } ]
  }, {
    "tStartMs": 1112810,
    "dDurationMs": 3330,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1112820,
    "dDurationMs": 6960,
    "wWinId": 1,
    "segs": [ {
      "utf8": "personal",
      "acAsrConf": 230
    }, {
      "utf8": " feedback",
      "tOffsetMs": 810,
      "acAsrConf": 247
    }, {
      "utf8": " so",
      "tOffsetMs": 1170,
      "acAsrConf": 167
    }, {
      "utf8": " that",
      "tOffsetMs": 1920,
      "acAsrConf": 201
    }, {
      "utf8": " could",
      "tOffsetMs": 2280,
      "acAsrConf": 249
    }, {
      "utf8": " be",
      "tOffsetMs": 2609,
      "acAsrConf": 157
    }, {
      "utf8": " a",
      "tOffsetMs": 2849,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1116130,
    "dDurationMs": 3650,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1116140,
    "dDurationMs": 6400,
    "wWinId": 1,
    "segs": [ {
      "utf8": "teacher",
      "acAsrConf": 249
    }, {
      "utf8": " or",
      "tOffsetMs": 1000,
      "acAsrConf": 252
    }, {
      "utf8": " someone",
      "tOffsetMs": 1600,
      "acAsrConf": 252
    }, {
      "utf8": " saying",
      "tOffsetMs": 1960,
      "acAsrConf": 210
    }, {
      "utf8": " hey",
      "tOffsetMs": 2409,
      "acAsrConf": 248
    }, {
      "utf8": " I",
      "tOffsetMs": 3310,
      "acAsrConf": 217
    } ]
  }, {
    "tStartMs": 1119770,
    "dDurationMs": 2770,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1119780,
    "dDurationMs": 5460,
    "wWinId": 1,
    "segs": [ {
      "utf8": "understood",
      "acAsrConf": 252
    }, {
      "utf8": " what",
      "tOffsetMs": 750,
      "acAsrConf": 255
    }, {
      "utf8": " you",
      "tOffsetMs": 1019,
      "acAsrConf": 252
    }, {
      "utf8": " said",
      "tOffsetMs": 1050,
      "acAsrConf": 104
    }, {
      "utf8": " but",
      "tOffsetMs": 1170,
      "acAsrConf": 187
    }, {
      "utf8": " it's",
      "tOffsetMs": 1920,
      "acAsrConf": 250
    }, {
      "utf8": " not",
      "tOffsetMs": 2460,
      "acAsrConf": 240
    } ]
  }, {
    "tStartMs": 1122530,
    "dDurationMs": 2710,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1122540,
    "dDurationMs": 5750,
    "wWinId": 1,
    "segs": [ {
      "utf8": "very",
      "acAsrConf": 238
    }, {
      "utf8": " natural",
      "tOffsetMs": 269,
      "acAsrConf": 229
    }, {
      "utf8": " to",
      "tOffsetMs": 629,
      "acAsrConf": 200
    }, {
      "utf8": " say",
      "tOffsetMs": 1590,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 1800,
      "acAsrConf": 252
    }, {
      "utf8": " like",
      "tOffsetMs": 1950,
      "acAsrConf": 221
    }, {
      "utf8": " that",
      "tOffsetMs": 2129,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 2160,
      "acAsrConf": 130
    } ]
  }, {
    "tStartMs": 1125230,
    "dDurationMs": 3060,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1125240,
    "dDurationMs": 5730,
    "wWinId": 1,
    "segs": [ {
      "utf8": "should",
      "acAsrConf": 96
    }, {
      "utf8": " say",
      "tOffsetMs": 689,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 929,
      "acAsrConf": 216
    }, {
      "utf8": " like",
      "tOffsetMs": 1080,
      "acAsrConf": 249
    }, {
      "utf8": " this",
      "tOffsetMs": 1439,
      "acAsrConf": 129
    }, {
      "utf8": " instead",
      "tOffsetMs": 1740,
      "acAsrConf": 208
    }, {
      "utf8": " and",
      "tOffsetMs": 2280,
      "acAsrConf": 206
    } ]
  }, {
    "tStartMs": 1128280,
    "dDurationMs": 2690,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1128290,
    "dDurationMs": 5259,
    "wWinId": 1,
    "segs": [ {
      "utf8": "maybe",
      "acAsrConf": 224
    }, {
      "utf8": " some",
      "tOffsetMs": 1000,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 1360,
      "acAsrConf": 243
    }, {
      "utf8": " your",
      "tOffsetMs": 1480,
      "acAsrConf": 208
    }, {
      "utf8": " friends",
      "tOffsetMs": 1629,
      "acAsrConf": 200
    }, {
      "utf8": " if",
      "tOffsetMs": 2139,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 2259,
      "acAsrConf": 162
    }, {
      "utf8": " have",
      "tOffsetMs": 2530,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1130960,
    "dDurationMs": 2589,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1130970,
    "dDurationMs": 5640,
    "wWinId": 1,
    "segs": [ {
      "utf8": "English",
      "acAsrConf": 166
    }, {
      "utf8": " speaking",
      "tOffsetMs": 420,
      "acAsrConf": 200
    }, {
      "utf8": " friends",
      "tOffsetMs": 780,
      "acAsrConf": 218
    }, {
      "utf8": " maybe",
      "tOffsetMs": 1160,
      "acAsrConf": 200
    }, {
      "utf8": " they",
      "tOffsetMs": 2160,
      "acAsrConf": 221
    } ]
  }, {
    "tStartMs": 1133539,
    "dDurationMs": 3071,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1133549,
    "dDurationMs": 7171,
    "wWinId": 1,
    "segs": [ {
      "utf8": "don't",
      "acAsrConf": 252
    }, {
      "utf8": " do",
      "tOffsetMs": 541,
      "acAsrConf": 241
    }, {
      "utf8": " that",
      "tOffsetMs": 1231,
      "acAsrConf": 226
    }, {
      "utf8": " with",
      "tOffsetMs": 1471,
      "acAsrConf": 207
    }, {
      "utf8": " you",
      "tOffsetMs": 1741,
      "acAsrConf": 252
    }, {
      "utf8": " they",
      "tOffsetMs": 1981,
      "acAsrConf": 249
    }, {
      "utf8": " just",
      "tOffsetMs": 2401,
      "acAsrConf": 65
    }, {
      "utf8": " say",
      "tOffsetMs": 2731,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1136600,
    "dDurationMs": 4120,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1136610,
    "dDurationMs": 5790,
    "wWinId": 1,
    "segs": [ {
      "utf8": "uh-huh",
      "acAsrConf": 166
    }, {
      "utf8": " oh",
      "tOffsetMs": 840,
      "acAsrConf": 206
    }, {
      "utf8": " really",
      "tOffsetMs": 1170,
      "acAsrConf": 246
    }, {
      "utf8": " oh",
      "tOffsetMs": 1980,
      "acAsrConf": 167
    }, {
      "utf8": " oh",
      "tOffsetMs": 2309,
      "acAsrConf": 255
    }, {
      "utf8": " because",
      "tOffsetMs": 2880,
      "acAsrConf": 249
    }, {
      "utf8": " they",
      "tOffsetMs": 3720,
      "acAsrConf": 200
    } ]
  }, {
    "tStartMs": 1140710,
    "dDurationMs": 1690,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1140720,
    "dDurationMs": 3810,
    "wWinId": 1,
    "segs": [ {
      "utf8": "don't",
      "acAsrConf": 234
    }, {
      "utf8": " want",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 449,
      "acAsrConf": 214
    }, {
      "utf8": " stop",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 900,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 930,
      "acAsrConf": 186
    }, {
      "utf8": " might",
      "tOffsetMs": 1320,
      "acAsrConf": 252
    }, {
      "utf8": " feel",
      "tOffsetMs": 1470,
      "acAsrConf": 217
    } ]
  }, {
    "tStartMs": 1142390,
    "dDurationMs": 2140,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1142400,
    "dDurationMs": 5010,
    "wWinId": 1,
    "segs": [ {
      "utf8": "uncomfortable",
      "acAsrConf": 252
    }, {
      "utf8": " but",
      "tOffsetMs": 50,
      "acAsrConf": 187
    }, {
      "utf8": " it's",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1260,
      "acAsrConf": 252
    }, {
      "utf8": " valuable",
      "tOffsetMs": 1500,
      "acAsrConf": 167
    }, {
      "utf8": " to",
      "tOffsetMs": 1860,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1144520,
    "dDurationMs": 2890,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1144530,
    "dDurationMs": 5070,
    "wWinId": 1,
    "segs": [ {
      "utf8": "get",
      "acAsrConf": 252
    }, {
      "utf8": " feedback",
      "tOffsetMs": 120,
      "acAsrConf": 226
    }, {
      "utf8": " so",
      "tOffsetMs": 360,
      "acAsrConf": 200
    }, {
      "utf8": " that",
      "tOffsetMs": 1110,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1139,
      "acAsrConf": 205
    }, {
      "utf8": " can",
      "tOffsetMs": 1470,
      "acAsrConf": 252
    }, {
      "utf8": " know",
      "tOffsetMs": 1680,
      "acAsrConf": 252
    }, {
      "utf8": " where",
      "tOffsetMs": 1889,
      "acAsrConf": 249
    } ]
  }, {
    "tStartMs": 1147400,
    "dDurationMs": 2200,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1147410,
    "dDurationMs": 5399,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 252
    }, {
      "utf8": " should",
      "tOffsetMs": 300,
      "acAsrConf": 213
    }, {
      "utf8": " improve",
      "tOffsetMs": 540,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 870,
      "acAsrConf": 236
    }, {
      "utf8": " feedback",
      "tOffsetMs": 1350,
      "acAsrConf": 205
    }, {
      "utf8": " is",
      "tOffsetMs": 1889,
      "acAsrConf": 213
    }, {
      "utf8": " very",
      "tOffsetMs": 2070,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1149590,
    "dDurationMs": 3219,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1149600,
    "dDurationMs": 6510,
    "wWinId": 1,
    "segs": [ {
      "utf8": "important",
      "acAsrConf": 248
    }, {
      "utf8": " and",
      "tOffsetMs": 480,
      "acAsrConf": 200
    }, {
      "utf8": " let's",
      "tOffsetMs": 829,
      "acAsrConf": 109
    }, {
      "utf8": " go",
      "tOffsetMs": 1829,
      "acAsrConf": 208
    }, {
      "utf8": " to",
      "tOffsetMs": 2370,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 2790,
      "acAsrConf": 252
    }, {
      "utf8": " next",
      "tOffsetMs": 3030,
      "acAsrConf": 209
    } ]
  }, {
    "tStartMs": 1152799,
    "dDurationMs": 3311,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1152809,
    "dDurationMs": 7201,
    "wWinId": 1,
    "segs": [ {
      "utf8": "points",
      "acAsrConf": 136
    }, {
      "utf8": " yes",
      "tOffsetMs": 481,
      "acAsrConf": 210
    }, {
      "utf8": " yes",
      "tOffsetMs": 811,
      "acAsrConf": 236
    }, {
      "utf8": " the",
      "tOffsetMs": 1201,
      "acAsrConf": 255
    }, {
      "utf8": " next",
      "tOffsetMs": 1441,
      "acAsrConf": 200
    }, {
      "utf8": " point",
      "tOffsetMs": 1831,
      "acAsrConf": 245
    }, {
      "utf8": " are",
      "tOffsetMs": 2101,
      "acAsrConf": 185
    }, {
      "utf8": " some",
      "tOffsetMs": 2521,
      "acAsrConf": 247
    } ]
  }, {
    "tStartMs": 1156100,
    "dDurationMs": 3910,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1156110,
    "dDurationMs": 5250,
    "wWinId": 1,
    "segs": [ {
      "utf8": "of",
      "acAsrConf": 205
    }, {
      "utf8": " my",
      "tOffsetMs": 270,
      "acAsrConf": 252
    }, {
      "utf8": " best",
      "tOffsetMs": 300,
      "acAsrConf": 216
    }, {
      "utf8": " conversation",
      "tOffsetMs": 1400,
      "acAsrConf": 205
    }, {
      "utf8": " tips",
      "tOffsetMs": 2400,
      "acAsrConf": 239
    }, {
      "utf8": " what",
      "tOffsetMs": 2790,
      "acAsrConf": 240
    }, {
      "utf8": " works",
      "tOffsetMs": 3630,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1160000,
    "dDurationMs": 1360,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1160010,
    "dDurationMs": 4230,
    "wWinId": 1,
    "segs": [ {
      "utf8": "and",
      "acAsrConf": 252
    }, {
      "utf8": " what",
      "tOffsetMs": 299,
      "acAsrConf": 237
    }, {
      "utf8": " doesn't",
      "tOffsetMs": 810,
      "acAsrConf": 230
    }, {
      "utf8": " work",
      "tOffsetMs": 1020,
      "acAsrConf": 77
    } ]
  }, {
    "tStartMs": 1161350,
    "dDurationMs": 2890,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1161360,
    "dDurationMs": 6540,
    "wWinId": 1,
    "segs": [ {
      "utf8": "and",
      "acAsrConf": 226
    }, {
      "utf8": " from",
      "tOffsetMs": 439,
      "acAsrConf": 252
    }, {
      "utf8": " my",
      "tOffsetMs": 1439,
      "acAsrConf": 252
    }, {
      "utf8": " experience",
      "tOffsetMs": 1620,
      "acAsrConf": 252
    }, {
      "utf8": " learning",
      "tOffsetMs": 1650,
      "acAsrConf": 226
    }, {
      "utf8": " French",
      "tOffsetMs": 2580,
      "acAsrConf": 247
    } ]
  }, {
    "tStartMs": 1164230,
    "dDurationMs": 3670,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1164240,
    "dDurationMs": 5760,
    "wWinId": 1,
    "segs": [ {
      "utf8": "and",
      "acAsrConf": 206
    }, {
      "utf8": " my",
      "tOffsetMs": 240,
      "acAsrConf": 228
    }, {
      "utf8": " not-so-great",
      "tOffsetMs": 1370,
      "acAsrConf": 0
    }, {
      "utf8": " experience",
      "tOffsetMs": 2370,
      "acAsrConf": 71
    }, {
      "utf8": " trying",
      "tOffsetMs": 3090,
      "acAsrConf": 230
    }, {
      "utf8": " to",
      "tOffsetMs": 3450,
      "acAsrConf": 202
    } ]
  }, {
    "tStartMs": 1167890,
    "dDurationMs": 2110,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1167900,
    "dDurationMs": 6420,
    "wWinId": 1,
    "segs": [ {
      "utf8": "learn",
      "acAsrConf": 252
    }, {
      "utf8": " Spanish",
      "tOffsetMs": 120,
      "acAsrConf": 236
    }, {
      "utf8": " I",
      "tOffsetMs": 420,
      "acAsrConf": 65
    }, {
      "utf8": " think",
      "tOffsetMs": 960,
      "acAsrConf": 235
    }, {
      "utf8": " we",
      "tOffsetMs": 1409,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 1529,
      "acAsrConf": 252
    }, {
      "utf8": " some",
      "tOffsetMs": 1680,
      "acAsrConf": 255
    }, {
      "utf8": " good",
      "tOffsetMs": 1890,
      "acAsrConf": 204
    } ]
  }, {
    "tStartMs": 1169990,
    "dDurationMs": 4330,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1170000,
    "dDurationMs": 7940,
    "wWinId": 1,
    "segs": [ {
      "utf8": "tips",
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 150,
      "acAsrConf": 166
    }, {
      "utf8": " you",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 630,
      "acAsrConf": 210
    }, {
      "utf8": " number",
      "tOffsetMs": 1230,
      "acAsrConf": 255
    }, {
      "utf8": " one",
      "tOffsetMs": 1950,
      "acAsrConf": 200
    }, {
      "utf8": " is",
      "tOffsetMs": 2130,
      "acAsrConf": 244
    }, {
      "utf8": " if",
      "tOffsetMs": 2400,
      "acAsrConf": 251
    }, {
      "utf8": " you",
      "tOffsetMs": 3330,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1174310,
    "dDurationMs": 3630,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1174320,
    "dDurationMs": 10020,
    "wWinId": 1,
    "segs": [ {
      "utf8": "want",
      "acAsrConf": 214
    }, {
      "utf8": " to",
      "tOffsetMs": 239,
      "acAsrConf": 252
    }, {
      "utf8": " be",
      "tOffsetMs": 300,
      "acAsrConf": 200
    }, {
      "utf8": " a",
      "tOffsetMs": 570,
      "acAsrConf": 238
    }, {
      "utf8": " fluent",
      "tOffsetMs": 750,
      "acAsrConf": 255
    }, {
      "utf8": " native",
      "tOffsetMs": 1380,
      "acAsrConf": 203
    }, {
      "utf8": " speaker",
      "tOffsetMs": 2070,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 2670,
      "acAsrConf": 200
    } ]
  }, {
    "tStartMs": 1177930,
    "dDurationMs": 6410,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1177940,
    "dDurationMs": 8440,
    "wWinId": 1,
    "segs": [ {
      "utf8": "have",
      "acAsrConf": 72
    }, {
      "utf8": " to",
      "tOffsetMs": 1000,
      "acAsrConf": 222
    }, {
      "utf8": " listen",
      "tOffsetMs": 1359,
      "acAsrConf": 234
    }, {
      "utf8": " to",
      "tOffsetMs": 1750,
      "acAsrConf": 252
    }, {
      "utf8": " what",
      "tOffsetMs": 1780,
      "acAsrConf": 226
    }, {
      "utf8": " natural",
      "tOffsetMs": 3960,
      "acAsrConf": 240
    }, {
      "utf8": " fluent",
      "tOffsetMs": 5400,
      "acAsrConf": 239
    } ]
  }, {
    "tStartMs": 1184330,
    "dDurationMs": 2050,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1184340,
    "dDurationMs": 5430,
    "wWinId": 1,
    "segs": [ {
      "utf8": "conversation",
      "acAsrConf": 252
    }, {
      "utf8": " yes",
      "tOffsetMs": 360,
      "acAsrConf": 143
    }, {
      "utf8": " just",
      "tOffsetMs": 1140,
      "acAsrConf": 236
    }, {
      "utf8": " like",
      "tOffsetMs": 1620,
      "acAsrConf": 247
    } ]
  }, {
    "tStartMs": 1186370,
    "dDurationMs": 3400,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1186380,
    "dDurationMs": 6600,
    "wWinId": 1,
    "segs": [ {
      "utf8": "just",
      "acAsrConf": 237
    }, {
      "utf8": " exactly",
      "tOffsetMs": 540,
      "acAsrConf": 255
    }, {
      "utf8": " like",
      "tOffsetMs": 750,
      "acAsrConf": 154
    }, {
      "utf8": " this",
      "tOffsetMs": 1320,
      "acAsrConf": 249
    }, {
      "utf8": " conversation",
      "tOffsetMs": 1500,
      "acAsrConf": 227
    }, {
      "utf8": " yes",
      "tOffsetMs": 2390,
      "acAsrConf": 143
    } ]
  }, {
    "tStartMs": 1189760,
    "dDurationMs": 3220,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1189770,
    "dDurationMs": 6480,
    "wWinId": 1,
    "segs": [ {
      "utf8": "yes",
      "acAsrConf": 208
    }, {
      "utf8": " so",
      "tOffsetMs": 510,
      "acAsrConf": 189
    }, {
      "utf8": " you",
      "tOffsetMs": 870,
      "acAsrConf": 248
    }, {
      "utf8": " need",
      "tOffsetMs": 1230,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1440,
      "acAsrConf": 252
    }, {
      "utf8": " listen",
      "tOffsetMs": 1590,
      "acAsrConf": 223
    }, {
      "utf8": " to",
      "tOffsetMs": 1800,
      "acAsrConf": 222
    }, {
      "utf8": " what",
      "tOffsetMs": 1980,
      "acAsrConf": 204
    }, {
      "utf8": " you",
      "tOffsetMs": 2400,
      "acAsrConf": 144
    } ]
  }, {
    "tStartMs": 1192970,
    "dDurationMs": 3280,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1192980,
    "dDurationMs": 5340,
    "wWinId": 1,
    "segs": [ {
      "utf8": "want",
      "acAsrConf": 85
    }, {
      "utf8": " to",
      "tOffsetMs": 300,
      "acAsrConf": 252
    }, {
      "utf8": " be",
      "tOffsetMs": 390,
      "acAsrConf": 200
    }, {
      "utf8": " so",
      "tOffsetMs": 690,
      "acAsrConf": 255
    }, {
      "utf8": " if",
      "tOffsetMs": 1020,
      "acAsrConf": 255
    }, {
      "utf8": " you",
      "tOffsetMs": 1200,
      "acAsrConf": 207
    }, {
      "utf8": " want",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1740,
      "acAsrConf": 219
    }, {
      "utf8": " be",
      "tOffsetMs": 1950,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 2490,
      "acAsrConf": 249
    }, {
      "utf8": " fluent",
      "tOffsetMs": 2760,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1196240,
    "dDurationMs": 2080,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1196250,
    "dDurationMs": 6240,
    "wWinId": 1,
    "segs": [ {
      "utf8": "speaker",
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 30,
      "acAsrConf": 187
    }, {
      "utf8": " need",
      "tOffsetMs": 570,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 780,
      "acAsrConf": 252
    }, {
      "utf8": " listen",
      "tOffsetMs": 930,
      "acAsrConf": 234
    }, {
      "utf8": " to",
      "tOffsetMs": 1110,
      "acAsrConf": 145
    }, {
      "utf8": " fluent",
      "tOffsetMs": 1410,
      "acAsrConf": 240
    } ]
  }, {
    "tStartMs": 1198310,
    "dDurationMs": 4180,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1198320,
    "dDurationMs": 6600,
    "wWinId": 1,
    "segs": [ {
      "utf8": "speaking",
      "acAsrConf": 227
    }, {
      "utf8": " but",
      "tOffsetMs": 600,
      "acAsrConf": 240
    }, {
      "utf8": " a",
      "tOffsetMs": 1230,
      "acAsrConf": 241
    }, {
      "utf8": " lot",
      "tOffsetMs": 2330,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 3330,
      "acAsrConf": 252
    }, {
      "utf8": " people",
      "tOffsetMs": 3360,
      "acAsrConf": 200
    }, {
      "utf8": " think",
      "tOffsetMs": 3810,
      "acAsrConf": 252
    }, {
      "utf8": " oh",
      "tOffsetMs": 4050,
      "acAsrConf": 245
    } ]
  }, {
    "tStartMs": 1202480,
    "dDurationMs": 2440,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1202490,
    "dDurationMs": 5160,
    "wWinId": 1,
    "segs": [ {
      "utf8": "I'll",
      "acAsrConf": 239
    }, {
      "utf8": " read",
      "tOffsetMs": 270,
      "acAsrConf": 252
    }, {
      "utf8": " some",
      "tOffsetMs": 630,
      "acAsrConf": 252
    }, {
      "utf8": " news",
      "tOffsetMs": 930,
      "acAsrConf": 252
    }, {
      "utf8": " articles",
      "tOffsetMs": 1110,
      "acAsrConf": 215
    }, {
      "utf8": " or",
      "tOffsetMs": 1350,
      "acAsrConf": 202
    }, {
      "utf8": " I'll",
      "tOffsetMs": 2010,
      "acAsrConf": 216
    } ]
  }, {
    "tStartMs": 1204910,
    "dDurationMs": 2740,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1204920,
    "dDurationMs": 4860,
    "wWinId": 1,
    "segs": [ {
      "utf8": "I'll",
      "acAsrConf": 241
    }, {
      "utf8": " read",
      "tOffsetMs": 630,
      "acAsrConf": 252
    }, {
      "utf8": " something",
      "tOffsetMs": 1380,
      "acAsrConf": 187
    }, {
      "utf8": " or",
      "tOffsetMs": 1890,
      "acAsrConf": 252
    }, {
      "utf8": " I'll",
      "tOffsetMs": 2040,
      "acAsrConf": 252
    }, {
      "utf8": " do",
      "tOffsetMs": 2190,
      "acAsrConf": 252
    }, {
      "utf8": " something",
      "tOffsetMs": 2340,
      "acAsrConf": 154
    } ]
  }, {
    "tStartMs": 1207640,
    "dDurationMs": 2140,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1207650,
    "dDurationMs": 5370,
    "wWinId": 1,
    "segs": [ {
      "utf8": "else",
      "acAsrConf": 204
    }, {
      "utf8": " that's",
      "tOffsetMs": 180,
      "acAsrConf": 252
    }, {
      "utf8": " not",
      "tOffsetMs": 360,
      "acAsrConf": 204
    }, {
      "utf8": " exactly",
      "tOffsetMs": 450,
      "acAsrConf": 203
    }, {
      "utf8": " what",
      "tOffsetMs": 870,
      "acAsrConf": 242
    }, {
      "utf8": " you",
      "tOffsetMs": 1290,
      "acAsrConf": 252
    }, {
      "utf8": " want",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 2040,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1209770,
    "dDurationMs": 3250,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1209780,
    "dDurationMs": 6360,
    "wWinId": 1,
    "segs": [ {
      "utf8": "do",
      "acAsrConf": 208
    }, {
      "utf8": " in",
      "tOffsetMs": 330,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 540,
      "acAsrConf": 252
    }, {
      "utf8": " future",
      "tOffsetMs": 660,
      "acAsrConf": 180
    }, {
      "utf8": " so",
      "tOffsetMs": 1020,
      "acAsrConf": 144
    }, {
      "utf8": " I",
      "tOffsetMs": 1320,
      "acAsrConf": 243
    }, {
      "utf8": " recommend",
      "tOffsetMs": 1350,
      "acAsrConf": 238
    }, {
      "utf8": " if",
      "tOffsetMs": 2040,
      "acAsrConf": 255
    }, {
      "utf8": " you",
      "tOffsetMs": 2250,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1213010,
    "dDurationMs": 3130,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1213020,
    "dDurationMs": 6300,
    "wWinId": 1,
    "segs": [ {
      "utf8": "want",
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 30,
      "acAsrConf": 167
    }, {
      "utf8": " be",
      "tOffsetMs": 360,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " fluent",
      "tOffsetMs": 540,
      "acAsrConf": 252
    }, {
      "utf8": " speaker",
      "tOffsetMs": 990,
      "acAsrConf": 252
    }, {
      "utf8": " listen",
      "tOffsetMs": 1610,
      "acAsrConf": 227
    }, {
      "utf8": " to",
      "tOffsetMs": 2610,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 1216130,
    "dDurationMs": 3190,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1216140,
    "dDurationMs": 6120,
    "wWinId": 1,
    "segs": [ {
      "utf8": "fluent",
      "acAsrConf": 237
    }, {
      "utf8": " English",
      "tOffsetMs": 780,
      "acAsrConf": 240
    }, {
      "utf8": " participate",
      "tOffsetMs": 1080,
      "acAsrConf": 198
    }, {
      "utf8": " in",
      "tOffsetMs": 2040,
      "acAsrConf": 227
    }, {
      "utf8": " fluent",
      "tOffsetMs": 2340,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1219310,
    "dDurationMs": 2950,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1219320,
    "dDurationMs": 4680,
    "wWinId": 1,
    "segs": [ {
      "utf8": "conversations",
      "acAsrConf": 246
    }, {
      "utf8": " like",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " meetup",
      "tOffsetMs": 1260,
      "acAsrConf": 249
    }, {
      "utf8": " comm",
      "tOffsetMs": 1770,
      "acAsrConf": 222
    }, {
      "utf8": " or",
      "tOffsetMs": 2160,
      "acAsrConf": 226
    }, {
      "utf8": " Skype",
      "tOffsetMs": 2430,
      "acAsrConf": 167
    } ]
  }, {
    "tStartMs": 1222250,
    "dDurationMs": 1750,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1222260,
    "dDurationMs": 5430,
    "wWinId": 1,
    "segs": [ {
      "utf8": "lessons",
      "acAsrConf": 188
    }, {
      "utf8": " something",
      "tOffsetMs": 480,
      "acAsrConf": 206
    }, {
      "utf8": " where",
      "tOffsetMs": 930,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1140,
      "acAsrConf": 252
    }, {
      "utf8": " will",
      "tOffsetMs": 1410,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 1590,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1223990,
    "dDurationMs": 3700,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1224000,
    "dDurationMs": 6090,
    "wWinId": 1,
    "segs": [ {
      "utf8": "the",
      "acAsrConf": 201
    }, {
      "utf8": " chance",
      "tOffsetMs": 210,
      "acAsrConf": 203
    }, {
      "utf8": " to",
      "tOffsetMs": 480,
      "acAsrConf": 207
    }, {
      "utf8": " experience",
      "tOffsetMs": 750,
      "acAsrConf": 201
    }, {
      "utf8": " what",
      "tOffsetMs": 1850,
      "acAsrConf": 132
    }, {
      "utf8": " you",
      "tOffsetMs": 2850,
      "acAsrConf": 247
    }, {
      "utf8": " want",
      "tOffsetMs": 3300,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1227680,
    "dDurationMs": 2410,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1227690,
    "dDurationMs": 5640,
    "wWinId": 1,
    "segs": [ {
      "utf8": "to",
      "acAsrConf": 252
    }, {
      "utf8": " be",
      "tOffsetMs": 210,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 390,
      "acAsrConf": 226
    }, {
      "utf8": " there's",
      "tOffsetMs": 660,
      "acAsrConf": 248
    }, {
      "utf8": " a",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    }, {
      "utf8": " lot",
      "tOffsetMs": 1620,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 1860,
      "acAsrConf": 252
    }, {
      "utf8": " great",
      "tOffsetMs": 1890,
      "acAsrConf": 192
    }, {
      "utf8": " ways",
      "tOffsetMs": 2250,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1230080,
    "dDurationMs": 3250,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1230090,
    "dDurationMs": 6420,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 236
    }, {
      "utf8": " can",
      "tOffsetMs": 210,
      "acAsrConf": 252
    }, {
      "utf8": " do",
      "tOffsetMs": 360,
      "acAsrConf": 252
    }, {
      "utf8": " this",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " YouTube",
      "tOffsetMs": 1310,
      "acAsrConf": 217
    }, {
      "utf8": " has",
      "tOffsetMs": 2310,
      "acAsrConf": 134
    }, {
      "utf8": " a",
      "tOffsetMs": 2610,
      "acAsrConf": 214
    }, {
      "utf8": " lot",
      "tOffsetMs": 2790,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 3060,
      "acAsrConf": 124
    } ]
  }, {
    "tStartMs": 1233320,
    "dDurationMs": 3190,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1233330,
    "dDurationMs": 5970,
    "wWinId": 1,
    "segs": [ {
      "utf8": "great",
      "acAsrConf": 236
    }, {
      "utf8": " videos",
      "tOffsetMs": 300,
      "acAsrConf": 204
    }, {
      "utf8": " that",
      "tOffsetMs": 720,
      "acAsrConf": 145
    }, {
      "utf8": " you",
      "tOffsetMs": 1650,
      "acAsrConf": 252
    }, {
      "utf8": " can",
      "tOffsetMs": 1830,
      "acAsrConf": 252
    }, {
      "utf8": " use",
      "tOffsetMs": 2010,
      "acAsrConf": 166
    }, {
      "utf8": " to",
      "tOffsetMs": 2280,
      "acAsrConf": 249
    }, {
      "utf8": " hear",
      "tOffsetMs": 2640,
      "acAsrConf": 247
    } ]
  }, {
    "tStartMs": 1236500,
    "dDurationMs": 2800,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1236510,
    "dDurationMs": 5190,
    "wWinId": 1,
    "segs": [ {
      "utf8": "natural",
      "acAsrConf": 252
    }, {
      "utf8": " conversations",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1620,
      "acAsrConf": 230
    }, {
      "utf8": " learn",
      "tOffsetMs": 2190,
      "acAsrConf": 232
    }, {
      "utf8": " those",
      "tOffsetMs": 2430,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1239290,
    "dDurationMs": 2410,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1239300,
    "dDurationMs": 6170,
    "wWinId": 1,
    "segs": [ {
      "utf8": "little",
      "acAsrConf": 252
    }, {
      "utf8": " tips",
      "tOffsetMs": 270,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 510,
      "acAsrConf": 204
    }, {
      "utf8": " tricks",
      "tOffsetMs": 810,
      "acAsrConf": 113
    }, {
      "utf8": " that",
      "tOffsetMs": 1140,
      "acAsrConf": 238
    }, {
      "utf8": " native",
      "tOffsetMs": 1770,
      "acAsrConf": 240
    } ]
  }, {
    "tStartMs": 1241690,
    "dDurationMs": 3780,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1241700,
    "dDurationMs": 6060,
    "wWinId": 1,
    "segs": [ {
      "utf8": "speakers",
      "acAsrConf": 124
    }, {
      "utf8": " use",
      "tOffsetMs": 420,
      "acAsrConf": 252
    }, {
      "utf8": " little",
      "tOffsetMs": 450,
      "acAsrConf": 221
    }, {
      "utf8": " dialogue",
      "tOffsetMs": 1050,
      "acAsrConf": 187
    }, {
      "utf8": " bits",
      "tOffsetMs": 1650,
      "acAsrConf": 235
    }, {
      "utf8": " and",
      "tOffsetMs": 2160,
      "acAsrConf": 217
    }, {
      "utf8": " I",
      "tOffsetMs": 2460,
      "acAsrConf": 246
    } ]
  }, {
    "tStartMs": 1245460,
    "dDurationMs": 2300,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1245470,
    "dDurationMs": 6400,
    "wWinId": 1,
    "segs": [ {
      "utf8": "think",
      "acAsrConf": 245
    }, {
      "utf8": " that",
      "tOffsetMs": 1000,
      "acAsrConf": 252
    }, {
      "utf8": " this",
      "tOffsetMs": 1180,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 1510,
      "acAsrConf": 236
    }, {
      "utf8": " a",
      "tOffsetMs": 1750,
      "acAsrConf": 252
    }, {
      "utf8": " great",
      "tOffsetMs": 1780,
      "acAsrConf": 239
    }, {
      "utf8": " resource",
      "tOffsetMs": 2080,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1247750,
    "dDurationMs": 4120,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1247760,
    "dDurationMs": 7950,
    "wWinId": 1,
    "segs": [ {
      "utf8": "because",
      "acAsrConf": 212
    }, {
      "utf8": " it's",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " free",
      "tOffsetMs": 1350,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1770,
      "acAsrConf": 237
    }, {
      "utf8": " of",
      "tOffsetMs": 2070,
      "acAsrConf": 248
    }, {
      "utf8": " course",
      "tOffsetMs": 2850,
      "acAsrConf": 252
    }, {
      "utf8": " there",
      "tOffsetMs": 2910,
      "acAsrConf": 96
    }, {
      "utf8": " is",
      "tOffsetMs": 3900,
      "acAsrConf": 199
    } ]
  }, {
    "tStartMs": 1251860,
    "dDurationMs": 3850,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1251870,
    "dDurationMs": 7650,
    "wWinId": 1,
    "segs": [ {
      "utf8": "some",
      "acAsrConf": 252
    }, {
      "utf8": " negative",
      "tOffsetMs": 270,
      "acAsrConf": 224
    }, {
      "utf8": " sides",
      "tOffsetMs": 710,
      "acAsrConf": 226
    }, {
      "utf8": " to",
      "tOffsetMs": 1710,
      "acAsrConf": 249
    }, {
      "utf8": " it",
      "tOffsetMs": 2130,
      "acAsrConf": 252
    }, {
      "utf8": " because",
      "tOffsetMs": 2310,
      "acAsrConf": 252
    }, {
      "utf8": " if",
      "tOffsetMs": 2790,
      "acAsrConf": 245
    }, {
      "utf8": " you",
      "tOffsetMs": 3150,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1255700,
    "dDurationMs": 3820,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1255710,
    "dDurationMs": 5760,
    "wWinId": 1,
    "segs": [ {
      "utf8": "understand",
      "acAsrConf": 252
    }, {
      "utf8": " less",
      "tOffsetMs": 720,
      "acAsrConf": 221
    }, {
      "utf8": " than",
      "tOffsetMs": 1020,
      "acAsrConf": 252
    }, {
      "utf8": " 70",
      "tOffsetMs": 1350,
      "acAsrConf": 143
    }, {
      "utf8": " percent",
      "tOffsetMs": 1830,
      "acAsrConf": 147
    }, {
      "utf8": " you're",
      "tOffsetMs": 2810,
      "acAsrConf": 169
    } ]
  }, {
    "tStartMs": 1259510,
    "dDurationMs": 1960,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1259520,
    "dDurationMs": 5040,
    "wWinId": 1,
    "segs": [ {
      "utf8": "not",
      "acAsrConf": 252
    }, {
      "utf8": " going",
      "tOffsetMs": 210,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 450,
      "acAsrConf": 230
    }, {
      "utf8": " be",
      "tOffsetMs": 570,
      "acAsrConf": 252
    }, {
      "utf8": " able",
      "tOffsetMs": 690,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 870,
      "acAsrConf": 201
    }, {
      "utf8": " always",
      "tOffsetMs": 1080,
      "acAsrConf": 151
    }, {
      "utf8": " have",
      "tOffsetMs": 1680,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1261460,
    "dDurationMs": 3100,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1261470,
    "dDurationMs": 5660,
    "wWinId": 1,
    "segs": [ {
      "utf8": "accurate",
      "acAsrConf": 244
    }, {
      "utf8": " subtitles",
      "tOffsetMs": 810,
      "acAsrConf": 233
    }, {
      "utf8": " so",
      "tOffsetMs": 1560,
      "acAsrConf": 231
    }, {
      "utf8": " YouTube",
      "tOffsetMs": 2070,
      "acAsrConf": 248
    }, {
      "utf8": " sometimes",
      "tOffsetMs": 2580,
      "acAsrConf": 213
    } ]
  }, {
    "tStartMs": 1264550,
    "dDurationMs": 2580,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1264560,
    "dDurationMs": 5430,
    "wWinId": 1,
    "segs": [ {
      "utf8": "has",
      "acAsrConf": 252
    }, {
      "utf8": " automatic",
      "tOffsetMs": 270,
      "acAsrConf": 255
    }, {
      "utf8": " subtitles",
      "tOffsetMs": 1170,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1440,
      "acAsrConf": 227
    } ]
  }, {
    "tStartMs": 1267120,
    "dDurationMs": 2870,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1267130,
    "dDurationMs": 5110,
    "wWinId": 1,
    "segs": [ {
      "utf8": "unfortunately",
      "acAsrConf": 235
    }, {
      "utf8": " it",
      "tOffsetMs": 1000,
      "acAsrConf": 229
    }, {
      "utf8": " might",
      "tOffsetMs": 1540,
      "acAsrConf": 248
    }, {
      "utf8": " be",
      "tOffsetMs": 2500,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 2650,
      "acAsrConf": 252
    }, {
      "utf8": " little",
      "tOffsetMs": 2680,
      "acAsrConf": 234
    } ]
  }, {
    "tStartMs": 1269980,
    "dDurationMs": 2260,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1269990,
    "dDurationMs": 4770,
    "wWinId": 1,
    "segs": [ {
      "utf8": "difficult",
      "acAsrConf": 160
    }, {
      "utf8": " to",
      "tOffsetMs": 570,
      "acAsrConf": 252
    }, {
      "utf8": " find",
      "tOffsetMs": 630,
      "acAsrConf": 164
    }, {
      "utf8": " accurate",
      "tOffsetMs": 930,
      "acAsrConf": 252
    }, {
      "utf8": " subtitles",
      "tOffsetMs": 1230,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 2070,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1272230,
    "dDurationMs": 2530,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1272240,
    "dDurationMs": 5460,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 252
    }, {
      "utf8": " need",
      "tOffsetMs": 30,
      "acAsrConf": 219
    }, {
      "utf8": " someone",
      "tOffsetMs": 420,
      "acAsrConf": 204
    }, {
      "utf8": " to",
      "tOffsetMs": 840,
      "acAsrConf": 230
    }, {
      "utf8": " explain",
      "tOffsetMs": 900,
      "acAsrConf": 128
    }, {
      "utf8": " or",
      "tOffsetMs": 1620,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 2070,
      "acAsrConf": 252
    }, {
      "utf8": " talk",
      "tOffsetMs": 2340,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1274750,
    "dDurationMs": 2950,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1274760,
    "dDurationMs": 6420,
    "wWinId": 1,
    "segs": [ {
      "utf8": "about",
      "acAsrConf": 232
    }, {
      "utf8": " expressions",
      "tOffsetMs": 210,
      "acAsrConf": 227
    }, {
      "utf8": " in",
      "tOffsetMs": 1140,
      "acAsrConf": 213
    }, {
      "utf8": " the",
      "tOffsetMs": 1320,
      "acAsrConf": 252
    }, {
      "utf8": " video",
      "tOffsetMs": 1620,
      "acAsrConf": 252
    }, {
      "utf8": " but",
      "tOffsetMs": 1890,
      "acAsrConf": 201
    }, {
      "utf8": " on",
      "tOffsetMs": 2580,
      "acAsrConf": 175
    } ]
  }, {
    "tStartMs": 1277690,
    "dDurationMs": 3490,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1277700,
    "dDurationMs": 6480,
    "wWinId": 1,
    "segs": [ {
      "utf8": "the",
      "acAsrConf": 241
    }, {
      "utf8": " mind",
      "tOffsetMs": 330,
      "acAsrConf": 226
    }, {
      "utf8": " map",
      "tOffsetMs": 630,
      "acAsrConf": 226
    }, {
      "utf8": " that",
      "tOffsetMs": 810,
      "acAsrConf": 141
    }, {
      "utf8": " I",
      "tOffsetMs": 1680,
      "acAsrConf": 248
    }, {
      "utf8": " sent",
      "tOffsetMs": 1710,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 2130,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 2250,
      "acAsrConf": 186
    }, {
      "utf8": " so",
      "tOffsetMs": 2520,
      "acAsrConf": 145
    }, {
      "utf8": " on",
      "tOffsetMs": 2760,
      "acAsrConf": 142
    }, {
      "utf8": " in",
      "tOffsetMs": 3270,
      "acAsrConf": 247
    } ]
  }, {
    "tStartMs": 1281170,
    "dDurationMs": 3010,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1281180,
    "dDurationMs": 5250,
    "wWinId": 1,
    "segs": [ {
      "utf8": "the",
      "acAsrConf": 252
    }, {
      "utf8": " document",
      "tOffsetMs": 180,
      "acAsrConf": 149
    }, {
      "utf8": " I",
      "tOffsetMs": 720,
      "acAsrConf": 107
    }, {
      "utf8": " sent",
      "tOffsetMs": 930,
      "acAsrConf": 252
    }, {
      "utf8": " three",
      "tOffsetMs": 1800,
      "acAsrConf": 206
    }, {
      "utf8": " YouTube",
      "tOffsetMs": 2250,
      "acAsrConf": 197
    } ]
  }, {
    "tStartMs": 1284170,
    "dDurationMs": 2260,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1284180,
    "dDurationMs": 5640,
    "wWinId": 1,
    "segs": [ {
      "utf8": "channels",
      "acAsrConf": 252
    }, {
      "utf8": " that",
      "tOffsetMs": 450,
      "acAsrConf": 173
    }, {
      "utf8": " I",
      "tOffsetMs": 690,
      "acAsrConf": 239
    }, {
      "utf8": " recommend",
      "tOffsetMs": 720,
      "acAsrConf": 234
    }, {
      "utf8": " now",
      "tOffsetMs": 1350,
      "acAsrConf": 224
    }, {
      "utf8": " these",
      "tOffsetMs": 2070,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1286420,
    "dDurationMs": 3400,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1286430,
    "dDurationMs": 7400,
    "wWinId": 1,
    "segs": [ {
      "utf8": "three",
      "acAsrConf": 218
    }, {
      "utf8": " YouTube",
      "tOffsetMs": 330,
      "acAsrConf": 165
    }, {
      "utf8": " channels",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " are",
      "tOffsetMs": 1110,
      "acAsrConf": 242
    }, {
      "utf8": " for",
      "tOffsetMs": 1520,
      "acAsrConf": 242
    }, {
      "utf8": " native",
      "tOffsetMs": 2520,
      "acAsrConf": 239
    } ]
  }, {
    "tStartMs": 1289810,
    "dDurationMs": 4020,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1289820,
    "dDurationMs": 7170,
    "wWinId": 1,
    "segs": [ {
      "utf8": "speakers",
      "acAsrConf": 232
    }, {
      "utf8": " so",
      "tOffsetMs": 780,
      "acAsrConf": 204
    }, {
      "utf8": " they",
      "tOffsetMs": 1620,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 2010,
      "acAsrConf": 252
    }, {
      "utf8": " very",
      "tOffsetMs": 2340,
      "acAsrConf": 243
    }, {
      "utf8": " natural",
      "tOffsetMs": 2960,
      "acAsrConf": 208
    } ]
  }, {
    "tStartMs": 1293820,
    "dDurationMs": 3170,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1293830,
    "dDurationMs": 6190,
    "wWinId": 1,
    "segs": [ {
      "utf8": "fluent",
      "acAsrConf": 242
    }, {
      "utf8": " conversation",
      "tOffsetMs": 1000,
      "acAsrConf": 252
    }, {
      "utf8": " they",
      "tOffsetMs": 1650,
      "acAsrConf": 201
    }, {
      "utf8": " have",
      "tOffsetMs": 2650,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 2830,
      "acAsrConf": 252
    }, {
      "utf8": " lot",
      "tOffsetMs": 2860,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 3130,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1296980,
    "dDurationMs": 3040,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1296990,
    "dDurationMs": 5640,
    "wWinId": 1,
    "segs": [ {
      "utf8": "great",
      "acAsrConf": 252
    }, {
      "utf8": " dialogues",
      "tOffsetMs": 330,
      "acAsrConf": 245
    }, {
      "utf8": " and",
      "tOffsetMs": 1020,
      "acAsrConf": 201
    }, {
      "utf8": " these",
      "tOffsetMs": 1410,
      "acAsrConf": 240
    }, {
      "utf8": " three",
      "tOffsetMs": 2100,
      "acAsrConf": 201
    }, {
      "utf8": " YouTube",
      "tOffsetMs": 2610,
      "acAsrConf": 202
    } ]
  }, {
    "tStartMs": 1300010,
    "dDurationMs": 2620,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1300020,
    "dDurationMs": 3930,
    "wWinId": 1,
    "segs": [ {
      "utf8": "channels",
      "acAsrConf": 252
    }, {
      "utf8": " on",
      "tOffsetMs": 480,
      "acAsrConf": 228
    }, {
      "utf8": " the",
      "tOffsetMs": 630,
      "acAsrConf": 252
    }, {
      "utf8": " mind",
      "tOffsetMs": 1020,
      "acAsrConf": 217
    }, {
      "utf8": " map",
      "tOffsetMs": 1260,
      "acAsrConf": 217
    }, {
      "utf8": " if",
      "tOffsetMs": 1410,
      "acAsrConf": 201
    }, {
      "utf8": " you",
      "tOffsetMs": 1830,
      "acAsrConf": 218
    }, {
      "utf8": " didn't",
      "tOffsetMs": 2310,
      "acAsrConf": 245
    } ]
  }, {
    "tStartMs": 1302620,
    "dDurationMs": 1330,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1302630,
    "dDurationMs": 5010,
    "wWinId": 1,
    "segs": [ {
      "utf8": "get",
      "acAsrConf": 241
    }, {
      "utf8": " it",
      "tOffsetMs": 150,
      "acAsrConf": 252
    }, {
      "utf8": " I'll",
      "tOffsetMs": 180,
      "acAsrConf": 187
    }, {
      "utf8": " make",
      "tOffsetMs": 450,
      "acAsrConf": 252
    }, {
      "utf8": " sure",
      "tOffsetMs": 600,
      "acAsrConf": 166
    }, {
      "utf8": " I",
      "tOffsetMs": 780,
      "acAsrConf": 252
    }, {
      "utf8": " send",
      "tOffsetMs": 810,
      "acAsrConf": 236
    }, {
      "utf8": " it",
      "tOffsetMs": 960,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 1303940,
    "dDurationMs": 3700,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1303950,
    "dDurationMs": 6360,
    "wWinId": 1,
    "segs": [ {
      "utf8": "afterwards",
      "acAsrConf": 243
    }, {
      "utf8": " in",
      "tOffsetMs": 630,
      "acAsrConf": 201
    }, {
      "utf8": " the",
      "tOffsetMs": 1080,
      "acAsrConf": 215
    }, {
      "utf8": " mind",
      "tOffsetMs": 1980,
      "acAsrConf": 248
    }, {
      "utf8": " map",
      "tOffsetMs": 2220,
      "acAsrConf": 248
    }, {
      "utf8": " I",
      "tOffsetMs": 2490,
      "acAsrConf": 145
    }, {
      "utf8": " talked",
      "tOffsetMs": 2760,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1307630,
    "dDurationMs": 2680,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1307640,
    "dDurationMs": 5400,
    "wWinId": 1,
    "segs": [ {
      "utf8": "about",
      "acAsrConf": 252
    }, {
      "utf8": " three",
      "tOffsetMs": 90,
      "acAsrConf": 211
    }, {
      "utf8": " channels",
      "tOffsetMs": 360,
      "acAsrConf": 252
    }, {
      "utf8": " one",
      "tOffsetMs": 720,
      "acAsrConf": 247
    }, {
      "utf8": " is",
      "tOffsetMs": 930,
      "acAsrConf": 252
    }, {
      "utf8": " soul",
      "tOffsetMs": 1140,
      "acAsrConf": 227
    }, {
      "utf8": " pancake",
      "tOffsetMs": 2100,
      "acAsrConf": 239
    } ]
  }, {
    "tStartMs": 1310300,
    "dDurationMs": 2740,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1310310,
    "dDurationMs": 7370,
    "wWinId": 1,
    "segs": [ {
      "utf8": "in",
      "acAsrConf": 217
    }, {
      "utf8": " so",
      "tOffsetMs": 540,
      "acAsrConf": 252
    }, {
      "utf8": " pancakes",
      "tOffsetMs": 1170,
      "acAsrConf": 248
    }, {
      "utf8": " channel",
      "tOffsetMs": 1650,
      "acAsrConf": 203
    }, {
      "utf8": " they",
      "tOffsetMs": 2040,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 2250,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 2460,
      "acAsrConf": 252
    }, {
      "utf8": " lot",
      "tOffsetMs": 2490,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1313030,
    "dDurationMs": 4650,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1313040,
    "dDurationMs": 6430,
    "wWinId": 1,
    "segs": [ {
      "utf8": "of",
      "acAsrConf": 252
    }, {
      "utf8": " family-friendly",
      "tOffsetMs": 740,
      "acAsrConf": 190
    }, {
      "utf8": " happy",
      "tOffsetMs": 2330,
      "acAsrConf": 167
    }, {
      "utf8": " interesting",
      "tOffsetMs": 3360,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1317670,
    "dDurationMs": 1800,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1317680,
    "dDurationMs": 4420,
    "wWinId": 1,
    "segs": [ {
      "utf8": "interviews",
      "acAsrConf": 238
    } ]
  }, {
    "tStartMs": 1319460,
    "dDurationMs": 2640,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1319470,
    "dDurationMs": 5670,
    "wWinId": 1,
    "segs": [ {
      "utf8": "our",
      "acAsrConf": 226
    }, {
      "utf8": " conversations",
      "tOffsetMs": 140,
      "acAsrConf": 252
    }, {
      "utf8": " lots",
      "tOffsetMs": 1140,
      "acAsrConf": 216
    }, {
      "utf8": " of",
      "tOffsetMs": 1680,
      "acAsrConf": 252
    }, {
      "utf8": " stuff",
      "tOffsetMs": 1829,
      "acAsrConf": 252
    }, {
      "utf8": " going",
      "tOffsetMs": 2069,
      "acAsrConf": 160
    }, {
      "utf8": " on",
      "tOffsetMs": 2370,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1322090,
    "dDurationMs": 3050,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1322100,
    "dDurationMs": 5829,
    "wWinId": 1,
    "segs": [ {
      "utf8": "yes",
      "acAsrConf": 200
    }, {
      "utf8": " it's",
      "tOffsetMs": 1000,
      "acAsrConf": 247
    }, {
      "utf8": " supposed",
      "tOffsetMs": 1209,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1329,
      "acAsrConf": 163
    }, {
      "utf8": " be",
      "tOffsetMs": 1600,
      "acAsrConf": 217
    }, {
      "utf8": " inspiring",
      "tOffsetMs": 1780,
      "acAsrConf": 230
    }, {
      "utf8": " yes",
      "tOffsetMs": 2410,
      "acAsrConf": 199
    } ]
  }, {
    "tStartMs": 1325130,
    "dDurationMs": 2799,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1325140,
    "dDurationMs": 4440,
    "wWinId": 1,
    "segs": [ {
      "utf8": "I'm",
      "acAsrConf": 231
    }, {
      "utf8": " seeing",
      "tOffsetMs": 180,
      "acAsrConf": 249
    }, {
      "utf8": " ideas",
      "tOffsetMs": 539,
      "acAsrConf": 236
    }, {
      "utf8": " make",
      "tOffsetMs": 1500,
      "acAsrConf": 217
    }, {
      "utf8": " you",
      "tOffsetMs": 2100,
      "acAsrConf": 252
    }, {
      "utf8": " feel",
      "tOffsetMs": 2220,
      "acAsrConf": 232
    }, {
      "utf8": " like",
      "tOffsetMs": 2399,
      "acAsrConf": 222
    }, {
      "utf8": " I'm",
      "tOffsetMs": 2610,
      "acAsrConf": 235
    } ]
  }, {
    "tStartMs": 1327919,
    "dDurationMs": 1661,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1327929,
    "dDurationMs": 3691,
    "wWinId": 1,
    "segs": [ {
      "utf8": "happy",
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 360,
      "acAsrConf": 201
    }, {
      "utf8": " be",
      "tOffsetMs": 691,
      "acAsrConf": 252
    }, {
      "utf8": " alive",
      "tOffsetMs": 721,
      "acAsrConf": 142
    } ]
  }, {
    "tStartMs": 1329570,
    "dDurationMs": 2050,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1329580,
    "dDurationMs": 5550,
    "wWinId": 1,
    "segs": [ {
      "utf8": "oh",
      "acAsrConf": 252
    }, {
      "utf8": " yes",
      "tOffsetMs": 30,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 540,
      "acAsrConf": 225
    }, {
      "utf8": " can",
      "tOffsetMs": 900,
      "acAsrConf": 186
    }, {
      "utf8": " do",
      "tOffsetMs": 1140,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 1290,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 1320,
      "acAsrConf": 164
    }, {
      "utf8": " can",
      "tOffsetMs": 1530,
      "acAsrConf": 252
    }, {
      "utf8": " learn",
      "tOffsetMs": 1560,
      "acAsrConf": 196
    }, {
      "utf8": " English",
      "tOffsetMs": 1890,
      "acAsrConf": 246
    } ]
  }, {
    "tStartMs": 1331610,
    "dDurationMs": 3520,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1331620,
    "dDurationMs": 5400,
    "wWinId": 1,
    "segs": [ {
      "utf8": "yes",
      "acAsrConf": 167
    }, {
      "utf8": " but",
      "tOffsetMs": 510,
      "acAsrConf": 128
    }, {
      "utf8": " it",
      "tOffsetMs": 1470,
      "acAsrConf": 252
    }, {
      "utf8": " has",
      "tOffsetMs": 1590,
      "acAsrConf": 206
    }, {
      "utf8": " very",
      "tOffsetMs": 1740,
      "acAsrConf": 244
    }, {
      "utf8": " real",
      "tOffsetMs": 1770,
      "acAsrConf": 243
    }, {
      "utf8": " English",
      "tOffsetMs": 2370,
      "acAsrConf": 233
    }, {
      "utf8": " yes",
      "tOffsetMs": 2880,
      "acAsrConf": 208
    }, {
      "utf8": " so",
      "tOffsetMs": 3270,
      "acAsrConf": 234
    } ]
  }, {
    "tStartMs": 1335120,
    "dDurationMs": 1900,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1335130,
    "dDurationMs": 4440,
    "wWinId": 1,
    "segs": [ {
      "utf8": "this",
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 480,
      "acAsrConf": 201
    }, {
      "utf8": " one",
      "tOffsetMs": 539,
      "acAsrConf": 186
    }, {
      "utf8": " of",
      "tOffsetMs": 810,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 840,
      "acAsrConf": 202
    }, {
      "utf8": " channels",
      "tOffsetMs": 1020,
      "acAsrConf": 212
    }, {
      "utf8": " I",
      "tOffsetMs": 1350,
      "acAsrConf": 236
    }, {
      "utf8": " recommend",
      "tOffsetMs": 1440,
      "acAsrConf": 245
    } ]
  }, {
    "tStartMs": 1337010,
    "dDurationMs": 2560,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1337020,
    "dDurationMs": 4920,
    "wWinId": 1,
    "segs": [ {
      "utf8": "and",
      "acAsrConf": 236
    }, {
      "utf8": " the",
      "tOffsetMs": 269,
      "acAsrConf": 252
    }, {
      "utf8": " second",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " one",
      "tOffsetMs": 1080,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 1200,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    }, {
      "utf8": " BuzzFeed",
      "tOffsetMs": 1529,
      "acAsrConf": 167
    }, {
      "utf8": " video",
      "tOffsetMs": 2279,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1339560,
    "dDurationMs": 2380,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1339570,
    "dDurationMs": 5700,
    "wWinId": 1,
    "segs": [ {
      "utf8": "this",
      "acAsrConf": 201
    }, {
      "utf8": " is",
      "tOffsetMs": 510,
      "acAsrConf": 207
    }, {
      "utf8": " like",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " pop",
      "tOffsetMs": 810,
      "acAsrConf": 241
    }, {
      "utf8": " culture",
      "tOffsetMs": 1050,
      "acAsrConf": 241
    }, {
      "utf8": " kind",
      "tOffsetMs": 1080,
      "acAsrConf": 224
    }, {
      "utf8": " of",
      "tOffsetMs": 1920,
      "acAsrConf": 252
    }, {
      "utf8": " videos",
      "tOffsetMs": 2040,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1341930,
    "dDurationMs": 3340,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1341940,
    "dDurationMs": 8150,
    "wWinId": 1,
    "segs": [ {
      "utf8": "and",
      "acAsrConf": 237
    }, {
      "utf8": " the",
      "tOffsetMs": 330,
      "acAsrConf": 240
    }, {
      "utf8": " final",
      "tOffsetMs": 1109,
      "acAsrConf": 252
    }, {
      "utf8": " one",
      "tOffsetMs": 1469,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 1619,
      "acAsrConf": 216
    }, {
      "utf8": " the",
      "tOffsetMs": 1830,
      "acAsrConf": 222
    }, {
      "utf8": " fine",
      "tOffsetMs": 1920,
      "acAsrConf": 187
    }, {
      "utf8": " brothers",
      "tOffsetMs": 2550,
      "acAsrConf": 83
    } ]
  }, {
    "tStartMs": 1345260,
    "dDurationMs": 4830,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1345270,
    "dDurationMs": 9000,
    "wWinId": 1,
    "segs": [ {
      "utf8": "so",
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " fine",
      "tOffsetMs": 360,
      "acAsrConf": 252
    }, {
      "utf8": " brothers",
      "tOffsetMs": 570,
      "acAsrConf": 201
    }, {
      "utf8": " videos",
      "tOffsetMs": 1019,
      "acAsrConf": 252
    }, {
      "utf8": " they",
      "tOffsetMs": 1850,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 2850,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1350080,
    "dDurationMs": 4190,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1350090,
    "dDurationMs": 6420,
    "wWinId": 1,
    "segs": [ {
      "utf8": "interviews",
      "acAsrConf": 229
    }, {
      "utf8": " and",
      "tOffsetMs": 1000,
      "acAsrConf": 235
    }, {
      "utf8": " reaction",
      "tOffsetMs": 1620,
      "acAsrConf": 251
    }, {
      "utf8": " videos",
      "tOffsetMs": 2620,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 3190,
      "acAsrConf": 241
    }, {
      "utf8": " their",
      "tOffsetMs": 3640,
      "acAsrConf": 238
    } ]
  }, {
    "tStartMs": 1354260,
    "dDurationMs": 2250,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1354270,
    "dDurationMs": 7350,
    "wWinId": 1,
    "segs": [ {
      "utf8": "videos",
      "acAsrConf": 252
    }, {
      "utf8": " are",
      "tOffsetMs": 389,
      "acAsrConf": 252
    }, {
      "utf8": " very",
      "tOffsetMs": 570,
      "acAsrConf": 252
    }, {
      "utf8": " famous",
      "tOffsetMs": 659,
      "acAsrConf": 152
    }, {
      "utf8": " for",
      "tOffsetMs": 1170,
      "acAsrConf": 166
    }, {
      "utf8": " showing",
      "tOffsetMs": 1620,
      "acAsrConf": 242
    } ]
  }, {
    "tStartMs": 1356500,
    "dDurationMs": 5120,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1356510,
    "dDurationMs": 7840,
    "wWinId": 1,
    "segs": [ {
      "utf8": "people",
      "acAsrConf": 204
    }, {
      "utf8": " reacting",
      "tOffsetMs": 1000,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1840,
      "acAsrConf": 216
    }, {
      "utf8": " for",
      "tOffsetMs": 2500,
      "acAsrConf": 231
    }, {
      "utf8": " example",
      "tOffsetMs": 3070,
      "acAsrConf": 245
    }, {
      "utf8": " a",
      "tOffsetMs": 3220,
      "acAsrConf": 142
    }, {
      "utf8": " hint",
      "tOffsetMs": 4029,
      "acAsrConf": 248
    } ]
  }, {
    "tStartMs": 1361610,
    "dDurationMs": 2740,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1361620,
    "dDurationMs": 5220,
    "wWinId": 1,
    "segs": [ {
      "utf8": "pop",
      "acAsrConf": 221
    }, {
      "utf8": " song",
      "tOffsetMs": 299,
      "acAsrConf": 221
    }, {
      "utf8": " like",
      "tOffsetMs": 689,
      "acAsrConf": 245
    }, {
      "utf8": " Gangnam",
      "tOffsetMs": 1110,
      "acAsrConf": 251
    }, {
      "utf8": " style",
      "tOffsetMs": 1410,
      "acAsrConf": 238
    }, {
      "utf8": " yes",
      "tOffsetMs": 1950,
      "acAsrConf": 132
    }, {
      "utf8": " like",
      "tOffsetMs": 2370,
      "acAsrConf": 224
    } ]
  }, {
    "tStartMs": 1364340,
    "dDurationMs": 2500,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1364350,
    "dDurationMs": 4079,
    "wWinId": 1,
    "segs": [ {
      "utf8": "Gangnam",
      "acAsrConf": 248
    }, {
      "utf8": " style",
      "tOffsetMs": 240,
      "acAsrConf": 221
    }, {
      "utf8": " or",
      "tOffsetMs": 750,
      "acAsrConf": 217
    }, {
      "utf8": " something",
      "tOffsetMs": 1380,
      "acAsrConf": 218
    }, {
      "utf8": " else",
      "tOffsetMs": 2160,
      "acAsrConf": 255
    }, {
      "utf8": " they'll",
      "tOffsetMs": 2309,
      "acAsrConf": 240
    } ]
  }, {
    "tStartMs": 1366830,
    "dDurationMs": 1599,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1366840,
    "dDurationMs": 4650,
    "wWinId": 1,
    "segs": [ {
      "utf8": "show",
      "acAsrConf": 252
    }, {
      "utf8": " people",
      "tOffsetMs": 209,
      "acAsrConf": 252
    }, {
      "utf8": " reacting",
      "tOffsetMs": 420,
      "acAsrConf": 225
    }, {
      "utf8": " so",
      "tOffsetMs": 1020,
      "acAsrConf": 237
    }, {
      "utf8": " you'll",
      "tOffsetMs": 1170,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 1350,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 1530,
      "acAsrConf": 206
    } ]
  }, {
    "tStartMs": 1368419,
    "dDurationMs": 3071,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1368429,
    "dDurationMs": 5161,
    "wWinId": 1,
    "segs": [ {
      "utf8": "chance",
      "acAsrConf": 198
    }, {
      "utf8": " to",
      "tOffsetMs": 271,
      "acAsrConf": 206
    }, {
      "utf8": " see",
      "tOffsetMs": 541,
      "acAsrConf": 252
    }, {
      "utf8": " natural",
      "tOffsetMs": 750,
      "acAsrConf": 201
    }, {
      "utf8": " reactions",
      "tOffsetMs": 1380,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 2130,
      "acAsrConf": 240
    }, {
      "utf8": " that",
      "tOffsetMs": 2461,
      "acAsrConf": 136
    } ]
  }, {
    "tStartMs": 1371480,
    "dDurationMs": 2110,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1371490,
    "dDurationMs": 5160,
    "wWinId": 1,
    "segs": [ {
      "utf8": "kind",
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 270,
      "acAsrConf": 252
    }, {
      "utf8": " dialogue",
      "tOffsetMs": 390,
      "acAsrConf": 252
    }, {
      "utf8": " back",
      "tOffsetMs": 660,
      "acAsrConf": 216
    }, {
      "utf8": " and",
      "tOffsetMs": 1260,
      "acAsrConf": 252
    }, {
      "utf8": " forth",
      "tOffsetMs": 1290,
      "acAsrConf": 231
    }, {
      "utf8": " so",
      "tOffsetMs": 1530,
      "acAsrConf": 113
    }, {
      "utf8": " these",
      "tOffsetMs": 1919,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1373580,
    "dDurationMs": 3070,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1373590,
    "dDurationMs": 6420,
    "wWinId": 1,
    "segs": [ {
      "utf8": "three",
      "acAsrConf": 236
    }, {
      "utf8": " channels",
      "tOffsetMs": 329,
      "acAsrConf": 201
    }, {
      "utf8": " I",
      "tOffsetMs": 780,
      "acAsrConf": 188
    }, {
      "utf8": " recommend",
      "tOffsetMs": 930,
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 1589,
      "acAsrConf": 243
    }, {
      "utf8": " listening",
      "tOffsetMs": 2250,
      "acAsrConf": 236
    } ]
  }, {
    "tStartMs": 1376640,
    "dDurationMs": 3370,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1376650,
    "dDurationMs": 7860,
    "wWinId": 1,
    "segs": [ {
      "utf8": "to",
      "acAsrConf": 248
    }, {
      "utf8": " natural",
      "tOffsetMs": 180,
      "acAsrConf": 252
    }, {
      "utf8": " English",
      "tOffsetMs": 810,
      "acAsrConf": 244
    }, {
      "utf8": " alright",
      "tOffsetMs": 1519,
      "acAsrConf": 100
    }, {
      "utf8": " let's",
      "tOffsetMs": 2519,
      "acAsrConf": 211
    }, {
      "utf8": " move",
      "tOffsetMs": 2850,
      "acAsrConf": 237
    }, {
      "utf8": " on",
      "tOffsetMs": 3120,
      "acAsrConf": 249
    } ]
  }, {
    "tStartMs": 1380000,
    "dDurationMs": 4510,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1380010,
    "dDurationMs": 8549,
    "wWinId": 1,
    "segs": [ {
      "utf8": "to",
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 360,
      "acAsrConf": 212
    }, {
      "utf8": " next",
      "tOffsetMs": 1350,
      "acAsrConf": 167
    }, {
      "utf8": " point",
      "tOffsetMs": 1680,
      "acAsrConf": 222
    }, {
      "utf8": " the",
      "tOffsetMs": 1919,
      "acAsrConf": 129
    }, {
      "utf8": " next",
      "tOffsetMs": 2299,
      "acAsrConf": 148
    }, {
      "utf8": " thing",
      "tOffsetMs": 3299,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 3570,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 4110,
      "acAsrConf": 245
    } ]
  }, {
    "tStartMs": 1384500,
    "dDurationMs": 4059,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1384510,
    "dDurationMs": 7590,
    "wWinId": 1,
    "segs": [ {
      "utf8": "have",
      "acAsrConf": 250
    }, {
      "utf8": " three",
      "tOffsetMs": 600,
      "acAsrConf": 144
    }, {
      "utf8": " tips",
      "tOffsetMs": 1409,
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 2039,
      "acAsrConf": 239
    }, {
      "utf8": " you",
      "tOffsetMs": 2340,
      "acAsrConf": 252
    }, {
      "utf8": " three",
      "tOffsetMs": 2549,
      "acAsrConf": 201
    }, {
      "utf8": " tips",
      "tOffsetMs": 3539,
      "acAsrConf": 221
    }, {
      "utf8": " of",
      "tOffsetMs": 3840,
      "acAsrConf": 228
    } ]
  }, {
    "tStartMs": 1388549,
    "dDurationMs": 3551,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1388559,
    "dDurationMs": 5941,
    "wWinId": 1,
    "segs": [ {
      "utf8": "what",
      "acAsrConf": 252
    }, {
      "utf8": " my",
      "tOffsetMs": 61,
      "acAsrConf": 165
    }, {
      "utf8": " most",
      "tOffsetMs": 841,
      "acAsrConf": 252
    }, {
      "utf8": " successful",
      "tOffsetMs": 1341,
      "acAsrConf": 166
    }, {
      "utf8": " students",
      "tOffsetMs": 2341,
      "acAsrConf": 246
    }, {
      "utf8": " do",
      "tOffsetMs": 2911,
      "acAsrConf": 252
    }, {
      "utf8": " this",
      "tOffsetMs": 3211,
      "acAsrConf": 216
    } ]
  }, {
    "tStartMs": 1392090,
    "dDurationMs": 2410,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1392100,
    "dDurationMs": 4260,
    "wWinId": 1,
    "segs": [ {
      "utf8": "is",
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 180,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 360,
      "acAsrConf": 211
    }, {
      "utf8": " yes",
      "tOffsetMs": 600,
      "acAsrConf": 233
    }, {
      "utf8": " for",
      "tOffsetMs": 1079,
      "acAsrConf": 202
    }, {
      "utf8": " you",
      "tOffsetMs": 1380,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1470,
      "acAsrConf": 226
    }, {
      "utf8": " you",
      "tOffsetMs": 1709,
      "acAsrConf": 252
    }, {
      "utf8": " want",
      "tOffsetMs": 1770,
      "acAsrConf": 225
    }, {
      "utf8": " to",
      "tOffsetMs": 2160,
      "acAsrConf": 235
    }, {
      "utf8": " be",
      "tOffsetMs": 2280,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1394490,
    "dDurationMs": 1870,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1394500,
    "dDurationMs": 4950,
    "wWinId": 1,
    "segs": [ {
      "utf8": "a",
      "acAsrConf": 252
    }, {
      "utf8": " successful",
      "tOffsetMs": 30,
      "acAsrConf": 252
    }, {
      "utf8": " english",
      "tOffsetMs": 450,
      "acAsrConf": 144
    }, {
      "utf8": " learner",
      "tOffsetMs": 990,
      "acAsrConf": 147
    }, {
      "utf8": " of",
      "tOffsetMs": 1260,
      "acAsrConf": 219
    }, {
      "utf8": " course",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1396350,
    "dDurationMs": 3100,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1396360,
    "dDurationMs": 4500,
    "wWinId": 1,
    "segs": [ {
      "utf8": "so",
      "acAsrConf": 238
    }, {
      "utf8": " there",
      "tOffsetMs": 809,
      "acAsrConf": 252
    }, {
      "utf8": " are",
      "tOffsetMs": 1559,
      "acAsrConf": 214
    }, {
      "utf8": " three",
      "tOffsetMs": 1830,
      "acAsrConf": 237
    }, {
      "utf8": " things",
      "tOffsetMs": 2220,
      "acAsrConf": 252
    }, {
      "utf8": " that",
      "tOffsetMs": 2250,
      "acAsrConf": 162
    }, {
      "utf8": " my",
      "tOffsetMs": 2580,
      "acAsrConf": 214
    } ]
  }, {
    "tStartMs": 1399440,
    "dDurationMs": 1420,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1399450,
    "dDurationMs": 5130,
    "wWinId": 1,
    "segs": [ {
      "utf8": "successful",
      "acAsrConf": 207
    }, {
      "utf8": " students",
      "tOffsetMs": 570,
      "acAsrConf": 252
    }, {
      "utf8": " do",
      "tOffsetMs": 870,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 990,
      "acAsrConf": 252
    }, {
      "utf8": " successful",
      "tOffsetMs": 1020,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1400850,
    "dDurationMs": 3730,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1400860,
    "dDurationMs": 5429,
    "wWinId": 1,
    "segs": [ {
      "utf8": "student",
      "acAsrConf": 165
    }, {
      "utf8": " is",
      "tOffsetMs": 660,
      "acAsrConf": 227
    }, {
      "utf8": " someone",
      "tOffsetMs": 1500,
      "acAsrConf": 216
    }, {
      "utf8": " who",
      "tOffsetMs": 2280,
      "acAsrConf": 252
    }, {
      "utf8": " wants",
      "tOffsetMs": 2640,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 3120,
      "acAsrConf": 153
    }, {
      "utf8": " improve",
      "tOffsetMs": 3330,
      "acAsrConf": 138
    } ]
  }, {
    "tStartMs": 1404570,
    "dDurationMs": 1719,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1404580,
    "dDurationMs": 4229,
    "wWinId": 1,
    "segs": [ {
      "utf8": "their",
      "acAsrConf": 130
    }, {
      "utf8": " English",
      "tOffsetMs": 180,
      "acAsrConf": 212
    }, {
      "utf8": " and",
      "tOffsetMs": 209,
      "acAsrConf": 130
    }, {
      "utf8": " then",
      "tOffsetMs": 630,
      "acAsrConf": 205
    }, {
      "utf8": " actually",
      "tOffsetMs": 900,
      "acAsrConf": 239
    }, {
      "utf8": " does",
      "tOffsetMs": 1500,
      "acAsrConf": 251
    } ]
  }, {
    "tStartMs": 1406279,
    "dDurationMs": 2530,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1406289,
    "dDurationMs": 7831,
    "wWinId": 1,
    "segs": [ {
      "utf8": "improve",
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 451,
      "acAsrConf": 227
    }, {
      "utf8": " these",
      "tOffsetMs": 750,
      "acAsrConf": 214
    }, {
      "utf8": " three",
      "tOffsetMs": 1620,
      "acAsrConf": 247
    }, {
      "utf8": " things",
      "tOffsetMs": 1831,
      "acAsrConf": 207
    }, {
      "utf8": " the",
      "tOffsetMs": 2101,
      "acAsrConf": 252
    }, {
      "utf8": " first",
      "tOffsetMs": 2311,
      "acAsrConf": 165
    } ]
  }, {
    "tStartMs": 1408799,
    "dDurationMs": 5321,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1408809,
    "dDurationMs": 8370,
    "wWinId": 1,
    "segs": [ {
      "utf8": "one",
      "acAsrConf": 238
    }, {
      "utf8": " is",
      "tOffsetMs": 151,
      "acAsrConf": 205
    }, {
      "utf8": " they",
      "tOffsetMs": 451,
      "acAsrConf": 236
    }, {
      "utf8": " ask",
      "tOffsetMs": 1441,
      "acAsrConf": 223
    }, {
      "utf8": " a",
      "tOffsetMs": 2431,
      "acAsrConf": 252
    }, {
      "utf8": " lot",
      "tOffsetMs": 3120,
      "acAsrConf": 171
    }, {
      "utf8": " of",
      "tOffsetMs": 3571,
      "acAsrConf": 252
    }, {
      "utf8": " questions",
      "tOffsetMs": 3811,
      "acAsrConf": 252
    }, {
      "utf8": " they",
      "tOffsetMs": 4311,
      "acAsrConf": 202
    } ]
  }, {
    "tStartMs": 1414110,
    "dDurationMs": 3069,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1414120,
    "dDurationMs": 5760,
    "wWinId": 1,
    "segs": [ {
      "utf8": "ask",
      "acAsrConf": 244
    }, {
      "utf8": " questions",
      "tOffsetMs": 210,
      "acAsrConf": 122
    }, {
      "utf8": " about",
      "tOffsetMs": 780,
      "acAsrConf": 248
    }, {
      "utf8": " some",
      "tOffsetMs": 1160,
      "acAsrConf": 220
    }, {
      "utf8": " English",
      "tOffsetMs": 2160,
      "acAsrConf": 239
    }, {
      "utf8": " video",
      "tOffsetMs": 2669,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1417169,
    "dDurationMs": 2711,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1417179,
    "dDurationMs": 5101,
    "wWinId": 1,
    "segs": [ {
      "utf8": "they",
      "acAsrConf": 252
    }, {
      "utf8": " saw",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 480,
      "acAsrConf": 227
    }, {
      "utf8": " podcast",
      "tOffsetMs": 691,
      "acAsrConf": 250
    }, {
      "utf8": " they",
      "tOffsetMs": 1321,
      "acAsrConf": 237
    }, {
      "utf8": " listen",
      "tOffsetMs": 1471,
      "acAsrConf": 143
    }, {
      "utf8": " to",
      "tOffsetMs": 1801,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1831,
      "acAsrConf": 215
    }, {
      "utf8": " oh",
      "tOffsetMs": 2250,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 1419870,
    "dDurationMs": 2410,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1419880,
    "dDurationMs": 5970,
    "wWinId": 1,
    "segs": [ {
      "utf8": "why",
      "acAsrConf": 233
    }, {
      "utf8": " do",
      "tOffsetMs": 690,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " think",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " it's",
      "tOffsetMs": 1110,
      "acAsrConf": 214
    }, {
      "utf8": " important",
      "tOffsetMs": 1289,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1770,
      "acAsrConf": 252
    }, {
      "utf8": " ask",
      "tOffsetMs": 2070,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1422270,
    "dDurationMs": 3580,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1422280,
    "dDurationMs": 6300,
    "wWinId": 1,
    "segs": [ {
      "utf8": "questions",
      "acAsrConf": 222
    }, {
      "utf8": " well",
      "tOffsetMs": 470,
      "acAsrConf": 151
    }, {
      "utf8": " if",
      "tOffsetMs": 1470,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1769,
      "acAsrConf": 252
    }, {
      "utf8": " ask",
      "tOffsetMs": 1920,
      "acAsrConf": 227
    }, {
      "utf8": " questions",
      "tOffsetMs": 2160,
      "acAsrConf": 204
    }, {
      "utf8": " then",
      "tOffsetMs": 3120,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1425840,
    "dDurationMs": 2740,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1425850,
    "dDurationMs": 5970,
    "wWinId": 1,
    "segs": [ {
      "utf8": "that",
      "acAsrConf": 252
    }, {
      "utf8": " means",
      "tOffsetMs": 209,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " are",
      "tOffsetMs": 689,
      "acAsrConf": 203
    }, {
      "utf8": " curious",
      "tOffsetMs": 780,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1370,
      "acAsrConf": 179
    }, {
      "utf8": " you",
      "tOffsetMs": 2370,
      "acAsrConf": 252
    }, {
      "utf8": " are",
      "tOffsetMs": 2430,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1428570,
    "dDurationMs": 3250,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1428580,
    "dDurationMs": 6000,
    "wWinId": 1,
    "segs": [ {
      "utf8": "not",
      "acAsrConf": 252
    }, {
      "utf8": " just",
      "tOffsetMs": 180,
      "acAsrConf": 208
    }, {
      "utf8": " thinking",
      "tOffsetMs": 780,
      "acAsrConf": 206
    }, {
      "utf8": " give",
      "tOffsetMs": 1469,
      "acAsrConf": 237
    }, {
      "utf8": " me",
      "tOffsetMs": 2250,
      "acAsrConf": 252
    }, {
      "utf8": " all",
      "tOffsetMs": 2579,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 2940,
      "acAsrConf": 227
    } ]
  }, {
    "tStartMs": 1431810,
    "dDurationMs": 2770,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1431820,
    "dDurationMs": 6000,
    "wWinId": 1,
    "segs": [ {
      "utf8": "information",
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 30,
      "acAsrConf": 167
    }, {
      "utf8": " are",
      "tOffsetMs": 989,
      "acAsrConf": 252
    }, {
      "utf8": " thinking",
      "tOffsetMs": 1020,
      "acAsrConf": 219
    }, {
      "utf8": " hmm",
      "tOffsetMs": 1469,
      "acAsrConf": 201
    }, {
      "utf8": " what",
      "tOffsetMs": 2400,
      "acAsrConf": 204
    } ]
  }, {
    "tStartMs": 1434570,
    "dDurationMs": 3250,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1434580,
    "dDurationMs": 6479,
    "wWinId": 1,
    "segs": [ {
      "utf8": "can",
      "acAsrConf": 243
    }, {
      "utf8": " I",
      "tOffsetMs": 329,
      "acAsrConf": 235
    }, {
      "utf8": " ask",
      "tOffsetMs": 420,
      "acAsrConf": 252
    }, {
      "utf8": " or",
      "tOffsetMs": 719,
      "acAsrConf": 233
    }, {
      "utf8": " how",
      "tOffsetMs": 1320,
      "acAsrConf": 147
    }, {
      "utf8": " can",
      "tOffsetMs": 1920,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 1979,
      "acAsrConf": 209
    }, {
      "utf8": " think",
      "tOffsetMs": 2400,
      "acAsrConf": 124
    }, {
      "utf8": " to",
      "tOffsetMs": 2760,
      "acAsrConf": 248
    }, {
      "utf8": " get",
      "tOffsetMs": 3060,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1437810,
    "dDurationMs": 3249,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1437820,
    "dDurationMs": 5400,
    "wWinId": 1,
    "segs": [ {
      "utf8": "better",
      "acAsrConf": 225
    }, {
      "utf8": " information",
      "tOffsetMs": 719,
      "acAsrConf": 252
    }, {
      "utf8": " myself",
      "tOffsetMs": 1370,
      "acAsrConf": 179
    }, {
      "utf8": " yes",
      "tOffsetMs": 2370,
      "acAsrConf": 178
    }, {
      "utf8": " so",
      "tOffsetMs": 2820,
      "acAsrConf": 201
    }, {
      "utf8": " a",
      "tOffsetMs": 3210,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1441049,
    "dDurationMs": 2171,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1441059,
    "dDurationMs": 4921,
    "wWinId": 1,
    "segs": [ {
      "utf8": "question",
      "acAsrConf": 252
    }, {
      "utf8": " puts",
      "tOffsetMs": 391,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 870,
      "acAsrConf": 252
    }, {
      "utf8": " power",
      "tOffsetMs": 1021,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 1201,
      "acAsrConf": 224
    }, {
      "utf8": " your",
      "tOffsetMs": 1381,
      "acAsrConf": 201
    }, {
      "utf8": " hands",
      "tOffsetMs": 1860,
      "acAsrConf": 236
    } ]
  }, {
    "tStartMs": 1443210,
    "dDurationMs": 2770,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1443220,
    "dDurationMs": 4500,
    "wWinId": 1,
    "segs": [ {
      "utf8": "yes",
      "acAsrConf": 147
    }, {
      "utf8": " so",
      "tOffsetMs": 420,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 630,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 689,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 930,
      "acAsrConf": 172
    }, {
      "utf8": " power",
      "tOffsetMs": 1290,
      "acAsrConf": 236
    }, {
      "utf8": " yourself",
      "tOffsetMs": 1500,
      "acAsrConf": 226
    }, {
      "utf8": " for",
      "tOffsetMs": 2100,
      "acAsrConf": 154
    } ]
  }, {
    "tStartMs": 1445970,
    "dDurationMs": 1750,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1445980,
    "dDurationMs": 3300,
    "wWinId": 1,
    "segs": [ {
      "utf8": "learning",
      "acAsrConf": 237
    }, {
      "utf8": " English",
      "tOffsetMs": 270,
      "acAsrConf": 244
    }, {
      "utf8": " when",
      "tOffsetMs": 569,
      "acAsrConf": 228
    }, {
      "utf8": " you",
      "tOffsetMs": 840,
      "acAsrConf": 137
    }, {
      "utf8": " ask",
      "tOffsetMs": 1050,
      "acAsrConf": 255
    }, {
      "utf8": " questions",
      "tOffsetMs": 1290,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 1447710,
    "dDurationMs": 1570,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1447720,
    "dDurationMs": 2100,
    "wWinId": 1,
    "segs": [ {
      "utf8": "so",
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 150,
      "acAsrConf": 244
    }, {
      "utf8": " know",
      "tOffsetMs": 179,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 360,
      "acAsrConf": 252
    }, {
      "utf8": " lot",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 689,
      "acAsrConf": 138
    }, {
      "utf8": " you",
      "tOffsetMs": 750,
      "acAsrConf": 228
    }, {
      "utf8": " here",
      "tOffsetMs": 959,
      "acAsrConf": 230
    }, {
      "utf8": " in",
      "tOffsetMs": 1230,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 1410,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1449270,
    "dDurationMs": 550,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1449280,
    "dDurationMs": 2699,
    "wWinId": 1,
    "segs": [ {
      "utf8": "question",
      "acAsrConf": 208
    } ]
  }, {
    "tStartMs": 1449810,
    "dDurationMs": 2169,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1449820,
    "dDurationMs": 4500,
    "wWinId": 1,
    "segs": [ {
      "utf8": "and",
      "acAsrConf": 241
    }, {
      "utf8": " Starbucks",
      "tOffsetMs": 89,
      "acAsrConf": 197
    }, {
      "utf8": " have",
      "tOffsetMs": 599,
      "acAsrConf": 221
    }, {
      "utf8": " asked",
      "tOffsetMs": 1260,
      "acAsrConf": 255
    }, {
      "utf8": " a",
      "tOffsetMs": 1589,
      "acAsrConf": 252
    }, {
      "utf8": " lot",
      "tOffsetMs": 1650,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 1799,
      "acAsrConf": 200
    }, {
      "utf8": " great",
      "tOffsetMs": 1859,
      "acAsrConf": 213
    } ]
  }, {
    "tStartMs": 1451969,
    "dDurationMs": 2351,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1451979,
    "dDurationMs": 4440,
    "wWinId": 1,
    "segs": [ {
      "utf8": "questions",
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 180,
      "acAsrConf": 146
    }, {
      "utf8": " great",
      "tOffsetMs": 601,
      "acAsrConf": 213
    }, {
      "utf8": " I'm",
      "tOffsetMs": 1320,
      "acAsrConf": 237
    }, {
      "utf8": " excited",
      "tOffsetMs": 1530,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1591,
      "acAsrConf": 148
    }, {
      "utf8": " answer",
      "tOffsetMs": 2070,
      "acAsrConf": 171
    } ]
  }, {
    "tStartMs": 1454310,
    "dDurationMs": 2109,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1454320,
    "dDurationMs": 4140,
    "wWinId": 1,
    "segs": [ {
      "utf8": "those",
      "acAsrConf": 252
    }, {
      "utf8": " when",
      "tOffsetMs": 150,
      "acAsrConf": 252
    }, {
      "utf8": " we",
      "tOffsetMs": 390,
      "acAsrConf": 166
    }, {
      "utf8": " finish",
      "tOffsetMs": 539,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 839,
      "acAsrConf": 233
    }, {
      "utf8": " asking",
      "tOffsetMs": 1109,
      "acAsrConf": 249
    } ]
  }, {
    "tStartMs": 1456409,
    "dDurationMs": 2051,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1456419,
    "dDurationMs": 6541,
    "wWinId": 1,
    "segs": [ {
      "utf8": "questions",
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 31,
      "acAsrConf": 165
    }, {
      "utf8": " great",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 781,
      "acAsrConf": 228
    }, {
      "utf8": " the",
      "tOffsetMs": 931,
      "acAsrConf": 252
    }, {
      "utf8": " second",
      "tOffsetMs": 1171,
      "acAsrConf": 255
    }, {
      "utf8": " thing",
      "tOffsetMs": 1771,
      "acAsrConf": 223
    } ]
  }, {
    "tStartMs": 1458450,
    "dDurationMs": 4510,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1458460,
    "dDurationMs": 8250,
    "wWinId": 1,
    "segs": [ {
      "utf8": "is",
      "acAsrConf": 252
    }, {
      "utf8": " they",
      "tOffsetMs": 269,
      "acAsrConf": 255
    }, {
      "utf8": " use",
      "tOffsetMs": 1199,
      "acAsrConf": 252
    }, {
      "utf8": " English",
      "tOffsetMs": 1679,
      "acAsrConf": 241
    }, {
      "utf8": " regularly",
      "tOffsetMs": 3289,
      "acAsrConf": 180
    }, {
      "utf8": " regularly",
      "tOffsetMs": 4289,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1462950,
    "dDurationMs": 3760,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1462960,
    "dDurationMs": 7110,
    "wWinId": 1,
    "segs": [ {
      "utf8": "exactly",
      "acAsrConf": 202
    }, {
      "utf8": " often",
      "tOffsetMs": 929,
      "acAsrConf": 249
    }, {
      "utf8": " it's",
      "tOffsetMs": 1409,
      "acAsrConf": 241
    }, {
      "utf8": " a",
      "tOffsetMs": 1829,
      "acAsrConf": 252
    }, {
      "utf8": " habit",
      "tOffsetMs": 2010,
      "acAsrConf": 252
    }, {
      "utf8": " maybe",
      "tOffsetMs": 2240,
      "acAsrConf": 187
    }, {
      "utf8": " even",
      "tOffsetMs": 3240,
      "acAsrConf": 232
    } ]
  }, {
    "tStartMs": 1466700,
    "dDurationMs": 3370,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1466710,
    "dDurationMs": 6750,
    "wWinId": 1,
    "segs": [ {
      "utf8": "every",
      "acAsrConf": 149
    }, {
      "utf8": " day",
      "tOffsetMs": 990,
      "acAsrConf": 160
    }, {
      "utf8": " yes",
      "tOffsetMs": 1169,
      "acAsrConf": 204
    }, {
      "utf8": " yes",
      "tOffsetMs": 1740,
      "acAsrConf": 208
    }, {
      "utf8": " they",
      "tOffsetMs": 2159,
      "acAsrConf": 255
    }, {
      "utf8": " use",
      "tOffsetMs": 2429,
      "acAsrConf": 252
    }, {
      "utf8": " at",
      "tOffsetMs": 2819,
      "acAsrConf": 234
    }, {
      "utf8": " least",
      "tOffsetMs": 3089,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 3179,
      "acAsrConf": 226
    } ]
  }, {
    "tStartMs": 1470060,
    "dDurationMs": 3400,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1470070,
    "dDurationMs": 6030,
    "wWinId": 1,
    "segs": [ {
      "utf8": "little",
      "acAsrConf": 252
    }, {
      "utf8": " yes",
      "tOffsetMs": 120,
      "acAsrConf": 227
    }, {
      "utf8": " as",
      "tOffsetMs": 1049,
      "acAsrConf": 149
    }, {
      "utf8": " much",
      "tOffsetMs": 1289,
      "acAsrConf": 252
    }, {
      "utf8": " as",
      "tOffsetMs": 1739,
      "acAsrConf": 252
    }, {
      "utf8": " possible",
      "tOffsetMs": 1979,
      "acAsrConf": 231
    }, {
      "utf8": " regularly",
      "tOffsetMs": 2390,
      "acAsrConf": 144
    } ]
  }, {
    "tStartMs": 1473450,
    "dDurationMs": 2650,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1473460,
    "dDurationMs": 5760,
    "wWinId": 1,
    "segs": [ {
      "utf8": "and",
      "acAsrConf": 182
    }, {
      "utf8": " maybe",
      "tOffsetMs": 209,
      "acAsrConf": 223
    }, {
      "utf8": " they",
      "tOffsetMs": 1140,
      "acAsrConf": 255
    }, {
      "utf8": " listen",
      "tOffsetMs": 1409,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1829,
      "acAsrConf": 252
    }, {
      "utf8": " English",
      "tOffsetMs": 1860,
      "acAsrConf": 204
    }, {
      "utf8": " every",
      "tOffsetMs": 2309,
      "acAsrConf": 227
    } ]
  }, {
    "tStartMs": 1476090,
    "dDurationMs": 3130,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1476100,
    "dDurationMs": 5370,
    "wWinId": 1,
    "segs": [ {
      "utf8": "day",
      "acAsrConf": 208
    }, {
      "utf8": " at",
      "tOffsetMs": 120,
      "acAsrConf": 252
    }, {
      "utf8": " breakfast",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " or",
      "tOffsetMs": 779,
      "acAsrConf": 203
    }, {
      "utf8": " they",
      "tOffsetMs": 1169,
      "acAsrConf": 237
    }, {
      "utf8": " listen",
      "tOffsetMs": 1949,
      "acAsrConf": 228
    }, {
      "utf8": " to",
      "tOffsetMs": 2909,
      "acAsrConf": 142
    }, {
      "utf8": " a",
      "tOffsetMs": 3090,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1479210,
    "dDurationMs": 2260,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1479220,
    "dDurationMs": 5689,
    "wWinId": 1,
    "segs": [ {
      "utf8": "video",
      "acAsrConf": 252
    }, {
      "utf8": " or",
      "tOffsetMs": 539,
      "acAsrConf": 207
    }, {
      "utf8": " watch",
      "tOffsetMs": 1049,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 1259,
      "acAsrConf": 252
    }, {
      "utf8": " video",
      "tOffsetMs": 1439,
      "acAsrConf": 252
    }, {
      "utf8": " while",
      "tOffsetMs": 1559,
      "acAsrConf": 252
    }, {
      "utf8": " they're",
      "tOffsetMs": 2039,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1481460,
    "dDurationMs": 3449,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1481470,
    "dDurationMs": 7230,
    "wWinId": 1,
    "segs": [ {
      "utf8": "sitting",
      "acAsrConf": 204
    }, {
      "utf8": " on",
      "tOffsetMs": 329,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 389,
      "acAsrConf": 220
    }, {
      "utf8": " train",
      "tOffsetMs": 569,
      "acAsrConf": 252
    }, {
      "utf8": " they",
      "tOffsetMs": 870,
      "acAsrConf": 232
    }, {
      "utf8": " they",
      "tOffsetMs": 1289,
      "acAsrConf": 243
    }, {
      "utf8": " reserved",
      "tOffsetMs": 2269,
      "acAsrConf": 250
    } ]
  }, {
    "tStartMs": 1484899,
    "dDurationMs": 3801,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1484909,
    "dDurationMs": 5830,
    "wWinId": 1,
    "segs": [ {
      "utf8": "five",
      "acAsrConf": 158
    }, {
      "utf8": " minutes",
      "tOffsetMs": 1000,
      "acAsrConf": 252
    }, {
      "utf8": " 30",
      "tOffsetMs": 1541,
      "acAsrConf": 186
    }, {
      "utf8": " minutes",
      "tOffsetMs": 2080,
      "acAsrConf": 156
    }, {
      "utf8": " a",
      "tOffsetMs": 2561,
      "acAsrConf": 232
    }, {
      "utf8": " short",
      "tOffsetMs": 2801,
      "acAsrConf": 252
    }, {
      "utf8": " period",
      "tOffsetMs": 3311,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 1488690,
    "dDurationMs": 2049,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1488700,
    "dDurationMs": 5189,
    "wWinId": 1,
    "segs": [ {
      "utf8": "of",
      "acAsrConf": 252
    }, {
      "utf8": " time",
      "tOffsetMs": 89,
      "acAsrConf": 214
    }, {
      "utf8": " that",
      "tOffsetMs": 149,
      "acAsrConf": 165
    }, {
      "utf8": " they",
      "tOffsetMs": 419,
      "acAsrConf": 145
    }, {
      "utf8": " can",
      "tOffsetMs": 1020,
      "acAsrConf": 112
    }, {
      "utf8": " use",
      "tOffsetMs": 1409,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1620,
      "acAsrConf": 252
    }, {
      "utf8": " learn",
      "tOffsetMs": 1649,
      "acAsrConf": 166
    } ]
  }, {
    "tStartMs": 1490729,
    "dDurationMs": 3160,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1490739,
    "dDurationMs": 5971,
    "wWinId": 1,
    "segs": [ {
      "utf8": "English",
      "acAsrConf": 244
    }, {
      "utf8": " regularly",
      "tOffsetMs": 620,
      "acAsrConf": 227
    }, {
      "utf8": " and",
      "tOffsetMs": 1620,
      "acAsrConf": 208
    }, {
      "utf8": " I",
      "tOffsetMs": 2190,
      "acAsrConf": 255
    }, {
      "utf8": " know",
      "tOffsetMs": 2430,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 2490,
      "acAsrConf": 255
    }, {
      "utf8": " have",
      "tOffsetMs": 2610,
      "acAsrConf": 249
    }, {
      "utf8": " some",
      "tOffsetMs": 2910,
      "acAsrConf": 251
    } ]
  }, {
    "tStartMs": 1493879,
    "dDurationMs": 2831,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1493889,
    "dDurationMs": 7081,
    "wWinId": 1,
    "segs": [ {
      "utf8": "students",
      "acAsrConf": 219
    }, {
      "utf8": " who",
      "tOffsetMs": 540,
      "acAsrConf": 252
    }, {
      "utf8": " really",
      "tOffsetMs": 750,
      "acAsrConf": 211
    }, {
      "utf8": " feel",
      "tOffsetMs": 1710,
      "acAsrConf": 252
    }, {
      "utf8": " passionate",
      "tOffsetMs": 2010,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 2250,
      "acAsrConf": 227
    } ]
  }, {
    "tStartMs": 1496700,
    "dDurationMs": 4270,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1496710,
    "dDurationMs": 6209,
    "wWinId": 1,
    "segs": [ {
      "utf8": "they",
      "acAsrConf": 252
    }, {
      "utf8": " watch",
      "tOffsetMs": 649,
      "acAsrConf": 251
    }, {
      "utf8": " an",
      "tOffsetMs": 1649,
      "acAsrConf": 249
    }, {
      "utf8": " english",
      "tOffsetMs": 2240,
      "acAsrConf": 252
    }, {
      "utf8": " movie",
      "tOffsetMs": 3240,
      "acAsrConf": 252
    }, {
      "utf8": " every",
      "tOffsetMs": 3419,
      "acAsrConf": 206
    }, {
      "utf8": " day",
      "tOffsetMs": 4079,
      "acAsrConf": 239
    } ]
  }, {
    "tStartMs": 1500960,
    "dDurationMs": 1959,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1500970,
    "dDurationMs": 6120,
    "wWinId": 1,
    "segs": [ {
      "utf8": "and",
      "acAsrConf": 200
    }, {
      "utf8": " listen",
      "tOffsetMs": 360,
      "acAsrConf": 247
    }, {
      "utf8": " to",
      "tOffsetMs": 509,
      "acAsrConf": 146
    }, {
      "utf8": " a",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " podcast",
      "tOffsetMs": 750,
      "acAsrConf": 248
    }, {
      "utf8": " every",
      "tOffsetMs": 929,
      "acAsrConf": 217
    }, {
      "utf8": " day",
      "tOffsetMs": 1379,
      "acAsrConf": 163
    } ]
  }, {
    "tStartMs": 1502909,
    "dDurationMs": 4181,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1502919,
    "dDurationMs": 8070,
    "wWinId": 1,
    "segs": [ {
      "utf8": "and",
      "acAsrConf": 252
    }, {
      "utf8": " after",
      "tOffsetMs": 90,
      "acAsrConf": 226
    }, {
      "utf8": " one",
      "tOffsetMs": 960,
      "acAsrConf": 233
    }, {
      "utf8": " week",
      "tOffsetMs": 1230,
      "acAsrConf": 152
    }, {
      "utf8": " they",
      "tOffsetMs": 2240,
      "acAsrConf": 209
    }, {
      "utf8": " feel",
      "tOffsetMs": 3240,
      "acAsrConf": 252
    }, {
      "utf8": " really",
      "tOffsetMs": 3570,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1507080,
    "dDurationMs": 3909,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1507090,
    "dDurationMs": 5219,
    "wWinId": 1,
    "segs": [ {
      "utf8": "tired",
      "acAsrConf": 201
    }, {
      "utf8": " and",
      "tOffsetMs": 539,
      "acAsrConf": 220
    }, {
      "utf8": " burnt",
      "tOffsetMs": 1049,
      "acAsrConf": 252
    }, {
      "utf8": " out",
      "tOffsetMs": 1589,
      "acAsrConf": 252
    }, {
      "utf8": " how",
      "tOffsetMs": 2630,
      "acAsrConf": 201
    }, {
      "utf8": " would",
      "tOffsetMs": 3630,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 3689,
      "acAsrConf": 204
    } ]
  }, {
    "tStartMs": 1510979,
    "dDurationMs": 1330,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1510989,
    "dDurationMs": 4620,
    "wWinId": 1,
    "segs": [ {
      "utf8": "describe",
      "acAsrConf": 252
    }, {
      "utf8": " burnt",
      "tOffsetMs": 270,
      "acAsrConf": 205
    }, {
      "utf8": " out",
      "tOffsetMs": 721,
      "acAsrConf": 243
    } ]
  }, {
    "tStartMs": 1512299,
    "dDurationMs": 3310,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1512309,
    "dDurationMs": 6330,
    "wWinId": 1,
    "segs": [ {
      "utf8": "it",
      "acAsrConf": 252
    }, {
      "utf8": " just",
      "tOffsetMs": 210,
      "acAsrConf": 252
    }, {
      "utf8": " means",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " very",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " tired",
      "tOffsetMs": 1141,
      "acAsrConf": 252
    }, {
      "utf8": " right",
      "tOffsetMs": 1500,
      "acAsrConf": 179
    }, {
      "utf8": " so",
      "tOffsetMs": 1980,
      "acAsrConf": 197
    }, {
      "utf8": " they",
      "tOffsetMs": 2670,
      "acAsrConf": 249
    } ]
  }, {
    "tStartMs": 1515599,
    "dDurationMs": 3040,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1515609,
    "dDurationMs": 5101,
    "wWinId": 1,
    "segs": [ {
      "utf8": "don't",
      "acAsrConf": 236
    }, {
      "utf8": " want",
      "tOffsetMs": 331,
      "acAsrConf": 237
    }, {
      "utf8": " to",
      "tOffsetMs": 690,
      "acAsrConf": 200
    }, {
      "utf8": " do",
      "tOffsetMs": 841,
      "acAsrConf": 217
    }, {
      "utf8": " more",
      "tOffsetMs": 1111,
      "acAsrConf": 238
    }, {
      "utf8": " work",
      "tOffsetMs": 1440,
      "acAsrConf": 252
    }, {
      "utf8": " right",
      "tOffsetMs": 1680,
      "acAsrConf": 201
    }, {
      "utf8": " right",
      "tOffsetMs": 2700,
      "acAsrConf": 238
    } ]
  }, {
    "tStartMs": 1518629,
    "dDurationMs": 2081,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1518639,
    "dDurationMs": 4321,
    "wWinId": 1,
    "segs": [ {
      "utf8": "so",
      "acAsrConf": 252
    }, {
      "utf8": " they",
      "tOffsetMs": 210,
      "acAsrConf": 252
    }, {
      "utf8": " did",
      "tOffsetMs": 390,
      "acAsrConf": 252
    }, {
      "utf8": " too",
      "tOffsetMs": 600,
      "acAsrConf": 249
    }, {
      "utf8": " much",
      "tOffsetMs": 870,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1260,
      "acAsrConf": 227
    }, {
      "utf8": " then",
      "tOffsetMs": 1590,
      "acAsrConf": 252
    }, {
      "utf8": " they",
      "tOffsetMs": 1770,
      "acAsrConf": 252
    }, {
      "utf8": " feel",
      "tOffsetMs": 1860,
      "acAsrConf": 255
    } ]
  }, {
    "tStartMs": 1520700,
    "dDurationMs": 2260,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1520710,
    "dDurationMs": 4049,
    "wWinId": 1,
    "segs": [ {
      "utf8": "burnt",
      "acAsrConf": 255
    }, {
      "utf8": " out",
      "tOffsetMs": 329,
      "acAsrConf": 249
    }, {
      "utf8": " or",
      "tOffsetMs": 539,
      "acAsrConf": 248
    }, {
      "utf8": " maybe",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " they",
      "tOffsetMs": 899,
      "acAsrConf": 236
    }, {
      "utf8": " feel",
      "tOffsetMs": 1169,
      "acAsrConf": 252
    }, {
      "utf8": " frustrated",
      "tOffsetMs": 1409,
      "acAsrConf": 143
    } ]
  }, {
    "tStartMs": 1522950,
    "dDurationMs": 1809,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1522960,
    "dDurationMs": 4799,
    "wWinId": 1,
    "segs": [ {
      "utf8": "because",
      "acAsrConf": 252
    }, {
      "utf8": " they",
      "tOffsetMs": 149,
      "acAsrConf": 183
    }, {
      "utf8": " worked",
      "tOffsetMs": 959,
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 1230,
      "acAsrConf": 252
    }, {
      "utf8": " eight",
      "tOffsetMs": 1500,
      "acAsrConf": 206
    }, {
      "utf8": " hours",
      "tOffsetMs": 1769,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1524749,
    "dDurationMs": 3010,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1524759,
    "dDurationMs": 6571,
    "wWinId": 1,
    "segs": [ {
      "utf8": "every",
      "acAsrConf": 201
    }, {
      "utf8": " day",
      "tOffsetMs": 691,
      "acAsrConf": 154
    }, {
      "utf8": " for",
      "tOffsetMs": 1380,
      "acAsrConf": 166
    }, {
      "utf8": " a",
      "tOffsetMs": 1770,
      "acAsrConf": 252
    }, {
      "utf8": " week",
      "tOffsetMs": 1800,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 2071,
      "acAsrConf": 227
    }, {
      "utf8": " their",
      "tOffsetMs": 2370,
      "acAsrConf": 226
    }, {
      "utf8": " English",
      "tOffsetMs": 2670,
      "acAsrConf": 203
    } ]
  }, {
    "tStartMs": 1527749,
    "dDurationMs": 3581,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1527759,
    "dDurationMs": 6900,
    "wWinId": 1,
    "segs": [ {
      "utf8": "isn't",
      "acAsrConf": 248
    }, {
      "utf8": " improving",
      "tOffsetMs": 270,
      "acAsrConf": 207
    }, {
      "utf8": " so",
      "tOffsetMs": 961,
      "acAsrConf": 223
    }, {
      "utf8": " take",
      "tOffsetMs": 1951,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 2340,
      "acAsrConf": 252
    }, {
      "utf8": " slowly",
      "tOffsetMs": 2490,
      "acAsrConf": 252
    }, {
      "utf8": " take",
      "tOffsetMs": 2821,
      "acAsrConf": 187
    } ]
  }, {
    "tStartMs": 1531320,
    "dDurationMs": 3339,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1531330,
    "dDurationMs": 6679,
    "wWinId": 1,
    "segs": [ {
      "utf8": "it",
      "acAsrConf": 252
    }, {
      "utf8": " step",
      "tOffsetMs": 29,
      "acAsrConf": 187
    }, {
      "utf8": " by",
      "tOffsetMs": 630,
      "acAsrConf": 226
    }, {
      "utf8": " step",
      "tOffsetMs": 659,
      "acAsrConf": 226
    }, {
      "utf8": " take",
      "tOffsetMs": 949,
      "acAsrConf": 145
    }, {
      "utf8": " English",
      "tOffsetMs": 1949,
      "acAsrConf": 248
    }, {
      "utf8": " in",
      "tOffsetMs": 2219,
      "acAsrConf": 206
    }, {
      "utf8": " small",
      "tOffsetMs": 2640,
      "acAsrConf": 244
    } ]
  }, {
    "tStartMs": 1534649,
    "dDurationMs": 3360,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1534659,
    "dDurationMs": 6240,
    "wWinId": 1,
    "segs": [ {
      "utf8": "sustainable",
      "acAsrConf": 252
    }, {
      "utf8": " pieces",
      "tOffsetMs": 980,
      "acAsrConf": 205
    }, {
      "utf8": " okay",
      "tOffsetMs": 2010,
      "acAsrConf": 0
    }, {
      "utf8": " bite-size",
      "tOffsetMs": 2460,
      "acAsrConf": 166
    } ]
  }, {
    "tStartMs": 1537999,
    "dDurationMs": 2900,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1538009,
    "dDurationMs": 5380,
    "wWinId": 1,
    "segs": [ {
      "utf8": "bite-sized",
      "acAsrConf": 163
    }, {
      "utf8": " chalk",
      "tOffsetMs": 1000,
      "acAsrConf": 205
    }, {
      "utf8": " sium",
      "tOffsetMs": 1451,
      "acAsrConf": 109
    }, {
      "utf8": " bite-size",
      "tOffsetMs": 1870,
      "acAsrConf": 75
    }, {
      "utf8": " pieces",
      "tOffsetMs": 2530,
      "acAsrConf": 249
    } ]
  }, {
    "tStartMs": 1540889,
    "dDurationMs": 2500,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1540899,
    "dDurationMs": 4700,
    "wWinId": 1,
    "segs": [ {
      "utf8": "and",
      "acAsrConf": 186
    }, {
      "utf8": " you",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " need",
      "tOffsetMs": 1200,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1410,
      "acAsrConf": 238
    }, {
      "utf8": " do",
      "tOffsetMs": 1590,
      "acAsrConf": 252
    }, {
      "utf8": " something",
      "tOffsetMs": 1740,
      "acAsrConf": 207
    }, {
      "utf8": " that",
      "tOffsetMs": 2160,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 2191,
      "acAsrConf": 165
    } ]
  }, {
    "tStartMs": 1543379,
    "dDurationMs": 2220,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1543389,
    "dDurationMs": 4681,
    "wWinId": 1,
    "segs": [ {
      "utf8": "sustainable",
      "acAsrConf": 200
    }, {
      "utf8": " don't",
      "tOffsetMs": 660,
      "acAsrConf": 200
    }, {
      "utf8": " forget",
      "tOffsetMs": 1140,
      "acAsrConf": 236
    }, {
      "utf8": " to",
      "tOffsetMs": 1650,
      "acAsrConf": 252
    }, {
      "utf8": " be",
      "tOffsetMs": 1831,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1545589,
    "dDurationMs": 2481,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1545599,
    "dDurationMs": 5050,
    "wWinId": 1,
    "segs": [ {
      "utf8": "sustainable",
      "acAsrConf": 207
    }, {
      "utf8": " something",
      "tOffsetMs": 1000,
      "acAsrConf": 218
    }, {
      "utf8": " that",
      "tOffsetMs": 1810,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1960,
      "acAsrConf": 252
    }, {
      "utf8": " can",
      "tOffsetMs": 2200,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1548060,
    "dDurationMs": 2589,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1548070,
    "dDurationMs": 5640,
    "wWinId": 1,
    "segs": [ {
      "utf8": "continue",
      "acAsrConf": 207
    }, {
      "utf8": " doing",
      "tOffsetMs": 719,
      "acAsrConf": 146
    }, {
      "utf8": " for",
      "tOffsetMs": 900,
      "acAsrConf": 232
    }, {
      "utf8": " a",
      "tOffsetMs": 1439,
      "acAsrConf": 252
    }, {
      "utf8": " long",
      "tOffsetMs": 1679,
      "acAsrConf": 223
    }, {
      "utf8": " time",
      "tOffsetMs": 1979,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1550639,
    "dDurationMs": 3071,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1550649,
    "dDurationMs": 6510,
    "wWinId": 1,
    "segs": [ {
      "utf8": "not",
      "acAsrConf": 239
    }, {
      "utf8": " ten",
      "tOffsetMs": 571,
      "acAsrConf": 144
    }, {
      "utf8": " hours",
      "tOffsetMs": 1561,
      "acAsrConf": 208
    }, {
      "utf8": " every",
      "tOffsetMs": 1921,
      "acAsrConf": 234
    }, {
      "utf8": " day",
      "tOffsetMs": 2370,
      "acAsrConf": 234
    }, {
      "utf8": " it's",
      "tOffsetMs": 2400,
      "acAsrConf": 219
    }, {
      "utf8": " impossible",
      "tOffsetMs": 2730,
      "acAsrConf": 249
    } ]
  }, {
    "tStartMs": 1553700,
    "dDurationMs": 3459,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1553710,
    "dDurationMs": 7409,
    "wWinId": 1,
    "segs": [ {
      "utf8": "to",
      "acAsrConf": 215
    }, {
      "utf8": " continue",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " like",
      "tOffsetMs": 959,
      "acAsrConf": 252
    }, {
      "utf8": " that",
      "tOffsetMs": 1139,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1169,
      "acAsrConf": 128
    }, {
      "utf8": " find",
      "tOffsetMs": 1919,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 2279,
      "acAsrConf": 255
    }, {
      "utf8": " time",
      "tOffsetMs": 2610,
      "acAsrConf": 248
    } ]
  }, {
    "tStartMs": 1557149,
    "dDurationMs": 3970,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1557159,
    "dDurationMs": 5880,
    "wWinId": 1,
    "segs": [ {
      "utf8": "period",
      "acAsrConf": 252
    }, {
      "utf8": " five",
      "tOffsetMs": 510,
      "acAsrConf": 201
    }, {
      "utf8": " minutes",
      "tOffsetMs": 1140,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    }, {
      "utf8": " day",
      "tOffsetMs": 1620,
      "acAsrConf": 166
    }, {
      "utf8": " great",
      "tOffsetMs": 1770,
      "acAsrConf": 248
    }, {
      "utf8": " maybe",
      "tOffsetMs": 2431,
      "acAsrConf": 235
    }, {
      "utf8": " 30",
      "tOffsetMs": 3150,
      "acAsrConf": 236
    } ]
  }, {
    "tStartMs": 1561109,
    "dDurationMs": 1930,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1561119,
    "dDurationMs": 5130,
    "wWinId": 1,
    "segs": [ {
      "utf8": "minutes",
      "acAsrConf": 252
    }, {
      "utf8": " on",
      "tOffsetMs": 270,
      "acAsrConf": 187
    }, {
      "utf8": " the",
      "tOffsetMs": 660,
      "acAsrConf": 205
    }, {
      "utf8": " weekend",
      "tOffsetMs": 930,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1290,
      "acAsrConf": 217
    }, {
      "utf8": " can",
      "tOffsetMs": 1471,
      "acAsrConf": 252
    }, {
      "utf8": " do",
      "tOffsetMs": 1620,
      "acAsrConf": 207
    }, {
      "utf8": " it",
      "tOffsetMs": 1800,
      "acAsrConf": 222
    } ]
  }, {
    "tStartMs": 1563029,
    "dDurationMs": 3220,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1563039,
    "dDurationMs": 6210,
    "wWinId": 1,
    "segs": [ {
      "utf8": "every",
      "acAsrConf": 207
    }, {
      "utf8": " weekend",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " great",
      "tOffsetMs": 990,
      "acAsrConf": 217
    }, {
      "utf8": " find",
      "tOffsetMs": 1531,
      "acAsrConf": 206
    }, {
      "utf8": " something",
      "tOffsetMs": 2340,
      "acAsrConf": 203
    }, {
      "utf8": " that",
      "tOffsetMs": 3060,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 1566239,
    "dDurationMs": 3010,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1566249,
    "dDurationMs": 5520,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 206
    }, {
      "utf8": " can",
      "tOffsetMs": 180,
      "acAsrConf": 177
    }, {
      "utf8": " do",
      "tOffsetMs": 390,
      "acAsrConf": 252
    }, {
      "utf8": " continually",
      "tOffsetMs": 510,
      "acAsrConf": 238
    }, {
      "utf8": " and",
      "tOffsetMs": 1260,
      "acAsrConf": 226
    }, {
      "utf8": " take",
      "tOffsetMs": 1441,
      "acAsrConf": 237
    }, {
      "utf8": " action",
      "tOffsetMs": 2370,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1569239,
    "dDurationMs": 2530,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1569249,
    "dDurationMs": 5400,
    "wWinId": 1,
    "segs": [ {
      "utf8": "right",
      "acAsrConf": 222
    }, {
      "utf8": " action",
      "tOffsetMs": 240,
      "acAsrConf": 210
    }, {
      "utf8": " is",
      "tOffsetMs": 1201,
      "acAsrConf": 217
    }, {
      "utf8": " important",
      "tOffsetMs": 1321,
      "acAsrConf": 252
    }, {
      "utf8": " right",
      "tOffsetMs": 1800,
      "acAsrConf": 201
    }, {
      "utf8": " right",
      "tOffsetMs": 1980,
      "acAsrConf": 218
    }, {
      "utf8": " we",
      "tOffsetMs": 2341,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1571759,
    "dDurationMs": 2890,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1571769,
    "dDurationMs": 3931,
    "wWinId": 1,
    "segs": [ {
      "utf8": "have",
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 210,
      "acAsrConf": 252
    }, {
      "utf8": " saying",
      "tOffsetMs": 240,
      "acAsrConf": 230
    }, {
      "utf8": " in",
      "tOffsetMs": 600,
      "acAsrConf": 204
    }, {
      "utf8": " English",
      "tOffsetMs": 840,
      "acAsrConf": 252
    }, {
      "utf8": " called",
      "tOffsetMs": 960,
      "acAsrConf": 226
    }, {
      "utf8": " use",
      "tOffsetMs": 1801,
      "acAsrConf": 209
    }, {
      "utf8": " it",
      "tOffsetMs": 2610,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1574639,
    "dDurationMs": 1061,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1574649,
    "dDurationMs": 4010,
    "wWinId": 1,
    "segs": [ {
      "utf8": "or",
      "acAsrConf": 252
    }, {
      "utf8": " lose",
      "tOffsetMs": 210,
      "acAsrConf": 233
    }, {
      "utf8": " it",
      "tOffsetMs": 811,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1575690,
    "dDurationMs": 2969,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1575700,
    "dDurationMs": 7440,
    "wWinId": 1,
    "segs": [ {
      "utf8": "ah",
      "acAsrConf": 202
    }, {
      "utf8": " huh",
      "tOffsetMs": 209,
      "acAsrConf": 0
    }, {
      "utf8": " yeah",
      "tOffsetMs": 689,
      "acAsrConf": 203
    }, {
      "utf8": " this",
      "tOffsetMs": 1620,
      "acAsrConf": 217
    }, {
      "utf8": " means",
      "tOffsetMs": 1799,
      "acAsrConf": 235
    }, {
      "utf8": " you",
      "tOffsetMs": 2039,
      "acAsrConf": 252
    }, {
      "utf8": " must",
      "tOffsetMs": 2250,
      "acAsrConf": 246
    }, {
      "utf8": " practice",
      "tOffsetMs": 2490,
      "acAsrConf": 228
    } ]
  }, {
    "tStartMs": 1578649,
    "dDurationMs": 4491,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1578659,
    "dDurationMs": 6101,
    "wWinId": 1,
    "segs": [ {
      "utf8": "English",
      "acAsrConf": 242
    }, {
      "utf8": " regularly",
      "tOffsetMs": 1051,
      "acAsrConf": 235
    }, {
      "utf8": " or",
      "tOffsetMs": 2051,
      "acAsrConf": 200
    }, {
      "utf8": " you",
      "tOffsetMs": 2460,
      "acAsrConf": 244
    }, {
      "utf8": " will",
      "tOffsetMs": 3460,
      "acAsrConf": 248
    }, {
      "utf8": " lose",
      "tOffsetMs": 3671,
      "acAsrConf": 229
    } ]
  }, {
    "tStartMs": 1583130,
    "dDurationMs": 1630,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1583140,
    "dDurationMs": 4289,
    "wWinId": 1,
    "segs": [ {
      "utf8": "English",
      "acAsrConf": 243
    }, {
      "utf8": " it",
      "tOffsetMs": 360,
      "acAsrConf": 251
    }, {
      "utf8": " will",
      "tOffsetMs": 539,
      "acAsrConf": 252
    }, {
      "utf8": " just",
      "tOffsetMs": 630,
      "acAsrConf": 200
    }, {
      "utf8": " go",
      "tOffsetMs": 840,
      "acAsrConf": 252
    }, {
      "utf8": " away",
      "tOffsetMs": 1080,
      "acAsrConf": 201
    }, {
      "utf8": " in",
      "tOffsetMs": 1260,
      "acAsrConf": 216
    }, {
      "utf8": " your",
      "tOffsetMs": 1500,
      "acAsrConf": 216
    } ]
  }, {
    "tStartMs": 1584750,
    "dDurationMs": 2679,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1584760,
    "dDurationMs": 3780,
    "wWinId": 1,
    "segs": [ {
      "utf8": "head",
      "acAsrConf": 252
    }, {
      "utf8": " or",
      "tOffsetMs": 210,
      "acAsrConf": 240
    }, {
      "utf8": " it",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " will",
      "tOffsetMs": 1110,
      "acAsrConf": 242
    }, {
      "utf8": " not",
      "tOffsetMs": 1380,
      "acAsrConf": 213
    }, {
      "utf8": " be",
      "tOffsetMs": 1530,
      "acAsrConf": 208
    }, {
      "utf8": " effective",
      "tOffsetMs": 1770,
      "acAsrConf": 252
    }, {
      "utf8": " site",
      "tOffsetMs": 2280,
      "acAsrConf": 165
    } ]
  }, {
    "tStartMs": 1587419,
    "dDurationMs": 1121,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1587429,
    "dDurationMs": 5191,
    "wWinId": 1,
    "segs": [ {
      "utf8": "use",
      "acAsrConf": 188
    }, {
      "utf8": " it",
      "tOffsetMs": 391,
      "acAsrConf": 244
    }, {
      "utf8": " or",
      "tOffsetMs": 571,
      "acAsrConf": 252
    }, {
      "utf8": " lose",
      "tOffsetMs": 661,
      "acAsrConf": 232
    }, {
      "utf8": " it",
      "tOffsetMs": 870,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1588530,
    "dDurationMs": 4090,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1588540,
    "dDurationMs": 6900,
    "wWinId": 1,
    "segs": [ {
      "utf8": "oh",
      "acAsrConf": 227
    }, {
      "utf8": " yes",
      "tOffsetMs": 230,
      "acAsrConf": 238
    }, {
      "utf8": " yes",
      "tOffsetMs": 1820,
      "acAsrConf": 200
    }, {
      "utf8": " you",
      "tOffsetMs": 2820,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 3000,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 3030,
      "acAsrConf": 167
    }, {
      "utf8": " use",
      "tOffsetMs": 3240,
      "acAsrConf": 147
    }, {
      "utf8": " it",
      "tOffsetMs": 3420,
      "acAsrConf": 252
    }, {
      "utf8": " or",
      "tOffsetMs": 3630,
      "acAsrConf": 252
    }, {
      "utf8": " lose",
      "tOffsetMs": 3840,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 4050,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1592610,
    "dDurationMs": 2830,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1592620,
    "dDurationMs": 4800,
    "wWinId": 1,
    "segs": [ {
      "utf8": "so",
      "acAsrConf": 212
    }, {
      "utf8": " the",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " final",
      "tOffsetMs": 1050,
      "acAsrConf": 224
    }, {
      "utf8": " thing",
      "tOffsetMs": 1770,
      "acAsrConf": 252
    }, {
      "utf8": " number",
      "tOffsetMs": 1800,
      "acAsrConf": 222
    }, {
      "utf8": " three",
      "tOffsetMs": 2250,
      "acAsrConf": 188
    }, {
      "utf8": " so",
      "tOffsetMs": 2400,
      "acAsrConf": 206
    }, {
      "utf8": " the",
      "tOffsetMs": 2700,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1595430,
    "dDurationMs": 1990,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1595440,
    "dDurationMs": 6420,
    "wWinId": 1,
    "segs": [ {
      "utf8": "first",
      "acAsrConf": 252
    }, {
      "utf8": " thing",
      "tOffsetMs": 300,
      "acAsrConf": 252
    }, {
      "utf8": " that",
      "tOffsetMs": 600,
      "acAsrConf": 140
    }, {
      "utf8": " my",
      "tOffsetMs": 869,
      "acAsrConf": 252
    }, {
      "utf8": " successful",
      "tOffsetMs": 1020,
      "acAsrConf": 202
    }, {
      "utf8": " students",
      "tOffsetMs": 1619,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1597410,
    "dDurationMs": 4450,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1597420,
    "dDurationMs": 7040,
    "wWinId": 1,
    "segs": [ {
      "utf8": "do",
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 210,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 770,
      "acAsrConf": 252
    }, {
      "utf8": " ask",
      "tOffsetMs": 1770,
      "acAsrConf": 252
    }, {
      "utf8": " questions",
      "tOffsetMs": 2220,
      "acAsrConf": 157
    }, {
      "utf8": " the",
      "tOffsetMs": 3120,
      "acAsrConf": 216
    }, {
      "utf8": " second",
      "tOffsetMs": 3629,
      "acAsrConf": 252
    }, {
      "utf8": " one",
      "tOffsetMs": 4020,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 4170,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1601850,
    "dDurationMs": 2610,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1601860,
    "dDurationMs": 10500,
    "wWinId": 1,
    "segs": [ {
      "utf8": "that",
      "acAsrConf": 178
    }, {
      "utf8": " they",
      "tOffsetMs": 210,
      "acAsrConf": 252
    }, {
      "utf8": " use",
      "tOffsetMs": 360,
      "acAsrConf": 252
    }, {
      "utf8": " English",
      "tOffsetMs": 660,
      "acAsrConf": 239
    }, {
      "utf8": " regularly",
      "tOffsetMs": 1100,
      "acAsrConf": 226
    }, {
      "utf8": " and",
      "tOffsetMs": 2100,
      "acAsrConf": 220
    } ]
  }, {
    "tStartMs": 1604450,
    "dDurationMs": 7910,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1604460,
    "dDurationMs": 10750,
    "wWinId": 1,
    "segs": [ {
      "utf8": "finally",
      "acAsrConf": 248
    }, {
      "utf8": " they",
      "tOffsetMs": 1000,
      "acAsrConf": 250
    }, {
      "utf8": " have",
      "tOffsetMs": 1900,
      "acAsrConf": 218
    }, {
      "utf8": " diligence",
      "tOffsetMs": 3180,
      "acAsrConf": 189
    }, {
      "utf8": " no",
      "tOffsetMs": 4180,
      "acAsrConf": 96
    }, {
      "utf8": " and",
      "tOffsetMs": 5880,
      "acAsrConf": 242
    }, {
      "utf8": " they",
      "tOffsetMs": 6900,
      "acAsrConf": 241
    } ]
  }, {
    "tStartMs": 1612350,
    "dDurationMs": 2860,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1612360,
    "dDurationMs": 4740,
    "wWinId": 1,
    "segs": [ {
      "utf8": "have",
      "acAsrConf": 252
    }, {
      "utf8": " fun",
      "tOffsetMs": 270,
      "acAsrConf": 243
    }, {
      "utf8": " they",
      "tOffsetMs": 660,
      "acAsrConf": 242
    }, {
      "utf8": " have",
      "tOffsetMs": 990,
      "acAsrConf": 176
    }, {
      "utf8": " fun",
      "tOffsetMs": 1680,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1920,
      "acAsrConf": 252
    }, {
      "utf8": " they",
      "tOffsetMs": 2130,
      "acAsrConf": 246
    }, {
      "utf8": " laugh",
      "tOffsetMs": 2250,
      "acAsrConf": 252
    }, {
      "utf8": " at",
      "tOffsetMs": 2550,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1615200,
    "dDurationMs": 1900,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1615210,
    "dDurationMs": 5459,
    "wWinId": 1,
    "segs": [ {
      "utf8": "themselves",
      "acAsrConf": 252
    }, {
      "utf8": " they're",
      "tOffsetMs": 209,
      "acAsrConf": 252
    }, {
      "utf8": " not",
      "tOffsetMs": 810,
      "acAsrConf": 144
    }, {
      "utf8": " stressed",
      "tOffsetMs": 959,
      "acAsrConf": 252
    }, {
      "utf8": " they're",
      "tOffsetMs": 1469,
      "acAsrConf": 174
    } ]
  }, {
    "tStartMs": 1617090,
    "dDurationMs": 3579,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1617100,
    "dDurationMs": 5699,
    "wWinId": 1,
    "segs": [ {
      "utf8": "not",
      "acAsrConf": 252
    }, {
      "utf8": " worried",
      "tOffsetMs": 230,
      "acAsrConf": 252
    }, {
      "utf8": " about",
      "tOffsetMs": 1230,
      "acAsrConf": 252
    }, {
      "utf8": " improving",
      "tOffsetMs": 1260,
      "acAsrConf": 242
    }, {
      "utf8": " all",
      "tOffsetMs": 2190,
      "acAsrConf": 228
    }, {
      "utf8": " their",
      "tOffsetMs": 3000,
      "acAsrConf": 247
    } ]
  }, {
    "tStartMs": 1620659,
    "dDurationMs": 2140,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1620669,
    "dDurationMs": 5791,
    "wWinId": 1,
    "segs": [ {
      "utf8": "English",
      "acAsrConf": 251
    }, {
      "utf8": " in",
      "tOffsetMs": 390,
      "acAsrConf": 252
    }, {
      "utf8": " one",
      "tOffsetMs": 571,
      "acAsrConf": 234
    }, {
      "utf8": " week",
      "tOffsetMs": 811,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 1051,
      "acAsrConf": 215
    }, {
      "utf8": " have",
      "tOffsetMs": 1260,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1531,
      "acAsrConf": 252
    }, {
      "utf8": " do",
      "tOffsetMs": 1711,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 1831,
      "acAsrConf": 252
    }, {
      "utf8": " now",
      "tOffsetMs": 1981,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1622789,
    "dDurationMs": 3671,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1622799,
    "dDurationMs": 5760,
    "wWinId": 1,
    "segs": [ {
      "utf8": "no",
      "acAsrConf": 220
    }, {
      "utf8": " they",
      "tOffsetMs": 241,
      "acAsrConf": 241
    }, {
      "utf8": " just",
      "tOffsetMs": 781,
      "acAsrConf": 249
    }, {
      "utf8": " have",
      "tOffsetMs": 1411,
      "acAsrConf": 231
    }, {
      "utf8": " fun",
      "tOffsetMs": 1801,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 2130,
      "acAsrConf": 209
    }, {
      "utf8": " realized",
      "tOffsetMs": 2491,
      "acAsrConf": 248
    }, {
      "utf8": " that",
      "tOffsetMs": 3361,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1626450,
    "dDurationMs": 2109,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1626460,
    "dDurationMs": 6930,
    "wWinId": 1,
    "segs": [ {
      "utf8": "yes",
      "acAsrConf": 245
    }, {
      "utf8": " it's",
      "tOffsetMs": 599,
      "acAsrConf": 239
    }, {
      "utf8": " going",
      "tOffsetMs": 870,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1080,
      "acAsrConf": 207
    }, {
      "utf8": " take",
      "tOffsetMs": 1170,
      "acAsrConf": 201
    }, {
      "utf8": " time",
      "tOffsetMs": 1410,
      "acAsrConf": 229
    }, {
      "utf8": " but",
      "tOffsetMs": 1650,
      "acAsrConf": 252
    }, {
      "utf8": " don't",
      "tOffsetMs": 1680,
      "acAsrConf": 200
    } ]
  }, {
    "tStartMs": 1628549,
    "dDurationMs": 4841,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1628559,
    "dDurationMs": 8101,
    "wWinId": 1,
    "segs": [ {
      "utf8": "worry",
      "acAsrConf": 230
    }, {
      "utf8": " you",
      "tOffsetMs": 240,
      "acAsrConf": 121
    }, {
      "utf8": " can",
      "tOffsetMs": 781,
      "acAsrConf": 252
    }, {
      "utf8": " do",
      "tOffsetMs": 1021,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 1141,
      "acAsrConf": 229
    }, {
      "utf8": " so",
      "tOffsetMs": 1291,
      "acAsrConf": 216
    }, {
      "utf8": " for",
      "tOffsetMs": 1891,
      "acAsrConf": 252
    }, {
      "utf8": " example",
      "tOffsetMs": 2791,
      "acAsrConf": 203
    }, {
      "utf8": " one",
      "tOffsetMs": 3831,
      "acAsrConf": 240
    } ]
  }, {
    "tStartMs": 1633380,
    "dDurationMs": 3280,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1633390,
    "dDurationMs": 6419,
    "wWinId": 1,
    "segs": [ {
      "utf8": "of",
      "acAsrConf": 252
    }, {
      "utf8": " my",
      "tOffsetMs": 29,
      "acAsrConf": 216
    }, {
      "utf8": " students",
      "tOffsetMs": 180,
      "acAsrConf": 229
    }, {
      "utf8": " said",
      "tOffsetMs": 779,
      "acAsrConf": 188
    }, {
      "utf8": " last",
      "tOffsetMs": 990,
      "acAsrConf": 239
    }, {
      "utf8": " week",
      "tOffsetMs": 1380,
      "acAsrConf": 252
    }, {
      "utf8": " ah",
      "tOffsetMs": 1740,
      "acAsrConf": 216
    }, {
      "utf8": " my",
      "tOffsetMs": 2270,
      "acAsrConf": 216
    } ]
  }, {
    "tStartMs": 1636650,
    "dDurationMs": 3159,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1636660,
    "dDurationMs": 9210,
    "wWinId": 1,
    "segs": [ {
      "utf8": "favorite",
      "acAsrConf": 219
    }, {
      "utf8": " place",
      "tOffsetMs": 300,
      "acAsrConf": 201
    }, {
      "utf8": " in",
      "tOffsetMs": 930,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 1320,
      "acAsrConf": 200
    }, {
      "utf8": " earth",
      "tOffsetMs": 1769,
      "acAsrConf": 216
    }, {
      "utf8": " is",
      "tOffsetMs": 2040,
      "acAsrConf": 222
    }, {
      "utf8": " the",
      "tOffsetMs": 2670,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1639799,
    "dDurationMs": 6071,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1639809,
    "dDurationMs": 9901,
    "wWinId": 1,
    "segs": [ {
      "utf8": "mountains",
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 561,
      "acAsrConf": 206
    }, {
      "utf8": " I",
      "tOffsetMs": 1561,
      "acAsrConf": 244
    }, {
      "utf8": " said",
      "tOffsetMs": 2071,
      "acAsrConf": 252
    }, {
      "utf8": " wow",
      "tOffsetMs": 2870,
      "acAsrConf": 204
    }, {
      "utf8": " you",
      "tOffsetMs": 3870,
      "acAsrConf": 202
    }, {
      "utf8": " went",
      "tOffsetMs": 4500,
      "acAsrConf": 141
    }, {
      "utf8": " in",
      "tOffsetMs": 5101,
      "acAsrConf": 226
    }, {
      "utf8": " the",
      "tOffsetMs": 5490,
      "acAsrConf": 246
    } ]
  }, {
    "tStartMs": 1645860,
    "dDurationMs": 3850,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1645870,
    "dDurationMs": 6179,
    "wWinId": 1,
    "segs": [ {
      "utf8": "earth",
      "acAsrConf": 233
    }, {
      "utf8": " it's",
      "tOffsetMs": 330,
      "acAsrConf": 205
    }, {
      "utf8": " amazing",
      "tOffsetMs": 780,
      "acAsrConf": 252
    }, {
      "utf8": " but",
      "tOffsetMs": 1670,
      "acAsrConf": 237
    }, {
      "utf8": " I",
      "tOffsetMs": 2670,
      "acAsrConf": 216
    }, {
      "utf8": " think",
      "tOffsetMs": 2939,
      "acAsrConf": 221
    }, {
      "utf8": " now",
      "tOffsetMs": 3210,
      "acAsrConf": 252
    }, {
      "utf8": " he'll",
      "tOffsetMs": 3480,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1649700,
    "dDurationMs": 2349,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1649710,
    "dDurationMs": 4410,
    "wWinId": 1,
    "segs": [ {
      "utf8": "never",
      "acAsrConf": 252
    }, {
      "utf8": " forget",
      "tOffsetMs": 180,
      "acAsrConf": 229
    }, {
      "utf8": " because",
      "tOffsetMs": 750,
      "acAsrConf": 227
    }, {
      "utf8": " we",
      "tOffsetMs": 1050,
      "acAsrConf": 126
    }, {
      "utf8": " laughed",
      "tOffsetMs": 1440,
      "acAsrConf": 252
    }, {
      "utf8": " about",
      "tOffsetMs": 1830,
      "acAsrConf": 201
    }, {
      "utf8": " it",
      "tOffsetMs": 2280,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1652039,
    "dDurationMs": 2081,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1652049,
    "dDurationMs": 6901,
    "wWinId": 1,
    "segs": [ {
      "utf8": "he'll",
      "acAsrConf": 186
    }, {
      "utf8": " never",
      "tOffsetMs": 541,
      "acAsrConf": 226
    }, {
      "utf8": " forget",
      "tOffsetMs": 721,
      "acAsrConf": 252
    }, {
      "utf8": " that",
      "tOffsetMs": 1171,
      "acAsrConf": 227
    }, {
      "utf8": " he",
      "tOffsetMs": 1231,
      "acAsrConf": 202
    }, {
      "utf8": " should",
      "tOffsetMs": 1561,
      "acAsrConf": 252
    }, {
      "utf8": " say",
      "tOffsetMs": 1801,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1654110,
    "dDurationMs": 4840,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1654120,
    "dDurationMs": 6990,
    "wWinId": 1,
    "segs": [ {
      "utf8": "wow",
      "acAsrConf": 167
    }, {
      "utf8": " my",
      "tOffsetMs": 750,
      "acAsrConf": 201
    }, {
      "utf8": " favorite",
      "tOffsetMs": 1559,
      "acAsrConf": 252
    }, {
      "utf8": " place",
      "tOffsetMs": 1830,
      "acAsrConf": 200
    }, {
      "utf8": " on",
      "tOffsetMs": 2340,
      "acAsrConf": 252
    }, {
      "utf8": " earth",
      "tOffsetMs": 2750,
      "acAsrConf": 231
    }, {
      "utf8": " is",
      "tOffsetMs": 3750,
      "acAsrConf": 243
    }, {
      "utf8": " the",
      "tOffsetMs": 4170,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1658940,
    "dDurationMs": 2170,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1658950,
    "dDurationMs": 2849,
    "wWinId": 1,
    "segs": [ {
      "utf8": "mountains",
      "acAsrConf": 252
    }, {
      "utf8": " though",
      "tOffsetMs": 479,
      "acAsrConf": 230
    }, {
      "utf8": " because",
      "tOffsetMs": 1260,
      "acAsrConf": 252
    }, {
      "utf8": " we",
      "tOffsetMs": 1500,
      "acAsrConf": 204
    }, {
      "utf8": " had",
      "tOffsetMs": 1800,
      "acAsrConf": 252
    }, {
      "utf8": " fun",
      "tOffsetMs": 1950,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1661100,
    "dDurationMs": 699,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1661110,
    "dDurationMs": 4050,
    "wWinId": 1,
    "segs": [ {
      "utf8": "about",
      "acAsrConf": 226
    }, {
      "utf8": " it",
      "tOffsetMs": 270,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1661789,
    "dDurationMs": 3371,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1661799,
    "dDurationMs": 6031,
    "wWinId": 1,
    "segs": [ {
      "utf8": "ah",
      "acAsrConf": 208
    }, {
      "utf8": " he",
      "tOffsetMs": 291,
      "acAsrConf": 207
    }, {
      "utf8": " I",
      "tOffsetMs": 1321,
      "acAsrConf": 161
    }, {
      "utf8": " think",
      "tOffsetMs": 1741,
      "acAsrConf": 217
    }, {
      "utf8": " he'll",
      "tOffsetMs": 2281,
      "acAsrConf": 252
    }, {
      "utf8": " never",
      "tOffsetMs": 2401,
      "acAsrConf": 252
    }, {
      "utf8": " forget",
      "tOffsetMs": 2581,
      "acAsrConf": 218
    }, {
      "utf8": " it",
      "tOffsetMs": 3000,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 3151,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1665150,
    "dDurationMs": 2680,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1665160,
    "dDurationMs": 4920,
    "wWinId": 1,
    "segs": [ {
      "utf8": "having",
      "acAsrConf": 207
    }, {
      "utf8": " fun",
      "tOffsetMs": 540,
      "acAsrConf": 255
    }, {
      "utf8": " is",
      "tOffsetMs": 750,
      "acAsrConf": 236
    }, {
      "utf8": " really",
      "tOffsetMs": 1019,
      "acAsrConf": 196
    }, {
      "utf8": " important",
      "tOffsetMs": 1620,
      "acAsrConf": 252
    }, {
      "utf8": " maybe",
      "tOffsetMs": 2370,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1667820,
    "dDurationMs": 2260,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1667830,
    "dDurationMs": 4680,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you've",
      "acAsrConf": 236
    }, {
      "utf8": " been",
      "tOffsetMs": 270,
      "acAsrConf": 144
    }, {
      "utf8": " playing",
      "tOffsetMs": 420,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 719,
      "acAsrConf": 238
    }, {
      "utf8": " game",
      "tOffsetMs": 930,
      "acAsrConf": 252
    }, {
      "utf8": " yeah",
      "tOffsetMs": 960,
      "acAsrConf": 200
    }, {
      "utf8": " playing",
      "tOffsetMs": 1770,
      "acAsrConf": 223
    } ]
  }, {
    "tStartMs": 1670070,
    "dDurationMs": 2440,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1670080,
    "dDurationMs": 6000,
    "wWinId": 1,
    "segs": [ {
      "utf8": "a",
      "acAsrConf": 217
    }, {
      "utf8": " game",
      "tOffsetMs": 180,
      "acAsrConf": 199
    }, {
      "utf8": " in",
      "tOffsetMs": 390,
      "acAsrConf": 252
    }, {
      "utf8": " English",
      "tOffsetMs": 599,
      "acAsrConf": 243
    }, {
      "utf8": " yes",
      "tOffsetMs": 960,
      "acAsrConf": 235
    }, {
      "utf8": " yes",
      "tOffsetMs": 1349,
      "acAsrConf": 252
    }, {
      "utf8": " which",
      "tOffsetMs": 1830,
      "acAsrConf": 238
    }, {
      "utf8": " games",
      "tOffsetMs": 2070,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1672500,
    "dDurationMs": 3580,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1672510,
    "dDurationMs": 6390,
    "wWinId": 1,
    "segs": [ {
      "utf8": "video",
      "acAsrConf": 252
    }, {
      "utf8": " games",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " boobies",
      "tOffsetMs": 840,
      "acAsrConf": 0
    }, {
      "utf8": " are",
      "tOffsetMs": 1710,
      "acAsrConf": 252
    }, {
      "utf8": " fun",
      "tOffsetMs": 1919,
      "acAsrConf": 224
    }, {
      "utf8": " you",
      "tOffsetMs": 2190,
      "acAsrConf": 152
    }, {
      "utf8": " could",
      "tOffsetMs": 2669,
      "acAsrConf": 80
    }, {
      "utf8": " go",
      "tOffsetMs": 3330,
      "acAsrConf": 212
    } ]
  }, {
    "tStartMs": 1676070,
    "dDurationMs": 2830,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1676080,
    "dDurationMs": 6270,
    "wWinId": 1,
    "segs": [ {
      "utf8": "to",
      "acAsrConf": 226
    }, {
      "utf8": " a",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " meet-up",
      "tOffsetMs": 690,
      "acAsrConf": 74
    }, {
      "utf8": " and",
      "tOffsetMs": 990,
      "acAsrConf": 187
    }, {
      "utf8": " play",
      "tOffsetMs": 1320,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 1530,
      "acAsrConf": 252
    }, {
      "utf8": " board",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    }, {
      "utf8": " game",
      "tOffsetMs": 1979,
      "acAsrConf": 252
    }, {
      "utf8": " haha",
      "tOffsetMs": 2070,
      "acAsrConf": 0
    } ]
  }, {
    "tStartMs": 1678890,
    "dDurationMs": 3460,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1678900,
    "dDurationMs": 6930,
    "wWinId": 1,
    "segs": [ {
      "utf8": "in",
      "acAsrConf": 218
    }, {
      "utf8": " English",
      "tOffsetMs": 510,
      "acAsrConf": 236
    }, {
      "utf8": " that",
      "tOffsetMs": 870,
      "acAsrConf": 201
    }, {
      "utf8": " would",
      "tOffsetMs": 1710,
      "acAsrConf": 252
    }, {
      "utf8": " be",
      "tOffsetMs": 1830,
      "acAsrConf": 205
    }, {
      "utf8": " very",
      "tOffsetMs": 1980,
      "acAsrConf": 205
    }, {
      "utf8": " fun",
      "tOffsetMs": 2159,
      "acAsrConf": 252
    }, {
      "utf8": " useful",
      "tOffsetMs": 2460,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 1682340,
    "dDurationMs": 3490,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1682350,
    "dDurationMs": 5910,
    "wWinId": 1,
    "segs": [ {
      "utf8": "so",
      "acAsrConf": 236
    }, {
      "utf8": " yeah",
      "tOffsetMs": 530,
      "acAsrConf": 255
    }, {
      "utf8": " keep",
      "tOffsetMs": 1530,
      "acAsrConf": 226
    }, {
      "utf8": " it",
      "tOffsetMs": 1800,
      "acAsrConf": 206
    }, {
      "utf8": " fun",
      "tOffsetMs": 1920,
      "acAsrConf": 216
    }, {
      "utf8": " keep",
      "tOffsetMs": 2130,
      "acAsrConf": 183
    }, {
      "utf8": " it",
      "tOffsetMs": 2579,
      "acAsrConf": 207
    }, {
      "utf8": " fun",
      "tOffsetMs": 2760,
      "acAsrConf": 236
    }, {
      "utf8": " keep",
      "tOffsetMs": 3000,
      "acAsrConf": 164
    }, {
      "utf8": " it",
      "tOffsetMs": 3300,
      "acAsrConf": 206
    } ]
  }, {
    "tStartMs": 1685820,
    "dDurationMs": 2440,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1685830,
    "dDurationMs": 5490,
    "wWinId": 1,
    "segs": [ {
      "utf8": "light",
      "acAsrConf": 252
    }, {
      "utf8": " realize",
      "tOffsetMs": 270,
      "acAsrConf": 201
    }, {
      "utf8": " that",
      "tOffsetMs": 1260,
      "acAsrConf": 157
    }, {
      "utf8": " yes",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1830,
      "acAsrConf": 227
    }, {
      "utf8": " want",
      "tOffsetMs": 2070,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 2250,
      "acAsrConf": 252
    }, {
      "utf8": " be",
      "tOffsetMs": 2280,
      "acAsrConf": 214
    } ]
  }, {
    "tStartMs": 1688250,
    "dDurationMs": 3070,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1688260,
    "dDurationMs": 4890,
    "wWinId": 1,
    "segs": [ {
      "utf8": "serious",
      "acAsrConf": 174
    }, {
      "utf8": " about",
      "tOffsetMs": 450,
      "acAsrConf": 252
    }, {
      "utf8": " studying",
      "tOffsetMs": 539,
      "acAsrConf": 216
    }, {
      "utf8": " English",
      "tOffsetMs": 870,
      "acAsrConf": 207
    }, {
      "utf8": " but",
      "tOffsetMs": 1200,
      "acAsrConf": 163
    }, {
      "utf8": " you",
      "tOffsetMs": 2070,
      "acAsrConf": 219
    } ]
  }, {
    "tStartMs": 1691310,
    "dDurationMs": 1840,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1691320,
    "dDurationMs": 4260,
    "wWinId": 1,
    "segs": [ {
      "utf8": "have",
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 180,
      "acAsrConf": 252
    }, {
      "utf8": " enjoy",
      "tOffsetMs": 330,
      "acAsrConf": 252
    }, {
      "utf8": " yourself",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " too",
      "tOffsetMs": 1140,
      "acAsrConf": 218
    }, {
      "utf8": " or",
      "tOffsetMs": 1349,
      "acAsrConf": 242
    }, {
      "utf8": " else",
      "tOffsetMs": 1530,
      "acAsrConf": 151
    } ]
  }, {
    "tStartMs": 1693140,
    "dDurationMs": 2440,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1693150,
    "dDurationMs": 5159,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you're",
      "acAsrConf": 252
    }, {
      "utf8": " not",
      "tOffsetMs": 210,
      "acAsrConf": 252
    }, {
      "utf8": " going",
      "tOffsetMs": 240,
      "acAsrConf": 232
    }, {
      "utf8": " to",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " continue",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1019,
      "acAsrConf": 201
    }, {
      "utf8": " let's",
      "tOffsetMs": 1830,
      "acAsrConf": 252
    }, {
      "utf8": " go",
      "tOffsetMs": 2190,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1695570,
    "dDurationMs": 2739,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1695580,
    "dDurationMs": 7200,
    "wWinId": 1,
    "segs": [ {
      "utf8": "to",
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 60,
      "acAsrConf": 252
    }, {
      "utf8": " next",
      "tOffsetMs": 599,
      "acAsrConf": 251
    }, {
      "utf8": " step",
      "tOffsetMs": 780,
      "acAsrConf": 202
    }, {
      "utf8": " this",
      "tOffsetMs": 1080,
      "acAsrConf": 236
    }, {
      "utf8": " is",
      "tOffsetMs": 1260,
      "acAsrConf": 251
    }, {
      "utf8": " on",
      "tOffsetMs": 1560,
      "acAsrConf": 255
    }, {
      "utf8": " the",
      "tOffsetMs": 1800,
      "acAsrConf": 252
    }, {
      "utf8": " mind",
      "tOffsetMs": 2160,
      "acAsrConf": 252
    }, {
      "utf8": " map",
      "tOffsetMs": 2460,
      "acAsrConf": 228
    } ]
  }, {
    "tStartMs": 1698299,
    "dDurationMs": 4481,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1698309,
    "dDurationMs": 7291,
    "wWinId": 1,
    "segs": [ {
      "utf8": "this",
      "acAsrConf": 225
    }, {
      "utf8": " is",
      "tOffsetMs": 421,
      "acAsrConf": 252
    }, {
      "utf8": " part",
      "tOffsetMs": 631,
      "acAsrConf": 248
    }, {
      "utf8": " 4",
      "tOffsetMs": 1521,
      "acAsrConf": 147
    }, {
      "utf8": " how",
      "tOffsetMs": 2521,
      "acAsrConf": 247
    }, {
      "utf8": " to",
      "tOffsetMs": 3331,
      "acAsrConf": 252
    }, {
      "utf8": " improve",
      "tOffsetMs": 3411,
      "acAsrConf": 239
    }, {
      "utf8": " your",
      "tOffsetMs": 4411,
      "acAsrConf": 207
    } ]
  }, {
    "tStartMs": 1702770,
    "dDurationMs": 2830,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1702780,
    "dDurationMs": 4279,
    "wWinId": 1,
    "segs": [ {
      "utf8": "English",
      "acAsrConf": 187
    }, {
      "utf8": " even",
      "tOffsetMs": 560,
      "acAsrConf": 153
    }, {
      "utf8": " if",
      "tOffsetMs": 1560,
      "acAsrConf": 247
    }, {
      "utf8": " you're",
      "tOffsetMs": 1769,
      "acAsrConf": 252
    }, {
      "utf8": " really",
      "tOffsetMs": 2070,
      "acAsrConf": 219
    }, {
      "utf8": " busy",
      "tOffsetMs": 2250,
      "acAsrConf": 206
    }, {
      "utf8": " I",
      "tOffsetMs": 2399,
      "acAsrConf": 158
    } ]
  }, {
    "tStartMs": 1705590,
    "dDurationMs": 1469,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1705600,
    "dDurationMs": 5550,
    "wWinId": 1,
    "segs": [ {
      "utf8": "know",
      "acAsrConf": 207
    }, {
      "utf8": " a",
      "tOffsetMs": 90,
      "acAsrConf": 252
    }, {
      "utf8": " lot",
      "tOffsetMs": 209,
      "acAsrConf": 251
    }, {
      "utf8": " of",
      "tOffsetMs": 569,
      "acAsrConf": 109
    }, {
      "utf8": " you",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " are",
      "tOffsetMs": 900,
      "acAsrConf": 252
    }, {
      "utf8": " very",
      "tOffsetMs": 1110,
      "acAsrConf": 252
    }, {
      "utf8": " busy",
      "tOffsetMs": 1140,
      "acAsrConf": 130
    } ]
  }, {
    "tStartMs": 1707049,
    "dDurationMs": 4101,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1707059,
    "dDurationMs": 6581,
    "wWinId": 1,
    "segs": [ {
      "utf8": "everyone's",
      "acAsrConf": 215
    }, {
      "utf8": " busy",
      "tOffsetMs": 1000,
      "acAsrConf": 206
    }, {
      "utf8": " right",
      "tOffsetMs": 1360,
      "acAsrConf": 238
    }, {
      "utf8": " so",
      "tOffsetMs": 1891,
      "acAsrConf": 187
    }, {
      "utf8": " the",
      "tOffsetMs": 2891,
      "acAsrConf": 248
    }, {
      "utf8": " most",
      "tOffsetMs": 3851,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1711140,
    "dDurationMs": 2500,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1711150,
    "dDurationMs": 5790,
    "wWinId": 1,
    "segs": [ {
      "utf8": "important",
      "acAsrConf": 241
    }, {
      "utf8": " thing",
      "tOffsetMs": 540,
      "acAsrConf": 200
    }, {
      "utf8": " for",
      "tOffsetMs": 690,
      "acAsrConf": 252
    }, {
      "utf8": " studying",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " English",
      "tOffsetMs": 1320,
      "acAsrConf": 226
    }, {
      "utf8": " if",
      "tOffsetMs": 1560,
      "acAsrConf": 152
    } ]
  }, {
    "tStartMs": 1713630,
    "dDurationMs": 3310,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1713640,
    "dDurationMs": 5759,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you're",
      "acAsrConf": 242
    }, {
      "utf8": " really",
      "tOffsetMs": 450,
      "acAsrConf": 201
    }, {
      "utf8": " busy",
      "tOffsetMs": 750,
      "acAsrConf": 178
    }, {
      "utf8": " is",
      "tOffsetMs": 930,
      "acAsrConf": 216
    }, {
      "utf8": " what",
      "tOffsetMs": 1260,
      "acAsrConf": 247
    }, {
      "utf8": " Dan",
      "tOffsetMs": 2190,
      "acAsrConf": 252
    }, {
      "utf8": " said",
      "tOffsetMs": 2430,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1716930,
    "dDurationMs": 2469,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1716940,
    "dDurationMs": 4770,
    "wWinId": 1,
    "segs": [ {
      "utf8": "what",
      "acAsrConf": 252
    }, {
      "utf8": " was",
      "tOffsetMs": 119,
      "acAsrConf": 252
    }, {
      "utf8": " that",
      "tOffsetMs": 239,
      "acAsrConf": 252
    }, {
      "utf8": " expression",
      "tOffsetMs": 270,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 750,
      "acAsrConf": 243
    }, {
      "utf8": " used",
      "tOffsetMs": 929,
      "acAsrConf": 201
    }, {
      "utf8": " use",
      "tOffsetMs": 1260,
      "acAsrConf": 195
    }, {
      "utf8": " it",
      "tOffsetMs": 2130,
      "acAsrConf": 223
    } ]
  }, {
    "tStartMs": 1719389,
    "dDurationMs": 2321,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1719399,
    "dDurationMs": 7441,
    "wWinId": 1,
    "segs": [ {
      "utf8": "or",
      "acAsrConf": 223
    }, {
      "utf8": " lose",
      "tOffsetMs": 240,
      "acAsrConf": 226
    }, {
      "utf8": " it",
      "tOffsetMs": 630,
      "acAsrConf": 223
    }, {
      "utf8": " use",
      "tOffsetMs": 660,
      "acAsrConf": 151
    }, {
      "utf8": " it",
      "tOffsetMs": 1081,
      "acAsrConf": 252
    }, {
      "utf8": " or",
      "tOffsetMs": 1230,
      "acAsrConf": 252
    }, {
      "utf8": " lose",
      "tOffsetMs": 1380,
      "acAsrConf": 205
    }, {
      "utf8": " it",
      "tOffsetMs": 1561,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1590,
      "acAsrConf": 209
    }, {
      "utf8": " you",
      "tOffsetMs": 1921,
      "acAsrConf": 214
    } ]
  }, {
    "tStartMs": 1721700,
    "dDurationMs": 5140,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1721710,
    "dDurationMs": 5730,
    "wWinId": 1,
    "segs": [ {
      "utf8": "also",
      "acAsrConf": 252
    }, {
      "utf8": " said",
      "tOffsetMs": 209,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 689,
      "acAsrConf": 236
    }, {
      "utf8": " bite-sized",
      "tOffsetMs": 3140,
      "acAsrConf": 0
    }, {
      "utf8": " chunks",
      "tOffsetMs": 4140,
      "acAsrConf": 108
    }, {
      "utf8": " oh",
      "tOffsetMs": 4860,
      "acAsrConf": 201
    }, {
      "utf8": " yes",
      "tOffsetMs": 4919,
      "acAsrConf": 241
    } ]
  }, {
    "tStartMs": 1726830,
    "dDurationMs": 610,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1726840,
    "dDurationMs": 2909,
    "wWinId": 1,
    "segs": [ {
      "utf8": "yes",
      "acAsrConf": 236
    } ]
  }, {
    "tStartMs": 1727430,
    "dDurationMs": 2319,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1727440,
    "dDurationMs": 4799,
    "wWinId": 1,
    "segs": [ {
      "utf8": "it's",
      "acAsrConf": 201
    }, {
      "utf8": " practice",
      "tOffsetMs": 229,
      "acAsrConf": 249
    }, {
      "utf8": " English",
      "tOffsetMs": 1229,
      "acAsrConf": 244
    }, {
      "utf8": " in",
      "tOffsetMs": 1619,
      "acAsrConf": 164
    }, {
      "utf8": " bite-sized",
      "tOffsetMs": 1799,
      "acAsrConf": 131
    } ]
  }, {
    "tStartMs": 1729739,
    "dDurationMs": 2500,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1729749,
    "dDurationMs": 3841,
    "wWinId": 1,
    "segs": [ {
      "utf8": "chunks",
      "acAsrConf": 139
    }, {
      "utf8": " yes",
      "tOffsetMs": 630,
      "acAsrConf": 200
    }, {
      "utf8": " and",
      "tOffsetMs": 961,
      "acAsrConf": 241
    }, {
      "utf8": " manageable",
      "tOffsetMs": 1260,
      "acAsrConf": 236
    }, {
      "utf8": " something",
      "tOffsetMs": 1951,
      "acAsrConf": 243
    } ]
  }, {
    "tStartMs": 1732229,
    "dDurationMs": 1361,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1732239,
    "dDurationMs": 2971,
    "wWinId": 1,
    "segs": [ {
      "utf8": "sustainable",
      "acAsrConf": 252
    }, {
      "utf8": " yes",
      "tOffsetMs": 510,
      "acAsrConf": 137
    } ]
  }, {
    "tStartMs": 1733580,
    "dDurationMs": 1630,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1733590,
    "dDurationMs": 4140,
    "wWinId": 1,
    "segs": [ {
      "utf8": "something",
      "acAsrConf": 221
    }, {
      "utf8": " sustainable",
      "tOffsetMs": 329,
      "acAsrConf": 168
    }, {
      "utf8": " so",
      "tOffsetMs": 899,
      "acAsrConf": 252
    }, {
      "utf8": " using",
      "tOffsetMs": 1019,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 1380,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1735200,
    "dDurationMs": 2530,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1735210,
    "dDurationMs": 4380,
    "wWinId": 1,
    "segs": [ {
      "utf8": "bite-size",
      "acAsrConf": 141
    }, {
      "utf8": " chunks",
      "tOffsetMs": 480,
      "acAsrConf": 255
    }, {
      "utf8": " small",
      "tOffsetMs": 510,
      "acAsrConf": 205
    }, {
      "utf8": " pieces",
      "tOffsetMs": 1199,
      "acAsrConf": 241
    }, {
      "utf8": " and",
      "tOffsetMs": 1649,
      "acAsrConf": 208
    }, {
      "utf8": " I'm",
      "tOffsetMs": 2039,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1737720,
    "dDurationMs": 1870,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1737730,
    "dDurationMs": 4079,
    "wWinId": 1,
    "segs": [ {
      "utf8": "really",
      "acAsrConf": 233
    }, {
      "utf8": " excited",
      "tOffsetMs": 210,
      "acAsrConf": 154
    }, {
      "utf8": " because",
      "tOffsetMs": 480,
      "acAsrConf": 206
    }, {
      "utf8": " on",
      "tOffsetMs": 750,
      "acAsrConf": 248
    }, {
      "utf8": " the",
      "tOffsetMs": 1199,
      "acAsrConf": 250
    }, {
      "utf8": " course",
      "tOffsetMs": 1620,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1739580,
    "dDurationMs": 2229,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1739590,
    "dDurationMs": 4919,
    "wWinId": 1,
    "segs": [ {
      "utf8": "that",
      "acAsrConf": 203
    }, {
      "utf8": " we",
      "tOffsetMs": 179,
      "acAsrConf": 227
    }, {
      "utf8": " made",
      "tOffsetMs": 329,
      "acAsrConf": 252
    }, {
      "utf8": " together",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " 30",
      "tOffsetMs": 659,
      "acAsrConf": 142
    }, {
      "utf8": " days",
      "tOffsetMs": 1350,
      "acAsrConf": 255
    }, {
      "utf8": " of",
      "tOffsetMs": 1679,
      "acAsrConf": 255
    }, {
      "utf8": " English",
      "tOffsetMs": 1890,
      "acAsrConf": 243
    } ]
  }, {
    "tStartMs": 1741799,
    "dDurationMs": 2710,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1741809,
    "dDurationMs": 4680,
    "wWinId": 1,
    "segs": [ {
      "utf8": "we",
      "acAsrConf": 204
    }, {
      "utf8": " did",
      "tOffsetMs": 511,
      "acAsrConf": 252
    }, {
      "utf8": " all",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 901,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 960,
      "acAsrConf": 246
    }, {
      "utf8": " hard",
      "tOffsetMs": 1350,
      "acAsrConf": 252
    }, {
      "utf8": " work",
      "tOffsetMs": 1381,
      "acAsrConf": 147
    }, {
      "utf8": " for",
      "tOffsetMs": 1560,
      "acAsrConf": 216
    }, {
      "utf8": " you",
      "tOffsetMs": 1740,
      "acAsrConf": 201
    }, {
      "utf8": " so",
      "tOffsetMs": 2220,
      "acAsrConf": 221
    } ]
  }, {
    "tStartMs": 1744499,
    "dDurationMs": 1990,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1744509,
    "dDurationMs": 3841,
    "wWinId": 1,
    "segs": [ {
      "utf8": "that",
      "acAsrConf": 200
    }, {
      "utf8": " you",
      "tOffsetMs": 181,
      "acAsrConf": 187
    }, {
      "utf8": " can",
      "tOffsetMs": 300,
      "acAsrConf": 216
    }, {
      "utf8": " improve",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " your",
      "tOffsetMs": 600,
      "acAsrConf": 203
    }, {
      "utf8": " English",
      "tOffsetMs": 1020,
      "acAsrConf": 196
    }, {
      "utf8": " just",
      "tOffsetMs": 1081,
      "acAsrConf": 164
    }, {
      "utf8": " a",
      "tOffsetMs": 1770,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1746479,
    "dDurationMs": 1871,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1746489,
    "dDurationMs": 3841,
    "wWinId": 1,
    "segs": [ {
      "utf8": "few",
      "acAsrConf": 226
    }, {
      "utf8": " minutes",
      "tOffsetMs": 180,
      "acAsrConf": 221
    }, {
      "utf8": " every",
      "tOffsetMs": 451,
      "acAsrConf": 214
    }, {
      "utf8": " day",
      "tOffsetMs": 660,
      "acAsrConf": 144
    }, {
      "utf8": " while",
      "tOffsetMs": 900,
      "acAsrConf": 201
    }, {
      "utf8": " you're",
      "tOffsetMs": 1680,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1748340,
    "dDurationMs": 1990,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1748350,
    "dDurationMs": 5039,
    "wWinId": 1,
    "segs": [ {
      "utf8": "eating",
      "acAsrConf": 167
    }, {
      "utf8": " breakfast",
      "tOffsetMs": 59,
      "acAsrConf": 216
    }, {
      "utf8": " or",
      "tOffsetMs": 840,
      "acAsrConf": 213
    }, {
      "utf8": " commuting",
      "tOffsetMs": 1110,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1590,
      "acAsrConf": 252
    }, {
      "utf8": " work",
      "tOffsetMs": 1740,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1750320,
    "dDurationMs": 3069,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1750330,
    "dDurationMs": 6059,
    "wWinId": 1,
    "segs": [ {
      "utf8": "it's",
      "acAsrConf": 201
    }, {
      "utf8": " something",
      "tOffsetMs": 599,
      "acAsrConf": 252
    }, {
      "utf8": " that",
      "tOffsetMs": 779,
      "acAsrConf": 132
    }, {
      "utf8": " you",
      "tOffsetMs": 1110,
      "acAsrConf": 252
    }, {
      "utf8": " can",
      "tOffsetMs": 1349,
      "acAsrConf": 252
    }, {
      "utf8": " use",
      "tOffsetMs": 1650,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1949,
      "acAsrConf": 255
    }, {
      "utf8": " you",
      "tOffsetMs": 2279,
      "acAsrConf": 251
    } ]
  }, {
    "tStartMs": 1753379,
    "dDurationMs": 3010,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1753389,
    "dDurationMs": 5821,
    "wWinId": 1,
    "segs": [ {
      "utf8": "won't",
      "acAsrConf": 201
    }, {
      "utf8": " feel",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " stressed",
      "tOffsetMs": 870,
      "acAsrConf": 252
    }, {
      "utf8": " about",
      "tOffsetMs": 1380,
      "acAsrConf": 231
    }, {
      "utf8": " how",
      "tOffsetMs": 1591,
      "acAsrConf": 206
    }, {
      "utf8": " much",
      "tOffsetMs": 2250,
      "acAsrConf": 244
    } ]
  }, {
    "tStartMs": 1756379,
    "dDurationMs": 2831,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1756389,
    "dDurationMs": 7860,
    "wWinId": 1,
    "segs": [ {
      "utf8": "information",
      "acAsrConf": 209
    }, {
      "utf8": " you",
      "tOffsetMs": 660,
      "acAsrConf": 138
    }, {
      "utf8": " need",
      "tOffsetMs": 1201,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1410,
      "acAsrConf": 252
    }, {
      "utf8": " learn",
      "tOffsetMs": 1441,
      "acAsrConf": 203
    }, {
      "utf8": " every",
      "tOffsetMs": 1831,
      "acAsrConf": 234
    }, {
      "utf8": " day",
      "tOffsetMs": 2581,
      "acAsrConf": 241
    } ]
  }, {
    "tStartMs": 1759200,
    "dDurationMs": 5049,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1759210,
    "dDurationMs": 6689,
    "wWinId": 1,
    "segs": [ {
      "utf8": "so",
      "acAsrConf": 234
    }, {
      "utf8": " in",
      "tOffsetMs": 630,
      "acAsrConf": 255
    }, {
      "utf8": " in",
      "tOffsetMs": 1199,
      "acAsrConf": 250
    }, {
      "utf8": " this",
      "tOffsetMs": 2069,
      "acAsrConf": 252
    }, {
      "utf8": " idea",
      "tOffsetMs": 2669,
      "acAsrConf": 239
    }, {
      "utf8": " you",
      "tOffsetMs": 3120,
      "acAsrConf": 202
    }, {
      "utf8": " can",
      "tOffsetMs": 3750,
      "acAsrConf": 252
    }, {
      "utf8": " use",
      "tOffsetMs": 3959,
      "acAsrConf": 227
    }, {
      "utf8": " all",
      "tOffsetMs": 4140,
      "acAsrConf": 252
    }, {
      "utf8": " three",
      "tOffsetMs": 4469,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1764239,
    "dDurationMs": 1660,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1764249,
    "dDurationMs": 3451,
    "wWinId": 1,
    "segs": [ {
      "utf8": "tips",
      "acAsrConf": 252
    }, {
      "utf8": " we're",
      "tOffsetMs": 300,
      "acAsrConf": 252
    }, {
      "utf8": " talking",
      "tOffsetMs": 540,
      "acAsrConf": 252
    }, {
      "utf8": " about",
      "tOffsetMs": 841,
      "acAsrConf": 252
    }, {
      "utf8": " today",
      "tOffsetMs": 961,
      "acAsrConf": 230
    }, {
      "utf8": " in",
      "tOffsetMs": 1201,
      "acAsrConf": 202
    }, {
      "utf8": " the",
      "tOffsetMs": 1471,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1765889,
    "dDurationMs": 1811,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1765899,
    "dDurationMs": 5071,
    "wWinId": 1,
    "segs": [ {
      "utf8": "course",
      "acAsrConf": 252
    }, {
      "utf8": " you'll",
      "tOffsetMs": 210,
      "acAsrConf": 252
    }, {
      "utf8": " learn",
      "tOffsetMs": 421,
      "acAsrConf": 252
    }, {
      "utf8": " speaking",
      "tOffsetMs": 480,
      "acAsrConf": 206
    }, {
      "utf8": " from",
      "tOffsetMs": 1110,
      "acAsrConf": 202
    } ]
  }, {
    "tStartMs": 1767690,
    "dDurationMs": 3280,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1767700,
    "dDurationMs": 6000,
    "wWinId": 1,
    "segs": [ {
      "utf8": "hearing",
      "acAsrConf": 252
    }, {
      "utf8": " natural",
      "tOffsetMs": 539,
      "acAsrConf": 252
    }, {
      "utf8": " conversations",
      "tOffsetMs": 1020,
      "acAsrConf": 252
    }, {
      "utf8": " how",
      "tOffsetMs": 1890,
      "acAsrConf": 226
    }, {
      "utf8": " to",
      "tOffsetMs": 2729,
      "acAsrConf": 252
    }, {
      "utf8": " use",
      "tOffsetMs": 2789,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1770960,
    "dDurationMs": 2740,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1770970,
    "dDurationMs": 4829,
    "wWinId": 1,
    "segs": [ {
      "utf8": "specific",
      "acAsrConf": 136
    }, {
      "utf8": " expressions",
      "tOffsetMs": 689,
      "acAsrConf": 255
    }, {
      "utf8": " from",
      "tOffsetMs": 1620,
      "acAsrConf": 252
    }, {
      "utf8": " conversations",
      "tOffsetMs": 1889,
      "acAsrConf": 233
    } ]
  }, {
    "tStartMs": 1773690,
    "dDurationMs": 2109,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1773700,
    "dDurationMs": 4109,
    "wWinId": 1,
    "segs": [ {
      "utf8": "and",
      "acAsrConf": 228
    }, {
      "utf8": " you'll",
      "tOffsetMs": 209,
      "acAsrConf": 252
    }, {
      "utf8": " get",
      "tOffsetMs": 659,
      "acAsrConf": 217
    }, {
      "utf8": " really",
      "tOffsetMs": 839,
      "acAsrConf": 246
    }, {
      "utf8": " valuable",
      "tOffsetMs": 1319,
      "acAsrConf": 210
    }, {
      "utf8": " feedback",
      "tOffsetMs": 1589,
      "acAsrConf": 203
    } ]
  }, {
    "tStartMs": 1775789,
    "dDurationMs": 2020,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1775799,
    "dDurationMs": 4830,
    "wWinId": 1,
    "segs": [ {
      "utf8": "because",
      "acAsrConf": 252
    }, {
      "utf8": " we",
      "tOffsetMs": 421,
      "acAsrConf": 160
    }, {
      "utf8": " said",
      "tOffsetMs": 570,
      "acAsrConf": 252
    }, {
      "utf8": " that's",
      "tOffsetMs": 810,
      "acAsrConf": 251
    }, {
      "utf8": " really",
      "tOffsetMs": 1260,
      "acAsrConf": 198
    }, {
      "utf8": " important",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1777799,
    "dDurationMs": 2830,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1777809,
    "dDurationMs": 5521,
    "wWinId": 1,
    "segs": [ {
      "utf8": "right",
      "acAsrConf": 229
    }, {
      "utf8": " yeah",
      "tOffsetMs": 151,
      "acAsrConf": 209
    }, {
      "utf8": " getting",
      "tOffsetMs": 1080,
      "acAsrConf": 200
    }, {
      "utf8": " feedback",
      "tOffsetMs": 1500,
      "acAsrConf": 225
    }, {
      "utf8": " is",
      "tOffsetMs": 1740,
      "acAsrConf": 219
    }, {
      "utf8": " so",
      "tOffsetMs": 2250,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1780619,
    "dDurationMs": 2711,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1780629,
    "dDurationMs": 5880,
    "wWinId": 1,
    "segs": [ {
      "utf8": "important",
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 510,
      "acAsrConf": 183
    }, {
      "utf8": " having",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    }, {
      "utf8": " someone",
      "tOffsetMs": 1500,
      "acAsrConf": 218
    }, {
      "utf8": " to",
      "tOffsetMs": 1920,
      "acAsrConf": 214
    }, {
      "utf8": " tell",
      "tOffsetMs": 2250,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 2490,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1783320,
    "dDurationMs": 3189,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1783330,
    "dDurationMs": 5279,
    "wWinId": 1,
    "segs": [ {
      "utf8": "hey",
      "acAsrConf": 240
    }, {
      "utf8": " it",
      "tOffsetMs": 630,
      "acAsrConf": 205
    }, {
      "utf8": " sounds",
      "tOffsetMs": 1219,
      "acAsrConf": 227
    }, {
      "utf8": " a",
      "tOffsetMs": 2219,
      "acAsrConf": 252
    }, {
      "utf8": " little",
      "tOffsetMs": 2279,
      "acAsrConf": 252
    }, {
      "utf8": " strange",
      "tOffsetMs": 2370,
      "acAsrConf": 134
    }, {
      "utf8": " when",
      "tOffsetMs": 2849,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 3059,
      "acAsrConf": 200
    } ]
  }, {
    "tStartMs": 1786499,
    "dDurationMs": 2110,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1786509,
    "dDurationMs": 3841,
    "wWinId": 1,
    "segs": [ {
      "utf8": "say",
      "acAsrConf": 252
    }, {
      "utf8": " that",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 451,
      "acAsrConf": 226
    }, {
      "utf8": " should",
      "tOffsetMs": 811,
      "acAsrConf": 252
    }, {
      "utf8": " say",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    }, {
      "utf8": " this",
      "tOffsetMs": 1230,
      "acAsrConf": 252
    }, {
      "utf8": " instead",
      "tOffsetMs": 1410,
      "acAsrConf": 212
    }, {
      "utf8": " so",
      "tOffsetMs": 1800,
      "acAsrConf": 194
    } ]
  }, {
    "tStartMs": 1788599,
    "dDurationMs": 1751,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1788609,
    "dDurationMs": 4410,
    "wWinId": 1,
    "segs": [ {
      "utf8": "in",
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 241,
      "acAsrConf": 252
    }, {
      "utf8": " course",
      "tOffsetMs": 450,
      "acAsrConf": 237
    }, {
      "utf8": " thirty",
      "tOffsetMs": 721,
      "acAsrConf": 152
    }, {
      "utf8": " days",
      "tOffsetMs": 1020,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 1231,
      "acAsrConf": 252
    }, {
      "utf8": " English",
      "tOffsetMs": 1380,
      "acAsrConf": 203
    } ]
  }, {
    "tStartMs": 1790340,
    "dDurationMs": 2679,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1790350,
    "dDurationMs": 3959,
    "wWinId": 1,
    "segs": [ {
      "utf8": "that",
      "acAsrConf": 145
    }, {
      "utf8": " we",
      "tOffsetMs": 209,
      "acAsrConf": 252
    }, {
      "utf8": " put",
      "tOffsetMs": 329,
      "acAsrConf": 252
    }, {
      "utf8": " together",
      "tOffsetMs": 569,
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1019,
      "acAsrConf": 209
    }, {
      "utf8": " we",
      "tOffsetMs": 1559,
      "acAsrConf": 216
    }, {
      "utf8": " will",
      "tOffsetMs": 2490,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1793009,
    "dDurationMs": 1300,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1793019,
    "dDurationMs": 2610,
    "wWinId": 1,
    "segs": [ {
      "utf8": "give",
      "acAsrConf": 207
    }, {
      "utf8": " you",
      "tOffsetMs": 181,
      "acAsrConf": 226
    }, {
      "utf8": " feedback",
      "tOffsetMs": 301,
      "acAsrConf": 252
    }, {
      "utf8": " you'll",
      "tOffsetMs": 540,
      "acAsrConf": 219
    }, {
      "utf8": " hear",
      "tOffsetMs": 1110,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1794299,
    "dDurationMs": 1330,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1794309,
    "dDurationMs": 4261,
    "wWinId": 1,
    "segs": [ {
      "utf8": "conversations",
      "acAsrConf": 252
    }, {
      "utf8": " you'll",
      "tOffsetMs": 781,
      "acAsrConf": 249
    }, {
      "utf8": " learn",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    }, {
      "utf8": " new",
      "tOffsetMs": 1171,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1795619,
    "dDurationMs": 2951,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1795629,
    "dDurationMs": 4351,
    "wWinId": 1,
    "segs": [ {
      "utf8": "expressions",
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 571,
      "acAsrConf": 197
    }, {
      "utf8": " that's",
      "tOffsetMs": 780,
      "acAsrConf": 233
    }, {
      "utf8": " so",
      "tOffsetMs": 1711,
      "acAsrConf": 236
    }, {
      "utf8": " relieving",
      "tOffsetMs": 2370,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1798560,
    "dDurationMs": 1420,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1798570,
    "dDurationMs": 3839,
    "wWinId": 1,
    "segs": [ {
      "utf8": "because",
      "acAsrConf": 202
    }, {
      "utf8": " you",
      "tOffsetMs": 479,
      "acAsrConf": 252
    }, {
      "utf8": " don't",
      "tOffsetMs": 630,
      "acAsrConf": 233
    }, {
      "utf8": " need",
      "tOffsetMs": 809,
      "acAsrConf": 217
    }, {
      "utf8": " to",
      "tOffsetMs": 929,
      "acAsrConf": 214
    }, {
      "utf8": " worry",
      "tOffsetMs": 1079,
      "acAsrConf": 252
    }, {
      "utf8": " about",
      "tOffsetMs": 1319,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1799970,
    "dDurationMs": 2439,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1799980,
    "dDurationMs": 5519,
    "wWinId": 1,
    "segs": [ {
      "utf8": "what",
      "acAsrConf": 227
    }, {
      "utf8": " you",
      "tOffsetMs": 449,
      "acAsrConf": 220
    }, {
      "utf8": " should",
      "tOffsetMs": 689,
      "acAsrConf": 71
    }, {
      "utf8": " be",
      "tOffsetMs": 929,
      "acAsrConf": 201
    }, {
      "utf8": " doing",
      "tOffsetMs": 1019,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1049,
      "acAsrConf": 124
    }, {
      "utf8": " I",
      "tOffsetMs": 1500,
      "acAsrConf": 245
    }, {
      "utf8": " have",
      "tOffsetMs": 2220,
      "acAsrConf": 229
    }, {
      "utf8": " a",
      "tOffsetMs": 2399,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1802399,
    "dDurationMs": 3100,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1802409,
    "dDurationMs": 5760,
    "wWinId": 1,
    "segs": [ {
      "utf8": "special",
      "acAsrConf": 186
    }, {
      "utf8": " exciting",
      "tOffsetMs": 600,
      "acAsrConf": 241
    }, {
      "utf8": " thing",
      "tOffsetMs": 1321,
      "acAsrConf": 208
    }, {
      "utf8": " to",
      "tOffsetMs": 1770,
      "acAsrConf": 146
    }, {
      "utf8": " share",
      "tOffsetMs": 1980,
      "acAsrConf": 252
    }, {
      "utf8": " with",
      "tOffsetMs": 2370,
      "acAsrConf": 195
    }, {
      "utf8": " you",
      "tOffsetMs": 2431,
      "acAsrConf": 166
    } ]
  }, {
    "tStartMs": 1805489,
    "dDurationMs": 2680,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1805499,
    "dDurationMs": 4860,
    "wWinId": 1,
    "segs": [ {
      "utf8": "yes",
      "acAsrConf": 252
    }, {
      "utf8": " because",
      "tOffsetMs": 630,
      "acAsrConf": 200
    }, {
      "utf8": " you",
      "tOffsetMs": 1110,
      "acAsrConf": 165
    }, {
      "utf8": " are",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    }, {
      "utf8": " here",
      "tOffsetMs": 1591,
      "acAsrConf": 202
    }, {
      "utf8": " wonderful",
      "tOffsetMs": 2130,
      "acAsrConf": 249
    } ]
  }, {
    "tStartMs": 1808159,
    "dDurationMs": 2200,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1808169,
    "dDurationMs": 5250,
    "wWinId": 1,
    "segs": [ {
      "utf8": "people",
      "acAsrConf": 172
    }, {
      "utf8": " who",
      "tOffsetMs": 330,
      "acAsrConf": 249
    }, {
      "utf8": " have",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " asked",
      "tOffsetMs": 630,
      "acAsrConf": 208
    }, {
      "utf8": " great",
      "tOffsetMs": 1140,
      "acAsrConf": 252
    }, {
      "utf8": " questions",
      "tOffsetMs": 1350,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 1380,
      "acAsrConf": 207
    } ]
  }, {
    "tStartMs": 1810349,
    "dDurationMs": 3070,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1810359,
    "dDurationMs": 5280,
    "wWinId": 1,
    "segs": [ {
      "utf8": "this",
      "acAsrConf": 218
    }, {
      "utf8": " course",
      "tOffsetMs": 690,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 1020,
      "acAsrConf": 226
    }, {
      "utf8": " have",
      "tOffsetMs": 1231,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 1800,
      "acAsrConf": 252
    }, {
      "utf8": " special",
      "tOffsetMs": 1831,
      "acAsrConf": 233
    }, {
      "utf8": " gift",
      "tOffsetMs": 2371,
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 2700,
      "acAsrConf": 160
    } ]
  }, {
    "tStartMs": 1813409,
    "dDurationMs": 2230,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1813419,
    "dDurationMs": 5580,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 235
    }, {
      "utf8": " some",
      "tOffsetMs": 61,
      "acAsrConf": 201
    }, {
      "utf8": " free",
      "tOffsetMs": 570,
      "acAsrConf": 248
    }, {
      "utf8": " things",
      "tOffsetMs": 990,
      "acAsrConf": 228
    }, {
      "utf8": " I",
      "tOffsetMs": 1320,
      "acAsrConf": 244
    }, {
      "utf8": " want",
      "tOffsetMs": 1531,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1771,
      "acAsrConf": 252
    }, {
      "utf8": " offer",
      "tOffsetMs": 1801,
      "acAsrConf": 201
    }, {
      "utf8": " you",
      "tOffsetMs": 1980,
      "acAsrConf": 237
    } ]
  }, {
    "tStartMs": 1815629,
    "dDurationMs": 3370,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1815639,
    "dDurationMs": 7441,
    "wWinId": 1,
    "segs": [ {
      "utf8": "so",
      "acAsrConf": 207
    }, {
      "utf8": " usually",
      "tOffsetMs": 841,
      "acAsrConf": 237
    }, {
      "utf8": " in",
      "tOffsetMs": 1470,
      "acAsrConf": 252
    }, {
      "utf8": " our",
      "tOffsetMs": 1681,
      "acAsrConf": 216
    }, {
      "utf8": " 30",
      "tOffsetMs": 2071,
      "acAsrConf": 167
    }, {
      "utf8": " days",
      "tOffsetMs": 2610,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 2850,
      "acAsrConf": 235
    }, {
      "utf8": " English",
      "tOffsetMs": 3030,
      "acAsrConf": 225
    } ]
  }, {
    "tStartMs": 1818989,
    "dDurationMs": 4091,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1818999,
    "dDurationMs": 6091,
    "wWinId": 1,
    "segs": [ {
      "utf8": "course",
      "acAsrConf": 252
    }, {
      "utf8": " there",
      "tOffsetMs": 360,
      "acAsrConf": 249
    }, {
      "utf8": " are",
      "tOffsetMs": 1170,
      "acAsrConf": 252
    }, {
      "utf8": " how",
      "tOffsetMs": 2211,
      "acAsrConf": 252
    }, {
      "utf8": " many",
      "tOffsetMs": 3211,
      "acAsrConf": 252
    }, {
      "utf8": " conversation",
      "tOffsetMs": 3240,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1823070,
    "dDurationMs": 2020,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1823080,
    "dDurationMs": 4679,
    "wWinId": 1,
    "segs": [ {
      "utf8": "videos",
      "acAsrConf": 252
    }, {
      "utf8": " are",
      "tOffsetMs": 360,
      "acAsrConf": 252
    }, {
      "utf8": " there",
      "tOffsetMs": 569,
      "acAsrConf": 226
    }, {
      "utf8": " there's",
      "tOffsetMs": 750,
      "acAsrConf": 226
    }, {
      "utf8": " 12",
      "tOffsetMs": 1409,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 1825080,
    "dDurationMs": 2679,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1825090,
    "dDurationMs": 6390,
    "wWinId": 1,
    "segs": [ {
      "utf8": "yeah",
      "acAsrConf": 235
    }, {
      "utf8": " there's",
      "tOffsetMs": 120,
      "acAsrConf": 252
    }, {
      "utf8": " 12",
      "tOffsetMs": 329,
      "acAsrConf": 208
    }, {
      "utf8": " conversation",
      "tOffsetMs": 630,
      "acAsrConf": 252
    }, {
      "utf8": " videos",
      "tOffsetMs": 1380,
      "acAsrConf": 251
    }, {
      "utf8": " 12",
      "tOffsetMs": 1769,
      "acAsrConf": 149
    } ]
  }, {
    "tStartMs": 1827749,
    "dDurationMs": 3731,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1827759,
    "dDurationMs": 7800,
    "wWinId": 1,
    "segs": [ {
      "utf8": "expression",
      "acAsrConf": 238
    }, {
      "utf8": " videos",
      "tOffsetMs": 870,
      "acAsrConf": 252
    }, {
      "utf8": " where",
      "tOffsetMs": 1230,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1711,
      "acAsrConf": 252
    }, {
      "utf8": " can",
      "tOffsetMs": 1740,
      "acAsrConf": 200
    }, {
      "utf8": " learn",
      "tOffsetMs": 2010,
      "acAsrConf": 146
    }, {
      "utf8": " 24",
      "tOffsetMs": 2721,
      "acAsrConf": 240
    } ]
  }, {
    "tStartMs": 1831470,
    "dDurationMs": 4089,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1831480,
    "dDurationMs": 6149,
    "wWinId": 1,
    "segs": [ {
      "utf8": "new",
      "acAsrConf": 252
    }, {
      "utf8": " expressions",
      "tOffsetMs": 299,
      "acAsrConf": 252
    }, {
      "utf8": " with",
      "tOffsetMs": 1019,
      "acAsrConf": 252
    }, {
      "utf8": " us",
      "tOffsetMs": 1169,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1350,
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 2210,
      "acAsrConf": 248
    }, {
      "utf8": " quick",
      "tOffsetMs": 3210,
      "acAsrConf": 217
    } ]
  }, {
    "tStartMs": 1835549,
    "dDurationMs": 2080,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1835559,
    "dDurationMs": 4620,
    "wWinId": 1,
    "segs": [ {
      "utf8": "response",
      "acAsrConf": 252
    }, {
      "utf8": " videos",
      "tOffsetMs": 240,
      "acAsrConf": 251
    }, {
      "utf8": " so",
      "tOffsetMs": 901,
      "acAsrConf": 142
    }, {
      "utf8": " those",
      "tOffsetMs": 1171,
      "acAsrConf": 252
    }, {
      "utf8": " are",
      "tOffsetMs": 1381,
      "acAsrConf": 252
    }, {
      "utf8": " going",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1710,
      "acAsrConf": 194
    }, {
      "utf8": " be",
      "tOffsetMs": 1921,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1837619,
    "dDurationMs": 2560,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1837629,
    "dDurationMs": 5490,
    "wWinId": 1,
    "segs": [ {
      "utf8": "speaking",
      "acAsrConf": 252
    }, {
      "utf8": " videos",
      "tOffsetMs": 420,
      "acAsrConf": 237
    }, {
      "utf8": " where",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1890,
      "acAsrConf": 252
    }, {
      "utf8": " can",
      "tOffsetMs": 1920,
      "acAsrConf": 202
    }, {
      "utf8": " answer",
      "tOffsetMs": 2101,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 1840169,
    "dDurationMs": 2950,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1840179,
    "dDurationMs": 5281,
    "wWinId": 1,
    "segs": [ {
      "utf8": "some",
      "acAsrConf": 206
    }, {
      "utf8": " questions",
      "tOffsetMs": 360,
      "acAsrConf": 236
    }, {
      "utf8": " and",
      "tOffsetMs": 870,
      "acAsrConf": 208
    }, {
      "utf8": " record",
      "tOffsetMs": 1080,
      "acAsrConf": 215
    }, {
      "utf8": " your",
      "tOffsetMs": 2041,
      "acAsrConf": 219
    }, {
      "utf8": " voice",
      "tOffsetMs": 2281,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 2521,
      "acAsrConf": 227
    } ]
  }, {
    "tStartMs": 1843109,
    "dDurationMs": 2351,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1843119,
    "dDurationMs": 4321,
    "wWinId": 1,
    "segs": [ {
      "utf8": "then",
      "acAsrConf": 234
    }, {
      "utf8": " send",
      "tOffsetMs": 390,
      "acAsrConf": 244
    }, {
      "utf8": " it",
      "tOffsetMs": 721,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 841,
      "acAsrConf": 126
    }, {
      "utf8": " me",
      "tOffsetMs": 1020,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1201,
      "acAsrConf": 212
    }, {
      "utf8": " I'll",
      "tOffsetMs": 1410,
      "acAsrConf": 243
    }, {
      "utf8": " give",
      "tOffsetMs": 1890,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 2221,
      "acAsrConf": 241
    } ]
  }, {
    "tStartMs": 1845450,
    "dDurationMs": 1990,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1845460,
    "dDurationMs": 3779,
    "wWinId": 1,
    "segs": [ {
      "utf8": "some",
      "acAsrConf": 250
    }, {
      "utf8": " of",
      "tOffsetMs": 240,
      "acAsrConf": 203
    }, {
      "utf8": " that",
      "tOffsetMs": 360,
      "acAsrConf": 135
    }, {
      "utf8": " personal",
      "tOffsetMs": 510,
      "acAsrConf": 186
    }, {
      "utf8": " feedback",
      "tOffsetMs": 1169,
      "acAsrConf": 252
    }, {
      "utf8": " we",
      "tOffsetMs": 1439,
      "acAsrConf": 217
    }, {
      "utf8": " were",
      "tOffsetMs": 1889,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1847430,
    "dDurationMs": 1809,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1847440,
    "dDurationMs": 3420,
    "wWinId": 1,
    "segs": [ {
      "utf8": "talking",
      "acAsrConf": 249
    }, {
      "utf8": " about",
      "tOffsetMs": 359,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 449,
      "acAsrConf": 172
    }, {
      "utf8": " you'll",
      "tOffsetMs": 1349,
      "acAsrConf": 252
    }, {
      "utf8": " get",
      "tOffsetMs": 1500,
      "acAsrConf": 208
    }, {
      "utf8": " that",
      "tOffsetMs": 1679,
      "acAsrConf": 217
    } ]
  }, {
    "tStartMs": 1849229,
    "dDurationMs": 1631,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1849239,
    "dDurationMs": 3361,
    "wWinId": 1,
    "segs": [ {
      "utf8": "personalized",
      "acAsrConf": 251
    }, {
      "utf8": " feedback",
      "tOffsetMs": 601,
      "acAsrConf": 226
    }, {
      "utf8": " but",
      "tOffsetMs": 870,
      "acAsrConf": 230
    } ]
  }, {
    "tStartMs": 1850850,
    "dDurationMs": 1750,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1850860,
    "dDurationMs": 6810,
    "wWinId": 1,
    "segs": [ {
      "utf8": "have",
      "acAsrConf": 235
    }, {
      "utf8": " some",
      "tOffsetMs": 180,
      "acAsrConf": 252
    }, {
      "utf8": " free",
      "tOffsetMs": 360,
      "acAsrConf": 252
    }, {
      "utf8": " things",
      "tOffsetMs": 569,
      "acAsrConf": 252
    }, {
      "utf8": " I'd",
      "tOffsetMs": 810,
      "acAsrConf": 200
    }, {
      "utf8": " love",
      "tOffsetMs": 1110,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1350,
      "acAsrConf": 202
    }, {
      "utf8": " share",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1852590,
    "dDurationMs": 5080,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1852600,
    "dDurationMs": 8190,
    "wWinId": 1,
    "segs": [ {
      "utf8": "with",
      "acAsrConf": 208
    }, {
      "utf8": " you",
      "tOffsetMs": 60,
      "acAsrConf": 167
    }, {
      "utf8": " so",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 1530,
      "acAsrConf": 243
    }, {
      "utf8": " the",
      "tOffsetMs": 2130,
      "acAsrConf": 252
    }, {
      "utf8": " next",
      "tOffsetMs": 2430,
      "acAsrConf": 252
    }, {
      "utf8": " 24",
      "tOffsetMs": 2600,
      "acAsrConf": 215
    }, {
      "utf8": " hours",
      "tOffsetMs": 3709,
      "acAsrConf": 107
    }, {
      "utf8": " from",
      "tOffsetMs": 4709,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1857660,
    "dDurationMs": 3130,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1857670,
    "dDurationMs": 6240,
    "wWinId": 1,
    "segs": [ {
      "utf8": "now",
      "acAsrConf": 206
    }, {
      "utf8": " about",
      "tOffsetMs": 120,
      "acAsrConf": 167
    }, {
      "utf8": " 24",
      "tOffsetMs": 780,
      "acAsrConf": 239
    }, {
      "utf8": " hours",
      "tOffsetMs": 1320,
      "acAsrConf": 252
    }, {
      "utf8": " I'm",
      "tOffsetMs": 1440,
      "acAsrConf": 213
    }, {
      "utf8": " going",
      "tOffsetMs": 2370,
      "acAsrConf": 233
    }, {
      "utf8": " to",
      "tOffsetMs": 2850,
      "acAsrConf": 220
    }, {
      "utf8": " share",
      "tOffsetMs": 2940,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1860780,
    "dDurationMs": 3130,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1860790,
    "dDurationMs": 5250,
    "wWinId": 1,
    "segs": [ {
      "utf8": "with",
      "acAsrConf": 204
    }, {
      "utf8": " you",
      "tOffsetMs": 210,
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 330,
      "acAsrConf": 252
    }, {
      "utf8": " free",
      "tOffsetMs": 600,
      "acAsrConf": 206
    }, {
      "utf8": " a",
      "tOffsetMs": 870,
      "acAsrConf": 226
    }, {
      "utf8": " super",
      "tOffsetMs": 1550,
      "acAsrConf": 252
    }, {
      "utf8": " English",
      "tOffsetMs": 2550,
      "acAsrConf": 229
    } ]
  }, {
    "tStartMs": 1863900,
    "dDurationMs": 2140,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1863910,
    "dDurationMs": 7230,
    "wWinId": 1,
    "segs": [ {
      "utf8": "membership",
      "acAsrConf": 235
    }, {
      "utf8": " so",
      "tOffsetMs": 690,
      "acAsrConf": 245
    }, {
      "utf8": " super",
      "tOffsetMs": 899,
      "acAsrConf": 239
    }, {
      "utf8": " English",
      "tOffsetMs": 1260,
      "acAsrConf": 236
    }, {
      "utf8": " membership",
      "tOffsetMs": 1590,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1866030,
    "dDurationMs": 5110,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1866040,
    "dDurationMs": 9150,
    "wWinId": 1,
    "segs": [ {
      "utf8": "is",
      "acAsrConf": 251
    }, {
      "utf8": " for",
      "tOffsetMs": 440,
      "acAsrConf": 252
    }, {
      "utf8": " extra",
      "tOffsetMs": 1460,
      "acAsrConf": 245
    }, {
      "utf8": " conversation",
      "tOffsetMs": 2460,
      "acAsrConf": 248
    }, {
      "utf8": " videos",
      "tOffsetMs": 3210,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 3630,
      "acAsrConf": 246
    }, {
      "utf8": " for",
      "tOffsetMs": 4100,
      "acAsrConf": 226
    } ]
  }, {
    "tStartMs": 1871130,
    "dDurationMs": 4060,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1871140,
    "dDurationMs": 7080,
    "wWinId": 1,
    "segs": [ {
      "utf8": "expression",
      "acAsrConf": 246
    }, {
      "utf8": " guide",
      "tOffsetMs": 990,
      "acAsrConf": 240
    }, {
      "utf8": " downloads",
      "tOffsetMs": 1380,
      "acAsrConf": 236
    }, {
      "utf8": " and",
      "tOffsetMs": 2280,
      "acAsrConf": 252
    }, {
      "utf8": " then",
      "tOffsetMs": 2640,
      "acAsrConf": 252
    }, {
      "utf8": " mp3s",
      "tOffsetMs": 3150,
      "acAsrConf": 243
    } ]
  }, {
    "tStartMs": 1875180,
    "dDurationMs": 3040,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1875190,
    "dDurationMs": 5940,
    "wWinId": 1,
    "segs": [ {
      "utf8": "mp3s",
      "acAsrConf": 231
    }, {
      "utf8": " is",
      "tOffsetMs": 630,
      "acAsrConf": 166
    }, {
      "utf8": " a",
      "tOffsetMs": 840,
      "acAsrConf": 252
    }, {
      "utf8": " new",
      "tOffsetMs": 869,
      "acAsrConf": 252
    }, {
      "utf8": " addition",
      "tOffsetMs": 1430,
      "acAsrConf": 235
    }, {
      "utf8": " for",
      "tOffsetMs": 2430,
      "acAsrConf": 65
    }, {
      "utf8": " the",
      "tOffsetMs": 2700,
      "acAsrConf": 252
    }, {
      "utf8": " course",
      "tOffsetMs": 2760,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1878210,
    "dDurationMs": 2920,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1878220,
    "dDurationMs": 6690,
    "wWinId": 1,
    "segs": [ {
      "utf8": "and",
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 270,
      "acAsrConf": 236
    }, {
      "utf8": " mp3s",
      "tOffsetMs": 959,
      "acAsrConf": 200
    }, {
      "utf8": " are",
      "tOffsetMs": 1530,
      "acAsrConf": 201
    }, {
      "utf8": " where",
      "tOffsetMs": 1800,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 2160,
      "acAsrConf": 252
    }, {
      "utf8": " can",
      "tOffsetMs": 2430,
      "acAsrConf": 237
    }, {
      "utf8": " download",
      "tOffsetMs": 2670,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1881120,
    "dDurationMs": 3790,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1881130,
    "dDurationMs": 5850,
    "wWinId": 1,
    "segs": [ {
      "utf8": "the",
      "acAsrConf": 236
    }, {
      "utf8": " mp3",
      "tOffsetMs": 900,
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 1320,
      "acAsrConf": 252
    }, {
      "utf8": " each",
      "tOffsetMs": 1770,
      "acAsrConf": 230
    }, {
      "utf8": " video",
      "tOffsetMs": 1980,
      "acAsrConf": 207
    }, {
      "utf8": " and",
      "tOffsetMs": 2460,
      "acAsrConf": 226
    }, {
      "utf8": " listen",
      "tOffsetMs": 2789,
      "acAsrConf": 215
    }, {
      "utf8": " to",
      "tOffsetMs": 3510,
      "acAsrConf": 238
    }, {
      "utf8": " it",
      "tOffsetMs": 3660,
      "acAsrConf": 248
    } ]
  }, {
    "tStartMs": 1884900,
    "dDurationMs": 2080,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1884910,
    "dDurationMs": 6800,
    "wWinId": 1,
    "segs": [ {
      "utf8": "so",
      "acAsrConf": 236
    }, {
      "utf8": " if",
      "tOffsetMs": 269,
      "acAsrConf": 237
    }, {
      "utf8": " you",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 870,
      "acAsrConf": 227
    }, {
      "utf8": " time",
      "tOffsetMs": 1080,
      "acAsrConf": 219
    }, {
      "utf8": " to",
      "tOffsetMs": 1320,
      "acAsrConf": 252
    }, {
      "utf8": " watch",
      "tOffsetMs": 1560,
      "acAsrConf": 220
    }, {
      "utf8": " the",
      "tOffsetMs": 1740,
      "acAsrConf": 252
    }, {
      "utf8": " video",
      "tOffsetMs": 1950,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1886970,
    "dDurationMs": 4740,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1886980,
    "dDurationMs": 7590,
    "wWinId": 1,
    "segs": [ {
      "utf8": "great",
      "acAsrConf": 147
    }, {
      "utf8": " but",
      "tOffsetMs": 810,
      "acAsrConf": 217
    }, {
      "utf8": " if",
      "tOffsetMs": 1590,
      "acAsrConf": 250
    }, {
      "utf8": " you",
      "tOffsetMs": 2220,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 2490,
      "acAsrConf": 227
    }, {
      "utf8": " just",
      "tOffsetMs": 3079,
      "acAsrConf": 211
    }, {
      "utf8": " time",
      "tOffsetMs": 4079,
      "acAsrConf": 249
    }, {
      "utf8": " to",
      "tOffsetMs": 4650,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1891700,
    "dDurationMs": 2870,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1891710,
    "dDurationMs": 6730,
    "wWinId": 1,
    "segs": [ {
      "utf8": "listen",
      "acAsrConf": 191
    }, {
      "utf8": " to",
      "tOffsetMs": 1000,
      "acAsrConf": 188
    }, {
      "utf8": " an",
      "tOffsetMs": 1120,
      "acAsrConf": 252
    }, {
      "utf8": " mp3",
      "tOffsetMs": 1240,
      "acAsrConf": 252
    }, {
      "utf8": " while",
      "tOffsetMs": 1599,
      "acAsrConf": 252
    }, {
      "utf8": " you're",
      "tOffsetMs": 2080,
      "acAsrConf": 252
    }, {
      "utf8": " waiting",
      "tOffsetMs": 2290,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 2709,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1894560,
    "dDurationMs": 3880,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1894570,
    "dDurationMs": 7020,
    "wWinId": 1,
    "segs": [ {
      "utf8": "line",
      "acAsrConf": 252
    }, {
      "utf8": " or",
      "tOffsetMs": 30,
      "acAsrConf": 202
    }, {
      "utf8": " while",
      "tOffsetMs": 750,
      "acAsrConf": 255
    }, {
      "utf8": " you",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    }, {
      "utf8": " are",
      "tOffsetMs": 1979,
      "acAsrConf": 252
    }, {
      "utf8": " falling",
      "tOffsetMs": 2010,
      "acAsrConf": 205
    }, {
      "utf8": " asleep",
      "tOffsetMs": 2820,
      "acAsrConf": 244
    }, {
      "utf8": " or",
      "tOffsetMs": 3120,
      "acAsrConf": 226
    } ]
  }, {
    "tStartMs": 1898430,
    "dDurationMs": 3160,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1898440,
    "dDurationMs": 5130,
    "wWinId": 1,
    "segs": [ {
      "utf8": "listening",
      "acAsrConf": 236
    }, {
      "utf8": " to",
      "tOffsetMs": 510,
      "acAsrConf": 249
    }, {
      "utf8": " my",
      "tOffsetMs": 570,
      "acAsrConf": 172
    }, {
      "utf8": " voice",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 1290,
      "acAsrConf": 238
    }, {
      "utf8": " your",
      "tOffsetMs": 1650,
      "acAsrConf": 252
    }, {
      "utf8": " dreams",
      "tOffsetMs": 1770,
      "acAsrConf": 239
    }, {
      "utf8": " you",
      "tOffsetMs": 2190,
      "acAsrConf": 207
    } ]
  }, {
    "tStartMs": 1901580,
    "dDurationMs": 1990,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1901590,
    "dDurationMs": 4140,
    "wWinId": 1,
    "segs": [ {
      "utf8": "can",
      "acAsrConf": 139
    }, {
      "utf8": " listen",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 420,
      "acAsrConf": 200
    }, {
      "utf8": " the",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " mp3",
      "tOffsetMs": 750,
      "acAsrConf": 136
    }, {
      "utf8": " so",
      "tOffsetMs": 1140,
      "acAsrConf": 252
    }, {
      "utf8": " it's",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    }, {
      "utf8": " another",
      "tOffsetMs": 1740,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1903560,
    "dDurationMs": 2170,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1903570,
    "dDurationMs": 5160,
    "wWinId": 1,
    "segs": [ {
      "utf8": "addition",
      "acAsrConf": 74
    }, {
      "utf8": " for",
      "tOffsetMs": 599,
      "acAsrConf": 240
    }, {
      "utf8": " super",
      "tOffsetMs": 930,
      "acAsrConf": 215
    }, {
      "utf8": " English",
      "tOffsetMs": 1260,
      "acAsrConf": 247
    }, {
      "utf8": " members",
      "tOffsetMs": 1560,
      "acAsrConf": 228
    }, {
      "utf8": " that",
      "tOffsetMs": 1920,
      "acAsrConf": 255
    } ]
  }, {
    "tStartMs": 1905720,
    "dDurationMs": 3010,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1905730,
    "dDurationMs": 5850,
    "wWinId": 1,
    "segs": [ {
      "utf8": "I",
      "acAsrConf": 224
    }, {
      "utf8": " want",
      "tOffsetMs": 210,
      "acAsrConf": 243
    }, {
      "utf8": " to",
      "tOffsetMs": 420,
      "acAsrConf": 252
    }, {
      "utf8": " give",
      "tOffsetMs": 540,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 720,
      "acAsrConf": 200
    }, {
      "utf8": " for",
      "tOffsetMs": 870,
      "acAsrConf": 252
    }, {
      "utf8": " free",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1530,
      "acAsrConf": 199
    }, {
      "utf8": " usually",
      "tOffsetMs": 2000,
      "acAsrConf": 227
    } ]
  }, {
    "tStartMs": 1908720,
    "dDurationMs": 2860,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1908730,
    "dDurationMs": 5340,
    "wWinId": 1,
    "segs": [ {
      "utf8": "people",
      "acAsrConf": 173
    }, {
      "utf8": " have",
      "tOffsetMs": 390,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1020,
      "acAsrConf": 225
    }, {
      "utf8": " pay",
      "tOffsetMs": 1260,
      "acAsrConf": 252
    }, {
      "utf8": " 15",
      "tOffsetMs": 1500,
      "acAsrConf": 219
    }, {
      "utf8": " dollars",
      "tOffsetMs": 2040,
      "acAsrConf": 234
    }, {
      "utf8": " for",
      "tOffsetMs": 2069,
      "acAsrConf": 142
    }, {
      "utf8": " this",
      "tOffsetMs": 2760,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1911570,
    "dDurationMs": 2500,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1911580,
    "dDurationMs": 4770,
    "wWinId": 1,
    "segs": [ {
      "utf8": "edition",
      "acAsrConf": 210
    }, {
      "utf8": " but",
      "tOffsetMs": 450,
      "acAsrConf": 236
    }, {
      "utf8": " it's",
      "tOffsetMs": 690,
      "acAsrConf": 253
    }, {
      "utf8": " free",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " excellent",
      "tOffsetMs": 1200,
      "acAsrConf": 236
    }, {
      "utf8": " over",
      "tOffsetMs": 1890,
      "acAsrConf": 200
    }, {
      "utf8": " 20",
      "tOffsetMs": 2219,
      "acAsrConf": 224
    } ]
  }, {
    "tStartMs": 1914060,
    "dDurationMs": 2290,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1914070,
    "dDurationMs": 4830,
    "wWinId": 1,
    "segs": [ {
      "utf8": "bro",
      "acAsrConf": 243
    }, {
      "utf8": " yes",
      "tOffsetMs": 180,
      "acAsrConf": 216
    }, {
      "utf8": " only",
      "tOffsetMs": 690,
      "acAsrConf": 208
    }, {
      "utf8": " for",
      "tOffsetMs": 930,
      "acAsrConf": 181
    }, {
      "utf8": " 24",
      "tOffsetMs": 1170,
      "acAsrConf": 202
    }, {
      "utf8": " hours",
      "tOffsetMs": 1470,
      "acAsrConf": 252
    }, {
      "utf8": " but",
      "tOffsetMs": 1650,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 1830,
      "acAsrConf": 228
    }, {
      "utf8": " want",
      "tOffsetMs": 2010,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 2190,
      "acAsrConf": 255
    } ]
  }, {
    "tStartMs": 1916340,
    "dDurationMs": 2560,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1916350,
    "dDurationMs": 5940,
    "wWinId": 1,
    "segs": [ {
      "utf8": "give",
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 150,
      "acAsrConf": 200
    }, {
      "utf8": " two",
      "tOffsetMs": 240,
      "acAsrConf": 243
    }, {
      "utf8": " more",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " things",
      "tOffsetMs": 930,
      "acAsrConf": 208
    }, {
      "utf8": " one",
      "tOffsetMs": 1260,
      "acAsrConf": 225
    }, {
      "utf8": " of",
      "tOffsetMs": 2010,
      "acAsrConf": 252
    }, {
      "utf8": " them",
      "tOffsetMs": 2040,
      "acAsrConf": 202
    }, {
      "utf8": " is",
      "tOffsetMs": 2250,
      "acAsrConf": 204
    } ]
  }, {
    "tStartMs": 1918890,
    "dDurationMs": 3400,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1918900,
    "dDurationMs": 6630,
    "wWinId": 1,
    "segs": [ {
      "utf8": "the",
      "acAsrConf": 238
    }, {
      "utf8": " audiobook",
      "tOffsetMs": 890,
      "acAsrConf": 226
    }, {
      "utf8": " so",
      "tOffsetMs": 1890,
      "acAsrConf": 224
    }, {
      "utf8": " a",
      "tOffsetMs": 2190,
      "acAsrConf": 246
    }, {
      "utf8": " lot",
      "tOffsetMs": 2370,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 2730,
      "acAsrConf": 167
    }, {
      "utf8": " you",
      "tOffsetMs": 2850,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 2970,
      "acAsrConf": 252
    }, {
      "utf8": " read",
      "tOffsetMs": 3149,
      "acAsrConf": 220
    } ]
  }, {
    "tStartMs": 1922280,
    "dDurationMs": 3250,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1922290,
    "dDurationMs": 5310,
    "wWinId": 1,
    "segs": [ {
      "utf8": "my",
      "acAsrConf": 252
    }, {
      "utf8": " free",
      "tOffsetMs": 30,
      "acAsrConf": 186
    }, {
      "utf8": " ebook",
      "tOffsetMs": 540,
      "acAsrConf": 228
    }, {
      "utf8": " 5",
      "tOffsetMs": 750,
      "acAsrConf": 166
    }, {
      "utf8": " steps",
      "tOffsetMs": 1680,
      "acAsrConf": 243
    }, {
      "utf8": " to",
      "tOffsetMs": 2220,
      "acAsrConf": 226
    }, {
      "utf8": " becoming",
      "tOffsetMs": 2580,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 2940,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1925520,
    "dDurationMs": 2080,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1925530,
    "dDurationMs": 5160,
    "wWinId": 1,
    "segs": [ {
      "utf8": "confident",
      "acAsrConf": 252
    }, {
      "utf8": " English",
      "tOffsetMs": 480,
      "acAsrConf": 227
    }, {
      "utf8": " speaker",
      "tOffsetMs": 870,
      "acAsrConf": 252
    }, {
      "utf8": " but",
      "tOffsetMs": 1320,
      "acAsrConf": 145
    }, {
      "utf8": " there's",
      "tOffsetMs": 1529,
      "acAsrConf": 98
    } ]
  }, {
    "tStartMs": 1927590,
    "dDurationMs": 3100,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1927600,
    "dDurationMs": 7230,
    "wWinId": 1,
    "segs": [ {
      "utf8": "also",
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 270,
      "acAsrConf": 252
    }, {
      "utf8": " 22",
      "tOffsetMs": 449,
      "acAsrConf": 223
    }, {
      "utf8": " minute",
      "tOffsetMs": 1410,
      "acAsrConf": 153
    }, {
      "utf8": " audio",
      "tOffsetMs": 1740,
      "acAsrConf": 204
    }, {
      "utf8": " book",
      "tOffsetMs": 1920,
      "acAsrConf": 168
    }, {
      "utf8": " that",
      "tOffsetMs": 2550,
      "acAsrConf": 236
    }, {
      "utf8": " goes",
      "tOffsetMs": 2850,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1930680,
    "dDurationMs": 4150,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1930690,
    "dDurationMs": 6630,
    "wWinId": 1,
    "segs": [ {
      "utf8": "with",
      "acAsrConf": 238
    }, {
      "utf8": " that",
      "tOffsetMs": 540,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 570,
      "acAsrConf": 204
    }, {
      "utf8": " the",
      "tOffsetMs": 1670,
      "acAsrConf": 236
    }, {
      "utf8": " audio",
      "tOffsetMs": 2670,
      "acAsrConf": 217
    }, {
      "utf8": " book",
      "tOffsetMs": 3060,
      "acAsrConf": 217
    }, {
      "utf8": " is",
      "tOffsetMs": 3330,
      "acAsrConf": 238
    }, {
      "utf8": " usually",
      "tOffsetMs": 3540,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1934820,
    "dDurationMs": 2500,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1934830,
    "dDurationMs": 4469,
    "wWinId": 1,
    "segs": [ {
      "utf8": "seven",
      "acAsrConf": 251
    }, {
      "utf8": " dollars",
      "tOffsetMs": 570,
      "acAsrConf": 251
    }, {
      "utf8": " and",
      "tOffsetMs": 1050,
      "acAsrConf": 251
    }, {
      "utf8": " fifty",
      "tOffsetMs": 1320,
      "acAsrConf": 251
    }, {
      "utf8": " cents",
      "tOffsetMs": 1440,
      "acAsrConf": 175
    }, {
      "utf8": " but",
      "tOffsetMs": 1860,
      "acAsrConf": 247
    }, {
      "utf8": " I",
      "tOffsetMs": 2040,
      "acAsrConf": 217
    }, {
      "utf8": " want",
      "tOffsetMs": 2219,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1937310,
    "dDurationMs": 1989,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1937320,
    "dDurationMs": 4650,
    "wWinId": 1,
    "segs": [ {
      "utf8": "to",
      "acAsrConf": 251
    }, {
      "utf8": " give",
      "tOffsetMs": 90,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 210,
      "acAsrConf": 231
    }, {
      "utf8": " to",
      "tOffsetMs": 330,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 359,
      "acAsrConf": 144
    }, {
      "utf8": " for",
      "tOffsetMs": 540,
      "acAsrConf": 252
    }, {
      "utf8": " free",
      "tOffsetMs": 780,
      "acAsrConf": 217
    }, {
      "utf8": " so",
      "tOffsetMs": 1080,
      "acAsrConf": 236
    }, {
      "utf8": " you",
      "tOffsetMs": 1560,
      "acAsrConf": 128
    }, {
      "utf8": " can",
      "tOffsetMs": 1710,
      "acAsrConf": 166
    } ]
  }, {
    "tStartMs": 1939289,
    "dDurationMs": 2681,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1939299,
    "dDurationMs": 5461,
    "wWinId": 1,
    "segs": [ {
      "utf8": "listen",
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 181,
      "acAsrConf": 218
    }, {
      "utf8": " it",
      "tOffsetMs": 391,
      "acAsrConf": 158
    }, {
      "utf8": " and",
      "tOffsetMs": 661,
      "acAsrConf": 252
    }, {
      "utf8": " enjoy",
      "tOffsetMs": 841,
      "acAsrConf": 235
    }, {
      "utf8": " it",
      "tOffsetMs": 1260,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1291,
      "acAsrConf": 187
    }, {
      "utf8": " to",
      "tOffsetMs": 1651,
      "acAsrConf": 238
    }, {
      "utf8": " thank",
      "tOffsetMs": 2431,
      "acAsrConf": 209
    } ]
  }, {
    "tStartMs": 1941960,
    "dDurationMs": 2800,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1941970,
    "dDurationMs": 5570,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 216
    }, {
      "utf8": " for",
      "tOffsetMs": 180,
      "acAsrConf": 201
    }, {
      "utf8": " being",
      "tOffsetMs": 449,
      "acAsrConf": 202
    }, {
      "utf8": " on",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " this",
      "tOffsetMs": 870,
      "acAsrConf": 216
    }, {
      "utf8": " webinar",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1370,
      "acAsrConf": 248
    }, {
      "utf8": " we",
      "tOffsetMs": 2370,
      "acAsrConf": 173
    }, {
      "utf8": " have",
      "tOffsetMs": 2459,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1944750,
    "dDurationMs": 2790,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1944760,
    "dDurationMs": 5610,
    "wWinId": 1,
    "segs": [ {
      "utf8": "super",
      "acAsrConf": 248
    }, {
      "utf8": " English",
      "tOffsetMs": 360,
      "acAsrConf": 255
    }, {
      "utf8": " Membership",
      "tOffsetMs": 630,
      "acAsrConf": 185
    }, {
      "utf8": " a",
      "tOffsetMs": 1169,
      "acAsrConf": 214
    }, {
      "utf8": " free",
      "tOffsetMs": 1380,
      "acAsrConf": 252
    }, {
      "utf8": " ebook",
      "tOffsetMs": 2070,
      "acAsrConf": 248
    } ]
  }, {
    "tStartMs": 1947530,
    "dDurationMs": 2840,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1947540,
    "dDurationMs": 4600,
    "wWinId": 1,
    "segs": [ {
      "utf8": "audiobook",
      "acAsrConf": 211
    }, {
      "utf8": " download",
      "tOffsetMs": 1000,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1030,
      "acAsrConf": 227
    }, {
      "utf8": " finally",
      "tOffsetMs": 1780,
      "acAsrConf": 252
    }, {
      "utf8": " this",
      "tOffsetMs": 2470,
      "acAsrConf": 236
    }, {
      "utf8": " is",
      "tOffsetMs": 2650,
      "acAsrConf": 112
    } ]
  }, {
    "tStartMs": 1950360,
    "dDurationMs": 1780,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1950370,
    "dDurationMs": 6510,
    "wWinId": 1,
    "segs": [ {
      "utf8": "the",
      "acAsrConf": 170
    }, {
      "utf8": " most",
      "tOffsetMs": 120,
      "acAsrConf": 252
    }, {
      "utf8": " exciting",
      "tOffsetMs": 299,
      "acAsrConf": 234
    }, {
      "utf8": " one",
      "tOffsetMs": 840,
      "acAsrConf": 252
    }, {
      "utf8": " I'm",
      "tOffsetMs": 870,
      "acAsrConf": 93
    }, {
      "utf8": " looking",
      "tOffsetMs": 1260,
      "acAsrConf": 255
    } ]
  }, {
    "tStartMs": 1952130,
    "dDurationMs": 4750,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1952140,
    "dDurationMs": 8960,
    "wWinId": 1,
    "segs": [ {
      "utf8": "forward",
      "acAsrConf": 232
    }, {
      "utf8": " to",
      "tOffsetMs": 390,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 419,
      "acAsrConf": 200
    }, {
      "utf8": " free",
      "tOffsetMs": 1669,
      "acAsrConf": 180
    }, {
      "utf8": " 10",
      "tOffsetMs": 2669,
      "acAsrConf": 145
    }, {
      "utf8": " minute",
      "tOffsetMs": 3419,
      "acAsrConf": 251
    }, {
      "utf8": " lesson",
      "tOffsetMs": 3900,
      "acAsrConf": 245
    }, {
      "utf8": " with",
      "tOffsetMs": 4409,
      "acAsrConf": 234
    } ]
  }, {
    "tStartMs": 1956870,
    "dDurationMs": 4230,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1956880,
    "dDurationMs": 7080,
    "wWinId": 1,
    "segs": [ {
      "utf8": "me",
      "acAsrConf": 252
    }, {
      "utf8": " on",
      "tOffsetMs": 270,
      "acAsrConf": 252
    }, {
      "utf8": " skype",
      "tOffsetMs": 539,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 2090,
      "acAsrConf": 212
    }, {
      "utf8": " this",
      "tOffsetMs": 3090,
      "acAsrConf": 250
    }, {
      "utf8": " is",
      "tOffsetMs": 3510,
      "acAsrConf": 251
    }, {
      "utf8": " exciting",
      "tOffsetMs": 3660,
      "acAsrConf": 227
    }, {
      "utf8": " because",
      "tOffsetMs": 4049,
      "acAsrConf": 227
    } ]
  }, {
    "tStartMs": 1961090,
    "dDurationMs": 2870,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1961100,
    "dDurationMs": 5650,
    "wWinId": 1,
    "segs": [ {
      "utf8": "usually",
      "acAsrConf": 187
    }, {
      "utf8": " this",
      "tOffsetMs": 1000,
      "acAsrConf": 252
    }, {
      "utf8": " costs",
      "tOffsetMs": 1209,
      "acAsrConf": 241
    }, {
      "utf8": " about",
      "tOffsetMs": 1690,
      "acAsrConf": 252
    }, {
      "utf8": " ten",
      "tOffsetMs": 1750,
      "acAsrConf": 0
    }, {
      "utf8": " dollars",
      "tOffsetMs": 2199,
      "acAsrConf": 0
    }, {
      "utf8": " for",
      "tOffsetMs": 2380,
      "acAsrConf": 165
    } ]
  }, {
    "tStartMs": 1963950,
    "dDurationMs": 2800,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1963960,
    "dDurationMs": 5219,
    "wWinId": 1,
    "segs": [ {
      "utf8": "10",
      "acAsrConf": 201
    }, {
      "utf8": " minutes",
      "tOffsetMs": 240,
      "acAsrConf": 227
    }, {
      "utf8": " and",
      "tOffsetMs": 570,
      "acAsrConf": 216
    }, {
      "utf8": " you'll",
      "tOffsetMs": 950,
      "acAsrConf": 239
    }, {
      "utf8": " have",
      "tOffsetMs": 1950,
      "acAsrConf": 243
    }, {
      "utf8": " a",
      "tOffsetMs": 2190,
      "acAsrConf": 252
    }, {
      "utf8": " chance",
      "tOffsetMs": 2219,
      "acAsrConf": 199
    }, {
      "utf8": " to",
      "tOffsetMs": 2459,
      "acAsrConf": 200
    } ]
  }, {
    "tStartMs": 1966740,
    "dDurationMs": 2439,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1966750,
    "dDurationMs": 4860,
    "wWinId": 1,
    "segs": [ {
      "utf8": "meet",
      "acAsrConf": 252
    }, {
      "utf8": " with",
      "tOffsetMs": 720,
      "acAsrConf": 243
    }, {
      "utf8": " me",
      "tOffsetMs": 870,
      "acAsrConf": 225
    }, {
      "utf8": " on",
      "tOffsetMs": 1080,
      "acAsrConf": 115
    }, {
      "utf8": " skype",
      "tOffsetMs": 1230,
      "acAsrConf": 252
    }, {
      "utf8": " ask",
      "tOffsetMs": 1559,
      "acAsrConf": 208
    }, {
      "utf8": " me",
      "tOffsetMs": 1919,
      "acAsrConf": 247
    }, {
      "utf8": " any",
      "tOffsetMs": 2250,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1969169,
    "dDurationMs": 2441,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1969179,
    "dDurationMs": 5581,
    "wWinId": 1,
    "segs": [ {
      "utf8": "questions",
      "acAsrConf": 247
    }, {
      "utf8": " I",
      "tOffsetMs": 571,
      "acAsrConf": 204
    }, {
      "utf8": " can",
      "tOffsetMs": 781,
      "acAsrConf": 239
    }, {
      "utf8": " give",
      "tOffsetMs": 870,
      "acAsrConf": 164
    }, {
      "utf8": " you",
      "tOffsetMs": 1261,
      "acAsrConf": 235
    }, {
      "utf8": " feedback",
      "tOffsetMs": 1411,
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 1711,
      "acAsrConf": 235
    } ]
  }, {
    "tStartMs": 1971600,
    "dDurationMs": 3160,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1971610,
    "dDurationMs": 5610,
    "wWinId": 1,
    "segs": [ {
      "utf8": "your",
      "acAsrConf": 246
    }, {
      "utf8": " speaking",
      "tOffsetMs": 330,
      "acAsrConf": 236
    }, {
      "utf8": " and",
      "tOffsetMs": 840,
      "acAsrConf": 213
    }, {
      "utf8": " you'll",
      "tOffsetMs": 1520,
      "acAsrConf": 221
    }, {
      "utf8": " be",
      "tOffsetMs": 2520,
      "acAsrConf": 252
    }, {
      "utf8": " able",
      "tOffsetMs": 2670,
      "acAsrConf": 154
    }, {
      "utf8": " to",
      "tOffsetMs": 2939,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1974750,
    "dDurationMs": 2470,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1974760,
    "dDurationMs": 4470,
    "wWinId": 1,
    "segs": [ {
      "utf8": "listen",
      "acAsrConf": 238
    }, {
      "utf8": " and",
      "tOffsetMs": 840,
      "acAsrConf": 205
    }, {
      "utf8": " hear",
      "tOffsetMs": 1260,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 1680,
      "acAsrConf": 252
    }, {
      "utf8": " native",
      "tOffsetMs": 1710,
      "acAsrConf": 201
    }, {
      "utf8": " English",
      "tOffsetMs": 2159,
      "acAsrConf": 245
    }, {
      "utf8": " speaker",
      "tOffsetMs": 2430,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1977210,
    "dDurationMs": 2020,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1977220,
    "dDurationMs": 4080,
    "wWinId": 1,
    "segs": [ {
      "utf8": "maybe",
      "acAsrConf": 243
    }, {
      "utf8": " some",
      "tOffsetMs": 540,
      "acAsrConf": 228
    }, {
      "utf8": " of",
      "tOffsetMs": 810,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 930,
      "acAsrConf": 248
    }, {
      "utf8": " have",
      "tOffsetMs": 1020,
      "acAsrConf": 252
    }, {
      "utf8": " never",
      "tOffsetMs": 1199,
      "acAsrConf": 252
    }, {
      "utf8": " spoken",
      "tOffsetMs": 1380,
      "acAsrConf": 201
    }, {
      "utf8": " with",
      "tOffsetMs": 1860,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1979220,
    "dDurationMs": 2080,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1979230,
    "dDurationMs": 4380,
    "wWinId": 1,
    "segs": [ {
      "utf8": "a",
      "acAsrConf": 211
    }, {
      "utf8": " native",
      "tOffsetMs": 300,
      "acAsrConf": 252
    }, {
      "utf8": " English",
      "tOffsetMs": 660,
      "acAsrConf": 239
    }, {
      "utf8": " speaker",
      "tOffsetMs": 1020,
      "acAsrConf": 252
    }, {
      "utf8": " before",
      "tOffsetMs": 1050,
      "acAsrConf": 243
    }, {
      "utf8": " so",
      "tOffsetMs": 1500,
      "acAsrConf": 201
    }, {
      "utf8": " it's",
      "tOffsetMs": 1920,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1981290,
    "dDurationMs": 2320,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1981300,
    "dDurationMs": 3720,
    "wWinId": 1,
    "segs": [ {
      "utf8": "a",
      "acAsrConf": 252
    }, {
      "utf8": " great",
      "tOffsetMs": 90,
      "acAsrConf": 252
    }, {
      "utf8": " chance",
      "tOffsetMs": 360,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 629,
      "acAsrConf": 217
    }, {
      "utf8": " test",
      "tOffsetMs": 930,
      "acAsrConf": 202
    }, {
      "utf8": " it",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    }, {
      "utf8": " out",
      "tOffsetMs": 1710,
      "acAsrConf": 208
    }, {
      "utf8": " for",
      "tOffsetMs": 1830,
      "acAsrConf": 186
    }, {
      "utf8": " 10",
      "tOffsetMs": 2160,
      "acAsrConf": 216
    } ]
  }, {
    "tStartMs": 1983600,
    "dDurationMs": 1420,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1983610,
    "dDurationMs": 3390,
    "wWinId": 1,
    "segs": [ {
      "utf8": "minutes",
      "acAsrConf": 223
    }, {
      "utf8": " not",
      "tOffsetMs": 270,
      "acAsrConf": 252
    }, {
      "utf8": " too",
      "tOffsetMs": 449,
      "acAsrConf": 226
    }, {
      "utf8": " long",
      "tOffsetMs": 690,
      "acAsrConf": 227
    } ]
  }, {
    "tStartMs": 1985010,
    "dDurationMs": 1990,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1985020,
    "dDurationMs": 5280,
    "wWinId": 1,
    "segs": [ {
      "utf8": "too",
      "acAsrConf": 213
    }, {
      "utf8": " short",
      "tOffsetMs": 90,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 300,
      "acAsrConf": 200
    }, {
      "utf8": " it'll",
      "tOffsetMs": 630,
      "acAsrConf": 203
    }, {
      "utf8": " be",
      "tOffsetMs": 1260,
      "acAsrConf": 235
    }, {
      "utf8": " a",
      "tOffsetMs": 1350,
      "acAsrConf": 252
    }, {
      "utf8": " good",
      "tOffsetMs": 1380,
      "acAsrConf": 252
    }, {
      "utf8": " chance",
      "tOffsetMs": 1590,
      "acAsrConf": 207
    }, {
      "utf8": " for",
      "tOffsetMs": 1860,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1986990,
    "dDurationMs": 3310,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1987000,
    "dDurationMs": 5670,
    "wWinId": 1,
    "segs": [ {
      "utf8": "a",
      "acAsrConf": 233
    }, {
      "utf8": " lot",
      "tOffsetMs": 150,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 300,
      "acAsrConf": 173
    }, {
      "utf8": " you",
      "tOffsetMs": 360,
      "acAsrConf": 201
    }, {
      "utf8": " so",
      "tOffsetMs": 600,
      "acAsrConf": 255
    }, {
      "utf8": " I",
      "tOffsetMs": 1590,
      "acAsrConf": 255
    }, {
      "utf8": " am",
      "tOffsetMs": 1860,
      "acAsrConf": 252
    }, {
      "utf8": " really",
      "tOffsetMs": 2640,
      "acAsrConf": 221
    }, {
      "utf8": " excited",
      "tOffsetMs": 2820,
      "acAsrConf": 201
    }, {
      "utf8": " to",
      "tOffsetMs": 2940,
      "acAsrConf": 223
    } ]
  }, {
    "tStartMs": 1990290,
    "dDurationMs": 2380,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1990300,
    "dDurationMs": 5130,
    "wWinId": 1,
    "segs": [ {
      "utf8": "share",
      "acAsrConf": 252
    }, {
      "utf8": " this",
      "tOffsetMs": 150,
      "acAsrConf": 252
    }, {
      "utf8": " with",
      "tOffsetMs": 330,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 390,
      "acAsrConf": 155
    }, {
      "utf8": " and",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 750,
      "acAsrConf": 204
    }, {
      "utf8": " the",
      "tOffsetMs": 1410,
      "acAsrConf": 252
    }, {
      "utf8": " mind",
      "tOffsetMs": 1800,
      "acAsrConf": 214
    }, {
      "utf8": " map",
      "tOffsetMs": 2070,
      "acAsrConf": 214
    } ]
  }, {
    "tStartMs": 1992660,
    "dDurationMs": 2770,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1992670,
    "dDurationMs": 5790,
    "wWinId": 1,
    "segs": [ {
      "utf8": "I",
      "acAsrConf": 218
    }, {
      "utf8": " gave",
      "tOffsetMs": 240,
      "acAsrConf": 245
    }, {
      "utf8": " you",
      "tOffsetMs": 690,
      "acAsrConf": 255
    }, {
      "utf8": " a",
      "tOffsetMs": 900,
      "acAsrConf": 252
    }, {
      "utf8": " link",
      "tOffsetMs": 930,
      "acAsrConf": 238
    }, {
      "utf8": " where",
      "tOffsetMs": 1620,
      "acAsrConf": 202
    }, {
      "utf8": " you",
      "tOffsetMs": 2220,
      "acAsrConf": 252
    }, {
      "utf8": " can",
      "tOffsetMs": 2250,
      "acAsrConf": 252
    }, {
      "utf8": " go",
      "tOffsetMs": 2400,
      "acAsrConf": 187
    }, {
      "utf8": " to",
      "tOffsetMs": 2730,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 1995420,
    "dDurationMs": 3040,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1995430,
    "dDurationMs": 5130,
    "wWinId": 1,
    "segs": [ {
      "utf8": "that",
      "acAsrConf": 179
    }, {
      "utf8": " and",
      "tOffsetMs": 330,
      "acAsrConf": 208
    }, {
      "utf8": " I'll",
      "tOffsetMs": 690,
      "acAsrConf": 239
    }, {
      "utf8": " send",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 2070,
      "acAsrConf": 217
    }, {
      "utf8": " again",
      "tOffsetMs": 2220,
      "acAsrConf": 239
    }, {
      "utf8": " when",
      "tOffsetMs": 2370,
      "acAsrConf": 164
    }, {
      "utf8": " we",
      "tOffsetMs": 2850,
      "acAsrConf": 238
    } ]
  }, {
    "tStartMs": 1998450,
    "dDurationMs": 2110,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 1998460,
    "dDurationMs": 5040,
    "wWinId": 1,
    "segs": [ {
      "utf8": "start",
      "acAsrConf": 236
    }, {
      "utf8": " answering",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " questions",
      "tOffsetMs": 450,
      "acAsrConf": 252
    }, {
      "utf8": " I'll",
      "tOffsetMs": 750,
      "acAsrConf": 104
    }, {
      "utf8": " start",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2000550,
    "dDurationMs": 2950,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2000560,
    "dDurationMs": 5940,
    "wWinId": 1,
    "segs": [ {
      "utf8": "I'll",
      "acAsrConf": 216
    }, {
      "utf8": " give",
      "tOffsetMs": 270,
      "acAsrConf": 165
    }, {
      "utf8": " you",
      "tOffsetMs": 1230,
      "acAsrConf": 248
    }, {
      "utf8": " the",
      "tOffsetMs": 1410,
      "acAsrConf": 248
    }, {
      "utf8": " link",
      "tOffsetMs": 1740,
      "acAsrConf": 245
    }, {
      "utf8": " so",
      "tOffsetMs": 1980,
      "acAsrConf": 252
    }, {
      "utf8": " that",
      "tOffsetMs": 2070,
      "acAsrConf": 61
    }, {
      "utf8": " you",
      "tOffsetMs": 2340,
      "acAsrConf": 132
    }, {
      "utf8": " can",
      "tOffsetMs": 2550,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2003490,
    "dDurationMs": 3010,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2003500,
    "dDurationMs": 5910,
    "wWinId": 1,
    "segs": [ {
      "utf8": "join",
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 930,
      "acAsrConf": 239
    }, {
      "utf8": " 30",
      "tOffsetMs": 1410,
      "acAsrConf": 252
    }, {
      "utf8": " days",
      "tOffsetMs": 1650,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 1860,
      "acAsrConf": 235
    }, {
      "utf8": " English",
      "tOffsetMs": 2040,
      "acAsrConf": 237
    }, {
      "utf8": " and",
      "tOffsetMs": 2370,
      "acAsrConf": 231
    }, {
      "utf8": " get",
      "tOffsetMs": 2580,
      "acAsrConf": 214
    }, {
      "utf8": " all",
      "tOffsetMs": 2820,
      "acAsrConf": 255
    } ]
  }, {
    "tStartMs": 2006490,
    "dDurationMs": 2920,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2006500,
    "dDurationMs": 5550,
    "wWinId": 1,
    "segs": [ {
      "utf8": "of",
      "acAsrConf": 252
    }, {
      "utf8": " those",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " extra",
      "tOffsetMs": 440,
      "acAsrConf": 255
    }, {
      "utf8": " free",
      "tOffsetMs": 1440,
      "acAsrConf": 227
    }, {
      "utf8": " things",
      "tOffsetMs": 1800,
      "acAsrConf": 252
    }, {
      "utf8": " super",
      "tOffsetMs": 2100,
      "acAsrConf": 252
    }, {
      "utf8": " English",
      "tOffsetMs": 2580,
      "acAsrConf": 242
    } ]
  }, {
    "tStartMs": 2009400,
    "dDurationMs": 2650,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2009410,
    "dDurationMs": 4860,
    "wWinId": 1,
    "segs": [ {
      "utf8": "membership",
      "acAsrConf": 144
    }, {
      "utf8": " audiobook",
      "tOffsetMs": 540,
      "acAsrConf": 127
    }, {
      "utf8": " and",
      "tOffsetMs": 1320,
      "acAsrConf": 249
    }, {
      "utf8": " the",
      "tOffsetMs": 1590,
      "acAsrConf": 236
    }, {
      "utf8": " ten",
      "tOffsetMs": 2220,
      "acAsrConf": 215
    }, {
      "utf8": " minute",
      "tOffsetMs": 2400,
      "acAsrConf": 165
    } ]
  }, {
    "tStartMs": 2012040,
    "dDurationMs": 2230,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2012050,
    "dDurationMs": 4440,
    "wWinId": 1,
    "segs": [ {
      "utf8": "Skype",
      "acAsrConf": 231
    }, {
      "utf8": " I",
      "tOffsetMs": 330,
      "acAsrConf": 90
    }, {
      "utf8": " would",
      "tOffsetMs": 570,
      "acAsrConf": 252
    }, {
      "utf8": " love",
      "tOffsetMs": 840,
      "acAsrConf": 236
    }, {
      "utf8": " to",
      "tOffsetMs": 1200,
      "acAsrConf": 214
    }, {
      "utf8": " meet",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1620,
      "acAsrConf": 201
    }, {
      "utf8": " on",
      "tOffsetMs": 1740,
      "acAsrConf": 252
    }, {
      "utf8": " skype",
      "tOffsetMs": 1920,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2014260,
    "dDurationMs": 2230,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2014270,
    "dDurationMs": 5450,
    "wWinId": 1,
    "segs": [ {
      "utf8": "it",
      "acAsrConf": 247
    }, {
      "utf8": " would",
      "tOffsetMs": 270,
      "acAsrConf": 252
    }, {
      "utf8": " be",
      "tOffsetMs": 420,
      "acAsrConf": 252
    }, {
      "utf8": " great",
      "tOffsetMs": 540,
      "acAsrConf": 252
    }, {
      "utf8": " maybe",
      "tOffsetMs": 870,
      "acAsrConf": 202
    }, {
      "utf8": " if",
      "tOffsetMs": 1680,
      "acAsrConf": 203
    }, {
      "utf8": " you're",
      "tOffsetMs": 2010,
      "acAsrConf": 252
    }, {
      "utf8": " lucky",
      "tOffsetMs": 2190,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2016480,
    "dDurationMs": 3240,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2016490,
    "dDurationMs": 6870,
    "wWinId": 1,
    "segs": [ {
      "utf8": "Dan",
      "acAsrConf": 202
    }, {
      "utf8": " will",
      "tOffsetMs": 600,
      "acAsrConf": 237
    }, {
      "utf8": " be",
      "tOffsetMs": 870,
      "acAsrConf": 252
    }, {
      "utf8": " there",
      "tOffsetMs": 900,
      "acAsrConf": 200
    }, {
      "utf8": " too",
      "tOffsetMs": 1170,
      "acAsrConf": 249
    }, {
      "utf8": " but",
      "tOffsetMs": 1410,
      "acAsrConf": 248
    }, {
      "utf8": " we",
      "tOffsetMs": 1680,
      "acAsrConf": 178
    }, {
      "utf8": " don't",
      "tOffsetMs": 1950,
      "acAsrConf": 252
    }, {
      "utf8": " know",
      "tOffsetMs": 1980,
      "acAsrConf": 202
    } ]
  }, {
    "tStartMs": 2019710,
    "dDurationMs": 3650,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2019720,
    "dDurationMs": 6910,
    "wWinId": 1,
    "segs": [ {
      "utf8": "definitely",
      "acAsrConf": 211
    }, {
      "utf8": " you",
      "tOffsetMs": 1000,
      "acAsrConf": 252
    }, {
      "utf8": " can",
      "tOffsetMs": 1120,
      "acAsrConf": 252
    }, {
      "utf8": " meet",
      "tOffsetMs": 1330,
      "acAsrConf": 252
    }, {
      "utf8": " me",
      "tOffsetMs": 1480,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1660,
      "acAsrConf": 240
    }, {
      "utf8": " yes",
      "tOffsetMs": 2340,
      "acAsrConf": 240
    }, {
      "utf8": " so",
      "tOffsetMs": 3340,
      "acAsrConf": 247
    } ]
  }, {
    "tStartMs": 2023350,
    "dDurationMs": 3280,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2023360,
    "dDurationMs": 5670,
    "wWinId": 1,
    "segs": [ {
      "utf8": "this",
      "acAsrConf": 228
    }, {
      "utf8": " is",
      "tOffsetMs": 240,
      "acAsrConf": 187
    }, {
      "utf8": " our",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " special",
      "tOffsetMs": 860,
      "acAsrConf": 252
    }, {
      "utf8": " gift",
      "tOffsetMs": 1860,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 2100,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 2220,
      "acAsrConf": 207
    }, {
      "utf8": " and",
      "tOffsetMs": 2490,
      "acAsrConf": 223
    }, {
      "utf8": " all",
      "tOffsetMs": 2790,
      "acAsrConf": 249
    } ]
  }, {
    "tStartMs": 2026620,
    "dDurationMs": 2410,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2026630,
    "dDurationMs": 4590,
    "wWinId": 1,
    "segs": [ {
      "utf8": "of",
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 30,
      "acAsrConf": 143
    }, {
      "utf8": " tips",
      "tOffsetMs": 330,
      "acAsrConf": 252
    }, {
      "utf8": " we've",
      "tOffsetMs": 540,
      "acAsrConf": 252
    }, {
      "utf8": " given",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 990,
      "acAsrConf": 252
    }, {
      "utf8": " today",
      "tOffsetMs": 1170,
      "acAsrConf": 128
    }, {
      "utf8": " I",
      "tOffsetMs": 1500,
      "acAsrConf": 175
    }, {
      "utf8": " hope",
      "tOffsetMs": 1740,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2029020,
    "dDurationMs": 2200,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2029030,
    "dDurationMs": 4800,
    "wWinId": 1,
    "segs": [ {
      "utf8": "really",
      "acAsrConf": 219
    }, {
      "utf8": " helpful",
      "tOffsetMs": 510,
      "acAsrConf": 216
    }, {
      "utf8": " for",
      "tOffsetMs": 1110,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1140,
      "acAsrConf": 204
    }, {
      "utf8": " I'd",
      "tOffsetMs": 1470,
      "acAsrConf": 239
    }, {
      "utf8": " like",
      "tOffsetMs": 1890,
      "acAsrConf": 202
    }, {
      "utf8": " to",
      "tOffsetMs": 2160,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2031210,
    "dDurationMs": 2620,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2031220,
    "dDurationMs": 4230,
    "wWinId": 1,
    "segs": [ {
      "utf8": "answer",
      "acAsrConf": 203
    }, {
      "utf8": " some",
      "tOffsetMs": 660,
      "acAsrConf": 200
    }, {
      "utf8": " questions",
      "tOffsetMs": 1110,
      "acAsrConf": 252
    }, {
      "utf8": " now",
      "tOffsetMs": 1140,
      "acAsrConf": 125
    }, {
      "utf8": " yeah",
      "tOffsetMs": 1680,
      "acAsrConf": 140
    }, {
      "utf8": " let's",
      "tOffsetMs": 2340,
      "acAsrConf": 245
    } ]
  }, {
    "tStartMs": 2033820,
    "dDurationMs": 1630,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2033830,
    "dDurationMs": 3420,
    "wWinId": 1,
    "segs": [ {
      "utf8": "take",
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 120,
      "acAsrConf": 207
    }, {
      "utf8": " look",
      "tOffsetMs": 180,
      "acAsrConf": 252
    }, {
      "utf8": " at",
      "tOffsetMs": 450,
      "acAsrConf": 252
    }, {
      "utf8": " them",
      "tOffsetMs": 540,
      "acAsrConf": 248
    }, {
      "utf8": " yes",
      "tOffsetMs": 690,
      "acAsrConf": 205
    }, {
      "utf8": " let's",
      "tOffsetMs": 930,
      "acAsrConf": 230
    }, {
      "utf8": " look",
      "tOffsetMs": 1170,
      "acAsrConf": 176
    }, {
      "utf8": " at",
      "tOffsetMs": 1230,
      "acAsrConf": 144
    } ]
  }, {
    "tStartMs": 2035440,
    "dDurationMs": 1810,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2035450,
    "dDurationMs": 5670,
    "wWinId": 1,
    "segs": [ {
      "utf8": "these",
      "acAsrConf": 252
    }, {
      "utf8": " questions",
      "tOffsetMs": 150,
      "acAsrConf": 206
    }, {
      "utf8": " Dan's",
      "tOffsetMs": 630,
      "acAsrConf": 150
    }, {
      "utf8": " going",
      "tOffsetMs": 1050,
      "acAsrConf": 214
    }, {
      "utf8": " to",
      "tOffsetMs": 1260,
      "acAsrConf": 202
    }, {
      "utf8": " help",
      "tOffsetMs": 1350,
      "acAsrConf": 252
    }, {
      "utf8": " read",
      "tOffsetMs": 1530,
      "acAsrConf": 237
    } ]
  }, {
    "tStartMs": 2037240,
    "dDurationMs": 3880,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2037250,
    "dDurationMs": 6060,
    "wWinId": 1,
    "segs": [ {
      "utf8": "the",
      "acAsrConf": 252
    }, {
      "utf8": " questions",
      "tOffsetMs": 210,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 660,
      "acAsrConf": 242
    }, {
      "utf8": " some",
      "tOffsetMs": 2240,
      "acAsrConf": 251
    }, {
      "utf8": " of",
      "tOffsetMs": 3240,
      "acAsrConf": 227
    }, {
      "utf8": " these",
      "tOffsetMs": 3450,
      "acAsrConf": 252
    }, {
      "utf8": " we",
      "tOffsetMs": 3570,
      "acAsrConf": 249
    } ]
  }, {
    "tStartMs": 2041110,
    "dDurationMs": 2200,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2041120,
    "dDurationMs": 4950,
    "wWinId": 1,
    "segs": [ {
      "utf8": "probably",
      "acAsrConf": 255
    }, {
      "utf8": " already",
      "tOffsetMs": 300,
      "acAsrConf": 186
    }, {
      "utf8": " answered",
      "tOffsetMs": 1050,
      "acAsrConf": 219
    }, {
      "utf8": " during",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    }, {
      "utf8": " our",
      "tOffsetMs": 1800,
      "acAsrConf": 216
    } ]
  }, {
    "tStartMs": 2043300,
    "dDurationMs": 2770,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2043310,
    "dDurationMs": 4740,
    "wWinId": 1,
    "segs": [ {
      "utf8": "presentation",
      "acAsrConf": 252
    }, {
      "utf8": " but",
      "tOffsetMs": 1010,
      "acAsrConf": 147
    }, {
      "utf8": " if",
      "tOffsetMs": 2010,
      "acAsrConf": 244
    }, {
      "utf8": " there",
      "tOffsetMs": 2310,
      "acAsrConf": 252
    }, {
      "utf8": " are",
      "tOffsetMs": 2610,
      "acAsrConf": 252
    }, {
      "utf8": " any",
      "tOffsetMs": 2640,
      "acAsrConf": 207
    } ]
  }, {
    "tStartMs": 2046060,
    "dDurationMs": 1990,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2046070,
    "dDurationMs": 3839,
    "wWinId": 1,
    "segs": [ {
      "utf8": "questions",
      "acAsrConf": 188
    }, {
      "utf8": " we",
      "tOffsetMs": 270,
      "acAsrConf": 167
    }, {
      "utf8": " haven't",
      "tOffsetMs": 630,
      "acAsrConf": 207
    }, {
      "utf8": " answered",
      "tOffsetMs": 990,
      "acAsrConf": 252
    }, {
      "utf8": " yet",
      "tOffsetMs": 1140,
      "acAsrConf": 252
    }, {
      "utf8": " let's",
      "tOffsetMs": 1500,
      "acAsrConf": 216
    } ]
  }, {
    "tStartMs": 2048040,
    "dDurationMs": 1869,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2048050,
    "dDurationMs": 4170,
    "wWinId": 1,
    "segs": [ {
      "utf8": "give",
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 210,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 240,
      "acAsrConf": 232
    }, {
      "utf8": " chance",
      "tOffsetMs": 390,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 660,
      "acAsrConf": 201
    }, {
      "utf8": " answer",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " that",
      "tOffsetMs": 1410,
      "acAsrConf": 252
    }, {
      "utf8": " let's",
      "tOffsetMs": 1530,
      "acAsrConf": 217
    } ]
  }, {
    "tStartMs": 2049899,
    "dDurationMs": 2321,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2049909,
    "dDurationMs": 6061,
    "wWinId": 1,
    "segs": [ {
      "utf8": "see",
      "acAsrConf": 252
    }, {
      "utf8": " yes",
      "tOffsetMs": 181,
      "acAsrConf": 216
    }, {
      "utf8": " would",
      "tOffsetMs": 1111,
      "acAsrConf": 232
    }, {
      "utf8": " you",
      "tOffsetMs": 1411,
      "acAsrConf": 252
    }, {
      "utf8": " like",
      "tOffsetMs": 1441,
      "acAsrConf": 201
    }, {
      "utf8": " to",
      "tOffsetMs": 1561,
      "acAsrConf": 169
    }, {
      "utf8": " Rita",
      "tOffsetMs": 1770,
      "acAsrConf": 212
    }, {
      "utf8": " this",
      "tOffsetMs": 1921,
      "acAsrConf": 222
    }, {
      "utf8": " is",
      "tOffsetMs": 2101,
      "acAsrConf": 165
    } ]
  }, {
    "tStartMs": 2052210,
    "dDurationMs": 3760,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2052220,
    "dDurationMs": 6600,
    "wWinId": 1,
    "segs": [ {
      "utf8": "from",
      "acAsrConf": 252
    }, {
      "utf8": " Jefferson",
      "tOffsetMs": 330,
      "acAsrConf": 127
    }, {
      "utf8": " Silva",
      "tOffsetMs": 1230,
      "acAsrConf": 60
    }, {
      "utf8": " okay",
      "tOffsetMs": 1890,
      "acAsrConf": 153
    }, {
      "utf8": " so",
      "tOffsetMs": 2459,
      "acAsrConf": 187
    }, {
      "utf8": " it",
      "tOffsetMs": 2880,
      "acAsrConf": 252
    }, {
      "utf8": " says",
      "tOffsetMs": 3209,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 3449,
      "acAsrConf": 219
    } ]
  }, {
    "tStartMs": 2055960,
    "dDurationMs": 2860,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2055970,
    "dDurationMs": 5010,
    "wWinId": 1,
    "segs": [ {
      "utf8": "have",
      "acAsrConf": 204
    }, {
      "utf8": " a",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " lot",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 780,
      "acAsrConf": 127
    }, {
      "utf8": " difficult",
      "tOffsetMs": 929,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1530,
      "acAsrConf": 240
    }, {
      "utf8": " heard",
      "tOffsetMs": 1770,
      "acAsrConf": 186
    }, {
      "utf8": " English",
      "tOffsetMs": 2250,
      "acAsrConf": 246
    } ]
  }, {
    "tStartMs": 2058810,
    "dDurationMs": 2170,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2058820,
    "dDurationMs": 5730,
    "wWinId": 1,
    "segs": [ {
      "utf8": "so",
      "acAsrConf": 234
    }, {
      "utf8": " I",
      "tOffsetMs": 359,
      "acAsrConf": 251
    }, {
      "utf8": " think",
      "tOffsetMs": 390,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 540,
      "acAsrConf": 222
    }, {
      "utf8": " mean",
      "tOffsetMs": 810,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 960,
      "acAsrConf": 228
    }, {
      "utf8": " have",
      "tOffsetMs": 1200,
      "acAsrConf": 252
    }, {
      "utf8": " difficulty",
      "tOffsetMs": 1380,
      "acAsrConf": 239
    } ]
  }, {
    "tStartMs": 2060970,
    "dDurationMs": 3580,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2060980,
    "dDurationMs": 4590,
    "wWinId": 1,
    "segs": [ {
      "utf8": "listening",
      "acAsrConf": 246
    }, {
      "utf8": " and",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " understanding",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 2270,
      "acAsrConf": 160
    }, {
      "utf8": " what",
      "tOffsetMs": 3270,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2064540,
    "dDurationMs": 1030,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2064550,
    "dDurationMs": 2550,
    "wWinId": 1,
    "segs": [ {
      "utf8": "would",
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 180,
      "acAsrConf": 248
    }, {
      "utf8": " recommend",
      "tOffsetMs": 330,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2065560,
    "dDurationMs": 1540,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2065570,
    "dDurationMs": 3510,
    "wWinId": 1,
    "segs": [ {
      "utf8": "aha",
      "acAsrConf": 201
    }, {
      "utf8": " what",
      "tOffsetMs": 450,
      "acAsrConf": 125
    }, {
      "utf8": " would",
      "tOffsetMs": 780,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 900,
      "acAsrConf": 252
    }, {
      "utf8": " recommend",
      "tOffsetMs": 990,
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 1109,
      "acAsrConf": 229
    } ]
  }, {
    "tStartMs": 2067090,
    "dDurationMs": 1990,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2067100,
    "dDurationMs": 4860,
    "wWinId": 1,
    "segs": [ {
      "utf8": "listening",
      "acAsrConf": 244
    }, {
      "utf8": " and",
      "tOffsetMs": 569,
      "acAsrConf": 252
    }, {
      "utf8": " understanding",
      "tOffsetMs": 720,
      "acAsrConf": 200
    }, {
      "utf8": " English",
      "tOffsetMs": 870,
      "acAsrConf": 252
    }, {
      "utf8": " well",
      "tOffsetMs": 1410,
      "acAsrConf": 143
    } ]
  }, {
    "tStartMs": 2069070,
    "dDurationMs": 2890,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2069080,
    "dDurationMs": 5130,
    "wWinId": 1,
    "segs": [ {
      "utf8": "I",
      "acAsrConf": 252
    }, {
      "utf8": " think",
      "tOffsetMs": 30,
      "acAsrConf": 252
    }, {
      "utf8": " what",
      "tOffsetMs": 210,
      "acAsrConf": 237
    }, {
      "utf8": " we",
      "tOffsetMs": 540,
      "acAsrConf": 252
    }, {
      "utf8": " talked",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " about",
      "tOffsetMs": 930,
      "acAsrConf": 204
    }, {
      "utf8": " a",
      "tOffsetMs": 1110,
      "acAsrConf": 207
    }, {
      "utf8": " 70%",
      "tOffsetMs": 1260,
      "acAsrConf": 228
    }, {
      "utf8": " rule",
      "tOffsetMs": 2069,
      "acAsrConf": 206
    } ]
  }, {
    "tStartMs": 2071950,
    "dDurationMs": 2260,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2071960,
    "dDurationMs": 6480,
    "wWinId": 1,
    "segs": [ {
      "utf8": "that's",
      "acAsrConf": 239
    }, {
      "utf8": " really",
      "tOffsetMs": 570,
      "acAsrConf": 245
    }, {
      "utf8": " important",
      "tOffsetMs": 990,
      "acAsrConf": 252
    }, {
      "utf8": " because",
      "tOffsetMs": 1500,
      "acAsrConf": 236
    }, {
      "utf8": " if",
      "tOffsetMs": 1709,
      "acAsrConf": 248
    } ]
  }, {
    "tStartMs": 2074200,
    "dDurationMs": 4240,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2074210,
    "dDurationMs": 7980,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you're",
      "acAsrConf": 252
    }, {
      "utf8": " listening",
      "tOffsetMs": 660,
      "acAsrConf": 131
    }, {
      "utf8": " to",
      "tOffsetMs": 1199,
      "acAsrConf": 252
    }, {
      "utf8": " native",
      "tOffsetMs": 2030,
      "acAsrConf": 252
    }, {
      "utf8": " English",
      "tOffsetMs": 3030,
      "acAsrConf": 241
    }, {
      "utf8": " TV",
      "tOffsetMs": 3810,
      "acAsrConf": 219
    } ]
  }, {
    "tStartMs": 2078430,
    "dDurationMs": 3760,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2078440,
    "dDurationMs": 7200,
    "wWinId": 1,
    "segs": [ {
      "utf8": "show",
      "acAsrConf": 252
    }, {
      "utf8": " but",
      "tOffsetMs": 209,
      "acAsrConf": 207
    }, {
      "utf8": " you",
      "tOffsetMs": 630,
      "acAsrConf": 252
    }, {
      "utf8": " only",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " understand",
      "tOffsetMs": 959,
      "acAsrConf": 252
    }, {
      "utf8": " 70%",
      "tOffsetMs": 2030,
      "acAsrConf": 172
    }, {
      "utf8": " it's",
      "tOffsetMs": 3030,
      "acAsrConf": 228
    } ]
  }, {
    "tStartMs": 2082180,
    "dDurationMs": 3460,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2082190,
    "dDurationMs": 5160,
    "wWinId": 1,
    "segs": [ {
      "utf8": "going",
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 209,
      "acAsrConf": 252
    }, {
      "utf8": " be",
      "tOffsetMs": 300,
      "acAsrConf": 252
    }, {
      "utf8": " difficult",
      "tOffsetMs": 420,
      "acAsrConf": 200
    }, {
      "utf8": " to",
      "tOffsetMs": 959,
      "acAsrConf": 172
    }, {
      "utf8": " get",
      "tOffsetMs": 1380,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 2130,
      "acAsrConf": 233
    }, {
      "utf8": " 80%",
      "tOffsetMs": 2880,
      "acAsrConf": 86
    } ]
  }, {
    "tStartMs": 2085630,
    "dDurationMs": 1720,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2085640,
    "dDurationMs": 4920,
    "wWinId": 1,
    "segs": [ {
      "utf8": "more",
      "acAsrConf": 160
    }, {
      "utf8": " so",
      "tOffsetMs": 360,
      "acAsrConf": 217
    }, {
      "utf8": " I",
      "tOffsetMs": 630,
      "acAsrConf": 248
    }, {
      "utf8": " recommend",
      "tOffsetMs": 660,
      "acAsrConf": 236
    }, {
      "utf8": " listening",
      "tOffsetMs": 1289,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1470,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 2087340,
    "dDurationMs": 3220,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2087350,
    "dDurationMs": 7020,
    "wWinId": 1,
    "segs": [ {
      "utf8": "something",
      "acAsrConf": 147
    }, {
      "utf8": " that",
      "tOffsetMs": 540,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 569,
      "acAsrConf": 201
    }, {
      "utf8": " closer",
      "tOffsetMs": 930,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1799,
      "acAsrConf": 252
    }, {
      "utf8": " your",
      "tOffsetMs": 2040,
      "acAsrConf": 252
    }, {
      "utf8": " level",
      "tOffsetMs": 2850,
      "acAsrConf": 237
    } ]
  }, {
    "tStartMs": 2090550,
    "dDurationMs": 3820,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2090560,
    "dDurationMs": 7110,
    "wWinId": 1,
    "segs": [ {
      "utf8": "and",
      "acAsrConf": 181
    }, {
      "utf8": " after",
      "tOffsetMs": 800,
      "acAsrConf": 205
    }, {
      "utf8": " this",
      "tOffsetMs": 1800,
      "acAsrConf": 252
    }, {
      "utf8": " presentation",
      "tOffsetMs": 2099,
      "acAsrConf": 238
    }, {
      "utf8": " I",
      "tOffsetMs": 2460,
      "acAsrConf": 144
    }, {
      "utf8": " can",
      "tOffsetMs": 3180,
      "acAsrConf": 252
    }, {
      "utf8": " send",
      "tOffsetMs": 3240,
      "acAsrConf": 142
    } ]
  }, {
    "tStartMs": 2094360,
    "dDurationMs": 3310,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2094370,
    "dDurationMs": 6510,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 252
    }, {
      "utf8": " some",
      "tOffsetMs": 90,
      "acAsrConf": 145
    }, {
      "utf8": " links",
      "tOffsetMs": 180,
      "acAsrConf": 201
    }, {
      "utf8": " to",
      "tOffsetMs": 690,
      "acAsrConf": 252
    }, {
      "utf8": " websites",
      "tOffsetMs": 1790,
      "acAsrConf": 200
    }, {
      "utf8": " where",
      "tOffsetMs": 2790,
      "acAsrConf": 227
    }, {
      "utf8": " you",
      "tOffsetMs": 3090,
      "acAsrConf": 251
    }, {
      "utf8": " can",
      "tOffsetMs": 3270,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2097660,
    "dDurationMs": 3220,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2097670,
    "dDurationMs": 5700,
    "wWinId": 1,
    "segs": [ {
      "utf8": "listen",
      "acAsrConf": 187
    }, {
      "utf8": " to",
      "tOffsetMs": 420,
      "acAsrConf": 210
    }, {
      "utf8": " conversations",
      "tOffsetMs": 1040,
      "acAsrConf": 204
    }, {
      "utf8": " or",
      "tOffsetMs": 2040,
      "acAsrConf": 214
    }, {
      "utf8": " you",
      "tOffsetMs": 2910,
      "acAsrConf": 252
    }, {
      "utf8": " can",
      "tOffsetMs": 3030,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2100870,
    "dDurationMs": 2500,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2100880,
    "dDurationMs": 5460,
    "wWinId": 1,
    "segs": [ {
      "utf8": "listen",
      "acAsrConf": 223
    }, {
      "utf8": " to",
      "tOffsetMs": 180,
      "acAsrConf": 226
    }, {
      "utf8": " like",
      "tOffsetMs": 360,
      "acAsrConf": 188
    }, {
      "utf8": " a",
      "tOffsetMs": 810,
      "acAsrConf": 252
    }, {
      "utf8": " dialogue",
      "tOffsetMs": 840,
      "acAsrConf": 224
    }, {
      "utf8": " or",
      "tOffsetMs": 1470,
      "acAsrConf": 228
    }, {
      "utf8": " short",
      "tOffsetMs": 1830,
      "acAsrConf": 236
    }, {
      "utf8": " story",
      "tOffsetMs": 2130,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2103360,
    "dDurationMs": 2980,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2103370,
    "dDurationMs": 5820,
    "wWinId": 1,
    "segs": [ {
      "utf8": "with",
      "acAsrConf": 216
    }, {
      "utf8": " different",
      "tOffsetMs": 780,
      "acAsrConf": 212
    }, {
      "utf8": " levels",
      "tOffsetMs": 1440,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1620,
      "acAsrConf": 200
    }, {
      "utf8": " maybe",
      "tOffsetMs": 1920,
      "acAsrConf": 207
    }, {
      "utf8": " if",
      "tOffsetMs": 2370,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 2700,
      "acAsrConf": 248
    } ]
  }, {
    "tStartMs": 2106330,
    "dDurationMs": 2860,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2106340,
    "dDurationMs": 5460,
    "wWinId": 1,
    "segs": [ {
      "utf8": "are",
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 30,
      "acAsrConf": 200
    }, {
      "utf8": " beginner",
      "tOffsetMs": 690,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1410,
      "acAsrConf": 187
    }, {
      "utf8": " can",
      "tOffsetMs": 2250,
      "acAsrConf": 145
    }, {
      "utf8": " listen",
      "tOffsetMs": 2490,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 2700,
      "acAsrConf": 227
    } ]
  }, {
    "tStartMs": 2109180,
    "dDurationMs": 2620,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2109190,
    "dDurationMs": 4890,
    "wWinId": 1,
    "segs": [ {
      "utf8": "beginner",
      "acAsrConf": 196
    }, {
      "utf8": " stories",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1470,
      "acAsrConf": 245
    }, {
      "utf8": " then",
      "tOffsetMs": 1860,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 2160,
      "acAsrConf": 252
    }, {
      "utf8": " can",
      "tOffsetMs": 2340,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2111790,
    "dDurationMs": 2290,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2111800,
    "dDurationMs": 4140,
    "wWinId": 1,
    "segs": [ {
      "utf8": "improve",
      "acAsrConf": 237
    }, {
      "utf8": " your",
      "tOffsetMs": 840,
      "acAsrConf": 252
    }, {
      "utf8": " English",
      "tOffsetMs": 1020,
      "acAsrConf": 247
    }, {
      "utf8": " in",
      "tOffsetMs": 1080,
      "acAsrConf": 249
    }, {
      "utf8": " that",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    }, {
      "utf8": " way",
      "tOffsetMs": 1620,
      "acAsrConf": 232
    }, {
      "utf8": " so",
      "tOffsetMs": 1830,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 1860,
      "acAsrConf": 228
    } ]
  }, {
    "tStartMs": 2114070,
    "dDurationMs": 1870,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2114080,
    "dDurationMs": 4029,
    "wWinId": 1,
    "segs": [ {
      "utf8": "recommend",
      "acAsrConf": 252
    }, {
      "utf8": " using",
      "tOffsetMs": 480,
      "acAsrConf": 236
    }, {
      "utf8": " something",
      "tOffsetMs": 810,
      "acAsrConf": 252
    }, {
      "utf8": " that's",
      "tOffsetMs": 1200,
      "acAsrConf": 175
    }, {
      "utf8": " closer",
      "tOffsetMs": 1530,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2115930,
    "dDurationMs": 2179,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2115940,
    "dDurationMs": 5349,
    "wWinId": 1,
    "segs": [ {
      "utf8": "to",
      "acAsrConf": 252
    }, {
      "utf8": " your",
      "tOffsetMs": 180,
      "acAsrConf": 252
    }, {
      "utf8": " level",
      "tOffsetMs": 330,
      "acAsrConf": 146
    }, {
      "utf8": " and",
      "tOffsetMs": 810,
      "acAsrConf": 107
    }, {
      "utf8": " I",
      "tOffsetMs": 1350,
      "acAsrConf": 130
    } ]
  }, {
    "tStartMs": 2118099,
    "dDurationMs": 3190,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2118109,
    "dDurationMs": 6630,
    "wWinId": 1,
    "segs": [ {
      "utf8": "would",
      "acAsrConf": 252
    }, {
      "utf8": " also",
      "tOffsetMs": 91,
      "acAsrConf": 252
    }, {
      "utf8": " recommend",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " maybe",
      "tOffsetMs": 720,
      "acAsrConf": 228
    }, {
      "utf8": " just",
      "tOffsetMs": 1291,
      "acAsrConf": 228
    }, {
      "utf8": " practice",
      "tOffsetMs": 2180,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2121279,
    "dDurationMs": 3460,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2121289,
    "dDurationMs": 6060,
    "wWinId": 1,
    "segs": [ {
      "utf8": "practice",
      "acAsrConf": 252
    }, {
      "utf8": " track",
      "tOffsetMs": 240,
      "acAsrConf": 167
    }, {
      "utf8": " as",
      "tOffsetMs": 1080,
      "acAsrConf": 208
    }, {
      "utf8": " we",
      "tOffsetMs": 1921,
      "acAsrConf": 252
    }, {
      "utf8": " said",
      "tOffsetMs": 2220,
      "acAsrConf": 210
    }, {
      "utf8": " hard",
      "tOffsetMs": 2431,
      "acAsrConf": 236
    }, {
      "utf8": " truth",
      "tOffsetMs": 2700,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 3151,
      "acAsrConf": 217
    } ]
  }, {
    "tStartMs": 2124729,
    "dDurationMs": 2620,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2124739,
    "dDurationMs": 6901,
    "wWinId": 1,
    "segs": [ {
      "utf8": "only",
      "acAsrConf": 211
    }, {
      "utf8": " way",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 691,
      "acAsrConf": 252
    }, {
      "utf8": " can",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " really",
      "tOffsetMs": 1201,
      "acAsrConf": 201
    }, {
      "utf8": " get",
      "tOffsetMs": 1681,
      "acAsrConf": 252
    }, {
      "utf8": " better",
      "tOffsetMs": 1860,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 2100,
      "acAsrConf": 226
    }, {
      "utf8": " to",
      "tOffsetMs": 2580,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2127339,
    "dDurationMs": 4301,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2127349,
    "dDurationMs": 6900,
    "wWinId": 1,
    "segs": [ {
      "utf8": "listen",
      "acAsrConf": 210
    }, {
      "utf8": " more",
      "tOffsetMs": 450,
      "acAsrConf": 229
    }, {
      "utf8": " hmm",
      "tOffsetMs": 1410,
      "acAsrConf": 208
    }, {
      "utf8": " and",
      "tOffsetMs": 2190,
      "acAsrConf": 112
    }, {
      "utf8": " regularly",
      "tOffsetMs": 2430,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 3210,
      "acAsrConf": 215
    }, {
      "utf8": " it's",
      "tOffsetMs": 3720,
      "acAsrConf": 244
    } ]
  }, {
    "tStartMs": 2131630,
    "dDurationMs": 2619,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2131640,
    "dDurationMs": 6229,
    "wWinId": 1,
    "segs": [ {
      "utf8": "very",
      "acAsrConf": 252
    }, {
      "utf8": " difficult",
      "tOffsetMs": 209,
      "acAsrConf": 186
    }, {
      "utf8": " but",
      "tOffsetMs": 869,
      "acAsrConf": 233
    }, {
      "utf8": " it's",
      "tOffsetMs": 1080,
      "acAsrConf": 188
    }, {
      "utf8": " start",
      "tOffsetMs": 1260,
      "acAsrConf": 234
    }, {
      "utf8": " at",
      "tOffsetMs": 1949,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 2189,
      "acAsrConf": 252
    }, {
      "utf8": " level",
      "tOffsetMs": 2219,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2134239,
    "dDurationMs": 3630,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2134249,
    "dDurationMs": 6360,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 78
    }, {
      "utf8": " can",
      "tOffsetMs": 360,
      "acAsrConf": 252
    }, {
      "utf8": " understand",
      "tOffsetMs": 570,
      "acAsrConf": 252
    }, {
      "utf8": " most",
      "tOffsetMs": 1290,
      "acAsrConf": 232
    }, {
      "utf8": " of",
      "tOffsetMs": 1981,
      "acAsrConf": 222
    }, {
      "utf8": " it",
      "tOffsetMs": 2550,
      "acAsrConf": 197
    }, {
      "utf8": " but",
      "tOffsetMs": 3151,
      "acAsrConf": 233
    } ]
  }, {
    "tStartMs": 2137859,
    "dDurationMs": 2750,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2137869,
    "dDurationMs": 5111,
    "wWinId": 1,
    "segs": [ {
      "utf8": "eventually",
      "acAsrConf": 213
    }, {
      "utf8": " try",
      "tOffsetMs": 1000,
      "acAsrConf": 228
    }, {
      "utf8": " to",
      "tOffsetMs": 1541,
      "acAsrConf": 203
    }, {
      "utf8": " keep",
      "tOffsetMs": 1601,
      "acAsrConf": 252
    }, {
      "utf8": " on",
      "tOffsetMs": 1900,
      "acAsrConf": 252
    }, {
      "utf8": " going",
      "tOffsetMs": 2111,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 2380,
      "acAsrConf": 209
    } ]
  }, {
    "tStartMs": 2140599,
    "dDurationMs": 2381,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2140609,
    "dDurationMs": 5101,
    "wWinId": 1,
    "segs": [ {
      "utf8": "build",
      "acAsrConf": 201
    }, {
      "utf8": " yourself",
      "tOffsetMs": 271,
      "acAsrConf": 252
    }, {
      "utf8": " up",
      "tOffsetMs": 480,
      "acAsrConf": 165
    }, {
      "utf8": " and",
      "tOffsetMs": 930,
      "acAsrConf": 194
    }, {
      "utf8": " then",
      "tOffsetMs": 1200,
      "acAsrConf": 188
    }, {
      "utf8": " you",
      "tOffsetMs": 1351,
      "acAsrConf": 252
    }, {
      "utf8": " also",
      "tOffsetMs": 1950,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 2101,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2142970,
    "dDurationMs": 2740,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2142980,
    "dDurationMs": 5250,
    "wWinId": 1,
    "segs": [ {
      "utf8": "to",
      "acAsrConf": 252
    }, {
      "utf8": " practice",
      "tOffsetMs": 29,
      "acAsrConf": 226
    }, {
      "utf8": " speaking",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " with",
      "tOffsetMs": 1349,
      "acAsrConf": 252
    }, {
      "utf8": " real",
      "tOffsetMs": 1559,
      "acAsrConf": 252
    }, {
      "utf8": " people",
      "tOffsetMs": 2069,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 2309,
      "acAsrConf": 209
    } ]
  }, {
    "tStartMs": 2145700,
    "dDurationMs": 2530,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2145710,
    "dDurationMs": 4980,
    "wWinId": 1,
    "segs": [ {
      "utf8": "actually",
      "acAsrConf": 252
    }, {
      "utf8": " speaking",
      "tOffsetMs": 680,
      "acAsrConf": 218
    }, {
      "utf8": " yourself",
      "tOffsetMs": 1680,
      "acAsrConf": 252
    }, {
      "utf8": " knows",
      "tOffsetMs": 2250,
      "acAsrConf": 83
    } ]
  }, {
    "tStartMs": 2148220,
    "dDurationMs": 2470,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2148230,
    "dDurationMs": 4859,
    "wWinId": 1,
    "segs": [ {
      "utf8": "actually",
      "acAsrConf": 226
    }, {
      "utf8": " doing",
      "tOffsetMs": 329,
      "acAsrConf": 147
    }, {
      "utf8": " it",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " creates",
      "tOffsetMs": 1109,
      "acAsrConf": 63
    }, {
      "utf8": " uh",
      "tOffsetMs": 1529,
      "acAsrConf": 224
    }, {
      "utf8": " and",
      "tOffsetMs": 1859,
      "acAsrConf": 242
    } ]
  }, {
    "tStartMs": 2150680,
    "dDurationMs": 2409,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2150690,
    "dDurationMs": 5040,
    "wWinId": 1,
    "segs": [ {
      "utf8": "speaking",
      "acAsrConf": 218
    }, {
      "utf8": " it",
      "tOffsetMs": 869,
      "acAsrConf": 252
    }, {
      "utf8": " will",
      "tOffsetMs": 1079,
      "acAsrConf": 252
    }, {
      "utf8": " help",
      "tOffsetMs": 1349,
      "acAsrConf": 252
    }, {
      "utf8": " your",
      "tOffsetMs": 1589,
      "acAsrConf": 227
    }, {
      "utf8": " listening",
      "tOffsetMs": 1710,
      "acAsrConf": 167
    }, {
      "utf8": " to",
      "tOffsetMs": 2159,
      "acAsrConf": 208
    } ]
  }, {
    "tStartMs": 2153079,
    "dDurationMs": 2651,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2153089,
    "dDurationMs": 4710,
    "wWinId": 1,
    "segs": [ {
      "utf8": "so",
      "acAsrConf": 249
    }, {
      "utf8": " just",
      "tOffsetMs": 361,
      "acAsrConf": 236
    }, {
      "utf8": " general",
      "tOffsetMs": 571,
      "acAsrConf": 226
    }, {
      "utf8": " practice",
      "tOffsetMs": 1381,
      "acAsrConf": 247
    }, {
      "utf8": " yes",
      "tOffsetMs": 1861,
      "acAsrConf": 202
    }, {
      "utf8": " great",
      "tOffsetMs": 2190,
      "acAsrConf": 234
    } ]
  }, {
    "tStartMs": 2155720,
    "dDurationMs": 2079,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2155730,
    "dDurationMs": 5819,
    "wWinId": 1,
    "segs": [ {
      "utf8": "question",
      "acAsrConf": 252
    }, {
      "utf8": " Jefferson",
      "tOffsetMs": 259,
      "acAsrConf": 238
    }, {
      "utf8": " let's",
      "tOffsetMs": 1259,
      "acAsrConf": 227
    }, {
      "utf8": " go",
      "tOffsetMs": 1529,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1710,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 1769,
      "acAsrConf": 252
    }, {
      "utf8": " next",
      "tOffsetMs": 1920,
      "acAsrConf": 216
    } ]
  }, {
    "tStartMs": 2157789,
    "dDurationMs": 3760,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2157799,
    "dDurationMs": 7681,
    "wWinId": 1,
    "segs": [ {
      "utf8": "question",
      "acAsrConf": 217
    }, {
      "utf8": " all",
      "tOffsetMs": 121,
      "acAsrConf": 84
    }, {
      "utf8": " right",
      "tOffsetMs": 631,
      "acAsrConf": 252
    }, {
      "utf8": " we",
      "tOffsetMs": 690,
      "acAsrConf": 156
    }, {
      "utf8": " have",
      "tOffsetMs": 1020,
      "acAsrConf": 203
    }, {
      "utf8": " evaldo",
      "tOffsetMs": 1940,
      "acAsrConf": 240
    }, {
      "utf8": " evaldo",
      "tOffsetMs": 2940,
      "acAsrConf": 226
    } ]
  }, {
    "tStartMs": 2161539,
    "dDurationMs": 3941,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2161549,
    "dDurationMs": 6960,
    "wWinId": 1,
    "segs": [ {
      "utf8": "hello",
      "acAsrConf": 252
    }, {
      "utf8": " London",
      "tOffsetMs": 450,
      "acAsrConf": 186
    }, {
      "utf8": " book",
      "tOffsetMs": 1351,
      "acAsrConf": 78
    }, {
      "utf8": " Korea",
      "tOffsetMs": 1740,
      "acAsrConf": 230
    }, {
      "utf8": " yes",
      "tOffsetMs": 2161,
      "acAsrConf": 215
    }, {
      "utf8": " Korea",
      "tOffsetMs": 2820,
      "acAsrConf": 239
    }, {
      "utf8": " maybe",
      "tOffsetMs": 3480,
      "acAsrConf": 123
    } ]
  }, {
    "tStartMs": 2165470,
    "dDurationMs": 3039,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2165480,
    "dDurationMs": 4799,
    "wWinId": 1,
    "segs": [ {
      "utf8": "I",
      "acAsrConf": 215
    }, {
      "utf8": " think",
      "tOffsetMs": 359,
      "acAsrConf": 252
    }, {
      "utf8": " you're",
      "tOffsetMs": 539,
      "acAsrConf": 201
    }, {
      "utf8": " from",
      "tOffsetMs": 779,
      "acAsrConf": 252
    }, {
      "utf8": " Brazil",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " okay",
      "tOffsetMs": 1109,
      "acAsrConf": 102
    }, {
      "utf8": " let's",
      "tOffsetMs": 2029,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 2168499,
    "dDurationMs": 1780,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2168509,
    "dDurationMs": 4530,
    "wWinId": 1,
    "segs": [ {
      "utf8": "see",
      "acAsrConf": 252
    }, {
      "utf8": " what",
      "tOffsetMs": 151,
      "acAsrConf": 212
    }, {
      "utf8": " does",
      "tOffsetMs": 600,
      "acAsrConf": 242
    }, {
      "utf8": " she",
      "tOffsetMs": 780,
      "acAsrConf": 201
    }, {
      "utf8": " say",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    }, {
      "utf8": " would",
      "tOffsetMs": 1290,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1471,
      "acAsrConf": 189
    }, {
      "utf8": " like",
      "tOffsetMs": 1530,
      "acAsrConf": 206
    }, {
      "utf8": " to",
      "tOffsetMs": 1651,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2170269,
    "dDurationMs": 2770,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2170279,
    "dDurationMs": 5730,
    "wWinId": 1,
    "segs": [ {
      "utf8": "research",
      "acAsrConf": 0
    }, {
      "utf8": " do",
      "tOffsetMs": 330,
      "acAsrConf": 200
    }, {
      "utf8": " you",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " think",
      "tOffsetMs": 901,
      "acAsrConf": 206
    }, {
      "utf8": " music",
      "tOffsetMs": 1230,
      "acAsrConf": 255
    }, {
      "utf8": " it",
      "tOffsetMs": 1560,
      "acAsrConf": 200
    }, {
      "utf8": " could",
      "tOffsetMs": 2191,
      "acAsrConf": 237
    }, {
      "utf8": " be",
      "tOffsetMs": 2550,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2173029,
    "dDurationMs": 2980,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2173039,
    "dDurationMs": 6391,
    "wWinId": 1,
    "segs": [ {
      "utf8": "an",
      "acAsrConf": 252
    }, {
      "utf8": " effective",
      "tOffsetMs": 270,
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 950,
      "acAsrConf": 230
    }, {
      "utf8": " learning",
      "tOffsetMs": 1950,
      "acAsrConf": 252
    }, {
      "utf8": " English",
      "tOffsetMs": 2101,
      "acAsrConf": 206
    }, {
      "utf8": " or",
      "tOffsetMs": 2730,
      "acAsrConf": 207
    } ]
  }, {
    "tStartMs": 2175999,
    "dDurationMs": 3431,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2176009,
    "dDurationMs": 4951,
    "wWinId": 1,
    "segs": [ {
      "utf8": "better",
      "acAsrConf": 227
    }, {
      "utf8": " improve",
      "tOffsetMs": 480,
      "acAsrConf": 240
    }, {
      "utf8": " ha",
      "tOffsetMs": 1230,
      "acAsrConf": 147
    }, {
      "utf8": " ha",
      "tOffsetMs": 1711,
      "acAsrConf": 151
    }, {
      "utf8": " music",
      "tOffsetMs": 1971,
      "acAsrConf": 249
    }, {
      "utf8": " could",
      "tOffsetMs": 2971,
      "acAsrConf": 144
    }, {
      "utf8": " be",
      "tOffsetMs": 3360,
      "acAsrConf": 200
    } ]
  }, {
    "tStartMs": 2179420,
    "dDurationMs": 1540,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2179430,
    "dDurationMs": 4260,
    "wWinId": 1,
    "segs": [ {
      "utf8": "effective",
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 419,
      "acAsrConf": 252
    }, {
      "utf8": " learning",
      "tOffsetMs": 540,
      "acAsrConf": 213
    }, {
      "utf8": " English",
      "tOffsetMs": 659,
      "acAsrConf": 202
    }, {
      "utf8": " better",
      "tOffsetMs": 1169,
      "acAsrConf": 224
    } ]
  }, {
    "tStartMs": 2180950,
    "dDurationMs": 2740,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2180960,
    "dDurationMs": 5069,
    "wWinId": 1,
    "segs": [ {
      "utf8": "improving",
      "acAsrConf": 251
    }, {
      "utf8": " yes",
      "tOffsetMs": 569,
      "acAsrConf": 217
    }, {
      "utf8": " definitely",
      "tOffsetMs": 1490,
      "acAsrConf": 202
    }, {
      "utf8": " I",
      "tOffsetMs": 2490,
      "acAsrConf": 237
    }, {
      "utf8": " think",
      "tOffsetMs": 2670,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2183680,
    "dDurationMs": 2349,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2183690,
    "dDurationMs": 4679,
    "wWinId": 1,
    "segs": [ {
      "utf8": "learning",
      "acAsrConf": 200
    }, {
      "utf8": " English",
      "tOffsetMs": 629,
      "acAsrConf": 236
    }, {
      "utf8": " is",
      "tOffsetMs": 960,
      "acAsrConf": 240
    }, {
      "utf8": " a",
      "tOffsetMs": 1169,
      "acAsrConf": 252
    }, {
      "utf8": " great",
      "tOffsetMs": 1230,
      "acAsrConf": 252
    }, {
      "utf8": " way",
      "tOffsetMs": 1710,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1889,
      "acAsrConf": 252
    }, {
      "utf8": " hear",
      "tOffsetMs": 1919,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2186019,
    "dDurationMs": 2350,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2186029,
    "dDurationMs": 5730,
    "wWinId": 1,
    "segs": [ {
      "utf8": "the",
      "acAsrConf": 252
    }, {
      "utf8": " rhythm",
      "tOffsetMs": 181,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 421,
      "acAsrConf": 201
    }, {
      "utf8": " to",
      "tOffsetMs": 990,
      "acAsrConf": 252
    }, {
      "utf8": " memorize",
      "tOffsetMs": 1171,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 1560,
      "acAsrConf": 193
    }, {
      "utf8": " words",
      "tOffsetMs": 2101,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2188359,
    "dDurationMs": 3400,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2188369,
    "dDurationMs": 6031,
    "wWinId": 1,
    "segs": [ {
      "utf8": "well",
      "acAsrConf": 216
    }, {
      "utf8": " especially",
      "tOffsetMs": 720,
      "acAsrConf": 249
    }, {
      "utf8": " if",
      "tOffsetMs": 1170,
      "acAsrConf": 219
    }, {
      "utf8": " you",
      "tOffsetMs": 1831,
      "acAsrConf": 252
    }, {
      "utf8": " really",
      "tOffsetMs": 2160,
      "acAsrConf": 252
    }, {
      "utf8": " enjoy",
      "tOffsetMs": 2611,
      "acAsrConf": 202
    }, {
      "utf8": " it",
      "tOffsetMs": 3150,
      "acAsrConf": 204
    } ]
  }, {
    "tStartMs": 2191749,
    "dDurationMs": 2651,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2191759,
    "dDurationMs": 7020,
    "wWinId": 1,
    "segs": [ {
      "utf8": "then",
      "acAsrConf": 226
    }, {
      "utf8": " you",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " will",
      "tOffsetMs": 840,
      "acAsrConf": 215
    }, {
      "utf8": " listen",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    }, {
      "utf8": " more",
      "tOffsetMs": 1320,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1530,
      "acAsrConf": 201
    }, {
      "utf8": " you",
      "tOffsetMs": 2131,
      "acAsrConf": 218
    }, {
      "utf8": " can",
      "tOffsetMs": 2431,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2194390,
    "dDurationMs": 4389,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2194400,
    "dDurationMs": 6510,
    "wWinId": 1,
    "segs": [ {
      "utf8": "maybe",
      "acAsrConf": 252
    }, {
      "utf8": " buy",
      "tOffsetMs": 540,
      "acAsrConf": 250
    }, {
      "utf8": " or",
      "tOffsetMs": 929,
      "acAsrConf": 251
    }, {
      "utf8": " find",
      "tOffsetMs": 1669,
      "acAsrConf": 251
    }, {
      "utf8": " the",
      "tOffsetMs": 2669,
      "acAsrConf": 207
    }, {
      "utf8": " lyrics",
      "tOffsetMs": 2909,
      "acAsrConf": 252
    }, {
      "utf8": " online",
      "tOffsetMs": 3119,
      "acAsrConf": 200
    }, {
      "utf8": " and",
      "tOffsetMs": 4109,
      "acAsrConf": 249
    } ]
  }, {
    "tStartMs": 2198769,
    "dDurationMs": 2141,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2198779,
    "dDurationMs": 4111,
    "wWinId": 1,
    "segs": [ {
      "utf8": "read",
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 151,
      "acAsrConf": 252
    }, {
      "utf8": " lyrics",
      "tOffsetMs": 181,
      "acAsrConf": 128
    }, {
      "utf8": " and",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " speak",
      "tOffsetMs": 990,
      "acAsrConf": 236
    }, {
      "utf8": " the",
      "tOffsetMs": 1351,
      "acAsrConf": 252
    }, {
      "utf8": " lyrics",
      "tOffsetMs": 1530,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1891,
      "acAsrConf": 225
    } ]
  }, {
    "tStartMs": 2200900,
    "dDurationMs": 1990,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2200910,
    "dDurationMs": 4010,
    "wWinId": 1,
    "segs": [ {
      "utf8": "sing",
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " songs",
      "tOffsetMs": 419,
      "acAsrConf": 227
    }, {
      "utf8": " your",
      "tOffsetMs": 689,
      "acAsrConf": 166
    }, {
      "utf8": " I",
      "tOffsetMs": 1050,
      "acAsrConf": 239
    }, {
      "utf8": " think",
      "tOffsetMs": 1320,
      "acAsrConf": 207
    }, {
      "utf8": " the",
      "tOffsetMs": 1530,
      "acAsrConf": 236
    }, {
      "utf8": " only",
      "tOffsetMs": 1619,
      "acAsrConf": 161
    } ]
  }, {
    "tStartMs": 2202880,
    "dDurationMs": 2040,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2202890,
    "dDurationMs": 7580,
    "wWinId": 1,
    "segs": [ {
      "utf8": "negative",
      "acAsrConf": 252
    }, {
      "utf8": " thing",
      "tOffsetMs": 570,
      "acAsrConf": 252
    }, {
      "utf8": " about",
      "tOffsetMs": 599,
      "acAsrConf": 78
    }, {
      "utf8": " music",
      "tOffsetMs": 1080,
      "acAsrConf": 126
    }, {
      "utf8": " is",
      "tOffsetMs": 1590,
      "acAsrConf": 227
    }, {
      "utf8": " that",
      "tOffsetMs": 1800,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2204910,
    "dDurationMs": 5560,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2204920,
    "dDurationMs": 8470,
    "wWinId": 1,
    "segs": [ {
      "utf8": "sometimes",
      "acAsrConf": 206
    }, {
      "utf8": " the",
      "tOffsetMs": 1000,
      "acAsrConf": 241
    }, {
      "utf8": " lyrics",
      "tOffsetMs": 1839,
      "acAsrConf": 252
    }, {
      "utf8": " are",
      "tOffsetMs": 2079,
      "acAsrConf": 201
    }, {
      "utf8": " like",
      "tOffsetMs": 2740,
      "acAsrConf": 223
    }, {
      "utf8": " poetry",
      "tOffsetMs": 3399,
      "acAsrConf": 245
    }, {
      "utf8": " so",
      "tOffsetMs": 4199,
      "acAsrConf": 166
    } ]
  }, {
    "tStartMs": 2210460,
    "dDurationMs": 2930,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2210470,
    "dDurationMs": 5740,
    "wWinId": 1,
    "segs": [ {
      "utf8": "maybe",
      "acAsrConf": 240
    }, {
      "utf8": " it's",
      "tOffsetMs": 1000,
      "acAsrConf": 255
    }, {
      "utf8": " it's",
      "tOffsetMs": 1389,
      "acAsrConf": 244
    }, {
      "utf8": " good",
      "tOffsetMs": 1839,
      "acAsrConf": 252
    }, {
      "utf8": " grammar",
      "tOffsetMs": 2079,
      "acAsrConf": 228
    }, {
      "utf8": " but",
      "tOffsetMs": 2379,
      "acAsrConf": 252
    }, {
      "utf8": " it's",
      "tOffsetMs": 2740,
      "acAsrConf": 215
    } ]
  }, {
    "tStartMs": 2213380,
    "dDurationMs": 2830,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2213390,
    "dDurationMs": 5699,
    "wWinId": 1,
    "segs": [ {
      "utf8": "more",
      "acAsrConf": 252
    }, {
      "utf8": " poetic",
      "tOffsetMs": 209,
      "acAsrConf": 216
    }, {
      "utf8": " grammar",
      "tOffsetMs": 629,
      "acAsrConf": 242
    }, {
      "utf8": " so",
      "tOffsetMs": 1229,
      "acAsrConf": 200
    }, {
      "utf8": " that's",
      "tOffsetMs": 2010,
      "acAsrConf": 252
    }, {
      "utf8": " just",
      "tOffsetMs": 2580,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2216200,
    "dDurationMs": 2889,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2216210,
    "dDurationMs": 5309,
    "wWinId": 1,
    "segs": [ {
      "utf8": "something",
      "acAsrConf": 226
    }, {
      "utf8": " to",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " keep",
      "tOffsetMs": 450,
      "acAsrConf": 177
    }, {
      "utf8": " in",
      "tOffsetMs": 659,
      "acAsrConf": 231
    }, {
      "utf8": " mind",
      "tOffsetMs": 869,
      "acAsrConf": 252
    }, {
      "utf8": " hmm",
      "tOffsetMs": 1099,
      "acAsrConf": 201
    }, {
      "utf8": " yes",
      "tOffsetMs": 2159,
      "acAsrConf": 226
    }, {
      "utf8": " maybe",
      "tOffsetMs": 2579,
      "acAsrConf": 246
    } ]
  }, {
    "tStartMs": 2219079,
    "dDurationMs": 2440,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2219089,
    "dDurationMs": 5821,
    "wWinId": 1,
    "segs": [ {
      "utf8": "not",
      "acAsrConf": 250
    }, {
      "utf8": " conversational",
      "tOffsetMs": 291,
      "acAsrConf": 233
    }, {
      "utf8": " English",
      "tOffsetMs": 1291,
      "acAsrConf": 243
    }, {
      "utf8": " but",
      "tOffsetMs": 1710,
      "acAsrConf": 220
    }, {
      "utf8": " you",
      "tOffsetMs": 2041,
      "acAsrConf": 252
    }, {
      "utf8": " can",
      "tOffsetMs": 2071,
      "acAsrConf": 133
    } ]
  }, {
    "tStartMs": 2221509,
    "dDurationMs": 3401,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2221519,
    "dDurationMs": 5671,
    "wWinId": 1,
    "segs": [ {
      "utf8": "certainly",
      "acAsrConf": 201
    }, {
      "utf8": " learn",
      "tOffsetMs": 330,
      "acAsrConf": 252
    }, {
      "utf8": " idioms",
      "tOffsetMs": 840,
      "acAsrConf": 252
    }, {
      "utf8": " or",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    }, {
      "utf8": " phrases",
      "tOffsetMs": 2191,
      "acAsrConf": 242
    }, {
      "utf8": " right",
      "tOffsetMs": 2580,
      "acAsrConf": 195
    } ]
  }, {
    "tStartMs": 2224900,
    "dDurationMs": 2290,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2224910,
    "dDurationMs": 4260,
    "wWinId": 1,
    "segs": [ {
      "utf8": "right",
      "acAsrConf": 203
    }, {
      "utf8": " stuff",
      "tOffsetMs": 570,
      "acAsrConf": 224
    }, {
      "utf8": " like",
      "tOffsetMs": 959,
      "acAsrConf": 251
    }, {
      "utf8": " that",
      "tOffsetMs": 1139,
      "acAsrConf": 207
    }, {
      "utf8": " I",
      "tOffsetMs": 1320,
      "acAsrConf": 123
    }, {
      "utf8": " know",
      "tOffsetMs": 1530,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 1679,
      "acAsrConf": 255
    }, {
      "utf8": " listen",
      "tOffsetMs": 1770,
      "acAsrConf": 227
    }, {
      "utf8": " to",
      "tOffsetMs": 2040,
      "acAsrConf": 148
    } ]
  }, {
    "tStartMs": 2227180,
    "dDurationMs": 1990,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2227190,
    "dDurationMs": 4530,
    "wWinId": 1,
    "segs": [ {
      "utf8": "a",
      "acAsrConf": 252
    }, {
      "utf8": " lot",
      "tOffsetMs": 30,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 119,
      "acAsrConf": 126
    }, {
      "utf8": " French",
      "tOffsetMs": 270,
      "acAsrConf": 206
    }, {
      "utf8": " music",
      "tOffsetMs": 540,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 839,
      "acAsrConf": 208
    }, {
      "utf8": " I",
      "tOffsetMs": 1200,
      "acAsrConf": 249
    }, {
      "utf8": " love",
      "tOffsetMs": 1379,
      "acAsrConf": 241
    }, {
      "utf8": " singing",
      "tOffsetMs": 1679,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2229160,
    "dDurationMs": 2560,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2229170,
    "dDurationMs": 5040,
    "wWinId": 1,
    "segs": [ {
      "utf8": "along",
      "acAsrConf": 207
    }, {
      "utf8": " with",
      "tOffsetMs": 419,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 449,
      "acAsrConf": 202
    }, {
      "utf8": " and",
      "tOffsetMs": 839,
      "acAsrConf": 222
    }, {
      "utf8": " hearing",
      "tOffsetMs": 1079,
      "acAsrConf": 252
    }, {
      "utf8": " phrases",
      "tOffsetMs": 1740,
      "acAsrConf": 226
    }, {
      "utf8": " that",
      "tOffsetMs": 2099,
      "acAsrConf": 207
    }, {
      "utf8": " I",
      "tOffsetMs": 2490,
      "acAsrConf": 206
    } ]
  }, {
    "tStartMs": 2231710,
    "dDurationMs": 2500,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2231720,
    "dDurationMs": 4649,
    "wWinId": 1,
    "segs": [ {
      "utf8": "already",
      "acAsrConf": 252
    }, {
      "utf8": " know",
      "tOffsetMs": 329,
      "acAsrConf": 206
    }, {
      "utf8": " in",
      "tOffsetMs": 779,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 1019,
      "acAsrConf": 252
    }, {
      "utf8": " music",
      "tOffsetMs": 1139,
      "acAsrConf": 204
    }, {
      "utf8": " so",
      "tOffsetMs": 1529,
      "acAsrConf": 201
    }, {
      "utf8": " yes",
      "tOffsetMs": 1829,
      "acAsrConf": 241
    }, {
      "utf8": " go",
      "tOffsetMs": 2220,
      "acAsrConf": 245
    } ]
  }, {
    "tStartMs": 2234200,
    "dDurationMs": 2169,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2234210,
    "dDurationMs": 7170,
    "wWinId": 1,
    "segs": [ {
      "utf8": "ahead",
      "acAsrConf": 248
    }, {
      "utf8": " and",
      "tOffsetMs": 119,
      "acAsrConf": 201
    }, {
      "utf8": " listen",
      "tOffsetMs": 389,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 539,
      "acAsrConf": 233
    }, {
      "utf8": " music",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " let's",
      "tOffsetMs": 869,
      "acAsrConf": 167
    }, {
      "utf8": " see",
      "tOffsetMs": 1829,
      "acAsrConf": 252
    }, {
      "utf8": " what",
      "tOffsetMs": 1950,
      "acAsrConf": 244
    } ]
  }, {
    "tStartMs": 2236359,
    "dDurationMs": 5021,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2236369,
    "dDurationMs": 8551,
    "wWinId": 1,
    "segs": [ {
      "utf8": "is",
      "acAsrConf": 252
    }, {
      "utf8": " another",
      "tOffsetMs": 120,
      "acAsrConf": 255
    }, {
      "utf8": " great",
      "tOffsetMs": 361,
      "acAsrConf": 226
    }, {
      "utf8": " question",
      "tOffsetMs": 660,
      "acAsrConf": 201
    }, {
      "utf8": " we",
      "tOffsetMs": 870,
      "acAsrConf": 227
    }, {
      "utf8": " have",
      "tOffsetMs": 1200,
      "acAsrConf": 252
    }, {
      "utf8": " mm-hmm",
      "tOffsetMs": 2930,
      "acAsrConf": 226
    } ]
  }, {
    "tStartMs": 2241370,
    "dDurationMs": 3550,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2241380,
    "dDurationMs": 5369,
    "wWinId": 1,
    "segs": [ {
      "utf8": "all",
      "acAsrConf": 217
    }, {
      "utf8": " right",
      "tOffsetMs": 149,
      "acAsrConf": 216
    }, {
      "utf8": " let's",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " see",
      "tOffsetMs": 750,
      "acAsrConf": 220
    }, {
      "utf8": " we",
      "tOffsetMs": 989,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 1199,
      "acAsrConf": 252
    }, {
      "utf8": " how",
      "tOffsetMs": 2300,
      "acAsrConf": 234
    }, {
      "utf8": " do",
      "tOffsetMs": 3300,
      "acAsrConf": 252
    }, {
      "utf8": " i",
      "tOffsetMs": 3359,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2244910,
    "dDurationMs": 1839,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2244920,
    "dDurationMs": 3329,
    "wWinId": 1,
    "segs": [ {
      "utf8": "improve",
      "acAsrConf": 249
    }, {
      "utf8": " my",
      "tOffsetMs": 270,
      "acAsrConf": 252
    }, {
      "utf8": " speaking",
      "tOffsetMs": 329,
      "acAsrConf": 208
    }, {
      "utf8": " oh",
      "tOffsetMs": 689,
      "acAsrConf": 135
    }, {
      "utf8": " I",
      "tOffsetMs": 1020,
      "acAsrConf": 238
    }, {
      "utf8": " hope",
      "tOffsetMs": 1109,
      "acAsrConf": 249
    }, {
      "utf8": " that",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1589,
      "acAsrConf": 207
    } ]
  }, {
    "tStartMs": 2246739,
    "dDurationMs": 1510,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2246749,
    "dDurationMs": 2131,
    "wWinId": 1,
    "segs": [ {
      "utf8": "learned",
      "acAsrConf": 213
    }, {
      "utf8": " some",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " tips",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 661,
      "acAsrConf": 223
    }, {
      "utf8": " this",
      "tOffsetMs": 901,
      "acAsrConf": 252
    }, {
      "utf8": " presentation",
      "tOffsetMs": 1020,
      "acAsrConf": 240
    } ]
  }, {
    "tStartMs": 2248239,
    "dDurationMs": 641,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2248249,
    "dDurationMs": 2671,
    "wWinId": 1,
    "segs": [ {
      "utf8": "already",
      "acAsrConf": 247
    } ]
  }, {
    "tStartMs": 2248870,
    "dDurationMs": 2050,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2248880,
    "dDurationMs": 3600,
    "wWinId": 1,
    "segs": [ {
      "utf8": "Oh",
      "acAsrConf": 152
    } ]
  }, {
    "tStartMs": 2250910,
    "dDurationMs": 1570,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2250920,
    "dDurationMs": 6480,
    "wWinId": 1,
    "segs": [ {
      "utf8": "all",
      "acAsrConf": 217
    }, {
      "utf8": " right",
      "tOffsetMs": 149,
      "acAsrConf": 217
    }, {
      "utf8": " we",
      "tOffsetMs": 330,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 510,
      "acAsrConf": 88
    }, {
      "utf8": " good",
      "tOffsetMs": 629,
      "acAsrConf": 223
    }, {
      "utf8": " question",
      "tOffsetMs": 929,
      "acAsrConf": 205
    }, {
      "utf8": " here",
      "tOffsetMs": 1379,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2252470,
    "dDurationMs": 4930,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2252480,
    "dDurationMs": 6810,
    "wWinId": 1,
    "segs": [ {
      "utf8": "from",
      "acAsrConf": 215
    }, {
      "utf8": " less",
      "tOffsetMs": 560,
      "acAsrConf": 212
    }, {
      "utf8": " less",
      "tOffsetMs": 1700,
      "acAsrConf": 238
    }, {
      "utf8": " sick",
      "tOffsetMs": 2700,
      "acAsrConf": 255
    }, {
      "utf8": " let",
      "tOffsetMs": 3240,
      "acAsrConf": 202
    }, {
      "utf8": " sick",
      "tOffsetMs": 3839,
      "acAsrConf": 239
    }, {
      "utf8": " mystic",
      "tOffsetMs": 4050,
      "acAsrConf": 170
    }, {
      "utf8": " I",
      "tOffsetMs": 4740,
      "acAsrConf": 209
    } ]
  }, {
    "tStartMs": 2257390,
    "dDurationMs": 1900,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2257400,
    "dDurationMs": 5790,
    "wWinId": 1,
    "segs": [ {
      "utf8": "hope",
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " said",
      "tOffsetMs": 390,
      "acAsrConf": 252
    }, {
      "utf8": " your",
      "tOffsetMs": 449,
      "acAsrConf": 252
    }, {
      "utf8": " name",
      "tOffsetMs": 630,
      "acAsrConf": 252
    }, {
      "utf8": " correctly",
      "tOffsetMs": 810,
      "acAsrConf": 217
    }, {
      "utf8": " house",
      "tOffsetMs": 1380,
      "acAsrConf": 130
    } ]
  }, {
    "tStartMs": 2259280,
    "dDurationMs": 3910,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2259290,
    "dDurationMs": 8549,
    "wWinId": 1,
    "segs": [ {
      "utf8": "key",
      "acAsrConf": 164
    }, {
      "utf8": " quite",
      "tOffsetMs": 920,
      "acAsrConf": 214
    }, {
      "utf8": " cows",
      "tOffsetMs": 1920,
      "acAsrConf": 122
    }, {
      "utf8": " key",
      "tOffsetMs": 2490,
      "acAsrConf": 124
    }, {
      "utf8": " sounds",
      "tOffsetMs": 2819,
      "acAsrConf": 252
    }, {
      "utf8": " like",
      "tOffsetMs": 3180,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 3360,
      "acAsrConf": 217
    }, {
      "utf8": " coke",
      "tOffsetMs": 3539,
      "acAsrConf": 214
    } ]
  }, {
    "tStartMs": 2263180,
    "dDurationMs": 4659,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2263190,
    "dDurationMs": 7440,
    "wWinId": 1,
    "segs": [ {
      "utf8": "Red",
      "acAsrConf": 171
    }, {
      "utf8": " Cross",
      "tOffsetMs": 540,
      "acAsrConf": 171
    }, {
      "utf8": " key",
      "tOffsetMs": 960,
      "acAsrConf": 157
    }, {
      "utf8": " good",
      "tOffsetMs": 1550,
      "acAsrConf": 245
    }, {
      "utf8": " try",
      "tOffsetMs": 2550,
      "acAsrConf": 251
    }, {
      "utf8": " let's",
      "tOffsetMs": 3320,
      "acAsrConf": 187
    }, {
      "utf8": " see",
      "tOffsetMs": 4320,
      "acAsrConf": 252
    }, {
      "utf8": " she",
      "tOffsetMs": 4440,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2267829,
    "dDurationMs": 2801,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2267839,
    "dDurationMs": 4141,
    "wWinId": 1,
    "segs": [ {
      "utf8": "says",
      "acAsrConf": 248
    }, {
      "utf8": " he",
      "tOffsetMs": 301,
      "acAsrConf": 252
    }, {
      "utf8": " says",
      "tOffsetMs": 691,
      "acAsrConf": 252
    }, {
      "utf8": " hmm",
      "tOffsetMs": 990,
      "acAsrConf": 201
    }, {
      "utf8": " I'm",
      "tOffsetMs": 1651,
      "acAsrConf": 194
    }, {
      "utf8": " not",
      "tOffsetMs": 1770,
      "acAsrConf": 252
    }, {
      "utf8": " sure",
      "tOffsetMs": 1831,
      "acAsrConf": 148
    }, {
      "utf8": " I",
      "tOffsetMs": 1980,
      "acAsrConf": 146
    }, {
      "utf8": " think",
      "tOffsetMs": 2311,
      "acAsrConf": 252
    }, {
      "utf8": " he",
      "tOffsetMs": 2371,
      "acAsrConf": 168
    } ]
  }, {
    "tStartMs": 2270620,
    "dDurationMs": 1360,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2270630,
    "dDurationMs": 3719,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 246
    }, {
      "utf8": " think",
      "tOffsetMs": 390,
      "acAsrConf": 252
    }, {
      "utf8": " he",
      "tOffsetMs": 419,
      "acAsrConf": 187
    }, {
      "utf8": " alright",
      "tOffsetMs": 719,
      "acAsrConf": 121
    } ]
  }, {
    "tStartMs": 2271970,
    "dDurationMs": 2379,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2271980,
    "dDurationMs": 4020,
    "wWinId": 1,
    "segs": [ {
      "utf8": "this",
      "acAsrConf": 251
    }, {
      "utf8": " wonderful",
      "tOffsetMs": 450,
      "acAsrConf": 255
    }, {
      "utf8": " English",
      "tOffsetMs": 1079,
      "acAsrConf": 160
    }, {
      "utf8": " Learner",
      "tOffsetMs": 1530,
      "acAsrConf": 233
    }, {
      "utf8": " says",
      "tOffsetMs": 1770,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 2160,
      "acAsrConf": 213
    } ]
  }, {
    "tStartMs": 2274339,
    "dDurationMs": 1661,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2274349,
    "dDurationMs": 3871,
    "wWinId": 1,
    "segs": [ {
      "utf8": "started",
      "acAsrConf": 252
    }, {
      "utf8": " learning",
      "tOffsetMs": 601,
      "acAsrConf": 252
    }, {
      "utf8": " whole",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " phrases",
      "tOffsetMs": 1081,
      "acAsrConf": 202
    }, {
      "utf8": " with",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2275990,
    "dDurationMs": 2230,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2276000,
    "dDurationMs": 4829,
    "wWinId": 1,
    "segs": [ {
      "utf8": "particular",
      "acAsrConf": 252
    }, {
      "utf8": " words",
      "tOffsetMs": 569,
      "acAsrConf": 252
    }, {
      "utf8": " I'd",
      "tOffsetMs": 869,
      "acAsrConf": 163
    }, {
      "utf8": " like",
      "tOffsetMs": 1170,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1230,
      "acAsrConf": 87
    }, {
      "utf8": " remember",
      "tOffsetMs": 1440,
      "acAsrConf": 167
    }, {
      "utf8": " I",
      "tOffsetMs": 1920,
      "acAsrConf": 159
    } ]
  }, {
    "tStartMs": 2278210,
    "dDurationMs": 2619,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2278220,
    "dDurationMs": 5099,
    "wWinId": 1,
    "segs": [ {
      "utf8": "use",
      "acAsrConf": 252
    }, {
      "utf8": " flashcards",
      "tOffsetMs": 180,
      "acAsrConf": 240
    }, {
      "utf8": " to",
      "tOffsetMs": 480,
      "acAsrConf": 198
    }, {
      "utf8": " remember",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    }, {
      "utf8": " them",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    }, {
      "utf8": " what",
      "tOffsetMs": 1680,
      "acAsrConf": 236
    }, {
      "utf8": " are",
      "tOffsetMs": 2399,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2280819,
    "dDurationMs": 2500,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2280829,
    "dDurationMs": 5101,
    "wWinId": 1,
    "segs": [ {
      "utf8": "your",
      "acAsrConf": 252
    }, {
      "utf8": " methods",
      "tOffsetMs": 211,
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 540,
      "acAsrConf": 228
    }, {
      "utf8": " this",
      "tOffsetMs": 1081,
      "acAsrConf": 227
    }, {
      "utf8": " so",
      "tOffsetMs": 1260,
      "acAsrConf": 197
    }, {
      "utf8": " what",
      "tOffsetMs": 2010,
      "acAsrConf": 236
    }, {
      "utf8": " are",
      "tOffsetMs": 2250,
      "acAsrConf": 252
    }, {
      "utf8": " your",
      "tOffsetMs": 2371,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2283309,
    "dDurationMs": 2621,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2283319,
    "dDurationMs": 4891,
    "wWinId": 1,
    "segs": [ {
      "utf8": "methods",
      "acAsrConf": 145
    }, {
      "utf8": " for",
      "tOffsetMs": 331,
      "acAsrConf": 252
    }, {
      "utf8": " learning",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " particular",
      "tOffsetMs": 1280,
      "acAsrConf": 219
    }, {
      "utf8": " words",
      "tOffsetMs": 2280,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2285920,
    "dDurationMs": 2290,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2285930,
    "dDurationMs": 4080,
    "wWinId": 1,
    "segs": [ {
      "utf8": "and",
      "acAsrConf": 214
    }, {
      "utf8": " learning",
      "tOffsetMs": 510,
      "acAsrConf": 250
    }, {
      "utf8": " you",
      "tOffsetMs": 1110,
      "acAsrConf": 244
    }, {
      "utf8": " said",
      "tOffsetMs": 1530,
      "acAsrConf": 252
    }, {
      "utf8": " you're",
      "tOffsetMs": 1800,
      "acAsrConf": 243
    }, {
      "utf8": " learning",
      "tOffsetMs": 1980,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2288200,
    "dDurationMs": 1810,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2288210,
    "dDurationMs": 3869,
    "wWinId": 1,
    "segs": [ {
      "utf8": "them",
      "acAsrConf": 251
    }, {
      "utf8": " with",
      "tOffsetMs": 119,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 330,
      "acAsrConf": 228
    }, {
      "utf8": " whole",
      "tOffsetMs": 510,
      "acAsrConf": 249
    }, {
      "utf8": " phrase",
      "tOffsetMs": 810,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 1290,
      "acAsrConf": 249
    }, {
      "utf8": " whole",
      "tOffsetMs": 1649,
      "acAsrConf": 185
    } ]
  }, {
    "tStartMs": 2290000,
    "dDurationMs": 2079,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2290010,
    "dDurationMs": 5490,
    "wWinId": 1,
    "segs": [ {
      "utf8": "phrase",
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 210,
      "acAsrConf": 215
    }, {
      "utf8": " I",
      "tOffsetMs": 839,
      "acAsrConf": 248
    }, {
      "utf8": " think",
      "tOffsetMs": 1140,
      "acAsrConf": 135
    }, {
      "utf8": " you're",
      "tOffsetMs": 1470,
      "acAsrConf": 251
    }, {
      "utf8": " on",
      "tOffsetMs": 1620,
      "acAsrConf": 207
    }, {
      "utf8": " the",
      "tOffsetMs": 1740,
      "acAsrConf": 252
    }, {
      "utf8": " right",
      "tOffsetMs": 1770,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 2292069,
    "dDurationMs": 3431,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2292079,
    "dDurationMs": 6091,
    "wWinId": 1,
    "segs": [ {
      "utf8": "path",
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 411,
      "acAsrConf": 249
    }, {
      "utf8": " way",
      "tOffsetMs": 1411,
      "acAsrConf": 167
    }, {
      "utf8": " not",
      "tOffsetMs": 1621,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 2101,
      "acAsrConf": 205
    }, {
      "utf8": " learn",
      "tOffsetMs": 2371,
      "acAsrConf": 226
    }, {
      "utf8": " words",
      "tOffsetMs": 2581,
      "acAsrConf": 236
    }, {
      "utf8": " is",
      "tOffsetMs": 2970,
      "acAsrConf": 225
    }, {
      "utf8": " just",
      "tOffsetMs": 3361,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2295490,
    "dDurationMs": 2680,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2295500,
    "dDurationMs": 5280,
    "wWinId": 1,
    "segs": [ {
      "utf8": "to",
      "acAsrConf": 155
    }, {
      "utf8": " memorize",
      "tOffsetMs": 540,
      "acAsrConf": 153
    }, {
      "utf8": " the",
      "tOffsetMs": 1049,
      "acAsrConf": 252
    }, {
      "utf8": " word",
      "tOffsetMs": 1349,
      "acAsrConf": 202
    }, {
      "utf8": " there's",
      "tOffsetMs": 1530,
      "acAsrConf": 218
    }, {
      "utf8": " no",
      "tOffsetMs": 1829,
      "acAsrConf": 252
    }, {
      "utf8": " context",
      "tOffsetMs": 2010,
      "acAsrConf": 230
    } ]
  }, {
    "tStartMs": 2298160,
    "dDurationMs": 2620,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2298170,
    "dDurationMs": 5250,
    "wWinId": 1,
    "segs": [ {
      "utf8": "there's",
      "acAsrConf": 252
    }, {
      "utf8": " no",
      "tOffsetMs": 179,
      "acAsrConf": 252
    }, {
      "utf8": " idea",
      "tOffsetMs": 330,
      "acAsrConf": 252
    }, {
      "utf8": " how",
      "tOffsetMs": 659,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1320,
      "acAsrConf": 252
    }, {
      "utf8": " put",
      "tOffsetMs": 1379,
      "acAsrConf": 252
    }, {
      "utf8": " them",
      "tOffsetMs": 2100,
      "acAsrConf": 226
    }, {
      "utf8": " into",
      "tOffsetMs": 2310,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 2490,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2300770,
    "dDurationMs": 2650,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2300780,
    "dDurationMs": 4559,
    "wWinId": 1,
    "segs": [ {
      "utf8": "sentence",
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 450,
      "acAsrConf": 226
    }, {
      "utf8": " you're",
      "tOffsetMs": 720,
      "acAsrConf": 240
    }, {
      "utf8": " on",
      "tOffsetMs": 1440,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 1710,
      "acAsrConf": 252
    }, {
      "utf8": " right",
      "tOffsetMs": 1950,
      "acAsrConf": 252
    }, {
      "utf8": " path",
      "tOffsetMs": 2100,
      "acAsrConf": 216
    }, {
      "utf8": " I",
      "tOffsetMs": 2339,
      "acAsrConf": 200
    } ]
  }, {
    "tStartMs": 2303410,
    "dDurationMs": 1929,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2303420,
    "dDurationMs": 5310,
    "wWinId": 1,
    "segs": [ {
      "utf8": "think",
      "acAsrConf": 233
    }, {
      "utf8": " it's",
      "tOffsetMs": 300,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 419,
      "acAsrConf": 252
    }, {
      "utf8": " very",
      "tOffsetMs": 510,
      "acAsrConf": 220
    }, {
      "utf8": " good",
      "tOffsetMs": 689,
      "acAsrConf": 223
    }, {
      "utf8": " thing",
      "tOffsetMs": 929,
      "acAsrConf": 227
    }, {
      "utf8": " to",
      "tOffsetMs": 990,
      "acAsrConf": 62
    }, {
      "utf8": " do",
      "tOffsetMs": 1439,
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 1590,
      "acAsrConf": 180
    } ]
  }, {
    "tStartMs": 2305329,
    "dDurationMs": 3401,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2305339,
    "dDurationMs": 5490,
    "wWinId": 1,
    "segs": [ {
      "utf8": "beginners",
      "acAsrConf": 252
    }, {
      "utf8": " especially",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " flashcards",
      "tOffsetMs": 1671,
      "acAsrConf": 206
    }, {
      "utf8": " you",
      "tOffsetMs": 2671,
      "acAsrConf": 167
    }, {
      "utf8": " know",
      "tOffsetMs": 3270,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2308720,
    "dDurationMs": 2109,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2308730,
    "dDurationMs": 5220,
    "wWinId": 1,
    "segs": [ {
      "utf8": "just",
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 300,
      "acAsrConf": 252
    }, {
      "utf8": " try",
      "tOffsetMs": 420,
      "acAsrConf": 204
    }, {
      "utf8": " to",
      "tOffsetMs": 599,
      "acAsrConf": 204
    }, {
      "utf8": " remember",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " very",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " basic",
      "tOffsetMs": 1319,
      "acAsrConf": 252
    }, {
      "utf8": " or",
      "tOffsetMs": 1859,
      "acAsrConf": 248
    } ]
  }, {
    "tStartMs": 2310819,
    "dDurationMs": 3131,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2310829,
    "dDurationMs": 4591,
    "wWinId": 1,
    "segs": [ {
      "utf8": "very",
      "acAsrConf": 237
    }, {
      "utf8": " important",
      "tOffsetMs": 500,
      "acAsrConf": 252
    }, {
      "utf8": " phrases",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    }, {
      "utf8": " that's",
      "tOffsetMs": 1831,
      "acAsrConf": 140
    }, {
      "utf8": " very",
      "tOffsetMs": 2551,
      "acAsrConf": 226
    }, {
      "utf8": " good",
      "tOffsetMs": 2911,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2313940,
    "dDurationMs": 1480,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2313950,
    "dDurationMs": 4050,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 252
    }, {
      "utf8": " were",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " going",
      "tOffsetMs": 359,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 570,
      "acAsrConf": 252
    }, {
      "utf8": " use",
      "tOffsetMs": 629,
      "acAsrConf": 252
    }, {
      "utf8": " important",
      "tOffsetMs": 780,
      "acAsrConf": 245
    }, {
      "utf8": " phrases",
      "tOffsetMs": 1230,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2315410,
    "dDurationMs": 2590,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2315420,
    "dDurationMs": 5159,
    "wWinId": 1,
    "segs": [ {
      "utf8": "all",
      "acAsrConf": 208
    }, {
      "utf8": " the",
      "tOffsetMs": 270,
      "acAsrConf": 252
    }, {
      "utf8": " time",
      "tOffsetMs": 389,
      "acAsrConf": 252
    }, {
      "utf8": " but",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " again",
      "tOffsetMs": 870,
      "acAsrConf": 248
    }, {
      "utf8": " for",
      "tOffsetMs": 1590,
      "acAsrConf": 204
    }, {
      "utf8": " fluently",
      "tOffsetMs": 2040,
      "acAsrConf": 236
    } ]
  }, {
    "tStartMs": 2317990,
    "dDurationMs": 2589,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2318000,
    "dDurationMs": 5579,
    "wWinId": 1,
    "segs": [ {
      "utf8": "speaking",
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 750,
      "acAsrConf": 236
    }, {
      "utf8": " you",
      "tOffsetMs": 1109,
      "acAsrConf": 249
    }, {
      "utf8": " have",
      "tOffsetMs": 1589,
      "acAsrConf": 233
    }, {
      "utf8": " to",
      "tOffsetMs": 1680,
      "acAsrConf": 252
    }, {
      "utf8": " really",
      "tOffsetMs": 1859,
      "acAsrConf": 237
    }, {
      "utf8": " practice",
      "tOffsetMs": 2220,
      "acAsrConf": 226
    } ]
  }, {
    "tStartMs": 2320569,
    "dDurationMs": 3010,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2320579,
    "dDurationMs": 5071,
    "wWinId": 1,
    "segs": [ {
      "utf8": "and",
      "acAsrConf": 187
    }, {
      "utf8": " do",
      "tOffsetMs": 571,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 1381,
      "acAsrConf": 240
    }, {
      "utf8": " sure",
      "tOffsetMs": 1591,
      "acAsrConf": 217
    }, {
      "utf8": " maybe",
      "tOffsetMs": 1951,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 2191,
      "acAsrConf": 235
    }, {
      "utf8": " can",
      "tOffsetMs": 2311,
      "acAsrConf": 199
    }, {
      "utf8": " read",
      "tOffsetMs": 2581,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 2821,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2323569,
    "dDurationMs": 2081,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2323579,
    "dDurationMs": 4681,
    "wWinId": 1,
    "segs": [ {
      "utf8": "flashcards",
      "acAsrConf": 250
    }, {
      "utf8": " out",
      "tOffsetMs": 571,
      "acAsrConf": 175
    }, {
      "utf8": " loud",
      "tOffsetMs": 720,
      "acAsrConf": 231
    }, {
      "utf8": " using",
      "tOffsetMs": 961,
      "acAsrConf": 232
    }, {
      "utf8": " those",
      "tOffsetMs": 1351,
      "acAsrConf": 245
    }, {
      "utf8": " steps",
      "tOffsetMs": 1530,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 1861,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2325640,
    "dDurationMs": 2620,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2325650,
    "dDurationMs": 4290,
    "wWinId": 1,
    "segs": [ {
      "utf8": "listening",
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 449,
      "acAsrConf": 202
    }, {
      "utf8": " reading",
      "tOffsetMs": 480,
      "acAsrConf": 207
    }, {
      "utf8": " speaking",
      "tOffsetMs": 719,
      "acAsrConf": 217
    }, {
      "utf8": " so",
      "tOffsetMs": 1230,
      "acAsrConf": 223
    }, {
      "utf8": " using",
      "tOffsetMs": 1980,
      "acAsrConf": 237
    } ]
  }, {
    "tStartMs": 2328250,
    "dDurationMs": 1690,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2328260,
    "dDurationMs": 4260,
    "wWinId": 1,
    "segs": [ {
      "utf8": "the",
      "acAsrConf": 252
    }, {
      "utf8": " flashcards",
      "tOffsetMs": 89,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 569,
      "acAsrConf": 206
    }, {
      "utf8": " be",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " able",
      "tOffsetMs": 809,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 930,
      "acAsrConf": 187
    }, {
      "utf8": " read",
      "tOffsetMs": 1200,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1410,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2329930,
    "dDurationMs": 2590,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2329940,
    "dDurationMs": 5520,
    "wWinId": 1,
    "segs": [ {
      "utf8": "learning",
      "acAsrConf": 246
    }, {
      "utf8": " the",
      "tOffsetMs": 389,
      "acAsrConf": 210
    }, {
      "utf8": " vocabulary",
      "tOffsetMs": 540,
      "acAsrConf": 208
    }, {
      "utf8": " words",
      "tOffsetMs": 1230,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 1980,
      "acAsrConf": 237
    }, {
      "utf8": " whole",
      "tOffsetMs": 2190,
      "acAsrConf": 214
    } ]
  }, {
    "tStartMs": 2332510,
    "dDurationMs": 2950,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2332520,
    "dDurationMs": 5130,
    "wWinId": 1,
    "segs": [ {
      "utf8": "phrases",
      "acAsrConf": 207
    }, {
      "utf8": " sentences",
      "tOffsetMs": 539,
      "acAsrConf": 252
    }, {
      "utf8": " even",
      "tOffsetMs": 750,
      "acAsrConf": 208
    }, {
      "utf8": " a",
      "tOffsetMs": 1380,
      "acAsrConf": 149
    }, {
      "utf8": " story",
      "tOffsetMs": 1740,
      "acAsrConf": 252
    }, {
      "utf8": " that's",
      "tOffsetMs": 2220,
      "acAsrConf": 214
    } ]
  }, {
    "tStartMs": 2335450,
    "dDurationMs": 2200,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2335460,
    "dDurationMs": 4230,
    "wWinId": 1,
    "segs": [ {
      "utf8": "better",
      "acAsrConf": 242
    }, {
      "utf8": " the",
      "tOffsetMs": 330,
      "acAsrConf": 252
    }, {
      "utf8": " more",
      "tOffsetMs": 599,
      "acAsrConf": 252
    }, {
      "utf8": " context",
      "tOffsetMs": 839,
      "acAsrConf": 242
    }, {
      "utf8": " you",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 1649,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 1680,
      "acAsrConf": 182
    } ]
  }, {
    "tStartMs": 2337640,
    "dDurationMs": 2050,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2337650,
    "dDurationMs": 4860,
    "wWinId": 1,
    "segs": [ {
      "utf8": "better",
      "acAsrConf": 226
    }, {
      "utf8": " all",
      "tOffsetMs": 179,
      "acAsrConf": 188
    }, {
      "utf8": " right",
      "tOffsetMs": 1020,
      "acAsrConf": 235
    }, {
      "utf8": " let's",
      "tOffsetMs": 1140,
      "acAsrConf": 252
    }, {
      "utf8": " see",
      "tOffsetMs": 1320,
      "acAsrConf": 252
    }, {
      "utf8": " what",
      "tOffsetMs": 1469,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 1500,
      "acAsrConf": 171
    } ]
  }, {
    "tStartMs": 2339680,
    "dDurationMs": 2830,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2339690,
    "dDurationMs": 5820,
    "wWinId": 1,
    "segs": [ {
      "utf8": "another",
      "acAsrConf": 247
    }, {
      "utf8": " question",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " right",
      "tOffsetMs": 1590,
      "acAsrConf": 163
    }, {
      "utf8": " can",
      "tOffsetMs": 2040,
      "acAsrConf": 230
    }, {
      "utf8": " you",
      "tOffsetMs": 2280,
      "acAsrConf": 176
    }, {
      "utf8": " continue",
      "tOffsetMs": 2399,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2342500,
    "dDurationMs": 3010,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2342510,
    "dDurationMs": 5849,
    "wWinId": 1,
    "segs": [ {
      "utf8": "let's",
      "acAsrConf": 229
    }, {
      "utf8": " see",
      "tOffsetMs": 599,
      "acAsrConf": 242
    }, {
      "utf8": " I",
      "tOffsetMs": 809,
      "acAsrConf": 243
    }, {
      "utf8": " like",
      "tOffsetMs": 1020,
      "acAsrConf": 251
    }, {
      "utf8": " this",
      "tOffsetMs": 1260,
      "acAsrConf": 252
    }, {
      "utf8": " one",
      "tOffsetMs": 1440,
      "acAsrConf": 243
    }, {
      "utf8": " though",
      "tOffsetMs": 1650,
      "acAsrConf": 224
    }, {
      "utf8": " speak",
      "tOffsetMs": 2000,
      "acAsrConf": 240
    } ]
  }, {
    "tStartMs": 2345500,
    "dDurationMs": 2859,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2345510,
    "dDurationMs": 4740,
    "wWinId": 1,
    "segs": [ {
      "utf8": "this",
      "acAsrConf": 204
    }, {
      "utf8": " is",
      "tOffsetMs": 539,
      "acAsrConf": 208
    }, {
      "utf8": " my",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " afraid",
      "tOffsetMs": 900,
      "acAsrConf": 222
    }, {
      "utf8": " oh",
      "tOffsetMs": 1440,
      "acAsrConf": 81
    }, {
      "utf8": " alright",
      "tOffsetMs": 1740,
      "acAsrConf": 80
    }, {
      "utf8": " listen",
      "tOffsetMs": 2280,
      "acAsrConf": 205
    }, {
      "utf8": " I",
      "tOffsetMs": 2670,
      "acAsrConf": 236
    } ]
  }, {
    "tStartMs": 2348349,
    "dDurationMs": 1901,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2348359,
    "dDurationMs": 2851,
    "wWinId": 1,
    "segs": [ {
      "utf8": "think",
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 61,
      "acAsrConf": 208
    }, {
      "utf8": " mean",
      "tOffsetMs": 331,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 601,
      "acAsrConf": 202
    }, {
      "utf8": " are",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " scared",
      "tOffsetMs": 1111,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 1651,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2350240,
    "dDurationMs": 970,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2350250,
    "dDurationMs": 3829,
    "wWinId": 1,
    "segs": [ {
      "utf8": "speaking",
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2351200,
    "dDurationMs": 2879,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2351210,
    "dDurationMs": 8010,
    "wWinId": 1,
    "segs": [ {
      "utf8": "alright",
      "acAsrConf": 66
    }, {
      "utf8": " we",
      "tOffsetMs": 359,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 540,
      "acAsrConf": 111
    }, {
      "utf8": " good",
      "tOffsetMs": 720,
      "acAsrConf": 239
    }, {
      "utf8": " question",
      "tOffsetMs": 1080,
      "acAsrConf": 208
    }, {
      "utf8": " here",
      "tOffsetMs": 1710,
      "acAsrConf": 236
    } ]
  }, {
    "tStartMs": 2354069,
    "dDurationMs": 5151,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2354079,
    "dDurationMs": 7931,
    "wWinId": 1,
    "segs": [ {
      "utf8": "Marcelo",
      "acAsrConf": 208
    }, {
      "utf8": " Barbosa",
      "tOffsetMs": 1020,
      "acAsrConf": 255
    }, {
      "utf8": " so",
      "tOffsetMs": 2020,
      "acAsrConf": 180
    }, {
      "utf8": " are",
      "tOffsetMs": 2891,
      "acAsrConf": 201
    }, {
      "utf8": " eight",
      "tOffsetMs": 3341,
      "acAsrConf": 69
    }, {
      "utf8": " boys",
      "tOffsetMs": 3881,
      "acAsrConf": 87
    } ]
  }, {
    "tStartMs": 2359210,
    "dDurationMs": 2800,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2359220,
    "dDurationMs": 5609,
    "wWinId": 1,
    "segs": [ {
      "utf8": "alright",
      "acAsrConf": 90
    }, {
      "utf8": " he",
      "tOffsetMs": 389,
      "acAsrConf": 233
    }, {
      "utf8": " says",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " speak",
      "tOffsetMs": 899,
      "acAsrConf": 233
    }, {
      "utf8": " this",
      "tOffsetMs": 1589,
      "acAsrConf": 224
    }, {
      "utf8": " is",
      "tOffsetMs": 2070,
      "acAsrConf": 207
    }, {
      "utf8": " my",
      "tOffsetMs": 2250,
      "acAsrConf": 252
    }, {
      "utf8": " afraid",
      "tOffsetMs": 2399,
      "acAsrConf": 231
    } ]
  }, {
    "tStartMs": 2362000,
    "dDurationMs": 2829,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2362010,
    "dDurationMs": 5970,
    "wWinId": 1,
    "segs": [ {
      "utf8": "I",
      "acAsrConf": 215
    }, {
      "utf8": " think",
      "tOffsetMs": 299,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " mean",
      "tOffsetMs": 599,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 720,
      "acAsrConf": 223
    }, {
      "utf8": " say",
      "tOffsetMs": 900,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 1140,
      "acAsrConf": 251
    }, {
      "utf8": " am",
      "tOffsetMs": 1470,
      "acAsrConf": 248
    }, {
      "utf8": " afraid",
      "tOffsetMs": 2309,
      "acAsrConf": 205
    }, {
      "utf8": " of",
      "tOffsetMs": 2789,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2364819,
    "dDurationMs": 3161,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2364829,
    "dDurationMs": 5671,
    "wWinId": 1,
    "segs": [ {
      "utf8": "speaking",
      "acAsrConf": 200
    }, {
      "utf8": " more",
      "tOffsetMs": 540,
      "acAsrConf": 145
    }, {
      "utf8": " I'm",
      "tOffsetMs": 871,
      "acAsrConf": 248
    }, {
      "utf8": " scared",
      "tOffsetMs": 1141,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 1921,
      "acAsrConf": 216
    }, {
      "utf8": " speaking",
      "tOffsetMs": 2341,
      "acAsrConf": 233
    }, {
      "utf8": " I'm",
      "tOffsetMs": 2970,
      "acAsrConf": 234
    } ]
  }, {
    "tStartMs": 2367970,
    "dDurationMs": 2530,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2367980,
    "dDurationMs": 4530,
    "wWinId": 1,
    "segs": [ {
      "utf8": "scared",
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 300,
      "acAsrConf": 252
    }, {
      "utf8": " speaking",
      "tOffsetMs": 389,
      "acAsrConf": 249
    }, {
      "utf8": " yes",
      "tOffsetMs": 839,
      "acAsrConf": 252
    }, {
      "utf8": " yes",
      "tOffsetMs": 1020,
      "acAsrConf": 226
    }, {
      "utf8": " and",
      "tOffsetMs": 1530,
      "acAsrConf": 216
    }, {
      "utf8": " that's",
      "tOffsetMs": 1859,
      "acAsrConf": 217
    } ]
  }, {
    "tStartMs": 2370490,
    "dDurationMs": 2020,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2370500,
    "dDurationMs": 3839,
    "wWinId": 1,
    "segs": [ {
      "utf8": "very",
      "acAsrConf": 200
    }, {
      "utf8": " difficult",
      "tOffsetMs": 270,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 480,
      "acAsrConf": 126
    }, {
      "utf8": " lot",
      "tOffsetMs": 990,
      "acAsrConf": 248
    }, {
      "utf8": " of",
      "tOffsetMs": 1380,
      "acAsrConf": 217
    }, {
      "utf8": " people",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    }, {
      "utf8": " are",
      "tOffsetMs": 1770,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 1859,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2372500,
    "dDurationMs": 1839,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2372510,
    "dDurationMs": 4620,
    "wWinId": 1,
    "segs": [ {
      "utf8": "the",
      "acAsrConf": 207
    }, {
      "utf8": " same",
      "tOffsetMs": 120,
      "acAsrConf": 229
    }, {
      "utf8": " boat",
      "tOffsetMs": 329,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 539,
      "acAsrConf": 156
    }, {
      "utf8": " don't",
      "tOffsetMs": 780,
      "acAsrConf": 201
    }, {
      "utf8": " worry",
      "tOffsetMs": 1140,
      "acAsrConf": 201
    }, {
      "utf8": " a",
      "tOffsetMs": 1349,
      "acAsrConf": 252
    }, {
      "utf8": " lot",
      "tOffsetMs": 1440,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 1650,
      "acAsrConf": 209
    } ]
  }, {
    "tStartMs": 2374329,
    "dDurationMs": 2801,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2374339,
    "dDurationMs": 5520,
    "wWinId": 1,
    "segs": [ {
      "utf8": "people",
      "acAsrConf": 252
    }, {
      "utf8": " feel",
      "tOffsetMs": 361,
      "acAsrConf": 252
    }, {
      "utf8": " similarly",
      "tOffsetMs": 571,
      "acAsrConf": 201
    }, {
      "utf8": " but",
      "tOffsetMs": 1351,
      "acAsrConf": 239
    }, {
      "utf8": " you",
      "tOffsetMs": 1921,
      "acAsrConf": 238
    }, {
      "utf8": " know",
      "tOffsetMs": 2460,
      "acAsrConf": 252
    }, {
      "utf8": " what",
      "tOffsetMs": 2581,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2377120,
    "dDurationMs": 2739,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2377130,
    "dDurationMs": 5250,
    "wWinId": 1,
    "segs": [ {
      "utf8": "it's",
      "acAsrConf": 226
    }, {
      "utf8": " best",
      "tOffsetMs": 300,
      "acAsrConf": 247
    }, {
      "utf8": " just",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 810,
      "acAsrConf": 241
    }, {
      "utf8": " try",
      "tOffsetMs": 1199,
      "acAsrConf": 252
    }, {
      "utf8": " just",
      "tOffsetMs": 1530,
      "acAsrConf": 167
    }, {
      "utf8": " to",
      "tOffsetMs": 2250,
      "acAsrConf": 252
    }, {
      "utf8": " try",
      "tOffsetMs": 2400,
      "acAsrConf": 246
    }, {
      "utf8": " to",
      "tOffsetMs": 2669,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 2379849,
    "dDurationMs": 2531,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2379859,
    "dDurationMs": 4021,
    "wWinId": 1,
    "segs": [ {
      "utf8": "speak",
      "acAsrConf": 252
    }, {
      "utf8": " maybe",
      "tOffsetMs": 391,
      "acAsrConf": 234
    }, {
      "utf8": " to",
      "tOffsetMs": 720,
      "acAsrConf": 220
    }, {
      "utf8": " yourself",
      "tOffsetMs": 1021,
      "acAsrConf": 248
    }, {
      "utf8": " this",
      "tOffsetMs": 1470,
      "acAsrConf": 203
    }, {
      "utf8": " actors",
      "tOffsetMs": 2101,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 2382370,
    "dDurationMs": 1510,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2382380,
    "dDurationMs": 3689,
    "wWinId": 1,
    "segs": [ {
      "utf8": "reading",
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 239,
      "acAsrConf": 211
    }, {
      "utf8": " book",
      "tOffsetMs": 479,
      "acAsrConf": 221
    }, {
      "utf8": " out",
      "tOffsetMs": 660,
      "acAsrConf": 208
    } ]
  }, {
    "tStartMs": 2383870,
    "dDurationMs": 2199,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2383880,
    "dDurationMs": 6000,
    "wWinId": 1,
    "segs": [ {
      "utf8": "out",
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 90,
      "acAsrConf": 230
    }, {
      "utf8": " just",
      "tOffsetMs": 330,
      "acAsrConf": 252
    }, {
      "utf8": " use",
      "tOffsetMs": 479,
      "acAsrConf": 233
    }, {
      "utf8": " your",
      "tOffsetMs": 1080,
      "acAsrConf": 0
    }, {
      "utf8": " mouth",
      "tOffsetMs": 1379,
      "acAsrConf": 255
    }, {
      "utf8": " in",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    }, {
      "utf8": " that",
      "tOffsetMs": 1800,
      "acAsrConf": 146
    }, {
      "utf8": " way",
      "tOffsetMs": 1979,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2386059,
    "dDurationMs": 3821,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2386069,
    "dDurationMs": 6440,
    "wWinId": 1,
    "segs": [ {
      "utf8": "yes",
      "acAsrConf": 236
    }, {
      "utf8": " maybe",
      "tOffsetMs": 601,
      "acAsrConf": 228
    }, {
      "utf8": " find",
      "tOffsetMs": 841,
      "acAsrConf": 227
    }, {
      "utf8": " a",
      "tOffsetMs": 1290,
      "acAsrConf": 252
    }, {
      "utf8": " friend",
      "tOffsetMs": 1530,
      "acAsrConf": 246
    }, {
      "utf8": " you",
      "tOffsetMs": 2190,
      "acAsrConf": 216
    }, {
      "utf8": " like",
      "tOffsetMs": 2550,
      "acAsrConf": 223
    }, {
      "utf8": " and",
      "tOffsetMs": 2821,
      "acAsrConf": 227
    }, {
      "utf8": " try",
      "tOffsetMs": 3210,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2389870,
    "dDurationMs": 2639,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2389880,
    "dDurationMs": 4830,
    "wWinId": 1,
    "segs": [ {
      "utf8": "to",
      "acAsrConf": 178
    }, {
      "utf8": " speak",
      "tOffsetMs": 120,
      "acAsrConf": 252
    }, {
      "utf8": " English",
      "tOffsetMs": 510,
      "acAsrConf": 238
    }, {
      "utf8": " with",
      "tOffsetMs": 810,
      "acAsrConf": 206
    }, {
      "utf8": " him",
      "tOffsetMs": 1320,
      "acAsrConf": 236
    }, {
      "utf8": " or",
      "tOffsetMs": 1590,
      "acAsrConf": 252
    }, {
      "utf8": " her",
      "tOffsetMs": 1800,
      "acAsrConf": 252
    }, {
      "utf8": " first",
      "tOffsetMs": 1979,
      "acAsrConf": 227
    } ]
  }, {
    "tStartMs": 2392499,
    "dDurationMs": 2211,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2392509,
    "dDurationMs": 4181,
    "wWinId": 1,
    "segs": [ {
      "utf8": "somebody",
      "acAsrConf": 207
    }, {
      "utf8": " you",
      "tOffsetMs": 1000,
      "acAsrConf": 252
    }, {
      "utf8": " feel",
      "tOffsetMs": 1151,
      "acAsrConf": 252
    }, {
      "utf8": " comfortable",
      "tOffsetMs": 1181,
      "acAsrConf": 188
    }, {
      "utf8": " with",
      "tOffsetMs": 1600,
      "acAsrConf": 228
    }, {
      "utf8": " yeah",
      "tOffsetMs": 1901,
      "acAsrConf": 135
    } ]
  }, {
    "tStartMs": 2394700,
    "dDurationMs": 1990,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2394710,
    "dDurationMs": 4379,
    "wWinId": 1,
    "segs": [ {
      "utf8": "I",
      "acAsrConf": 235
    }, {
      "utf8": " think",
      "tOffsetMs": 59,
      "acAsrConf": 247
    }, {
      "utf8": " if",
      "tOffsetMs": 180,
      "acAsrConf": 201
    }, {
      "utf8": " you",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " feel",
      "tOffsetMs": 629,
      "acAsrConf": 243
    }, {
      "utf8": " afraid",
      "tOffsetMs": 869,
      "acAsrConf": 207
    }, {
      "utf8": " to",
      "tOffsetMs": 1409,
      "acAsrConf": 252
    }, {
      "utf8": " speak",
      "tOffsetMs": 1680,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2396680,
    "dDurationMs": 2409,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2396690,
    "dDurationMs": 4950,
    "wWinId": 1,
    "segs": [ {
      "utf8": "follow",
      "acAsrConf": 225
    }, {
      "utf8": " those",
      "tOffsetMs": 389,
      "acAsrConf": 242
    }, {
      "utf8": " three",
      "tOffsetMs": 780,
      "acAsrConf": 202
    }, {
      "utf8": " steps",
      "tOffsetMs": 1109,
      "acAsrConf": 252
    }, {
      "utf8": " first",
      "tOffsetMs": 1319,
      "acAsrConf": 212
    }, {
      "utf8": " listen",
      "tOffsetMs": 1710,
      "acAsrConf": 236
    } ]
  }, {
    "tStartMs": 2399079,
    "dDurationMs": 2561,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2399089,
    "dDurationMs": 4710,
    "wWinId": 1,
    "segs": [ {
      "utf8": "then",
      "acAsrConf": 234
    }, {
      "utf8": " read",
      "tOffsetMs": 631,
      "acAsrConf": 252
    }, {
      "utf8": " out",
      "tOffsetMs": 1081,
      "acAsrConf": 225
    }, {
      "utf8": " loud",
      "tOffsetMs": 1321,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1351,
      "acAsrConf": 206
    }, {
      "utf8": " that",
      "tOffsetMs": 1831,
      "acAsrConf": 161
    }, {
      "utf8": " you",
      "tOffsetMs": 2010,
      "acAsrConf": 252
    }, {
      "utf8": " get",
      "tOffsetMs": 2131,
      "acAsrConf": 252
    }, {
      "utf8": " used",
      "tOffsetMs": 2311,
      "acAsrConf": 208
    } ]
  }, {
    "tStartMs": 2401630,
    "dDurationMs": 2169,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2401640,
    "dDurationMs": 4679,
    "wWinId": 1,
    "segs": [ {
      "utf8": "to",
      "acAsrConf": 204
    }, {
      "utf8": " hearing",
      "tOffsetMs": 179,
      "acAsrConf": 252
    }, {
      "utf8": " your",
      "tOffsetMs": 389,
      "acAsrConf": 226
    }, {
      "utf8": " voice",
      "tOffsetMs": 629,
      "acAsrConf": 218
    }, {
      "utf8": " speaking",
      "tOffsetMs": 899,
      "acAsrConf": 201
    }, {
      "utf8": " and",
      "tOffsetMs": 1439,
      "acAsrConf": 244
    }, {
      "utf8": " then",
      "tOffsetMs": 1619,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2403789,
    "dDurationMs": 2530,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2403799,
    "dDurationMs": 5010,
    "wWinId": 1,
    "segs": [ {
      "utf8": "try",
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 391,
      "acAsrConf": 252
    }, {
      "utf8": " speak",
      "tOffsetMs": 421,
      "acAsrConf": 227
    }, {
      "utf8": " with",
      "tOffsetMs": 841,
      "acAsrConf": 174
    }, {
      "utf8": " someone",
      "tOffsetMs": 1020,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1201,
      "acAsrConf": 218
    }, {
      "utf8": " Marcello",
      "tOffsetMs": 1591,
      "acAsrConf": 156
    } ]
  }, {
    "tStartMs": 2406309,
    "dDurationMs": 2500,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2406319,
    "dDurationMs": 4171,
    "wWinId": 1,
    "segs": [ {
      "utf8": "that's",
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " wonderful",
      "tOffsetMs": 540,
      "acAsrConf": 231
    }, {
      "utf8": " question",
      "tOffsetMs": 1530,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1681,
      "acAsrConf": 217
    }, {
      "utf8": " a",
      "tOffsetMs": 2101,
      "acAsrConf": 252
    }, {
      "utf8": " lot",
      "tOffsetMs": 2161,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 2190,
      "acAsrConf": 122
    } ]
  }, {
    "tStartMs": 2408799,
    "dDurationMs": 1691,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2408809,
    "dDurationMs": 3960,
    "wWinId": 1,
    "segs": [ {
      "utf8": "people",
      "acAsrConf": 252
    }, {
      "utf8": " here",
      "tOffsetMs": 300,
      "acAsrConf": 255
    }, {
      "utf8": " wrote",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " similar",
      "tOffsetMs": 780,
      "acAsrConf": 252
    }, {
      "utf8": " things",
      "tOffsetMs": 1230,
      "acAsrConf": 252
    }, {
      "utf8": " they",
      "tOffsetMs": 1500,
      "acAsrConf": 246
    } ]
  }, {
    "tStartMs": 2410480,
    "dDurationMs": 2289,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2410490,
    "dDurationMs": 3839,
    "wWinId": 1,
    "segs": [ {
      "utf8": "feel",
      "acAsrConf": 252
    }, {
      "utf8": " afraid",
      "tOffsetMs": 180,
      "acAsrConf": 201
    }, {
      "utf8": " to",
      "tOffsetMs": 539,
      "acAsrConf": 252
    }, {
      "utf8": " speak",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 990,
      "acAsrConf": 214
    }, {
      "utf8": " I",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    }, {
      "utf8": " hope",
      "tOffsetMs": 1529,
      "acAsrConf": 240
    }, {
      "utf8": " that",
      "tOffsetMs": 1799,
      "acAsrConf": 239
    }, {
      "utf8": " that",
      "tOffsetMs": 2010,
      "acAsrConf": 217
    } ]
  }, {
    "tStartMs": 2412759,
    "dDurationMs": 1570,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2412769,
    "dDurationMs": 4830,
    "wWinId": 1,
    "segs": [ {
      "utf8": "advice",
      "acAsrConf": 205
    }, {
      "utf8": " is",
      "tOffsetMs": 361,
      "acAsrConf": 252
    }, {
      "utf8": " helpful",
      "tOffsetMs": 570,
      "acAsrConf": 115
    }, {
      "utf8": " for",
      "tOffsetMs": 961,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 990,
      "acAsrConf": 212
    }, {
      "utf8": " lot",
      "tOffsetMs": 1171,
      "acAsrConf": 128
    }, {
      "utf8": " of",
      "tOffsetMs": 1290,
      "acAsrConf": 129
    }, {
      "utf8": " other",
      "tOffsetMs": 1411,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2414319,
    "dDurationMs": 3280,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2414329,
    "dDurationMs": 7381,
    "wWinId": 1,
    "segs": [ {
      "utf8": "people",
      "acAsrConf": 252
    }, {
      "utf8": " let's",
      "tOffsetMs": 891,
      "acAsrConf": 243
    }, {
      "utf8": " see",
      "tOffsetMs": 1891,
      "acAsrConf": 252
    }, {
      "utf8": " what",
      "tOffsetMs": 2131,
      "acAsrConf": 252
    }, {
      "utf8": " some",
      "tOffsetMs": 2430,
      "acAsrConf": 222
    }, {
      "utf8": " other",
      "tOffsetMs": 2671,
      "acAsrConf": 232
    }, {
      "utf8": " people",
      "tOffsetMs": 2790,
      "acAsrConf": 144
    } ]
  }, {
    "tStartMs": 2417589,
    "dDurationMs": 4121,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2417599,
    "dDurationMs": 9720,
    "wWinId": 1,
    "segs": [ {
      "utf8": "have",
      "acAsrConf": 252
    }, {
      "utf8": " said",
      "tOffsetMs": 151,
      "acAsrConf": 214
    }, {
      "utf8": " here",
      "tOffsetMs": 420,
      "acAsrConf": 244
    }, {
      "utf8": " all",
      "tOffsetMs": 1641,
      "acAsrConf": 227
    }, {
      "utf8": " right",
      "tOffsetMs": 2641,
      "acAsrConf": 252
    }, {
      "utf8": " all",
      "tOffsetMs": 3210,
      "acAsrConf": 205
    }, {
      "utf8": " right",
      "tOffsetMs": 3631,
      "acAsrConf": 234
    }, {
      "utf8": " we",
      "tOffsetMs": 3690,
      "acAsrConf": 167
    } ]
  }, {
    "tStartMs": 2421700,
    "dDurationMs": 5619,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2421710,
    "dDurationMs": 7200,
    "wWinId": 1,
    "segs": [ {
      "utf8": "have",
      "acAsrConf": 227
    }, {
      "utf8": " ha",
      "tOffsetMs": 1639,
      "acAsrConf": 210
    }, {
      "utf8": " ha",
      "tOffsetMs": 2639,
      "acAsrConf": 211
    }, {
      "utf8": " very",
      "tOffsetMs": 2879,
      "acAsrConf": 238
    }, {
      "utf8": " general",
      "tOffsetMs": 3540,
      "acAsrConf": 250
    }, {
      "utf8": " yes",
      "tOffsetMs": 4170,
      "acAsrConf": 204
    }, {
      "utf8": " many",
      "tOffsetMs": 5010,
      "acAsrConf": 252
    }, {
      "utf8": " people",
      "tOffsetMs": 5280,
      "acAsrConf": 140
    } ]
  }, {
    "tStartMs": 2427309,
    "dDurationMs": 1601,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2427319,
    "dDurationMs": 3030,
    "wWinId": 1,
    "segs": [ {
      "utf8": "had",
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 121,
      "acAsrConf": 252
    }, {
      "utf8": " similar",
      "tOffsetMs": 151,
      "acAsrConf": 252
    }, {
      "utf8": " question",
      "tOffsetMs": 391,
      "acAsrConf": 146
    }, {
      "utf8": " to",
      "tOffsetMs": 931,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 960,
      "acAsrConf": 201
    }, {
      "utf8": " Marcello",
      "tOffsetMs": 1141,
      "acAsrConf": 218
    }, {
      "utf8": " I",
      "tOffsetMs": 1561,
      "acAsrConf": 167
    } ]
  }, {
    "tStartMs": 2428900,
    "dDurationMs": 1449,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2428910,
    "dDurationMs": 3439,
    "wWinId": 1,
    "segs": [ {
      "utf8": "feel",
      "acAsrConf": 207
    }, {
      "utf8": " worried",
      "tOffsetMs": 330,
      "acAsrConf": 252
    }, {
      "utf8": " about",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " speaking",
      "tOffsetMs": 899,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 1199,
      "acAsrConf": 210
    }, {
      "utf8": " feel",
      "tOffsetMs": 1409,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2430339,
    "dDurationMs": 2010,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2430349,
    "dDurationMs": 4531,
    "wWinId": 1,
    "segs": [ {
      "utf8": "scared",
      "acAsrConf": 131
    }, {
      "utf8": " about",
      "tOffsetMs": 391,
      "acAsrConf": 207
    }, {
      "utf8": " speaking",
      "tOffsetMs": 631,
      "acAsrConf": 252
    }, {
      "utf8": " how",
      "tOffsetMs": 930,
      "acAsrConf": 211
    }, {
      "utf8": " can",
      "tOffsetMs": 1291,
      "acAsrConf": 166
    }, {
      "utf8": " I",
      "tOffsetMs": 1561,
      "acAsrConf": 226
    }, {
      "utf8": " improve",
      "tOffsetMs": 1591,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2432339,
    "dDurationMs": 2541,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2432349,
    "dDurationMs": 6520,
    "wWinId": 1,
    "segs": [ {
      "utf8": "so",
      "acAsrConf": 167
    }, {
      "utf8": " I",
      "tOffsetMs": 1000,
      "acAsrConf": 243
    }, {
      "utf8": " hope",
      "tOffsetMs": 1031,
      "acAsrConf": 234
    }, {
      "utf8": " we",
      "tOffsetMs": 1361,
      "acAsrConf": 226
    }, {
      "utf8": " gave",
      "tOffsetMs": 1480,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 1601,
      "acAsrConf": 205
    }, {
      "utf8": " lot",
      "tOffsetMs": 1660,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 1841,
      "acAsrConf": 147
    }, {
      "utf8": " tips",
      "tOffsetMs": 1901,
      "acAsrConf": 205
    }, {
      "utf8": " already",
      "tOffsetMs": 2170,
      "acAsrConf": 251
    } ]
  }, {
    "tStartMs": 2434870,
    "dDurationMs": 3999,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2434880,
    "dDurationMs": 8459,
    "wWinId": 1,
    "segs": [ {
      "utf8": "about",
      "acAsrConf": 218
    }, {
      "utf8": " that",
      "tOffsetMs": 300,
      "acAsrConf": 252
    }, {
      "utf8": " let's",
      "tOffsetMs": 2209,
      "acAsrConf": 247
    }, {
      "utf8": " see",
      "tOffsetMs": 3209,
      "acAsrConf": 252
    }, {
      "utf8": " this",
      "tOffsetMs": 3419,
      "acAsrConf": 202
    }, {
      "utf8": " is",
      "tOffsetMs": 3719,
      "acAsrConf": 216
    }, {
      "utf8": " like",
      "tOffsetMs": 3840,
      "acAsrConf": 219
    } ]
  }, {
    "tStartMs": 2438859,
    "dDurationMs": 4480,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2438869,
    "dDurationMs": 6900,
    "wWinId": 1,
    "segs": [ {
      "utf8": "grammar",
      "acAsrConf": 245
    }, {
      "utf8": " aha",
      "tOffsetMs": 271,
      "acAsrConf": 208
    }, {
      "utf8": " so",
      "tOffsetMs": 960,
      "acAsrConf": 224
    }, {
      "utf8": " we",
      "tOffsetMs": 1470,
      "acAsrConf": 167
    }, {
      "utf8": " have",
      "tOffsetMs": 1920,
      "acAsrConf": 252
    }, {
      "utf8": " mindo",
      "tOffsetMs": 2450,
      "acAsrConf": 143
    }, {
      "utf8": " Marcellus",
      "tOffsetMs": 3470,
      "acAsrConf": 207
    } ]
  }, {
    "tStartMs": 2443329,
    "dDurationMs": 2440,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2443339,
    "dDurationMs": 5371,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 206
    }, {
      "utf8": " know",
      "tOffsetMs": 450,
      "acAsrConf": 195
    }, {
      "utf8": " I",
      "tOffsetMs": 841,
      "acAsrConf": 202
    }, {
      "utf8": " hope",
      "tOffsetMs": 1141,
      "acAsrConf": 242
    }, {
      "utf8": " I",
      "tOffsetMs": 1321,
      "acAsrConf": 252
    }, {
      "utf8": " said",
      "tOffsetMs": 1470,
      "acAsrConf": 252
    }, {
      "utf8": " that",
      "tOffsetMs": 1530,
      "acAsrConf": 252
    }, {
      "utf8": " right",
      "tOffsetMs": 1680,
      "acAsrConf": 252
    }, {
      "utf8": " do",
      "tOffsetMs": 1801,
      "acAsrConf": 208
    }, {
      "utf8": " you",
      "tOffsetMs": 2341,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2445759,
    "dDurationMs": 2951,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2445769,
    "dDurationMs": 5760,
    "wWinId": 1,
    "segs": [ {
      "utf8": "know",
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 121,
      "acAsrConf": 244
    }, {
      "utf8": " as",
      "tOffsetMs": 631,
      "acAsrConf": 236
    }, {
      "utf8": " you",
      "tOffsetMs": 780,
      "acAsrConf": 251
    }, {
      "utf8": " know",
      "tOffsetMs": 901,
      "acAsrConf": 244
    }, {
      "utf8": " says",
      "tOffsetMs": 1020,
      "acAsrConf": 207
    }, {
      "utf8": " so",
      "tOffsetMs": 1260,
      "acAsrConf": 219
    }, {
      "utf8": " I",
      "tOffsetMs": 1951,
      "acAsrConf": 234
    }, {
      "utf8": " think",
      "tOffsetMs": 2221,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 2371,
      "acAsrConf": 214
    } ]
  }, {
    "tStartMs": 2448700,
    "dDurationMs": 2829,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2448710,
    "dDurationMs": 4770,
    "wWinId": 1,
    "segs": [ {
      "utf8": "will",
      "acAsrConf": 252
    }, {
      "utf8": " not",
      "tOffsetMs": 119,
      "acAsrConf": 148
    }, {
      "utf8": " get",
      "tOffsetMs": 480,
      "acAsrConf": 222
    }, {
      "utf8": " well",
      "tOffsetMs": 839,
      "acAsrConf": 252
    }, {
      "utf8": " structure",
      "tOffsetMs": 1200,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 2069,
      "acAsrConf": 252
    }, {
      "utf8": " sentence",
      "tOffsetMs": 2339,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2451519,
    "dDurationMs": 1961,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2451529,
    "dDurationMs": 4770,
    "wWinId": 1,
    "segs": [ {
      "utf8": "it",
      "acAsrConf": 236
    }, {
      "utf8": " makes",
      "tOffsetMs": 300,
      "acAsrConf": 248
    }, {
      "utf8": " me",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " give",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " up",
      "tOffsetMs": 1020,
      "acAsrConf": 252
    }, {
      "utf8": " talking",
      "tOffsetMs": 1050,
      "acAsrConf": 227
    }, {
      "utf8": " to",
      "tOffsetMs": 1621,
      "acAsrConf": 215
    }, {
      "utf8": " someone",
      "tOffsetMs": 1740,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2453470,
    "dDurationMs": 2829,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2453480,
    "dDurationMs": 4920,
    "wWinId": 1,
    "segs": [ {
      "utf8": "hmm",
      "acAsrConf": 218
    }, {
      "utf8": " so",
      "tOffsetMs": 660,
      "acAsrConf": 239
    }, {
      "utf8": " I",
      "tOffsetMs": 779,
      "acAsrConf": 252
    }, {
      "utf8": " think",
      "tOffsetMs": 809,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1049,
      "acAsrConf": 227
    }, {
      "utf8": " mean",
      "tOffsetMs": 1410,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 1710,
      "acAsrConf": 201
    }, {
      "utf8": " think",
      "tOffsetMs": 2129,
      "acAsrConf": 167
    }, {
      "utf8": " my",
      "tOffsetMs": 2609,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2456289,
    "dDurationMs": 2111,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2456299,
    "dDurationMs": 5520,
    "wWinId": 1,
    "segs": [ {
      "utf8": "sentence",
      "acAsrConf": 252
    }, {
      "utf8": " structure",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " isn't",
      "tOffsetMs": 990,
      "acAsrConf": 237
    }, {
      "utf8": " good",
      "tOffsetMs": 1260,
      "acAsrConf": 165
    }, {
      "utf8": " our",
      "tOffsetMs": 1770,
      "acAsrConf": 227
    } ]
  }, {
    "tStartMs": 2458390,
    "dDurationMs": 3429,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2458400,
    "dDurationMs": 6119,
    "wWinId": 1,
    "segs": [ {
      "utf8": "grammar",
      "acAsrConf": 247
    }, {
      "utf8": " my",
      "tOffsetMs": 330,
      "acAsrConf": 230
    }, {
      "utf8": " grammar",
      "tOffsetMs": 899,
      "acAsrConf": 252
    }, {
      "utf8": " isn't",
      "tOffsetMs": 1260,
      "acAsrConf": 180
    }, {
      "utf8": " good",
      "tOffsetMs": 1889,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 2100,
      "acAsrConf": 251
    }, {
      "utf8": " I",
      "tOffsetMs": 2340,
      "acAsrConf": 236
    }, {
      "utf8": " I",
      "tOffsetMs": 2969,
      "acAsrConf": 0
    } ]
  }, {
    "tStartMs": 2461809,
    "dDurationMs": 2710,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2461819,
    "dDurationMs": 5540,
    "wWinId": 1,
    "segs": [ {
      "utf8": "give",
      "acAsrConf": 252
    }, {
      "utf8": " up",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " speaking",
      "tOffsetMs": 690,
      "acAsrConf": 201
    }, {
      "utf8": " with",
      "tOffsetMs": 1290,
      "acAsrConf": 152
    }, {
      "utf8": " people",
      "tOffsetMs": 1591,
      "acAsrConf": 252
    }, {
      "utf8": " because",
      "tOffsetMs": 1800,
      "acAsrConf": 200
    }, {
      "utf8": " I",
      "tOffsetMs": 2190,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 2464509,
    "dDurationMs": 2850,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2464519,
    "dDurationMs": 6270,
    "wWinId": 1,
    "segs": [ {
      "utf8": "feel",
      "acAsrConf": 213
    }, {
      "utf8": " frustrated",
      "tOffsetMs": 391,
      "acAsrConf": 247
    }, {
      "utf8": " so",
      "tOffsetMs": 961,
      "acAsrConf": 231
    }, {
      "utf8": " basically",
      "tOffsetMs": 1141,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1621,
      "acAsrConf": 252
    }, {
      "utf8": " are",
      "tOffsetMs": 2040,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2467349,
    "dDurationMs": 3440,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2467359,
    "dDurationMs": 4331,
    "wWinId": 1,
    "segs": [ {
      "utf8": "scared",
      "acAsrConf": 228
    }, {
      "utf8": " to",
      "tOffsetMs": 1000,
      "acAsrConf": 145
    }, {
      "utf8": " make",
      "tOffsetMs": 1271,
      "acAsrConf": 252
    }, {
      "utf8": " mistakes",
      "tOffsetMs": 1391,
      "acAsrConf": 201
    }, {
      "utf8": " oh",
      "tOffsetMs": 1930,
      "acAsrConf": 0
    }, {
      "utf8": " that's",
      "tOffsetMs": 2430,
      "acAsrConf": 243
    } ]
  }, {
    "tStartMs": 2470779,
    "dDurationMs": 911,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2470789,
    "dDurationMs": 3121,
    "wWinId": 1,
    "segs": [ {
      "utf8": "interesting",
      "acAsrConf": 208
    } ]
  }, {
    "tStartMs": 2471680,
    "dDurationMs": 2230,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2471690,
    "dDurationMs": 4260,
    "wWinId": 1,
    "segs": [ {
      "utf8": "grammar",
      "acAsrConf": 252
    }, {
      "utf8": " mistakes",
      "tOffsetMs": 329,
      "acAsrConf": 252
    }, {
      "utf8": " sure",
      "tOffsetMs": 869,
      "acAsrConf": 166
    }, {
      "utf8": " giving",
      "tOffsetMs": 1200,
      "acAsrConf": 252
    }, {
      "utf8": " up",
      "tOffsetMs": 1710,
      "acAsrConf": 252
    }, {
      "utf8": " speaking",
      "tOffsetMs": 1829,
      "acAsrConf": 228
    } ]
  }, {
    "tStartMs": 2473900,
    "dDurationMs": 2050,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2473910,
    "dDurationMs": 5220,
    "wWinId": 1,
    "segs": [ {
      "utf8": "with",
      "acAsrConf": 252
    }, {
      "utf8": " someone",
      "tOffsetMs": 149,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 330,
      "acAsrConf": 213
    }, {
      "utf8": " I",
      "tOffsetMs": 689,
      "acAsrConf": 230
    }, {
      "utf8": " really",
      "tOffsetMs": 899,
      "acAsrConf": 252
    }, {
      "utf8": " encourage",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1709,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2475940,
    "dDurationMs": 3190,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2475950,
    "dDurationMs": 5639,
    "wWinId": 1,
    "segs": [ {
      "utf8": "to",
      "acAsrConf": 146
    }, {
      "utf8": " try",
      "tOffsetMs": 569,
      "acAsrConf": 223
    }, {
      "utf8": " to",
      "tOffsetMs": 869,
      "acAsrConf": 203
    }, {
      "utf8": " find",
      "tOffsetMs": 930,
      "acAsrConf": 225
    }, {
      "utf8": " someone",
      "tOffsetMs": 1530,
      "acAsrConf": 227
    }, {
      "utf8": " to",
      "tOffsetMs": 2010,
      "acAsrConf": 252
    }, {
      "utf8": " speak",
      "tOffsetMs": 2069,
      "acAsrConf": 195
    }, {
      "utf8": " with",
      "tOffsetMs": 2430,
      "acAsrConf": 252
    }, {
      "utf8": " who",
      "tOffsetMs": 2639,
      "acAsrConf": 246
    } ]
  }, {
    "tStartMs": 2479120,
    "dDurationMs": 2469,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2479130,
    "dDurationMs": 4979,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 236
    }, {
      "utf8": " won't",
      "tOffsetMs": 479,
      "acAsrConf": 255
    }, {
      "utf8": " feel",
      "tOffsetMs": 810,
      "acAsrConf": 252
    }, {
      "utf8": " stressed",
      "tOffsetMs": 1199,
      "acAsrConf": 252
    }, {
      "utf8": " speaking",
      "tOffsetMs": 1709,
      "acAsrConf": 211
    }, {
      "utf8": " with",
      "tOffsetMs": 2310,
      "acAsrConf": 223
    } ]
  }, {
    "tStartMs": 2481579,
    "dDurationMs": 2530,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2481589,
    "dDurationMs": 4591,
    "wWinId": 1,
    "segs": [ {
      "utf8": "maybe",
      "acAsrConf": 245
    }, {
      "utf8": " that's",
      "tOffsetMs": 720,
      "acAsrConf": 251
    }, {
      "utf8": " going",
      "tOffsetMs": 1020,
      "acAsrConf": 242
    }, {
      "utf8": " to",
      "tOffsetMs": 1291,
      "acAsrConf": 250
    }, {
      "utf8": " be",
      "tOffsetMs": 1411,
      "acAsrConf": 252
    }, {
      "utf8": " someone",
      "tOffsetMs": 1530,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 1950,
      "acAsrConf": 227
    }, {
      "utf8": " a",
      "tOffsetMs": 2250,
      "acAsrConf": 255
    } ]
  }, {
    "tStartMs": 2484099,
    "dDurationMs": 2081,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2484109,
    "dDurationMs": 4621,
    "wWinId": 1,
    "segs": [ {
      "utf8": "meetup",
      "acAsrConf": 239
    }, {
      "utf8": " group",
      "tOffsetMs": 331,
      "acAsrConf": 203
    }, {
      "utf8": " from",
      "tOffsetMs": 480,
      "acAsrConf": 165
    }, {
      "utf8": " Meetup",
      "tOffsetMs": 990,
      "acAsrConf": 142
    }, {
      "utf8": " calm",
      "tOffsetMs": 1291,
      "acAsrConf": 227
    }, {
      "utf8": " maybe",
      "tOffsetMs": 1710,
      "acAsrConf": 218
    } ]
  }, {
    "tStartMs": 2486170,
    "dDurationMs": 2560,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2486180,
    "dDurationMs": 5399,
    "wWinId": 1,
    "segs": [ {
      "utf8": "that's",
      "acAsrConf": 247
    }, {
      "utf8": " going",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 480,
      "acAsrConf": 227
    }, {
      "utf8": " be",
      "tOffsetMs": 599,
      "acAsrConf": 248
    }, {
      "utf8": " online",
      "tOffsetMs": 720,
      "acAsrConf": 227
    }, {
      "utf8": " but",
      "tOffsetMs": 929,
      "acAsrConf": 208
    }, {
      "utf8": " if",
      "tOffsetMs": 1859,
      "acAsrConf": 210
    }, {
      "utf8": " you",
      "tOffsetMs": 2070,
      "acAsrConf": 252
    }, {
      "utf8": " are",
      "tOffsetMs": 2280,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2488720,
    "dDurationMs": 2859,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2488730,
    "dDurationMs": 5220,
    "wWinId": 1,
    "segs": [ {
      "utf8": "giving",
      "acAsrConf": 252
    }, {
      "utf8": " up",
      "tOffsetMs": 900,
      "acAsrConf": 252
    }, {
      "utf8": " speaking",
      "tOffsetMs": 1079,
      "acAsrConf": 221
    }, {
      "utf8": " with",
      "tOffsetMs": 1799,
      "acAsrConf": 243
    }, {
      "utf8": " someone",
      "tOffsetMs": 1950,
      "acAsrConf": 252
    }, {
      "utf8": " because",
      "tOffsetMs": 2220,
      "acAsrConf": 145
    } ]
  }, {
    "tStartMs": 2491569,
    "dDurationMs": 2381,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2491579,
    "dDurationMs": 6091,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you're",
      "acAsrConf": 238
    }, {
      "utf8": " so",
      "tOffsetMs": 601,
      "acAsrConf": 252
    }, {
      "utf8": " worried",
      "tOffsetMs": 901,
      "acAsrConf": 252
    }, {
      "utf8": " about",
      "tOffsetMs": 1351,
      "acAsrConf": 252
    }, {
      "utf8": " your",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    }, {
      "utf8": " sentence",
      "tOffsetMs": 1710,
      "acAsrConf": 200
    } ]
  }, {
    "tStartMs": 2493940,
    "dDurationMs": 3730,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2493950,
    "dDurationMs": 5669,
    "wWinId": 1,
    "segs": [ {
      "utf8": "structure",
      "acAsrConf": 252
    }, {
      "utf8": " just",
      "tOffsetMs": 859,
      "acAsrConf": 204
    }, {
      "utf8": " try",
      "tOffsetMs": 1859,
      "acAsrConf": 225
    }, {
      "utf8": " to",
      "tOffsetMs": 2669,
      "acAsrConf": 232
    }, {
      "utf8": " find",
      "tOffsetMs": 2730,
      "acAsrConf": 230
    }, {
      "utf8": " someone",
      "tOffsetMs": 3180,
      "acAsrConf": 252
    }, {
      "utf8": " who",
      "tOffsetMs": 3510,
      "acAsrConf": 248
    } ]
  }, {
    "tStartMs": 2497660,
    "dDurationMs": 1959,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2497670,
    "dDurationMs": 3510,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 186
    }, {
      "utf8": " can",
      "tOffsetMs": 149,
      "acAsrConf": 252
    }, {
      "utf8": " speak",
      "tOffsetMs": 359,
      "acAsrConf": 252
    }, {
      "utf8": " with",
      "tOffsetMs": 750,
      "acAsrConf": 237
    }, {
      "utf8": " and",
      "tOffsetMs": 990,
      "acAsrConf": 249
    }, {
      "utf8": " you're",
      "tOffsetMs": 1230,
      "acAsrConf": 252
    }, {
      "utf8": " not",
      "tOffsetMs": 1439,
      "acAsrConf": 201
    }, {
      "utf8": " going",
      "tOffsetMs": 1649,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2499609,
    "dDurationMs": 1571,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2499619,
    "dDurationMs": 3750,
    "wWinId": 1,
    "segs": [ {
      "utf8": "to",
      "acAsrConf": 200
    }, {
      "utf8": " feel",
      "tOffsetMs": 121,
      "acAsrConf": 227
    }, {
      "utf8": " stressed",
      "tOffsetMs": 301,
      "acAsrConf": 252
    }, {
      "utf8": " with",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " because",
      "tOffsetMs": 841,
      "acAsrConf": 239
    }, {
      "utf8": " stress",
      "tOffsetMs": 1170,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 1531,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2501170,
    "dDurationMs": 2199,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2501180,
    "dDurationMs": 6329,
    "wWinId": 1,
    "segs": [ {
      "utf8": "not",
      "acAsrConf": 206
    }, {
      "utf8": " going",
      "tOffsetMs": 389,
      "acAsrConf": 235
    }, {
      "utf8": " to",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " help",
      "tOffsetMs": 869,
      "acAsrConf": 230
    }, {
      "utf8": " yes",
      "tOffsetMs": 1050,
      "acAsrConf": 214
    }, {
      "utf8": " we",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 1710,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 1859,
      "acAsrConf": 252
    }, {
      "utf8": " saying",
      "tOffsetMs": 1889,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2503359,
    "dDurationMs": 4150,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2503369,
    "dDurationMs": 8450,
    "wWinId": 1,
    "segs": [ {
      "utf8": "in",
      "acAsrConf": 198
    }, {
      "utf8": " English",
      "tOffsetMs": 361,
      "acAsrConf": 240
    }, {
      "utf8": " take",
      "tOffsetMs": 571,
      "acAsrConf": 180
    }, {
      "utf8": " baby",
      "tOffsetMs": 1381,
      "acAsrConf": 236
    }, {
      "utf8": " steps",
      "tOffsetMs": 1950,
      "acAsrConf": 218
    }, {
      "utf8": " which",
      "tOffsetMs": 2900,
      "acAsrConf": 200
    }, {
      "utf8": " means",
      "tOffsetMs": 3900,
      "acAsrConf": 243
    } ]
  }, {
    "tStartMs": 2507499,
    "dDurationMs": 4320,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2507509,
    "dDurationMs": 6991,
    "wWinId": 1,
    "segs": [ {
      "utf8": "try",
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 360,
      "acAsrConf": 252
    }, {
      "utf8": " take",
      "tOffsetMs": 421,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 691,
      "acAsrConf": 252
    }, {
      "utf8": " small",
      "tOffsetMs": 721,
      "acAsrConf": 252
    }, {
      "utf8": " step",
      "tOffsetMs": 1201,
      "acAsrConf": 227
    }, {
      "utf8": " if",
      "tOffsetMs": 2040,
      "acAsrConf": 246
    }, {
      "utf8": " you",
      "tOffsetMs": 2371,
      "acAsrConf": 252
    }, {
      "utf8": " want",
      "tOffsetMs": 2431,
      "acAsrConf": 145
    }, {
      "utf8": " to",
      "tOffsetMs": 3080,
      "acAsrConf": 145
    } ]
  }, {
    "tStartMs": 2511809,
    "dDurationMs": 2691,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2511819,
    "dDurationMs": 5050,
    "wWinId": 1,
    "segs": [ {
      "utf8": "improve",
      "acAsrConf": 252
    }, {
      "utf8": " your",
      "tOffsetMs": 1000,
      "acAsrConf": 252
    }, {
      "utf8": " grammar",
      "tOffsetMs": 1210,
      "acAsrConf": 216
    }, {
      "utf8": " or",
      "tOffsetMs": 1450,
      "acAsrConf": 252
    }, {
      "utf8": " structure",
      "tOffsetMs": 1841,
      "acAsrConf": 252
    }, {
      "utf8": " don't",
      "tOffsetMs": 2351,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2514490,
    "dDurationMs": 2379,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2514500,
    "dDurationMs": 3240,
    "wWinId": 1,
    "segs": [ {
      "utf8": "take",
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 299,
      "acAsrConf": 252
    }, {
      "utf8": " big",
      "tOffsetMs": 329,
      "acAsrConf": 249
    }, {
      "utf8": " step",
      "tOffsetMs": 779,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1170,
      "acAsrConf": 228
    }, {
      "utf8": " feel",
      "tOffsetMs": 1440,
      "acAsrConf": 217
    }, {
      "utf8": " very",
      "tOffsetMs": 1859,
      "acAsrConf": 252
    }, {
      "utf8": " scared",
      "tOffsetMs": 1890,
      "acAsrConf": 179
    } ]
  }, {
    "tStartMs": 2516859,
    "dDurationMs": 881,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2516869,
    "dDurationMs": 2731,
    "wWinId": 1,
    "segs": [ {
      "utf8": "right",
      "acAsrConf": 147
    } ]
  }, {
    "tStartMs": 2517730,
    "dDurationMs": 1870,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2517740,
    "dDurationMs": 5040,
    "wWinId": 1,
    "segs": [ {
      "utf8": "maybe",
      "acAsrConf": 252
    }, {
      "utf8": " just",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " take",
      "tOffsetMs": 780,
      "acAsrConf": 209
    }, {
      "utf8": " a",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " small",
      "tOffsetMs": 990,
      "acAsrConf": 213
    }, {
      "utf8": " step",
      "tOffsetMs": 1290,
      "acAsrConf": 252
    }, {
      "utf8": " try",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1800,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2519590,
    "dDurationMs": 3190,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2519600,
    "dDurationMs": 5910,
    "wWinId": 1,
    "segs": [ {
      "utf8": "listen",
      "acAsrConf": 214
    }, {
      "utf8": " to",
      "tOffsetMs": 420,
      "acAsrConf": 201
    }, {
      "utf8": " something",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " regularly",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " uh",
      "tOffsetMs": 1590,
      "acAsrConf": 147
    }, {
      "utf8": " maybe",
      "tOffsetMs": 2180,
      "acAsrConf": 237
    } ]
  }, {
    "tStartMs": 2522770,
    "dDurationMs": 2740,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2522780,
    "dDurationMs": 6960,
    "wWinId": 1,
    "segs": [ {
      "utf8": "speak",
      "acAsrConf": 247
    }, {
      "utf8": " with",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " somebody",
      "tOffsetMs": 900,
      "acAsrConf": 252
    }, {
      "utf8": " online",
      "tOffsetMs": 1110,
      "acAsrConf": 201
    }, {
      "utf8": " and",
      "tOffsetMs": 1500,
      "acAsrConf": 207
    }, {
      "utf8": " again",
      "tOffsetMs": 2190,
      "acAsrConf": 201
    }, {
      "utf8": " a",
      "tOffsetMs": 2700,
      "acAsrConf": 149
    } ]
  }, {
    "tStartMs": 2525500,
    "dDurationMs": 4240,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2525510,
    "dDurationMs": 6330,
    "wWinId": 1,
    "segs": [ {
      "utf8": "friend",
      "acAsrConf": 200
    }, {
      "utf8": " or",
      "tOffsetMs": 420,
      "acAsrConf": 207
    }, {
      "utf8": " somebody",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " who",
      "tOffsetMs": 990,
      "acAsrConf": 248
    }, {
      "utf8": " who",
      "tOffsetMs": 2150,
      "acAsrConf": 232
    }, {
      "utf8": " you",
      "tOffsetMs": 3150,
      "acAsrConf": 226
    }, {
      "utf8": " are",
      "tOffsetMs": 3420,
      "acAsrConf": 252
    }, {
      "utf8": " more",
      "tOffsetMs": 3750,
      "acAsrConf": 242
    } ]
  }, {
    "tStartMs": 2529730,
    "dDurationMs": 2110,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2529740,
    "dDurationMs": 3950,
    "wWinId": 1,
    "segs": [ {
      "utf8": "comfortable",
      "acAsrConf": 252
    }, {
      "utf8": " making",
      "tOffsetMs": 240,
      "acAsrConf": 209
    }, {
      "utf8": " a",
      "tOffsetMs": 810,
      "acAsrConf": 208
    }, {
      "utf8": " mistake",
      "tOffsetMs": 990,
      "acAsrConf": 252
    }, {
      "utf8": " with",
      "tOffsetMs": 1110,
      "acAsrConf": 252
    }, {
      "utf8": " by",
      "tOffsetMs": 1350,
      "acAsrConf": 225
    } ]
  }, {
    "tStartMs": 2531830,
    "dDurationMs": 1860,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2531840,
    "dDurationMs": 6660,
    "wWinId": 1,
    "segs": [ {
      "utf8": "step",
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 360,
      "acAsrConf": 230
    }, {
      "utf8": " you",
      "tOffsetMs": 480,
      "acAsrConf": 200
    }, {
      "utf8": " also",
      "tOffsetMs": 630,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 840,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1140,
      "acAsrConf": 252
    }, {
      "utf8": " remember",
      "tOffsetMs": 1290,
      "acAsrConf": 207
    } ]
  }, {
    "tStartMs": 2533680,
    "dDurationMs": 4820,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2533690,
    "dDurationMs": 8410,
    "wWinId": 1,
    "segs": [ {
      "utf8": "everybody",
      "acAsrConf": 212
    }, {
      "utf8": " makes",
      "tOffsetMs": 1080,
      "acAsrConf": 252
    }, {
      "utf8": " mistakes",
      "tOffsetMs": 2080,
      "acAsrConf": 207
    }, {
      "utf8": " yes",
      "tOffsetMs": 2590,
      "acAsrConf": 210
    }, {
      "utf8": " yes",
      "tOffsetMs": 3400,
      "acAsrConf": 210
    }, {
      "utf8": " I",
      "tOffsetMs": 4240,
      "acAsrConf": 202
    }, {
      "utf8": " make",
      "tOffsetMs": 4450,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2538490,
    "dDurationMs": 3610,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2538500,
    "dDurationMs": 4590,
    "wWinId": 1,
    "segs": [ {
      "utf8": "English",
      "acAsrConf": 245
    }, {
      "utf8": " mistakes",
      "tOffsetMs": 360,
      "acAsrConf": 252
    }, {
      "utf8": " oh",
      "tOffsetMs": 570,
      "acAsrConf": 127
    }, {
      "utf8": " I",
      "tOffsetMs": 1369,
      "acAsrConf": 240
    }, {
      "utf8": " am",
      "tOffsetMs": 2369,
      "acAsrConf": 243
    }, {
      "utf8": " a",
      "tOffsetMs": 3000,
      "acAsrConf": 252
    }, {
      "utf8": " very",
      "tOffsetMs": 3030,
      "acAsrConf": 250
    }, {
      "utf8": " good",
      "tOffsetMs": 3330,
      "acAsrConf": 145
    } ]
  }, {
    "tStartMs": 2542090,
    "dDurationMs": 1000,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2542100,
    "dDurationMs": 3060,
    "wWinId": 1,
    "segs": [ {
      "utf8": "English",
      "acAsrConf": 252
    }, {
      "utf8": " speaker",
      "tOffsetMs": 300,
      "acAsrConf": 154
    } ]
  }, {
    "tStartMs": 2543080,
    "dDurationMs": 2080,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2543090,
    "dDurationMs": 3630,
    "wWinId": 1,
    "segs": [ {
      "utf8": "he",
      "acAsrConf": 252
    }, {
      "utf8": " has",
      "tOffsetMs": 630,
      "acAsrConf": 252
    }, {
      "utf8": " everyone",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " makes",
      "tOffsetMs": 1140,
      "acAsrConf": 252
    }, {
      "utf8": " mistakes",
      "tOffsetMs": 1350,
      "acAsrConf": 252
    }, {
      "utf8": " alright",
      "tOffsetMs": 1890,
      "acAsrConf": 66
    } ]
  }, {
    "tStartMs": 2545150,
    "dDurationMs": 1570,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2545160,
    "dDurationMs": 4260,
    "wWinId": 1,
    "segs": [ {
      "utf8": "let's",
      "acAsrConf": 252
    }, {
      "utf8": " go",
      "tOffsetMs": 150,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 270,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 300,
      "acAsrConf": 217
    }, {
      "utf8": " next",
      "tOffsetMs": 480,
      "acAsrConf": 143
    }, {
      "utf8": " question",
      "tOffsetMs": 870,
      "acAsrConf": 200
    }, {
      "utf8": " all",
      "tOffsetMs": 1050,
      "acAsrConf": 214
    }, {
      "utf8": " we",
      "tOffsetMs": 1410,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2546710,
    "dDurationMs": 2710,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2546720,
    "dDurationMs": 4710,
    "wWinId": 1,
    "segs": [ {
      "utf8": "have",
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 120,
      "acAsrConf": 252
    }, {
      "utf8": " nice",
      "tOffsetMs": 149,
      "acAsrConf": 252
    }, {
      "utf8": " comment",
      "tOffsetMs": 390,
      "acAsrConf": 252
    }, {
      "utf8": " from",
      "tOffsetMs": 660,
      "acAsrConf": 203
    }, {
      "utf8": " Marcia",
      "tOffsetMs": 1050,
      "acAsrConf": 199
    }, {
      "utf8": " Lewis",
      "tOffsetMs": 1920,
      "acAsrConf": 208
    }, {
      "utf8": " he",
      "tOffsetMs": 2399,
      "acAsrConf": 142
    } ]
  }, {
    "tStartMs": 2549410,
    "dDurationMs": 2020,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2549420,
    "dDurationMs": 6660,
    "wWinId": 1,
    "segs": [ {
      "utf8": "says",
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 240,
      "acAsrConf": 234
    }, {
      "utf8": " bet",
      "tOffsetMs": 449,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 870,
      "acAsrConf": 228
    }, {
      "utf8": " will",
      "tOffsetMs": 1080,
      "acAsrConf": 252
    }, {
      "utf8": " open",
      "tOffsetMs": 1230,
      "acAsrConf": 252
    }, {
      "utf8": " your",
      "tOffsetMs": 1410,
      "acAsrConf": 208
    }, {
      "utf8": " coffee",
      "tOffsetMs": 1650,
      "acAsrConf": 223
    } ]
  }, {
    "tStartMs": 2551420,
    "dDurationMs": 4660,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2551430,
    "dDurationMs": 6810,
    "wWinId": 1,
    "segs": [ {
      "utf8": "shop",
      "acAsrConf": 252
    }, {
      "utf8": " oh",
      "tOffsetMs": 210,
      "acAsrConf": 98
    }, {
      "utf8": " yes",
      "tOffsetMs": 840,
      "acAsrConf": 247
    }, {
      "utf8": " yes",
      "tOffsetMs": 1830,
      "acAsrConf": 247
    }, {
      "utf8": " yes",
      "tOffsetMs": 2250,
      "acAsrConf": 207
    }, {
      "utf8": " alright",
      "tOffsetMs": 2840,
      "acAsrConf": 0
    }, {
      "utf8": " and",
      "tOffsetMs": 3840,
      "acAsrConf": 144
    }, {
      "utf8": " we",
      "tOffsetMs": 4200,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 4560,
      "acAsrConf": 211
    } ]
  }, {
    "tStartMs": 2556070,
    "dDurationMs": 2170,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2556080,
    "dDurationMs": 3930,
    "wWinId": 1,
    "segs": [ {
      "utf8": "another",
      "acAsrConf": 252
    }, {
      "utf8": " question",
      "tOffsetMs": 240,
      "acAsrConf": 151
    }, {
      "utf8": " from",
      "tOffsetMs": 690,
      "acAsrConf": 161
    }, {
      "utf8": " Allah",
      "tOffsetMs": 960,
      "acAsrConf": 243
    }, {
      "utf8": " from",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    }, {
      "utf8": " evaldo",
      "tOffsetMs": 1710,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2558230,
    "dDurationMs": 1780,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2558240,
    "dDurationMs": 3900,
    "wWinId": 1,
    "segs": [ {
      "utf8": "thank",
      "acAsrConf": 226
    }, {
      "utf8": " you",
      "tOffsetMs": 270,
      "acAsrConf": 226
    }, {
      "utf8": " she",
      "tOffsetMs": 360,
      "acAsrConf": 252
    }, {
      "utf8": " says",
      "tOffsetMs": 570,
      "acAsrConf": 252
    }, {
      "utf8": " I've",
      "tOffsetMs": 780,
      "acAsrConf": 242
    }, {
      "utf8": " been",
      "tOffsetMs": 960,
      "acAsrConf": 99
    }, {
      "utf8": " to",
      "tOffsetMs": 1170,
      "acAsrConf": 201
    }, {
      "utf8": " Madison",
      "tOffsetMs": 1320,
      "acAsrConf": 236
    } ]
  }, {
    "tStartMs": 2560000,
    "dDurationMs": 2140,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2560010,
    "dDurationMs": 4230,
    "wWinId": 1,
    "segs": [ {
      "utf8": "Wisconsin",
      "acAsrConf": 236
    }, {
      "utf8": " once",
      "tOffsetMs": 150,
      "acAsrConf": 252
    }, {
      "utf8": " are",
      "tOffsetMs": 840,
      "acAsrConf": 189
    }, {
      "utf8": " there",
      "tOffsetMs": 1380,
      "acAsrConf": 252
    }, {
      "utf8": " any",
      "tOffsetMs": 1950,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2562130,
    "dDurationMs": 2110,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2562140,
    "dDurationMs": 4260,
    "wWinId": 1,
    "segs": [ {
      "utf8": "similarities",
      "acAsrConf": 239
    }, {
      "utf8": " between",
      "tOffsetMs": 330,
      "acAsrConf": 215
    }, {
      "utf8": " this",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1350,
      "acAsrConf": 252
    }, {
      "utf8": " your",
      "tOffsetMs": 1590,
      "acAsrConf": 252
    }, {
      "utf8": " city",
      "tOffsetMs": 1800,
      "acAsrConf": 206
    } ]
  }, {
    "tStartMs": 2564230,
    "dDurationMs": 2170,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2564240,
    "dDurationMs": 3510,
    "wWinId": 1,
    "segs": [ {
      "utf8": "oh",
      "acAsrConf": 93
    }, {
      "utf8": " I'm",
      "tOffsetMs": 180,
      "acAsrConf": 255
    }, {
      "utf8": " sorry",
      "tOffsetMs": 660,
      "acAsrConf": 245
    }, {
      "utf8": " it",
      "tOffsetMs": 900,
      "acAsrConf": 218
    }, {
      "utf8": " is",
      "tOffsetMs": 1170,
      "acAsrConf": 206
    }, {
      "utf8": " warmer",
      "tOffsetMs": 1320,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2566390,
    "dDurationMs": 1360,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2566400,
    "dDurationMs": 3660,
    "wWinId": 1,
    "segs": [ {
      "utf8": "oh",
      "acAsrConf": 231
    }, {
      "utf8": " it's",
      "tOffsetMs": 30,
      "acAsrConf": 252
    }, {
      "utf8": " warmer",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " I've",
      "tOffsetMs": 600,
      "acAsrConf": 248
    }, {
      "utf8": " never",
      "tOffsetMs": 750,
      "acAsrConf": 247
    }, {
      "utf8": " been",
      "tOffsetMs": 990,
      "acAsrConf": 249
    }, {
      "utf8": " to",
      "tOffsetMs": 1230,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2567740,
    "dDurationMs": 2320,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2567750,
    "dDurationMs": 5010,
    "wWinId": 1,
    "segs": [ {
      "utf8": "Madison",
      "acAsrConf": 251
    }, {
      "utf8": " Wisconsin",
      "tOffsetMs": 420,
      "acAsrConf": 251
    }, {
      "utf8": " but",
      "tOffsetMs": 570,
      "acAsrConf": 123
    }, {
      "utf8": " what",
      "tOffsetMs": 1320,
      "acAsrConf": 246
    }, {
      "utf8": " we",
      "tOffsetMs": 1800,
      "acAsrConf": 252
    }, {
      "utf8": " know",
      "tOffsetMs": 1980,
      "acAsrConf": 252
    }, {
      "utf8": " about",
      "tOffsetMs": 2130,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2570050,
    "dDurationMs": 2710,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2570060,
    "dDurationMs": 6000,
    "wWinId": 1,
    "segs": [ {
      "utf8": "Madison",
      "acAsrConf": 236
    }, {
      "utf8": " Wisconsin",
      "tOffsetMs": 660,
      "acAsrConf": 236
    }, {
      "utf8": " is",
      "tOffsetMs": 1290,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 1410,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 1650,
      "acAsrConf": 237
    }, {
      "utf8": " is",
      "tOffsetMs": 1830,
      "acAsrConf": 252
    }, {
      "utf8": " very",
      "tOffsetMs": 2010,
      "acAsrConf": 200
    }, {
      "utf8": " cold",
      "tOffsetMs": 2340,
      "acAsrConf": 248
    } ]
  }, {
    "tStartMs": 2572750,
    "dDurationMs": 3310,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2572760,
    "dDurationMs": 6500,
    "wWinId": 1,
    "segs": [ {
      "utf8": "for",
      "acAsrConf": 187
    }, {
      "utf8": " much",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 1140,
      "acAsrConf": 203
    }, {
      "utf8": " the",
      "tOffsetMs": 1320,
      "acAsrConf": 221
    }, {
      "utf8": " year",
      "tOffsetMs": 1470,
      "acAsrConf": 252
    }, {
      "utf8": " yes",
      "tOffsetMs": 1770,
      "acAsrConf": 207
    }, {
      "utf8": " yes",
      "tOffsetMs": 2190,
      "acAsrConf": 247
    }, {
      "utf8": " it's",
      "tOffsetMs": 2850,
      "acAsrConf": 252
    }, {
      "utf8": " very",
      "tOffsetMs": 3210,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2576050,
    "dDurationMs": 3210,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2576060,
    "dDurationMs": 4770,
    "wWinId": 1,
    "segs": [ {
      "utf8": "cold",
      "acAsrConf": 230
    }, {
      "utf8": " so",
      "tOffsetMs": 740,
      "acAsrConf": 151
    }, {
      "utf8": " let's",
      "tOffsetMs": 1740,
      "acAsrConf": 203
    }, {
      "utf8": " see",
      "tOffsetMs": 2040,
      "acAsrConf": 252
    }, {
      "utf8": " we",
      "tOffsetMs": 2190,
      "acAsrConf": 240
    }, {
      "utf8": " have",
      "tOffsetMs": 2400,
      "acAsrConf": 252
    }, {
      "utf8": " another",
      "tOffsetMs": 2580,
      "acAsrConf": 226
    } ]
  }, {
    "tStartMs": 2579250,
    "dDurationMs": 1580,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2579260,
    "dDurationMs": 4120,
    "wWinId": 1,
    "segs": [ {
      "utf8": "question",
      "acAsrConf": 166
    }, {
      "utf8": " here",
      "tOffsetMs": 1000,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2580820,
    "dDurationMs": 2560,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2580830,
    "dDurationMs": 6900,
    "wWinId": 1,
    "segs": [ {
      "utf8": "aha",
      "acAsrConf": 217
    }, {
      "utf8": " would",
      "tOffsetMs": 480,
      "acAsrConf": 236
    }, {
      "utf8": " you",
      "tOffsetMs": 1440,
      "acAsrConf": 252
    }, {
      "utf8": " like",
      "tOffsetMs": 1560,
      "acAsrConf": 243
    }, {
      "utf8": " to",
      "tOffsetMs": 1710,
      "acAsrConf": 252
    }, {
      "utf8": " read",
      "tOffsetMs": 1860,
      "acAsrConf": 252
    }, {
      "utf8": " this",
      "tOffsetMs": 2039,
      "acAsrConf": 212
    }, {
      "utf8": " question",
      "tOffsetMs": 2100,
      "acAsrConf": 148
    } ]
  }, {
    "tStartMs": 2583370,
    "dDurationMs": 4360,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2583380,
    "dDurationMs": 7950,
    "wWinId": 1,
    "segs": [ {
      "utf8": "hi",
      "acAsrConf": 205
    }, {
      "utf8": " this",
      "tOffsetMs": 960,
      "acAsrConf": 245
    }, {
      "utf8": " is",
      "tOffsetMs": 1050,
      "acAsrConf": 206
    }, {
      "utf8": " from",
      "tOffsetMs": 1200,
      "acAsrConf": 231
    }, {
      "utf8": " Jer",
      "tOffsetMs": 2840,
      "acAsrConf": 151
    } ]
  }, {
    "tStartMs": 2587720,
    "dDurationMs": 3610,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2587730,
    "dDurationMs": 6330,
    "wWinId": 1,
    "segs": [ {
      "utf8": "where",
      "acAsrConf": 165
    }, {
      "utf8": " I",
      "tOffsetMs": 300,
      "acAsrConf": 142
    }, {
      "utf8": " you",
      "tOffsetMs": 570,
      "acAsrConf": 84
    }, {
      "utf8": " know",
      "tOffsetMs": 1080,
      "acAsrConf": 0
    }, {
      "utf8": " I",
      "tOffsetMs": 1590,
      "acAsrConf": 243
    }, {
      "utf8": " you",
      "tOffsetMs": 1620,
      "acAsrConf": 0
    }, {
      "utf8": " I",
      "tOffsetMs": 2270,
      "acAsrConf": 240
    }, {
      "utf8": " don't",
      "tOffsetMs": 3270,
      "acAsrConf": 197
    }, {
      "utf8": " know",
      "tOffsetMs": 3510,
      "acAsrConf": 240
    } ]
  }, {
    "tStartMs": 2591320,
    "dDurationMs": 2740,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2591330,
    "dDurationMs": 5070,
    "wWinId": 1,
    "segs": [ {
      "utf8": "maybe",
      "acAsrConf": 236
    }, {
      "utf8": " we",
      "tOffsetMs": 270,
      "acAsrConf": 237
    }, {
      "utf8": " need",
      "tOffsetMs": 510,
      "acAsrConf": 206
    }, {
      "utf8": " to",
      "tOffsetMs": 660,
      "acAsrConf": 137
    }, {
      "utf8": " practice",
      "tOffsetMs": 779,
      "acAsrConf": 252
    }, {
      "utf8": " our",
      "tOffsetMs": 1039,
      "acAsrConf": 122
    }, {
      "utf8": " Spanish",
      "tOffsetMs": 2039,
      "acAsrConf": 232
    }, {
      "utf8": " or",
      "tOffsetMs": 2340,
      "acAsrConf": 143
    } ]
  }, {
    "tStartMs": 2594050,
    "dDurationMs": 2350,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2594060,
    "dDurationMs": 4500,
    "wWinId": 1,
    "segs": [ {
      "utf8": "Portuguese",
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 30,
      "acAsrConf": 244
    }, {
      "utf8": " think",
      "tOffsetMs": 780,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1049,
      "acAsrConf": 252
    }, {
      "utf8": " yeah",
      "tOffsetMs": 1260,
      "acAsrConf": 222
    }, {
      "utf8": " like",
      "tOffsetMs": 1680,
      "acAsrConf": 145
    }, {
      "utf8": " I",
      "tOffsetMs": 2250,
      "acAsrConf": 207
    } ]
  }, {
    "tStartMs": 2596390,
    "dDurationMs": 2170,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2596400,
    "dDurationMs": 2880,
    "wWinId": 1,
    "segs": [ {
      "utf8": "realized",
      "acAsrConf": 221
    }, {
      "utf8": " I",
      "tOffsetMs": 630,
      "acAsrConf": 245
    }, {
      "utf8": " understand",
      "tOffsetMs": 870,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    }, {
      "utf8": " better",
      "tOffsetMs": 1680,
      "acAsrConf": 252
    }, {
      "utf8": " than",
      "tOffsetMs": 1860,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 2598550,
    "dDurationMs": 730,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2598560,
    "dDurationMs": 4520,
    "wWinId": 1,
    "segs": [ {
      "utf8": "other",
      "acAsrConf": 238
    }, {
      "utf8": " people",
      "tOffsetMs": 270,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2599270,
    "dDurationMs": 3810,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2599280,
    "dDurationMs": 7079,
    "wWinId": 1,
    "segs": [ {
      "utf8": "Oh",
      "acAsrConf": 0
    }, {
      "utf8": " I",
      "tOffsetMs": 120,
      "acAsrConf": 147
    }, {
      "utf8": " think",
      "tOffsetMs": 600,
      "acAsrConf": 202
    }, {
      "utf8": " your",
      "tOffsetMs": 1170,
      "acAsrConf": 252
    }, {
      "utf8": " accent",
      "tOffsetMs": 1440,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 2070,
      "acAsrConf": 236
    }, {
      "utf8": " perfect",
      "tOffsetMs": 2250,
      "acAsrConf": 247
    }, {
      "utf8": " tone",
      "tOffsetMs": 3089,
      "acAsrConf": 230
    } ]
  }, {
    "tStartMs": 2603070,
    "dDurationMs": 3289,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2603080,
    "dDurationMs": 5770,
    "wWinId": 1,
    "segs": [ {
      "utf8": "perfect",
      "acAsrConf": 216
    }, {
      "utf8": " to",
      "tOffsetMs": 1000,
      "acAsrConf": 207
    }, {
      "utf8": " me",
      "tOffsetMs": 1090,
      "acAsrConf": 252
    }, {
      "utf8": " Oh",
      "tOffsetMs": 1240,
      "acAsrConf": 202
    }, {
      "utf8": " to",
      "tOffsetMs": 1420,
      "acAsrConf": 252
    }, {
      "utf8": " me",
      "tOffsetMs": 1810,
      "acAsrConf": 252
    }, {
      "utf8": " oh",
      "tOffsetMs": 2140,
      "acAsrConf": 91
    }, {
      "utf8": " thank",
      "tOffsetMs": 2380,
      "acAsrConf": 242
    }, {
      "utf8": " you",
      "tOffsetMs": 2890,
      "acAsrConf": 219
    } ]
  }, {
    "tStartMs": 2606349,
    "dDurationMs": 2501,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2606359,
    "dDurationMs": 5131,
    "wWinId": 1,
    "segs": [ {
      "utf8": "I'd",
      "acAsrConf": 252
    }, {
      "utf8": " like",
      "tOffsetMs": 391,
      "acAsrConf": 234
    }, {
      "utf8": " to",
      "tOffsetMs": 691,
      "acAsrConf": 207
    }, {
      "utf8": " know",
      "tOffsetMs": 841,
      "acAsrConf": 242
    }, {
      "utf8": " where",
      "tOffsetMs": 1021,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1291,
      "acAsrConf": 252
    }, {
      "utf8": " are",
      "tOffsetMs": 1321,
      "acAsrConf": 65
    }, {
      "utf8": " born",
      "tOffsetMs": 1591,
      "acAsrConf": 252
    }, {
      "utf8": " what",
      "tOffsetMs": 1801,
      "acAsrConf": 200
    } ]
  }, {
    "tStartMs": 2608840,
    "dDurationMs": 2650,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2608850,
    "dDurationMs": 5670,
    "wWinId": 1,
    "segs": [ {
      "utf8": "city",
      "acAsrConf": 252
    }, {
      "utf8": " do",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 390,
      "acAsrConf": 252
    }, {
      "utf8": " live",
      "tOffsetMs": 450,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 690,
      "acAsrConf": 227
    }, {
      "utf8": " sounds",
      "tOffsetMs": 1080,
      "acAsrConf": 252
    }, {
      "utf8": " like",
      "tOffsetMs": 1470,
      "acAsrConf": 228
    }, {
      "utf8": " there",
      "tOffsetMs": 1680,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 2220,
      "acAsrConf": 153
    } ]
  }, {
    "tStartMs": 2611480,
    "dDurationMs": 3040,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2611490,
    "dDurationMs": 5040,
    "wWinId": 1,
    "segs": [ {
      "utf8": "a",
      "acAsrConf": 252
    }, {
      "utf8": " perfect",
      "tOffsetMs": 30,
      "acAsrConf": 252
    }, {
      "utf8": " accent",
      "tOffsetMs": 600,
      "acAsrConf": 226
    }, {
      "utf8": " for",
      "tOffsetMs": 1080,
      "acAsrConf": 252
    }, {
      "utf8": " each",
      "tOffsetMs": 1170,
      "acAsrConf": 225
    }, {
      "utf8": " ear",
      "tOffsetMs": 1440,
      "acAsrConf": 241
    }, {
      "utf8": " I",
      "tOffsetMs": 1890,
      "acAsrConf": 225
    }, {
      "utf8": " love",
      "tOffsetMs": 2100,
      "acAsrConf": 227
    }, {
      "utf8": " you",
      "tOffsetMs": 2879,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2614510,
    "dDurationMs": 2020,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2614520,
    "dDurationMs": 4650,
    "wWinId": 1,
    "segs": [ {
      "utf8": "so",
      "acAsrConf": 252
    }, {
      "utf8": " much",
      "tOffsetMs": 270,
      "acAsrConf": 252
    }, {
      "utf8": " thank",
      "tOffsetMs": 560,
      "acAsrConf": 200
    }, {
      "utf8": " you",
      "tOffsetMs": 1560,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 2616520,
    "dDurationMs": 2650,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2616530,
    "dDurationMs": 6060,
    "wWinId": 1,
    "segs": [ {
      "utf8": "Jerry",
      "acAsrConf": 204
    }, {
      "utf8": " from",
      "tOffsetMs": 329,
      "acAsrConf": 218
    }, {
      "utf8": " Brazil",
      "tOffsetMs": 570,
      "acAsrConf": 233
    }, {
      "utf8": " oh",
      "tOffsetMs": 1380,
      "acAsrConf": 216
    }, {
      "utf8": " thank",
      "tOffsetMs": 1589,
      "acAsrConf": 234
    }, {
      "utf8": " you",
      "tOffsetMs": 2160,
      "acAsrConf": 210
    }, {
      "utf8": " EJ",
      "tOffsetMs": 2310,
      "acAsrConf": 239
    }, {
      "utf8": " oh",
      "tOffsetMs": 2610,
      "acAsrConf": 147
    } ]
  }, {
    "tStartMs": 2619160,
    "dDurationMs": 3430,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2619170,
    "dDurationMs": 5370,
    "wWinId": 1,
    "segs": [ {
      "utf8": "that's",
      "acAsrConf": 200
    }, {
      "utf8": " very",
      "tOffsetMs": 720,
      "acAsrConf": 197
    }, {
      "utf8": " strong",
      "tOffsetMs": 930,
      "acAsrConf": 220
    }, {
      "utf8": " feelings",
      "tOffsetMs": 1290,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1939,
      "acAsrConf": 236
    }, {
      "utf8": " I'm",
      "tOffsetMs": 2939,
      "acAsrConf": 247
    }, {
      "utf8": " glad",
      "tOffsetMs": 3210,
      "acAsrConf": 187
    } ]
  }, {
    "tStartMs": 2622580,
    "dDurationMs": 1960,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2622590,
    "dDurationMs": 4650,
    "wWinId": 1,
    "segs": [ {
      "utf8": "that",
      "acAsrConf": 140
    }, {
      "utf8": " you",
      "tOffsetMs": 180,
      "acAsrConf": 207
    }, {
      "utf8": " can",
      "tOffsetMs": 480,
      "acAsrConf": 227
    }, {
      "utf8": " understand",
      "tOffsetMs": 690,
      "acAsrConf": 252
    }, {
      "utf8": " my",
      "tOffsetMs": 1140,
      "acAsrConf": 252
    }, {
      "utf8": " accent",
      "tOffsetMs": 1290,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 1650,
      "acAsrConf": 250
    }, {
      "utf8": " hope",
      "tOffsetMs": 1800,
      "acAsrConf": 213
    } ]
  }, {
    "tStartMs": 2624530,
    "dDurationMs": 2710,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2624540,
    "dDurationMs": 6720,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 252
    }, {
      "utf8": " can",
      "tOffsetMs": 150,
      "acAsrConf": 252
    }, {
      "utf8": " understand",
      "tOffsetMs": 300,
      "acAsrConf": 252
    }, {
      "utf8": " damn",
      "tOffsetMs": 720,
      "acAsrConf": 225
    }, {
      "utf8": " -",
      "tOffsetMs": 930,
      "acAsrConf": 251
    }, {
      "utf8": " we",
      "tOffsetMs": 1200,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 1950,
      "acAsrConf": 252
    }, {
      "utf8": " very",
      "tOffsetMs": 2250,
      "acAsrConf": 246
    } ]
  }, {
    "tStartMs": 2627230,
    "dDurationMs": 4030,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2627240,
    "dDurationMs": 6660,
    "wWinId": 1,
    "segs": [ {
      "utf8": "Jin",
      "acAsrConf": 181
    }, {
      "utf8": " general",
      "tOffsetMs": 780,
      "acAsrConf": 252
    }, {
      "utf8": " American",
      "tOffsetMs": 1560,
      "acAsrConf": 241
    }, {
      "utf8": " accents",
      "tOffsetMs": 2220,
      "acAsrConf": 255
    }, {
      "utf8": " nothing",
      "tOffsetMs": 2760,
      "acAsrConf": 228
    }, {
      "utf8": " too",
      "tOffsetMs": 3720,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2631250,
    "dDurationMs": 2650,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2631260,
    "dDurationMs": 6210,
    "wWinId": 1,
    "segs": [ {
      "utf8": "strong",
      "acAsrConf": 252
    }, {
      "utf8": " nothing",
      "tOffsetMs": 300,
      "acAsrConf": 208
    }, {
      "utf8": " too",
      "tOffsetMs": 810,
      "acAsrConf": 252
    }, {
      "utf8": " specific",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " we",
      "tOffsetMs": 1310,
      "acAsrConf": 152
    }, {
      "utf8": " would",
      "tOffsetMs": 2310,
      "acAsrConf": 252
    }, {
      "utf8": " say",
      "tOffsetMs": 2460,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2633890,
    "dDurationMs": 3580,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2633900,
    "dDurationMs": 8240,
    "wWinId": 1,
    "segs": [ {
      "utf8": "um",
      "acAsrConf": 94
    }, {
      "utf8": " maybe",
      "tOffsetMs": 270,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 719,
      "acAsrConf": 252
    }, {
      "utf8": " Midwestern",
      "tOffsetMs": 870,
      "acAsrConf": 252
    }, {
      "utf8": " accent",
      "tOffsetMs": 1740,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 2670,
      "acAsrConf": 227
    }, {
      "utf8": " America",
      "tOffsetMs": 3000,
      "acAsrConf": 255
    } ]
  }, {
    "tStartMs": 2637460,
    "dDurationMs": 4680,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2637470,
    "dDurationMs": 8910,
    "wWinId": 1,
    "segs": [ {
      "utf8": "yeah",
      "acAsrConf": 228
    }, {
      "utf8": " like",
      "tOffsetMs": 300,
      "acAsrConf": 245
    }, {
      "utf8": " a",
      "tOffsetMs": 570,
      "acAsrConf": 252
    }, {
      "utf8": " standard",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 1139,
      "acAsrConf": 201
    }, {
      "utf8": " movie",
      "tOffsetMs": 1700,
      "acAsrConf": 237
    }, {
      "utf8": " or",
      "tOffsetMs": 2700,
      "acAsrConf": 252
    }, {
      "utf8": " news",
      "tOffsetMs": 3480,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2642130,
    "dDurationMs": 4250,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2642140,
    "dDurationMs": 7510,
    "wWinId": 1,
    "segs": [ {
      "utf8": "broadcast",
      "acAsrConf": 250
    }, {
      "utf8": " accent",
      "tOffsetMs": 1000,
      "acAsrConf": 202
    }, {
      "utf8": " nothing",
      "tOffsetMs": 1600,
      "acAsrConf": 139
    }, {
      "utf8": " too",
      "tOffsetMs": 1960,
      "acAsrConf": 252
    }, {
      "utf8": " strong",
      "tOffsetMs": 2200,
      "acAsrConf": 243
    }, {
      "utf8": " so",
      "tOffsetMs": 3100,
      "acAsrConf": 146
    }, {
      "utf8": " I",
      "tOffsetMs": 3460,
      "acAsrConf": 145
    } ]
  }, {
    "tStartMs": 2646370,
    "dDurationMs": 3280,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2646380,
    "dDurationMs": 5280,
    "wWinId": 1,
    "segs": [ {
      "utf8": "was",
      "acAsrConf": 252
    }, {
      "utf8": " born",
      "tOffsetMs": 690,
      "acAsrConf": 234
    }, {
      "utf8": " in",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    }, {
      "utf8": " Pittsburgh",
      "tOffsetMs": 1410,
      "acAsrConf": 167
    }, {
      "utf8": " Pennsylvania",
      "tOffsetMs": 2190,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 2489,
      "acAsrConf": 126
    } ]
  }, {
    "tStartMs": 2649640,
    "dDurationMs": 2020,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2649650,
    "dDurationMs": 5280,
    "wWinId": 1,
    "segs": [ {
      "utf8": "the",
      "acAsrConf": 202
    }, {
      "utf8": " north",
      "tOffsetMs": 360,
      "acAsrConf": 236
    } ]
  }, {
    "tStartMs": 2651650,
    "dDurationMs": 3280,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2651660,
    "dDurationMs": 3840,
    "wWinId": 1,
    "segs": [ {
      "utf8": "and",
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 90,
      "acAsrConf": 248
    }, {
      "utf8": " moved",
      "tOffsetMs": 630,
      "acAsrConf": 233
    }, {
      "utf8": " to",
      "tOffsetMs": 1380,
      "acAsrConf": 220
    }, {
      "utf8": " the",
      "tOffsetMs": 1950,
      "acAsrConf": 223
    }, {
      "utf8": " south",
      "tOffsetMs": 2070,
      "acAsrConf": 202
    }, {
      "utf8": " when",
      "tOffsetMs": 2340,
      "acAsrConf": 201
    }, {
      "utf8": " I",
      "tOffsetMs": 2970,
      "acAsrConf": 255
    }, {
      "utf8": " was",
      "tOffsetMs": 3000,
      "acAsrConf": 251
    } ]
  }, {
    "tStartMs": 2654920,
    "dDurationMs": 580,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2654930,
    "dDurationMs": 3030,
    "wWinId": 1,
    "segs": [ {
      "utf8": "really",
      "acAsrConf": 252
    }, {
      "utf8": " young",
      "tOffsetMs": 210,
      "acAsrConf": 202
    } ]
  }, {
    "tStartMs": 2655490,
    "dDurationMs": 2470,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2655500,
    "dDurationMs": 4830,
    "wWinId": 1,
    "segs": [ {
      "utf8": "so",
      "acAsrConf": 146
    }, {
      "utf8": " I",
      "tOffsetMs": 630,
      "acAsrConf": 239
    }, {
      "utf8": " don't",
      "tOffsetMs": 660,
      "acAsrConf": 226
    }, {
      "utf8": " really",
      "tOffsetMs": 1140,
      "acAsrConf": 255
    }, {
      "utf8": " have",
      "tOffsetMs": 1260,
      "acAsrConf": 200
    }, {
      "utf8": " a",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    }, {
      "utf8": " northern",
      "tOffsetMs": 1590,
      "acAsrConf": 252
    }, {
      "utf8": " accent",
      "tOffsetMs": 2100,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2657950,
    "dDurationMs": 2380,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2657960,
    "dDurationMs": 5160,
    "wWinId": 1,
    "segs": [ {
      "utf8": "or",
      "acAsrConf": 252
    }, {
      "utf8": " southern",
      "tOffsetMs": 210,
      "acAsrConf": 255
    }, {
      "utf8": " accent",
      "tOffsetMs": 510,
      "acAsrConf": 251
    }, {
      "utf8": " just",
      "tOffsetMs": 990,
      "acAsrConf": 216
    }, {
      "utf8": " from",
      "tOffsetMs": 1380,
      "acAsrConf": 249
    }, {
      "utf8": " anywhere",
      "tOffsetMs": 1920,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2660320,
    "dDurationMs": 2800,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2660330,
    "dDurationMs": 4170,
    "wWinId": 1,
    "segs": [ {
      "utf8": "what",
      "acAsrConf": 220
    }, {
      "utf8": " what",
      "tOffsetMs": 330,
      "acAsrConf": 248
    }, {
      "utf8": " about",
      "tOffsetMs": 510,
      "acAsrConf": 242
    }, {
      "utf8": " you",
      "tOffsetMs": 600,
      "acAsrConf": 187
    }, {
      "utf8": " and",
      "tOffsetMs": 870,
      "acAsrConf": 216
    }, {
      "utf8": " I",
      "tOffsetMs": 1170,
      "acAsrConf": 252
    }, {
      "utf8": " was",
      "tOffsetMs": 1440,
      "acAsrConf": 252
    }, {
      "utf8": " born",
      "tOffsetMs": 1590,
      "acAsrConf": 227
    }, {
      "utf8": " in",
      "tOffsetMs": 1920,
      "acAsrConf": 252
    }, {
      "utf8": " LA",
      "tOffsetMs": 2100,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2663110,
    "dDurationMs": 1390,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2663120,
    "dDurationMs": 4140,
    "wWinId": 1,
    "segs": [ {
      "utf8": "in",
      "acAsrConf": 222
    }, {
      "utf8": " California",
      "tOffsetMs": 260,
      "acAsrConf": 245
    } ]
  }, {
    "tStartMs": 2664490,
    "dDurationMs": 2770,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2664500,
    "dDurationMs": 5280,
    "wWinId": 1,
    "segs": [ {
      "utf8": "but",
      "acAsrConf": 236
    }, {
      "utf8": " I",
      "tOffsetMs": 150,
      "acAsrConf": 237
    }, {
      "utf8": " was",
      "tOffsetMs": 450,
      "acAsrConf": 252
    }, {
      "utf8": " five",
      "tOffsetMs": 810,
      "acAsrConf": 143
    }, {
      "utf8": " years",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    }, {
      "utf8": " old",
      "tOffsetMs": 1080,
      "acAsrConf": 252
    }, {
      "utf8": " when",
      "tOffsetMs": 1500,
      "acAsrConf": 210
    }, {
      "utf8": " I",
      "tOffsetMs": 1710,
      "acAsrConf": 167
    }, {
      "utf8": " left",
      "tOffsetMs": 1890,
      "acAsrConf": 158
    }, {
      "utf8": " so",
      "tOffsetMs": 2190,
      "acAsrConf": 203
    } ]
  }, {
    "tStartMs": 2667250,
    "dDurationMs": 2530,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2667260,
    "dDurationMs": 5370,
    "wWinId": 1,
    "segs": [ {
      "utf8": "I",
      "acAsrConf": 244
    }, {
      "utf8": " don't",
      "tOffsetMs": 30,
      "acAsrConf": 224
    }, {
      "utf8": " remember",
      "tOffsetMs": 480,
      "acAsrConf": 216
    }, {
      "utf8": " California",
      "tOffsetMs": 1130,
      "acAsrConf": 201
    }, {
      "utf8": " very",
      "tOffsetMs": 2130,
      "acAsrConf": 248
    }, {
      "utf8": " much",
      "tOffsetMs": 2310,
      "acAsrConf": 215
    } ]
  }, {
    "tStartMs": 2669770,
    "dDurationMs": 2860,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2669780,
    "dDurationMs": 5790,
    "wWinId": 1,
    "segs": [ {
      "utf8": "yes",
      "acAsrConf": 205
    }, {
      "utf8": " yeah",
      "tOffsetMs": 330,
      "acAsrConf": 198
    }, {
      "utf8": " I",
      "tOffsetMs": 870,
      "acAsrConf": 243
    }, {
      "utf8": " would",
      "tOffsetMs": 1140,
      "acAsrConf": 252
    }, {
      "utf8": " say",
      "tOffsetMs": 1440,
      "acAsrConf": 252
    }, {
      "utf8": " we",
      "tOffsetMs": 1470,
      "acAsrConf": 202
    }, {
      "utf8": " also",
      "tOffsetMs": 2010,
      "acAsrConf": 252
    }, {
      "utf8": " probably",
      "tOffsetMs": 2220,
      "acAsrConf": 221
    } ]
  }, {
    "tStartMs": 2672620,
    "dDurationMs": 2950,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2672630,
    "dDurationMs": 6480,
    "wWinId": 1,
    "segs": [ {
      "utf8": "sound",
      "acAsrConf": 218
    }, {
      "utf8": " very",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " clear",
      "tOffsetMs": 990,
      "acAsrConf": 252
    }, {
      "utf8": " because",
      "tOffsetMs": 1410,
      "acAsrConf": 151
    }, {
      "utf8": " we",
      "tOffsetMs": 1920,
      "acAsrConf": 200
    }, {
      "utf8": " are",
      "tOffsetMs": 2490,
      "acAsrConf": 252
    }, {
      "utf8": " trying",
      "tOffsetMs": 2640,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2675560,
    "dDurationMs": 3550,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2675570,
    "dDurationMs": 7410,
    "wWinId": 1,
    "segs": [ {
      "utf8": "to",
      "acAsrConf": 200
    }, {
      "utf8": " speak",
      "tOffsetMs": 210,
      "acAsrConf": 166
    }, {
      "utf8": " clearer",
      "tOffsetMs": 660,
      "acAsrConf": 109
    }, {
      "utf8": " now",
      "tOffsetMs": 1620,
      "acAsrConf": 251
    }, {
      "utf8": " so",
      "tOffsetMs": 2190,
      "acAsrConf": 252
    }, {
      "utf8": " there",
      "tOffsetMs": 2520,
      "acAsrConf": 237
    }, {
      "utf8": " are",
      "tOffsetMs": 3210,
      "acAsrConf": 252
    }, {
      "utf8": " many",
      "tOffsetMs": 3360,
      "acAsrConf": 221
    } ]
  }, {
    "tStartMs": 2679100,
    "dDurationMs": 3880,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2679110,
    "dDurationMs": 7200,
    "wWinId": 1,
    "segs": [ {
      "utf8": "search",
      "acAsrConf": 250
    }, {
      "utf8": " situations",
      "tOffsetMs": 650,
      "acAsrConf": 220
    }, {
      "utf8": " where",
      "tOffsetMs": 1650,
      "acAsrConf": 252
    }, {
      "utf8": " people",
      "tOffsetMs": 2010,
      "acAsrConf": 252
    }, {
      "utf8": " mumble",
      "tOffsetMs": 2540,
      "acAsrConf": 218
    }, {
      "utf8": " in",
      "tOffsetMs": 3540,
      "acAsrConf": 166
    } ]
  }, {
    "tStartMs": 2682970,
    "dDurationMs": 3340,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2682980,
    "dDurationMs": 7380,
    "wWinId": 1,
    "segs": [ {
      "utf8": "English",
      "acAsrConf": 243
    }, {
      "utf8": " so",
      "tOffsetMs": 540,
      "acAsrConf": 240
    }, {
      "utf8": " people",
      "tOffsetMs": 1220,
      "acAsrConf": 248
    }, {
      "utf8": " who",
      "tOffsetMs": 2220,
      "acAsrConf": 201
    }, {
      "utf8": " speak",
      "tOffsetMs": 2460,
      "acAsrConf": 252
    }, {
      "utf8": " English",
      "tOffsetMs": 2700,
      "acAsrConf": 255
    }, {
      "utf8": " are",
      "tOffsetMs": 2880,
      "acAsrConf": 202
    } ]
  }, {
    "tStartMs": 2686300,
    "dDurationMs": 4060,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2686310,
    "dDurationMs": 6930,
    "wWinId": 1,
    "segs": [ {
      "utf8": "very",
      "acAsrConf": 252
    }, {
      "utf8": " bad",
      "tOffsetMs": 750,
      "acAsrConf": 205
    }, {
      "utf8": " at",
      "tOffsetMs": 1590,
      "acAsrConf": 252
    }, {
      "utf8": " taking",
      "tOffsetMs": 1890,
      "acAsrConf": 252
    }, {
      "utf8": " words",
      "tOffsetMs": 2160,
      "acAsrConf": 216
    }, {
      "utf8": " and",
      "tOffsetMs": 2550,
      "acAsrConf": 235
    }, {
      "utf8": " smashing",
      "tOffsetMs": 3050,
      "acAsrConf": 243
    } ]
  }, {
    "tStartMs": 2690350,
    "dDurationMs": 2890,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2690360,
    "dDurationMs": 5550,
    "wWinId": 1,
    "segs": [ {
      "utf8": "them",
      "acAsrConf": 252
    }, {
      "utf8": " together",
      "tOffsetMs": 150,
      "acAsrConf": 215
    }, {
      "utf8": " so",
      "tOffsetMs": 420,
      "acAsrConf": 201
    }, {
      "utf8": " nobody",
      "tOffsetMs": 1350,
      "acAsrConf": 252
    }, {
      "utf8": " understands",
      "tOffsetMs": 1890,
      "acAsrConf": 246
    }, {
      "utf8": " them",
      "tOffsetMs": 2640,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2693230,
    "dDurationMs": 2680,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2693240,
    "dDurationMs": 5730,
    "wWinId": 1,
    "segs": [ {
      "utf8": "so",
      "acAsrConf": 201
    }, {
      "utf8": " yeah",
      "tOffsetMs": 360,
      "acAsrConf": 251
    }, {
      "utf8": " a",
      "tOffsetMs": 1170,
      "acAsrConf": 252
    }, {
      "utf8": " lot",
      "tOffsetMs": 1200,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 1530,
      "acAsrConf": 139
    }, {
      "utf8": " people",
      "tOffsetMs": 1740,
      "acAsrConf": 166
    }, {
      "utf8": " speak",
      "tOffsetMs": 2160,
      "acAsrConf": 252
    }, {
      "utf8": " very",
      "tOffsetMs": 2430,
      "acAsrConf": 109
    } ]
  }, {
    "tStartMs": 2695900,
    "dDurationMs": 3070,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2695910,
    "dDurationMs": 5820,
    "wWinId": 1,
    "segs": [ {
      "utf8": "quickly",
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 330,
      "acAsrConf": 236
    }, {
      "utf8": " take",
      "tOffsetMs": 1190,
      "acAsrConf": 255
    }, {
      "utf8": " words",
      "tOffsetMs": 2190,
      "acAsrConf": 238
    }, {
      "utf8": " and",
      "tOffsetMs": 2490,
      "acAsrConf": 252
    }, {
      "utf8": " put",
      "tOffsetMs": 2790,
      "acAsrConf": 252
    }, {
      "utf8": " them",
      "tOffsetMs": 2910,
      "acAsrConf": 188
    } ]
  }, {
    "tStartMs": 2698960,
    "dDurationMs": 2770,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2698970,
    "dDurationMs": 4890,
    "wWinId": 1,
    "segs": [ {
      "utf8": "together",
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 180,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 540,
      "acAsrConf": 141
    }, {
      "utf8": " can",
      "tOffsetMs": 810,
      "acAsrConf": 200
    }, {
      "utf8": " be",
      "tOffsetMs": 990,
      "acAsrConf": 252
    }, {
      "utf8": " confusing",
      "tOffsetMs": 1140,
      "acAsrConf": 252
    }, {
      "utf8": " but",
      "tOffsetMs": 1590,
      "acAsrConf": 160
    }, {
      "utf8": " I",
      "tOffsetMs": 2520,
      "acAsrConf": 244
    } ]
  }, {
    "tStartMs": 2701720,
    "dDurationMs": 2140,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2701730,
    "dDurationMs": 4110,
    "wWinId": 1,
    "segs": [ {
      "utf8": "know",
      "acAsrConf": 252
    }, {
      "utf8": " we",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " try",
      "tOffsetMs": 930,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1110,
      "acAsrConf": 252
    }, {
      "utf8": " make",
      "tOffsetMs": 1140,
      "acAsrConf": 216
    }, {
      "utf8": " it",
      "tOffsetMs": 1290,
      "acAsrConf": 94
    }, {
      "utf8": " a",
      "tOffsetMs": 1380,
      "acAsrConf": 150
    }, {
      "utf8": " little",
      "tOffsetMs": 1560,
      "acAsrConf": 228
    }, {
      "utf8": " easier",
      "tOffsetMs": 1710,
      "acAsrConf": 202
    } ]
  }, {
    "tStartMs": 2703850,
    "dDurationMs": 1990,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2703860,
    "dDurationMs": 3930,
    "wWinId": 1,
    "segs": [ {
      "utf8": "for",
      "acAsrConf": 246
    }, {
      "utf8": " this",
      "tOffsetMs": 270,
      "acAsrConf": 216
    }, {
      "utf8": " these",
      "tOffsetMs": 570,
      "acAsrConf": 238
    }, {
      "utf8": " kinds",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 1290,
      "acAsrConf": 252
    }, {
      "utf8": " videos",
      "tOffsetMs": 1320,
      "acAsrConf": 228
    }, {
      "utf8": " sure",
      "tOffsetMs": 1440,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 2705830,
    "dDurationMs": 1960,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2705840,
    "dDurationMs": 3930,
    "wWinId": 1,
    "segs": [ {
      "utf8": "having",
      "acAsrConf": 252
    }, {
      "utf8": " some",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " natural",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " conversation",
      "tOffsetMs": 1080,
      "acAsrConf": 252
    }, {
      "utf8": " but",
      "tOffsetMs": 1740,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2707780,
    "dDurationMs": 1990,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2707790,
    "dDurationMs": 4620,
    "wWinId": 1,
    "segs": [ {
      "utf8": "also",
      "acAsrConf": 177
    }, {
      "utf8": " having",
      "tOffsetMs": 210,
      "acAsrConf": 216
    }, {
      "utf8": " something",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " that",
      "tOffsetMs": 900,
      "acAsrConf": 170
    }, {
      "utf8": " you",
      "tOffsetMs": 1290,
      "acAsrConf": 252
    }, {
      "utf8": " can",
      "tOffsetMs": 1800,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2709760,
    "dDurationMs": 2650,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2709770,
    "dDurationMs": 5730,
    "wWinId": 1,
    "segs": [ {
      "utf8": "understand",
      "acAsrConf": 252
    }, {
      "utf8": " because",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " like",
      "tOffsetMs": 720,
      "acAsrConf": 202
    }, {
      "utf8": " I",
      "tOffsetMs": 1260,
      "acAsrConf": 252
    }, {
      "utf8": " said",
      "tOffsetMs": 1410,
      "acAsrConf": 252
    }, {
      "utf8": " about",
      "tOffsetMs": 1470,
      "acAsrConf": 219
    }, {
      "utf8": " 70%",
      "tOffsetMs": 1800,
      "acAsrConf": 165
    } ]
  }, {
    "tStartMs": 2712400,
    "dDurationMs": 3100,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2712410,
    "dDurationMs": 5610,
    "wWinId": 1,
    "segs": [ {
      "utf8": "I",
      "acAsrConf": 220
    }, {
      "utf8": " want",
      "tOffsetMs": 150,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 390,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 480,
      "acAsrConf": 248
    }, {
      "utf8": " understand",
      "tOffsetMs": 630,
      "acAsrConf": 252
    }, {
      "utf8": " at",
      "tOffsetMs": 1110,
      "acAsrConf": 252
    }, {
      "utf8": " least",
      "tOffsetMs": 1290,
      "acAsrConf": 207
    }, {
      "utf8": " 70%",
      "tOffsetMs": 1530,
      "acAsrConf": 236
    }, {
      "utf8": " so",
      "tOffsetMs": 2130,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 2715490,
    "dDurationMs": 2530,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2715500,
    "dDurationMs": 3540,
    "wWinId": 1,
    "segs": [ {
      "utf8": "that",
      "acAsrConf": 255
    }, {
      "utf8": " you",
      "tOffsetMs": 210,
      "acAsrConf": 252
    }, {
      "utf8": " can",
      "tOffsetMs": 330,
      "acAsrConf": 252
    }, {
      "utf8": " improve",
      "tOffsetMs": 540,
      "acAsrConf": 202
    }, {
      "utf8": " that",
      "tOffsetMs": 870,
      "acAsrConf": 142
    }, {
      "utf8": " 30%",
      "tOffsetMs": 1260,
      "acAsrConf": 163
    }, {
      "utf8": " more",
      "tOffsetMs": 2070,
      "acAsrConf": 252
    }, {
      "utf8": " that",
      "tOffsetMs": 2280,
      "acAsrConf": 242
    } ]
  }, {
    "tStartMs": 2718010,
    "dDurationMs": 1030,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2718020,
    "dDurationMs": 3510,
    "wWinId": 1,
    "segs": [ {
      "utf8": "will",
      "acAsrConf": 252
    }, {
      "utf8": " be",
      "tOffsetMs": 120,
      "acAsrConf": 201
    }, {
      "utf8": " really",
      "tOffsetMs": 240,
      "acAsrConf": 237
    }, {
      "utf8": " helpful",
      "tOffsetMs": 540,
      "acAsrConf": 165
    } ]
  }, {
    "tStartMs": 2719030,
    "dDurationMs": 2500,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2719040,
    "dDurationMs": 4440,
    "wWinId": 1,
    "segs": [ {
      "utf8": "hmm",
      "acAsrConf": 165
    }, {
      "utf8": " so",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " all",
      "tOffsetMs": 900,
      "acAsrConf": 232
    }, {
      "utf8": " right",
      "tOffsetMs": 1440,
      "acAsrConf": 186
    }, {
      "utf8": " let's",
      "tOffsetMs": 1890,
      "acAsrConf": 166
    }, {
      "utf8": " see",
      "tOffsetMs": 2070,
      "acAsrConf": 252
    }, {
      "utf8": " if",
      "tOffsetMs": 2190,
      "acAsrConf": 243
    }, {
      "utf8": " we",
      "tOffsetMs": 2310,
      "acAsrConf": 167
    }, {
      "utf8": " have",
      "tOffsetMs": 2460,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2721520,
    "dDurationMs": 1960,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2721530,
    "dDurationMs": 4560,
    "wWinId": 1,
    "segs": [ {
      "utf8": "any",
      "acAsrConf": 201
    }, {
      "utf8": " other",
      "tOffsetMs": 300,
      "acAsrConf": 252
    }, {
      "utf8": " questions",
      "tOffsetMs": 390,
      "acAsrConf": 155
    }, {
      "utf8": " and",
      "tOffsetMs": 600,
      "acAsrConf": 106
    }, {
      "utf8": " then",
      "tOffsetMs": 1170,
      "acAsrConf": 226
    }, {
      "utf8": " I",
      "tOffsetMs": 1470,
      "acAsrConf": 246
    }, {
      "utf8": " think",
      "tOffsetMs": 1500,
      "acAsrConf": 244
    }, {
      "utf8": " it",
      "tOffsetMs": 1830,
      "acAsrConf": 232
    } ]
  }, {
    "tStartMs": 2723470,
    "dDurationMs": 2620,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2723480,
    "dDurationMs": 4560,
    "wWinId": 1,
    "segs": [ {
      "utf8": "might",
      "acAsrConf": 252
    }, {
      "utf8": " be",
      "tOffsetMs": 30,
      "acAsrConf": 88
    }, {
      "utf8": " time",
      "tOffsetMs": 180,
      "acAsrConf": 217
    }, {
      "utf8": " for",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " us",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 720,
      "acAsrConf": 227
    }, {
      "utf8": " go",
      "tOffsetMs": 1020,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1200,
      "acAsrConf": 233
    }, {
      "utf8": " Lizzy",
      "tOffsetMs": 1830,
      "acAsrConf": 77
    }, {
      "utf8": " if",
      "tOffsetMs": 2430,
      "acAsrConf": 249
    } ]
  }, {
    "tStartMs": 2726080,
    "dDurationMs": 1960,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2726090,
    "dDurationMs": 10160,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 215
    }, {
      "utf8": " have",
      "tOffsetMs": 150,
      "acAsrConf": 252
    }, {
      "utf8": " any",
      "tOffsetMs": 300,
      "acAsrConf": 252
    }, {
      "utf8": " final",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " questions",
      "tOffsetMs": 840,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1110,
      "acAsrConf": 252
    }, {
      "utf8": " can",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    }, {
      "utf8": " ask",
      "tOffsetMs": 1740,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2728030,
    "dDurationMs": 8220,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2728040,
    "dDurationMs": 8210,
    "wWinId": 1,
    "segs": [ {
      "utf8": "them",
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 330,
      "acAsrConf": 211
    }, {
      "utf8": " the",
      "tOffsetMs": 840,
      "acAsrConf": 248
    }, {
      "utf8": " question",
      "tOffsetMs": 2450,
      "acAsrConf": 224
    }, {
      "utf8": " question",
      "tOffsetMs": 3650,
      "acAsrConf": 245
    }, {
      "utf8": " box",
      "tOffsetMs": 4650,
      "acAsrConf": 230
    }, {
      "utf8": " and",
      "tOffsetMs": 5210,
      "acAsrConf": 236
    } ]
  }, {
    "tStartMs": 2736300,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2736310,
    "dDurationMs": 5650,
    "wWinId": 1,
    "segs": [ {
      "utf8": "what",
      "acAsrConf": 69
    }, {
      "utf8": " about",
      "tOffsetMs": 1000,
      "acAsrConf": 217
    }, {
      "utf8": " vocabulary",
      "tOffsetMs": 1210,
      "acAsrConf": 252
    }, {
      "utf8": " okay",
      "tOffsetMs": 1560,
      "acAsrConf": 145
    }, {
      "utf8": " we",
      "tOffsetMs": 2560,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 2740,
      "acAsrConf": 235
    }, {
      "utf8": " a",
      "tOffsetMs": 2830,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2739160,
    "dDurationMs": 2800,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2739170,
    "dDurationMs": 3780,
    "wWinId": 1,
    "segs": [ {
      "utf8": "question",
      "acAsrConf": 216
    }, {
      "utf8": " about",
      "tOffsetMs": 210,
      "acAsrConf": 207
    }, {
      "utf8": " how",
      "tOffsetMs": 360,
      "acAsrConf": 204
    }, {
      "utf8": " to",
      "tOffsetMs": 600,
      "acAsrConf": 178
    }, {
      "utf8": " learn",
      "tOffsetMs": 660,
      "acAsrConf": 252
    }, {
      "utf8": " how",
      "tOffsetMs": 1640,
      "acAsrConf": 181
    }, {
      "utf8": " do",
      "tOffsetMs": 2640,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 2700,
      "acAsrConf": 247
    } ]
  }, {
    "tStartMs": 2741950,
    "dDurationMs": 1000,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2741960,
    "dDurationMs": 5970,
    "wWinId": 1,
    "segs": [ {
      "utf8": "remember",
      "acAsrConf": 250
    }, {
      "utf8": " it",
      "tOffsetMs": 360,
      "acAsrConf": 218
    }, {
      "utf8": " okay",
      "tOffsetMs": 600,
      "acAsrConf": 165
    } ]
  }, {
    "tStartMs": 2742940,
    "dDurationMs": 4990,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2742950,
    "dDurationMs": 8850,
    "wWinId": 1,
    "segs": [ {
      "utf8": "I",
      "acAsrConf": 60
    }, {
      "utf8": " think",
      "tOffsetMs": 600,
      "acAsrConf": 246
    }, {
      "utf8": " it",
      "tOffsetMs": 1050,
      "acAsrConf": 231
    }, {
      "utf8": " was",
      "tOffsetMs": 1140,
      "acAsrConf": 232
    }, {
      "utf8": " uh-huh",
      "tOffsetMs": 2540,
      "acAsrConf": 203
    }, {
      "utf8": " all",
      "tOffsetMs": 3570,
      "acAsrConf": 203
    }, {
      "utf8": " right",
      "tOffsetMs": 3960,
      "acAsrConf": 146
    }, {
      "utf8": " Amir",
      "tOffsetMs": 4260,
      "acAsrConf": 237
    } ]
  }, {
    "tStartMs": 2747920,
    "dDurationMs": 3880,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2747930,
    "dDurationMs": 6480,
    "wWinId": 1,
    "segs": [ {
      "utf8": "Amir",
      "acAsrConf": 248
    }, {
      "utf8": " Wadi",
      "tOffsetMs": 420,
      "acAsrConf": 240
    }, {
      "utf8": " yes",
      "tOffsetMs": 1080,
      "acAsrConf": 205
    }, {
      "utf8": " Amir",
      "tOffsetMs": 1650,
      "acAsrConf": 235
    }, {
      "utf8": " Amir",
      "tOffsetMs": 2220,
      "acAsrConf": 235
    }, {
      "utf8": " says",
      "tOffsetMs": 2550,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 3000,
      "acAsrConf": 206
    }, {
      "utf8": " what",
      "tOffsetMs": 3360,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 3720,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2751790,
    "dDurationMs": 2620,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2751800,
    "dDurationMs": 6510,
    "wWinId": 1,
    "segs": [ {
      "utf8": "the",
      "acAsrConf": 252
    }, {
      "utf8": " best",
      "tOffsetMs": 120,
      "acAsrConf": 252
    }, {
      "utf8": " way",
      "tOffsetMs": 300,
      "acAsrConf": 206
    }, {
      "utf8": " to",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " remember",
      "tOffsetMs": 510,
      "acAsrConf": 240
    }, {
      "utf8": " vocabulary",
      "tOffsetMs": 870,
      "acAsrConf": 207
    }, {
      "utf8": " well",
      "tOffsetMs": 1610,
      "acAsrConf": 144
    } ]
  }, {
    "tStartMs": 2754400,
    "dDurationMs": 3910,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2754410,
    "dDurationMs": 6810,
    "wWinId": 1,
    "segs": [ {
      "utf8": "in",
      "acAsrConf": 250
    }, {
      "utf8": " my",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " opinion",
      "tOffsetMs": 1440,
      "acAsrConf": 167
    }, {
      "utf8": " there",
      "tOffsetMs": 1830,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 2100,
      "acAsrConf": 252
    }, {
      "utf8": " not",
      "tOffsetMs": 2130,
      "acAsrConf": 202
    }, {
      "utf8": " one",
      "tOffsetMs": 2550,
      "acAsrConf": 252
    }, {
      "utf8": " perfect",
      "tOffsetMs": 3120,
      "acAsrConf": 187
    } ]
  }, {
    "tStartMs": 2758300,
    "dDurationMs": 2920,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2758310,
    "dDurationMs": 6180,
    "wWinId": 1,
    "segs": [ {
      "utf8": "way",
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 180,
      "acAsrConf": 147
    }, {
      "utf8": " remember",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " vocabulary",
      "tOffsetMs": 1020,
      "acAsrConf": 217
    }, {
      "utf8": " because",
      "tOffsetMs": 1550,
      "acAsrConf": 200
    }, {
      "utf8": " it's",
      "tOffsetMs": 2550,
      "acAsrConf": 169
    } ]
  }, {
    "tStartMs": 2761210,
    "dDurationMs": 3280,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2761220,
    "dDurationMs": 5100,
    "wWinId": 1,
    "segs": [ {
      "utf8": "going",
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 240,
      "acAsrConf": 224
    }, {
      "utf8": " depend",
      "tOffsetMs": 330,
      "acAsrConf": 204
    }, {
      "utf8": " your",
      "tOffsetMs": 950,
      "acAsrConf": 231
    }, {
      "utf8": " learning",
      "tOffsetMs": 1950,
      "acAsrConf": 234
    }, {
      "utf8": " style",
      "tOffsetMs": 2550,
      "acAsrConf": 252
    }, {
      "utf8": " but",
      "tOffsetMs": 2910,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2764480,
    "dDurationMs": 1840,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2764490,
    "dDurationMs": 4080,
    "wWinId": 1,
    "segs": [ {
      "utf8": "there",
      "acAsrConf": 252
    }, {
      "utf8": " are",
      "tOffsetMs": 180,
      "acAsrConf": 252
    }, {
      "utf8": " definitely",
      "tOffsetMs": 330,
      "acAsrConf": 252
    }, {
      "utf8": " ways",
      "tOffsetMs": 870,
      "acAsrConf": 252
    }, {
      "utf8": " that",
      "tOffsetMs": 1140,
      "acAsrConf": 238
    }, {
      "utf8": " don't",
      "tOffsetMs": 1440,
      "acAsrConf": 208
    } ]
  }, {
    "tStartMs": 2766310,
    "dDurationMs": 2260,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2766320,
    "dDurationMs": 5790,
    "wWinId": 1,
    "segs": [ {
      "utf8": "work",
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 240,
      "acAsrConf": 187
    }, {
      "utf8": " like",
      "tOffsetMs": 690,
      "acAsrConf": 252
    }, {
      "utf8": " we",
      "tOffsetMs": 1230,
      "acAsrConf": 252
    }, {
      "utf8": " said",
      "tOffsetMs": 1410,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 1590,
      "acAsrConf": 196
    }, {
      "utf8": " don't",
      "tOffsetMs": 1650,
      "acAsrConf": 211
    }, {
      "utf8": " recommend",
      "tOffsetMs": 1980,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2768560,
    "dDurationMs": 3550,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2768570,
    "dDurationMs": 6390,
    "wWinId": 1,
    "segs": [ {
      "utf8": "memorizing",
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 420,
      "acAsrConf": 252
    }, {
      "utf8": " list",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " but",
      "tOffsetMs": 1250,
      "acAsrConf": 236
    }, {
      "utf8": " one",
      "tOffsetMs": 2250,
      "acAsrConf": 245
    }, {
      "utf8": " of",
      "tOffsetMs": 2850,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 2880,
      "acAsrConf": 207
    }, {
      "utf8": " better",
      "tOffsetMs": 3210,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2772100,
    "dDurationMs": 2860,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2772110,
    "dDurationMs": 7650,
    "wWinId": 1,
    "segs": [ {
      "utf8": "ways",
      "acAsrConf": 223
    }, {
      "utf8": " to",
      "tOffsetMs": 390,
      "acAsrConf": 252
    }, {
      "utf8": " learn",
      "tOffsetMs": 420,
      "acAsrConf": 200
    }, {
      "utf8": " vocabulary",
      "tOffsetMs": 900,
      "acAsrConf": 252
    }, {
      "utf8": " is",
      "tOffsetMs": 1140,
      "acAsrConf": 217
    }, {
      "utf8": " to",
      "tOffsetMs": 1980,
      "acAsrConf": 252
    }, {
      "utf8": " learn",
      "tOffsetMs": 2550,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2774950,
    "dDurationMs": 4810,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2774960,
    "dDurationMs": 7320,
    "wWinId": 1,
    "segs": [ {
      "utf8": "phrases",
      "acAsrConf": 252
    }, {
      "utf8": " or",
      "tOffsetMs": 900,
      "acAsrConf": 231
    }, {
      "utf8": " learn",
      "tOffsetMs": 1530,
      "acAsrConf": 252
    }, {
      "utf8": " words",
      "tOffsetMs": 2070,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 2520,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 2880,
      "acAsrConf": 252
    }, {
      "utf8": " story",
      "tOffsetMs": 3480,
      "acAsrConf": 252
    }, {
      "utf8": " or",
      "tOffsetMs": 4110,
      "acAsrConf": 208
    }, {
      "utf8": " in",
      "tOffsetMs": 4560,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2779750,
    "dDurationMs": 2530,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2779760,
    "dDurationMs": 4260,
    "wWinId": 1,
    "segs": [ {
      "utf8": "a",
      "acAsrConf": 252
    }, {
      "utf8": " conversation",
      "tOffsetMs": 90,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 470,
      "acAsrConf": 165
    }, {
      "utf8": " for",
      "tOffsetMs": 1470,
      "acAsrConf": 252
    }, {
      "utf8": " example",
      "tOffsetMs": 1800,
      "acAsrConf": 226
    }, {
      "utf8": " in",
      "tOffsetMs": 2160,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 2310,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2782270,
    "dDurationMs": 1750,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2782280,
    "dDurationMs": 3360,
    "wWinId": 1,
    "segs": [ {
      "utf8": "course",
      "acAsrConf": 229
    }, {
      "utf8": " we",
      "tOffsetMs": 270,
      "acAsrConf": 252
    }, {
      "utf8": " talked",
      "tOffsetMs": 420,
      "acAsrConf": 252
    }, {
      "utf8": " about",
      "tOffsetMs": 630,
      "acAsrConf": 211
    }, {
      "utf8": " in",
      "tOffsetMs": 780,
      "acAsrConf": 226
    }, {
      "utf8": " the",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    }, {
      "utf8": " 30",
      "tOffsetMs": 1170,
      "acAsrConf": 252
    }, {
      "utf8": " days",
      "tOffsetMs": 1380,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 1470,
      "acAsrConf": 144
    } ]
  }, {
    "tStartMs": 2784010,
    "dDurationMs": 1630,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2784020,
    "dDurationMs": 3900,
    "wWinId": 1,
    "segs": [ {
      "utf8": "English",
      "acAsrConf": 204
    }, {
      "utf8": " course",
      "tOffsetMs": 300,
      "acAsrConf": 252
    }, {
      "utf8": " then",
      "tOffsetMs": 630,
      "acAsrConf": 186
    }, {
      "utf8": " a",
      "tOffsetMs": 1200,
      "acAsrConf": 204
    }, {
      "utf8": " nine",
      "tOffsetMs": 1230,
      "acAsrConf": 209
    } ]
  }, {
    "tStartMs": 2785630,
    "dDurationMs": 2290,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2785640,
    "dDurationMs": 4860,
    "wWinId": 1,
    "segs": [ {
      "utf8": "talk",
      "acAsrConf": 221
    }, {
      "utf8": " about",
      "tOffsetMs": 120,
      "acAsrConf": 250
    }, {
      "utf8": " a",
      "tOffsetMs": 300,
      "acAsrConf": 207
    }, {
      "utf8": " specific",
      "tOffsetMs": 810,
      "acAsrConf": 243
    }, {
      "utf8": " topic",
      "tOffsetMs": 1139,
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 1679,
      "acAsrConf": 231
    }, {
      "utf8": " each",
      "tOffsetMs": 2189,
      "acAsrConf": 236
    } ]
  }, {
    "tStartMs": 2787910,
    "dDurationMs": 2590,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2787920,
    "dDurationMs": 4340,
    "wWinId": 1,
    "segs": [ {
      "utf8": "video",
      "acAsrConf": 187
    }, {
      "utf8": " so",
      "tOffsetMs": 419,
      "acAsrConf": 126
    }, {
      "utf8": " one",
      "tOffsetMs": 659,
      "acAsrConf": 216
    }, {
      "utf8": " topic",
      "tOffsetMs": 1080,
      "acAsrConf": 252
    }, {
      "utf8": " we",
      "tOffsetMs": 1500,
      "acAsrConf": 214
    }, {
      "utf8": " talked",
      "tOffsetMs": 1649,
      "acAsrConf": 237
    }, {
      "utf8": " about",
      "tOffsetMs": 1830,
      "acAsrConf": 252
    }, {
      "utf8": " like",
      "tOffsetMs": 1980,
      "acAsrConf": 200
    } ]
  }, {
    "tStartMs": 2790490,
    "dDurationMs": 1770,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2790500,
    "dDurationMs": 4740,
    "wWinId": 1,
    "segs": [ {
      "utf8": "friendship",
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 690,
      "acAsrConf": 218
    }, {
      "utf8": " we",
      "tOffsetMs": 900,
      "acAsrConf": 252
    }, {
      "utf8": " use",
      "tOffsetMs": 1019,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 1230,
      "acAsrConf": 255
    }, {
      "utf8": " lot",
      "tOffsetMs": 1260,
      "acAsrConf": 255
    }, {
      "utf8": " of",
      "tOffsetMs": 1559,
      "acAsrConf": 141
    } ]
  }, {
    "tStartMs": 2792250,
    "dDurationMs": 2990,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2792260,
    "dDurationMs": 8650,
    "wWinId": 1,
    "segs": [ {
      "utf8": "friendship",
      "acAsrConf": 243
    }, {
      "utf8": " words",
      "tOffsetMs": 1000,
      "acAsrConf": 243
    }, {
      "utf8": " in",
      "tOffsetMs": 1420,
      "acAsrConf": 230
    }, {
      "utf8": " the",
      "tOffsetMs": 2019,
      "acAsrConf": 252
    }, {
      "utf8": " video",
      "tOffsetMs": 2349,
      "acAsrConf": 252
    }, {
      "utf8": " like",
      "tOffsetMs": 2470,
      "acAsrConf": 209
    } ]
  }, {
    "tStartMs": 2795230,
    "dDurationMs": 5680,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2795240,
    "dDurationMs": 7980,
    "wWinId": 1,
    "segs": [ {
      "utf8": "respect",
      "acAsrConf": 240
    }, {
      "utf8": " a",
      "tOffsetMs": 990,
      "acAsrConf": 232
    }, {
      "utf8": " good",
      "tOffsetMs": 1230,
      "acAsrConf": 252
    }, {
      "utf8": " listener",
      "tOffsetMs": 1589,
      "acAsrConf": 236
    }, {
      "utf8": " kind",
      "tOffsetMs": 2220,
      "acAsrConf": 218
    }, {
      "utf8": " gentle",
      "tOffsetMs": 3079,
      "acAsrConf": 252
    }, {
      "utf8": " same",
      "tOffsetMs": 4670,
      "acAsrConf": 228
    } ]
  }, {
    "tStartMs": 2800900,
    "dDurationMs": 2320,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2800910,
    "dDurationMs": 4080,
    "wWinId": 1,
    "segs": [ {
      "utf8": "interest",
      "acAsrConf": 240
    }, {
      "utf8": " we",
      "tOffsetMs": 689,
      "acAsrConf": 239
    }, {
      "utf8": " use",
      "tOffsetMs": 840,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 990,
      "acAsrConf": 252
    }, {
      "utf8": " lot",
      "tOffsetMs": 1020,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 1320,
      "acAsrConf": 252
    }, {
      "utf8": " words",
      "tOffsetMs": 1350,
      "acAsrConf": 235
    }, {
      "utf8": " about",
      "tOffsetMs": 1649,
      "acAsrConf": 186
    } ]
  }, {
    "tStartMs": 2803210,
    "dDurationMs": 1780,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2803220,
    "dDurationMs": 5760,
    "wWinId": 1,
    "segs": [ {
      "utf8": "that",
      "acAsrConf": 144
    }, {
      "utf8": " topic",
      "tOffsetMs": 240,
      "acAsrConf": 184
    }, {
      "utf8": " so",
      "tOffsetMs": 690,
      "acAsrConf": 252
    }, {
      "utf8": " that's",
      "tOffsetMs": 720,
      "acAsrConf": 201
    }, {
      "utf8": " a",
      "tOffsetMs": 1049,
      "acAsrConf": 252
    }, {
      "utf8": " great",
      "tOffsetMs": 1170,
      "acAsrConf": 252
    }, {
      "utf8": " way",
      "tOffsetMs": 1500,
      "acAsrConf": 165
    }, {
      "utf8": " to",
      "tOffsetMs": 1559,
      "acAsrConf": 0
    } ]
  }, {
    "tStartMs": 2804980,
    "dDurationMs": 4000,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2804990,
    "dDurationMs": 6329,
    "wWinId": 1,
    "segs": [ {
      "utf8": "learn",
      "acAsrConf": 252
    }, {
      "utf8": " vocabulary",
      "tOffsetMs": 299,
      "acAsrConf": 227
    }, {
      "utf8": " is",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 1260,
      "acAsrConf": 255
    }, {
      "utf8": " context",
      "tOffsetMs": 1579,
      "acAsrConf": 245
    }, {
      "utf8": " so",
      "tOffsetMs": 2579,
      "acAsrConf": 176
    }, {
      "utf8": " don't",
      "tOffsetMs": 3270,
      "acAsrConf": 240
    } ]
  }, {
    "tStartMs": 2808970,
    "dDurationMs": 2349,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2808980,
    "dDurationMs": 5900,
    "wWinId": 1,
    "segs": [ {
      "utf8": "study",
      "acAsrConf": 252
    }, {
      "utf8": " vocabulary",
      "tOffsetMs": 420,
      "acAsrConf": 201
    }, {
      "utf8": " words",
      "tOffsetMs": 1170,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 1530,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 1740,
      "acAsrConf": 207
    }, {
      "utf8": " list",
      "tOffsetMs": 1829,
      "acAsrConf": 252
    }, {
      "utf8": " or",
      "tOffsetMs": 2129,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2811309,
    "dDurationMs": 3571,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2811319,
    "dDurationMs": 6540,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you'll",
      "acAsrConf": 217
    }, {
      "utf8": " probably",
      "tOffsetMs": 181,
      "acAsrConf": 252
    }, {
      "utf8": " be",
      "tOffsetMs": 450,
      "acAsrConf": 252
    }, {
      "utf8": " very",
      "tOffsetMs": 931,
      "acAsrConf": 252
    }, {
      "utf8": " bored",
      "tOffsetMs": 1081,
      "acAsrConf": 215
    }, {
      "utf8": " it",
      "tOffsetMs": 1970,
      "acAsrConf": 106
    }, {
      "utf8": " is",
      "tOffsetMs": 2970,
      "acAsrConf": 252
    }, {
      "utf8": " also",
      "tOffsetMs": 3121,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2814870,
    "dDurationMs": 2989,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2814880,
    "dDurationMs": 6219,
    "wWinId": 1,
    "segs": [ {
      "utf8": "difficult",
      "acAsrConf": 214
    }, {
      "utf8": " to",
      "tOffsetMs": 1000,
      "acAsrConf": 252
    }, {
      "utf8": " remember",
      "tOffsetMs": 1270,
      "acAsrConf": 167
    }, {
      "utf8": " vocabulary",
      "tOffsetMs": 1780,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 1929,
      "acAsrConf": 224
    }, {
      "utf8": " think",
      "tOffsetMs": 2590,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2817849,
    "dDurationMs": 3250,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2817859,
    "dDurationMs": 5551,
    "wWinId": 1,
    "segs": [ {
      "utf8": "if",
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 180,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 420,
      "acAsrConf": 255
    }, {
      "utf8": " too",
      "tOffsetMs": 690,
      "acAsrConf": 248
    }, {
      "utf8": " much",
      "tOffsetMs": 1291,
      "acAsrConf": 252
    }, {
      "utf8": " pressure",
      "tOffsetMs": 1470,
      "acAsrConf": 226
    }, {
      "utf8": " you",
      "tOffsetMs": 2091,
      "acAsrConf": 137
    }, {
      "utf8": " have",
      "tOffsetMs": 3091,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 2821089,
    "dDurationMs": 2321,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2821099,
    "dDurationMs": 5161,
    "wWinId": 1,
    "segs": [ {
      "utf8": "to",
      "acAsrConf": 228
    }, {
      "utf8": " maybe",
      "tOffsetMs": 180,
      "acAsrConf": 251
    }, {
      "utf8": " feel",
      "tOffsetMs": 361,
      "acAsrConf": 245
    }, {
      "utf8": " a",
      "tOffsetMs": 811,
      "acAsrConf": 252
    }, {
      "utf8": " little",
      "tOffsetMs": 841,
      "acAsrConf": 250
    }, {
      "utf8": " more",
      "tOffsetMs": 1111,
      "acAsrConf": 217
    }, {
      "utf8": " relaxed",
      "tOffsetMs": 1440,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2823400,
    "dDurationMs": 2860,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2823410,
    "dDurationMs": 5159,
    "wWinId": 1,
    "segs": [ {
      "utf8": "there's",
      "acAsrConf": 164
    }, {
      "utf8": " not",
      "tOffsetMs": 990,
      "acAsrConf": 215
    }, {
      "utf8": " a",
      "tOffsetMs": 1290,
      "acAsrConf": 252
    }, {
      "utf8": " big",
      "tOffsetMs": 1320,
      "acAsrConf": 252
    }, {
      "utf8": " test",
      "tOffsetMs": 1679,
      "acAsrConf": 246
    }, {
      "utf8": " you",
      "tOffsetMs": 2159,
      "acAsrConf": 226
    }, {
      "utf8": " have",
      "tOffsetMs": 2429,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 2550,
      "acAsrConf": 186
    }, {
      "utf8": " study",
      "tOffsetMs": 2669,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2826250,
    "dDurationMs": 2319,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2826260,
    "dDurationMs": 4710,
    "wWinId": 1,
    "segs": [ {
      "utf8": "for",
      "acAsrConf": 229
    }, {
      "utf8": " and",
      "tOffsetMs": 120,
      "acAsrConf": 200
    }, {
      "utf8": " you're",
      "tOffsetMs": 509,
      "acAsrConf": 175
    }, {
      "utf8": " you",
      "tOffsetMs": 630,
      "acAsrConf": 239
    }, {
      "utf8": " are",
      "tOffsetMs": 1230,
      "acAsrConf": 200
    }, {
      "utf8": " studying",
      "tOffsetMs": 1380,
      "acAsrConf": 252
    }, {
      "utf8": " all",
      "tOffsetMs": 1799,
      "acAsrConf": 212
    }, {
      "utf8": " the",
      "tOffsetMs": 2190,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2828559,
    "dDurationMs": 2411,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2828569,
    "dDurationMs": 5040,
    "wWinId": 1,
    "segs": [ {
      "utf8": "vocabulary",
      "acAsrConf": 252
    }, {
      "utf8": " where",
      "tOffsetMs": 270,
      "acAsrConf": 255
    }, {
      "utf8": " it's",
      "tOffsetMs": 811,
      "acAsrConf": 222
    }, {
      "utf8": " you",
      "tOffsetMs": 960,
      "acAsrConf": 222
    }, {
      "utf8": " know",
      "tOffsetMs": 1621,
      "acAsrConf": 204
    }, {
      "utf8": " you",
      "tOffsetMs": 1770,
      "acAsrConf": 222
    }, {
      "utf8": " have",
      "tOffsetMs": 2161,
      "acAsrConf": 233
    } ]
  }, {
    "tStartMs": 2830960,
    "dDurationMs": 2649,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2830970,
    "dDurationMs": 5639,
    "wWinId": 1,
    "segs": [ {
      "utf8": "to",
      "acAsrConf": 252
    }, {
      "utf8": " really",
      "tOffsetMs": 210,
      "acAsrConf": 204
    }, {
      "utf8": " be",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " wanting",
      "tOffsetMs": 1079,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1859,
      "acAsrConf": 252
    }, {
      "utf8": " learn",
      "tOffsetMs": 1889,
      "acAsrConf": 156
    }, {
      "utf8": " the",
      "tOffsetMs": 2399,
      "acAsrConf": 225
    } ]
  }, {
    "tStartMs": 2833599,
    "dDurationMs": 3010,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2833609,
    "dDurationMs": 5521,
    "wWinId": 1,
    "segs": [ {
      "utf8": "vocabulary",
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 121,
      "acAsrConf": 229
    }, {
      "utf8": " having",
      "tOffsetMs": 1010,
      "acAsrConf": 248
    }, {
      "utf8": " it",
      "tOffsetMs": 2010,
      "acAsrConf": 252
    }, {
      "utf8": " be",
      "tOffsetMs": 2190,
      "acAsrConf": 252
    }, {
      "utf8": " part",
      "tOffsetMs": 2460,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 2490,
      "acAsrConf": 144
    }, {
      "utf8": " your",
      "tOffsetMs": 2851,
      "acAsrConf": 239
    } ]
  }, {
    "tStartMs": 2836599,
    "dDurationMs": 2531,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2836609,
    "dDurationMs": 4801,
    "wWinId": 1,
    "segs": [ {
      "utf8": "life",
      "acAsrConf": 252
    }, {
      "utf8": " part",
      "tOffsetMs": 180,
      "acAsrConf": 201
    }, {
      "utf8": " of",
      "tOffsetMs": 960,
      "acAsrConf": 208
    }, {
      "utf8": " your",
      "tOffsetMs": 1081,
      "acAsrConf": 221
    }, {
      "utf8": " habit",
      "tOffsetMs": 1230,
      "acAsrConf": 231
    }, {
      "utf8": " you",
      "tOffsetMs": 1440,
      "acAsrConf": 133
    }, {
      "utf8": " know",
      "tOffsetMs": 2160,
      "acAsrConf": 252
    }, {
      "utf8": " that's",
      "tOffsetMs": 2281,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2839120,
    "dDurationMs": 2290,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2839130,
    "dDurationMs": 5370,
    "wWinId": 1,
    "segs": [ {
      "utf8": "why",
      "acAsrConf": 252
    }, {
      "utf8": " we",
      "tOffsetMs": 149,
      "acAsrConf": 252
    }, {
      "utf8": " said",
      "tOffsetMs": 179,
      "acAsrConf": 146
    }, {
      "utf8": " make",
      "tOffsetMs": 719,
      "acAsrConf": 216
    }, {
      "utf8": " it",
      "tOffsetMs": 1530,
      "acAsrConf": 205
    }, {
      "utf8": " part",
      "tOffsetMs": 1679,
      "acAsrConf": 244
    }, {
      "utf8": " of",
      "tOffsetMs": 1979,
      "acAsrConf": 252
    }, {
      "utf8": " your",
      "tOffsetMs": 2070,
      "acAsrConf": 252
    }, {
      "utf8": " routine",
      "tOffsetMs": 2130,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2841400,
    "dDurationMs": 3100,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2841410,
    "dDurationMs": 5520,
    "wWinId": 1,
    "segs": [ {
      "utf8": "and",
      "acAsrConf": 207
    }, {
      "utf8": " your",
      "tOffsetMs": 780,
      "acAsrConf": 239
    }, {
      "utf8": " habit",
      "tOffsetMs": 1290,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1679,
      "acAsrConf": 202
    }, {
      "utf8": " that",
      "tOffsetMs": 1919,
      "acAsrConf": 216
    }, {
      "utf8": " will",
      "tOffsetMs": 2369,
      "acAsrConf": 148
    }, {
      "utf8": " be",
      "tOffsetMs": 2760,
      "acAsrConf": 252
    }, {
      "utf8": " more",
      "tOffsetMs": 2879,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2844490,
    "dDurationMs": 2440,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2844500,
    "dDurationMs": 6000,
    "wWinId": 1,
    "segs": [ {
      "utf8": "sustainable",
      "acAsrConf": 205
    }, {
      "utf8": " so",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " yeah",
      "tOffsetMs": 990,
      "acAsrConf": 255
    }, {
      "utf8": " sure",
      "tOffsetMs": 1619,
      "acAsrConf": 252
    }, {
      "utf8": " if",
      "tOffsetMs": 1890,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 2069,
      "acAsrConf": 202
    }, {
      "utf8": " have",
      "tOffsetMs": 2220,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 2400,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2846920,
    "dDurationMs": 3580,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2846930,
    "dDurationMs": 5849,
    "wWinId": 1,
    "segs": [ {
      "utf8": "list",
      "acAsrConf": 236
    }, {
      "utf8": " of",
      "tOffsetMs": 330,
      "acAsrConf": 226
    }, {
      "utf8": " maybe",
      "tOffsetMs": 599,
      "acAsrConf": 252
    }, {
      "utf8": " five",
      "tOffsetMs": 889,
      "acAsrConf": 222
    }, {
      "utf8": " words",
      "tOffsetMs": 1889,
      "acAsrConf": 255
    }, {
      "utf8": " in",
      "tOffsetMs": 2250,
      "acAsrConf": 219
    }, {
      "utf8": " one",
      "tOffsetMs": 2669,
      "acAsrConf": 252
    }, {
      "utf8": " day",
      "tOffsetMs": 3119,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 3329,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2850490,
    "dDurationMs": 2289,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2850500,
    "dDurationMs": 4170,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 204
    }, {
      "utf8": " want",
      "tOffsetMs": 119,
      "acAsrConf": 200
    }, {
      "utf8": " to",
      "tOffsetMs": 329,
      "acAsrConf": 252
    }, {
      "utf8": " use",
      "tOffsetMs": 480,
      "acAsrConf": 252
    }, {
      "utf8": " those",
      "tOffsetMs": 690,
      "acAsrConf": 252
    }, {
      "utf8": " words",
      "tOffsetMs": 990,
      "acAsrConf": 201
    }, {
      "utf8": " throughout",
      "tOffsetMs": 1319,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 2852769,
    "dDurationMs": 1901,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2852779,
    "dDurationMs": 4800,
    "wWinId": 1,
    "segs": [ {
      "utf8": "the",
      "acAsrConf": 228
    }, {
      "utf8": " day",
      "tOffsetMs": 181,
      "acAsrConf": 252
    }, {
      "utf8": " make",
      "tOffsetMs": 361,
      "acAsrConf": 246
    }, {
      "utf8": " a",
      "tOffsetMs": 721,
      "acAsrConf": 252
    }, {
      "utf8": " little",
      "tOffsetMs": 750,
      "acAsrConf": 252
    }, {
      "utf8": " game",
      "tOffsetMs": 990,
      "acAsrConf": 208
    }, {
      "utf8": " of",
      "tOffsetMs": 1260,
      "acAsrConf": 77
    }, {
      "utf8": " it",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    }, {
      "utf8": " right",
      "tOffsetMs": 1651,
      "acAsrConf": 236
    } ]
  }, {
    "tStartMs": 2854660,
    "dDurationMs": 2919,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2854670,
    "dDurationMs": 5099,
    "wWinId": 1,
    "segs": [ {
      "utf8": "so",
      "acAsrConf": 207
    }, {
      "utf8": " maybe",
      "tOffsetMs": 270,
      "acAsrConf": 229
    }, {
      "utf8": " you",
      "tOffsetMs": 449,
      "acAsrConf": 218
    }, {
      "utf8": " could",
      "tOffsetMs": 629,
      "acAsrConf": 226
    }, {
      "utf8": " make",
      "tOffsetMs": 839,
      "acAsrConf": 203
    }, {
      "utf8": " it",
      "tOffsetMs": 990,
      "acAsrConf": 252
    }, {
      "utf8": " more",
      "tOffsetMs": 1139,
      "acAsrConf": 252
    }, {
      "utf8": " fun",
      "tOffsetMs": 1349,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1770,
      "acAsrConf": 203
    }, {
      "utf8": " I",
      "tOffsetMs": 2460,
      "acAsrConf": 236
    } ]
  }, {
    "tStartMs": 2857569,
    "dDurationMs": 2200,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2857579,
    "dDurationMs": 4921,
    "wWinId": 1,
    "segs": [ {
      "utf8": "think",
      "acAsrConf": 252
    }, {
      "utf8": " that's",
      "tOffsetMs": 270,
      "acAsrConf": 236
    }, {
      "utf8": " a",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " good",
      "tOffsetMs": 601,
      "acAsrConf": 252
    }, {
      "utf8": " way",
      "tOffsetMs": 720,
      "acAsrConf": 207
    }, {
      "utf8": " to",
      "tOffsetMs": 780,
      "acAsrConf": 93
    }, {
      "utf8": " remember",
      "tOffsetMs": 930,
      "acAsrConf": 216
    }, {
      "utf8": " to",
      "tOffsetMs": 1621,
      "acAsrConf": 205
    } ]
  }, {
    "tStartMs": 2859759,
    "dDurationMs": 2741,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2859769,
    "dDurationMs": 4500,
    "wWinId": 1,
    "segs": [ {
      "utf8": "our",
      "acAsrConf": 239
    }, {
      "utf8": " shortcut",
      "tOffsetMs": 151,
      "acAsrConf": 200
    }, {
      "utf8": " is",
      "tOffsetMs": 661,
      "acAsrConf": 243
    }, {
      "utf8": " on",
      "tOffsetMs": 810,
      "acAsrConf": 242
    }, {
      "utf8": " some",
      "tOffsetMs": 990,
      "acAsrConf": 255
    }, {
      "utf8": " words",
      "tOffsetMs": 1201,
      "acAsrConf": 255
    }, {
      "utf8": " great",
      "tOffsetMs": 1411,
      "acAsrConf": 216
    }, {
      "utf8": " yes",
      "tOffsetMs": 2010,
      "acAsrConf": 242
    } ]
  }, {
    "tStartMs": 2862490,
    "dDurationMs": 1779,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2862500,
    "dDurationMs": 4049,
    "wWinId": 1,
    "segs": [ {
      "utf8": "and",
      "acAsrConf": 252
    }, {
      "utf8": " thank",
      "tOffsetMs": 89,
      "acAsrConf": 246
    }, {
      "utf8": " you",
      "tOffsetMs": 539,
      "acAsrConf": 252
    }, {
      "utf8": " everyone",
      "tOffsetMs": 599,
      "acAsrConf": 199
    }, {
      "utf8": " who",
      "tOffsetMs": 960,
      "acAsrConf": 252
    }, {
      "utf8": " wrote",
      "tOffsetMs": 1109,
      "acAsrConf": 207
    }, {
      "utf8": " some",
      "tOffsetMs": 1470,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2864259,
    "dDurationMs": 2290,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2864269,
    "dDurationMs": 4020,
    "wWinId": 1,
    "segs": [ {
      "utf8": "some",
      "acAsrConf": 248
    }, {
      "utf8": " nice",
      "tOffsetMs": 691,
      "acAsrConf": 252
    }, {
      "utf8": " comments",
      "tOffsetMs": 901,
      "acAsrConf": 252
    }, {
      "utf8": " here",
      "tOffsetMs": 1201,
      "acAsrConf": 166
    }, {
      "utf8": " we",
      "tOffsetMs": 1711,
      "acAsrConf": 252
    }, {
      "utf8": " didn't",
      "tOffsetMs": 1861,
      "acAsrConf": 247
    }, {
      "utf8": " get",
      "tOffsetMs": 2161,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 2221,
      "acAsrConf": 218
    } ]
  }, {
    "tStartMs": 2866539,
    "dDurationMs": 1750,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2866549,
    "dDurationMs": 4201,
    "wWinId": 1,
    "segs": [ {
      "utf8": "chance",
      "acAsrConf": 180
    }, {
      "utf8": " to",
      "tOffsetMs": 270,
      "acAsrConf": 217
    }, {
      "utf8": " answer",
      "tOffsetMs": 451,
      "acAsrConf": 252
    }, {
      "utf8": " all",
      "tOffsetMs": 631,
      "acAsrConf": 225
    }, {
      "utf8": " of",
      "tOffsetMs": 1020,
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 1081,
      "acAsrConf": 252
    }, {
      "utf8": " questions",
      "tOffsetMs": 1710,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2868279,
    "dDurationMs": 2471,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2868289,
    "dDurationMs": 5401,
    "wWinId": 1,
    "segs": [ {
      "utf8": "but",
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 570,
      "acAsrConf": 142
    }, {
      "utf8": " hope",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " during",
      "tOffsetMs": 990,
      "acAsrConf": 206
    }, {
      "utf8": " our",
      "tOffsetMs": 1411,
      "acAsrConf": 248
    }, {
      "utf8": " presentation",
      "tOffsetMs": 1560,
      "acAsrConf": 252
    }, {
      "utf8": " we",
      "tOffsetMs": 1951,
      "acAsrConf": 210
    } ]
  }, {
    "tStartMs": 2870740,
    "dDurationMs": 2950,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2870750,
    "dDurationMs": 9809,
    "wWinId": 1,
    "segs": [ {
      "utf8": "answered",
      "acAsrConf": 252
    }, {
      "utf8": " some",
      "tOffsetMs": 359,
      "acAsrConf": 252
    }, {
      "utf8": " as",
      "tOffsetMs": 390,
      "acAsrConf": 201
    }, {
      "utf8": " well",
      "tOffsetMs": 900,
      "acAsrConf": 225
    }, {
      "utf8": " so",
      "tOffsetMs": 1140,
      "acAsrConf": 227
    }, {
      "utf8": " if",
      "tOffsetMs": 1920,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 2190,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 2579,
      "acAsrConf": 208
    } ]
  }, {
    "tStartMs": 2873680,
    "dDurationMs": 6879,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2873690,
    "dDurationMs": 9359,
    "wWinId": 1,
    "segs": [ {
      "utf8": "signed",
      "acAsrConf": 252
    }, {
      "utf8": " up",
      "tOffsetMs": 929,
      "acAsrConf": 234
    }, {
      "utf8": " for",
      "tOffsetMs": 1290,
      "acAsrConf": 237
    }, {
      "utf8": " our",
      "tOffsetMs": 1740,
      "acAsrConf": 238
    }, {
      "utf8": " our",
      "tOffsetMs": 4010,
      "acAsrConf": 187
    }, {
      "utf8": " list",
      "tOffsetMs": 5010,
      "acAsrConf": 204
    }, {
      "utf8": " where",
      "tOffsetMs": 6000,
      "acAsrConf": 252
    }, {
      "utf8": " I",
      "tOffsetMs": 6270,
      "acAsrConf": 236
    }, {
      "utf8": " sent",
      "tOffsetMs": 6300,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2880549,
    "dDurationMs": 2500,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2880559,
    "dDurationMs": 6571,
    "wWinId": 1,
    "segs": [ {
      "utf8": "some",
      "acAsrConf": 239
    }, {
      "utf8": " emails",
      "tOffsetMs": 571,
      "acAsrConf": 137
    }, {
      "utf8": " previously",
      "tOffsetMs": 1020,
      "acAsrConf": 248
    }, {
      "utf8": " about",
      "tOffsetMs": 1891,
      "acAsrConf": 245
    }, {
      "utf8": " this",
      "tOffsetMs": 2250,
      "acAsrConf": 226
    } ]
  }, {
    "tStartMs": 2883039,
    "dDurationMs": 4091,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2883049,
    "dDurationMs": 6000,
    "wWinId": 1,
    "segs": [ {
      "utf8": "training",
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 331,
      "acAsrConf": 197
    }, {
      "utf8": " will",
      "tOffsetMs": 810,
      "acAsrConf": 252
    }, {
      "utf8": " receive",
      "tOffsetMs": 841,
      "acAsrConf": 211
    }, {
      "utf8": " the",
      "tOffsetMs": 1831,
      "acAsrConf": 251
    }, {
      "utf8": " replay",
      "tOffsetMs": 2750,
      "acAsrConf": 245
    }, {
      "utf8": " of",
      "tOffsetMs": 3750,
      "acAsrConf": 255
    } ]
  }, {
    "tStartMs": 2887120,
    "dDurationMs": 1929,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2887130,
    "dDurationMs": 4800,
    "wWinId": 1,
    "segs": [ {
      "utf8": "this",
      "acAsrConf": 226
    }, {
      "utf8": " video",
      "tOffsetMs": 330,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 659,
      "acAsrConf": 236
    }, {
      "utf8": " you'll",
      "tOffsetMs": 1050,
      "acAsrConf": 252
    }, {
      "utf8": " receive",
      "tOffsetMs": 1260,
      "acAsrConf": 240
    }, {
      "utf8": " any",
      "tOffsetMs": 1469,
      "acAsrConf": 236
    } ]
  }, {
    "tStartMs": 2889039,
    "dDurationMs": 2891,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2889049,
    "dDurationMs": 6000,
    "wWinId": 1,
    "segs": [ {
      "utf8": "updates",
      "acAsrConf": 252
    }, {
      "utf8": " about",
      "tOffsetMs": 570,
      "acAsrConf": 229
    }, {
      "utf8": " this",
      "tOffsetMs": 871,
      "acAsrConf": 207
    }, {
      "utf8": " to",
      "tOffsetMs": 1381,
      "acAsrConf": 227
    }, {
      "utf8": " make",
      "tOffsetMs": 2010,
      "acAsrConf": 234
    }, {
      "utf8": " sure",
      "tOffsetMs": 2191,
      "acAsrConf": 200
    }, {
      "utf8": " that",
      "tOffsetMs": 2431,
      "acAsrConf": 207
    }, {
      "utf8": " you",
      "tOffsetMs": 2671,
      "acAsrConf": 242
    } ]
  }, {
    "tStartMs": 2891920,
    "dDurationMs": 3129,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2891930,
    "dDurationMs": 6780,
    "wWinId": 1,
    "segs": [ {
      "utf8": "go",
      "acAsrConf": 252
    }, {
      "utf8": " there",
      "tOffsetMs": 240,
      "acAsrConf": 221
    }, {
      "utf8": " it's",
      "tOffsetMs": 480,
      "acAsrConf": 245
    }, {
      "utf8": " speak",
      "tOffsetMs": 1099,
      "acAsrConf": 236
    }, {
      "utf8": " English",
      "tOffsetMs": 2099,
      "acAsrConf": 248
    }, {
      "utf8": " with",
      "tOffsetMs": 2520,
      "acAsrConf": 246
    }, {
      "utf8": " Vanessa",
      "tOffsetMs": 2760,
      "acAsrConf": 238
    } ]
  }, {
    "tStartMs": 2895039,
    "dDurationMs": 3671,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2895049,
    "dDurationMs": 6901,
    "wWinId": 1,
    "segs": [ {
      "utf8": "comm",
      "acAsrConf": 218
    }, {
      "utf8": " slash",
      "tOffsetMs": 540,
      "acAsrConf": 175
    }, {
      "utf8": " conversation",
      "tOffsetMs": 1250,
      "acAsrConf": 245
    }, {
      "utf8": " so",
      "tOffsetMs": 2250,
      "acAsrConf": 251
    }, {
      "utf8": " I",
      "tOffsetMs": 2941,
      "acAsrConf": 245
    }, {
      "utf8": " hope",
      "tOffsetMs": 2970,
      "acAsrConf": 236
    }, {
      "utf8": " that",
      "tOffsetMs": 3391,
      "acAsrConf": 167
    } ]
  }, {
    "tStartMs": 2898700,
    "dDurationMs": 3250,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2898710,
    "dDurationMs": 6869,
    "wWinId": 1,
    "segs": [ {
      "utf8": "if",
      "acAsrConf": 255
    }, {
      "utf8": " you",
      "tOffsetMs": 659,
      "acAsrConf": 252
    }, {
      "utf8": " are",
      "tOffsetMs": 930,
      "acAsrConf": 252
    }, {
      "utf8": " already",
      "tOffsetMs": 1200,
      "acAsrConf": 215
    }, {
      "utf8": " there",
      "tOffsetMs": 1470,
      "acAsrConf": 202
    }, {
      "utf8": " you'll",
      "tOffsetMs": 1980,
      "acAsrConf": 214
    }, {
      "utf8": " be",
      "tOffsetMs": 2909,
      "acAsrConf": 252
    }, {
      "utf8": " able",
      "tOffsetMs": 3000,
      "acAsrConf": 159
    } ]
  }, {
    "tStartMs": 2901940,
    "dDurationMs": 3639,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2901950,
    "dDurationMs": 5579,
    "wWinId": 1,
    "segs": [ {
      "utf8": "to",
      "acAsrConf": 239
    }, {
      "utf8": " receive",
      "tOffsetMs": 329,
      "acAsrConf": 226
    }, {
      "utf8": " the",
      "tOffsetMs": 690,
      "acAsrConf": 236
    }, {
      "utf8": " special",
      "tOffsetMs": 1409,
      "acAsrConf": 202
    }, {
      "utf8": " gifts",
      "tOffsetMs": 2369,
      "acAsrConf": 245
    }, {
      "utf8": " at",
      "tOffsetMs": 3030,
      "acAsrConf": 219
    }, {
      "utf8": " the",
      "tOffsetMs": 3210,
      "acAsrConf": 200
    }, {
      "utf8": " end",
      "tOffsetMs": 3450,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2905569,
    "dDurationMs": 1960,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2905579,
    "dDurationMs": 3990,
    "wWinId": 1,
    "segs": [ {
      "utf8": "of",
      "acAsrConf": 238
    }, {
      "utf8": " our",
      "tOffsetMs": 151,
      "acAsrConf": 252
    }, {
      "utf8": " presentation",
      "tOffsetMs": 210,
      "acAsrConf": 241
    }, {
      "utf8": " for",
      "tOffsetMs": 421,
      "acAsrConf": 140
    }, {
      "utf8": " the",
      "tOffsetMs": 1260,
      "acAsrConf": 252
    }, {
      "utf8": " course",
      "tOffsetMs": 1381,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1651,
      "acAsrConf": 205
    } ]
  }, {
    "tStartMs": 2907519,
    "dDurationMs": 2050,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2907529,
    "dDurationMs": 4411,
    "wWinId": 1,
    "segs": [ {
      "utf8": "if",
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 661,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 780,
      "acAsrConf": 252
    }, {
      "utf8": " any",
      "tOffsetMs": 901,
      "acAsrConf": 252
    }, {
      "utf8": " other",
      "tOffsetMs": 1080,
      "acAsrConf": 252
    }, {
      "utf8": " questions",
      "tOffsetMs": 1290,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1740,
      "acAsrConf": 252
    }, {
      "utf8": " can",
      "tOffsetMs": 1830,
      "acAsrConf": 140
    } ]
  }, {
    "tStartMs": 2909559,
    "dDurationMs": 2381,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2909569,
    "dDurationMs": 4710,
    "wWinId": 1,
    "segs": [ {
      "utf8": "send",
      "acAsrConf": 252
    }, {
      "utf8": " them",
      "tOffsetMs": 181,
      "acAsrConf": 208
    }, {
      "utf8": " to",
      "tOffsetMs": 331,
      "acAsrConf": 185
    }, {
      "utf8": " me",
      "tOffsetMs": 391,
      "acAsrConf": 202
    }, {
      "utf8": " so",
      "tOffsetMs": 651,
      "acAsrConf": 216
    }, {
      "utf8": " thanks",
      "tOffsetMs": 1651,
      "acAsrConf": 241
    }, {
      "utf8": " so",
      "tOffsetMs": 2190,
      "acAsrConf": 252
    }, {
      "utf8": " much",
      "tOffsetMs": 2250,
      "acAsrConf": 233
    } ]
  }, {
    "tStartMs": 2911930,
    "dDurationMs": 2349,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2911940,
    "dDurationMs": 4530,
    "wWinId": 1,
    "segs": [ {
      "utf8": "everyone",
      "acAsrConf": 203
    }, {
      "utf8": " this",
      "tOffsetMs": 510,
      "acAsrConf": 248
    }, {
      "utf8": " was",
      "tOffsetMs": 720,
      "acAsrConf": 207
    }, {
      "utf8": " great",
      "tOffsetMs": 899,
      "acAsrConf": 249
    }, {
      "utf8": " yeah",
      "tOffsetMs": 1200,
      "acAsrConf": 164
    }, {
      "utf8": " yeah",
      "tOffsetMs": 1800,
      "acAsrConf": 236
    }, {
      "utf8": " I",
      "tOffsetMs": 2310,
      "acAsrConf": 231
    } ]
  }, {
    "tStartMs": 2914269,
    "dDurationMs": 2201,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2914279,
    "dDurationMs": 4931,
    "wWinId": 1,
    "segs": [ {
      "utf8": "thank",
      "acAsrConf": 255
    }, {
      "utf8": " you",
      "tOffsetMs": 361,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 480,
      "acAsrConf": 244
    }, {
      "utf8": " much",
      "tOffsetMs": 631,
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 780,
      "acAsrConf": 216
    }, {
      "utf8": " joining",
      "tOffsetMs": 961,
      "acAsrConf": 252
    }, {
      "utf8": " us",
      "tOffsetMs": 1141,
      "acAsrConf": 227
    }, {
      "utf8": " sorry",
      "tOffsetMs": 1500,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2916460,
    "dDurationMs": 2750,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2916470,
    "dDurationMs": 4820,
    "wWinId": 1,
    "segs": [ {
      "utf8": "about",
      "acAsrConf": 252
    }, {
      "utf8": " the",
      "tOffsetMs": 210,
      "acAsrConf": 201
    }, {
      "utf8": " technical",
      "tOffsetMs": 660,
      "acAsrConf": 243
    }, {
      "utf8": " difficulties",
      "tOffsetMs": 1099,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2919200,
    "dDurationMs": 2090,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2919210,
    "dDurationMs": 4750,
    "wWinId": 1,
    "segs": [ {
      "utf8": "it's",
      "acAsrConf": 255
    }, {
      "utf8": " our",
      "tOffsetMs": 1000,
      "acAsrConf": 252
    }, {
      "utf8": " first",
      "tOffsetMs": 1149,
      "acAsrConf": 205
    }, {
      "utf8": " time",
      "tOffsetMs": 1389,
      "acAsrConf": 229
    }, {
      "utf8": " using",
      "tOffsetMs": 1600,
      "acAsrConf": 232
    }, {
      "utf8": " Google",
      "tOffsetMs": 1659,
      "acAsrConf": 146
    } ]
  }, {
    "tStartMs": 2921280,
    "dDurationMs": 2680,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2921290,
    "dDurationMs": 4680,
    "wWinId": 1,
    "segs": [ {
      "utf8": "Hangouts",
      "acAsrConf": 143
    }, {
      "utf8": " yes",
      "tOffsetMs": 780,
      "acAsrConf": 216
    }, {
      "utf8": " yes",
      "tOffsetMs": 1140,
      "acAsrConf": 201
    }, {
      "utf8": " thank",
      "tOffsetMs": 1650,
      "acAsrConf": 202
    }, {
      "utf8": " you",
      "tOffsetMs": 2100,
      "acAsrConf": 252
    }, {
      "utf8": " for",
      "tOffsetMs": 2160,
      "acAsrConf": 164
    }, {
      "utf8": " your",
      "tOffsetMs": 2549,
      "acAsrConf": 226
    } ]
  }, {
    "tStartMs": 2923950,
    "dDurationMs": 2020,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2923960,
    "dDurationMs": 3690,
    "wWinId": 1,
    "segs": [ {
      "utf8": "patience",
      "acAsrConf": 247
    }, {
      "utf8": " at",
      "tOffsetMs": 450,
      "acAsrConf": 250
    }, {
      "utf8": " the",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " beginning",
      "tOffsetMs": 720,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 1020,
      "acAsrConf": 225
    }, {
      "utf8": " for",
      "tOffsetMs": 1440,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2925960,
    "dDurationMs": 1690,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2925970,
    "dDurationMs": 4619,
    "wWinId": 1,
    "segs": [ {
      "utf8": "sticking",
      "acAsrConf": 200
    }, {
      "utf8": " through",
      "tOffsetMs": 510,
      "acAsrConf": 187
    }, {
      "utf8": " the",
      "tOffsetMs": 869,
      "acAsrConf": 252
    }, {
      "utf8": " whole",
      "tOffsetMs": 1260,
      "acAsrConf": 252
    }, {
      "utf8": " presentation",
      "tOffsetMs": 1440,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2927640,
    "dDurationMs": 2949,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2927650,
    "dDurationMs": 4800,
    "wWinId": 1,
    "segs": [ {
      "utf8": "yes",
      "acAsrConf": 143
    }, {
      "utf8": " we",
      "tOffsetMs": 959,
      "acAsrConf": 205
    }, {
      "utf8": " got",
      "tOffsetMs": 1620,
      "acAsrConf": 202
    }, {
      "utf8": " a",
      "tOffsetMs": 1770,
      "acAsrConf": 252
    }, {
      "utf8": " chance",
      "tOffsetMs": 1800,
      "acAsrConf": 238
    }, {
      "utf8": " to",
      "tOffsetMs": 2040,
      "acAsrConf": 200
    }, {
      "utf8": " talk",
      "tOffsetMs": 2250,
      "acAsrConf": 248
    }, {
      "utf8": " for",
      "tOffsetMs": 2459,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 2699,
      "acAsrConf": 206
    }, {
      "utf8": " long",
      "tOffsetMs": 2760,
      "acAsrConf": 228
    } ]
  }, {
    "tStartMs": 2930579,
    "dDurationMs": 1871,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2930589,
    "dDurationMs": 4441,
    "wWinId": 1,
    "segs": [ {
      "utf8": "time",
      "acAsrConf": 252
    }, {
      "utf8": " hopefully",
      "tOffsetMs": 211,
      "acAsrConf": 164
    }, {
      "utf8": " you",
      "tOffsetMs": 780,
      "acAsrConf": 64
    }, {
      "utf8": " got",
      "tOffsetMs": 1111,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 1260,
      "acAsrConf": 222
    }, {
      "utf8": " learn",
      "tOffsetMs": 1441,
      "acAsrConf": 252
    }, {
      "utf8": " some",
      "tOffsetMs": 1561,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2932440,
    "dDurationMs": 2590,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2932450,
    "dDurationMs": 4710,
    "wWinId": 1,
    "segs": [ {
      "utf8": "great",
      "acAsrConf": 217
    }, {
      "utf8": " tips",
      "tOffsetMs": 210,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 419,
      "acAsrConf": 168
    }, {
      "utf8": " I'll",
      "tOffsetMs": 840,
      "acAsrConf": 242
    }, {
      "utf8": " send",
      "tOffsetMs": 1649,
      "acAsrConf": 211
    }, {
      "utf8": " everyone",
      "tOffsetMs": 1919,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 2040,
      "acAsrConf": 202
    } ]
  }, {
    "tStartMs": 2935020,
    "dDurationMs": 2140,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2935030,
    "dDurationMs": 6240,
    "wWinId": 1,
    "segs": [ {
      "utf8": "message",
      "acAsrConf": 252
    }, {
      "utf8": " some",
      "tOffsetMs": 360,
      "acAsrConf": 252
    }, {
      "utf8": " information",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " after",
      "tOffsetMs": 1170,
      "acAsrConf": 249
    }, {
      "utf8": " this",
      "tOffsetMs": 1650,
      "acAsrConf": 243
    }, {
      "utf8": " so",
      "tOffsetMs": 1890,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2937150,
    "dDurationMs": 4120,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2937160,
    "dDurationMs": 6390,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 252
    }, {
      "utf8": " can",
      "tOffsetMs": 60,
      "acAsrConf": 252
    }, {
      "utf8": " review",
      "tOffsetMs": 1550,
      "acAsrConf": 219
    }, {
      "utf8": " what",
      "tOffsetMs": 2550,
      "acAsrConf": 248
    }, {
      "utf8": " we",
      "tOffsetMs": 2880,
      "acAsrConf": 211
    }, {
      "utf8": " talked",
      "tOffsetMs": 3150,
      "acAsrConf": 252
    }, {
      "utf8": " about",
      "tOffsetMs": 3480,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 3630,
      "acAsrConf": 214
    } ]
  }, {
    "tStartMs": 2941260,
    "dDurationMs": 2290,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2941270,
    "dDurationMs": 7260,
    "wWinId": 1,
    "segs": [ {
      "utf8": "hopefully",
      "acAsrConf": 200
    }, {
      "utf8": " you'll",
      "tOffsetMs": 450,
      "acAsrConf": 252
    }, {
      "utf8": " be",
      "tOffsetMs": 630,
      "acAsrConf": 252
    }, {
      "utf8": " able",
      "tOffsetMs": 690,
      "acAsrConf": 202
    }, {
      "utf8": " to",
      "tOffsetMs": 810,
      "acAsrConf": 200
    }, {
      "utf8": " take",
      "tOffsetMs": 1319,
      "acAsrConf": 252
    }, {
      "utf8": " action",
      "tOffsetMs": 2069,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2943540,
    "dDurationMs": 4990,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2943550,
    "dDurationMs": 6600,
    "wWinId": 1,
    "segs": [ {
      "utf8": "take",
      "acAsrConf": 206
    }, {
      "utf8": " action",
      "tOffsetMs": 990,
      "acAsrConf": 85
    }, {
      "utf8": " and",
      "tOffsetMs": 1440,
      "acAsrConf": 252
    }, {
      "utf8": " find",
      "tOffsetMs": 1650,
      "acAsrConf": 240
    }, {
      "utf8": " some",
      "tOffsetMs": 2220,
      "acAsrConf": 252
    }, {
      "utf8": " way",
      "tOffsetMs": 3120,
      "acAsrConf": 238
    }, {
      "utf8": " to",
      "tOffsetMs": 3990,
      "acAsrConf": 252
    }, {
      "utf8": " speak",
      "tOffsetMs": 4350,
      "acAsrConf": 219
    } ]
  }, {
    "tStartMs": 2948520,
    "dDurationMs": 1630,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2948530,
    "dDurationMs": 3839,
    "wWinId": 1,
    "segs": [ {
      "utf8": "whether",
      "acAsrConf": 205
    }, {
      "utf8": " it's",
      "tOffsetMs": 540,
      "acAsrConf": 252
    }, {
      "utf8": " listening",
      "tOffsetMs": 780,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 990,
      "acAsrConf": 186
    }, {
      "utf8": " reading",
      "tOffsetMs": 1410,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2950140,
    "dDurationMs": 2229,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2950150,
    "dDurationMs": 4199,
    "wWinId": 1,
    "segs": [ {
      "utf8": "speaking",
      "acAsrConf": 203
    }, {
      "utf8": " some",
      "tOffsetMs": 900,
      "acAsrConf": 228
    }, {
      "utf8": " way",
      "tOffsetMs": 1530,
      "acAsrConf": 228
    }, {
      "utf8": " that",
      "tOffsetMs": 1650,
      "acAsrConf": 146
    }, {
      "utf8": " you'll",
      "tOffsetMs": 1860,
      "acAsrConf": 252
    }, {
      "utf8": " have",
      "tOffsetMs": 2010,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 2160,
      "acAsrConf": 214
    } ]
  }, {
    "tStartMs": 2952359,
    "dDurationMs": 1990,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2952369,
    "dDurationMs": 3871,
    "wWinId": 1,
    "segs": [ {
      "utf8": "chance",
      "acAsrConf": 216
    }, {
      "utf8": " to",
      "tOffsetMs": 211,
      "acAsrConf": 208
    }, {
      "utf8": " speak",
      "tOffsetMs": 451,
      "acAsrConf": 252
    }, {
      "utf8": " use",
      "tOffsetMs": 631,
      "acAsrConf": 187
    }, {
      "utf8": " it",
      "tOffsetMs": 1111,
      "acAsrConf": 252
    }, {
      "utf8": " or",
      "tOffsetMs": 1291,
      "acAsrConf": 252
    }, {
      "utf8": " lose",
      "tOffsetMs": 1411,
      "acAsrConf": 202
    }, {
      "utf8": " it",
      "tOffsetMs": 1621,
      "acAsrConf": 252
    }, {
      "utf8": " yes",
      "tOffsetMs": 1651,
      "acAsrConf": 200
    } ]
  }, {
    "tStartMs": 2954339,
    "dDurationMs": 1901,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2954349,
    "dDurationMs": 3721,
    "wWinId": 1,
    "segs": [ {
      "utf8": "you",
      "acAsrConf": 239
    }, {
      "utf8": " doesn't",
      "tOffsetMs": 211,
      "acAsrConf": 252
    }, {
      "utf8": " lose",
      "tOffsetMs": 481,
      "acAsrConf": 218
    }, {
      "utf8": " it",
      "tOffsetMs": 811,
      "acAsrConf": 252
    }, {
      "utf8": " so",
      "tOffsetMs": 1020,
      "acAsrConf": 242
    }, {
      "utf8": " thank",
      "tOffsetMs": 1260,
      "acAsrConf": 245
    }, {
      "utf8": " you",
      "tOffsetMs": 1591,
      "acAsrConf": 234
    }, {
      "utf8": " for",
      "tOffsetMs": 1711,
      "acAsrConf": 218
    } ]
  }, {
    "tStartMs": 2956230,
    "dDurationMs": 1840,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2956240,
    "dDurationMs": 4830,
    "wWinId": 1,
    "segs": [ {
      "utf8": "joining",
      "acAsrConf": 201
    }, {
      "utf8": " us",
      "tOffsetMs": 240,
      "acAsrConf": 252
    }, {
      "utf8": " and",
      "tOffsetMs": 420,
      "acAsrConf": 239
    }, {
      "utf8": " we'll",
      "tOffsetMs": 720,
      "acAsrConf": 220
    }, {
      "utf8": " talk",
      "tOffsetMs": 1410,
      "acAsrConf": 229
    }, {
      "utf8": " to",
      "tOffsetMs": 1590,
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 1619,
      "acAsrConf": 187
    }, {
      "utf8": " again",
      "tOffsetMs": 1770,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2958060,
    "dDurationMs": 3010,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2958070,
    "dDurationMs": 5039,
    "wWinId": 1,
    "segs": [ {
      "utf8": "another",
      "acAsrConf": 227
    }, {
      "utf8": " time",
      "tOffsetMs": 330,
      "acAsrConf": 208
    }, {
      "utf8": " and",
      "tOffsetMs": 690,
      "acAsrConf": 236
    }, {
      "utf8": " see",
      "tOffsetMs": 1529,
      "acAsrConf": 239
    }, {
      "utf8": " you",
      "tOffsetMs": 2279,
      "acAsrConf": 252
    }, {
      "utf8": " if",
      "tOffsetMs": 2310,
      "acAsrConf": 206
    }, {
      "utf8": " you",
      "tOffsetMs": 2549,
      "acAsrConf": 252
    }, {
      "utf8": " want",
      "tOffsetMs": 2700,
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 2880,
      "acAsrConf": 229
    } ]
  }, {
    "tStartMs": 2961060,
    "dDurationMs": 2049,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2961070,
    "dDurationMs": 4680,
    "wWinId": 1,
    "segs": [ {
      "utf8": "join",
      "acAsrConf": 252
    }, {
      "utf8": " our",
      "tOffsetMs": 150,
      "acAsrConf": 62
    }, {
      "utf8": " first",
      "tOffsetMs": 360,
      "acAsrConf": 252
    }, {
      "utf8": " 30",
      "tOffsetMs": 600,
      "acAsrConf": 252
    }, {
      "utf8": " days",
      "tOffsetMs": 810,
      "acAsrConf": 252
    }, {
      "utf8": " of",
      "tOffsetMs": 900,
      "acAsrConf": 206
    }, {
      "utf8": " English",
      "tOffsetMs": 1170,
      "acAsrConf": 207
    }, {
      "utf8": " make",
      "tOffsetMs": 1470,
      "acAsrConf": 201
    } ]
  }, {
    "tStartMs": 2963099,
    "dDurationMs": 2651,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2963109,
    "dDurationMs": 4260,
    "wWinId": 1,
    "segs": [ {
      "utf8": "sure",
      "acAsrConf": 252
    }, {
      "utf8": " you",
      "tOffsetMs": 181,
      "acAsrConf": 252
    }, {
      "utf8": " do",
      "tOffsetMs": 331,
      "acAsrConf": 252
    }, {
      "utf8": " it",
      "tOffsetMs": 510,
      "acAsrConf": 252
    }, {
      "utf8": " in",
      "tOffsetMs": 631,
      "acAsrConf": 164
    }, {
      "utf8": " 24",
      "tOffsetMs": 841,
      "acAsrConf": 245
    }, {
      "utf8": " hours",
      "tOffsetMs": 1260,
      "acAsrConf": 213
    }, {
      "utf8": " you'll",
      "tOffsetMs": 1740,
      "acAsrConf": 177
    }, {
      "utf8": " have",
      "tOffsetMs": 2461,
      "acAsrConf": 252
    }, {
      "utf8": " a",
      "tOffsetMs": 2490,
      "acAsrConf": 157
    } ]
  }, {
    "tStartMs": 2965740,
    "dDurationMs": 1629,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2965750,
    "dDurationMs": 4770,
    "wWinId": 1,
    "segs": [ {
      "utf8": "chance",
      "acAsrConf": 252
    }, {
      "utf8": " to",
      "tOffsetMs": 240,
      "acAsrConf": 186
    }, {
      "utf8": " get",
      "tOffsetMs": 450,
      "acAsrConf": 252
    }, {
      "utf8": " those",
      "tOffsetMs": 570,
      "acAsrConf": 236
    }, {
      "utf8": " special",
      "tOffsetMs": 810,
      "acAsrConf": 150
    }, {
      "utf8": " free",
      "tOffsetMs": 1290,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2967359,
    "dDurationMs": 3161,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2967369,
    "dDurationMs": 7431,
    "wWinId": 1,
    "segs": [ {
      "utf8": "editions",
      "acAsrConf": 248
    }, {
      "utf8": " together",
      "tOffsetMs": 661,
      "acAsrConf": 206
    }, {
      "utf8": " all",
      "tOffsetMs": 1191,
      "acAsrConf": 184
    }, {
      "utf8": " right",
      "tOffsetMs": 2191,
      "acAsrConf": 200
    }, {
      "utf8": " thanks",
      "tOffsetMs": 2371,
      "acAsrConf": 204
    }, {
      "utf8": " so",
      "tOffsetMs": 3121,
      "acAsrConf": 252
    } ]
  }, {
    "tStartMs": 2970510,
    "dDurationMs": 4290,
    "wWinId": 1,
    "aAppend": 1,
    "segs": [ {
      "utf8": "\n"
    } ]
  }, {
    "tStartMs": 2970520,
    "dDurationMs": 4280,
    "wWinId": 1,
    "segs": [ {
      "utf8": "much",
      "acAsrConf": 187
    }, {
      "utf8": " thanks",
      "tOffsetMs": 230,
      "acAsrConf": 225
    }, {
      "utf8": " everybody",
      "tOffsetMs": 1230,
      "acAsrConf": 252
    }, {
      "utf8": " yes",
      "tOffsetMs": 1500,
      "acAsrConf": 163
    }, {
      "utf8": " goodbye",
      "tOffsetMs": 1920,
      "acAsrConf": 204
    } ]
  } ]
}`
)
