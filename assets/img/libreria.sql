-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Servidor: 127.0.0.1
-- Tiempo de generación: 14-04-2022 a las 00:33:39
-- Versión del servidor: 10.4.22-MariaDB
-- Versión de PHP: 8.1.2

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de datos: `libreria`
--

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `areac`
--

CREATE TABLE `areac` (
  `id_al` int(10) NOT NULL,
  `nombrec` varchar(50) NOT NULL,
  `descripcion` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `areal`
--

CREATE TABLE `areal` (
  `id_al` int(10) NOT NULL,
  `isbn` varchar(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `autolib`
--

CREATE TABLE `autolib` (
  `nombre` varchar(10) DEFAULT NULL,
  `id_a` int(10) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `autores`
--

CREATE TABLE `autores` (
  `id_a` int(10) NOT NULL,
  `isbn` varchar(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Volcado de datos para la tabla `autores`
--

INSERT INTO `autores` (`id_a`, `isbn`) VALUES
(11, '132en48r'),
(13, '3ryhv7ry'),
(10, '783hyrd7'),
(12, '9ifh38ry'),
(14, 'd45rfe45');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `editorial`
--

CREATE TABLE `editorial` (
  `id_e` int(10) NOT NULL,
  `nombre_e` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Volcado de datos para la tabla `editorial`
--

INSERT INTO `editorial` (`id_e`, `nombre_e`) VALUES
(1, 'fantasia y cuentos'),
(2, 'serpensortia'),
(3, 'salamandra'),
(4, 'porrua'),
(5, 'gandhi');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `libros`
--

CREATE TABLE `libros` (
  `isbn` varchar(10) NOT NULL,
  `titulo` varchar(100) NOT NULL,
  `edicion` int(10) NOT NULL,
  `id_e` int(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Volcado de datos para la tabla `libros`
--

INSERT INTO `libros` (`isbn`, `titulo`, `edicion`, `id_e`) VALUES
('132en48r', 'Las dos muertes de Lina Posada', 20, 3),
('3ryhv7ry', 'Hermoso Desastre', 32, 5),
('783hyrd7', 'Orgullo y Prejuicio', 11, 2),
('9ifh38ry', 'El conde de Montecristo', 45, 4),
('d45rfe45', 'Harry Potter', 4, 1);

--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `areac`
--
ALTER TABLE `areac`
  ADD KEY `id_al` (`id_al`);

--
-- Indices de la tabla `areal`
--
ALTER TABLE `areal`
  ADD PRIMARY KEY (`id_al`),
  ADD KEY `isbn` (`isbn`);

--
-- Indices de la tabla `autolib`
--
ALTER TABLE `autolib`
  ADD KEY `isbn` (`nombre`),
  ADD KEY `id_a` (`id_a`);

--
-- Indices de la tabla `autores`
--
ALTER TABLE `autores`
  ADD PRIMARY KEY (`id_a`),
  ADD KEY `isbn` (`isbn`);

--
-- Indices de la tabla `editorial`
--
ALTER TABLE `editorial`
  ADD PRIMARY KEY (`id_e`);

--
-- Indices de la tabla `libros`
--
ALTER TABLE `libros`
  ADD PRIMARY KEY (`isbn`),
  ADD KEY `id_e` (`id_e`);

--
-- Restricciones para tablas volcadas
--

--
-- Filtros para la tabla `areac`
--
ALTER TABLE `areac`
  ADD CONSTRAINT `areac_ibfk_1` FOREIGN KEY (`id_al`) REFERENCES `areal` (`id_al`);

--
-- Filtros para la tabla `areal`
--
ALTER TABLE `areal`
  ADD CONSTRAINT `areal_ibfk_1` FOREIGN KEY (`isbn`) REFERENCES `libros` (`isbn`);

--
-- Filtros para la tabla `autolib`
--
ALTER TABLE `autolib`
  ADD CONSTRAINT `autolib_ibfk_1` FOREIGN KEY (`nombre`) REFERENCES `libros` (`isbn`),
  ADD CONSTRAINT `autolib_ibfk_2` FOREIGN KEY (`id_a`) REFERENCES `autores` (`id_a`);

--
-- Filtros para la tabla `autores`
--
ALTER TABLE `autores`
  ADD CONSTRAINT `autores_ibfk_1` FOREIGN KEY (`isbn`) REFERENCES `libros` (`isbn`);

--
-- Filtros para la tabla `libros`
--
ALTER TABLE `libros`
  ADD CONSTRAINT `libros_ibfk_1` FOREIGN KEY (`id_e`) REFERENCES `editorial` (`id_e`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
