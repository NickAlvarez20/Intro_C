FROM gcc:14-bookworm AS build
WORKDIR /app
COPY copy_input_to_output.c .
RUN gcc -Wall -Wextra -O2 -o copy_input_to_output copy_input_to_output.c

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=build /app/copy_input_to_output .
ENTRYPOINT ["./copy_input_to_output"]
