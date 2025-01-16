# YouTube Learning Mode Quiz Service

This Go-based microservice provides AI-generated quizes using AI service. It generates questions with text, options, explanation for each option, and timestamps for each question.

## Features

- **AI Session Initialization**: Having the AI service generate quizes of fixed format and parse into json for easy rendering in the browser extention.
- **Redis Integration**: Caches AI-generated quizes to load from db when same video is loaded.
- **Future Enhancements**: Potentially able to generate short answer questions in the future.

## Tech Stack

- **Go**: Core language for the service.
- **AI service API**: For generating quizes.
- **Redis**: For storing conversation history.
- **Gorilla Mux**: Router for handling HTTP requests.

## Prerequisites

Make sure you have the following installed:

- **Go**: Version 1.16 or higher.
- **Docker & Docker Compose**: For containerization.
- **AI service**: Required for generate quizes.

## Installation

### 1. Clone the Repository

```bash
git clone https://github.com/Learning-Mode-AI/Learning-Mode-AI-quiz-service.git
cd Learning-Mode-AI-quiz-service
```

### 2. Set Up Environment Variables

Create a `.env` file in the root directory to store your environment variables:

```bash
QUIZ_SERVICE_PORT=8084
ENVIRONMENT=docker
```

> **Important:** Do not commit the `.env` file to version control as it contains sensitive information.

### 3. Run with Docker

To build and run the service using Docker:

```bash
sudo docker build -t quiz-service:latest .
sudo docker run -p 8084:8084 quiz-service:latest
```

Or use Docker Compose with other services:

```bash
docker-compose up --build
```

This will start the Quiz Service along with the other microservices in a shared network.

## Running the Service Without Docker (Optional)

If you want to run the service without Docker, make sure Redis is running locally and your `.env` file is properly set up:

```bash
cd cmd
go run main.go
```

The server will start on `http://localhost:8084`.

## API Endpoints

### 1. Initialize AI Session

- **Endpoint**: `POST /quiz/generate-quiz`
- **Description**: Create 10 multiple choice questions.
- **Request Body**:

  ```json
  {
    "video_id": "VIDEO_ID"
  }
  ```

- **Response**:

  ```json
  {
    "quiz_id": "6_2hzRopPbQ",
    "questions": [
        {
            "text": "What is Tensorflow?",
            "options": [
                {
                    "option": "a flexible and open source library for building deep learning models",
                    "explanation": "Tensorflow is used to simplify the process of building deep learning models."
                },
                {
                    "option": "a tool for data analysis and visualization",
                    "explanation": "Tensorflow is primarily used for data analysis and visualization."
                },
                {
                    "option": "a programming language for AI development",
                    "explanation": "Tensorflow is a programming language specifically for AI development."
                },
                {
                    "option": "a cloud computing platform",
                    "explanation": "Tensorflow is a cloud computing platform."
                }
            ],
            "answer": "a flexible and open source library for building deep learning models",
            "timestamp": "29.519"
        },
        {
            "text": "What is the primary objective of the Tensorflow model demonstrated in the video?",
            "options": [
                {
                    "option": "To predict customer churn",
                    "explanation": "The goal of the example model is to predict whether a customer will churn or not."
                },
                {
                    "option": "To predict sales figures",
                    "explanation": "The model is used for predicting sales figures."
                },
                {
                    "option": "To predict image categories",
                    "explanation": "The model is used for image recognition tasks."
                },
                {
                    "option": "To predict sentiment in text",
                    "explanation": "The model is used for natural language processing."
                }
            ],
            "answer": "To predict customer churn",
            "timestamp": "121.04"
        },
        {
            "text": "What is the first coding task that the presenter mentions?",
            "options": [
                {
                    "option": "Importing Tensorflow into a Notebook",
                    "explanation": "This step is crucial to use Tensorflow functions in the code."
                },
                {
                    "option": "Building a deep neural network",
                    "explanation": "This step is about defining the model architecture."
                },
                {
                    "option": "Saving the model to disk",
                    "explanation": "This step is about saving the model after training."
                },
                {
                    "option": "Calculating accuracy from predictions",
                    "explanation": "This step is about evaluating the model's performance."
                }
            ],
            "answer": "Importing Tensorflow into a Notebook",
            "timestamp": "2.47"
        },
        {
            "text": "What type of model class is used to build the neural network?",
            "options": [
                {
                    "option": "Sequential model",
                    "explanation": "The Sequential model is the core class used to build models in Tensorflow."
                },
                {
                    "option": "Functional model",
                    "explanation": "The Functional model is used for handling complex architectures."
                },
                {
                    "option": "Model class",
                    "explanation": "The Model class is for loading pre-trained models only."
                },
                {
                    "option": "Keras model",
                    "explanation": "The Keras model is an older version of Tensorflow's model."
                }
            ],
            "answer": "Sequential model",
            "timestamp": "185.519"
        },
        {
            "text": "Which loss function is used when compiling the model?",
            "options": [
                {
                    "option": "Binary_crossentropy",
                    "explanation": "Binary crossentropy is a common loss function used for binary classification problems."
                },
                {
                    "option": "Mean squared error",
                    "explanation": "Mean squared error is used for regression tasks."
                },
                {
                    "option": "Categorical crossentropy",
                    "explanation": "Categorical crossentropy is used for multi-class classification tasks."
                },
                {
                    "option": "Hinge loss",
                    "explanation": "Hinge loss is used for support vector machines."
                }
            ],
            "answer": "Binary_crossentropy",
            "timestamp": "396.72"
        },
        {
            "text": "What does the 'fit' method do in the context of Tensorflow?",
            "options": [
                {
                    "option": "Training the model",
                    "explanation": "This step involves fitting the model to the training data."
                },
                {
                    "option": "Saving the model",
                    "explanation": "This step is about saving the model after training."
                },
                {
                    "option": "Loading the model",
                    "explanation": "This step is about loading the model from disk."
                },
                {
                    "option": "Evaluating the model",
                    "explanation": "This step is about evaluating the model's accuracy."
                }
            ],
            "answer": "Training the model",
            "timestamp": "444.479"
        },
        {
            "text": "What threshold value is used to classify the predictions into binary outcomes?",
            "options": [
                {
                    "option": "0.5",
                    "explanation": "Values above 0.5 are classified as 1 (churn) and below as 0 (no churn)."
                },
                {
                    "option": "1",
                    "explanation": "Values above 1 indicate a positive prediction."
                },
                {
                    "option": "0",
                    "explanation": "Values below 0 indicate no prediction."
                },
                {
                    "option": "0",
                    "explanation": "Values of 0 indicate a neutral prediction."
                }
            ],
            "answer": "0.5",
            "timestamp": "528.08"
        },
        {
            "text": "What is the last coding task that the presenter demonstrates?",
            "options": [
                {
                    "option": "Saving the model to disk",
                    "explanation": "This step allows the model to be reused later without retraining."
                },
                {
                    "option": "Preparing the model for deployment",
                    "explanation": "This step is about preparing the model for deployment."
                },
                {
                    "option": "Generating visualizations",
                    "explanation": "This step generates visualizations of the model performance."
                },
                {
                    "option": "Optimizing the model",
                    "explanation": "This step optimizes the model for better performance."
                }
            ],
            "answer": "Saving the model to disk",
            "timestamp": "586.16"
        },
        {
            "text": "What is demonstrated after saving the model?",
            "options": [
                {
                    "option": "Loading Tensorflow Models",
                    "explanation": "The load_model function allows previously saved models to be reused."
                },
                {
                    "option": "Saving Tensorflow Models",
                    "explanation": "The save_model function is for saving new models only."
                },
                {
                    "option": "Training Tensorflow Models",
                    "explanation": "The fit method is for training the model only."
                },
                {
                    "option": "Evaluating Tensorflow Models",
                    "explanation": "The evaluate method is for assessing model performance only."
                }
            ],
            "answer": "Loading Tensorflow Models",
            "timestamp": "614.0"
        },
        {
            "text": "What metric does the presenter use to evaluate the model's performance?",
            "options": [
                {
                    "option": "Accuracy Score",
                    "explanation": "Accuracy score is used to determine how well the model is performing."
                },
                {
                    "option": "Loss Function",
                    "explanation": "Loss function helps to quantify the model's errors."
                },
                {
                    "option": "Precision",
                    "explanation": "Precision measures the model's relevancy in positive predictions."
                },
                {
                    "option": "Recall",
                    "explanation": "Recall measures the model's ability to find all positive instances."
                }
            ],
            "answer": "Accuracy Score",
            "timestamp": "571.6"
        }
    ]
    }
  ```



## Project Structure

```
├── cmd/
│   └── main.go                 # Entry point of the AI service
├── pkg/
│   └── config/
│       └── conifg.go           # Defines local or docker
│   ├── handlers/
│   │   └── quizhandlers.go         # Handles error logging
│   ├── services/
│   │   ├── quizservice.go      # Interacts with AI service to generate question and parse into json
│   │   ├── redis.go    # Handles Redis connections and operations
│   └── router/
│       └── router.go           # Defines API routes (if applicable)
├── .env                        # Environment variables (DO NOT COMMIT)
├── dockerfile                  # Docker configuration for containerizing the service
├── go.mod                      # Go module dependencies
├── go.sum                      # Checksums for Go modules
└── README.md                   # Project documentation
```

## Dependencies

- **Go Redis**: For connecting to the Redis database.
- **OpenAI Go SDK**: For communicating with the OpenAI GPT-4 API.
- **Gorilla Mux**: For HTTP request routing.

## Future Enhancements

- **More question types**: To provide more question types to improve engagement such as short answer question.
- **Custom Models**: Incorporate additional AI models or fine-tuned versions for more specific responses.

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.

2. Create a new branch:

   ```bash
   git checkout -b feature/YourFeature
   ```

3. Commit your changes:

   ```bash
   git commit -am 'Add some feature'
   ```

4. Push to the branch:

   ```bash
   git push origin feature/YourFeature
   ```

5. Open a Pull Request.

---

### Important Notes:

- Ensure Redis is off to run docker.
