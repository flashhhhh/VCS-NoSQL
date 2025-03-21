# Elasticsearch
## So sánh Elasticsearch với các hệ quản trị SQL/ NoSQL khác
### SQL

| Tiêu chí                 | Elasticsearch | SQL DBMS |
| ------------------------ | ------------------ | --------------- |
| Mô hình dữ liệu    | Document-based (JSON) | Relational (bảng) |
| Truy vấn               | DSL hoặc REST API | SQL |
| Tốc độ tìm kiếm   | **Nhanh, tối ưu cho tìm kiếm văn bản (FTS) ** | Nhanh hơn với truy vấn đơn giản nhưng chậm hơn khi tìm kiếm trên dữ liệu lớn |
| Tốc độ đọc ghi     | Chậm do có overhead của index, đặc biệt rất chậm với những truy vấn sửa hoặc xóa do phải index lại | **Nhanh do chỉ phải sửa đổi 1 bản ghi trong bảng** |
| Hỗ trợ giao dịch (tính ACID) | Không hỗ trợ ACID | **Hỗ trợ ACID đầy đủ** |
| Khả năng mở rộng | **Mở rộng theo chiều ngang** | Chủ yếu là theo chiều dọc |
| Tính nhất quán | Eventual Consistency | **Strong Consistency** |
| Use cases | Search engine, phân tích real-time | Ứng dụng cần tính giao dịch, quản lý dữ liệu có cấu trúc |

### NoSQL
#### MongoDB
| Tiêu chí                 | Elasticsearch | MongoDB |
| ------------------------ | ------------------ | --------------- |
| Mô hình dữ liệu    | Document-based (JSON) | Document-based (JSON) |
| Truy vấn               | DSL hoặc REST API | MQL (Mongodb Query Language) |
| Lược đồ               | Không cần lược đồ, nhưng vẫn cần có mapping | **Không cần lược đồ hoàn toàn** |
| Hỗ trợ giao dịch (tính ACID) | Không hỗ trợ ACID | **Có hỗ trợ tính ACID từ Mongo 4.0** |
| Khả năng mở rộng | Mở rộng theo chiều ngang (Replication + Sharding) | Mở rộng theo chiều ngang (Replication + Sharding) |
| Tính nhất quán | Eventual Consistency | **Strong Consistency hoặc Eventual Consistency tùy theo cấu hình** |
| Use cases | Search engine, phân tích real-time | NoSQL cho web app, IoT |

#### Redis
| Tiêu chí                 | Elasticsearch | Redis |
| ------------------------ | ------------------ | --------------- |
| Mô hình dữ liệu    | Document-based (JSON) | Key-value |
| Truy vấn               | DSL hoặc REST API | Tìm kiếm key-value |
| Lưu trữ                | **Lưu trữ vĩnh cửu vào ổ đĩa** | Lưu trữ tạm thời vào RAM |
| Tốc độ đọc          | Nhanh | **Cực kỳ nhanh do truy vấn trong RAM** |
| Tốc độ viết         | Chậm do overhead của index | **Cực kỳ nhanh do chỉ phải ghi vào RAM** |
| Time to live        | **Lưu trữ vĩnh cửu** | Cần TTL để lọc bớt dữ liệu |
| Khả năng mở rộng | Mở rộng theo chiều ngang (Replication + Sharding) | Mở rộng theo chiều ngang (Replication + Clustering) |
| Use cases | Search engine, phân tích real-time | Cache, lưu trữ phiên, real-time leaderboard, pub/sub |

## CRUD
Elasticsearch là một công cụ tìm kiếm, và có các thuật ngữ tương đồng với SQL như sau:

| Elasticsearch | SQL        |
| ------------------ | ----------- |
| cluster           | database |
| index             | table        |
| document      | row          |
| field               | column    |

Tạo một index mới với tên là companies:
```elasticsearch
PUT /companies
{
	"settings": {
		"number_of_shards": 1
	}
}
```

Để kiểm tra index vừa được tạo:
```elasticsearch
GET _cat/indices
```

### Create
Elasticsearch hỗ trợ insert nhiều document một lúc vào index.
```elasticsearch
POST /companies/_bulk
{ "index": { "_id": 1 }}
{ "name": "Web Technology Company", "found_year": "2012-12-28", "avg_salary" : "5000$","benefits" : "snacks, learning budget, sport discount", "num_emp": 2000, "city": "San Francisco","country": "USA", "brief_info":"You can improve yourself in web technology by using django, Elasticsearch, and CI/CD tools. You can work remotely in San Francisco by having too much fun!" }
{ "index": { "_id": 2 }}
{ "name": "Defense Industry Company", "found_year": "2001-10-10", "avg_salary" : "4000$","benefits" : "rented-house, snacks, learning budget", "num_emp": 100, "city": "Munich","country": "Germany", "brief_info":"In our company we are working for defense industry. We are using c++. Codes are protected and you can only work in the office!"  }
{ "index": { "_id": 3 }}
{ "name": "Machine Learning Company", "found_year": "2011-02-10", "avg_salary" : "9000$","benefits" : "rented-house, snacks, learning budget, rented car", "num_emp": 500, "city": "Ankara","country": "Turkey", "brief_info":"Do you want to imporove yourself in machine learning and web technology? So this is right place to work. Come and enjoy your workdays. You can work remotely and enjoy the view of the capital city, Ankara!"  }
{ "index": { "_id": 4 }}
{ "name": "Embedded System Development Company", "found_year": "2006-03-03", "avg_salary" : "10000$","benefits" : "snacks, learning budget", "num_emp": 200, "city": "London","country": "England", "brief_info": "Are you interested in embedded system development? If you know or desire to learn c++, and improve yourself in this area apply now!"  }
{ "index": { "_id": 5 }}
{ "name": "Internet of Things Company", "found_year": "2016-02-13", "avg_salary" : "8000$","benefits" : "snacks", "num_emp": 20, "city": "Amsterdam","country": "Holland", "brief_info":"Researching for Internet of Things with flexible working hours! Come to Amsterdam office and have fun while learning!"   }
```

### READ
Elasticsearch hỗ trợ đa dạng các kiểu tìm kiếm khác nhau:
#### Tìm kiếm toàn bộ bản ghi
```elasticsearch
GET /companies/_search
{ "query": { "match_all": {}}}
```

#### Truy vấn Match cơ bản
Truy vấn này cho phép tìm kiếm 1 từ trong 1 field cụ thể. Ví dụ nếu muốn tìm từ "Web Technology" trong field  **name**, và giới hạn chỉ trả về các field là **name** và **brief_info**:
```elasticsearch
GET /companies/_search
{
  "query": {
    "match": {
      "name": "Web Technology"
    }
  },
  "_source": ["name", "brief_info"]
}
```
Lưu ý: Truy vấn trên sẽ tìm kiếm các bản ghi có từ "Web" hoặc từ "Technology".

Nếu muốn truy vấn **match** trên nhiều field (chỉ cần 1 field thỏa mãn), ta sử dụng **multi_match**:
```elasticsearch
GET /companies/_search
{
  "query": {
    "multi_match": {
      "query": "San Francisco sport",
      "fields": ["city", "benefits"]
    }
  },
  "_source": ["city", "benefits"]
}
```

Nếu muốn truy vấn **match** trên tất cả các field, bỏ qua trường **fields** trong **multi_match**:
```elasticsearch
GET /companies/_search
{
  "query": {
    "multi_match": {
      "query": "Web Technology"
    }
  }
}
```

#### Truy vấn Bool
Truy vấn bool dùng để viết các truy vấn AND/ OR/ NOT để kết hợp các ràng buộc.

Ví dụ, nếu muốn tìm kiếm các bản ghi với **benefits** là sport HOẶC rented house, và **city** KHÔNG phải là Ankara VÀ Amsterdam:
```elasticsearch
GET /companies/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "bool": {
            "should": [
              {
                "match": {
                  "benefits": "sport"
                }
              },
              {
                "match": {
                  "benefits": "rented house"
                }
              }
            ]
          }
        }
      ],
      "must_not": [
        {
          "bool": {
            "should": [
              {
                "match": {
                  "city": "Ankara"
                }
              },
              {
                "match": {
                  "city": "Amsterdam"
                }
              }
            ]
          }
        }
      ]
    }
  }
}
```
#### Truy vấn Fuzzy
Truy vấn fuzzy cho phép truy vấn 1 ràng buộc được viết không đúng chính tả. Trong truy vấn fuzzy có 1 chỉ số **fuzziness** là khoảng cách tối đa được chấp nhận giữa từ tìm được với từ trong truy vấn. Khoảng cách giữa 2 từ chính được tính là số thao tác để đưa từ thứ nhất về từ thứ hai chỉ qua 3 thao tác:
* Thêm 1 ký tự vào vị trí bất kỳ
* Xóa 1 ký tự bất kỳ
* Sửa 1 ký tự bất kỳ thành 1 ký tự khác.

```elasticsearch
GET /companies/_search
{
  "query": {
    "multi_match": {
      "query": "developmnt",
      "fields": ["name", "brief_info"],
      "fuzziness": "AUTO"
    }
  },
  "_source": ["name", "brief_info"]
}
```

#### Wildcard
Truy vấn wildcard cho phép tìm kiểm 1 pattern chứ không nhất thiết là 1 từ cụ thể. Trong truy vấn wildcard thì ? đại diện cho 1 ký tự, còn * đại diện cho 0 hoặc nhiều ký tự.

```elasticsearch
GET /companies/_search
{
  "query": {
    "wildcard": {
      "brief_info": {
        "value": "developm*"
      }
    }
  },
  "_source": ["name", "brief_info"]
}
```
#### Truy vấn match_phrase
Truy vấn match_phrase yêu cầu cụm từ được tìm kiếm phải xuất hiện theo đúng thứ tự đó.
```elasticsearch
GET /companies/_search
{
  "query": {
    "match_phrase": {
      "brief_info": "web technology"
    }
  },
  "_source": ["name", "brief_info"]
}
```
### Update
### Delete

## Aggregation
## Index/ Search Algorithm
### Inverted Index
Inverted index là một danh sách có thứ tự của toàn bộ các từ phân biệt xuất hiện trong bất kỳ 1 document nào của index. Inverted index được tạo ra bởi Analyzer gồm 3 bước:
* **Character Filter**: sử dụng để tiền xử lý dữ liệu bằng cách thêm, xóa, sửa ký tự.
* **Tokenizer**: Phân tách các document ra thành các từ riêng biệt.
* **Token Filters**: Sử dụng để thêm, xóa, sửa token. Ví dụ:
	* **lowercase**: Chuyển toàn bộ ký tự sang dạng chũ thường.
	* **stop**: lược bỏ các stop word.
	* **synonym**: chuyển đổi thành các từ đồng nghĩa.

Ví dụ, với 2 bản ghi:
* Doc1: ``` Show me the code ```
* Doc2: ``` Write the code and keep coding ```

Inverted index sẽ như sau:
```txt
code -> [Doc1, Doc2]
me -> [Doc1]
show -> [Doc1]
the -> [Doc1, Doc2]
write -> [Doc2]
keep -> [Doc2]
and -> [Doc2]
```
### Search Algorithm
Với mỗi truy vấn tìm kiếm, mỗi bản ghi hợp lệ sẽ được elasticsearch gán 1 chỉ số là **relevance_score**. Các bản ghi có **relevance_score** cao hơn sẽ đứng trước các bản ghi có chỉ số này thấp hơn.

Elasticsearch hỗ trợ nhiều thuật toán để tính điểm, trong đó nổi bật là 2 thuật toán sau.
#### TF/ IDF
Công thức của TF/ IDF như sau:

relavance_score = TF_score * IDF_score * fieldNorms

Trong đó:
* **term frequency**: Căn bậc 2 số lần xuất hiện của term trong field. Công thức: **tf(t, d) = sqrt(frequency)**.
* **inverse document frequency**: Nghịch đảo số lần xuất hiện của term trên toàn bộ index. Công thức: **idf(t) = 1 + log ( numDocs / (docFreq + 1) )**.
* **field-length norm**: Nghịch đảo căn bậc 2 của độ dài của field đó. Công thức: **norm(d) = 1 / sqrt(numTerms)**.

#### BM25
Đây là công thức mặc định của Elasticsearch dựa trên cải tiến cho TF/ IDF:
* **IDF**: idf(t) = log(1 + (docCount - docFreq + 0.5) / (docFreq + 0.5)) với:
	* **docCount**: số lượng document.
	* **docFreq**: số lượng document chứa term.
* **TF**: tf = ((k + 1) * freq) / (k * (1.0 - b + b * L) + freq) với:
	* **freq**: số lần xuất hiện của term trong document.
	* **k**: hằng số (mặc định là 1.2)
	* **b**: hằng số (mặc định là 0.75)
	* **L**: tỉ lệ giữa độ dài của document với độ dài trung bình của tất cả documents trong index.

