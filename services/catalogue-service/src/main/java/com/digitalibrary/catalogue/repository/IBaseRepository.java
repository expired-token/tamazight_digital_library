package com.digitalibrary.catalogue.repository;

import com.digitalibrary.catalogue.entity.BaseEntity;
import org.jspecify.annotations.NonNull;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.repository.NoRepositoryBean;

import java.util.List;
import java.util.Optional;

@NoRepositoryBean
public interface IBaseRepository<T extends BaseEntity> extends JpaRepository<T, Long> {
    Optional<T> findById(@NonNull Long id);
    boolean existsById(@NonNull Long id);
    Optional<T> findByCode(@NonNull String code);
    List<T> findAll();
}
