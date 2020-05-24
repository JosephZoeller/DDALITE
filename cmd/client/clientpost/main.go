package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)

var SDNCAddr = "192.168.1.20" // <INSERT SDNC ADDRESS HERE>
const workFilePath = "work_order.json"

func init() {
	flag.Parse()
}

func main() {
	switch flag.Arg(0) {
	case "seek":
		workFile, err := os.Open(workFilePath)
		if err != nil {
			log.Printf("%s Error - %s\n", workFilePath, err.Error())
			os.Exit(1)
		}
		seekReq(workFile)
	case "seekEmbed":
		seekReqProgrammatic()
	case "teardown":
		teardownReq()
	case "genjson":
		OutputJson()
	//case "workerTest":
	//	workerTestReq()
	default:
		log.Println(usage)
		return
	}

	log.Println("DONE")
}

func seekReqProgrammatic() {
	groupSize := 3
	totalGroups := 10
	var waitInterval time.Duration = 30
	
	colliders := make([]cityhashutil.ColliderSpecifications, 0)

	for i := 0; i < totalGroups; i++ {
		colliders = append(colliders, cityhashutil.ColliderSpecifications{
			InputHashes: []uint64{100018398406407, 85894109417755, 100344186467618, 100705409720908, 10084092424946, 101038941345065, 101353819857458, 102173724970085, 102321488268028, 102715687304464, 102811854670790, 10310337011531, 104154504280890, 104666841482893, 105284202074245, 106076310574035, 106208077970767, 106587990740832, 107802774576470, 109035421338545, 110183014877418, 110187210323059, 113157661544328, 114772578968719, 115014120695510, 11536543653399, 115587439426589, 115621406961613, 116482325613822, 116845306561127, 117229107888337, 118065835591265, 119336623252032, 119731277046827, 119958107505363, 119984435706406, 120270760612510, 121032000668620, 121035842863494, 121610670284021, 122898547313985, 124747721014436, 12607858374059, 126558355994374, 127441289898194, 127596943633518, 127638733590100, 128247721523386, 129095633094216, 129498085446864, 129498678523132, 129818498658395, 130219926364552, 130430345148764, 131485665533707, 132130685949643, 13216472786127, 132524993802473, 133342682207414, 133942827548200, 134451830201100, 13546292615043, 135808022692213, 135812550361498, 135814319458652, 136045145264627, 136229993767345, 136331588267811, 136354801730908, 136770444341289, 136979138120404, 139081541402355, 139222428653447, 139733073712999, 140283266436887, 140610086406515, 140693398840635, 143678976442170, 144257574756652, 145083297758212, 145739801048304, 145864607999828, 147909351436460, 148072711192111, 148797272599418, 148935797357780, 149348065080275, 151115046508671, 151410448453480, 152380689078046, 152792419387252, 152852247188084, 155874073873252, 156130383350500, 156181678376861, 157310544915746, 157402054338947, 157482910679507, 157687679407703, 157766615537467, 158455652395075, 159889855567111, 159918592545707, 16022057869344, 16032351500764, 160497155249166, 160752404839235, 161117327810469, 161154109394771, 161942088228175, 162413521636424, 162627077207864, 163701056378428, 164051404194680, 165098553667181, 165231953886971, 165991983806716, 16787986784266, 168002599570263, 170433418243717, 172501886224886, 172585505229095, 173989471257169, 174787648043115, 178590574530940, 179673835993582, 18162403183311, 182220976625964, 182558667073806, 182721821583302, 182990262559578, 183697092816706, 183836951833349, 18387845381859, 184149668579875, 184260734637614, 184613706032485, 184723361925497, 186299151436286, 187239081157422, 188643438332811, 189805571181043, 189885888633548, 190390685184598, 190670263378087, 191372518393815, 191947579684016, 1923430442702, 193288166795997, 193407690640890, 193414341055349, 193995984300117, 195636469891810, 196148785507973, 197079104529403, 197128892969526, 197244572600781, 198520783384717, 202190862149068, 202384999040422, 203580553224979, 205956825449312, 208827360265579, 209491297356261, 211128440901290, 211389317961521, 211472136399331, 211692387495580, 212387881526675, 216064519912663, 216505082093396, 217147612833288, 217421127966789, 217986995660710, 218156423355872, 218627673754400, 218840997917537, 221227959471847, 221679689966392, 221966759100867, 223355182652541, 223852504351741, 224027568749162, 225207429120179, 227874969906556, 228788678577881, 23278352516925, 23463391709190, 236020746423326, 238472779820926, 240687601396086, 240767065887893, 242786796224364, 24300256674803, 243737399290754, 24422607134746, 244720175117252, 244835221446562, 245037001248461, 246452969109747, 246650590047933, 246807753043925, 246828295310136, 24807816550370, 248316319164044, 248977470868292, 249209013147774, 250395549009486, 250810208850172, 251991314440154, 252228964911159, 252945228709488, 253929859193491, 254810460981251, 255117744840406, 256723045212428, 257075118197811, 25789544013089, 258826363407794, 259633369427611, 259926065145866, 260482628825493, 262559729479142, 26323030650365, 263794709895511, 26511062512644, 267225850343878, 26762051400521, 268229368231798, 268585427932377, 269702399666585, 269921198977841, 269960630535862, 270062937684180, 272140506733628, 274383749473104, 274914826523517, 275146552507742, 2755808233795, 276086469541850, 276981779850611, 277409320524446, 27899082493518, 279335325210660, 281001333207226, 281390524378745, 281455042442475, 28425069143671, 28848249376204, 30033909465817, 31921291762616, 32627462378062, 33218873262556, 34780538976698, 36058030541339, 36316582427757, 3770146511559, 38853151620730, 39242550506188, 40084429642847, 40722697802122, 41156550617742, 41334545889889, 41603542035227, 45931149596213, 46358439943518, 49263949749627, 49863833055554, 50592559961664, 51144673887396, 52460443125074, 53487797509967, 5465524759509, 55533807448729, 55640619164809, 57613484266400, 60008566283269, 60563395080164, 60668380447603, 60795138794923, 60988805759440, 62447405246809, 63120590333790, 64004971506200, 64418549477795, 64942172097437, 64990148870029, 65339965475737, 65776824223389, 66440439364686, 6704539012570, 67128625567568, 69744545156671, 7107569878845, 71885869455063, 72081020218900, 72171198940931, 73927543969519, 74277219576128, 75290771918125, 75610718981125, 76516009669142, 76891595433441, 7733774688470, 78570593565868, 78953731717284, 81671944697182, 81697005424857, 81827248469923, 8287270209564, 83471246330419, 85996665853390, 86053755702769, 87067160542761, 88624349846462, 90830427877514, 91141748613160, 91231667148359, 91932972848672, 91956173861316, 92330539077970, 93600798934715, 94163123710410, 94485681156258, 94831980750732, 94854515111911, 95900451979075, 96314815898542, 96528799475605, 96841703780305, 97450771643871, 97599190796329, 97851076449043, 98005344675213, 98334132130808, 99020600922051, 99038264810477, 99426570292499},
			Dictionary:  []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"},
			Delimiter:   "",
			Depth:       5,
			StartsWith:  fmt.Sprintf("SKL_%3d_", i),
		})

		post, err := json.Marshal(cityhashutil.ClientSpecifications{Colliders: colliders})
		if err != nil {
			panic(err)
		}

		log.Println("Sending: " + string(post))
		seekReq(bytes.NewReader(post))

		if (i+1)%groupSize == 0 {
			colliders = make([]cityhashutil.ColliderSpecifications, 0)
			time.Sleep(waitInterval * time.Second)
		}
	}
}

func seekReq(post io.Reader) {
	msg := cityhashutil.MessageResponse{}
	rsp, err := http.Post(fmt.Sprintf("http://%s:666/ClientToSDNC", SDNCAddr), "application/json", post)
	if err != nil {
		panic(err)
	} else {
		err = json.NewDecoder(rsp.Body).Decode(&msg)
		if err != nil {
			log.Println("Failed to decode server response")
		} else {
			log.Println(msg.Message)
		}
	}
}

func OutputJson() {
	post, err := json.Marshal(cityhashutil.ClientSpecifications{
		Colliders: []cityhashutil.ColliderSpecifications{
			{
				InputHashes: []uint64{85894109417755},
				Dictionary:  []string{"A", "p", "l", "e"},
				Delimiter:   "",
				Depth:       5,
			},
		},
	})
	if err != nil {
		panic(err)
	}
	log.Println(string(post))
}

func teardownReq() {
	msg := cityhashutil.MessageResponse{}
	rsp, err := http.Post(fmt.Sprintf("http://%s:666/teardown", SDNCAddr), "application/json", nil)
	if err != nil {
		panic(err)
	} else {
		err = json.NewDecoder(rsp.Body).Decode(&msg)
		if err != nil {
			log.Println("Failed to decode server response")
		} else {
			log.Println(msg.Message)
		}
	}
}
