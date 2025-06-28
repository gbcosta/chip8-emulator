# Versão mínima do CMake
cmake_minimum_required(VERSION 3.12)
# Habilitamos explicitamente o suporte para C e C++
project(build CXX C)

# Definimos o padrão do C++
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Encontramos os pacotes necessários
find_package(OpenGL REQUIRED)
find_package(glfw3 3.3 REQUIRED)

# --- Biblioteca Estática para o GLAD ---
# Em vez de compilar glad.c junto com tudo, criamos uma biblioteca separada para ele.
# Isso resolve problemas de linkagem entre C e C++.
add_library(glad_lib STATIC
    src/glad.c
)

# A biblioteca glad_lib precisa saber onde encontrar seu próprio arquivo de cabeçalho, glad.h.
# Usamos PUBLIC para que qualquer alvo que se vincule a glad_lib herde este caminho.
target_include_directories(glad_lib PUBLIC
    "${PROJECT_SOURCE_DIR}/third_party/glad"
)

add_executable(build
    main.cpp
    src/chip8.cpp
    src/opengl.cpp
)

target_link_libraries(build PRIVATE
    glad_lib
    glfw
    OpenGL::GL
)

set_target_properties(build PROPERTIES POSITION_INDEPENDENT_CODE ON)

# --- Mensagens de Status ---
message(STATUS "Usando bibliotecas OpenGL: ${OPENGL_LIBRARIES}")
message(STATUS "Usando GLFW através do alvo 'glfw'")
message(STATUS "Compilando e vinculando a biblioteca estática 'glad_lib'")
