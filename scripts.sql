-- Tabla para almacenar información de EXPEDIENTES JUDICIALES
CREATE TABLE judicial_files (
    id SERIAL PRIMARY KEY,
    file_number VARCHAR(100) NOT NULL,          -- N° de expediente
    notification_number VARCHAR(100),           -- N° de notificación
    digitization_number VARCHAR(100),           -- N° de digitalización
    document_type VARCHAR(50),                  -- Ej: Resolución, Notificación
    headquarters VARCHAR(100),                  -- Ej: Villa El Salvador
    court VARCHAR(150),                         -- Ej: Juzgado Investigación Preparatoria
    notification_date TIMESTAMPTZ,              -- Fecha de notificación
    creation_date TIMESTAMPTZ DEFAULT NOW(),    -- Fecha de creación del registro
    update_date TIMESTAMPTZ DEFAULT NOW(),      -- Fecha de última actualización del registro
    court_id BIGINT NOT NULL,                   -- fk a la tabla courts (cortes judiciales)
    CONSTRAINT fk_court FOREIGN KEY (court_id) REFERENCES courts(id)
);

-- Tabla para almacenar información de CORTES JUDICIALES
CREATE TABLE courts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,                 -- Nombre  Ej: Corte Superior de Justicia Lima Sur
    headquarters VARCHAR(100)                   -- Sede   Ej: Villa El Salvador
);

-- Tabla para almacenar información de JUECES
CREATE TABLE judges (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(200) NOT NULL,
    specialty VARCHAR(100),                     -- Ej: Juzgado de Investigación Preparatoria
    court_id BIGINT NOT NULL,
    CONSTRAINT fk_judge_court FOREIGN KEY (court_id) REFERENCES courts(id)
);

-- Tabla para almacenar información de personas relacionadas con los expedientes
CREATE TABLE persons (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(200) NOT NULL,
    role VARCHAR(50) NOT NULL,                  -- Ej: IMPUTADO, AGRAVIADO, DESTINATARIO
    email VARCHAR(150),                         -- Dirección electrónica
    phone VARCHAR(20)
);

-- Tabla intermedia para relacionar expedientes con personas
CREATE TABLE judicial_file_persons (
    id SERIAL PRIMARY KEY,
    judicial_file_id BIGINT NOT NULL,
    person_id BIGINT NOT NULL,
    CONSTRAINT fk_jf FOREIGN KEY (judicial_file_id) REFERENCES judicial_files(id),
    CONSTRAINT fk_person FOREIGN KEY (person_id) REFERENCES persons(id)
);

-- Tabla para almacenar información de abogados
CREATE TABLE lawyers (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(200) NOT NULL,
    bar_number VARCHAR(50),                     -- Número de colegiatura
    email VARCHAR(150)
);

-- Tabla intermedia para relacionar expedientes con abogados
CREATE TABLE judicial_file_lawyers (
    id SERIAL PRIMARY KEY,
    judicial_file_id BIGINT NOT NULL,
    lawyer_id BIGINT NOT NULL,
    CONSTRAINT fk_jfl FOREIGN KEY (judicial_file_id) REFERENCES judicial_files(id),
    CONSTRAINT fk_lawyer FOREIGN KEY (lawyer_id) REFERENCES lawyers(id)
);