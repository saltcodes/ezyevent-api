# EzyEvent Main API
Main API for CRUD Operations.

## Version: 1.0

**Contact information:**  
API Support  
oseiyeboahjohnson@gmail.com

**License:** [Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0.html)

### /bookings

#### GET
##### Summary:

get all existing bookings

##### Description:

Get List of all bookings.

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

### /bookings/{id}

#### DELETE
##### Summary:

delete booking by id

##### Description:

delete bookings in the database by checking with id

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Booking ID | Yes | string |
| request | body | query params | Yes | [ezyevent-api_internal_model.Booking](#ezyevent-api_internal_model.Booking) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

#### GET
##### Summary:

get booking by given ID

##### Description:

Get Booking by given ID.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Booking ID | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

#### POST
##### Summary:

create a booking

##### Description:

create booking in the database

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Booking ID | Yes | string |
| request | body | query params | Yes | [ezyevent-api_internal_model.Booking](#ezyevent-api_internal_model.Booking) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

#### PUT
##### Summary:

update booking by id

##### Description:

updates bookings in the database by checking with id

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Booking ID | Yes | string |
| request | body | query params | Yes | [ezyevent-api_internal_model.Booking](#ezyevent-api_internal_model.Booking) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

### /categories

#### GET
##### Summary:

get all existing categories

##### Description:

Get List of all categories.

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

### /categories/{id}

#### POST
##### Summary:

create a category

##### Description:

create category in the database

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Category ID | Yes | string |
| request | body | query params | Yes | [ezyevent-api_internal_model.Category](#ezyevent-api_internal_model.Category) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

### /comments

#### GET
##### Summary:

get all existing comments

##### Description:

Get List of all comments.

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

### /events

#### GET
##### Summary:

get all existing events

##### Description:

Get List of all events.

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

### /events/{id}

#### DELETE
##### Summary:

delete event by id

##### Description:

delete events in the database by checking with id

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Event ID | Yes | string |
| request | body | query params | Yes | [ezyevent-api_internal_model.Event](#ezyevent-api_internal_model.Event) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

#### GET
##### Summary:

get event by given ID

##### Description:

Get Event by given ID.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Event ID | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

#### POST
##### Summary:

create an event

##### Description:

create event in the database

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Event ID | Yes | string |
| request | body | query params | Yes | [ezyevent-api_internal_model.Event](#ezyevent-api_internal_model.Event) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

#### PUT
##### Summary:

update event by id

##### Description:

updates events in the database by checking with id

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Event ID | Yes | string |
| request | body | query params | Yes | [ezyevent-api_internal_model.Event](#ezyevent-api_internal_model.Event) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

### /events/find

#### GET
##### Summary:

query events bases and location

##### Description:

find events using lng,lat and radius

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| radius | query | enter radius for event | No | string |
| lat | query | enter starting latitude | No | string |
| lng | query | enter starting  longitude | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

### /organizations

#### GET
##### Summary:

get all existing organizations

##### Description:

Get List of all organizations.

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

### /organizations/{id}

#### DELETE
##### Summary:

delete organization by id

##### Description:

delete organizations in the database by checking with id

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Organization ID | Yes | string |
| request | body | query params | Yes | [ezyevent-api_internal_model.Organization](#ezyevent-api_internal_model.Organization) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

#### GET
##### Summary:

get organization by given ID

##### Description:

Get Organization by given ID.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Organization ID | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

#### POST
##### Summary:

create an organization

##### Description:

create organization in the database

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Organization ID | Yes | string |
| request | body | query params | Yes | [ezyevent-api_internal_model.Organization](#ezyevent-api_internal_model.Organization) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

#### PUT
##### Summary:

update organization by id

##### Description:

updates organizations in the database by checking with id

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Organization ID | Yes | string |
| request | body | query params | Yes | [ezyevent-api_internal_model.Organization](#ezyevent-api_internal_model.Organization) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

### /schedules

#### GET
##### Summary:

get all existing schedules

##### Description:

Get List of all schedules.

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

### /schedules/{id}

#### DELETE
##### Summary:

delete schedule by id

##### Description:

delete schedules in the database by checking with id

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Schedule ID | Yes | string |
| request | body | query params | Yes | [ezyevent-api_internal_model.Schedule](#ezyevent-api_internal_model.Schedule) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

#### GET
##### Summary:

get schedule by given ID

##### Description:

Get Schedule by given ID.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Schedule ID | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

#### POST
##### Summary:

create a schedule

##### Description:

create schedule in the database

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Schedule ID | Yes | string |
| request | body | query params | Yes | [ezyevent-api_internal_model.Schedule](#ezyevent-api_internal_model.Schedule) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

#### PUT
##### Summary:

update schedule by id

##### Description:

updates schedules in the database by checking with id

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Schedule ID | Yes | string |
| request | body | query params | Yes | [ezyevent-api_internal_model.Schedule](#ezyevent-api_internal_model.Schedule) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

### /users

#### GET
##### Summary:

get all existing users

##### Description:

Get List of all users.

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

### /users/{id}

#### DELETE
##### Summary:

delete user by id

##### Description:

delete users in the database by checking with id

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | User ID | Yes | string |
| request | body | query params | Yes | [ezyevent-api_internal_model.User](#ezyevent-api_internal_model.User) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

#### GET
##### Summary:

get user by given ID

##### Description:

Get User by given ID.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | User ID | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

#### POST
##### Summary:

create a user

##### Description:

create user in the database

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | User ID | Yes | string |
| request | body | query params | Yes | [ezyevent-api_internal_model.User](#ezyevent-api_internal_model.User) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

#### PUT
##### Summary:

update user by id

##### Description:

updates users in the database by checking with id

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | User ID | Yes | string |
| request | body | query params | Yes | [ezyevent-api_internal_model.User](#ezyevent-api_internal_model.User) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | desc | [ezyevent-api_internal_model.ResponseObject](#ezyevent-api_internal_model.ResponseObject) & object |

### Models


#### ezyevent-api_internal_model.Booking

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| ID | string |  | No |
| createdAt | string |  | No |
| schedule | [ezyevent-api_internal_model.Schedule](#ezyevent-api_internal_model.Schedule) |  | No |
| scheduleId | string |  | No |
| updatedAt | string |  | No |
| user | [ezyevent-api_internal_model.User](#ezyevent-api_internal_model.User) |  | No |
| userId | string |  | No |

#### ezyevent-api_internal_model.Category

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| iconUrl | string |  | No |
| id | string |  | No |
| name | string |  | No |

#### ezyevent-api_internal_model.Comment

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| createdAt | integer |  | No |
| event_id | string |  | No |
| id | string |  | No |
| message | string |  | No |
| updatedAt | integer | Use unix nano seconds as updating time | No |
| user | [ezyevent-api_internal_model.User](#ezyevent-api_internal_model.User) |  | No |
| userId | string |  | No |

#### ezyevent-api_internal_model.Event

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| banner | string |  | No |
| category | [ezyevent-api_internal_model.Category](#ezyevent-api_internal_model.Category) |  | No |
| categoryID | string |  | No |
| date | integer |  | No |
| details | string |  | No |
| id | string |  | No |
| images | [ string ] |  | No |
| lat | number |  | No |
| lng | number |  | No |
| locationName | string |  | No |
| name | string |  | No |
| orgId | string |  | No |
| organization | [ezyevent-api_internal_model.Organization](#ezyevent-api_internal_model.Organization) |  | No |
| price | number |  | No |
| summary | string |  | No |

#### ezyevent-api_internal_model.Organization

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| bannerUrl | string |  | No |
| createdAt | string |  | No |
| id | string |  | No |
| name | string |  | No |
| orgImageUrl | string |  | No |
| updatedAt | string |  | No |
| user | [ezyevent-api_internal_model.User](#ezyevent-api_internal_model.User) |  | No |
| userId | string |  | No |

#### ezyevent-api_internal_model.ResponseObject

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| code | integer |  | No |
| data |  |  | No |
| message | string |  | No |

#### ezyevent-api_internal_model.Schedule

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| createdAt | string |  | No |
| date | string |  | No |
| event | [ezyevent-api_internal_model.Event](#ezyevent-api_internal_model.Event) |  | No |
| eventId | string |  | No |
| id | string |  | No |
| orgId | string |  | No |
| updatedAt | string |  | No |

#### ezyevent-api_internal_model.User

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| dob | integer |  | No |
| email | string |  | No |
| firstName | string |  | No |
| id | string |  | No |
| lastName | string |  | No |
| phone | string |  | No |
| profile_url | string |  | No |