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
	if string(j3.ToSrt())!=subtitleSrt{
		t.Fail()
	}
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
