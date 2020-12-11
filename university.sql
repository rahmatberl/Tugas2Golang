-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Dec 11, 2020 at 07:33 PM
-- Server version: 10.4.14-MariaDB
-- PHP Version: 7.4.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `university`
--

-- --------------------------------------------------------

--
-- Table structure for table `mahasiswa`
--

CREATE TABLE `mahasiswa` (
  `id_mahasiswa` varchar(20) NOT NULL,
  `nama` varchar(50) DEFAULT NULL,
  `jalan` varchar(100) DEFAULT NULL,
  `kelurahan` varchar(50) DEFAULT NULL,
  `kecamatan` varchar(50) DEFAULT NULL,
  `kabupaten` varchar(50) DEFAULT NULL,
  `provinsi` varchar(50) DEFAULT NULL,
  `fakultas` varchar(30) DEFAULT NULL,
  `jurusan` varchar(30) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `mahasiswa`
--

INSERT INTO `mahasiswa` (`id_mahasiswa`, `nama`, `jalan`, `kelurahan`, `kecamatan`, `kabupaten`, `provinsi`, `fakultas`, `jurusan`) VALUES
('1811082010', 'Rahmat Berlianda Avandi', 'Jln. Pagengsaan Timur no. 56', 'Pitalah', 'Batipuh', 'Tanah Datar', 'Sumatera Barat', 'Teknologi Informasi', 'Software Engineer');

-- --------------------------------------------------------

--
-- Table structure for table `mata_kuliah`
--

CREATE TABLE `mata_kuliah` (
  `id_mata_kuliah` varchar(15) NOT NULL,
  `nama` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `mata_kuliah`
--

INSERT INTO `mata_kuliah` (`id_mata_kuliah`, `nama`) VALUES
('11111', 'Android'),
('22222', 'Sistem Operasi');

-- --------------------------------------------------------

--
-- Table structure for table `nilai`
--

CREATE TABLE `nilai` (
  `id_mahasiswa` varchar(20) NOT NULL,
  `id_mata_kuliah` varchar(15) DEFAULT NULL,
  `nilai` double DEFAULT NULL,
  `semester` int(10) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `nilai`
--

INSERT INTO `nilai` (`id_mahasiswa`, `id_mata_kuliah`, `nilai`, `semester`) VALUES
('1811082010', '11111', 9.3, 4);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `mahasiswa`
--
ALTER TABLE `mahasiswa`
  ADD PRIMARY KEY (`id_mahasiswa`);

--
-- Indexes for table `mata_kuliah`
--
ALTER TABLE `mata_kuliah`
  ADD PRIMARY KEY (`id_mata_kuliah`);

--
-- Indexes for table `nilai`
--
ALTER TABLE `nilai`
  ADD PRIMARY KEY (`id_mahasiswa`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
