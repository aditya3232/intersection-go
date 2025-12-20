package intersection

import "strings"

/*
Soal:
Diberikan sebuah array berisi dua string, masing-masing string berisi
sekumpulan angka yang dipisahkan dengan tanda koma:

    strArr = [ "listA", "listB" ]

Tugas:
1. Pecah setiap string menjadi daftar angka.
2. Bersihkan spasi yang mungkin ada pada tiap elemen.
3. Temukan angka-angka yang muncul di *kedua* daftar tersebut (intersection / irisan).
4. Urutan hasil harus mengikuti urutan kemunculan pada list kedua (listB).
5. Jika tidak ada angka yang sama, return string kosong "".
6. Jika ada, gabungkan hasilnya dengan tanda koma ",".

Contoh:
Input:
    ["1,3,4,7,13", "1,2,4,13,15"]

Langkah:
    listA = [1,3,4,7,13]
    listB = [1,2,4,13,15]

    Angka yang muncul di kedua list:
        1 → ada di A dan B
        4 → ada di A dan B
        13 → ada di A dan B

Output:
    "1,4,13"

Kasus tambahan:
- Jika salah satu list kosong → output = ""
- Jika tidak ada angka yang sama → output = ""
- Spasi di dalam input harus diabaikan (misalnya "1, 2, 3")
*/

func FindIntersection(strArr []string) string {
	if len(strArr) != 2 {
		return ""
	}

	first := strings.Split(strArr[0], ",")
	second := strings.Split(strArr[1], ",")

	for i := range first {
		first[i] = strings.TrimSpace(first[i])
	}
	for i := range second {
		second[i] = strings.TrimSpace(second[i])
	}

	set := make(map[string]bool)
	for _, v := range first {
		set[v] = true
	}

	var result []string
	for _, v := range second {
		if set[v] {
			result = append(result, v)
		}
	}

	if len(result) == 0 {
		return ""
	}

	return strings.Join(result, ",")

}
