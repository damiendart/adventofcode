# This file was written by Damien Dart, <damiendart@pobox.com>. This is
# free and unencumbered software released into the public domain. For
# more information, please refer to the accompanying "UNLICENCE" file.

(def input-parser (peg/compile ~{:main (split "\n" (group (* (any (<- (set ".S^"))))))}))

(defn solve [input]
  (def parsed (peg/match input-parser input))
  (var lasers (array/new-filled (length (get parsed 0)) 0))

  (each line parsed
    (eachp [idx column] line
      (case column
        "S" (put lasers idx 1)
        "^" (do
              (if (pos? (get lasers idx)) (do
                                            (def v (get lasers idx))
                                            (put lasers idx 0)
                                            (put lasers (- idx 1) (+ (get lasers (- idx 1)) v))
                                            (put lasers (+ idx 1) (+ (get lasers (+ idx 1)) v))))))))

  (sum lasers))

(defn main [&] (->> (file/read stdin :all) string/trim solve pp))
