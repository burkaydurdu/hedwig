# Hedwig
Hedwig is a quick email interface.

### API

```http
GET /hello/:to/:title/:message/:img?
```
| Parameter | Type     | Description                     |
|:----------|:---------|:--------------------------------|
| `to`      | `string` | **Required**. receiver email    |
| `title`   | `string` | **Required**. mail body title   |
| `message` | `string` | **Required**. mail body message |
| `img`     | `string` | **Optional**. mail body image   |

---

```http
POST /hello
```

**Request Body**
```json
{
  "to": "example@example.com",
  "title": "Title",
  "message": "hello my best friend",
  "image": "http:/xxxx.png"
}
```

### K8s
```shell
kubectl apply -f deployment.yaml
```
#### Note for K8s

- First of all, you should setup ``cert-manager``
```shell
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.7.2/cert-manager.yaml
```