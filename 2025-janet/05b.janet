# This file was written by Damien Dart, <damiendart@pobox.com>. This is
# free and unencumbered software released into the public domain. For
# more information, please refer to the accompanying "UNLICENCE" file.

(def input-parser (peg/compile ~{:main (split :s :line)
                                 :line (any (+
                                              (group (* (number :d+) "-" (number :d+)))
                                              (number :d+)))}))

(defn solve [input]
  (def parsed (peg/match input-parser input))
  (var upper 0)
  (let [ranges (sort-by |(get $ 0) (filter array? parsed))]
    (reduce
      (fn [carry range]
        (def difference (- (get range 1) (if (> upper (get range 0)) upper (get range 0)) -1))
        (set upper (+ (if (> upper (get range 1)) upper (get range 1)) (if (> difference 0) 1 0)))
        (if (> difference 0) (+ carry difference) carry))
      0
      ranges)))

(defn main [&] (->> (file/read stdin :all) string/trim solve pp))
