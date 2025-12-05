# This file was written by Damien Dart, <damiendart@pobox.com>. This is
# free and unencumbered software released into the public domain. For
# more information, please refer to the accompanying "UNLICENCE" file.

(def input-parser (peg/compile ~{:main (split :s :line)
                                 :line (any (+
                                              (group (* (number :d+) "-" (number :d+)))
                                              (number :d+)))}))

(defn solve [input]
  (def parsed (peg/match input-parser input))
  (let [ranges (filter array? parsed)
        ids (filter number? parsed)]
    (count (fn [id] (find |(and (>= id (get $ 0)) (<= id (get $ 1))) ranges)) ids)))

(defn main [&] (->> (file/read stdin :all) string/trim solve pp))
