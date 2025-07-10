# 📘 Referensi
1. 📚 *Learn Golang and Python Quickly Coding for Beginners with Hands-On Projects* — J. J. Tam  
2. 🤖 ChatGPT  
3. 🎥 [Understanding Allocations: The Stack and the Heap - GopherCon SG 2019](https://www.youtube.com/watch?v=ZMZpH4yT7M0)  
4. 🎥 [Belajar Golang Concurrency Gampang](https://www.youtube.com/watch?v=fHxCpoF88Qg)  
5. 👨‍💻 Go for Devpos  
6. 🧱 [Belajar Clean Architecture dari YouTube Backend Magang](https://github.com/medivh13/for_learning_2)

---

# 🚀 Hal Penting yang Harus Dipelajari
🧠 *Dasar & Konsep Lanjutan Golang*:
- Goroutine dan sinkronisasi  
- Pointer (Constant Pointer vs. Pointer to a Constant)  
- Struct Pointer  
- Interface  
- Constructor *(Go tidak punya, jadi harus manual)*  
- Pointer Function  
- `new` Pointer  
- Rune  
- Multithread  
- Channel dan Channel Buffer  
- Heap & Stack  
- Stack Goroutine  
- Garbage Collection  
- Escape Heap / No Escape?  
- Currency = Multithreading?  
- Parallelism  
- Goroutines vs. Thread?  
- Asynchronous  
- `itoa`, `atoi`

🔁 *Concurrency & Asynchrony*:
- ✅ Goroutine  
- 🔄 Channel, Blocking Channel, Buffered Size Channel  
- 🔁 Range dan Channel Closing  
- 🔘 Select dan Default Selection  
- ❓ Async Await?  
- ❌ Deadlock?  
- ⚠️ Race Condition  
- ⏱️ Time Management: `time.Sleep`, `sync.WaitGroup`  
- 💡 Thread lebih ringan => Gunakan **Goroutine**

---

# 🧾 Format Verb dalam Go
| Verb | Fungsi |
|------|--------|
| `%v` | Menampilkan nilai dalam format default |
| `%T` | Menampilkan tipe dari nilai |
| `%d` | Menampilkan angka desimal *(integer)* |
| `%b` | Menampilkan angka dalam format biner |
| `%c` | Menampilkan karakter Unicode |
| `%x` | Menampilkan angka heksadesimal *(huruf kecil)* |
| `%X` | Menampilkan angka heksadesimal *(huruf besar)* |
| `%f` | Menampilkan angka desimal *(floating point)* |
| `%e` | Notasi ilmiah *(huruf kecil)* |
| `%E` | Notasi ilmiah *(huruf besar)* |
| `%s` | Menampilkan string — `tet jooko` |
| `%q` | Menampilkan string yang dikutip |
| `%p` | Menampilkan pointer *(alamat memori)* |

---

# ⚔️ New vs Make
![new vs make](image.png)

---

# 🧬 Go Lang: Passing by Value
> ⚠️ **Ingat:** Go **secara default menggunakan _passing by value_**, artinya perubahan terhadap variabel **tidak memengaruhi variabel asli** — karena yang diubah adalah salinannya.

<span style="color:red"><strong>🚨 KECUALI PAKAI POINTER!!!!!</strong></span>

---

### 🧷 Simbol Pointer `*`
> Dibaca **pointer ke ...**

Contoh:  
- `*int` → pointer ke `int`  
- `*intPtr` → pointer ke variabel `intPtr`

### 📍 Simbol Alamat `&`
> Dibaca **alamat dari ...**

Contoh:  
- `&employee` → alamat dari `employee`
