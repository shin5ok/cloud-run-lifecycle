# �x�[�X�C���[�W���w��
FROM golang:latest

# ��ƃf�B���N�g����ݒ�
WORKDIR /app

# �ˑ��֌W���R�s�[���A�C���X�g�[��
COPY go.mod go.sum ./
RUN go mod download

# �\�[�X�R�[�h���R�s�[
COPY . .

# �A�v���P�[�V�������r���h
RUN go build -o main .

# ���s����R�}���h���w��
CMD ["./main"]

